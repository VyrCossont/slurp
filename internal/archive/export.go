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
	"encoding/json"
	"errors"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/time/rate"

	"github.com/VyrCossont/slurp/client/accounts"
	"github.com/VyrCossont/slurp/client/statuses"
	"github.com/VyrCossont/slurp/internal/api"
	"github.com/VyrCossont/slurp/internal/auth"
	"github.com/VyrCossont/slurp/internal/own"
	"github.com/VyrCossont/slurp/internal/util"
	"github.com/VyrCossont/slurp/models"
)

// Export exports a vaguely Mastodon-compatible archive to file (actually a folder path).
func Export(
	authClient *auth.Client,
	archiveFolderPath string,
) error {
	err := checkArchiveFolder(archiveFolderPath)
	if err != nil {
		return err
	}

	account, err := own.Account(authClient)
	if err != nil {
		slog.Error("couldn't retrieve your account", "err", err)
		return err
	}

	// TODO: (Vyr) is this the only way to get the AP ID? this only works for Mastodon and GTS
	accountURI := strings.Replace(account.URL, "@", "users/", 1)
	followersURI := accountURI + "/followers"
	err = exportActor(account, accountURI, followersURI, archiveFolderPath)
	if err != nil {
		return err
	}

	// TODO: (Vyr) this could be a command-line flag
	defaultLanguage := "en"
	if account.Source != nil && account.Source.Language != "" {
		defaultLanguage = account.Source.Language
	}

	outbox := &Outbox{
		OrderedItems: nil,
	}

	// Retrieve all of the account's statuses, oldest to newest.
	// TODO: (Vyr) this reads every single status into RAM at once. Do better.
	// TODO: (Vyr) break this beast up into chunks
	pagedRequester := &accountStatusesPagedRequester{
		accountID:     account.ID,
		forwardPaging: true,
	}
	accountStatuses, err := api.ReadAllPaged(authClient, pagedRequester, nil, util.Ptr("0"))
	if err != nil {
		slog.Error("couldn't retrieve your account's statuses", "err", err)
		return err
	}

	// Create a media attachments folder.
	attachmentsFolder := path.Join(archiveFolderPath, "media_attachments")
	if err = os.MkdirAll(attachmentsFolder, 0755); err != nil {
		slog.Error("couldn't create attachments folder", "path", attachmentsFolder, "err", err)
		return err
	}

	// Create an emojis folder.
	emojisFolder := path.Join(archiveFolderPath, "custom_emojis")
	if err = os.MkdirAll(emojisFolder, 0755); err != nil {
		slog.Error("couldn't create emojis folder", "path", emojisFolder, "err", err)
		return err
	}
	// Map of shortcode (no colons) to cache value.
	emojiCache := map[string]*emojiCacheValue{}

	// TODO: (Vyr) make this configurable thru command-line flags
	mediaDownloadLimiter := rate.NewLimiter(1, 1)
	mediaDownloadClient := util.HttpClient

	// Convert each status into an ActivityPub object, and download their attachments.
StatusesLoop:
	for _, status := range accountStatuses {
		if status.Reblog != nil {
			// TODO: (Vyr) implement boosts
			slog.Error("export doesn't handle boosts yet, skipping", "status", status.URI)
			continue StatusesLoop
		}
		if status.Poll != nil {
			// TODO: (Vyr) implement polls
			slog.Error("export doesn't handle polls yet, skipping", "status", status.URI)
			continue StatusesLoop
		}
		if status.InteractionPolicy != nil {
			// TODO: (Vyr) implement interaction policies. check GTS AP serializer.
			slog.Warn("export can't handle interaction policies correctly, continuing anyway", "status", status.URI)
		}
		if status.LocalOnly {
			// TODO: (Vyr) this is a fun puzzle since non-federated statuses may not *have* an AP serialization.
			slog.Warn("export can't handle local-only statuses correctly, continuing anyway", "status", status.URI)
		}

		object := Object{
			Id:        status.URI,
			Type:      "Note",
			Url:       status.URL,
			Sensitive: status.Sensitive,
			Content:   status.Content,
		}
		if status.SpoilerText != "" {
			object.Summary = &status.SpoilerText
		}
		if object.Published, err = time.Parse(time.RFC3339Nano, status.CreatedAt); err != nil {
			slog.Error("couldn't parse status creation time, skipping", "status", status.URI, "err", err)
			continue StatusesLoop
		}
		language := status.Language
		if language == "" {
			language = defaultLanguage
		}
		object.ContentMap = map[string]string{
			language: status.Content,
		}

		switch Visibility(status.Visibility) {
		case VisibilityPublic:
			object.To = append(object.To, ASPublic)
			object.Cc = append(object.Cc, followersURI)
		case VisibilityUnlisted:
			object.To = append(object.To, followersURI)
			object.Cc = append(object.Cc, ASPublic)
		case VisibilityPrivate:
			for _, mention := range status.Mentions {
				object.To = append(object.To, mention.URL)
			}
			object.Cc = append(object.Cc, followersURI)
		case VisibilityDirect:
			for _, mention := range status.Mentions {
				object.To = append(object.To, mention.URL)
			}
			object.Cc = []string{}
			object.DirectMessage = true
		default:
			// TODO: (Vyr) check how GTS handles mutuals-only
			slog.Warn("export doesn't support this visibility level, skipping", "status", status.URI, "visibility", status.Visibility)
			continue
		}

		if status.InReplyToID != "" {
			// Fetch the replied-to status so we can get its URI.
			if err = authClient.Wait(); err != nil {
				return err
			}

			var response *statuses.StatusGetOK
			response, err = authClient.Client.Statuses.StatusGet(
				&statuses.StatusGetParams{
					ID: status.InReplyToID,
				},
				authClient.Auth,
			)
			if err != nil {
				slog.Warn("couldn't get status being replied to (it may have been deleted), continuing anyway", "status", status.URI, "inReplyToID", status.InReplyToID, "err", err)
				continue StatusesLoop
			}

			object.InReplyTo = &response.GetPayload().URI
		}

		// Add tags for mentions.
		for _, mention := range status.Mentions {
			tag := &MentionOrHashtag{
				Type: "Mention",
				Name: "@" + mention.Acct,
				Href: mention.URL,
			}
			var rawTag json.RawMessage
			rawTag, err = json.MarshalIndent(tag, "          ", "  ")
			if err != nil {
				slog.Error("couldn't serialize status mention as JSON (should never happen)", "status", status.URI, "mention", mention.Acct, "err", err)
				return err
			}
			object.RawTags = append(object.RawTags, rawTag)
		}

		// Add tags for hashtags.
		for _, hashtag := range status.Tags {
			tag := &MentionOrHashtag{
				Type: "Hashtag",
				Name: "#" + hashtag.Name,
				Href: hashtag.URL,
			}
			var rawTag json.RawMessage
			rawTag, err = json.MarshalIndent(tag, "          ", "  ")
			if err != nil {
				slog.Error("couldn't serialize status hashtag as JSON (should never happen)", "status", status.URI, "hashtag", hashtag.Name, "err", err)
				return err
			}
			object.RawTags = append(object.RawTags, rawTag)
		}

		// Download any emojis and add tags for them.
		for _, emoji := range status.Emojis {
			cached, ok := emojiCache[emoji.Shortcode]
			if !ok {
				// Go fetch it.
				var localPath, contentType string
				localPath, contentType, err = DownloadAttachment(
					context.TODO(),
					mediaDownloadLimiter,
					mediaDownloadClient,
					status.URI,
					emojisFolder,
					emoji.URL,
				)
				if err != nil {
					// TODO: (Vyr) DownloadAttachment has detailed error messages already, but they all say "attachment" instead of "emoji". lol. refactor.
					slog.Error("couldn't download an emoji, skipping emoji and continuing anyway", "status", status.URI, "emoji", emoji.Shortcode, "err", err)
					continue
				}

				// Construct the relative path, which is the "URL" we include in the export.
				var relPath string
				relPath, err = filepath.Rel(emojisFolder, localPath)
				if err != nil {
					slog.Error("couldn't derive emoji folder relative path for emoji (should never happen)", "status", status.URI, "emoji", emoji.Shortcode, "path", localPath, "err", err)
					return err
				}

				cached = &emojiCacheValue{
					relPath:     relPath,
					contentType: contentType,
				}
				emojiCache[emoji.Shortcode] = cached
			}

			tag := &Emoji{
				Type: "Emoji",
				Name: ":" + emoji.Shortcode + ":",
				Icon: Icon{
					Type:      "Image",
					MediaType: cached.contentType,
					Url:       cached.relPath,
				},
			}
			var rawTag json.RawMessage
			rawTag, err = json.MarshalIndent(tag, "          ", "  ")
			if err != nil {
				slog.Error("couldn't serialize status emoji as JSON (should never happen)", "status", status.URI, "emoji", emoji.Shortcode, "err", err)
				return err
			}
			object.RawTags = append(object.RawTags, rawTag)
		}

		// Download any attachments.
		statusAttachmentsFolder := path.Join(attachmentsFolder, status.ID)
		if len(status.MediaAttachments) > 0 {
			if err = os.MkdirAll(statusAttachmentsFolder, 0755); err != nil {
				slog.Error("couldn't create status attachment folder", "status", status.URI, "path", statusAttachmentsFolder, "err", err)
				return err
			}
		}
		for _, attachment := range status.MediaAttachments {
			var localPath, contentType string
			localPath, contentType, err = DownloadAttachment(
				context.TODO(),
				mediaDownloadLimiter,
				mediaDownloadClient,
				status.URI,
				statusAttachmentsFolder,
				attachment.URL,
			)
			if err != nil {
				slog.Error("couldn't download an attachment, skipping attachment and continuing anyway", "status", status.URI, "attachment", attachment.URL, "err", err)
				continue
			}

			// Construct the relative path, which is the "URL" we include in the export.
			var relPath string
			relPath, err = filepath.Rel(archiveFolderPath, localPath)
			if err != nil {
				slog.Error("couldn't derive archive folder relative path for attachment (should never happen)", "status", status.URI, "attachment", attachment.URL, "path", localPath, "err", err)
				return err
			}

			apAttachment := &Attachment{
				Type:      "Document",
				MediaType: contentType,
				Url:       relPath,
			}
			if attachment.Description != "" {
				apAttachment.Name = &attachment.Description
			}
			if attachment.Meta != nil && attachment.Meta.Focus != nil {
				// TODO: (Vyr) check how audio/video get handled in Mastodon archives. do we leave the focus array empty, or skip it entirely?
				apAttachment.RawFocalPoint = append(
					apAttachment.RawFocalPoint,
					float64(attachment.Meta.Focus.X),
					float64(attachment.Meta.Focus.Y),
				)
			}
			object.Attachments = append(object.Attachments, apAttachment)
		}

		// Create a wrapper activity for the object.
		activity := Activity{
			Type: "Create",
		}
		activity.RawObject, err = json.MarshalIndent(object, "        ", "  ")
		if err != nil {
			slog.Error("couldn't serialize AP object for status as JSON (should never happen)", "status", status.URI, "err", err)
			return err
		}
		outbox.OrderedItems = append(outbox.OrderedItems, activity)
	}

	return util.SaveJSON(path.Join(archiveFolderPath, "outbox.json"), outbox)

	// TODO: (Vyr) write out the emoji cache in slurp emojis format
}

