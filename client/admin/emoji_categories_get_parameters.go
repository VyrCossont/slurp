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

// NewEmojiCategoriesGetParams creates a new EmojiCategoriesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEmojiCategoriesGetParams() *EmojiCategoriesGetParams {
	return &EmojiCategoriesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEmojiCategoriesGetParamsWithTimeout creates a new EmojiCategoriesGetParams object
// with the ability to set a timeout on a request.
func NewEmojiCategoriesGetParamsWithTimeout(timeout time.Duration) *EmojiCategoriesGetParams {
	return &EmojiCategoriesGetParams{
		timeout: timeout,
	}
}

// NewEmojiCategoriesGetParamsWithContext creates a new EmojiCategoriesGetParams object
// with the ability to set a context for a request.
func NewEmojiCategoriesGetParamsWithContext(ctx context.Context) *EmojiCategoriesGetParams {
	return &EmojiCategoriesGetParams{
		Context: ctx,
	}
}

// NewEmojiCategoriesGetParamsWithHTTPClient creates a new EmojiCategoriesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewEmojiCategoriesGetParamsWithHTTPClient(client *http.Client) *EmojiCategoriesGetParams {
	return &EmojiCategoriesGetParams{
		HTTPClient: client,
	}
}

/*
EmojiCategoriesGetParams contains all the parameters to send to the API endpoint

	for the emoji categories get operation.

	Typically these are written to a http.Request.
*/
type EmojiCategoriesGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the emoji categories get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmojiCategoriesGetParams) WithDefaults() *EmojiCategoriesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the emoji categories get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmojiCategoriesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the emoji categories get params
func (o *EmojiCategoriesGetParams) WithTimeout(timeout time.Duration) *EmojiCategoriesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the emoji categories get params
func (o *EmojiCategoriesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the emoji categories get params
func (o *EmojiCategoriesGetParams) WithContext(ctx context.Context) *EmojiCategoriesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the emoji categories get params
func (o *EmojiCategoriesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the emoji categories get params
func (o *EmojiCategoriesGetParams) WithHTTPClient(client *http.Client) *EmojiCategoriesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the emoji categories get params
func (o *EmojiCategoriesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *EmojiCategoriesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
