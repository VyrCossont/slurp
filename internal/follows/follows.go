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

package follows

import (
	"encoding/csv"
	"errors"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/VyrCossont/slurp/client/accounts"
	"github.com/VyrCossont/slurp/internal/api"
	"github.com/VyrCossont/slurp/internal/auth"
	"github.com/VyrCossont/slurp/internal/own"
	"github.com/VyrCossont/slurp/models"
)

const relationshipBatchSize = 40

func Export(authClient *auth.Client, file string) error {
	ownAccount, err := own.Account(authClient)
	if err != nil {
		return err
	}
	pagedRequester := &accountFollowingPagedRequester{accountID: ownAccount.ID}

	ownInstance, err := own.Instance(authClient)
	if err != nil {
		return err
	}
	ownDomain := ownInstance.AccountDomain
	if ownDomain == "" {
		ownDomain = ownInstance.Domain
	}
	if ownDomain == "" {
		return errors.New("couldn't find domain for accounts on this instance")
	}

	followedAccounts, err := api.ReadAllPaged(authClient, pagedRequester)
	if err != nil {
		return err
	}

	// Fetch followed account relationship data in batches.
	relationships := make(map[string]*models.Relationship, len(followedAccounts))
	for i := 0; i < len(followedAccounts); i += relationshipBatchSize {
		params := &accounts.AccountRelationshipsParams{}
		for _, account := range followedAccounts[i:min(i+relationshipBatchSize, len(followedAccounts))] {
			params.ID = append(params.ID, account.ID)
		}

		err = authClient.Wait()
		if err != nil {
			return err
		}
		resp, err := authClient.Client.Accounts.AccountRelationships(params, authClient.Auth)
		if err != nil {
			return err
		}

		for _, relationship := range resp.GetPayload() {
			relationships[relationship.ID] = relationship
		}
	}

	follows := make([]*followData, 0, len(followedAccounts))
	for _, account := range followedAccounts {
		relationship, found := relationships[account.ID]
		if !found {
			slog.Warn("couldn't find relationship data", "account_id", account.ID)
		}
		follows = append(follows, newFollowData(ownDomain, account, relationship))
	}

	csvRows := make([][]string, 0, 1+len(follows))
	csvRows = append(csvRows, []string{
		"Account address",
		"Show boosts",
		"Notify on new posts",
		"Languages",
	})
	for _, follow := range follows {
		csvRows = append(csvRows, follow.csvFields())
	}

	out := os.Stdout
	if file != "" {
		out, err = os.Create(file)
		if err != nil {
			slog.Error("couldn't create output file", "error", err)
			return err
		}
		defer func() { _ = out.Close() }()
	}

	w := csv.NewWriter(out)
	err = w.WriteAll(csvRows)
	if err != nil {
		slog.Error("couldn't write to output file", "error", err)
		return err
	}

	return nil
}

func Import(authClient *auth.Client, file string) error {
	// TODO
	return errors.New("NOT IMPLEMENTED")
}

type accountFollowingPagedRequester struct {
	accountID string
}

func (pagedRequester *accountFollowingPagedRequester) Request(authClient *auth.Client, maxID *string) (*accountFollowingPagedResponse, error) {
	resp, err := authClient.Client.Accounts.AccountFollowing(&accounts.AccountFollowingParams{
		ID:    pagedRequester.accountID,
		MaxID: maxID,
	}, authClient.Auth)
	if err != nil {
		return nil, err
	}
	return &accountFollowingPagedResponse{resp}, nil
}

type accountFollowingPagedResponse struct {
	resp *accounts.AccountFollowingOK
}

func (pagedResponse *accountFollowingPagedResponse) Link() string {
	return pagedResponse.resp.Link
}

func (pagedResponse *accountFollowingPagedResponse) Elements() []*models.Account {
	return pagedResponse.resp.GetPayload()
}

type followData struct {
	accountAddress   string
	showBoosts       bool
	notifyOnNewPosts bool
	languages        []string
}

func newFollowData(ownDomain string, account *models.Account, relationship *models.Relationship) *followData {
	f := &followData{
		accountAddress:   account.Acct,
		showBoosts:       true,
		notifyOnNewPosts: false,
		// FUTURE: Not supported by GtS API.
		languages: nil,
	}
	if !strings.ContainsRune(f.accountAddress, '@') {
		f.accountAddress += "@" + ownDomain
	}
	if relationship != nil {
		f.showBoosts = relationship.ShowingReblogs
		f.notifyOnNewPosts = relationship.Notifying
	}
	return f
}

func (f *followData) csvFields() []string {
	return []string{
		f.accountAddress,
		strconv.FormatBool(f.showBoosts),
		strconv.FormatBool(f.notifyOnNewPosts),
		strings.Join(f.languages, ", "),
	}
}

func newFollowDataFromCsvFields(fields []string) (*followData, error) {
	var err error
	f := &followData{}

	if len(fields) == 0 {
		return nil, errors.New("not enough fields, expected at least 1")
	}
	f.accountAddress = fields[0]
	if !strings.ContainsRune(f.accountAddress, '@') {
		err = errors.New("account address expected to contain @")
		slog.Error("malformed account address", "fields", fields)
		return nil, err
	}

	if len(fields) > 1 {
		f.showBoosts, err = strconv.ParseBool(fields[1])
		if err != nil {
			err = errors.New("boolean expected")
			slog.Error("malformed show boosts", "fields", fields)
			return nil, err
		}
	}

	if len(fields) > 2 {
		f.notifyOnNewPosts, err = strconv.ParseBool(fields[2])
		if err != nil {
			err = errors.New("boolean expected")
			slog.Error("malformed notify on new posts", "fields", fields)
			return nil, err
		}
	}

	if len(fields) > 3 {
		for _, language := range strings.Split(fields[3], ",") {
			f.languages = append(f.languages, strings.TrimSpace(language))
		}
	}

	return f, nil
}
