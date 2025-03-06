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
	"time"

	"github.com/VyrCossont/slurp/client/media"
	"github.com/VyrCossont/slurp/client/statuses"
)

type ImportableReader interface {
	Statuses() []ImportableStatus

	// Fetch returns the local path for an attachment,
	// which may or may not involve downloading it from somewhere.
	Fetch(attachment ImportableAttachment) (string, error)
}

// ImportableStatus is something that can be converted to GTS API params for creating a status.
type ImportableStatus interface {
	// ID is the archive ID of the status, as used in the status map file.
	// Might be an URL for Mastodon ActivityPub archives, or an API ID for others.
	ID() string

	CreatedAt() time.Time

	// Visibility returns the GTS API visibility as a string.
	Visibility() string

	// HasPoll is true if the status has/is a poll.
	HasPoll() bool

	// IsBoost is true for boosts and quotes, and false otherwise.
	IsBoost() bool

	// IsReplyToOtherAccount is false for OPs and self-replies and true otherwise.
	IsReplyToOtherAccount() bool

	// MentionsOtherAccount is true if the status mentions an account that is not the creator.
	MentionsOtherAccount() bool

	// InReplyToID is the archive ID of the status being replied to, if there is one,
	// and empty otherwise.
	InReplyToID() string

	// EmojiNames returns the emoji names used by the status,
	// *with* leading and trailing colons.
	EmojiNames() []string

	Attachments() []ImportableAttachment

	// Params returns the parameters necessary to create a status,
	// *without* the media attachment IDs.
	Params() *statuses.StatusCreateParams
}

// ImportableAttachment is something that can be converted to GTS API params for uploading a media attachment.
type ImportableAttachment interface {
	// URL is the archive URL of the media attachment, as used in the attachment map file.
	// Might be relative to the archive for Mastodon ActivityPub archives,
	// or fully qualified for Pixelfed remote attachments that need to be downloaded.
	URL() string

	// Params returns the parameters necessary to upload a media attachment,
	// *without* the API version or the file upload reader.
	Params() *media.MediaCreateParams
}
