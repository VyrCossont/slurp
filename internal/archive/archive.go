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
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path"
	"regexp"
	"slices"
	"strings"

	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/base"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/commonmark"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/strikethrough"
	"github.com/VyrCossont/slurp/internal/archive/mastodon"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VyrCossont/slurp/client/media"
	"github.com/VyrCossont/slurp/client/statuses"
	"github.com/VyrCossont/slurp/internal/auth"
	"github.com/VyrCossont/slurp/internal/own"
	"github.com/VyrCossont/slurp/internal/util"
)

var MentionPattern = regexp.MustCompile(`@\w+`)
var AtLinkPattern = regexp.MustCompile(`https://[a-z0-9.-]+/@\w+`)

func Import(
	authClient *auth.Client,
	file string,
	statusMapFile string,
	attachmentMapFile string,
	allowMissingCustomEmojis bool,
) error {

	archiveIdToImportedApiId, mediaPathToImportedApiId, err := RequireMapFiles(statusMapFile, attachmentMapFile)
	if err != nil {
		return err
	}

	instance, err := own.Instance(authClient)
	if err != nil {
		return err
	}
	statusContentType := "text/markdown"
	if !slices.Contains(instance.Configuration.Statuses.SupportedMimeTypes, statusContentType) {
		return errors.New("instance must support posting in Markdown")
	}
	// TODO: (Vyr) some instance types (Glitch, Akkoma) can handle text/html as well, but that's another route
	// TODO: (Vyr) could we post HTML-in-Markdown to skip conversion, or does that mangle mentions/hashtags?
	markdownConverter := converter.NewConverter(
		converter.WithPlugins(
			base.NewBasePlugin(),
			commonmark.NewCommonmarkPlugin(),
			strikethrough.NewStrikethroughPlugin(),
		),
	)

	// TODO: (Vyr) detect backfill support.	Assuming GTS 0.18 for now. Run at your own risk on anything else.

	emojis, err := own.Emojis(authClient)
	if err != nil {
		return err
	}
	// Stores shortcodes wrapped in colons.
	instanceEmojiNames := map[string]struct{}{}
	for _, emoji := range emojis {
		instanceEmojiNames[":"+emoji.Shortcode+":"] = struct{}{}
	}

	// Put statuses in order so we do things consistently. Only the topo sort is required for correctness,
	// but the rest is a nice property when running the same job again and again while debugging.
	var reader ImportableReader
	statuses := reader.Statuses()
	statusMap := make(map[string]ImportableStatus, len(statuses))
	statusesInImportOrder := make([]ImportableStatus, 0, len(statuses))
	for _, status := range statuses {
		statusMap[status.ID()] = status
		statusesInImportOrder = append(statusesInImportOrder, status)
	}
	slices.SortFunc(statusesInImportOrder, func(a, b ImportableStatus) int {
		// Topo sort by reply graph.
		if b.InReplyToID() != "" && a.ID() == b.InReplyToID() {
			return -1
		} else if a.InReplyToID() != "" && a.InReplyToID() == b.ID() {
			return 1
		}

		// Sort by published date.
		publishedCmp := a.CreatedAt().Compare(b.CreatedAt())
		if publishedCmp != 0 {
			return publishedCmp
		}

		// AP outbox IDs are URLs and sorting them may not always be useful,
		// but after date sorting, it's usually okay.
		return cmp.Compare(a.ID(), b.ID())
	})

StatusesLoop:
	for _, status := range statusesInImportOrder {
		// Skip previously imported notes.
		if _, previouslyImported := archiveIdToImportedApiId[status.ID()]; previouslyImported {
			continue StatusesLoop
		}

		visibility := status.Visibility()
		// TODO: (Vyr) command-line visibility mapping
		if visibility != string(mastodon.VisibilityPublic) && visibility != string(mastodon.VisibilityUnlisted) {
			slog.Info("Skipping status due to visibility", "status", status.ID(), "visibility", string(visibility))
			continue StatusesLoop
		}

		// Allow OPs or self-replies to our other posts only.
		if status.InReplyToID() != "" && statusMap[status.InReplyToID()] == nil {
			slog.Info("Skipping status because it isn't a top-level status or self-reply", "status", status.Id)
			continue StatusesLoop
		}
		if status.TargetsSpecificUsersInToOrCc(actor) {
			var audience []string
			audience = append(audience, status.To...)
			audience = append(audience, status.Cc...)
			slog.Info("Skipping status because its audience includes other accounts", "status", status.Id, "audience", audience)
			continue StatusesLoop
		}

		// Convert to Markdown.
		markdown, err := markdownConverter.ConvertString(status.Content)
		if err != nil {
			slog.Error("Markdown conversion error", "status", status.Id, "err", err)
			continue StatusesLoop
		}

		// Check for messages that start with an @ and are clearly a non-self-reply.
		if strings.HasPrefix(markdown, "[@") {
			slog.Error("Skipping status because it looks like a reply", "status", status.Id)
			continue StatusesLoop
		}

		// Skip anything that mentions somebody else.
		// Skip other indicators of mentions.
		// Collect hashtags and emojis for later processing.
		noteTags := status.Tags()
		hashtags := make([]*mastodon.MentionOrHashtag, 0, len(noteTags))
		requiredEmojiNames := make([]string, 0, len(noteTags))
		for _, tag := range noteTags {
			tagType := tag.GetType()
			switch tagType {
			case "Mention":
				if mention, ok := tag.(*mastodon.MentionOrHashtag); ok {
					slog.Info("Skipping status because it mentions another account", "status", status.Id, "mention", mention)
					continue StatusesLoop
				}
			case "Hashtag":
				if hashtag, ok := tag.(*mastodon.MentionOrHashtag); ok {
					hashtags = append(hashtags, hashtag)
				}
			case "Emoji":
				if emoji, ok := tag.(*mastodon.Emoji); ok {
					requiredEmojiNames = append(requiredEmojiNames, emoji.Name)
				}
			default:
				slog.Error("Unexpected tag type", "status", status.Id, "tagType", tagType)
				continue StatusesLoop
			}
		}

		// Check for missing emojis.
		for _, name := range requiredEmojiNames {
			if _, foundOnInstance := instanceEmojiNames[name]; !foundOnInstance && !allowMissingCustomEmojis {
				slog.Error("Instance is missing required custom emoji", "status", status.Id, "emoji", name)
				continue StatusesLoop
			}
		}

		// Mentions from deleted instances or users don't show up as mention tags. Skip those too.
		// This transform allows the common case of URLs that contain @ but are probably links to Fedi posts.
		// TODO: (Vyr) Will give false positives on code blocks, Bluesky handles, mentions that never resolved, etc.
		atLinkPlaceholderMarkdown := AtLinkPattern.ReplaceAllString(markdown, "at_link_placeholder")
		if MentionPattern.MatchString(atLinkPlaceholderMarkdown) {
			slog.Info("Skipping status because it looks like it has a mention of a dead account", "status", status.Id)
			continue StatusesLoop
		}

		// Convert Markdown links to hashtags back into hashtags.
		// TODO: (Vyr) this might make more sense to do in HTML, and loses CamelCasing as is.
		for _, hashtag := range hashtags {
			linkPattern, err := regexp.Compile(`(?i)\Q` + `[` + hashtag.Name + `](` + hashtag.Href + `)\E`)
			if err != nil {
				slog.Error("Regex compilation error", "status", status.Id, "err", err)
				continue StatusesLoop
			}
			markdown = linkPattern.ReplaceAllString(markdown, hashtag.Name)
		}

		// Get the imported API ID of the post we're replying to, if applicable.
		var inReplyToApiId *string
		if status.InReplyTo != nil {
			inReplyToArchiveId := *status.InReplyTo
			if apiId, found := archiveIdToImportedApiId[inReplyToArchiveId]; found {
				inReplyToApiId = &apiId
			} else {
				slog.Error("Couldn't find API ID for status being replied to", "status", status.Id, "inReplyTo", inReplyToArchiveId)
				continue StatusesLoop
			}
		}

		// Upload any media attachments.
		mediaIDs := make([]string, 0, len(status.Attachments))
		for _, attachment := range status.Attachments {
			mediaID, err := uploadAttachment(
				authClient,
				attachmentMapFile,
				mediaPathToImportedApiId,
				attachment.Url,
				path.Join(file, attachment.Url),
				attachment.Name,
				attachment.FocalPointX(),
				attachment.FocalPointY(),
			)
			if err != nil {
				slog.Error("Error retrieving or uploading attachment", "status", status.Id, "attachment", attachment.Url, "err", err)
				continue StatusesLoop
			}
			mediaIDs = append(mediaIDs, mediaID)
		}

		if err := authClient.Wait(); err != nil {
			return err
		}

		response, err := authClient.Client.Statuses.StatusCreate(
			&statuses.StatusCreateParams{
				ContentType: &statusContentType,
				InReplyToID: inReplyToApiId,
				Language:    status.Language(),
				MediaIDs:    mediaIDs,
				ScheduledAt: util.Ptr(strfmt.DateTime(status.Published)),
				Sensitive:   &status.Sensitive,
				SpoilerText: status.Summary,
				Status:      &markdown,
				Visibility:  util.Ptr(string(status.Visibility())),
			},
			authClient.Auth,
			func(op *runtime.ClientOperation) {
				op.ConsumesMediaTypes = []string{"application/x-www-form-urlencoded"}
			},
		)
		if err != nil {
			slog.Error("Couldn't post converted status", "status", status.Id, "err", err)
			continue StatusesLoop
		}
		status := response.GetPayload()

		// Save the API ID of the status we just imported for future imports.
		archiveIdToImportedApiId[status.Id] = status.ID
		if err := writeMapFile(statusMapFile, archiveIdToImportedApiId); err != nil {
			slog.Error("Couldn't write status map file", "path", statusMapFile, "err", err)
			return err
		}

		slog.Info("Imported status", "status", status.Id, "url", status.URL)
	}

	return nil
}

