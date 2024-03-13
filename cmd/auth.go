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
		return auth.Login(User)
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

// authWhoami represents the auth whoami command
var authWhoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Display the default currently authenticated user, if there is one",
	RunE: func(cmd *cobra.Command, args []string) error {
		return auth.Whoami()
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	authCmd.AddCommand(authLoginCmd)

	authCmd.AddCommand(authLogoutCmd)
}
