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
	"github.com/go-openapi/runtime"

	apiclient "github.com/VyrCossont/slurp/client"
	"github.com/VyrCossont/slurp/models"
)

// Account returns the currently authenticated account.
func Account(client *apiclient.GoToSocialSwaggerDocumentation, auth runtime.ClientAuthInfoWriter) (*models.Account, error) {
	resp, err := client.Accounts.AccountVerify(nil, auth)
	if err != nil {
		return nil, err
	}
	return resp.GetPayload(), nil
}

// Instance returns the instance of the currently authenticated account.
func Instance(client *apiclient.GoToSocialSwaggerDocumentation) (*models.InstanceV2, error) {
	resp, err := client.Instance.InstanceGetV2(nil)
	if err != nil {
		return nil, err
	}
	return resp.GetPayload(), nil
}