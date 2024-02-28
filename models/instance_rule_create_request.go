// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InstanceRuleCreateRequest InstanceRuleCreateRequest represents a request to create a new instance rule, made through the admin API.
//
// swagger:model InstanceRuleCreateRequest
type InstanceRuleCreateRequest struct {

	// text
	Text string `json:"Text,omitempty"`
}

// Validate validates this instance rule create request
func (m *InstanceRuleCreateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this instance rule create request based on context it is used
func (m *InstanceRuleCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InstanceRuleCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceRuleCreateRequest) UnmarshalBinary(b []byte) error {
	var res InstanceRuleCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
