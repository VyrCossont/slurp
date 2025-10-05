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
	"bytes"
	"context"
	"errors"
	"log/slog"
	"os"
	"path"
	"slices"
	"strings"
	"time"

	"github.com/VyrCossont/slurp/client/media"
	"github.com/VyrCossont/slurp/client/statuses"
	"github.com/VyrCossont/slurp/internal/util"
	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/atproto/syntax"
	bskyrepo "github.com/bluesky-social/indigo/repo"
	"github.com/docker/go-units"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	gocid "github.com/ipfs/go-cid"

	"github.com/VyrCossont/slurp/internal/auth"
	"github.com/VyrCossont/slurp/internal/bluesky"
)

func BlueskyImport(
	authClient *auth.Client,
	archiveFolderPath string,
	statusMapFile string,
	attachmentMapFile string,
) error {
	ctx := context.TODO()

	postCIDToImportedApiId, blobCIDToImportedApiId, err := requireMapFiles(statusMapFile, attachmentMapFile)
	if err != nil {
		return err
	}

	// Load and sort posts.
	posts, err := loadPostsFromRepo(ctx, archiveFolderPath)
	if err != nil {
		return err
	}
	slices.SortFunc(posts, comparePosts)
	slog.Info("found some posts", "count", len(posts))

	// Import posts. Skip ones we can't process for some reason, without giving up.
	for _, post := range posts {
		err = importPost(authClient, archiveFolderPath, statusMapFile, postCIDToImportedApiId, attachmentMapFile, blobCIDToImportedApiId, post)
		if err != nil && !errors.Is(err, skipPost) {
			slog.Error("couldn't import post", "cid", post.cid.String(), "error", err)
			return err
		}
	}

	return nil
}

// sortablePost contains a post with metadata attached and parsed.
type sortablePost struct {
	cid       gocid.Cid
	recordKey syntax.RecordKey
	post      *bsky.FeedPost
	createdAt time.Time
	// parentCID is only present for posts that are replies.
	parentCID *gocid.Cid
}

// loadPostsFromRepo loads all the posts from a given archive's repo and returns them indexed by CID.
func loadPostsFromRepo(ctx context.Context, archiveFolderPath string) ([]*sortablePost, error) {
	carPath := path.Join(archiveFolderPath, bluesky.RepoCAR)
	f, err := os.Open(carPath)
	if err != nil {
		slog.Error("unable to open repo", "path", carPath, "error", err)
		return nil, err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			slog.Error("couldn't close repo file", "path", carPath, "error", err)
		}
	}()

	repo, err := bskyrepo.ReadRepoFromCar(ctx, f)
	if err != nil {
		slog.Error("unable to read repo", "path", carPath, "error", err)
		return nil, err
	}

	var posts []*sortablePost

	err = repo.ForEach(ctx, "", func(repoPath string, cid gocid.Cid) error {
		post, err := loadPostFromRepo(ctx, repo, repoPath, cid)
		if err != nil {
			if !errors.Is(err, skipPost) {
				// skipPost is an expected error, but this error wasn't expected.
				slog.Error("unable to load post, skipping it", "repoPath", repoPath, "cid", cid, "error", err)
			}
			return nil
		}
		posts = append(posts, post)
		return nil
	})
	if err != nil {
		slog.Error("failure while iterating over repo", "path", carPath, "error", err)
		return nil, err
	}

	return posts, nil
}

// loadPostFromRepo loads a single post from an open repo by repo path, including its metadata.
func loadPostFromRepo(ctx context.Context, repo *bskyrepo.Repo, repoPath string, cid gocid.Cid) (*sortablePost, error) {
	collection, recordKey, err := syntax.ParseRepoPath(repoPath)
	if err != nil {
		slog.Error("invalid path in repo", "repoPath", repoPath, "error", err)
		return nil, err
	}

	if collection != "app.bsky.feed.post" {
		// Not a post. Skip it.
		return nil, skipPost
	}

	_, recordCBOR, err := repo.GetRecordBytes(ctx, repoPath)
	if err != nil {
		slog.Error("couldn't get record from repo", "repoPath", repoPath, "error", err)
		return nil, err
	}

	post := &bsky.FeedPost{}
	err = post.UnmarshalCBOR(bytes.NewReader(*recordCBOR))
	if err != nil {
		slog.Error("couldn't decode post from repo", "error", err)
		return nil, err
	}

	sortable := sortablePost{cid: cid, post: post, recordKey: recordKey}

	// Parse timestamp. Supposed to be RFC 3339: https://atproto.com/specs/lexicon#:~:text=Datetime%20strings%20in%20atproto
	sortable.createdAt, err = time.Parse(time.RFC3339Nano, post.CreatedAt)
	if err != nil {
		sortable.createdAt, err = time.Parse(time.RFC3339, post.CreatedAt)
	}
	if err != nil {
		slog.Error("couldn't parse post's creation time", "cid", cid, "timestamp", post.CreatedAt, "error", err)
		return nil, err
	}

	// Parse reply parent CID if it has one.
	if post.Reply != nil && post.Reply.Parent != nil {
		if post.Reply.Parent == nil {
			slog.Error("post is a reply but has no parent", "cid", cid)
			return nil, err
		}
		parentCID, err := gocid.Parse(post.Reply.Parent.Cid)
		if err != nil {
			slog.Error("couldn't parse post's reply parent CID", "cid", cid, "parentCID", post.Reply.Parent.Cid)
			return nil, err
		}
		sortable.parentCID = &parentCID
	}

	return &sortable, nil
}

