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

package blocks

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/VyrCossont/slurp/client/blocks"
	"github.com/VyrCossont/slurp/internal/api"
	"github.com/VyrCossont/slurp/internal/auth"
	"github.com/VyrCossont/slurp/internal/own"
	"github.com/VyrCossont/slurp/internal/util"
	"github.com/VyrCossont/slurp/models"
)

// Note: Mastodon's block list exports currently don't have a CSV header.

func Export(authClient *auth.Client, file string) error {
	pagedRequester := &blocksPagedRequester{}
	blockedAccounts, err := api.ReadAllPaged(authClient, pagedRequester)
	if err != nil {
		return err
	}

	ownDomain, err := own.Domain(authClient)
	if err != nil {
		return err
	}

	//goland:noinspection GoImportUsedAsName
	blocks := make([]*blockListEntry, 0, len(blockedAccounts))
	for _, account := range blockedAccounts {
		blocks = append(blocks, newBlockListEntry(ownDomain, account))
	}

	csvRows := make([][]string, 0, len(blocks))
	for _, block := range blocks {
		csvRows = append(csvRows, block.csvFields())
	}

	return util.WriteCSV(file, csvRows)
}

func Import(authClient *auth.Client, file string) error {
	return nil
}

// blocksPagedRequester has no state.
type blocksPagedRequester struct {
}

func (pagedRequester *blocksPagedRequester) Request(authClient *auth.Client, maxID *string) (*blocksPagedResponse, error) {
	resp, err := authClient.Client.Blocks.BlocksGet(&blocks.BlocksGetParams{
		MaxID: maxID,
	}, authClient.Auth)
	if err != nil {
		return nil, err
	}
	return &blocksPagedResponse{resp}, nil
}

type blocksPagedResponse struct {
	resp *blocks.BlocksGetOK
}

func (pagedResponse *blocksPagedResponse) Link() string {
	return pagedResponse.resp.Link
}

func (pagedResponse *blocksPagedResponse) Elements() []*models.Account {
	return pagedResponse.resp.GetPayload()
}

type blockListEntry struct {
	accountAddress string
}

func newBlockListEntry(ownDomain string, account *models.Account) *blockListEntry {
	e := &blockListEntry{
		accountAddress: account.Acct,
	}
	if !strings.ContainsRune(e.accountAddress, '@') {
		e.accountAddress += "@" + ownDomain
	}
	return e
}

func (e *blockListEntry) csvFields() []string {
	return []string{
		e.accountAddress,
	}
}

func newBlockListEntryFromCsvFields(fields []string) (*blockListEntry, error) {
	e := &blockListEntry{}

	if len(fields) == 0 {
		return nil, errors.WithStack(errors.New("not enough fields, expected at least 1"))
	}
	e.accountAddress = fields[0]

	return e, nil
}