type emojiCacheValue struct {
	relPath     string
	contentType string
}

func checkArchiveFolder(archiveFolderPath string) error {
	// If the folder doesn't exist, create it.
	stat, err := os.Stat(archiveFolderPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(archiveFolderPath, 0755)
			if err != nil {
				slog.Error("archive folder didn't exist, failed to create it", "path", archiveFolderPath, "err", err)
				return err
			}
		} else {
			slog.Error("couldn't get archive info from filesystem", "path", archiveFolderPath, "err", err)
			return err
		}
	} else if !stat.IsDir() {
		return errors.New("archive folder path is already occupied by something that isn't a folder")
	}

	// Check that it's empty.
	empty, err := util.IsEmpty(archiveFolderPath)
	if err != nil {
		slog.Error("couldn't check whether archive folder is empty", "path", archiveFolderPath, "err", err)
		return err
	}
	if !empty {
		return errors.New("archive folder must be empty")
	}

	return nil
}

func exportActor(
	account *models.Account,
	accountURI string,
	followersURI string,
	archiveFolderPath string,
) error {
	actor := &Actor{
		Id:                accountURI,
		Followers:         followersURI,
		Outbox:            "outbox.json",
		PreferredUsername: account.Username,
		Url:               account.URL,
	}
	return util.SaveJSON(path.Join(archiveFolderPath, "actor.json"), actor)
}

type accountStatusesPagedRequester struct {
	accountID     string
	forwardPaging bool
}

func (pagedRequester *accountStatusesPagedRequester) Request(authClient *auth.Client, maxID *string, minID *string) (*accountStatusesPagedResponse, error) {
	resp, err := authClient.Client.Accounts.AccountStatuses(&accounts.AccountStatusesParams{
		ID:    pagedRequester.accountID,
		MaxID: maxID,
		MinID: minID,
	}, authClient.Auth)
	if err != nil {
		return nil, err
	}
	return &accountStatusesPagedResponse{resp}, nil
}

func (pagedRequester *accountStatusesPagedRequester) ForwardPaging() bool {
	return pagedRequester.forwardPaging
}

type accountStatusesPagedResponse struct {
	resp *accounts.AccountStatusesOK
}

func (pagedResponse *accountStatusesPagedResponse) Link() string {
	return pagedResponse.resp.Link
}

func (pagedResponse *accountStatusesPagedResponse) Elements() []*models.Status {
	return pagedResponse.resp.GetPayload()
}
