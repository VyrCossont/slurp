// GoToSocial
// Copyright (C) Vyr Cossont
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

//go:generate go run github.com/go-swagger/go-swagger/cmd/swagger generate client --spec https://raw.githubusercontent.com/superseriousbusiness/gotosocial/main/docs/api/swagger.yaml

import (
	"log/slog"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/peterhellberg/link"
	"github.com/pkg/errors"

	apiclient "github.com/VyrCossont/slurp/client"
	"github.com/VyrCossont/slurp/client/search"
	"github.com/VyrCossont/slurp/client/timelines"
	"github.com/VyrCossont/slurp/models"
)

type config struct {
	debug       bool
	srcInstance string
	srcToken    string
	dstInstance string
	dstToken    string
	count       int
	pageSize    int
}

// newConfigFromEnv captures important environment variables.
// If one or more are wrong, it returns nil, and the program should exit.
func newConfigFromEnv() *config {
	// If this is set, we're missing something important.
	fatal := false

	// Defaults.
	c := config{
		count:    100,
		pageSize: 40,
	}

	debug := os.Getenv("DEBUG")
	if debug != "" && debug != "true" && debug != "false" {
		fatal = true
		slog.Error("the DEBUG environment variable, if set, must be a boolean")
	}
	c.debug = debug == "true"

	c.srcInstance = os.Getenv("SRC_INSTANCE")
	if c.srcInstance == "" {
		fatal = true
		slog.Error("the SRC_INSTANCE environment variable must be set to the domain of your source instance")
	}

	// This can be empty: some instances still allow unauthenticated access.
	c.srcToken = os.Getenv("SRC_TOKEN")

	c.dstInstance = os.Getenv("DST_INSTANCE")
	if c.dstInstance == "" {
		fatal = true
		slog.Error("the DST_INSTANCE environment variable must be set to the domain of your destination instance")
	}

	c.dstToken = os.Getenv("DST_TOKEN")
	if c.dstToken == "" {
		fatal = true
		slog.Error("the DST_TOKEN environment variable must be set to an authentication token for your destination instance")
	}

	count := os.Getenv("COUNT")
	if count != "" {
		count, err := strconv.Atoi(count)
		msg := "the COUNT environment variable, if set, must be a positive integer"
		if err != nil {
			fatal = true
			slog.Error(msg, "error", err)
		} else if count <= 0 {
			fatal = true
			slog.Error(msg, "value", count)
		} else {
			c.count = count
		}
	}

	pageSize := os.Getenv("PAGE_SIZE")
	if pageSize != "" {
		pageSize, err := strconv.Atoi(pageSize)
		msg := "the PAGE_SIZE environment variable, if set, must be a positive integer"
		if err != nil {
			fatal = true
			slog.Error(msg, "error", err)
		} else if pageSize <= 0 {
			fatal = true
			slog.Error(msg, "value", pageSize)
		} else {
			c.pageSize = pageSize
		}
	}

	if fatal {
		return nil
	}

	return &c
}

func (c *config) srcClient() *apiclient.GoToSocialSwaggerDocumentation {
	return apiclient.New(httptransport.New(c.srcInstance, "", []string{"https"}), strfmt.Default)
}

func (c *config) srcAuth() runtime.ClientAuthInfoWriter {
	if c.srcToken == "" {
		var noopAuth runtime.ClientAuthInfoWriterFunc = func(request runtime.ClientRequest, registry strfmt.Registry) error {
			return nil
		}
		return noopAuth
	}
	return httptransport.BearerToken(c.srcToken)
}

func (c *config) dstClient() *apiclient.GoToSocialSwaggerDocumentation {
	return apiclient.New(httptransport.New(c.dstInstance, "", []string{"https"}), strfmt.Default)
}

func (c *config) dstAuth() runtime.ClientAuthInfoWriter {
	return httptransport.BearerToken(c.dstToken)
}

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, nil)))

	c := newConfigFromEnv()
	if c == nil {
		os.Exit(1)
	}
	if c.debug {
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		})))
	}

	urlsToResolve := make(chan string, c.count)
	statusesRequested := make(chan int, c.count)

	srcDone := make(chan error, 1)
	go srcTask(c, urlsToResolve, statusesRequested, srcDone)

	dstDone := make(chan error, 1)
	go dstTask(c, urlsToResolve, statusesRequested, dstDone)

	srcError := <-srcDone
	dstError := <-dstDone
	if srcError != nil || dstError != nil {
		os.Exit(1)
	}
}

