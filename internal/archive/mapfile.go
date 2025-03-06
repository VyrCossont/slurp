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
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"strings"
)

func readMapFile(mapFile string) (map[string]string, error) {
	jsonFile, err := os.Open(mapFile)
	if err != nil {
		return nil, err
	}
	defer func() { _ = jsonFile.Close() }()

	var doc map[string]string
	if err := json.NewDecoder(jsonFile).Decode(&doc); err != nil {
		return nil, err
	}
	return doc, nil
}

func writeMapFile(mapFile string, doc map[string]string) error {
	jsonFile, err := os.Create(mapFile)
	if err != nil {
		return err
	}
	defer func() { _ = jsonFile.Close() }()

	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)
	return encoder.Encode(doc)
}

func RequireMapFiles(statusMapFile string, attachmentMapFile string) (map[string]string, map[string]string, error) {
	// Require status map file.
	if !strings.HasSuffix(strings.ToLower(statusMapFile), ".json") {
		return nil, nil, errors.New("status map file is required and must have a .json extension")
	}
	archiveIdToImportedApiId, err := readMapFile(statusMapFile)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			archiveIdToImportedApiId = map[string]string{}
		} else {
			return nil, nil, err
		}
	}
	if err = writeMapFile(statusMapFile, archiveIdToImportedApiId); err != nil {
		return nil, nil, err
	}

	// Require attachment map file.
	if !strings.HasSuffix(strings.ToLower(attachmentMapFile), ".json") {
		return nil, nil, errors.New("attachment map file is required and must have a .json extension")
	}
	mediaPathToImportedApiId, err := readMapFile(attachmentMapFile)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			mediaPathToImportedApiId = map[string]string{}
		} else {
			return nil, nil, err
		}
	}
	if err = writeMapFile(attachmentMapFile, mediaPathToImportedApiId); err != nil {
		return nil, nil, err
	}

	return archiveIdToImportedApiId, mediaPathToImportedApiId, nil
}
