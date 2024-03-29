// Code generated by go-swagger; DO NOT EDIT.

package custom_emojis

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

// NewCustomEmojisGetParams creates a new CustomEmojisGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomEmojisGetParams() *CustomEmojisGetParams {
	return &CustomEmojisGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomEmojisGetParamsWithTimeout creates a new CustomEmojisGetParams object
// with the ability to set a timeout on a request.
func NewCustomEmojisGetParamsWithTimeout(timeout time.Duration) *CustomEmojisGetParams {
	return &CustomEmojisGetParams{
		timeout: timeout,
	}
}

// NewCustomEmojisGetParamsWithContext creates a new CustomEmojisGetParams object
// with the ability to set a context for a request.
func NewCustomEmojisGetParamsWithContext(ctx context.Context) *CustomEmojisGetParams {
	return &CustomEmojisGetParams{
		Context: ctx,
	}
}

// NewCustomEmojisGetParamsWithHTTPClient creates a new CustomEmojisGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomEmojisGetParamsWithHTTPClient(client *http.Client) *CustomEmojisGetParams {
	return &CustomEmojisGetParams{
		HTTPClient: client,
	}
}

/*
CustomEmojisGetParams contains all the parameters to send to the API endpoint

	for the custom emojis get operation.

	Typically these are written to a http.Request.
*/
type CustomEmojisGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the custom emojis get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomEmojisGetParams) WithDefaults() *CustomEmojisGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the custom emojis get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomEmojisGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the custom emojis get params
func (o *CustomEmojisGetParams) WithTimeout(timeout time.Duration) *CustomEmojisGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom emojis get params
func (o *CustomEmojisGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom emojis get params
func (o *CustomEmojisGetParams) WithContext(ctx context.Context) *CustomEmojisGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom emojis get params
func (o *CustomEmojisGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom emojis get params
func (o *CustomEmojisGetParams) WithHTTPClient(client *http.Client) *CustomEmojisGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom emojis get params
func (o *CustomEmojisGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CustomEmojisGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
