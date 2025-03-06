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

package misskey

import (
	"time"

	"github.com/VyrCossont/slurp/internal/archive/mastodon"
)

type MisskeyArchiveNote struct {
	ID string `json:"id"`

	// Text appears to be the original source, and may be null, at least for boosts.
	Text *string `json:"text"`

	CreatedAt time.Time `json:"createdAt"`

	// FileIDs is always an empty array in the sample.
	// TODO: Misskey "fileIds" field
	FileIDs []struct{} `json:"fileIds"`

	// Files is always an empty array in the sample.
	// TODO: Misskey "files" field
	Files []struct{} `json:"files"`

	// ReplyID is null if the post is not a reply.
	ReplyID *string `json:"replyId"`

	// RenoteID is null if the post is not a boost.
	RenoteID *string `json:"renoteId"`

	// Poll is always null in the sample.
	// TODO: Misskey "poll" field
	Poll *struct{} `json:"poll"`

	CW *string `json:"cw"`

	MisskeyVisibility Visibility `json:"visibility"`

	// VisibleUserIDs is set for `"visibility": "specified"` posts.
	// It is not a _mention_ list, since mentions in the sample don't use it.
	VisibleUserIDs []string

	LocalOnly bool `json:"localOnly"`

	// ReactionAcceptance is always null in the sample.
	// TODO: Misskey "reactionAcceptance" field
	ReactionAcceptance *struct{} `json:"reactionAcceptance"`
}

// Visibility returns the Mastodon-equivalent visibility.
func (n *MisskeyArchiveNote) Visibility() mastodon.Visibility {
	switch n.MisskeyVisibility {
	case VisibilityPublic:
		return mastodon.VisibilityPublic
	case VisibilityHome:
		return mastodon.VisibilityUnlisted
	case VisibilityFollowers:
		return mastodon.VisibilityPrivate
	case VisibilitySpecified:
		return mastodon.VisibilityDirect
	default:
		panic("unknown Misskey visibility level: " + n.MisskeyVisibility)
	}
}

type Visibility string

const (
	// VisibilityPublic is the equivalent of Mastodon "unlisted".
	VisibilityPublic Visibility = "public"
	// VisibilityHome is the equivalent of Mastodon "unlisted".
	VisibilityHome Visibility = "home"
	// VisibilityFollowers is the equivalent of Mastodon "private".
	// TODO: this isn't in the sample so the enum value might be wrong.
	VisibilityFollowers Visibility = "followers"
	// VisibilitySpecified is the equivalent of Mastodon "direct".
	VisibilitySpecified Visibility = "specified"
)
