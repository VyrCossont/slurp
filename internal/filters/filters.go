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

	"github.com/pkg/errors"

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
	filters := resp.GetPayload()

	csvRows := make([][]string, 0, 1+len(filters))
	csvRows = append(csvRows, csvHeader)
	for _, filter := range filters {
		csvRows = append(csvRows, newFilterListEntry(filter).csvFields())
	}

	return util.WriteCSV(file, csvRows)
}

func Import(authClient *auth.Client, file string) error {
	return nil
}

var csvHeader = []string{
	"Keyword",
	"Whole word",
	"Action",
	"Contexts",
}

type filterListEntry struct {
	keyword   string
	wholeWord bool
	action    string
	contexts  []string
}

func newFilterListEntry(filter *models.FilterV1) *filterListEntry {
	e := &filterListEntry{
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
	return []string{
		e.keyword,
		strconv.FormatBool(e.wholeWord),
		e.action,
		strings.Join(e.contexts, ", "),
	}
}

func newFilterListEntryFromCsvFields(fields []string) (*filterListEntry, error) {
	var err error
	e := &filterListEntry{}

	if len(fields) == 0 {
		return nil, errors.WithStack(errors.New("not enough fields, expected at least 1"))
	}
	e.keyword = fields[0]

	if len(fields) > 1 {
		e.wholeWord, err = strconv.ParseBool(fields[1])
		if err != nil {
			err = errors.WithStack(errors.New("boolean expected"))
			slog.Error("malformed whole word", "fields", fields)
			return nil, err
		}
	}

	if len(fields) > 2 {
		action := fields[2]
		switch action {
		case "hide", "drop":
			e.action = action
		default:
			err = errors.WithStack(errors.New("boolean expected"))
			slog.Error("unexpected action (allows \"hide\" or \"drop\")", "fields", fields)
			return nil, err
		}
	}

	if len(fields) > 3 {
		contexts := fields[3]
		if contexts != "" {
			for _, context := range strings.Split(contexts, ",") {
				e.contexts = append(e.contexts, strings.TrimSpace(context))
			}
		}
	}

	return e, nil
}
