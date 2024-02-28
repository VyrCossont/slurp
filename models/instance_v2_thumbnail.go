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

// InstanceV2Thumbnail An image used to represent this instance.
//
// swagger:model InstanceV2Thumbnail
type InstanceV2Thumbnail struct {

	// A hash computed by the BlurHash algorithm, for generating colorful preview thumbnails when media has not been downloaded yet.
	// Key/value not set if no blurhash available.
	// Example: UeKUpFxuo~R%0nW;WCnhF6RjaJt757oJodS$
	Blurhash string `json:"blurhash,omitempty"`

	// Description of the instance thumbnail.
	// Key/value not set if no description available.
	// Example: picture of a cute lil' friendly sloth
	Description string `json:"thumbnail_description,omitempty"`

	// MIME type of the instance thumbnail.
	// Key/value not set if thumbnail image type unknown.
	// Example: image/png
	Type string `json:"thumbnail_type,omitempty"`

	// The URL for the thumbnail image.
	// Example: https://example.org/fileserver/01BPSX2MKCRVMD4YN4D71G9CP5/attachment/original/01H88X0KQ2DFYYDSWYP93VDJZA.png
	URL string `json:"url,omitempty"`

	// versions
	Versions *InstanceV2ThumbnailVersions `json:"versions,omitempty"`
}

// Validate validates this instance v2 thumbnail
func (m *InstanceV2Thumbnail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceV2Thumbnail) validateVersions(formats strfmt.Registry) error {
	if swag.IsZero(m.Versions) { // not required
		return nil
	}

	if m.Versions != nil {
		if err := m.Versions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("versions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this instance v2 thumbnail based on the context it is used
func (m *InstanceV2Thumbnail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVersions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceV2Thumbnail) contextValidateVersions(ctx context.Context, formats strfmt.Registry) error {

	if m.Versions != nil {

		if swag.IsZero(m.Versions) { // not required
			return nil
		}

		if err := m.Versions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("versions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceV2Thumbnail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceV2Thumbnail) UnmarshalBinary(b []byte) error {
	var res InstanceV2Thumbnail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
