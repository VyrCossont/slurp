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

// Report Report models a moderation report submitted to the instance, either via the client API or via the federated API.
//
// swagger:model Report
type Report struct {

	// Whether an action has been taken by an admin in response to this report.
	// Example: false
	ActionTaken bool `json:"action_taken,omitempty"`

	// If an action was taken, at what time was this done? (ISO 8601 Datetime)
	// Will be null if not set / no action yet taken.
	// Example: 2021-07-30T09:20:25+00:00
	ActionTakenAt string `json:"action_taken_at,omitempty"`

	// If an action was taken, what comment was made by the admin on the taken action?
	// Will be null if not set / no action yet taken.
	// Example: Account was suspended.
	ActionTakenComment string `json:"action_taken_comment,omitempty"`

	// Under what category was this report created?
	// Example: spam
	Category string `json:"category,omitempty"`

	// Comment submitted when the report was created.
	// Will be empty if no comment was submitted.
	// Example: This person has been harassing me.
	Comment string `json:"comment,omitempty"`

	// The date when this report was created (ISO 8601 Datetime).
	// Example: 2021-07-30T09:20:25+00:00
	CreatedAt string `json:"created_at,omitempty"`

	// Bool to indicate that report should be federated to remote instance.
	// Example: true
	Forwarded bool `json:"forwarded,omitempty"`

	// ID of the report.
	// Example: 01FBVD42CQ3ZEEVMW180SBX03B
	ID string `json:"id,omitempty"`

	// Array of rule IDs that were submitted along with this report.
	// Will be empty if no rule IDs were submitted.
	// Example: ["01GPBN5YDY6JKBWE44H7YQBDCQ","01GPBN65PDWSBPWVDD0SQCFFY3"]
	RuleIDs []string `json:"rule_ids"`

	// Array of IDs of statuses that were submitted along with this report.
	// Will be empty if no status IDs were submitted.
	// Example: ["01GPBN5YDY6JKBWE44H7YQBDCQ","01GPBN65PDWSBPWVDD0SQCFFY3"]
	StatusIDs []string `json:"status_ids"`

	// target account
	TargetAccount *Account `json:"target_account,omitempty"`
}

// Validate validates this report
func (m *Report) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTargetAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Report) validateTargetAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetAccount) { // not required
		return nil
	}

	if m.TargetAccount != nil {
		if err := m.TargetAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target_account")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target_account")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this report based on the context it is used
func (m *Report) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTargetAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Report) contextValidateTargetAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.TargetAccount != nil {

		if swag.IsZero(m.TargetAccount) { // not required
			return nil
		}

		if err := m.TargetAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target_account")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target_account")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Report) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Report) UnmarshalBinary(b []byte) error {
	var res Report
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
