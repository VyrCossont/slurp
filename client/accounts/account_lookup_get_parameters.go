// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

// NewAccountLookupGetParams creates a new AccountLookupGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountLookupGetParams() *AccountLookupGetParams {
	return &AccountLookupGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountLookupGetParamsWithTimeout creates a new AccountLookupGetParams object
// with the ability to set a timeout on a request.
func NewAccountLookupGetParamsWithTimeout(timeout time.Duration) *AccountLookupGetParams {
	return &AccountLookupGetParams{
		timeout: timeout,
	}
}

// NewAccountLookupGetParamsWithContext creates a new AccountLookupGetParams object
// with the ability to set a context for a request.
func NewAccountLookupGetParamsWithContext(ctx context.Context) *AccountLookupGetParams {
	return &AccountLookupGetParams{
		Context: ctx,
	}
}

// NewAccountLookupGetParamsWithHTTPClient creates a new AccountLookupGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountLookupGetParamsWithHTTPClient(client *http.Client) *AccountLookupGetParams {
	return &AccountLookupGetParams{
		HTTPClient: client,
	}
}

/*
AccountLookupGetParams contains all the parameters to send to the API endpoint

	for the account lookup get operation.

	Typically these are written to a http.Request.
*/
type AccountLookupGetParams struct {

	/* Acct.

	   The username or Webfinger address to lookup.
	*/
	Acct string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the account lookup get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountLookupGetParams) WithDefaults() *AccountLookupGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the account lookup get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountLookupGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the account lookup get params
func (o *AccountLookupGetParams) WithTimeout(timeout time.Duration) *AccountLookupGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account lookup get params
func (o *AccountLookupGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account lookup get params
func (o *AccountLookupGetParams) WithContext(ctx context.Context) *AccountLookupGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account lookup get params
func (o *AccountLookupGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account lookup get params
func (o *AccountLookupGetParams) WithHTTPClient(client *http.Client) *AccountLookupGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account lookup get params
func (o *AccountLookupGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcct adds the acct to the account lookup get params
func (o *AccountLookupGetParams) WithAcct(acct string) *AccountLookupGetParams {
	o.SetAcct(acct)
	return o
}

// SetAcct adds the acct to the account lookup get params
func (o *AccountLookupGetParams) SetAcct(acct string) {
	o.Acct = acct
}

// WriteToRequest writes these params to a swagger request
func (o *AccountLookupGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param acct
	qrAcct := o.Acct
	qAcct := qrAcct
	if qAcct != "" {

		if err := r.SetQueryParam("acct", qAcct); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
