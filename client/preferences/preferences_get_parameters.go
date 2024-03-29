// Code generated by go-swagger; DO NOT EDIT.

package preferences

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

// NewPreferencesGetParams creates a new PreferencesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPreferencesGetParams() *PreferencesGetParams {
	return &PreferencesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPreferencesGetParamsWithTimeout creates a new PreferencesGetParams object
// with the ability to set a timeout on a request.
func NewPreferencesGetParamsWithTimeout(timeout time.Duration) *PreferencesGetParams {
	return &PreferencesGetParams{
		timeout: timeout,
	}
}

// NewPreferencesGetParamsWithContext creates a new PreferencesGetParams object
// with the ability to set a context for a request.
func NewPreferencesGetParamsWithContext(ctx context.Context) *PreferencesGetParams {
	return &PreferencesGetParams{
		Context: ctx,
	}
}

// NewPreferencesGetParamsWithHTTPClient creates a new PreferencesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPreferencesGetParamsWithHTTPClient(client *http.Client) *PreferencesGetParams {
	return &PreferencesGetParams{
		HTTPClient: client,
	}
}

/*
PreferencesGetParams contains all the parameters to send to the API endpoint

	for the preferences get operation.

	Typically these are written to a http.Request.
*/
type PreferencesGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the preferences get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PreferencesGetParams) WithDefaults() *PreferencesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the preferences get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PreferencesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the preferences get params
func (o *PreferencesGetParams) WithTimeout(timeout time.Duration) *PreferencesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the preferences get params
func (o *PreferencesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the preferences get params
func (o *PreferencesGetParams) WithContext(ctx context.Context) *PreferencesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the preferences get params
func (o *PreferencesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the preferences get params
func (o *PreferencesGetParams) WithHTTPClient(client *http.Client) *PreferencesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the preferences get params
func (o *PreferencesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PreferencesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
