// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Status Status models a status or post.
//
// swagger:model Status
type Status struct {

	// This status has been bookmarked by the account viewing it.
	Bookmarked bool `json:"bookmarked,omitempty"`

	// The content of this status. Should be HTML, but might also be plaintext in some cases.
	// Example: \u003cp\u003eHey this is a status!\u003c/p\u003e
	Content string `json:"content,omitempty"`

	// The date when this status was created (ISO 8601 Datetime).
	// Example: 2021-07-30T09:20:25+00:00
	CreatedAt string `json:"created_at,omitempty"`

	// Custom emoji to be used when rendering status content.
	Emojis []*Emoji `json:"emojis"`

	// This status has been favourited by the account viewing it.
	Favourited bool `json:"favourited,omitempty"`

	// Number of favourites/likes this status has received, according to our instance.
	FavouritesCount int64 `json:"favourites_count,omitempty"`

	// ID of the status.
	// Example: 01FBVD42CQ3ZEEVMW180SBX03B
	ID string `json:"id,omitempty"`

	// ID of the account being replied to.
	// Example: 01FBVD42CQ3ZEEVMW180SBX03B
	InReplyToAccountID string `json:"in_reply_to_account_id,omitempty"`

	// ID of the status being replied to.
	// Example: 01FBVD42CQ3ZEEVMW180SBX03B
	InReplyToID string `json:"in_reply_to_id,omitempty"`

	// Primary language of this status (ISO 639 Part 1 two-letter language code).
	// Will be null if language is not known.
	// Example: en
	Language string `json:"language,omitempty"`

	// Media that is attached to this status.
	MediaAttachments []*Attachment `json:"media_attachments"`

	// Mentions of users within the status content.
	Mentions []*Mention `json:"mentions"`

	// Replies to this status have been muted by the account viewing it.
	Muted bool `json:"muted,omitempty"`

	// This status has been pinned by the account viewing it (only relevant for your own statuses).
	Pinned bool `json:"pinned,omitempty"`

	// This status has been boosted/reblogged by the account viewing it.
	Reblogged bool `json:"reblogged,omitempty"`

	// Number of times this status has been boosted/reblogged, according to our instance.
	ReblogsCount int64 `json:"reblogs_count,omitempty"`

	// Number of replies to this status, according to our instance.
	RepliesCount int64 `json:"replies_count,omitempty"`

	// Status contains sensitive content.
	// Example: false
	Sensitive bool `json:"sensitive,omitempty"`

	// Subject, summary, or content warning for the status.
	// Example: warning nsfw
	SpoilerText string `json:"spoiler_text,omitempty"`

	// Hashtags used within the status content.
	Tags []*Tag `json:"tags"`

	// Plain-text source of a status. Returned instead of content when status is deleted,
	// so the user may redraft from the source text without the client having to reverse-engineer
	// the original text from the HTML content.
	Text string `json:"text,omitempty"`

	// ActivityPub URI of the status. Equivalent to the status's activitypub ID.
	// Example: https://example.org/users/some_user/statuses/01FBVD42CQ3ZEEVMW180SBX03B
	URI string `json:"uri,omitempty"`

	// The status's publicly available web URL. This link will only work if the visibility of the status is 'public'.
	// Example: https://example.org/@some_user/statuses/01FBVD42CQ3ZEEVMW180SBX03B
	URL string `json:"url,omitempty"`

	// Visibility of this status.
	// Example: unlisted
	Visibility string `json:"visibility,omitempty"`

	// account
	Account *Account `json:"account,omitempty"`

	// application
	Application *Application `json:"application,omitempty"`

	// card
	Card *Card `json:"card,omitempty"`

	// poll
	Poll *Poll `json:"poll,omitempty"`

	// reblog
	Reblog *StatusReblogged `json:"reblog,omitempty"`
}