// srcTask uses the public timeline API to discover statuses.
func srcTask(c *config, urlsToResolve chan string, statusesRequested chan int, done chan error) {
	slog.Debug("starting…")

	var err error
	client := c.srcClient()
	auth := c.srcAuth()
	limit := int64(c.pageSize)
	var maxID *string
	var statuses []*models.Status

loop:
	for {
		numRequested, open := <-statusesRequested
		if !open {
			slog.Debug("destination shut down")
			break
		}

		for numRequested > 0 {
			// Send all the statuses requested, or as many as we have.
			split := min(numRequested, len(statuses))
			for _, status := range statuses[:split] {
				urlsToResolve <- status.URI
			}
			statuses = statuses[split:]
			if split > 0 {
				slog.Debug("sent", "count", split)
			}
			numRequested -= split
			if numRequested <= 0 {
				break
			}

			// Get some more.
			var resp *timelines.PublicTimelineOK
			resp, err = client.Timelines.PublicTimeline(&timelines.PublicTimelineParams{
				Limit:   &limit,
				Local:   nil,
				MaxID:   maxID,
				MinID:   nil,
				SinceID: nil,
			}, auth)
			if err != nil {
				slog.Error("timeline error", "error", err)
				break loop
			}
			slog.Debug("fetched statuses", "count", len(resp.Payload))

			maxID, err = parseLink(resp.Link)
			if maxID == nil || err != nil {
				slog.Error("paging error", "error", err)
				break loop
			}

			for _, status := range resp.Payload {
				if skip(status) {
					continue
				}
				statuses = append(statuses, status)
			}
		}
	}

	close(urlsToResolve)
	done <- nil
	slog.Debug("finished")
}

// parseLink extracts the `max_id` from the `next` link for paging to older statuses.
func parseLink(linkHeader string) (*string, error) {
	next := link.Parse(linkHeader)["next"]
	if next == nil {
		// No link header in that direction means end of results.
		return nil, nil
	}
	nextUrl, err := url.Parse(next.URI)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't parse next page URL")
	}
	nextMaxID := nextUrl.Query().Get("max_id")
	if nextMaxID == "" {
		return nil, errors.New("couldn't find next page max ID")
	}
	return &nextMaxID, err
}

// skip returns true if we should ignore a status.
func skip(status *models.Status) bool {
	if status.Visibility != "public" {
		return true
	}
	if status.Sensitive {
		return true
	}
	if !strings.HasPrefix(status.Language, "en") {
		return true
	}
	if len(status.MediaAttachments) > 0 {
		return true
	}
	if status.Account.Bot {
		return true
	}

	return false
}

// dstTask uses the search API to resolve and import statuses.
func dstTask(c *config, urlsToResolve chan string, statusesRequested chan int, done chan error) {
	slog.Debug("starting…")

	var err error
	client := c.dstClient()
	auth := c.dstAuth()
	remaining := c.count
	var limit int64 = 1
	resolve := true
	type_ := "statuses"

	statusesRequested <- remaining

	for remaining > 0 {
		remoteURL, open := <-urlsToResolve
		if !open {
			slog.Debug("source shut down")
			break
		}

		var resp *search.SearchGetOK
		resp, err = client.Search.SearchGet(&search.SearchGetParams{
			APIVersion: "v2",
			Limit:      &limit,
			Q:          remoteURL,
			Resolve:    &resolve,
			Type:       &type_,
		}, auth)
		if err != nil {
			slog.Error("resolve error", "error", err)
			break
		}

		slog.Debug("statuses remaining to import", "count", remaining)

		for _, status := range resp.Payload.Statuses {
			slog.Info("resolved", "url", remoteURL, "id", status.ID)
		}

		numImported := len(resp.Payload.Statuses)
		remaining -= numImported

		numFailures := 1 - numImported
		if numFailures > 0 {
			statusesRequested <- numFailures
			slog.Debug("requested statuses", "count", numFailures)
		}
	}

	close(statusesRequested)
	done <- err
	slog.Debug("finished")
}
