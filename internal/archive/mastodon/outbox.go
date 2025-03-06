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

package mastodon

import (
	"encoding/json"
	"time"
)

// Outbox represents a Mastodon export's `outbox.json`.
type Outbox struct {
	OrderedItems []Activity `json:"orderedItems"`
}

// Notes returns a list of Note objects.
func (o *Outbox) Notes() []*Object {
	notes := make([]*Object, 0, len(o.OrderedItems))
	for _, activity := range o.OrderedItems {
		if activity.Type != "Create" {
			continue
		}

		object := activity.Object()
		if object == nil {
			continue
		}

		if object.Type != "Note" {
			continue
		}
	}
	return notes
}

type Activity struct {
	Type      string          `json:"type"`
	RawObject json.RawMessage `json:"object"`
	// TODO: (Vyr) for compat with Akkoma, might need to check To, Cc, and DirectMessage on activity as well as object
}

func (a *Activity) Object() *Object {
	var object Object
	if err := json.Unmarshal(a.RawObject, &object); err == nil {
		return &object
	}
	return nil
}

type Object struct {
	Id             string            `json:"id"`
	Type           string            `json:"type"`
	Summary        *string           `json:"summary"`
	InReplyTo      *string           `json:"inReplyTo"`
	Published      time.Time         `json:"published"`
	Url            string            `json:"url"`
	To             []string          `json:"to"`
	Cc             []string          `json:"cc"`
	Sensitive      bool              `json:"sensitive"`
	Conversation   string            `json:"conversation"`
	Content        string            `json:"content"`
	ContentMap     map[string]string `json:"contentMap"`
	DirectMessage  bool              `json:"directMessage"`
	AllAttachments []*Attachment     `json:"attachment"`
	RawTags        []json.RawMessage `json:"tag"`
}

const ASPublic = "https://www.w3.org/ns/activitystreams#Public"

type Visibility string

const (
	VisibilityPublic   Visibility = "public"
	VisibilityUnlisted Visibility = "unlisted"
	VisibilityPrivate  Visibility = "private"
	VisibilityDirect   Visibility = "direct"
)

func (o *Object) Visibility() Visibility {
	for _, url := range o.To {
		if url == ASPublic {
			return VisibilityPublic
		}
	}
	for _, url := range o.Cc {
		if url == ASPublic {
			return VisibilityUnlisted
		}
	}
	if o.DirectMessage {
		return VisibilityDirect
	}
	return VisibilityPrivate
}

func (o *Object) Language() *string {
	for language, content := range o.ContentMap {
		if content == o.Content {
			return &language
		}
	}
	return nil
}

// TargetsSpecificUsersInToOrCc is true if the status is addressed to anything that
// isn't the public or followers collections or the actor that posted it.
func (o *Object) TargetsSpecificUsersInToOrCc(a *Actor) bool {
	for _, url := range o.To {
		if url != ASPublic && url != a.Id && url != a.Followers {
			return true
		}
	}
	for _, url := range o.Cc {
		if url != ASPublic && url != a.Id && url != a.Followers {
			return true
		}
	}
	return false
}

type Attachment struct {
	Type          string    `json:"type"`
	MediaType     string    `json:"mediaType"`
	Url           string    `json:"url"`
	Name          *string   `json:"name"`
	RawFocalPoint []float64 `json:"focalPoint"`
	Icon          *Icon     `json:"icon"`
	// TODO: (Vyr) Akkoma archives may have Url as a list containing href, mediaType, etc.
}

func (a *Attachment) FocalPointX() *float64 {
	if len(a.RawFocalPoint) > 0 {
		return &a.RawFocalPoint[0]
	}
	return nil
}

func (a *Attachment) FocalPointY() *float64 {
	if len(a.RawFocalPoint) > 1 {
		return &a.RawFocalPoint[1]
	}
	return nil
}

func (o *Object) Tags() []Tag {
	tags := make([]Tag, 0, len(o.RawTags))
	for _, rawTag := range o.RawTags {
		var mentionOrHashtag MentionOrHashtag
		var emoji Emoji
		if err := json.Unmarshal(rawTag, &mentionOrHashtag); err == nil &&
			(mentionOrHashtag.Type == "Mention" || mentionOrHashtag.Type == "Hashtag") {
			tags = append(tags, &mentionOrHashtag)
		} else if err := json.Unmarshal(rawTag, &emoji); err == nil &&
			mentionOrHashtag.Type == "Emoji" {
			tags = append(tags, &emoji)
		}
	}
	return tags
}

type Tag interface {
	GetType() string
}

type MentionOrHashtag struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Href string `json:"href"`
}

func (m *MentionOrHashtag) GetType() string {
	return m.Type
}

type Emoji struct {
	Type    string    `json:"type"`
	Name    string    `json:"name"`
	Id      string    `json:"id"`
	Updated time.Time `json:"updated"`
	Icon    Icon      `json:"icon"`
}

func (e *Emoji) GetType() string {
	return e.Type
}

type Icon struct {
	Type      string `json:"type"`
	MediaType string `json:"mediaType"`
	Url       string `json:"url"`
}
