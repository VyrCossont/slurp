// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewEmojiUpdateParams creates a new EmojiUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEmojiUpdateParams() *EmojiUpdateParams {
	return &EmojiUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEmojiUpdateParamsWithTimeout creates a new EmojiUpdateParams object
// with the ability to set a timeout on a request.
func NewEmojiUpdateParamsWithTimeout(timeout time.Duration) *EmojiUpdateParams {
	return &EmojiUpdateParams{
		timeout: timeout,
	}
}

// NewEmojiUpdateParamsWithContext creates a new EmojiUpdateParams object
// with the ability to set a context for a request.
func NewEmojiUpdateParamsWithContext(ctx context.Context) *EmojiUpdateParams {
	return &EmojiUpdateParams{
		Context: ctx,
	}
}

// NewEmojiUpdateParamsWithHTTPClient creates a new EmojiUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewEmojiUpdateParamsWithHTTPClient(client *http.Client) *EmojiUpdateParams {
	return &EmojiUpdateParams{
		HTTPClient: client,
	}
}

/*
EmojiUpdateParams contains all the parameters to send to the API endpoint

	for the emoji update operation.

	Typically these are written to a http.Request.
*/
type EmojiUpdateParams struct {

	/* Category.

	   Category in which to place the emoji. If a category with the given name doesn't exist yet, it will be created.
	*/
	Category *string

	/* ID.

	   The id of the emoji.
	*/
	ID string

	/* Image.

	   A new png or gif image to use for the emoji. Animated pngs work too! To ensure compatibility with other fedi implementations, emoji size limit is 50kb by default. Works for LOCAL emojis only.
	*/
	Image runtime.NamedReadCloser

	/* Shortcode.

	   The code to use for the emoji, which will be used by instance denizens to select it. This must be unique on the instance. Works for the `copy` action type only.
	*/
	Shortcode *string

	/* Type.

	     Type of action to be taken. One of: (`disable`, `copy`, `modify`).
	For REMOTE emojis, `copy` or `disable` are supported.
	For LOCAL emojis, only `modify` is supported.
	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the emoji update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmojiUpdateParams) WithDefaults() *EmojiUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the emoji update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmojiUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the emoji update params
func (o *EmojiUpdateParams) WithTimeout(timeout time.Duration) *EmojiUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the emoji update params
func (o *EmojiUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the emoji update params
func (o *EmojiUpdateParams) WithContext(ctx context.Context) *EmojiUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the emoji update params
func (o *EmojiUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the emoji update params
func (o *EmojiUpdateParams) WithHTTPClient(client *http.Client) *EmojiUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the emoji update params
func (o *EmojiUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategory adds the category to the emoji update params
func (o *EmojiUpdateParams) WithCategory(category *string) *EmojiUpdateParams {
	o.SetCategory(category)
	return o
}

// SetCategory adds the category to the emoji update params
func (o *EmojiUpdateParams) SetCategory(category *string) {
	o.Category = category
}

// WithID adds the id to the emoji update params
func (o *EmojiUpdateParams) WithID(id string) *EmojiUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the emoji update params
func (o *EmojiUpdateParams) SetID(id string) {
	o.ID = id
}

// WithImage adds the image to the emoji update params
func (o *EmojiUpdateParams) WithImage(image runtime.NamedReadCloser) *EmojiUpdateParams {
	o.SetImage(image)
	return o
}

// SetImage adds the image to the emoji update params
func (o *EmojiUpdateParams) SetImage(image runtime.NamedReadCloser) {
	o.Image = image
}

// WithShortcode adds the shortcode to the emoji update params
func (o *EmojiUpdateParams) WithShortcode(shortcode *string) *EmojiUpdateParams {
	o.SetShortcode(shortcode)
	return o
}

// SetShortcode adds the shortcode to the emoji update params
func (o *EmojiUpdateParams) SetShortcode(shortcode *string) {
	o.Shortcode = shortcode
}

// WithType adds the typeVar to the emoji update params
func (o *EmojiUpdateParams) WithType(typeVar string) *EmojiUpdateParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the emoji update params
func (o *EmojiUpdateParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *EmojiUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Category != nil {

		// form param category
		var frCategory string
		if o.Category != nil {
			frCategory = *o.Category
		}
		fCategory := frCategory
		if fCategory != "" {
			if err := r.SetFormParam("category", fCategory); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Image != nil {

		if o.Image != nil {
			// form file param image
			if err := r.SetFileParam("image", o.Image); err != nil {
				return err
			}
		}
	}

	if o.Shortcode != nil {

		// form param shortcode
		var frShortcode string
		if o.Shortcode != nil {
			frShortcode = *o.Shortcode
		}
		fShortcode := frShortcode
		if fShortcode != "" {
			if err := r.SetFormParam("shortcode", fShortcode); err != nil {
				return err
			}
		}
	}

	// form param type
	frType := o.Type
	fType := frType
	if fType != "" {
		if err := r.SetFormParam("type", fType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
