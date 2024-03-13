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

package util

import (
	"net/url"

	"github.com/peterhellberg/link"
	"github.com/pkg/errors"
)

// ParseLinkMaxID extracts the `max_id` from the `next` link for paging to older items.
func ParseLinkMaxID(linkHeader string) (*string, error) {
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
