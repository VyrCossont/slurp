// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SwaggerCollectionPage SwaggerCollectionPage represents one page of a collection.
//
// swagger:model SwaggerCollectionPage
type SwaggerCollectionPage struct {

	// ActivityStreams ID.
	// Example: https://example.org/users/some_user/statuses/106717595988259568/replies?page=true
	ID string `json:"id,omitempty"`

	// Items on this page.
	// Example: ["https://example.org/users/some_other_user/statuses/086417595981111564","https://another.example.com/users/another_user/statuses/01FCN8XDV3YG7B4R42QA6YQZ9R"]
	Items []string `json:"items"`

	// Link to the next page.
	// Example: https://example.org/users/some_user/statuses/106717595988259568/replies?only_other_accounts=true\u0026page=true
	Next string `json:"next,omitempty"`

	// Collection this page belongs to.
	// Example: https://example.org/users/some_user/statuses/106717595988259568/replies
	PartOf string `json:"partOf,omitempty"`

	// ActivityStreams type.
	// Example: CollectionPage
	Type string `json:"type,omitempty"`
}

// Validate validates this swagger collection page
func (m *SwaggerCollectionPage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this swagger collection page based on context it is used
func (m *SwaggerCollectionPage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SwaggerCollectionPage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SwaggerCollectionPage) UnmarshalBinary(b []byte) error {
	var res SwaggerCollectionPage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
