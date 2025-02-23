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
	"os"

	"github.com/pkg/errors"
	"github.com/zalando/go-keyring"
)

// SystemKeyring forwards all operations to the go-keyring global system keyring.
var SystemKeyring keyring.Keyring

// CleartextFileKeyring stores everything in a cleartext JSON file.
var CleartextFileKeyring keyring.Keyring

func init() {
	SystemKeyring = systemKeyring{}
	CleartextFileKeyring = cleartextFileKeyring{}
}

// SystemKeyring forwards all operations to the go-keyring global system keyring.
type systemKeyring struct{}

func (s systemKeyring) Set(service, user, password string) error {
	return keyring.Set(service, user, password)
}

func (s systemKeyring) Get(service, user string) (string, error) {
	return keyring.Get(service, user)
}

func (s systemKeyring) Delete(service, user string) error {
	return keyring.Delete(service, user)
}

func (s systemKeyring) DeleteAll(service string) error {
	return keyring.DeleteAll(service)
}

type cleartextFileKeyring struct{}

// load returns the keyring from disk or an empty one if it's not there yet.
func load() (map[string]map[string]string, error) {
	f, err := os.Open(keyringPath)
	if err != nil {
		return map[string]map[string]string{}, nil
	}
	defer func() { _ = f.Close() }()

	var serviceMap map[string]map[string]string
	err = json.NewDecoder(f).Decode(&serviceMap)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return serviceMap, nil
}

// save creates on-disk preferences or overwrites existing ones.
func save(serviceMap map[string]map[string]string) error {
	err := os.MkdirAll(prefsDir, 0o755)
	if err != nil {
		return errors.WithStack(err)
	}

	f, err := os.OpenFile(keyringPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() { _ = f.Close() }()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)
	err = encoder.Encode(serviceMap)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (c cleartextFileKeyring) Set(service, user, password string) error {
	serviceMap, err := load()
	if err != nil {
		return err
	}

	userMap, found := serviceMap[service]
	if !found {
		userMap = map[string]string{}
	}

	userMap[user] = password

	serviceMap[service] = userMap
	return save(serviceMap)
}

func (c cleartextFileKeyring) Get(service, user string) (string, error) {
	serviceMap, err := load()
	if err != nil {
		return "", err
	}

	userMap, found := serviceMap[service]
	if !found {
		return "", keyring.ErrNotFound
	}

	password, found := userMap[user]
	if !found {
		return "", keyring.ErrNotFound
	}

	return password, nil
}

func (c cleartextFileKeyring) Delete(service, user string) error {
	serviceMap, err := load()
	if err != nil {
		return err
	}

	userMap, found := serviceMap[service]
	if !found {
		return keyring.ErrNotFound
	}

	delete(userMap, user)

	serviceMap[service] = userMap
	return save(serviceMap)
}

func (c cleartextFileKeyring) DeleteAll(service string) error {
	serviceMap, err := load()
	if err != nil {
		return err
	}

	delete(serviceMap, service)

	return save(serviceMap)
}