// sortPosts sorts posts topologically and then by date for consistent import order.
func comparePosts(a, b *sortablePost) int {
	// Topo sort by reply graph (required).
	if b.parentCID != nil && a.cid == *b.parentCID {
		return -1
	} else if a.parentCID != nil && *a.parentCID == b.cid {
		return 1
	}

	// Sort by creation date.
	createdAtCmp := a.createdAt.Compare(b.createdAt)
	if createdAtCmp != 0 {
		return createdAtCmp
	}

	// Tiebreak by sorting on repo record key.
	return strings.Compare(string(a.recordKey), string(b.recordKey))
}

// importPost creates a Fedi status from a Bluesky post, including attachments.
// Only returns showstopper errors that should stop the entire import process.
func importPost(
	authClient *auth.Client,
	archiveFolderPath string,
	statusMapFile string,
	postCIDToImportedApiId map[string]string,
	attachmentMapFile string,
	blobCIDToImportedApiId map[string]string,
	post *sortablePost,
) error {
	cidString := post.cid.String()

	// Skip posts that have already been imported.
	if _, alreadyImported := postCIDToImportedApiId[cidString]; alreadyImported {
		slog.Info("skipping already imported post", "cid", cidString)
		return skipPost
	}

	// If a post is a reply, we must have imported its parent.
	// Posts replying to other users will never meet this criterion and will be skipped.
	var optionalInReplyToID *string
	if post.parentCID != nil {
		inReplyToID, found := postCIDToImportedApiId[post.parentCID.String()]
		if !found {
			slog.Info("skipping reply post because we did not import the post is replying to", "cid", cidString,
				"text", post.post.Text)
			return skipPost
		}
		optionalInReplyToID = &inReplyToID
	}

	// Check for content labels.
	visibility := VisibilityPublic
	var sensitiveMedia *bool
	var cwText *string
	if post.post.Labels != nil {
		labels := post.post.Labels.LabelDefs_SelfLabels.Values
		cwTextParts := make([]string, 0, len(labels))
		for _, label := range labels {
			l := blueskySelfLabel(label.Val)
			switch l {
			case blueskySelfLabelNoUnauthenticated:
				// TODO: this label may be intended for *accounts*, not posts,
				//	in which case this will do effectively nothing.
				visibility = VisibilityPrivate
			case blueskySelfLabelPorn,
				blueskySelfLabelSexual,
				blueskySelfLabelNudity,
				blueskySelfLabelGraphicMedia:
				sensitiveMedia = util.Ptr(true)
				cwTextPart := l.CWText()
				if cwTextPart != "" {
					cwTextParts = append(cwTextParts, cwTextPart)
				}
			default:
				slog.Info("unknown Bluesky self label", "label", l)
			}
		}
		if len(cwTextParts) > 0 {
			cwText = util.Ptr("CW: " + strings.Join(cwTextParts, ", "))
		}
	}

	// Guess language.
	var language *string
	if len(post.post.Langs) == 1 {
		language = &post.post.Langs[0]
	} else if len(post.post.Langs) > 1 {
		language = util.Ptr("mul")
	}

	// Apply rich text facets to reconstruct post source.
	text, err := resolveFacets(post)
	if err != nil {
		return err
	}
	// Use the plain text type.
	// TODO: users may well want to override this if they post in Markdown on Bluesky.
	contentType := "text/plain"
	// Note: we don't have a way to prevent emoji processing in the destination server.

	// Import a post's attachments.
	mediaIDs, err := importPostMedia(authClient, archiveFolderPath, attachmentMapFile, blobCIDToImportedApiId, post)
	if err != nil {
		return err
	}

	response, err := authClient.Client.Statuses.StatusCreate(
		&statuses.StatusCreateParams{
			ContentType: &contentType,
			InReplyToID: optionalInReplyToID,
			Language:    language,
			MediaIDs:    mediaIDs,
			ScheduledAt: util.Ptr(strfmt.DateTime(post.createdAt)),
			Sensitive:   sensitiveMedia,
			SpoilerText: cwText,
			Status:      &text,
			Visibility:  util.Ptr(string(visibility)),
		},
		authClient.Auth,
		func(op *runtime.ClientOperation) {
			op.ConsumesMediaTypes = []string{"application/x-www-form-urlencoded"}
		},
	)
	if err != nil {
		slog.Error("failed to post converted post", "cid", cidString, "error", err)
		return err
	}
	status := response.GetPayload()

	postCIDToImportedApiId[cidString] = status.ID
	if err := writeMapFile(statusMapFile, postCIDToImportedApiId); err != nil {
		slog.Error("couldn't write status map file", "path", statusMapFile, "error", err)
		return err
	}

	slog.Info("imported post", "cid", cidString, "url", status.URL)

	return nil
}

