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
	"github.com/VyrCossont/slurp/internal/lists"
)

// listsCmd represents the lists command
var listsCmd = &cobra.Command{
	Use:   "lists",
	Short: "Import and export lists",
}

// listsExportCmd represents the lists export command
var listsExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a list of lists",
	RunE: func(cmd *cobra.Command, args []string) error {
		authClient, err := auth.NewAuthClient(User)
		if err != nil {
			return err
		}
		return lists.Export(authClient, File)
	},
}

// listsImportCmd represents the lists import command
var listsImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a list of lists",
	RunE: func(cmd *cobra.Command, args []string) error {
		authClient, err := auth.NewAuthClient(User)
		if err != nil {
			return err
		}
		return lists.Import(authClient, File)
	},
}

func init() {
	rootCmd.AddCommand(listsCmd)

	listsExportCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "path to export lists to (optional: stdout will be used if omitted)")
	listsCmd.AddCommand(listsExportCmd)

	listsImportCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "path to import lists from (optional: stdin will be used if omitted)")
	listsCmd.AddCommand(listsImportCmd)
}
