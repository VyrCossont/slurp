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

// NewStatusBoostedByParams creates a new StatusBoostedByParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStatusBoostedByParams() *StatusBoostedByParams {
	return &StatusBoostedByParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStatusBoostedByParamsWithTimeout creates a new StatusBoostedByParams object
// with the ability to set a timeout on a request.
func NewStatusBoostedByParamsWithTimeout(timeout time.Duration) *StatusBoostedByParams {
	return &StatusBoostedByParams{
		timeout: timeout,
	}
}

// NewStatusBoostedByParamsWithContext creates a new StatusBoostedByParams object
// with the ability to set a context for a request.
func NewStatusBoostedByParamsWithContext(ctx context.Context) *StatusBoostedByParams {
	return &StatusBoostedByParams{
		Context: ctx,
	}
}

// NewStatusBoostedByParamsWithHTTPClient creates a new StatusBoostedByParams object
// with the ability to set a custom HTTPClient for a request.
func NewStatusBoostedByParamsWithHTTPClient(client *http.Client) *StatusBoostedByParams {
	return &StatusBoostedByParams{
		HTTPClient: client,
	}
}

/*
StatusBoostedByParams contains all the parameters to send to the API endpoint

	for the status boosted by operation.

	Typically these are written to a http.Request.
*/
type StatusBoostedByParams struct {

	/* ID.

	   Target status ID.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the status boosted by params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StatusBoostedByParams) WithDefaults() *StatusBoostedByParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the status boosted by params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StatusBoostedByParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the status boosted by params
func (o *StatusBoostedByParams) WithTimeout(timeout time.Duration) *StatusBoostedByParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the status boosted by params
func (o *StatusBoostedByParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the status boosted by params
func (o *StatusBoostedByParams) WithContext(ctx context.Context) *StatusBoostedByParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the status boosted by params
func (o *StatusBoostedByParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the status boosted by params
func (o *StatusBoostedByParams) WithHTTPClient(client *http.Client) *StatusBoostedByParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the status boosted by params
func (o *StatusBoostedByParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the status boosted by params
func (o *StatusBoostedByParams) WithID(id string) *StatusBoostedByParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the status boosted by params
func (o *StatusBoostedByParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *StatusBoostedByParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
