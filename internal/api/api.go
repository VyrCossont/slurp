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

package api

import (
	"log/slog"

	"github.com/VyrCossont/slurp/client/accounts"
	"github.com/VyrCossont/slurp/models"
	"github.com/pkg/errors"

	"github.com/VyrCossont/slurp/internal/auth"
)

type PagedRequester[Response PagedResponse[Element], Element any] interface {
	Request(authClient *auth.Client, maxID *string) (Response, error)
}

type PagedResponse[Element any] interface {
	Link() string
	Elements() []Element
}

func ReadAllPaged[Requester PagedRequester[Response, Element], Response PagedResponse[Element], Element any](authClient *auth.Client, pagedRequester Requester) ([]Element, error) {
	var all []Element
	var maxID *string

	for {
		err := authClient.Wait()
		if err != nil {
			return all, errors.WithStack(err)
		}

		pagedResponse, err := pagedRequester.Request(authClient, maxID)
		if err != nil {
			slog.Error("error fetching page", "error", err)
			return all, errors.WithStack(err)
		}

		maxID, err = ParseLinkMaxID(pagedResponse.Link())
		if err != nil {
			slog.Error("error parsing Link header", "error", err)
			return all, errors.WithStack(err)
		}
		if maxID == nil {
			// End of pages.
			break
		}

		all = append(all, pagedResponse.Elements()...)
	}

	return all, nil
}

const relationshipBatchSize = 40

// GetBatchedRelationships fetches account relationships in batches.
// If some batches fail, it will keep going.
func GetBatchedRelationships(authClient *auth.Client, accountIDs []string) (relationships map[string]*models.Relationship, err error) {
	relationships = make(map[string]*models.Relationship, len(accountIDs))
	for i := 0; i < len(accountIDs); i += relationshipBatchSize {
		params := &accounts.AccountRelationshipsParams{
			ID: accountIDs[i:min(i+relationshipBatchSize, len(accountIDs))],
		}

		err = authClient.Wait()
		if err != nil {
			return
		}

		resp, err := authClient.Client.Accounts.AccountRelationships(params, authClient.Auth)
		if err != nil {
			slog.Warn("couldn't fetch relationships", "account_ids", params.ID)
			continue
		}

		for _, relationship := range resp.GetPayload() {
			relationships[relationship.ID] = relationship
		}
	}

	return
}
