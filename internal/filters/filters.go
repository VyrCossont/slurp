// Slurp
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

package filters

import (
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/pkg/errors"

	"github.com/VyrCossont/slurp/client/filters"
	"github.com/VyrCossont/slurp/internal/auth"
	"github.com/VyrCossont/slurp/internal/util"
	"github.com/VyrCossont/slurp/models"
)

func Export(authClient *auth.Client, file string) error {
	err := authClient.Wait()
	if err != nil {
		return err
	}

	resp, err := authClient.Client.Filters.FiltersV1Get(nil, authClient.Auth)
	if err != nil {
		slog.Error("error getting filters", "error", err)
		return err
	}
	//goland:noinspection GoImportUsedAsName
	filters := resp.GetPayload()

	csvRows := make([][]string, 0, 1+len(filters))
	csvRows = append(csvRows, csvHeader)
	for _, filter := range filters {
		csvRows = append(csvRows, newFilterListEntry(filter).csvFields())
	}

	return util.WriteCSV(file, csvRows)
}

func Import(authClient *auth.Client, file string) error {
	csvRows, err := util.ReadCSV(file)
	if err != nil {
		return err
	}

	csvRows, err = util.RemoveExpectedCSVHeader(csvHeader, csvRows)
	if err != nil {
		return err
	}

	for _, row := range csvRows {
		filter, err := newFilterListEntryFromCsvFields(row)
		if err != nil {
			slog.Warn("couldn't parse filter from CSV row", "row", row, "error", err)
			continue
		}

		params := &filters.FilterV1PostParams{
			FContext:     filter.contexts,
			Phrase:       filter.keyword,
			Irreversible: util.Ptr(filter.action == "drop"),
			WholeWord:    util.Ptr(filter.wholeWord),
		}
		if !filter.expiresAt.IsZero() {
			params.ExpiresIn = util.Ptr(filter.expiresAt.Sub(time.Now()).Seconds())
		}

		_, err = authClient.Client.Filters.FilterV1Post(
			params,
			authClient.Auth,
			func(op *runtime.ClientOperation) {
				op.ConsumesMediaTypes = []string{"application/x-www-form-urlencoded"}
			},
		)
		if err != nil {
			slog.Warn("couldn't create filter", "title", filter.title, "error", err)
		}
	}

	return nil
}

var csvHeader = []string{
	"Title",
	"Keyword",
	"Whole word",
	"Action",
	"Expires at",
	"Contexts",
}

type filterListEntry struct {
	title     string
	keyword   string
	wholeWord bool
	action    string
	expiresAt time.Time
	contexts  []string
}

func newFilterListEntry(filter *models.FilterV1) *filterListEntry {
	e := &filterListEntry{
		title:     filter.Phrase,
		keyword:   filter.Phrase,
		wholeWord: filter.WholeWord,
		action:    "hide",
	}
	if filter.Irreversible {
		e.action = "drop"
	}
	for _, context := range filter.Context {
		e.contexts = append(e.contexts, string(context))
	}
	return e
}

func (e *filterListEntry) csvFields() []string {
	expiresAt := ""
	if !e.expiresAt.IsZero() {
		expiresAt = e.expiresAt.Format(time.RFC3339)
	}
	return []string{
		e.title,
		e.keyword,
		strconv.FormatBool(e.wholeWord),
		e.action,
		expiresAt,
		strings.Join(e.contexts, ", "),
	}
}

func newFilterListEntryFromCsvFields(fields []string) (*filterListEntry, error) {
	var err error
	e := &filterListEntry{}

	if len(fields) == 0 {
		return nil, errors.WithStack(errors.New("not enough fields, expected at least 1"))
	}
	e.title = fields[0]
	e.keyword = e.title

	if len(fields) > 1 {
		e.keyword = fields[1]
	}

	if len(fields) > 2 {
		e.wholeWord, err = strconv.ParseBool(fields[2])
		if err != nil {
			err = errors.WithStack(errors.New("boolean expected"))
			slog.Error("malformed whole word", "fields", fields)
			return nil, err
		}
	}

	if len(fields) > 3 {
		action := fields[3]
		switch action {
		case "hide", "drop":
			e.action = action
		default:
			err = errors.WithStack(errors.New("unexpected action"))
			slog.Error("action must be either \"hide\" or \"drop\"", "fields", fields)
			return nil, err
		}
	}

	if len(fields) > 4 {
		timestamp := fields[4]
		if timestamp != "" {
			expiresAt, err := time.Parse(time.RFC3339Nano, timestamp)
			if err != nil {
				expiresAt, err = time.Parse(time.RFC3339, timestamp)
			}
			if err != nil {
				slog.Error("error parsing timestamp", "error", err, "timestamp", timestamp)
				return nil, err
			}
			e.expiresAt = expiresAt
		}
	}

	if len(fields) > 5 {
		contexts := fields[5]
		if contexts != "" {
			for _, context := range strings.Split(contexts, ",") {
				e.contexts = append(e.contexts, strings.TrimSpace(context))
			}
		}
	}

	return e, nil
}