func uploadAttachment(
	authClient *auth.Client,
	attachmentMapFile string,
	mediaPathToImportedApiId map[string]string,
	archiveURL string,
	localPath string,
	description *string,
	focusX *float64,
	focusY *float64,
) (string, error) {
	// Check the attachment map file for a previously uploaded copy.
	if id, previouslyUploaded := mediaPathToImportedApiId[archiveURL]; previouslyUploaded {
		return id, nil
	}

	filename := path.Base(archiveURL)
	file, err := os.Open(localPath)
	if err != nil {
		return "", err
	}

	if err := authClient.Wait(); err != nil {
		return "", err
	}

	var focusString *string
	if focusX != nil && focusY != nil {
		focusString = util.Ptr(fmt.Sprintf("%f,%f", *focusX, *focusY))
	}

	// TODO: (Vyr) we may have thumbnails as the .Icon property, but GTS doesn't support thumbnails yet.
	// 	Also note that as of Mastodon 4.3, thumbnails are exported as URLs on the original server,
	// 	and not as part of the archive, so video thumbnails require HTTPS calls to get,
	//	and may be lost if the server is down.
	response, err := authClient.Client.Media.MediaCreate(
		&media.MediaCreateParams{
			APIVersion:  "v2",
			Description: description,
			File:        runtime.NamedReader(filename, file),
			Focus:       focusString,
		},
		authClient.Auth,
		func(op *runtime.ClientOperation) {
			op.ConsumesMediaTypes = []string{"multipart/form-data"}
		},
	)
	if err != nil {
		slog.Error("Couldn't upload media attachment", "path", localPath, "err", err)
		return "", err
	}
	apiAttachment := response.GetPayload()

	// Save the API ID of the status we just imported for future imports.
	mediaPathToImportedApiId[archiveURL] = apiAttachment.ID
	if err := writeMapFile(attachmentMapFile, mediaPathToImportedApiId); err != nil {
		slog.Error("Couldn't write attachment map file", "path", attachmentMapFile, "err", err)
		return "", err
	}

	slog.Info("Imported attachment", "path", localPath, "url", apiAttachment.TextURL)
	return apiAttachment.ID, nil
}

func ReadJSON[T any](path string) (*T, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = jsonFile.Close() }()

	var doc T
	if err := json.NewDecoder(jsonFile).Decode(&doc); err != nil {
		return nil, err
	}
	return &doc, nil
}
