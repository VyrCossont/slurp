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
	"github.com/VyrCossont/slurp/internal/bookmarks"
)

// bookmarksCmd represents the bookmarks command
var bookmarksCmd = &cobra.Command{
	Use:   "bookmarks",
	Short: "Import and export bookmarks",
}

// bookmarksExportCmd represents the bookmarks export command
var bookmarksExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a list of bookmarks",
	RunE: func(cmd *cobra.Command, args []string) error {
		authClient, err := auth.NewAuthClient(User)
		if err != nil {
			return err
		}
		return bookmarks.Export(authClient, File)
	},
}

// bookmarksImportCmd represents the bookmarks import command
var bookmarksImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a list of bookmarks",
	RunE: func(cmd *cobra.Command, args []string) error {
		authClient, err := auth.NewAuthClient(User)
		if err != nil {
			return err
		}
		return bookmarks.Import(authClient, File)
	},
}

func init() {
	rootCmd.AddCommand(bookmarksCmd)

	bookmarksExportCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "path to export bookmarks to (optional: stdout will be used if omitted)")
	bookmarksCmd.AddCommand(bookmarksExportCmd)

	bookmarksImportCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "path to import bookmarks from (optional: stdin will be used if omitted)")
	bookmarksCmd.AddCommand(bookmarksImportCmd)
}
