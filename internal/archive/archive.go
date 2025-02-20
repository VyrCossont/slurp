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
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VyrCossont/slurp/client/media"
	"github.com/VyrCossont/slurp/client/statuses"
	"github.com/VyrCossont/slurp/internal/auth"
	"github.com/VyrCossont/slurp/internal/own"
	"github.com/VyrCossont/slurp/internal/util"
)

var mentionPattern = regexp.MustCompile(`@\w+`)
var atLinkPattern = regexp.MustCompile(`https://[a-z0-9.-]+/@\w+`)

func Import(authClient *auth.Client, file string, statusMapFile string, attachmentMapFile string) error {
	// Require archive to be already uncompressed.
	stat, err := os.Stat(file)
	if err != nil {
		slog.Error("couldn't get archive info from filesystem", "path", file, "err", err)
		return err
	}
	if !stat.IsDir() {
		return errors.New("archive must be an uncompressed folder")
	}

	archiveIdToImportedApiId, mediaPathToImportedApiId, err := requireMapFiles(statusMapFile, attachmentMapFile)
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

	actor, err := readActor(file)
	if err != nil {
		return err
	}

	outbox, err := readOutbox(file)
	if err != nil {
		return err
	}
	notes := outbox.Notes()

	// Put notes in order so we do things consistently. Only the topo sort is required for correctness,
	// but the rest is a nice property when running the same job again and again while debugging.
	notesInOrder := make([]*Object, 0, len(notes))
	for _, note := range notes {
		notesInOrder = append(notesInOrder, note)
	}
	slices.SortFunc(notesInOrder, func(a, b *Object) int {
		// Topo sort by reply graph.
		if b.InReplyTo != nil && a.Id == *b.InReplyTo {
			return -1
		} else if a.InReplyTo != nil && *a.InReplyTo == b.Id {
			return 1
		}

		// Sort by published date.
		publishedCmp := a.Published.Compare(b.Published)
		if publishedCmp != 0 {
			return publishedCmp
		}

		// AP outbox IDs are URLs and sorting them may not always be useful,
		// but after date sorting, it's usually okay.
		return cmp.Compare(a.Id, b.Id)
	})

NotesLoop:
	for _, note := range notesInOrder {
		// Skip previously imported notes.
		if _, previouslyImported := archiveIdToImportedApiId[note.Id]; previouslyImported {
			continue NotesLoop
		}

		visibility := note.Visibility()
		// TODO: (Vyr) command-line visibility mapping
		if visibility != VisibilityPublic && visibility != VisibilityUnlisted {
			slog.Info("Skipping status due to visibility", "status", note.Id, "visibility", string(visibility))
			continue NotesLoop
		}

		// Allow OPs or self-replies to our other posts only.
		if note.InReplyTo != nil && notes[*note.InReplyTo] == nil {
			slog.Info("Skipping status because it isn't a top-level status or self-reply", "status", note.Id)
			continue NotesLoop
		}
		if note.TargetsSpecificUsersInToOrCc(actor) {
			var audience []string
			audience = append(audience, note.To...)
			audience = append(audience, note.Cc...)
			slog.Info("Skipping status because its audience includes other accounts", "status", note.Id, "audience", audience)
			continue NotesLoop
		}

		// Convert to Markdown.
		markdown, err := markdownConverter.ConvertString(note.Content)
		if err != nil {
			slog.Error("Markdown conversion error", "status", note.Id, "err", err)
			continue NotesLoop
		}

		// Check for messages that start with an @ and are clearly a non-self-reply.
		if strings.HasPrefix(markdown, "[@") {
			slog.Error("Skipping status because it looks like a reply", "status", note.Id)
			continue NotesLoop
		}

		// Skip anything that mentions somebody else.
		// Skip other indicators of mentions.
		// Collect hashtags and emojis for later processing.
		noteTags := note.Tags()
		hashtags := make([]*MentionOrHashtag, 0, len(noteTags))
		requiredEmojiNames := make([]string, 0, len(noteTags))
		for _, tag := range noteTags {
			tagType := tag.GetType()
			switch tagType {
			case "Mention":
				if mention, ok := tag.(*MentionOrHashtag); ok {
					slog.Info("Skipping status because it mentions another account", "status", note.Id, "mention", mention)
					continue NotesLoop
				}
			case "Hashtag":
				if hashtag, ok := tag.(*MentionOrHashtag); ok {
					hashtags = append(hashtags, hashtag)
				}
			case "Emoji":
				if emoji, ok := tag.(*Emoji); ok {
					requiredEmojiNames = append(requiredEmojiNames, emoji.Name)
				}
			default:
				slog.Error("Unexpected tag type", "status", note.Id, "tagType", tagType)
				continue NotesLoop
			}
		}

		// Check for missing emojis.
		for _, name := range requiredEmojiNames {
			if _, foundOnInstance := instanceEmojiNames[name]; !foundOnInstance {
				slog.Error("Instance is missing required custom emoji", "status", note.Id, "emoji", name)
				continue NotesLoop
			}
		}

		// Mentions from deleted instances or users don't show up as mention tags. Skip those too.
		// This transform allows the common case of URLs that contain @ but are probably links to Fedi posts.
		// TODO: (Vyr) Will give false positives on code blocks, Bluesky handles, mentions that never resolved, etc.
		atLinkPlaceholderMarkdown := atLinkPattern.ReplaceAllString(markdown, "at_link_placeholder")
		if mentionPattern.MatchString(atLinkPlaceholderMarkdown) {
			slog.Info("Skipping status because it looks like it has a mention of a dead account", "status", note.Id)
			continue NotesLoop
		}

		// Convert Markdown links to hashtags back into hashtags.
		// TODO: (Vyr) this might make more sense to do in HTML, and loses CamelCasing as is.
		for _, hashtag := range hashtags {
			linkPattern, err := regexp.Compile(`(?i)\Q` + `[` + hashtag.Name + `](` + hashtag.Href + `)\E`)
			if err != nil {
				slog.Error("Regex compilation error", "status", note.Id, "err", err)
				continue NotesLoop
			}
			markdown = linkPattern.ReplaceAllString(markdown, hashtag.Name)
		}

		// Get the imported API ID of the post we're replying to, if applicable.
		var inReplyToApiId *string
		if note.InReplyTo != nil {
			inReplyToArchiveId := *note.InReplyTo
			if apiId, found := archiveIdToImportedApiId[inReplyToArchiveId]; found {
				inReplyToApiId = &apiId
			} else {
				slog.Error("Couldn't find API ID for status being replied to", "status", note.Id, "inReplyTo", inReplyToArchiveId)
				continue NotesLoop
			}
		}

		// Upload any media attachments.
		mediaIDs := make([]string, 0, len(note.Attachments))
		for _, attachment := range note.Attachments {
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
				slog.Error("Error retrieving or uploading attachment", "status", note.Id, "attachment", attachment.Url, "err", err)
				continue NotesLoop
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
				Language:    note.Language(),
				MediaIDs:    mediaIDs,
				ScheduledAt: util.Ptr(strfmt.DateTime(note.Published)),
				Sensitive:   &note.Sensitive,
				SpoilerText: note.Summary,
				Status:      &markdown,
				Visibility:  util.Ptr(string(note.Visibility())),
			},
			authClient.Auth,
			func(op *runtime.ClientOperation) {
				op.ConsumesMediaTypes = []string{"application/x-www-form-urlencoded"}
			},
		)
		if err != nil {
			slog.Error("Couldn't post converted status", "status", note.Id, "err", err)
			continue NotesLoop
		}
		status := response.GetPayload()

		// Save the API ID of the status we just imported for future imports.
		archiveIdToImportedApiId[note.Id] = status.ID
		if err := writeMapFile(statusMapFile, archiveIdToImportedApiId); err != nil {
			slog.Error("Couldn't write status map file", "path", statusMapFile, "err", err)
			return err
		}

		slog.Info("Imported status", "status", note.Id, "url", status.URL)
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

// TODO: (Vyr) can we collapse these two into a generic?
func readActor(file string) (*Actor, error) {
	jsonPath := path.Join(file, "actor.json")

	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		return nil, err
	}
	defer func() { _ = jsonFile.Close() }()

	var doc Actor
	if err := json.NewDecoder(jsonFile).Decode(&doc); err != nil {
		return nil, err
	}
	return &doc, nil
}

func readOutbox(file string) (*Outbox, error) {
	jsonPath := path.Join(file, "outbox.json")

	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		return nil, err
	}
	defer func() { _ = jsonFile.Close() }()

	var doc Outbox
	if err := json.NewDecoder(jsonFile).Decode(&doc); err != nil {
		return nil, err
	}
	return &doc, nil
}
