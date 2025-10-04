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
	"github.com/VyrCossont/slurp/internal/bluesky"
	"github.com/spf13/cobra"
)

// blueskyCmd represents the bluesky command
var blueskyCmd = &cobra.Command{
	Use:   "bluesky",
	Short: "Bluesky commands",
}

// blueskyDownloadCmd represents the bluesky download command
var blueskyDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a Bluesky account's post repo and media blobs",
	RunE: func(cmd *cobra.Command, args []string) error {
		return bluesky.Download(User, File)
	},
}

func init() {
	rootCmd.AddCommand(blueskyCmd)

	blueskyDownloadCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "path to save repo and blobs to (must be an empty folder, will be created if it doesn't exist)")
	blueskyCmd.AddCommand(blueskyDownloadCmd)
}
