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

// NewHeaderFilterBlockCreateParams creates a new HeaderFilterBlockCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeaderFilterBlockCreateParams() *HeaderFilterBlockCreateParams {
	return &HeaderFilterBlockCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeaderFilterBlockCreateParamsWithTimeout creates a new HeaderFilterBlockCreateParams object
// with the ability to set a timeout on a request.
func NewHeaderFilterBlockCreateParamsWithTimeout(timeout time.Duration) *HeaderFilterBlockCreateParams {
	return &HeaderFilterBlockCreateParams{
		timeout: timeout,
	}
}

// NewHeaderFilterBlockCreateParamsWithContext creates a new HeaderFilterBlockCreateParams object
// with the ability to set a context for a request.
func NewHeaderFilterBlockCreateParamsWithContext(ctx context.Context) *HeaderFilterBlockCreateParams {
	return &HeaderFilterBlockCreateParams{
		Context: ctx,
	}
}

// NewHeaderFilterBlockCreateParamsWithHTTPClient creates a new HeaderFilterBlockCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeaderFilterBlockCreateParamsWithHTTPClient(client *http.Client) *HeaderFilterBlockCreateParams {
	return &HeaderFilterBlockCreateParams{
		HTTPClient: client,
	}
}

/*
HeaderFilterBlockCreateParams contains all the parameters to send to the API endpoint

	for the header filter block create operation.

	Typically these are written to a http.Request.
*/
type HeaderFilterBlockCreateParams struct {

	/* Header.

	   The HTTP header to match against (e.g. User-Agent).
	*/
	Header string

	/* Regex.

	   The header value matching regular expression.
	*/
	Regex string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the header filter block create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeaderFilterBlockCreateParams) WithDefaults() *HeaderFilterBlockCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the header filter block create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeaderFilterBlockCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the header filter block create params
func (o *HeaderFilterBlockCreateParams) WithTimeout(timeout time.Duration) *HeaderFilterBlockCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the header filter block create params
func (o *HeaderFilterBlockCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the header filter block create params
func (o *HeaderFilterBlockCreateParams) WithContext(ctx context.Context) *HeaderFilterBlockCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the header filter block create params
func (o *HeaderFilterBlockCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the header filter block create params
func (o *HeaderFilterBlockCreateParams) WithHTTPClient(client *http.Client) *HeaderFilterBlockCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the header filter block create params
func (o *HeaderFilterBlockCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeader adds the header to the header filter block create params
func (o *HeaderFilterBlockCreateParams) WithHeader(header string) *HeaderFilterBlockCreateParams {
	o.SetHeader(header)
	return o
}

// SetHeader adds the header to the header filter block create params
func (o *HeaderFilterBlockCreateParams) SetHeader(header string) {
	o.Header = header
}

// WithRegex adds the regex to the header filter block create params
func (o *HeaderFilterBlockCreateParams) WithRegex(regex string) *HeaderFilterBlockCreateParams {
	o.SetRegex(regex)
	return o
}

// SetRegex adds the regex to the header filter block create params
func (o *HeaderFilterBlockCreateParams) SetRegex(regex string) {
	o.Regex = regex
}

// WriteToRequest writes these params to a swagger request
func (o *HeaderFilterBlockCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param header
	frHeader := o.Header
	fHeader := frHeader
	if fHeader != "" {
		if err := r.SetFormParam("header", fHeader); err != nil {
			return err
		}
	}

	// form param regex
	frRegex := o.Regex
	fRegex := frRegex
	if fRegex != "" {
		if err := r.SetFormParam("regex", fRegex); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