// blueskySelfLabel is one of the known self-label values for sensitive content.
// https://docs.bsky.app/docs/advanced-guides/moderation#self-labels
type blueskySelfLabel string

const (
	// blueskySelfLabelNoUnauthenticated is a request not to show this post to unauthenticated users.
	blueskySelfLabelNoUnauthenticated blueskySelfLabel = "!no-unauthenticated"
	blueskySelfLabelPorn              blueskySelfLabel = "porn"
	blueskySelfLabelSexual            blueskySelfLabel = "sexual"
	blueskySelfLabelNudity            blueskySelfLabel = "nudity"
	blueskySelfLabelGraphicMedia      blueskySelfLabel = "graphic-media"
)

// CWText returns the content warning text for that label.
// TODO: localize these and match to post language.
func (l blueskySelfLabel) CWText() string {
	switch l {
	case blueskySelfLabelNoUnauthenticated:
		// Does not appear in CWs.
		return ""
	case blueskySelfLabelPorn, blueskySelfLabelSexual, blueskySelfLabelNudity:
		return string(l)
	case blueskySelfLabelGraphicMedia:
		return "graphic media"
	default:
		return ""
	}
}

// skipPost marks a post that we are skipping but one that should not stop the whole process.
var skipPost = errors.New("skipping post")

// resolveFacets turns Bluesky special text spans corresponding to mentions, hashtags, or links into appropriate formatting.
func resolveFacets(post *sortablePost) (string, error) {
	builder := strings.Builder{}

	facets := post.post.Facets
	slices.SortFunc(facets, func(a, b *bsky.RichtextFacet) int {
		startCmp := a.Index.ByteStart - b.Index.ByteStart
		if startCmp != 0 {
			return int(startCmp)
		}
		return int(a.Index.ByteEnd - b.Index.ByteEnd)
	})

	text := post.post.Text
	cursor := 0
	for _, facet := range facets {
		// We only care about link facets right now.
		link := ""
		for _, feature := range facet.Features {
			// We can ignore feature.RichtextFacet_Tag since the text representation of Fedi tags and Bluesky tags is effectively identical.
			// TODO: is this true for #tags#close#together?

			if feature.RichtextFacet_Mention != nil {
				// TODO: mentions of the post's owner are okay, but should be fairly rare
				slog.Info("skipping post that contains a mention", "cid", post.cid.String())
				return "", skipPost
			}

			if feature.RichtextFacet_Link != nil {
				link = feature.RichtextFacet_Link.Uri
			}
		}

		start := int(facet.Index.ByteStart)
		end := int(facet.Index.ByteEnd)
		if link == "" {
			// Copy text from cursor thru end of facet.
			builder.WriteString(text[cursor:end])
		} else {
			// Copy text up to start of facet.
			if cursor < start {
				builder.WriteString(text[cursor:start])
			}
			// Replace facet's span in text with its link.
			builder.WriteString(link)
		}
		cursor = end
	}
	// Copy tail after last facet.
	if cursor < len(text) {
		builder.WriteString(text[cursor:])
	}

	return builder.String(), nil
}

