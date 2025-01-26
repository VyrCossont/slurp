// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccountDisplayRole AccountDisplayRole models a public, displayable role of an account.
//
// This is a subset of AccountRole.
//
// swagger:model AccountDisplayRole
type AccountDisplayRole struct {

	// Color is a 6-digit CSS-style hex color code with leading `#`, or an empty string if this role has no color.
	// Since GotoSocial doesn't use role colors, we leave this empty.
	Color string `json:"color,omitempty"`

	// ID of the role.
	// Not used by GotoSocial, but we set it to the role name, just in case a client expects a unique ID.
	ID string `json:"id,omitempty"`

	// Name of the role.
	Name string `json:"name,omitempty"`
}

// Validate validates this account display role
func (m *AccountDisplayRole) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this account display role based on context it is used
func (m *AccountDisplayRole) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountDisplayRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountDisplayRole) UnmarshalBinary(b []byte) error {
	var res AccountDisplayRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
