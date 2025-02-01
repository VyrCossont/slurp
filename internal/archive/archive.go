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
	"io/fs"
	"log"
	"os"
	"path"
	"regexp"
	"slices"
	"strings"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/VyrCossont/slurp/client/statuses"
	"github.com/VyrCossont/slurp/internal/auth"
	"github.com/VyrCossont/slurp/internal/own"
	"github.com/VyrCossont/slurp/internal/util"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

var mentionPattern = regexp.MustCompile(`@\w+`)

func Import(authClient *auth.Client, file string, mapFile string) error {
	// TODO: (Vyr) assume folder for now
	if file == "-" || strings.HasSuffix(strings.ToLower(file), ".zip") {
		return errors.New("can't handle compressed archives yet")
	}

	if !strings.HasSuffix(strings.ToLower(mapFile), ".json") {
		return errors.New("map file is required and must have a .json extension")
	}
	archiveIdToImportedApiId, err := readMapFile(mapFile)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			archiveIdToImportedApiId = map[string]string{}
		} else {
			return err
		}
	}
	if err = writeMapFile(mapFile, archiveIdToImportedApiId); err != nil {
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

	// TODO: (Vyr) do we need to detect scheduled status support *and* backfill support?
	// 	Assuming GTS 0.18 for now.

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

	// TODO: (Vyr) replace log with slog

NotesLoop:
	for _, note := range notesInOrder {
		if _, previouslyImported := archiveIdToImportedApiId[note.Id]; previouslyImported {
			continue
		}

		visibility := note.Visibility()
		if visibility != "public" && visibility != "unlisted" {
			continue
		}

		// Allow OPs or self-replies to our other posts only.
		if note.InReplyTo != nil {
			if notes[*note.InReplyTo] == nil {
				// Not a self reply.
				continue
			}
		}

		// Skip anything that mentions somebody else.
		if note.TargetsSpecificUsersInToOrCc(actor) {
			continue
		}
		// Skip other indicators of mentions.
		// Collect hashtags for later processing.
		noteTags := note.Tags()
		hashtags := make([]*MentionOrHashtag, 0, len(noteTags))
		requiredEmojiNames := make([]string, 0, len(noteTags))
		for _, tag := range noteTags {
			tagType := tag.GetType()
			switch tagType {
			case "Mention":
				continue NotesLoop
			case "Hashtag":
				if hashtag, ok := tag.(*MentionOrHashtag); ok {
					hashtags = append(hashtags, hashtag)
				}
			case "Emoji":
				if emoji, ok := tag.(*Emoji); ok {
					requiredEmojiNames = append(requiredEmojiNames, emoji.Name)
				}
			default:
				log.Printf("Unexpected tag type: %+v", tagType)
				continue
			}
		}

		// Mentions from deleted instances or users don't show up as mention tags. Skip those too.
		// Will give false positives on code blocks, Bluesky handles, mentions that never resolved,
		// @ in URLs in HTML-formatted posts, etc.
		if mentionPattern.MatchString(note.Content) {
			continue
		}

		// Convert to Markdown.
		markdown, err := htmltomarkdown.ConvertString(note.Content)
		if err != nil {
			log.Printf("Markdown error: %+v", err)
			continue
		}

		// Convert Markdown links to hashtags back into hashtags.
		// TODO: (Vyr) this might make more sense to do in HTML, and loses CamelCasing as is.
		for _, hashtag := range hashtags {
			linkPattern, err := regexp.Compile(`(?i)\Q` + `[` + hashtag.Name + `](` + hashtag.Href + `)\E`)
			if err != nil {
				log.Printf("Markdown error: %+v", err)
				continue
			}
			markdown = linkPattern.ReplaceAllString(markdown, hashtag.Name)
		}

		// TODO: (Vyr) handle attachments. For now, skip posts with them.
		if len(note.Attachments) > 0 {
			continue
		}

		// TODO: (Vyr) handle emoji. This will need admin privs, and access to a live original server,
		// 	or an emoji export of some kind if there is such a thing.
		if len(requiredEmojiNames) > 0 {
			continue
		}

		// Get the imported API ID of the post we're replying to, if applicable.
		var inReplyToApiId *string
		if note.InReplyTo != nil {
			inReplyToArchiveId := *note.InReplyTo
			if apiId, found := archiveIdToImportedApiId[inReplyToArchiveId]; found {
				inReplyToApiId = &apiId
			} else {
				log.Printf("couldn't find API ID for archive ID (post may have not been imported if not supported, or topo sort error): %+v", inReplyToArchiveId)
				continue
			}
		}

		if err := authClient.Wait(); err != nil {
			return err
		}

		response, err := authClient.Client.Statuses.StatusCreate(
			&statuses.StatusCreateParams{
				ContentType: &statusContentType,
				InReplyToID: inReplyToApiId,
				Language:    note.Language(),
				ScheduledAt: util.Ptr(strfmt.DateTime(note.Published)),
				Sensitive:   &note.Sensitive,
				SpoilerText: note.Summary,
				Status:      &markdown,
				Visibility:  &visibility,
			},
			authClient.Auth,
			func(op *runtime.ClientOperation) {
				op.ConsumesMediaTypes = []string{"application/x-www-form-urlencoded"}
			},
		)
		if err != nil {
			log.Printf("couldn't post converted status with archive ID %+v: %+v", note.Id, err)
			continue
		}
		status := response.GetPayload()

		// Save the API ID of the status we just imported for future imports.
		archiveIdToImportedApiId[note.Id] = status.ID
		if err := writeMapFile(mapFile, archiveIdToImportedApiId); err != nil {
			return err
		}

		log.Printf("imported %s as %s", note.Id, status.URL)
	}

	return nil
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
