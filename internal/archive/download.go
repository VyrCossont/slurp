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

package archive

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path"

	"golang.org/x/time/rate"
)

// DownloadAttachment downloads an attachment to a local directory, building the full path from the remote filename.
// It returns the local path and the MIME type (or "application/octet-stream" if the remote server doesn't send one).
// It uses slog and takes the status URI for log context.
func DownloadAttachment(
	ctx context.Context,
	mediaDownloadLimiter *rate.Limiter,
	mediaDownloadClient *http.Client,
	statusURI string,
	localDir string,
	url string,
) (string, string, error) {
	localPath := path.Join(localDir, path.Base(url))
	localFile, err := os.Create(localPath)
	if err != nil {
		slog.Error("Error creating local attachment file", "status", statusURI, "attachment", url, "localPath", localPath, "err", err)
	}
	defer func() {
		if err = localFile.Close(); err != nil {
			slog.Error("Error closing local attachment file", "status", statusURI, "attachment", url, "localPath", localPath, "err", err)
		}
	}()

	// TODO: (Vyr) add media download timeout
	if err = mediaDownloadLimiter.Wait(ctx); err != nil {
		return "", "", err
	}

	// Download the attachment from the original server.
	resp, err := mediaDownloadClient.Get(url)
	if err != nil {
		slog.Error("Error downloading attachment", "status", statusURI, "attachment", url, "err", err)
		return "", "", err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			slog.Error("Error closing attachment response body", "status", statusURI, "attachment", url, "err", err)
		}
	}()
	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	if _, err = io.Copy(localFile, resp.Body); err != nil {
		slog.Error("Error copying response to local file", "status", statusURI, "attachment", url, "localPath", localPath, "err", err)
		return "", "", err
	}

	return localPath, contentType, nil
}
