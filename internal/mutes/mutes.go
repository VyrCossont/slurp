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

package mutes

import (
	"log/slog"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/pkg/errors"

	"github.com/VyrCossont/slurp/client/accounts"
	"github.com/VyrCossont/slurp/client/mutes"
	"github.com/VyrCossont/slurp/internal/api"
	"github.com/VyrCossont/slurp/internal/auth"
	"github.com/VyrCossont/slurp/internal/own"
	"github.com/VyrCossont/slurp/internal/resolve"
	"github.com/VyrCossont/slurp/internal/util"
	"github.com/VyrCossont/slurp/models"
)

// Note: Mastodon's mute list exports currently don't include the expiration time.

func Export(authClient *auth.Client, file string) error {
	pagedRequester := &mutesPagedRequester{}
	mutedAccounts, err := api.ReadAllPaged(authClient, pagedRequester)
	if err != nil {
		return err
	}

	ownDomain, err := own.Domain(authClient)
	if err != nil {
		return err
	}

	mutedAccountIDs := make([]string, len(mutedAccounts))
	for _, account := range mutedAccounts {
		mutedAccountIDs = append(mutedAccountIDs, account.ID)
	}
	relationships, err := api.GetBatchedRelationships(authClient, mutedAccountIDs)
	if err != nil {
		return err
	}

	csvRows := make([][]string, 0, 1+len(mutedAccounts))
	csvRows = append(csvRows, csvHeader)
	for _, account := range mutedAccounts {
		relationship, found := relationships[account.ID]
		if !found {
			slog.Warn("couldn't find relationship data", "account_id", account.ID)
		}

		entry, err := newMuteListEntry(ownDomain, account, relationship)
		if err != nil {
			slog.Warn("couldn't convert muted account to list entry", "error", err)
			continue
		}
		csvRows = append(csvRows, entry.csvFields())
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
		mute, err := newMuteListEntryFromCsvFields(row)
		if err != nil {
			slog.Warn("couldn't parse mute from CSV row", "row", row, "error", err)
			continue
		}

		account, err := resolve.Account(authClient, "@"+mute.accountAddress)
		if err != nil {
			slog.Warn("couldn't resolve account", "account_address", mute.accountAddress, "error", err)
			continue
		}

		err = authClient.Wait()
		if err != nil {
			return err
		}

		duration := 0.0 // indefinite mute
		if !mute.expiresAt.IsZero() {
			// Form version of API expects an integer expires_in.
			duration = math.Trunc(time.Until(mute.expiresAt).Seconds())
		}
		_, err = authClient.Client.Accounts.AccountMute(
			&accounts.AccountMuteParams{
				ID:            account.ID,
				Notifications: &mute.notifications,
				Duration:      &duration,
			},
			authClient.Auth,
			func(op *runtime.ClientOperation) {
				op.ConsumesMediaTypes = []string{"application/x-www-form-urlencoded"}
			},
		)
		if err != nil {
			slog.Warn("couldn't mute account", "account_address", mute.accountAddress, "error", err)
		}
	}

	return nil
}

// mutesPagedRequester has no state.
type mutesPagedRequester struct {
}

func (pagedRequester *mutesPagedRequester) Request(authClient *auth.Client, maxID *string) (*mutesPagedResponse, error) {
	resp, err := authClient.Client.Mutes.MutesGet(&mutes.MutesGetParams{
		MaxID: maxID,
	}, authClient.Auth)
	if err != nil {
		return nil, err
	}
	return &mutesPagedResponse{resp}, nil
}

type mutesPagedResponse struct {
	resp *mutes.MutesGetOK
}

func (pagedResponse *mutesPagedResponse) Link() string {
	return pagedResponse.resp.Link
}

func (pagedResponse *mutesPagedResponse) Elements() []*models.MutedAccount {
	return pagedResponse.resp.GetPayload()
}

var csvHeader = []string{
	"Account address",
	"Hide notifications",
	"Expires at",
}

type muteListEntry struct {
	accountAddress string
	notifications  bool
	expiresAt      time.Time
}

func newMuteListEntry(ownDomain string, account *models.MutedAccount, relationship *models.Relationship) (*muteListEntry, error) {
	var err error
	e := &muteListEntry{
		accountAddress: account.Acct,
	}
	if !strings.ContainsRune(e.accountAddress, '@') {
		e.accountAddress += "@" + ownDomain
	}
	if account.MuteExpiresAt != "" {
		e.expiresAt, err = time.Parse(time.RFC3339Nano, account.MuteExpiresAt)
		if err != nil {
			e.expiresAt, err = time.Parse(time.RFC3339, account.MuteExpiresAt)
		}
		if err != nil {
			slog.Error("error parsing timestamp", "error", err, "timestamp", account.MuteExpiresAt)
			return nil, err
		}
	}
	if relationship != nil {
		e.notifications = relationship.MutingNotifications
	}
	return e, nil
}

func (e *muteListEntry) csvFields() []string {
	expiresAt := ""
	if !e.expiresAt.IsZero() {
		expiresAt = e.expiresAt.Format(time.RFC3339)
	}
	return []string{
		e.accountAddress,
		strconv.FormatBool(e.notifications),
		expiresAt,
	}
}

func newMuteListEntryFromCsvFields(fields []string) (*muteListEntry, error) {
	var err error
	e := &muteListEntry{}

	if len(fields) == 0 {
		return nil, errors.WithStack(errors.New("not enough fields, expected at least 1"))
	}
	e.accountAddress = fields[0]

	if len(fields) > 1 {
		e.notifications, err = strconv.ParseBool(fields[1])
		if err != nil {
			err = errors.WithStack(errors.New("boolean expected"))
			slog.Error("malformed whole word", "fields", fields)
			return nil, err
		}
	}

	if len(fields) > 2 {
		timestamp := fields[2]
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

	return e, nil
}
