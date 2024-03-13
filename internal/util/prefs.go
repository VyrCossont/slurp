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
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/pkg/errors"
)

// Prefs stores persisted slurp preferences.
type Prefs struct {
	// Instances is a map of instance names to instance preferences.
	Instances map[string]PrefsInstance `json:"instances,omitempty"`
	// Users is a map of user usernames@domains to instance data.
	Users map[string]PrefsUser `json:"users,omitempty"`
	// DefaultUser is the username@domain of the last user we successfully authenticated as,
	// if there is one.
	DefaultUser string `json:"default_user,omitempty"`
}

// PrefsInstance stores preferences for a given instance.
// OAuth2 app secrets are stored in the system keychain.
type PrefsInstance struct {
	// ClientID is the OAuth2 client ID for slurp on this instance.
	ClientID string `json:"client_id"`
}

// PrefsUser stores preferences for a given user.
// User access tokens are stored in the system keychain.
type PrefsUser struct {
	// Instance is the name of the instance that the user is on.
	Instance string `json:"instance"`
}

// prefsDir is the path to the directory containing all slurp preference files.
var prefsDir string

// prefsPath is the path to the file within that directory that stores all of our prefs.
var prefsPath string

func init() {
	prefsDir = filepath.Join(xdg.ConfigHome, "codes.catgirl.slurp")
	prefsPath = filepath.Join(prefsDir, "prefs.json")
}

// LoadPrefs returns preferences from disk or an empty prefs object if there are none stored or accessible.
func LoadPrefs() (*Prefs, error) {
	f, err := os.Open(prefsPath)
	if err != nil {
		return &Prefs{
			Instances: map[string]PrefsInstance{},
			Users:     map[string]PrefsUser{},
		}, nil
	}
	defer func() { _ = f.Close() }()

	var prefs Prefs
	err = json.NewDecoder(f).Decode(&prefs)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if prefs.Instances == nil {
		prefs.Instances = map[string]PrefsInstance{}
	}
	if prefs.Users == nil {
		prefs.Users = map[string]PrefsUser{}
	}

	return &prefs, nil
}

// SavePrefs creates on-disk preferences or overwrites existing ones.
func SavePrefs(prefs *Prefs) error {
	err := os.MkdirAll(prefsDir, 0o755)
	if err != nil {
		return errors.WithStack(err)
	}

	f, err := os.Create(prefsPath)
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() { _ = f.Close() }()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)
	err = encoder.Encode(prefs)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

var PrefNotFound = errors.New("preference value not found")

func getPrefValue(get func(prefs *Prefs) (string, bool)) (string, error) {
	prefs, err := LoadPrefs()
	if err != nil {
		return "", err
	}

	value, exists := get(prefs)
	if !exists {
		return "", errors.WithStack(PrefNotFound)
	}

	return value, nil
}

func setPrefValue(set func(prefs *Prefs)) error {
	prefs, err := LoadPrefs()
	if err != nil {
		return err
	}

	set(prefs)

	err = SavePrefs(prefs)
	if err != nil {
		return err
	}

	return nil
}

func GetDefaultUser() (string, error) {
	return getPrefValue(func(prefs *Prefs) (string, bool) {
		if prefs.DefaultUser == "" {
			return "", false
		}
		return prefs.DefaultUser, true
	})
}

func SetDefaultUser(user string) error {
	return setPrefValue(func(prefs *Prefs) {
		prefs.DefaultUser = user
	})
}

func GetInstanceClientID(instance string) (string, error) {
	return getPrefValue(func(prefs *Prefs) (string, bool) {
		prefsInstance, exists := prefs.Instances[instance]
		if !exists {
			return "", false
		}
		return prefsInstance.ClientID, true
	})
}

func SetInstanceClientID(instance string, clientID string) error {
	return setPrefValue(func(prefs *Prefs) {
		prefsInstance := prefs.Instances[instance]
		prefsInstance.ClientID = clientID
		prefs.Instances[instance] = prefsInstance
	})
}

func GetUserInstance(user string) (string, error) {
	return getPrefValue(func(prefs *Prefs) (string, bool) {
		prefsUser, exists := prefs.Users[user]
		if !exists {
			return "", false
		}
		return prefsUser.Instance, true
	})
}

func SetUserInstance(user string, instance string) error {
	return setPrefValue(func(prefs *Prefs) {
		prefsUser := prefs.Users[user]
		prefsUser.Instance = instance
		prefs.Users[user] = prefsUser
	})
}
