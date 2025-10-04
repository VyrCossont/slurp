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

package bluesky

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path"
	"strings"

	"github.com/VyrCossont/slurp/internal/util"
	"github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/atproto/identity"
	"github.com/bluesky-social/indigo/atproto/syntax"
	"github.com/bluesky-social/indigo/xrpc"
	"github.com/pkg/errors"
)

// Download downloads a user's post repo and media blobs into a new or empty folder.
// Copied from <https://docs.bsky.app/blog/repo-export>.
func Download(user string, archiveFolderPath string) error {
	var err error
	ctx := context.TODO()

	// Make sure the archive folder is ready.
	err = util.CheckArchiveFolder(archiveFolderPath)
	if err != nil {
		return err
	}

	// Identify the user and create an XRPC client talking to the right PDS.
	atid, err := parseUser(user)
	if err != nil {
		return err
	}
	ident, err := lookupIdentity(ctx, user, atid)
	if err != nil {
		return err
	}
	xrpcc := &xrpc.Client{
		Host: ident.PDSEndpoint(),
	}

	// Save the user's identity metadata.
	err = saveIdentity(archiveFolderPath, ident)
	if err != nil {
		return err
	}

	// Download the post repo.
	err = downloadRepo(ctx, user, archiveFolderPath, xrpcc, ident)
	if err != nil {
		return err
	}

	// Download the media blobs.
	err = downloadBlobs(ctx, user, archiveFolderPath, xrpcc, ident)
	if err != nil {
		return err
	}

	return nil
}

// parseUser converts a Bluesky username without leading @ to the identifier type used by Indigo.
func parseUser(user string) (*syntax.AtIdentifier, error) {
	if user[0] == '@' {
		return nil, errors.WithStack(errors.New("take the leading @ off the Bluesky account name and try again"))
	}
	if strings.ContainsRune(user, '@') || !strings.ContainsRune(user, '.') {
		return nil, errors.WithStack(errors.New("expected a Bluesky account name but this doesn't look like one"))
	}

	atid, err := syntax.ParseAtIdentifier(user)
	if err != nil {
		slog.Error("couldn't parse Bluesky account name", "user", user, "error", err)
		return nil, err
	}

	return atid, nil
}

// lookupIdentity looks up an identifier in Bluesky's default directory server.
func lookupIdentity(ctx context.Context, user string, atid *syntax.AtIdentifier) (*identity.Identity, error) {
	dir := identity.DefaultDirectory()
	ident, err := dir.Lookup(ctx, *atid)
	if err != nil {
		slog.Error("couldn't look up Bluesky account name in default directory", "user", user, "error", err)
		return nil, err
	}

	if ident.PDSEndpoint() == "" {
		return nil, errors.WithStack(fmt.Errorf("no PDS endpoint for identity"))
	}
	slog.Info("found PDS endpoint", "endpoint", ident.PDSEndpoint(), "user", user)

	return ident, nil
}

// IdentityJSON is the path inside a Bluesky archive folder containing identity metadata.
const IdentityJSON = "identity.json"

func saveIdentity(archiveFolderPath string, ident *identity.Identity) error {
	identityPath := path.Join(archiveFolderPath, IdentityJSON)
	return util.SaveJSON(identityPath, ident)
}

// RepoCAR is the path inside a Bluesky archive folder containing the user's post repo.
const RepoCAR = "repo.car"

// downloadRepo downloads a user's post repo into a file in the archive folder.
func downloadRepo(ctx context.Context, user string, archiveFolderPath string, xrpcc *xrpc.Client, ident *identity.Identity) error {
	carPath := path.Join(archiveFolderPath, RepoCAR)

	repoBytes, err := atproto.SyncGetRepo(ctx, xrpcc, ident.DID.String(), "")
	if err != nil {
		slog.Error("couldn't download user's post repo", "user", user, "error", err)
		return err
	}

	err = os.WriteFile(carPath, repoBytes, 0644)
	if err != nil {
		slog.Error("couldn't write user's post repo to disk", "user", user, "path", carPath, "error", err)
		return err
	}
	slog.Info("downloaded post repo", "path", carPath)

	// TODO: Bluesky says "you would probably want to confirm the commit signature using the account's signing public key (included in the resolved identity metadata)"

	return nil
}

// BlobsFolder is the subfolder inside a Bluesky archive folder containing media blobs.
const BlobsFolder = "blobs"

// BlobBatchSize is the number of blob CIDs to fetch at once.
const BlobBatchSize = 500

// downloadBlobs downloads the user's media blobs into a subfolder of the archive folder.
func downloadBlobs(ctx context.Context, user string, archiveFolderPath string, xrpcc *xrpc.Client, ident *identity.Identity) error {
	blobsPath := path.Join(archiveFolderPath, BlobsFolder)
	err := os.Mkdir(blobsPath, 0755)
	if err != nil {
		slog.Info("couldn't create blobs folder inside archive folder", "path", blobsPath, "error", err)
		return err
	}

	cursor := ""
	for {
		// Fetch a batch of blob CIDs.
		resp, err := atproto.SyncListBlobs(ctx, xrpcc, cursor, ident.DID.String(), BlobBatchSize, "")
		if err != nil {
			slog.Error("couldn't fetch batch of user's blob CIDs", "user", user, "error", err)
			return err
		}

		// Download each individual blob.
		for _, cid := range resp.Cids {
			blobBytes, err := atproto.SyncGetBlob(ctx, xrpcc, cid, ident.DID.String())
			if err != nil {
				slog.Error("couldn't download blob", "user", user, "cid", cid, "error", err)
				return err
			}

			blobPath := path.Join(blobsPath, cid)
			err = os.WriteFile(blobPath, blobBytes, 0644)
			if err != nil {
				slog.Error("couldn't write blob to disk", "user", user, "cid", cid, "path", blobPath, "error", err)
				return err
			}
			slog.Info("downloaded media blob", "path", blobPath)

			// TODO: Bluesky says "A more rigorous implementation should verify the blob CID (by hashing the downloaded bytes), at a minimum to detect corruption and errors."
		}

		// If there are more results, keep going.
		if resp.Cursor != nil && *resp.Cursor != "" {
			cursor = *resp.Cursor
		} else {
			break
		}
	}

	return nil
}
