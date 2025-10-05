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
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/VyrCossont/slurp/internal/archive"
	"github.com/VyrCossont/slurp/internal/auth"
)

// archiveCmd represents the archive command
var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Import and export post archives",
}

// archiveImportCmd represents the archive import command
var archiveImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a post archive",
	RunE: func(cmd *cobra.Command, args []string) error {
		keyring, err := auth.ClientKeyring()
		if err != nil {
			return err
		}

		authClient, err := auth.NewAuthClient(User, keyring)
		if err != nil {
			return err
		}

		switch Format {
		case "", "mastodon":
			return archive.Import(authClient, File, StatusMapFile, AttachmentMapFile, AllowMissingCustomEmojis)

		case "pixelfed":
			return archive.PixelfedImport(authClient, File, StatusMapFile, AttachmentMapFile, AttachmentDirectory)

		case "bluesky":
			return archive.BlueskyImport(authClient, File, StatusMapFile, AttachmentMapFile)

		default:
			return errors.Errorf("unknown archive format: %s", Format)
		}
	},
}

// archiveExportCmd represents the archive import command
var archiveExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a post archive",
	RunE: func(cmd *cobra.Command, args []string) error {
		keyring, err := auth.ClientKeyring()
		if err != nil {
			return err
		}

		authClient, err := auth.NewAuthClient(User, keyring)
		if err != nil {
			return err
		}

		return archive.Export(authClient, File)
	},
}

// Format is the file path for recording how archive media attachment paths map to imported media attachment IDs.
var Format string

// StatusMapFile is the file path for recording how archive status IDs map to imported status IDs.
var StatusMapFile string

// AttachmentMapFile is the file path for recording how archive media attachment paths map to imported media attachment IDs.
var AttachmentMapFile string

// AttachmentDirectory is the folder for attachments downloaded from Pixelfed during the course of an import.
var AttachmentDirectory string

// AllowMissingCustomEmojis allows importing statuses for which the instance doesn't have a matching custom emoji.
var AllowMissingCustomEmojis bool

func init() {
	rootCmd.AddCommand(archiveCmd)

	archiveImportCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "path to import archive from (this must be an uncompressed folder for Mastodon)")
	archiveImportCmd.PersistentFlags().StringVarP(&Format, "format", "t", "", "archive format can be one of: mastodon, pixelfed (default is mastodon)")
	archiveImportCmd.PersistentFlags().StringVarP(&StatusMapFile, "status-map-file", "m", "", "JSON file to store mapping of archive status IDs to imported status IDs")
	archiveImportCmd.PersistentFlags().StringVarP(&AttachmentMapFile, "attachment-map-file", "a", "", "JSON file to store mapping of archive media attachment paths IDs to media attachment IDs")
	archiveImportCmd.PersistentFlags().StringVarP(&AttachmentDirectory, "attachment-directory", "d", "", "folder to store media downloaded from Pixelfed")
	archiveImportCmd.PersistentFlags().BoolVarP(&AllowMissingCustomEmojis, "allow-missing-custom-emojis", "e", false, "import statuses for which the instance doesn't have all of the custom emojis")
	archiveCmd.AddCommand(archiveImportCmd)

	archiveExportCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "path to import archive from (this must be an empty folder; if it doesn't exist, it'll be created)")
	archiveCmd.AddCommand(archiveExportCmd)
}
