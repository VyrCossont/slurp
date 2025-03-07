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

package lists

import (
	"log/slog"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/pkg/errors"

	"github.com/VyrCossont/slurp/client/lists"
	"github.com/VyrCossont/slurp/internal/api"
	"github.com/VyrCossont/slurp/internal/auth"
	"github.com/VyrCossont/slurp/internal/own"
	"github.com/VyrCossont/slurp/internal/resolve"
	"github.com/VyrCossont/slurp/internal/util"
	"github.com/VyrCossont/slurp/models"
)

// Note: Mastodon's list list exports currently don't have a CSV header.

func Export(authClient *auth.Client, file string) error {
	ownDomain, err := own.Domain(authClient)
	if err != nil {
		return err
	}

	var entries []*listListEntry

	err = authClient.Wait()
	if err != nil {
		return err
	}

	resp, err := authClient.Client.Lists.Lists(nil, authClient.Auth)
	if err != nil {
		slog.Error("couldn't fetch existing lists", "error", err)
		return errors.WithStack(err)
	}
	for _, list := range resp.GetPayload() {
		pagedRequester := &listAccountsPagedRequester{listID: list.ID}
		listAccounts, err := api.ReadAllPaged(authClient, pagedRequester, nil, nil)
		if err != nil {
			slog.Error("couldn't fetch list accounts", "title", list.Title, "list_id", list.ID, "error", err)
			continue
		}
		for _, account := range listAccounts {
			entries = append(entries, newListListEntry(ownDomain, list, account))
		}
	}

	csvRows := make([][]string, 0, len(entries))
	for _, entry := range entries {
		csvRows = append(csvRows, entry.csvFields())
	}

	return util.WriteCSV(file, csvRows)
}

func Import(authClient *auth.Client, file string) error {
	csvRows, err := util.ReadCSV(file)
	if err != nil {
		return err
	}

	// Collect set of all accounts so we can resolve them only once each.
	listAccountAddressesSet := map[string]interface{}{}
	// Collect list of accounts in each list.
	listAccountAddressesByTitle := map[string][]string{}
	for _, row := range csvRows {
		listListEntry, err := newListListEntryFromCsvFields(row)
		if err != nil {
			slog.Warn("couldn't parse list entry from CSV row", "row", row, "error", err)
			continue
		}
		listAccountAddressesSet[listListEntry.accountAddress] = nil
		listAccountAddressesByTitle[listListEntry.title] = append(
			listAccountAddressesByTitle[listListEntry.title],
			listListEntry.accountAddress,
		)
	}

	// Resolve all accounts.
	listAccountIDsByAddress := map[string]string{}
	for accountAddress := range listAccountAddressesSet {
		account, err := resolve.Account(authClient, "@"+accountAddress)
		if err != nil {
			slog.Warn("couldn't resolve account", "account_address", accountAddress, "error", err)
			continue
		}
		listAccountIDsByAddress[accountAddress] = account.ID
	}

	// Fetch existing lists.
	existingListIDsByTitle := map[string]string{}

	err = authClient.Wait()
	if err != nil {
		return err
	}

	resp, err := authClient.Client.Lists.Lists(nil, authClient.Auth)
	if err != nil {
		slog.Warn("couldn't fetch existing lists", "error", err)
		// We can technically keep going without them, but it really should succeed.
	} else {
		for _, list := range resp.GetPayload() {
			existingListIDsByTitle[list.Title] = list.ID
		}
	}

	// Create each list if necessary and add accounts to them.
	for title, listAccountAddresses := range listAccountAddressesByTitle {
		var listID string
		var found bool
		if listID, found = existingListIDsByTitle[title]; !found {
			// Create a new list.
			err := authClient.Wait()
			if err != nil {
				return err
			}

			resp, err := authClient.Client.Lists.ListCreate(
				&lists.ListCreateParams{
					Title: title,
				},
				authClient.Auth,
				func(op *runtime.ClientOperation) {
					op.ConsumesMediaTypes = []string{"application/x-www-form-urlencoded"}
				},
			)
			if err != nil {
				slog.Warn("couldn't create list", "title", title, "error", err)
				continue
			}
			listID = resp.GetPayload().ID
		}

		// Add list accounts one by one.
		// The Mastodon API specifies that adding accounts to a list is not idempotent,
		// and will fail if the same account is already there. GtS is flaw-compatible.
		for _, accountAddress := range listAccountAddresses {
			var accountID string
			if accountID, found = listAccountIDsByAddress[accountAddress]; !found {
				// We've already logged a warning for not being able to resolve this account.
				continue
			}
			params := &lists.AddListAccountsParams{
				ID:         listID,
				AccountIds: []string{accountID},
			}

			err := authClient.Wait()
			if err != nil {
				return err
			}

			_, err = authClient.Client.Lists.AddListAccounts(
				params,
				authClient.Auth,
				func(op *runtime.ClientOperation) {
					op.ConsumesMediaTypes = []string{"application/x-www-form-urlencoded"}
				},
			)
			if err != nil {
				slog.Warn("couldn't add account to list", "title", title, "list_id", listID, "account_id", accountID, "error", err)
				continue
			}
		}
	}

	return nil
}

type listAccountsPagedRequester struct {
	listID        string
	forwardPaging bool
}

func (pagedRequester *listAccountsPagedRequester) Request(authClient *auth.Client, maxID *string, minID *string) (*listAccountsPagedResponse, error) {
	resp, err := authClient.Client.Lists.ListAccounts(&lists.ListAccountsParams{
		ID:    pagedRequester.listID,
		MaxID: maxID,
		MinID: minID,
	}, authClient.Auth)
	if err != nil {
		return nil, err
	}
	return &listAccountsPagedResponse{resp}, nil
}

func (pagedRequester *listAccountsPagedRequester) ForwardPaging() bool {
	return pagedRequester.forwardPaging
}

type listAccountsPagedResponse struct {
	resp *lists.ListAccountsOK
}

func (pagedResponse *listAccountsPagedResponse) Link() string {
	return pagedResponse.resp.Link
}

func (pagedResponse *listAccountsPagedResponse) Elements() []*models.Account {
	return pagedResponse.resp.GetPayload()
}

type listListEntry struct {
	// title is repeated as many times as there are accounts in a list.
	title          string
	accountAddress string
}

func newListListEntry(ownDomain string, list *models.List, account *models.Account) *listListEntry {
	e := &listListEntry{
		title:          list.Title,
		accountAddress: account.Acct,
	}
	if !strings.ContainsRune(e.accountAddress, '@') {
		e.accountAddress += "@" + ownDomain
	}
	return e
}

func (e *listListEntry) csvFields() []string {
	return []string{
		e.title,
		e.accountAddress,
	}
}

func newListListEntryFromCsvFields(fields []string) (*listListEntry, error) {
	e := &listListEntry{}

	if len(fields) < 2 {
		return nil, errors.WithStack(errors.New("not enough fields, expected at least 2"))
	}
	e.title = fields[0]
	e.accountAddress = fields[1]
	if !strings.ContainsRune(e.accountAddress, '@') {
		err := errors.WithStack(errors.New("account address expected to contain @"))
		slog.Error("malformed account address", "fields", fields)
		return nil, err
	}

	return e, nil
}
