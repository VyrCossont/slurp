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

	"github.com/VyrCossont/slurp/internal/blocks"
)

// blocksCmd represents the blocks command
var blocksCmd = &cobra.Command{
	Use:   "blocks",
	Short: "Import and export blocks",
}

// blocksExportCmd represents the blocks export command
var blocksExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a list of blocks",
	RunE: func(cmd *cobra.Command, args []string) error {
		return blocks.Export()
	},
}

// blocksImportCmd represents the blocks import command
var blocksImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a list of blocks",
	RunE: func(cmd *cobra.Command, args []string) error {
		return blocks.Import()
	},
}

func init() {
	rootCmd.AddCommand(blocksCmd)

	blocksExportCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "path to export blocks to (optional: stdout will be used if omitted)")
	blocksCmd.AddCommand(blocksExportCmd)

	blocksImportCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "path to import blocks from (optional: stdin will be used if omitted)")
	blocksCmd.AddCommand(blocksImportCmd)
}
