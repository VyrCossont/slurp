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
	"github.com/VyrCossont/slurp/internal/auth"
	"github.com/VyrCossont/slurp/internal/emojis"
	"github.com/spf13/cobra"
)

// emojisCmd represents the emojis command
var emojisCmd = &cobra.Command{
	Use:   "emojis",
	Short: "Import and export emojis",
}

// emojisExportCmd represents the emojis export command
var emojisExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a list of emojis",
	RunE: func(cmd *cobra.Command, args []string) error {
		keyring, err := auth.ClientKeyring()
		if err != nil {
			return err
		}

		authClient, err := auth.NewAuthClient(User, keyring)
		if err != nil {
			return err
		}

		return emojis.Export(authClient, File, Inline)
	},
}

// emojisImportCmd represents the emojis import command
var emojisImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a list of emojis",
	RunE: func(cmd *cobra.Command, args []string) error {
		keyring, err := auth.ClientKeyring()
		if err != nil {
			return err
		}

		authClient, err := auth.NewAuthClient(User, keyring)
		if err != nil {
			return err
		}

		return emojis.Import(authClient, File)
	},
}

// Inline includes emojis in the exported emoji JSON as data: URLs.
var Inline bool

func init() {
	rootCmd.AddCommand(emojisCmd)

	emojisExportCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "path to export emojis to (optional: stdout will be used if omitted)")
	emojisExportCmd.PersistentFlags().BoolVarP(&Inline, "inline", "i", false, "inline emojis as data: URLs")
	emojisCmd.AddCommand(emojisExportCmd)

	emojisImportCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "path to import emojis from (optional: stdin will be used if omitted)")
	emojisCmd.AddCommand(emojisImportCmd)
}
