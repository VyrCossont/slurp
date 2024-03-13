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
	"log/slog"
	"strconv"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/pkg/errors"

	"github.com/VyrCossont/slurp/client/accounts"
	"github.com/VyrCossont/slurp/internal/api"
	"github.com/VyrCossont/slurp/internal/auth"
	"github.com/VyrCossont/slurp/internal/own"
	"github.com/VyrCossont/slurp/internal/resolve"
	"github.com/VyrCossont/slurp/internal/util"
	"github.com/VyrCossont/slurp/models"
)

const relationshipBatchSize = 40

func Export(authClient *auth.Client, file string) error {
	ownAccount, err := own.Account(authClient)
	if err != nil {
		return err
	}
	pagedRequester := &accountFollowingPagedRequester{accountID: ownAccount.ID}

	ownDomain, err := own.Domain(authClient)
	if err != nil {
		return err
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

		err := authClient.Wait()
		if err != nil {
			return err
		}

		resp, err := authClient.Client.Accounts.AccountRelationships(params, authClient.Auth)
		if err != nil {
			return errors.WithStack(err)
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
	csvRows = append(csvRows, csvHeader)
	for _, follow := range follows {
		csvRows = append(csvRows, follow.csvFields())
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
		follow, err := newFollowDataFromCsvFields(row)
		if err != nil {
			slog.Warn("couldn't parse follow from CSV row", "row", row, "error", err)
			continue
		}

		account, err := resolve.Account(authClient, "@"+follow.accountAddress)
		if err != nil {
			slog.Warn("couldn't resolve account", "account_address", follow.accountAddress, "error", err)
			continue
		}

		err = authClient.Wait()
		if err != nil {
			return err
		}

		if account.Locked {
			slog.Info("can't follow account because it's locked, so a follow request will be sent", "account_address", follow.accountAddress)
		}

		if len(follow.languages) > 0 {
			slog.Warn("filtering followed account status languages isn't yet supported by the GtS API", "account_address", follow.accountAddress, "languages", len(follow.languages))
		}

		_, err = authClient.Client.Accounts.AccountFollow(
			&accounts.AccountFollowParams{
				ID:      account.ID,
				Notify:  util.Ptr(follow.notifyOnNewPosts),
				Reblogs: util.Ptr(follow.showBoosts),
			},
			authClient.Auth,
			func(op *runtime.ClientOperation) {
				op.ConsumesMediaTypes = []string{"application/x-www-form-urlencoded"}
			},
		)
		if err != nil {
			slog.Warn("couldn't follow account", "account_address", follow.accountAddress, "error", err)
		}
	}

	return nil
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

var csvHeader = []string{
	"Account address",
	"Show boosts",
	"Notify on new posts",
	"Languages",
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
		return nil, errors.WithStack(errors.New("not enough fields, expected at least 1"))
	}
	f.accountAddress = fields[0]
	if !strings.ContainsRune(f.accountAddress, '@') {
		err = errors.WithStack(errors.New("account address expected to contain @"))
		slog.Error("malformed account address", "fields", fields)
		return nil, err
	}

	if len(fields) > 1 {
		f.showBoosts, err = strconv.ParseBool(fields[1])
		if err != nil {
			err = errors.WithStack(errors.New("boolean expected"))
			slog.Error("malformed show boosts", "fields", fields)
			return nil, err
		}
	}

	if len(fields) > 2 {
		f.notifyOnNewPosts, err = strconv.ParseBool(fields[2])
		if err != nil {
			err = errors.WithStack(errors.New("boolean expected"))
			slog.Error("malformed notify on new posts", "fields", fields)
			return nil, err
		}
	}

	if len(fields) > 3 {
		languages := fields[3]
		if languages != "" {
			for _, language := range strings.Split(languages, ",") {
				f.languages = append(f.languages, strings.TrimSpace(language))
			}
		}
	}

	return f, nil
}
