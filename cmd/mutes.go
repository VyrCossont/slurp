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
	"github.com/VyrCossont/slurp/internal/mutes"
)

// mutesCmd represents the mutes command
var mutesCmd = &cobra.Command{
	Use:   "mutes",
	Short: "Import and export mutes",
}

// mutesExportCmd represents the mutes export command
var mutesExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a list of mutes",
	RunE: func(cmd *cobra.Command, args []string) error {
		authClient, err := auth.NewAuthClient(User)
		if err != nil {
			return err
		}
		return mutes.Export(authClient, File)
	},
}

// mutesImportCmd represents the mutes import command
var mutesImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a list of mutes",
	RunE: func(cmd *cobra.Command, args []string) error {
		authClient, err := auth.NewAuthClient(User)
		if err != nil {
			return err
		}
		return mutes.Import(authClient, File)
	},
}

func init() {
	rootCmd.AddCommand(mutesCmd)

	mutesExportCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "path to export mutes to (optional: stdout will be used if omitted)")
	mutesCmd.AddCommand(mutesExportCmd)

	mutesImportCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "path to import mutes from (optional: stdin will be used if omitted)")
	mutesCmd.AddCommand(mutesImportCmd)
}
