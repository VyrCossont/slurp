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
	"log/slog"
	"os"
	"path"

	"github.com/VyrCossont/slurp/internal/auth"
	"github.com/VyrCossont/slurp/internal/own"
	"github.com/VyrCossont/slurp/internal/util"
)

// Export exports a vaguely Mastodon-compatible archive to file (actually a folder path).
func Export(
	authClient *auth.Client,
	archiveFolderPath string,
) error {
	err := checkArchiveFolder(archiveFolderPath)
	if err != nil {
		return err
	}

	err = exportActor(authClient, archiveFolderPath)
	if err != nil {
		return err
	}

	// FIXME: (Vyr) implement this
	return nil
}

func checkArchiveFolder(archiveFolderPath string) error {
	// If the folder doesn't exist, create it.
	stat, err := os.Stat(archiveFolderPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(archiveFolderPath, 0755)
			if err != nil {
				slog.Error("archive folder didn't exist, failed to create it", "path", archiveFolderPath, "err", err)
				return err
			}
		} else {
			slog.Error("couldn't get archive info from filesystem", "path", archiveFolderPath, "err", err)
			return err
		}
	} else if !stat.IsDir() {
		return errors.New("archive folder path is already occupied by something that isn't a folder")
	}

	// Check that it's empty.
	empty, err := util.IsEmpty(archiveFolderPath)
	if err != nil {
		slog.Error("couldn't check whether archive folder is empty", "path", archiveFolderPath, "err", err)
		return err
	}
	if !empty {
		return errors.New("archive folder must be empty")
	}

	return nil
}

func exportActor(
	authClient *auth.Client,
	archiveFolderPath string,
) error {
	account, err := own.Account(authClient)
	if err != nil {
		slog.Error("couldn't retrieve your account", "err", err)
		return err
	}
	actor := &Actor{
		Id:                account.URL,
		Followers:         account.URL + "/followers",
		Outbox:            "outbox.json",
		PreferredUsername: account.Username,
		Url:               account.URL,
	}
	return util.SaveJSON(path.Join(archiveFolderPath, "actor.json"), actor)
}
