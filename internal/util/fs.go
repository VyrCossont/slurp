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

package util

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"os"
)

func CheckArchiveFolder(archiveFolderPath string) error {
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
	empty, err := IsEmpty(archiveFolderPath)
	if err != nil {
		slog.Error("couldn't check whether archive folder is empty", "path", archiveFolderPath, "err", err)
		return err
	}
	if !empty {
		return errors.New("archive folder must be empty")
	}

	return nil
}

// IsEmpty returns whether a folder is empty.
func IsEmpty(folderPath string) (bool, error) {
	archiveFolder, err := os.Open(folderPath)
	if err != nil {
		slog.Error("couldn't open folder to list it", "path", folderPath, "err", err)
		return false, err
	}
	defer func() {
		err = archiveFolder.Close()
		if err != nil {
			slog.Error("couldn't close folder", "path", folderPath, "err", err)
		}
	}()
	_, err = archiveFolder.Readdirnames(1)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return true, nil
		}
		slog.Error("couldn't list folder", "path", folderPath, "err", err)
		return false, err
	}
	return false, nil
}

func LoadJSON[T any](path string) (*T, error) {
	f, err := os.Open(path)
	if err != nil {
		slog.Error("couldn't open input file", "path", path, "err", err)
		return nil, err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			slog.Error("couldn't close input file", "path", path, "err", err)
		}
	}()

	var doc T
	if err := json.NewDecoder(f).Decode(&doc); err != nil {
		slog.Error("couldn't read data as JSON from input file", "path", path, "err", err)
		return nil, err
	}
	return &doc, nil
}

func SaveJSON(path string, data any) error {
	f, err := os.Create(path)
	if err != nil {
		slog.Error("couldn't create output file", "path", path, "err", err)
		return err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			slog.Error("couldn't close output file", "path", path, "err", err)
		}
	}()

	encoder := json.NewEncoder(f)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	if err = encoder.Encode(data); err != nil {
		slog.Error("couldn't write data as JSON to output file", "path", path, "data", data, "err", err)
		return err
	}
	return nil
}
