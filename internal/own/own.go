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

package own

import (
	"log/slog"

	"github.com/pkg/errors"

	"github.com/VyrCossont/slurp/internal/auth"
	"github.com/VyrCossont/slurp/models"
)

// Account returns the currently authenticated account.
func Account(authClient *auth.Client) (*models.Account, error) {
	err := authClient.Wait()
	if err != nil {
		return nil, err
	}

	resp, err := authClient.Client.Accounts.AccountVerify(nil, authClient.Auth)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return resp.GetPayload(), nil
}

// Instance returns the instance of the currently authenticated account.
func Instance(authClient *auth.Client) (*models.InstanceV2, error) {
	err := authClient.Wait()
	if err != nil {
		return nil, err
	}

	resp, err := authClient.Client.Instance.InstanceGetV2(nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return resp.GetPayload(), nil
}

func Domain(authClient *auth.Client) (string, error) {
	ownInstance, err := Instance(authClient)
	if err != nil {
		return "", err
	}

	ownDomain := ownInstance.AccountDomain
	if ownDomain == "" {
		ownDomain = ownInstance.Domain
	}
	if ownDomain == "" {
		return "", errors.WithStack(errors.New("couldn't find domain for accounts on this instance"))
	}

	return ownDomain, nil
}

// Emojis returns custom emojis available to the authenticated account.
func Emojis(authClient *auth.Client) ([]*models.Emoji, error) {
	if err := authClient.Wait(); err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := authClient.Client.CustomEmojis.CustomEmojisGet(nil, authClient.Auth)
	if err != nil {
		slog.Error("error getting emojis", "error", err)
		return nil, errors.WithStack(err)
	}

	return response.GetPayload(), nil
}
