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
	"github.com/VyrCossont/slurp/internal/filters"
)

// filtersCmd represents the filters command
var filtersCmd = &cobra.Command{
	Use:   "filters",
	Short: "Import and export filters",
}

// filtersExportCmd represents the filters export command
var filtersExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a list of filters",
	RunE: func(cmd *cobra.Command, args []string) error {
		authClient, err := auth.NewAuthClient(User)
		if err != nil {
			return err
		}
		return filters.Export(authClient, File)
	},
}

// filtersImportCmd represents the filters import command
var filtersImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a list of filters",
	RunE: func(cmd *cobra.Command, args []string) error {
		authClient, err := auth.NewAuthClient(User)
		if err != nil {
			return err
		}
		return filters.Import(authClient, File)
	},
}

func init() {
	rootCmd.AddCommand(filtersCmd)

	filtersExportCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "path to export filters to (optional: stdout will be used if omitted)")
	filtersCmd.AddCommand(filtersExportCmd)

	filtersImportCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "path to import filters from (optional: stdin will be used if omitted)")
	filtersCmd.AddCommand(filtersImportCmd)
}
