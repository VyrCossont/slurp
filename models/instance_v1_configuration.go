// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InstanceV1Configuration InstanceV1Configuration models instance configuration parameters.
//
// swagger:model InstanceV1Configuration
type InstanceV1Configuration struct {

	// accounts
	Accounts *InstanceConfigurationAccounts `json:"accounts,omitempty"`

	// emojis
	Emojis *InstanceConfigurationEmojis `json:"emojis,omitempty"`

	// media attachments
	MediaAttachments *InstanceConfigurationMediaAttachments `json:"media_attachments,omitempty"`

	// polls
	Polls *InstanceConfigurationPolls `json:"polls,omitempty"`

	// statuses
	Statuses *InstanceConfigurationStatuses `json:"statuses,omitempty"`
}

// Validate validates this instance v1 configuration
func (m *InstanceV1Configuration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmojis(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaAttachments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolls(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatuses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceV1Configuration) validateAccounts(formats strfmt.Registry) error {
	if swag.IsZero(m.Accounts) { // not required
		return nil
	}

	if m.Accounts != nil {
		if err := m.Accounts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accounts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accounts")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceV1Configuration) validateEmojis(formats strfmt.Registry) error {
	if swag.IsZero(m.Emojis) { // not required
		return nil
	}

	if m.Emojis != nil {
		if err := m.Emojis.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emojis")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("emojis")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceV1Configuration) validateMediaAttachments(formats strfmt.Registry) error {
	if swag.IsZero(m.MediaAttachments) { // not required
		return nil
	}

	if m.MediaAttachments != nil {
		if err := m.MediaAttachments.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("media_attachments")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("media_attachments")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceV1Configuration) validatePolls(formats strfmt.Registry) error {
	if swag.IsZero(m.Polls) { // not required
		return nil
	}

	if m.Polls != nil {
		if err := m.Polls.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("polls")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("polls")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceV1Configuration) validateStatuses(formats strfmt.Registry) error {
	if swag.IsZero(m.Statuses) { // not required
		return nil
	}

	if m.Statuses != nil {
		if err := m.Statuses.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statuses")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statuses")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this instance v1 configuration based on the context it is used
func (m *InstanceV1Configuration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmojis(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMediaAttachments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolls(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatuses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceV1Configuration) contextValidateAccounts(ctx context.Context, formats strfmt.Registry) error {

	if m.Accounts != nil {

		if swag.IsZero(m.Accounts) { // not required
			return nil
		}

		if err := m.Accounts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accounts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accounts")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceV1Configuration) contextValidateEmojis(ctx context.Context, formats strfmt.Registry) error {

	if m.Emojis != nil {

		if swag.IsZero(m.Emojis) { // not required
			return nil
		}

		if err := m.Emojis.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emojis")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("emojis")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceV1Configuration) contextValidateMediaAttachments(ctx context.Context, formats strfmt.Registry) error {

	if m.MediaAttachments != nil {

		if swag.IsZero(m.MediaAttachments) { // not required
			return nil
		}

		if err := m.MediaAttachments.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("media_attachments")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("media_attachments")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceV1Configuration) contextValidatePolls(ctx context.Context, formats strfmt.Registry) error {

	if m.Polls != nil {

		if swag.IsZero(m.Polls) { // not required
			return nil
		}

		if err := m.Polls.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("polls")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("polls")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceV1Configuration) contextValidateStatuses(ctx context.Context, formats strfmt.Registry) error {

	if m.Statuses != nil {

		if swag.IsZero(m.Statuses) { // not required
			return nil
		}

		if err := m.Statuses.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statuses")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statuses")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceV1Configuration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceV1Configuration) UnmarshalBinary(b []byte) error {
	var res InstanceV1Configuration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
