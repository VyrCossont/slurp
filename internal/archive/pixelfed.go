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
	"cmp"
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"path"
	"slices"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"golang.org/x/time/rate"

	"github.com/VyrCossont/slurp/client/statuses"
	"github.com/VyrCossont/slurp/internal/auth"
	"github.com/VyrCossont/slurp/internal/util"
	"github.com/VyrCossont/slurp/models"
)

func PixelfedImport(
	authClient *auth.Client,
	file string,
	statusMapFile string,
	attachmentMapFile string,
	attachmentDirectory string,
) error {
	archiveIdToImportedApiId, mediaPathToImportedApiId, err := requireMapFiles(statusMapFile, attachmentMapFile)
	if err != nil {
		return err
	}

	// TODO: (Vyr) make this configurable thru command-line flags
	mediaDownloadLimiter := rate.NewLimiter(1, 1)
	mediaDownloadClient := util.HttpClient

	if attachmentDirectory == "" {
		return errors.New("attachment directory is required for Pixelfed imports")
	}
	err = os.MkdirAll(attachmentDirectory, 0755)
	if err != nil {
		slog.Error("Couldn't create attachment directory", "path", attachmentDirectory, "err", err)
		return err
	}

	statusesToImport, err := pixelfedReadStatuses(file)
	if err != nil {
		slog.Error("Couldn't read Pixelfed statuses from JSON file", "path", file, "err", err)
		return err
	}

	// Assume the same user posted all the archive statuses,
	// which should be a fairly safe assumption.
	if len(statusesToImport) == 0 {
		return errors.New("no statuses to import")
	}
	archiveOwner := statusesToImport[0].Account

	// Map of archive API ID to status.
	statusGraph := map[string]*PixelfedArchiveStatus{}
	for _, status := range statusesToImport {
		statusGraph[status.ID] = status
	}

	// Put statuses in order so we do things consistently. Only the topo sort is required for correctness,
	// but the rest is a nice property when running the same job again and again while debugging.
	slices.SortFunc(statusesToImport, func(a, b *PixelfedArchiveStatus) int {
		// Topo sort by reply graph.
		if a.ID == b.InReplyToID {
			return -1
		} else if a.InReplyToID == b.ID {
			return 1
		}

		// Sort by published date.
		var err error
		aCreatedAt, err := time.Parse(time.RFC3339Nano, a.CreatedAt)
		if err != nil {
			slog.Error("Couldn't parse Pixelfed archive status created time", "status", a.ID, "created_at", a.CreatedAt)
			return 0
		}
		bCreatedAt, err := time.Parse(time.RFC3339Nano, b.CreatedAt)
		if err != nil {
			slog.Error("Couldn't parse Pixelfed archive status created time", "status", b.ID, "created_at", b.CreatedAt)
			return 0
		}

		publishedCmp := aCreatedAt.Compare(bCreatedAt)
		if publishedCmp != 0 {
			return publishedCmp
		}

		// Pixelfed API IDs are probably snowflake numeric IDs and sorting them as strings may not always be useful,
		// but after date sorting, it's usually okay.
		return cmp.Compare(a.ID, b.ID)
	})

StatusesLoop:
	for _, status := range statusesToImport {
		// Skip previously imported statuses.
		if _, previouslyImported := archiveIdToImportedApiId[status.ID]; previouslyImported {
			continue StatusesLoop
		}

		visibility := Visibility(status.Visibility)
		// TODO: (Vyr) command-line visibility mapping
		if visibility != VisibilityPublic && visibility != VisibilityUnlisted {
			slog.Info("Skipping status due to visibility", "status", status.URI, "visibility", string(visibility))
			continue StatusesLoop
		}

		// Allow OPs or self-replies to our other posts only.
		if status.InReplyToAccountID != "" && status.InReplyToAccountID != archiveOwner.ID {
			slog.Info("Skipping status because it's a reply to another account'", "status", status.URI)
			continue StatusesLoop
		}
		if status.InReplyToID != "" && statusGraph[status.InReplyToID] == nil {
			slog.Info("Skipping status because it isn't a top-level status or self-reply", "status", status.URI)
			continue StatusesLoop
		}
		if len(status.Mentions) > 0 {
			// TODO: (Vyr) doesn't handle edge case where you explicitly mention yourself
			slog.Info("Skipping status because it mentions other accounts", "status", status.URI, "mentions", status.Mentions)
			continue StatusesLoop
		}

		// Skip statuses with polls.
		if status.Poll != nil {
			slog.Info("Skipping status because it has a poll", "status", status.URI)
			continue StatusesLoop
		}

		// Mentions from deleted instances or users don't show up as mention tags. Skip those too.
		// This transform allows the common case of URLs that contain @ but are probably links to Fedi posts.
		// TODO: (Vyr) Will give false positives on code blocks, Bluesky handles, mentions that never resolved, etc.
		contentText := ""
		if status.ContentText != nil {
			contentText = *status.ContentText
		}
		atLinkPlaceholderText := atLinkPattern.ReplaceAllString(contentText, "at_link_placeholder")
		if mentionPattern.MatchString(atLinkPlaceholderText) {
			slog.Info("Skipping status because it looks like it has a mention of a dead account", "status", status.URI)
			continue StatusesLoop
		}

		// Get the imported API ID of the post we're replying to, if applicable.
		var inReplyToApiId *string
		if status.InReplyToID != "" {
			if apiId, found := archiveIdToImportedApiId[status.InReplyToID]; found {
				inReplyToApiId = &apiId
			} else {
				slog.Error("Couldn't find API ID for status being replied to", "status", status.URI, "inReplyToID", status.InReplyToID)
				continue StatusesLoop
			}
		}

		// Create a directory in which to put media attachments for this status, if it has any.
		localDir := path.Join(attachmentDirectory, status.ID)
		if len(status.MediaAttachments) > 0 {
			if err := os.MkdirAll(localDir, 0755); err != nil {
				slog.Error("Couldn't create attachment directory", "status", status.URI, "localDir", localDir, "err", err)
				continue StatusesLoop
			}
		}

		// Download and then upload any media attachments.
		mediaIDs := make([]string, 0, len(status.MediaAttachments))
		for _, attachment := range status.MediaAttachments {
			localPath, _, err := DownloadAttachment(
				context.TODO(),
				mediaDownloadLimiter,
				mediaDownloadClient,
				status.URI,
				localDir,
				attachment.URL,
			)
			if err != nil {
				continue StatusesLoop
			}

			var focusX *float64
			var focusY *float64
			if attachment.Meta != nil && attachment.Meta.Focus != nil {
				focusX = util.Ptr(float64(attachment.Meta.Focus.X))
				focusY = util.Ptr(float64(attachment.Meta.Focus.Y))
			}
			mediaID, err := uploadAttachment(
				authClient,
				attachmentMapFile,
				mediaPathToImportedApiId,
				attachment.URL,
				localPath,
				&attachment.Description,
				focusX,
				focusY,
			)
			if err != nil {
				slog.Error("Error retrieving or uploading attachment", "status", status.URI, "attachment", attachment.URL, "err", err)
				continue StatusesLoop
			}
			mediaIDs = append(mediaIDs, mediaID)
		}

		if err := authClient.Wait(); err != nil {
			return err
		}

		createParams, err := status.CreateParams(inReplyToApiId, mediaIDs)
		if err != nil {
			continue StatusesLoop
		}

		response, err := authClient.Client.Statuses.StatusCreate(
			createParams,
			authClient.Auth,
			func(op *runtime.ClientOperation) {
				op.ConsumesMediaTypes = []string{"application/x-www-form-urlencoded"}
			},
		)
		if err != nil {
			slog.Error("Couldn't post converted status", "status", status.URI, "err", err)
			continue StatusesLoop
		}
		importedStatus := response.GetPayload()

		// Save the API ID of the status we just imported for future imports.
		archiveIdToImportedApiId[status.ID] = importedStatus.ID
		if err := writeMapFile(statusMapFile, archiveIdToImportedApiId); err != nil {
			slog.Error("Couldn't write status map file", "path", statusMapFile, "err", err)
			return err
		}

		slog.Info("Imported status", "status", status.URI, "url", status.URL)
	}

	return nil
}

