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

package bookmarks

import (
	"log/slog"

	"github.com/go-openapi/runtime"
	"github.com/pkg/errors"

	"github.com/VyrCossont/slurp/client/bookmarks"
	"github.com/VyrCossont/slurp/client/statuses"
	"github.com/VyrCossont/slurp/internal/api"
	"github.com/VyrCossont/slurp/internal/auth"
	"github.com/VyrCossont/slurp/internal/resolve"
	"github.com/VyrCossont/slurp/internal/util"
	"github.com/VyrCossont/slurp/models"
)

// Note: Mastodon's bookmark list exports currently don't have a CSV header.

func Export(authClient *auth.Client, file string) error {
	pagedRequester := &bookmarksPagedRequester{}
	bookmarkedStatuses, err := api.ReadAllPaged(authClient, pagedRequester)
	if err != nil {
		return err
	}

	//goland:noinspection GoImportUsedAsName
	bookmarks := make([]*bookmarkListEntry, 0, len(bookmarkedStatuses))
	for _, status := range bookmarkedStatuses {
		bookmarks = append(bookmarks, newBookmarkListEntry(status))
	}

	csvRows := make([][]string, 0, len(bookmarks))
	for _, bookmark := range bookmarks {
		csvRows = append(csvRows, bookmark.csvFields())
	}

	return util.WriteCSV(file, csvRows)
}

func Import(authClient *auth.Client, file string) error {
	csvRows, err := util.ReadCSV(file)
	if err != nil {
		return err
	}

	for _, row := range csvRows {
		bookmark, err := newBookmarkListEntryFromCsvFields(row)
		if err != nil {
			slog.Warn("couldn't parse bookmark from CSV row", "row", row, "error", err)
			continue
		}

		status, err := resolve.Status(authClient, bookmark.statusURI)
		if err != nil {
			slog.Warn("couldn't resolve status", "status_uri", bookmark.statusURI, "error", err)
			continue
		}

		err = authClient.Wait()
		if err != nil {
			return err
		}

		_, err = authClient.Client.Statuses.StatusBookmark(
			&statuses.StatusBookmarkParams{
				ID: status.ID,
			},
			authClient.Auth,
			func(op *runtime.ClientOperation) {
				op.ConsumesMediaTypes = []string{"application/x-www-form-urlencoded"}
			},
		)
		if err != nil {
			slog.Warn("couldn't bookmark status", "status_uri", bookmark.statusURI, "error", err)
		}
	}

	return nil
}

// bookmarksPagedRequester has no state.
type bookmarksPagedRequester struct {
}

func (pagedRequester *bookmarksPagedRequester) Request(authClient *auth.Client, maxID *string) (*bookmarksPagedResponse, error) {
	resp, err := authClient.Client.Bookmarks.BookmarksGet(&bookmarks.BookmarksGetParams{
		MaxID: maxID,
	}, authClient.Auth)
	if err != nil {
		return nil, err
	}
	return &bookmarksPagedResponse{resp}, nil
}

type bookmarksPagedResponse struct {
	resp *bookmarks.BookmarksGetOK
}

func (pagedResponse *bookmarksPagedResponse) Link() string {
	return pagedResponse.resp.Link
}

func (pagedResponse *bookmarksPagedResponse) Elements() []*models.Status {
	return pagedResponse.resp.GetPayload()
}

type bookmarkListEntry struct {
	statusURI string
}

func newBookmarkListEntry(status *models.Status) *bookmarkListEntry {
	e := &bookmarkListEntry{
		statusURI: status.URI,
	}
	return e
}

func (e *bookmarkListEntry) csvFields() []string {
	return []string{
		e.statusURI,
	}
}

func newBookmarkListEntryFromCsvFields(fields []string) (*bookmarkListEntry, error) {
	e := &bookmarkListEntry{}

	if len(fields) == 0 {
		return nil, errors.WithStack(errors.New("not enough fields, expected at least 1"))
	}
	e.statusURI = fields[0]

	return e, nil
}
