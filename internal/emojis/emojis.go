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

package emojis

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/VyrCossont/slurp/client/admin"
	"github.com/VyrCossont/slurp/internal/auth"
	"github.com/VyrCossont/slurp/internal/own"
	"github.com/VyrCossont/slurp/models"
	"github.com/go-openapi/runtime"
	"github.com/pkg/errors"
	"github.com/vincent-petithory/dataurl"
)

func Export(authClient *auth.Client, file string, inline bool) error {
	emojis, err := own.Emojis(authClient)
	if err != nil {
		return err
	}

	if inline {
		for _, emoji := range emojis {
			emojiBytes, mimeType, err := getEmojiHttps(emoji.URL)
			if err != nil {
				return err
			}
			emoji.URL = dataurl.New(emojiBytes, mimeType).String()
			emoji.StaticURL = ""
		}
	}

	out := os.Stdout
	if file != "" {
		out, err = os.Create(file)
		if err != nil {
			slog.Error("couldn't create emojis file", "file", file, "error", err)
			return errors.WithStack(err)
		}
		defer func() {
			if err := out.Close(); err != nil {
				slog.Error("error closing emojis file", "file", file, "error", err)
			}
		}()
	}

	encoder := json.NewEncoder(out)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(emojis); err != nil {
		slog.Error("error writing emojis to file", "file", file, "error", err)
		return err
	}

	return nil
}

func Import(authClient *auth.Client, file string) error {
	var emojis []*models.Emoji

	var err error
	in := os.Stdin
	if file != "" {
		in, err = os.Open(file)
		if err != nil {
			slog.Error("couldn't open emojis file", "error", err)
			return errors.WithStack(err)
		}
		defer func() {
			if err := in.Close(); err != nil {
				slog.Error("error closing emojis file", "file", file, "error", err)
			}
		}()
	}

	if err := json.NewDecoder(in).Decode(&emojis); err != nil {
		slog.Error("couldn't read from emojis file", "error", err)
		return errors.WithStack(err)
	}

	for _, emoji := range emojis {
		var emojiBytes []byte
		var mimeType string

		urlParts, err := url.Parse(emoji.URL)
		if err != nil {
			return errors.WithStack(err)
		}
		switch urlParts.Scheme {
		case "https":
			emojiBytes, mimeType, err = getEmojiHttps(emoji.URL)
			if err != nil {
				return errors.WithStack(err)
			}

		case "data":
			data, err := dataurl.DecodeString(emoji.URL)
			if err != nil {
				slog.Error("couldn't decode data: URL from emojis file", "shortcode", emoji.Shortcode, "error", err)
				return errors.WithStack(err)
			}
			emojiBytes = data.Data
			mimeType = data.ContentType()

		default:
			err := errors.New("unknown scheme for URL from emojis file")
			slog.Error("unknown scheme for URL from emojis file", "shortcode", emoji.Shortcode, "scheme", urlParts.Scheme, "error", err)
			return errors.WithStack(err)
		}

		// Recreate a filename for the upload.
		exts, err := mime.ExtensionsByType(mimeType)
		if err != nil {
			slog.Error("couldn't get extensions for MIME type", "shortcode", emoji.Shortcode, "mimeType", mimeType, "error", err)
			return errors.WithStack(err)
		}
		if len(exts) == 0 {
			err := errors.New("no known extensions for MIME type")
			slog.Error("no known extensions for MIME type", "shortcode", emoji.Shortcode, "mimeType", mimeType, "error", err)
			return errors.WithStack(err)
		}
		uploadFilename := emoji.Shortcode + exts[0]

		response, err := authClient.Client.Admin.EmojiCreate(
			&admin.EmojiCreateParams{
				Category:  &emoji.Category,
				Image:     runtime.NamedReader(uploadFilename, bytes.NewReader(emojiBytes)),
				Shortcode: emoji.Shortcode,
			},
			authClient.Auth,
			func(op *runtime.ClientOperation) {
				op.ConsumesMediaTypes = []string{"multipart/form-data"}
			},
		)
		if err != nil {
			slog.Error("couldn't create emoji", "shortcode", emoji.Shortcode, "error", err)
			return errors.WithStack(err)
		}
		apiEmoji := response.GetPayload()

		slog.Info("imported emoji", "shortcode", apiEmoji.Shortcode)
	}

	return nil
}

func getEmojiHttps(url string) ([]byte, string, error) {
	response, err := http.Get(url)
	if err != nil {
		slog.Error("error getting emoji from source", "url", url, "error", err)
		return nil, "", errors.WithStack(err)
	}

	mimeType := response.Header.Get("Content-Type")
	if mimeType == "" {
		mimeType = mime.TypeByExtension(path.Ext(url))
	}
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}

	emojiBytes, err := io.ReadAll(response.Body)
	if err != nil {
		slog.Error("error reading emoji bytes from source", "url", url, "error", err)
		return nil, "", errors.WithStack(err)
	}

	return emojiBytes, mimeType, nil
}
