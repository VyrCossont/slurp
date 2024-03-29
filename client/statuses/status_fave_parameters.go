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

// NewStatusFaveParams creates a new StatusFaveParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStatusFaveParams() *StatusFaveParams {
	return &StatusFaveParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStatusFaveParamsWithTimeout creates a new StatusFaveParams object
// with the ability to set a timeout on a request.
func NewStatusFaveParamsWithTimeout(timeout time.Duration) *StatusFaveParams {
	return &StatusFaveParams{
		timeout: timeout,
	}
}

// NewStatusFaveParamsWithContext creates a new StatusFaveParams object
// with the ability to set a context for a request.
func NewStatusFaveParamsWithContext(ctx context.Context) *StatusFaveParams {
	return &StatusFaveParams{
		Context: ctx,
	}
}

// NewStatusFaveParamsWithHTTPClient creates a new StatusFaveParams object
// with the ability to set a custom HTTPClient for a request.
func NewStatusFaveParamsWithHTTPClient(client *http.Client) *StatusFaveParams {
	return &StatusFaveParams{
		HTTPClient: client,
	}
}

/*
StatusFaveParams contains all the parameters to send to the API endpoint

	for the status fave operation.

	Typically these are written to a http.Request.
*/
type StatusFaveParams struct {

	/* ID.

	   Target status ID.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the status fave params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StatusFaveParams) WithDefaults() *StatusFaveParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the status fave params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StatusFaveParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the status fave params
func (o *StatusFaveParams) WithTimeout(timeout time.Duration) *StatusFaveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the status fave params
func (o *StatusFaveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the status fave params
func (o *StatusFaveParams) WithContext(ctx context.Context) *StatusFaveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the status fave params
func (o *StatusFaveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the status fave params
func (o *StatusFaveParams) WithHTTPClient(client *http.Client) *StatusFaveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the status fave params
func (o *StatusFaveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the status fave params
func (o *StatusFaveParams) WithID(id string) *StatusFaveParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the status fave params
func (o *StatusFaveParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *StatusFaveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
