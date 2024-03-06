// Code generated by go-swagger; DO NOT EDIT.

package filters

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

// NewFilterV1DeleteParams creates a new FilterV1DeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFilterV1DeleteParams() *FilterV1DeleteParams {
	return &FilterV1DeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFilterV1DeleteParamsWithTimeout creates a new FilterV1DeleteParams object
// with the ability to set a timeout on a request.
func NewFilterV1DeleteParamsWithTimeout(timeout time.Duration) *FilterV1DeleteParams {
	return &FilterV1DeleteParams{
		timeout: timeout,
	}
}

// NewFilterV1DeleteParamsWithContext creates a new FilterV1DeleteParams object
// with the ability to set a context for a request.
func NewFilterV1DeleteParamsWithContext(ctx context.Context) *FilterV1DeleteParams {
	return &FilterV1DeleteParams{
		Context: ctx,
	}
}

// NewFilterV1DeleteParamsWithHTTPClient creates a new FilterV1DeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewFilterV1DeleteParamsWithHTTPClient(client *http.Client) *FilterV1DeleteParams {
	return &FilterV1DeleteParams{
		HTTPClient: client,
	}
}

/*
FilterV1DeleteParams contains all the parameters to send to the API endpoint

	for the filter v1 delete operation.

	Typically these are written to a http.Request.
*/
type FilterV1DeleteParams struct {

	/* ID.

	   ID of the list
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the filter v1 delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FilterV1DeleteParams) WithDefaults() *FilterV1DeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the filter v1 delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FilterV1DeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the filter v1 delete params
func (o *FilterV1DeleteParams) WithTimeout(timeout time.Duration) *FilterV1DeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the filter v1 delete params
func (o *FilterV1DeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the filter v1 delete params
func (o *FilterV1DeleteParams) WithContext(ctx context.Context) *FilterV1DeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the filter v1 delete params
func (o *FilterV1DeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the filter v1 delete params
func (o *FilterV1DeleteParams) WithHTTPClient(client *http.Client) *FilterV1DeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the filter v1 delete params
func (o *FilterV1DeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the filter v1 delete params
func (o *FilterV1DeleteParams) WithID(id string) *FilterV1DeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the filter v1 delete params
func (o *FilterV1DeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FilterV1DeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
