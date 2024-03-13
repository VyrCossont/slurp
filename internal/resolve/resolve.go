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

package resolve

import (
	"errors"

	"github.com/VyrCossont/slurp/client/search"
	"github.com/VyrCossont/slurp/internal/auth"
	"github.com/VyrCossont/slurp/internal/util"
	"github.com/VyrCossont/slurp/models"
)

// NoResults indicated we didn't find anything.
// This shouldn't stop a bulk operation.
var NoResults = errors.New("no search results")

// Account uses the search API to resolve an account.
// The account can be a URL, or a @-prefixed username with optional domain.
func Account(authClient auth.Client, accountOrURL string) (*models.Account, error) {
	results, err := search1(authClient, accountOrURL, "accounts")
	if err != nil {
		return nil, err
	}
	if len(results.Accounts) == 0 {
		return nil, NoResults
	}
	return results.Accounts[0], nil
}

// Status uses the search API to resolve a status by URL.
func Status(authClient auth.Client, url string) (*models.Status, error) {
	results, err := search1(authClient, url, "statuses")
	if err != nil {
		return nil, err
	}
	if len(results.Statuses) == 0 {
		return nil, NoResults
	}
	return results.Statuses[0], nil
}

func search1(authClient auth.Client, q string, searchType string) (*models.SearchResult, error) {
	resp, err := authClient.Client.Search.SearchGet(&search.SearchGetParams{
		APIVersion: "v2",
		Limit:      util.Ptr(int64(1)),
		Q:          q,
		Resolve:    util.Ptr(true),
		Type:       util.Ptr(searchType),
	}, authClient.Auth)
	if err != nil {
		return nil, err
	}
	return resp.GetPayload(), nil
}
