// Code generated by go-swagger; DO NOT EDIT.

package statuses

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

// NewStatusBookmarkParams creates a new StatusBookmarkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStatusBookmarkParams() *StatusBookmarkParams {
	return &StatusBookmarkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStatusBookmarkParamsWithTimeout creates a new StatusBookmarkParams object
// with the ability to set a timeout on a request.
func NewStatusBookmarkParamsWithTimeout(timeout time.Duration) *StatusBookmarkParams {
	return &StatusBookmarkParams{
		timeout: timeout,
	}
}

// NewStatusBookmarkParamsWithContext creates a new StatusBookmarkParams object
// with the ability to set a context for a request.
func NewStatusBookmarkParamsWithContext(ctx context.Context) *StatusBookmarkParams {
	return &StatusBookmarkParams{
		Context: ctx,
	}
}

// NewStatusBookmarkParamsWithHTTPClient creates a new StatusBookmarkParams object
// with the ability to set a custom HTTPClient for a request.
func NewStatusBookmarkParamsWithHTTPClient(client *http.Client) *StatusBookmarkParams {
	return &StatusBookmarkParams{
		HTTPClient: client,
	}
}

/*
StatusBookmarkParams contains all the parameters to send to the API endpoint

	for the status bookmark operation.

	Typically these are written to a http.Request.
*/
type StatusBookmarkParams struct {

	/* ID.

	   Target status ID.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the status bookmark params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StatusBookmarkParams) WithDefaults() *StatusBookmarkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the status bookmark params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StatusBookmarkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the status bookmark params
func (o *StatusBookmarkParams) WithTimeout(timeout time.Duration) *StatusBookmarkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the status bookmark params
func (o *StatusBookmarkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the status bookmark params
func (o *StatusBookmarkParams) WithContext(ctx context.Context) *StatusBookmarkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the status bookmark params
func (o *StatusBookmarkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the status bookmark params
func (o *StatusBookmarkParams) WithHTTPClient(client *http.Client) *StatusBookmarkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the status bookmark params
func (o *StatusBookmarkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the status bookmark params
func (o *StatusBookmarkParams) WithID(id string) *StatusBookmarkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the status bookmark params
func (o *StatusBookmarkParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *StatusBookmarkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
