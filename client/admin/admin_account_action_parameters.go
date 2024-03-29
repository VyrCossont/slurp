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

// NewAdminAccountActionParams creates a new AdminAccountActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminAccountActionParams() *AdminAccountActionParams {
	return &AdminAccountActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminAccountActionParamsWithTimeout creates a new AdminAccountActionParams object
// with the ability to set a timeout on a request.
func NewAdminAccountActionParamsWithTimeout(timeout time.Duration) *AdminAccountActionParams {
	return &AdminAccountActionParams{
		timeout: timeout,
	}
}

// NewAdminAccountActionParamsWithContext creates a new AdminAccountActionParams object
// with the ability to set a context for a request.
func NewAdminAccountActionParamsWithContext(ctx context.Context) *AdminAccountActionParams {
	return &AdminAccountActionParams{
		Context: ctx,
	}
}

// NewAdminAccountActionParamsWithHTTPClient creates a new AdminAccountActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminAccountActionParamsWithHTTPClient(client *http.Client) *AdminAccountActionParams {
	return &AdminAccountActionParams{
		HTTPClient: client,
	}
}

/*
AdminAccountActionParams contains all the parameters to send to the API endpoint

	for the admin account action operation.

	Typically these are written to a http.Request.
*/
type AdminAccountActionParams struct {

	/* ID.

	   ID of the account.
	*/
	ID string

	/* Text.

	   Optional text describing why this action was taken.
	*/
	Text *string

	/* Type.

	   Type of action to be taken, currently only supports `suspend`.
	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin account action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminAccountActionParams) WithDefaults() *AdminAccountActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin account action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminAccountActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin account action params
func (o *AdminAccountActionParams) WithTimeout(timeout time.Duration) *AdminAccountActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin account action params
func (o *AdminAccountActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin account action params
func (o *AdminAccountActionParams) WithContext(ctx context.Context) *AdminAccountActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin account action params
func (o *AdminAccountActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin account action params
func (o *AdminAccountActionParams) WithHTTPClient(client *http.Client) *AdminAccountActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin account action params
func (o *AdminAccountActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the admin account action params
func (o *AdminAccountActionParams) WithID(id string) *AdminAccountActionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the admin account action params
func (o *AdminAccountActionParams) SetID(id string) {
	o.ID = id
}

// WithText adds the text to the admin account action params
func (o *AdminAccountActionParams) WithText(text *string) *AdminAccountActionParams {
	o.SetText(text)
	return o
}

// SetText adds the text to the admin account action params
func (o *AdminAccountActionParams) SetText(text *string) {
	o.Text = text
}

// WithType adds the typeVar to the admin account action params
func (o *AdminAccountActionParams) WithType(typeVar string) *AdminAccountActionParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the admin account action params
func (o *AdminAccountActionParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *AdminAccountActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Text != nil {

		// form param text
		var frText string
		if o.Text != nil {
			frText = *o.Text
		}
		fText := frText
		if fText != "" {
			if err := r.SetFormParam("text", fText); err != nil {
				return err
			}
		}
	}

	// form param type
	frType := o.Type
	fType := frType
	if fType != "" {
		if err := r.SetFormParam("type", fType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