// importPostMedia imports a post's media attachments.
func importPostMedia(
	authClient *auth.Client,
	archiveFolderPath string,
	attachmentMapFile string,
	blobCIDToImportedApiId map[string]string,
	post *sortablePost,
) ([]string, error) {
	embed := post.post.Embed
	if embed == nil {
		return nil, nil
	}

	// As of writing there are five kinds.
	// embed.EmbedExternal is a link preview card and we ignore that.

	if embed.EmbedRecord != nil || embed.EmbedRecordWithMedia != nil {
		// TODO: are these quote skeets?
		slog.Error("don't know how to process quote skeets, skipping post", "cid", post.cid)
		return nil, skipPost
	}

	var attachables []attachableMedia
	if embed.EmbedImages != nil {
		for _, image := range embed.EmbedImages.Images {
			attachables = append(attachables, attachableImage{image})
		}
	}
	if embed.EmbedVideo != nil {
		attachables = append(attachables, attachableVideo{embed.EmbedVideo})
	}

	attachmentIDs := make([]string, 0, len(attachables))
	for _, attachable := range attachables {
		attachmentID, err := importMedia(authClient, archiveFolderPath, attachmentMapFile, blobCIDToImportedApiId, attachable)
		if err != nil {
			slog.Error("error uploading attachment, skipping post", "error", err)
			return nil, err
		}
		attachmentIDs = append(attachmentIDs, attachmentID)
	}

	return attachmentIDs, nil
}

// importMedia creates a Fedi media attachment from a Bluesky embedded image or video and its blob, returning the ID.
func importMedia(
	authClient *auth.Client,
	archiveFolderPath string,
	attachmentMapFile string,
	blobCIDToImportedApiId map[string]string,
	attachable attachableMedia,
) (string, error) {
	cidString := attachable.cidString()

	// Check the attachment map file for a previously uploaded copy.
	if attachmentID, previouslyUploaded := blobCIDToImportedApiId[cidString]; previouslyUploaded {
		return attachmentID, nil
	}

	blobPath := path.Join(archiveFolderPath, bluesky.BlobsFolder, cidString)
	stat, err := os.Stat(blobPath)
	if err != nil {
		slog.Error("couldn't find blob", "cid", cidString, "path", blobPath, "error", err)
		return "", err
	}
	file, err := os.Open(blobPath)
	if err != nil {
		slog.Error("couldn't open blob", "cid", cidString, "path", blobPath, "error", err)
		return "", err
	}
	slog.Info(
		"found blob",
		"cid", cidString,
		"path", blobPath,
		"size", units.BytesSize(float64(stat.Size())),
		"contentType", attachable.mimeString(),
	)

	response, err := authClient.Client.Media.MediaCreate(
		&media.MediaCreateParams{
			APIVersion:  "v2",
			Description: attachable.altText(),
			File: typedNamedReadCloser{
				NamedReadCloser: runtime.NamedReader(cidString, file),
				contentType:     attachable.mimeString(),
			},
		},
		authClient.Auth,
		func(op *runtime.ClientOperation) {
			op.ConsumesMediaTypes = []string{"multipart/form-data"}
		},
	)
	if err != nil {
		slog.Error("couldn't upload attachment", "cid", cidString, "path", blobPath, "error", err)
		return "", err
	}
	attachment := response.GetPayload()

	blobCIDToImportedApiId[attachable.cidString()] = attachment.ID
	if err := writeMapFile(attachmentMapFile, blobCIDToImportedApiId); err != nil {
		slog.Error("couldn't write attachment map file", "path", attachmentMapFile, "error", err)
		return "", err
	}

	slog.Info("uploaded attachment", "cid", cidString, "path", blobPath, "url", attachment.TextURL)
	return attachment.ID, nil
}

// attachableMedia is a file identified by blob CID, with MIME type and alt text.
type attachableMedia interface {
	cidString() string
	mimeString() string
	altText() *string
}

type attachableImage struct {
	image *bsky.EmbedImages_Image
}

func (a attachableImage) cidString() string {
	return a.image.Image.Ref.String()
}

func (a attachableImage) mimeString() string {
	return a.image.Image.MimeType
}

func (a attachableImage) altText() *string {
	if a.image.Alt == "" {
		return nil
	}
	return &a.image.Alt
}

type attachableVideo struct {
	video *bsky.EmbedVideo
}

func (a attachableVideo) cidString() string {
	return a.video.Video.Ref.String()
}

func (a attachableVideo) mimeString() string {
	return a.video.Video.MimeType
}

func (a attachableVideo) altText() *string {
	return a.video.Alt
}

type typedNamedReadCloser struct {
	runtime.NamedReadCloser
	contentType string
}

func (t typedNamedReadCloser) ContentType() string {
	return t.contentType
}
