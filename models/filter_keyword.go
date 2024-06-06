// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FilterKeyword FilterKeyword represents text to filter within a v2 filter.
//
// swagger:model FilterKeyword
type FilterKeyword struct {

	// The ID of the filter keyword entry in the database.
	ID string `json:"id,omitempty"`

	// The text to be filtered.
	// Example: fnord
	Keyword string `json:"keyword,omitempty"`

	// Should the filter keyword consider word boundaries?
	// Example: true
	WholeWord bool `json:"whole_word,omitempty"`
}

// Validate validates this filter keyword
func (m *FilterKeyword) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this filter keyword based on context it is used
func (m *FilterKeyword) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FilterKeyword) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilterKeyword) UnmarshalBinary(b []byte) error {
	var res FilterKeyword
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
