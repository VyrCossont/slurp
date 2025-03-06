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
	"errors"
	"log/slog"
	"os"
	"path"

	"github.com/VyrCossont/slurp/client/statuses"
	"github.com/VyrCossont/slurp/internal/archive"
)

func NewArchiveReader(base string) (archive.ImportableReader, error) {
	// Require archive to be already uncompressed.
	stat, err := os.Stat(base)
	if err != nil {
		slog.Error("couldn't get archive info from filesystem", "path", base, "err", err)
		return nil, err
	}
	if !stat.IsDir() {
		return nil, errors.New("archive must be an uncompressed folder")
	}

	reader := &archiveReader{base: base}

	reader.actor, err = archive.ReadJSON[Actor](path.Join(base, "actor.json"))
	if err != nil {
		return nil, err
	}

	reader.outbox, err = archive.ReadJSON[Outbox](path.Join(base, "outbox.json"))
	if err != nil {
		return nil, err
	}

	return reader, nil
}

type archiveReader struct {
	base   string
	actor  *Actor
	outbox *Outbox
}

func (a *archiveReader) Statuses() []archive.ImportableStatus {
	return a.outbox.Notes()
}

func (a *archiveReader) Fetch(attachment archive.ImportableAttachment) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (o *Object) ID() string {
	//TODO implement me
	panic("implement me")
}

func (o *Object) HasPoll() bool {
	//TODO implement me
	panic("implement me")
}

func (o *Object) IsBoost() bool {
	//TODO implement me
	panic("implement me")
}

func (o *Object) IsReplyToOtherAccount() bool {
	//TODO implement me
	panic("implement me")
}

func (o *Object) MentionsOtherAccount() bool {
	//TODO implement me
	panic("implement me")
}

func (o *Object) ReplyToID() string {
	//TODO implement me
	panic("implement me")
}

func (o *Object) EmojiNames() []string {
	//TODO implement me
	panic("implement me")
}

func (o *Object) Attachments() []archive.ImportableAttachment {
	//TODO implement me
	panic("implement me")
}

func (o *Object) Params() *statuses.StatusCreateParams {
	//TODO implement me
	panic("implement me")
}
