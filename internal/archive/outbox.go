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
	"encoding/json"
	"time"
)

type Outbox struct {
	OrderedItems []Activity `json:"orderedItems"`
}

func (o *Outbox) Notes() map[string]*Object {
	notes := make(map[string]*Object, len(o.OrderedItems))
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

		notes[object.Id] = object
	}
	return notes
}

type Activity struct {
	Type      string          `json:"type"`
	To        []string        `json:"to"`
	Cc        []string        `json:"cc"`
	RawObject json.RawMessage `json:"object"`
}

func (a *Activity) Object() *Object {
	var object Object
	if err := json.Unmarshal(a.RawObject, &object); err == nil {
		return &object
	}
	return nil
}

type Object struct {
	Id            string            `json:"id"`
	Type          string            `json:"type"`
	Summary       *string           `json:"summary"`
	InReplyTo     *string           `json:"inReplyTo"`
	Published     time.Time         `json:"published"`
	Url           string            `json:"url"`
	To            []string          `json:"to"`
	Cc            []string          `json:"cc"`
	Sensitive     bool              `json:"sensitive"`
	Conversation  string            `json:"conversation"`
	Content       string            `json:"content"`
	ContentMap    map[string]string `json:"contentMap"`
	DirectMessage bool              `json:"directMessage"`
	Attachments   []Attachment      `json:"attachment"`
	RawTags       []json.RawMessage `json:"tag"`
}

const ASPublic = "https://www.w3.org/ns/activitystreams#Public"

func (o *Object) Visibility() string {
	for _, url := range o.To {
		if url == ASPublic {
			return "public"
		}
	}
	for _, url := range o.Cc {
		if url == ASPublic {
			return "unlisted"
		}
	}
	if o.DirectMessage {
		return "direct"
	}
	return "private"
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
// isn't the public or followers collections.
func (o *Object) TargetsSpecificUsersInToOrCc(a *Actor) bool {
	for _, url := range o.To {
		if url != ASPublic && url != a.Followers {
			continue
		}
	}
	for _, url := range o.Cc {
		if url != ASPublic && url != a.Followers {
			return true
		}
	}
	return false
}

type Attachment struct {
	Type          string    `json:"type"`
	MediaType     string    `json:"mediaType"`
	Url           string    `json:"url"`
	Name          string    `json:"name"`
	RawFocalPoint []float64 `json:"focalPoint"`
}

func (a Attachment) FocalPointX() float64 {
	if len(a.RawFocalPoint) == 2 {
		return a.RawFocalPoint[0]
	}
	return 0
}

func (a Attachment) FocalPointY() float64 {
	if len(a.RawFocalPoint) == 2 {
		return a.RawFocalPoint[1]
	}
	return 0
}

func (o *Object) Tags() []Tag {
	tags := make([]Tag, 0, len(o.RawTags))
	for _, rawTag := range o.RawTags {
		var mentionOrHashtag MentionOrHashtag
		var emoji Emoji
		if err := json.Unmarshal(rawTag, &mentionOrHashtag); err == nil {
			tags = append(tags, &mentionOrHashtag)
		} else if err := json.Unmarshal(rawTag, &emoji); err == nil {
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
