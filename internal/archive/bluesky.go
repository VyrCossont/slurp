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

package archive

import (
	"errors"

	"github.com/VyrCossont/slurp/internal/auth"
)

func BlueskyImport(
	authClient *auth.Client,
	file string,
	statusMapFile string,
	attachmentMapFile string,
	attachmentDirectory string,
) error {
	// TODO: import from Bluesky CAR + blobs
	return errors.New("not yet implemented")
}
