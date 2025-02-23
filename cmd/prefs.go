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

package cmd

import (
	"errors"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/VyrCossont/slurp/internal/util"
)

// prefsCmd represents the auth command
var prefsCmd = &cobra.Command{
	Use:   "prefs",
	Short: "Get or set preferences",
}

// prefsGetCmd represents the prefs get command
var prefsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a preference",
}

// prefsGetRateLimitCmd represents the prefs get ratelimit command
var prefsGetRateLimitCmd = &cobra.Command{
	Use:   "ratelimit",
	Short: "Get the instance rate limit",
	RunE: func(cmd *cobra.Command, args []string) error {
		instance, err := prefsGetInstance()
		if err != nil {
			return err
		}

		rateLimit, err := util.GetInstanceRateLimit(instance)
		if err != nil {
			return err
		}

		_, err = cmd.OutOrStderr().Write([]byte(fmt.Sprintf("%f\n", rateLimit)))
		return err
	},
}

// prefsGetBurstCapCmd represents the prefs get burstcap command
var prefsGetBurstCapCmd = &cobra.Command{
	Use:   "burstcap",
	Short: "Get the instance burst capacity",
	RunE: func(cmd *cobra.Command, args []string) error {
		instance, err := prefsGetInstance()
		if err != nil {
			return err
		}

		burstCap, err := util.GetInstanceBurstCap(instance)
		if err != nil {
			return err
		}

		_, err = cmd.OutOrStderr().Write([]byte(fmt.Sprintf("%d\n", burstCap)))
		return err
	},
}

// prefsSetCmd represents the prefs set command
var prefsSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a preference",
}

// prefsSetRateLimitCmd represents the prefs set ratelimit command
var prefsSetRateLimitCmd = &cobra.Command{
	Use:   "ratelimit <float>",
	Short: "Set the instance rate limit",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		instance, err := prefsGetInstance()
		if err != nil {
			return err
		}

		rateLimit, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			return err
		}
		if rateLimit <= 0 {
			return errors.New("rate limit must be a positive number")
		}

		return util.SetInstanceRateLimit(instance, rateLimit)
	},
}

// prefsSetBurstCapCmd represents the prefs set burstcap command
var prefsSetBurstCapCmd = &cobra.Command{
	Use:   "burstcap <int>",
	Short: "Set the instance burst capacity",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		instance, err := prefsGetInstance()
		if err != nil {
			return err
		}

		burstCap, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		if burstCap <= 0 {
			return errors.New("burst capacity must be a positive integer")
		}

		return util.SetInstanceBurstCap(instance, burstCap)
	},
}

// prefsPathCmd represents the prefs path command
var prefsPathCmd = &cobra.Command{
	Use:   "path",
	Short: "Get the path to the preferences file",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := cmd.OutOrStderr().Write([]byte(fmt.Sprintf("%s\n", util.PrefsPath())))
		return err
	},
}

// prefsKeyringPathCmd represents the prefs keyring-path command
var prefsKeyringPathCmd = &cobra.Command{
	Use:   "keyring-path",
	Short: "Get the path to the keyring file (whether or not you're using the file-backed keyring)",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := cmd.OutOrStderr().Write([]byte(fmt.Sprintf("%s\n", util.KeyringPath())))
		return err
	},
}

// prefsGetInstance returns the instance for the provided or default user, if there is one.
func prefsGetInstance() (string, error) {
	var err error

	user := User
	if user == "" {
		user, err = util.GetDefaultUser()
		if err != nil {
			slog.Error("no user provided, couldn't get default user from prefs (did you log in first?)")
			return "", err
		}
	}

	instance, err := util.GetUserInstance(user)
	if err != nil {
		slog.Error("couldn't get user's instance from prefs (did you log in first?)", "user", user)
		return "", err
	}

	return instance, nil
}

func init() {
	rootCmd.AddCommand(prefsCmd)

	prefsCmd.AddCommand(prefsGetCmd)

	prefsGetCmd.AddCommand(prefsGetRateLimitCmd)
	prefsGetCmd.AddCommand(prefsGetBurstCapCmd)

	prefsCmd.AddCommand(prefsSetCmd)

	prefsSetCmd.AddCommand(prefsSetRateLimitCmd)
	prefsSetCmd.AddCommand(prefsSetBurstCapCmd)

	prefsCmd.AddCommand(prefsPathCmd)

	prefsCmd.AddCommand(prefsKeyringPathCmd)
}
