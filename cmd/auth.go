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
	"fmt"

	"github.com/spf13/cobra"

	"github.com/VyrCossont/slurp/internal/auth"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Log in or out",
}

// authLoginCmd represents the auth login command
var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in",
	RunE: func(cmd *cobra.Command, args []string) error {
		return auth.Login(User, AllowHTTP, UseCleartextFileKeyring)
	},
}

// authLogoutCmd represents the auth logout command
var authLogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out",
	RunE: func(cmd *cobra.Command, args []string) error {
		return auth.Logout(User)
	},
}

// authWhoamiCmd represents the auth whoami command
var authWhoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Display the default currently authenticated user, if there is one",
	RunE: func(cmd *cobra.Command, args []string) error {
		user, err := auth.Whoami()
		if err != nil {
			return err
		}

		_, err = cmd.OutOrStderr().Write([]byte(fmt.Sprintf("%s\n", user)))
		return err
	},
}

// authSwitchCmd represents the auth switch command
var authSwitchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Change the default user",
	RunE: func(cmd *cobra.Command, args []string) error {
		return auth.Switch(User)
	},
}

// AllowHTTP allows cleartext HTTP instead of HTTPS.
// Should be used for local testing only.
var AllowHTTP bool

// UseCleartextFileKeyring uses a cleartext file-backed keyring instead of a system keyring.
// Should be used only if you're
var UseCleartextFileKeyring bool

func init() {
	rootCmd.AddCommand(authCmd)

	authLoginCmd.PersistentFlags().BoolVarP(&AllowHTTP, "allow-http", "x", false, "allow cleartext HTTP for user's instance (should be used for local testing only)")
	authLoginCmd.PersistentFlags().BoolVarP(&UseCleartextFileKeyring, "use-cleartext-file-keyring", "k", false, "use cleartext file keyring instead of system keyring (should be used only if your system doesn't provide a keyring)")
	authCmd.AddCommand(authLoginCmd)

	authCmd.AddCommand(authLogoutCmd)

	authCmd.AddCommand(authWhoamiCmd)

	authCmd.AddCommand(authSwitchCmd)
}