// Validate validates this status
func (m *Status) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmojis(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaAttachments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMentions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePoll(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReblog(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Status) validateEmojis(formats strfmt.Registry) error {
	if swag.IsZero(m.Emojis) { // not required
		return nil
	}

	for i := 0; i < len(m.Emojis); i++ {
		if swag.IsZero(m.Emojis[i]) { // not required
			continue
		}

		if m.Emojis[i] != nil {
			if err := m.Emojis[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emojis" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("emojis" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Status) validateMediaAttachments(formats strfmt.Registry) error {
	if swag.IsZero(m.MediaAttachments) { // not required
		return nil
	}

	for i := 0; i < len(m.MediaAttachments); i++ {
		if swag.IsZero(m.MediaAttachments[i]) { // not required
			continue
		}

		if m.MediaAttachments[i] != nil {
			if err := m.MediaAttachments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("media_attachments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("media_attachments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Status) validateMentions(formats strfmt.Registry) error {
	if swag.IsZero(m.Mentions) { // not required
		return nil
	}

	for i := 0; i < len(m.Mentions); i++ {
		if swag.IsZero(m.Mentions[i]) { // not required
			continue
		}

		if m.Mentions[i] != nil {
			if err := m.Mentions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mentions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mentions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Status) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Status) validateAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("account")
			}
			return err
		}
	}

	return nil
}

func (m *Status) validateApplication(formats strfmt.Registry) error {
	if swag.IsZero(m.Application) { // not required
		return nil
	}

	if m.Application != nil {
		if err := m.Application.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

func (m *Status) validateCard(formats strfmt.Registry) error {
	if swag.IsZero(m.Card) { // not required
		return nil
	}

	if m.Card != nil {
		if err := m.Card.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("card")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("card")
			}
			return err
		}
	}

	return nil
}

func (m *Status) validatePoll(formats strfmt.Registry) error {
	if swag.IsZero(m.Poll) { // not required
		return nil
	}

	if m.Poll != nil {
		if err := m.Poll.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("poll")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("poll")
			}
			return err
		}
	}

	return nil
}

func (m *Status) validateReblog(formats strfmt.Registry) error {
	if swag.IsZero(m.Reblog) { // not required
		return nil
	}

	if m.Reblog != nil {
		if err := m.Reblog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reblog")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reblog")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this status based on the context it is used
func (m *Status) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEmojis(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMediaAttachments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMentions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateApplication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePoll(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReblog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Status) contextValidateEmojis(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Emojis); i++ {

		if m.Emojis[i] != nil {

			if swag.IsZero(m.Emojis[i]) { // not required
				return nil
			}

			if err := m.Emojis[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emojis" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("emojis" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Status) contextValidateMediaAttachments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MediaAttachments); i++ {

		if m.MediaAttachments[i] != nil {

			if swag.IsZero(m.MediaAttachments[i]) { // not required
				return nil
			}

			if err := m.MediaAttachments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("media_attachments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("media_attachments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Status) contextValidateMentions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Mentions); i++ {

		if m.Mentions[i] != nil {

			if swag.IsZero(m.Mentions[i]) { // not required
				return nil
			}

			if err := m.Mentions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mentions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mentions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Status) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {

			if swag.IsZero(m.Tags[i]) { // not required
				return nil
			}

			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Status) contextValidateAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.Account != nil {

		if swag.IsZero(m.Account) { // not required
			return nil
		}

		if err := m.Account.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("account")
			}
			return err
		}
	}

	return nil
}

func (m *Status) contextValidateApplication(ctx context.Context, formats strfmt.Registry) error {

	if m.Application != nil {

		if swag.IsZero(m.Application) { // not required
			return nil
		}

		if err := m.Application.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

func (m *Status) contextValidateCard(ctx context.Context, formats strfmt.Registry) error {

	if m.Card != nil {

		if swag.IsZero(m.Card) { // not required
			return nil
		}

		if err := m.Card.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("card")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("card")
			}
			return err
		}
	}

	return nil
}

func (m *Status) contextValidatePoll(ctx context.Context, formats strfmt.Registry) error {

	if m.Poll != nil {

		if swag.IsZero(m.Poll) { // not required
			return nil
		}

		if err := m.Poll.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("poll")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("poll")
			}
			return err
		}
	}

	return nil
}

func (m *Status) contextValidateReblog(ctx context.Context, formats strfmt.Registry) error {

	if m.Reblog != nil {

		if swag.IsZero(m.Reblog) { // not required
			return nil
		}

		if err := m.Reblog.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reblog")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reblog")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Status) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Status) UnmarshalBinary(b []byte) error {
	var res Status
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
