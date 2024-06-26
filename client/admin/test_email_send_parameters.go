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

// NewTestEmailSendParams creates a new TestEmailSendParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTestEmailSendParams() *TestEmailSendParams {
	return &TestEmailSendParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTestEmailSendParamsWithTimeout creates a new TestEmailSendParams object
// with the ability to set a timeout on a request.
func NewTestEmailSendParamsWithTimeout(timeout time.Duration) *TestEmailSendParams {
	return &TestEmailSendParams{
		timeout: timeout,
	}
}

// NewTestEmailSendParamsWithContext creates a new TestEmailSendParams object
// with the ability to set a context for a request.
func NewTestEmailSendParamsWithContext(ctx context.Context) *TestEmailSendParams {
	return &TestEmailSendParams{
		Context: ctx,
	}
}

// NewTestEmailSendParamsWithHTTPClient creates a new TestEmailSendParams object
// with the ability to set a custom HTTPClient for a request.
func NewTestEmailSendParamsWithHTTPClient(client *http.Client) *TestEmailSendParams {
	return &TestEmailSendParams{
		HTTPClient: client,
	}
}

/*
TestEmailSendParams contains all the parameters to send to the API endpoint

	for the test email send operation.

	Typically these are written to a http.Request.
*/
type TestEmailSendParams struct {

	/* Email.

	   The email address that the test email should be sent to.
	*/
	Email string

	/* Message.

	   Optional message to include in the email.
	*/
	Message *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the test email send params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestEmailSendParams) WithDefaults() *TestEmailSendParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the test email send params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestEmailSendParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the test email send params
func (o *TestEmailSendParams) WithTimeout(timeout time.Duration) *TestEmailSendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test email send params
func (o *TestEmailSendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test email send params
func (o *TestEmailSendParams) WithContext(ctx context.Context) *TestEmailSendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test email send params
func (o *TestEmailSendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test email send params
func (o *TestEmailSendParams) WithHTTPClient(client *http.Client) *TestEmailSendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test email send params
func (o *TestEmailSendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmail adds the email to the test email send params
func (o *TestEmailSendParams) WithEmail(email string) *TestEmailSendParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the test email send params
func (o *TestEmailSendParams) SetEmail(email string) {
	o.Email = email
}

// WithMessage adds the message to the test email send params
func (o *TestEmailSendParams) WithMessage(message *string) *TestEmailSendParams {
	o.SetMessage(message)
	return o
}

// SetMessage adds the message to the test email send params
func (o *TestEmailSendParams) SetMessage(message *string) {
	o.Message = message
}

// WriteToRequest writes these params to a swagger request
func (o *TestEmailSendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param email
	frEmail := o.Email
	fEmail := frEmail
	if fEmail != "" {
		if err := r.SetFormParam("email", fEmail); err != nil {
			return err
		}
	}

	if o.Message != nil {

		// form param message
		var frMessage string
		if o.Message != nil {
			frMessage = *o.Message
		}
		fMessage := frMessage
		if fMessage != "" {
			if err := r.SetFormParam("message", fMessage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
