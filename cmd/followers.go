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
	"github.com/VyrCossont/slurp/internal/followers"
)

// followersCmd represents the followers command
var followersCmd = &cobra.Command{
	Use:   "followers",
	Short: "Export followers",
}

// followersExportCmd represents the followers export command
var followersExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a list of followers",
	RunE: func(cmd *cobra.Command, args []string) error {
		keyring, err := auth.ClientKeyring()
		if err != nil {
			return err
		}

		authClient, err := auth.NewAuthClient(User, keyring)
		if err != nil {
			return err
		}

		return followers.Export(authClient, File)
	},
}

func init() {
	rootCmd.AddCommand(followersCmd)

	followersExportCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "path to export followers to (optional: stdout will be used if omitted)")
	followersCmd.AddCommand(followersExportCmd)
}