// PixelfedArchiveStatus is a Mastodon API status with a few extra fields.
type PixelfedArchiveStatus struct {
	models.Status

	// ContentText is probably the Pixelfed version of `source`.
	// If this is missing, the `content` will be an empty string.
	ContentText *string `json:"content_text"`

	// Place is an optional location.
	Place *PixelfedPlace `json:"place"`

	// CommentsDisabled surfaces Pixelfed's own interaction controls.
	CommentsDisabled bool `json:"comments_disabled"`
}

func (p *PixelfedArchiveStatus) CreateParams(inReplyToID *string, mediaIDs []string) (*statuses.StatusCreateParams, error) {
	createdAt, err := time.Parse(time.RFC3339Nano, p.CreatedAt)
	if err != nil {
		return nil, errors.Errorf("couldn't parse Pixelfed archive status created time: %s", p.CreatedAt)
	}

	s := &statuses.StatusCreateParams{
		ContentType: util.Ptr("text/plain"),
		InReplyToID: inReplyToID,
		Language:    &p.Language,
		MediaIDs:    mediaIDs,
		ScheduledAt: util.Ptr(strfmt.DateTime(createdAt)),
		Sensitive:   &p.Sensitive,
		SpoilerText: &p.SpoilerText,
		Status:      p.ContentText,
		Visibility:  &p.Visibility,
	}

	if p.CommentsDisabled {
		s.InteractionPolicyCanReplyAlways0 = util.Ptr("me")
	}

	return s, nil
}

// PixelfedPlace describes a geotagged location.
// TODO: (Vyr) not currently used. Also, figure out the full set of fields.
type PixelfedPlace struct {
	ID      int    `json:"id"`
	Slug    string `json:"slug"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func pixelfedReadStatuses(file string) ([]*PixelfedArchiveStatus, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer func() { _ = jsonFile.Close() }()

	var doc []*PixelfedArchiveStatus
	if err := json.NewDecoder(jsonFile).Decode(&doc); err != nil {
		return nil, err
	}
	return doc, nil
}
