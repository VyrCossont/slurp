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

// NewAccountUnfollowParams creates a new AccountUnfollowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountUnfollowParams() *AccountUnfollowParams {
	return &AccountUnfollowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountUnfollowParamsWithTimeout creates a new AccountUnfollowParams object
// with the ability to set a timeout on a request.
func NewAccountUnfollowParamsWithTimeout(timeout time.Duration) *AccountUnfollowParams {
	return &AccountUnfollowParams{
		timeout: timeout,
	}
}

// NewAccountUnfollowParamsWithContext creates a new AccountUnfollowParams object
// with the ability to set a context for a request.
func NewAccountUnfollowParamsWithContext(ctx context.Context) *AccountUnfollowParams {
	return &AccountUnfollowParams{
		Context: ctx,
	}
}

// NewAccountUnfollowParamsWithHTTPClient creates a new AccountUnfollowParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountUnfollowParamsWithHTTPClient(client *http.Client) *AccountUnfollowParams {
	return &AccountUnfollowParams{
		HTTPClient: client,
	}
}

/*
AccountUnfollowParams contains all the parameters to send to the API endpoint

	for the account unfollow operation.

	Typically these are written to a http.Request.
*/
type AccountUnfollowParams struct {

	/* ID.

	   The id of the account to unfollow.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the account unfollow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountUnfollowParams) WithDefaults() *AccountUnfollowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the account unfollow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountUnfollowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the account unfollow params
func (o *AccountUnfollowParams) WithTimeout(timeout time.Duration) *AccountUnfollowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account unfollow params
func (o *AccountUnfollowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account unfollow params
func (o *AccountUnfollowParams) WithContext(ctx context.Context) *AccountUnfollowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account unfollow params
func (o *AccountUnfollowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account unfollow params
func (o *AccountUnfollowParams) WithHTTPClient(client *http.Client) *AccountUnfollowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account unfollow params
func (o *AccountUnfollowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the account unfollow params
func (o *AccountUnfollowParams) WithID(id string) *AccountUnfollowParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the account unfollow params
func (o *AccountUnfollowParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AccountUnfollowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
