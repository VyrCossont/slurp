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

package emojis

import (
	"encoding/json"
	"log/slog"
	"os"

	"github.com/VyrCossont/slurp/internal/auth"
	"github.com/pkg/errors"
)

func Export(authClient *auth.Client, file string) error {
	err := authClient.Wait()
	if err != nil {
		return errors.WithStack(err)
	}

	resp, err := authClient.Client.CustomEmojis.CustomEmojisGet(nil, authClient.Auth)
	if err != nil {
		slog.Error("error getting filters", "error", err)
		return err
	}
	//goland:noinspection GoImportUsedAsName
	emojis := resp.GetPayload()

	out := os.Stdout
	if file != "" {
		out, err = os.Create(file)
		if err != nil {
			slog.Error("couldn't create emojis output file", "file", file, "error", err)
			return errors.WithStack(err)
		}
		defer func() {
			if err := out.Close(); err != nil {
				slog.Error("error closing emojis file", "file", file, "error", err)
			}
		}()
	}

	encoder := json.NewEncoder(out)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(emojis); err != nil {
		slog.Error("error writing emojis to file", "file", file, "error", err)
		return err
	}

	return nil
}

func Import(authClient *auth.Client, file string) error {
	// TODO: (Vyr) load emojis using admin API.

	return nil
}
