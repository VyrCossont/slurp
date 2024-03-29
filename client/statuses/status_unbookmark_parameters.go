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

// NewStatusUnbookmarkParams creates a new StatusUnbookmarkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStatusUnbookmarkParams() *StatusUnbookmarkParams {
	return &StatusUnbookmarkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStatusUnbookmarkParamsWithTimeout creates a new StatusUnbookmarkParams object
// with the ability to set a timeout on a request.
func NewStatusUnbookmarkParamsWithTimeout(timeout time.Duration) *StatusUnbookmarkParams {
	return &StatusUnbookmarkParams{
		timeout: timeout,
	}
}

// NewStatusUnbookmarkParamsWithContext creates a new StatusUnbookmarkParams object
// with the ability to set a context for a request.
func NewStatusUnbookmarkParamsWithContext(ctx context.Context) *StatusUnbookmarkParams {
	return &StatusUnbookmarkParams{
		Context: ctx,
	}
}

// NewStatusUnbookmarkParamsWithHTTPClient creates a new StatusUnbookmarkParams object
// with the ability to set a custom HTTPClient for a request.
func NewStatusUnbookmarkParamsWithHTTPClient(client *http.Client) *StatusUnbookmarkParams {
	return &StatusUnbookmarkParams{
		HTTPClient: client,
	}
}

/*
StatusUnbookmarkParams contains all the parameters to send to the API endpoint

	for the status unbookmark operation.

	Typically these are written to a http.Request.
*/
type StatusUnbookmarkParams struct {

	/* ID.

	   Target status ID.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the status unbookmark params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StatusUnbookmarkParams) WithDefaults() *StatusUnbookmarkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the status unbookmark params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StatusUnbookmarkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the status unbookmark params
func (o *StatusUnbookmarkParams) WithTimeout(timeout time.Duration) *StatusUnbookmarkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the status unbookmark params
func (o *StatusUnbookmarkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the status unbookmark params
func (o *StatusUnbookmarkParams) WithContext(ctx context.Context) *StatusUnbookmarkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the status unbookmark params
func (o *StatusUnbookmarkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the status unbookmark params
func (o *StatusUnbookmarkParams) WithHTTPClient(client *http.Client) *StatusUnbookmarkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the status unbookmark params
func (o *StatusUnbookmarkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the status unbookmark params
func (o *StatusUnbookmarkParams) WithID(id string) *StatusUnbookmarkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the status unbookmark params
func (o *StatusUnbookmarkParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *StatusUnbookmarkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
