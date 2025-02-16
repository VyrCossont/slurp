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
	"slices"

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
	ClientID string `json:"client_id,omitempty"`

	// Scheme is the URL scheme for this instance.
	Scheme string `json:"scheme,omitempty"`

	// RateLimit is the rate limit for requests to this instance, in requests per second.
	RateLimit float64 `json:"rate_limit,omitempty"`

	// RateLimit is the burst capacity for requests to this instance, in requests at once.
	BurstCap int `json:"burst_cap,omitempty"`
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

// PrefsPath returns the path to the preferences file, whether or not it exists yet.
func PrefsPath() string {
	return prefsPath
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

// getPrefValue loads the preferences file and uses the get function to extract a value from it.
// If the second return value of get is false, the preference is missing and has no default,
// in which case an error will be raised.
func getPrefValue[T any](get func(prefs *Prefs) (T, bool)) (T, error) {
	var value T

	prefs, err := LoadPrefs()
	if err != nil {
		return value, err
	}

	var exists bool
	value, exists = get(prefs)
	if !exists {
		return value, errors.WithStack(PrefNotFound)
	}

	return value, nil
}

// setPrefValue loads the preferences file, applies the set function to mutate it, and saves it.
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
	users, err := getPrefValue(func(prefs *Prefs) ([]string, bool) {
		var users []string
		for user := range prefs.Users {
			users = append(users, user)
		}
		return users, true
	})
	if err != nil {
		return err
	}

	if !slices.Contains(users, user) {
		return errors.Errorf("there is no authenticated user named %s", user)
	}

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

// DefaultScheme is the default instance URL scheme.
const DefaultScheme = "https"

func GetInstanceScheme(instance string) (string, error) {
	return getPrefValue(func(prefs *Prefs) (string, bool) {
		prefsInstance, exists := prefs.Instances[instance]
		if !exists {
			return "", false
		}
		if prefsInstance.Scheme == "" {
			// Older versions of `slurp` don't store this.
			return DefaultScheme, true
		}
		return prefsInstance.Scheme, true
	})
}

func SetInstanceScheme(instance string, scheme string) error {
	return setPrefValue(func(prefs *Prefs) {
		prefsInstance := prefs.Instances[instance]
		prefsInstance.Scheme = scheme
		prefs.Instances[instance] = prefsInstance
	})
}

// DefaultRateLimit is the default instance rate limit. Lower than the GTS default of 1.0.
const DefaultRateLimit float64 = 0.5

func GetInstanceRateLimit(instance string) (float64, error) {
	return getPrefValue(func(prefs *Prefs) (float64, bool) {
		prefsInstance, exists := prefs.Instances[instance]
		if !exists {
			return 0, false
		}
		if prefsInstance.RateLimit == 0 {
			// Older versions of `slurp` don't store this.
			return DefaultRateLimit, true
		}
		return prefsInstance.RateLimit, true
	})
}

func SetInstanceRateLimit(instance string, rateLimit float64) error {
	return setPrefValue(func(prefs *Prefs) {
		prefsInstance := prefs.Instances[instance]
		prefsInstance.RateLimit = rateLimit
		prefs.Instances[instance] = prefsInstance
	})
}

// DefaultBurstCap is the default instance burst capacity. Lower than the GTS default of 300.
const DefaultBurstCap int = 10

func GetInstanceBurstCap(instance string) (int, error) {
	return getPrefValue(func(prefs *Prefs) (int, bool) {
		prefsInstance, exists := prefs.Instances[instance]
		if !exists {
			return 0, false
		}
		if prefsInstance.BurstCap == 0 {
			// Older versions of `slurp` don't store this.
			return DefaultBurstCap, true
		}
		return prefsInstance.BurstCap, true
	})
}

func SetInstanceBurstCap(instance string, burstCap int) error {
	return setPrefValue(func(prefs *Prefs) {
		prefsInstance := prefs.Instances[instance]
		prefsInstance.BurstCap = burstCap
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
