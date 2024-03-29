// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InstanceV2Users Usage data related to users on this instance.
//
// swagger:model InstanceV2Users
type InstanceV2Users struct {

	// The number of active users in the past 4 weeks.
	// Currently not implemented: will always be 0.
	// Example: 0
	ActiveMonth int64 `json:"active_month,omitempty"`
}

// Validate validates this instance v2 users
func (m *InstanceV2Users) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this instance v2 users based on context it is used
func (m *InstanceV2Users) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InstanceV2Users) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceV2Users) UnmarshalBinary(b []byte) error {
	var res InstanceV2Users
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
