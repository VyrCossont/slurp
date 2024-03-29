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

// NewDomainKeysExpireParams creates a new DomainKeysExpireParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDomainKeysExpireParams() *DomainKeysExpireParams {
	return &DomainKeysExpireParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDomainKeysExpireParamsWithTimeout creates a new DomainKeysExpireParams object
// with the ability to set a timeout on a request.
func NewDomainKeysExpireParamsWithTimeout(timeout time.Duration) *DomainKeysExpireParams {
	return &DomainKeysExpireParams{
		timeout: timeout,
	}
}

// NewDomainKeysExpireParamsWithContext creates a new DomainKeysExpireParams object
// with the ability to set a context for a request.
func NewDomainKeysExpireParamsWithContext(ctx context.Context) *DomainKeysExpireParams {
	return &DomainKeysExpireParams{
		Context: ctx,
	}
}

// NewDomainKeysExpireParamsWithHTTPClient creates a new DomainKeysExpireParams object
// with the ability to set a custom HTTPClient for a request.
func NewDomainKeysExpireParamsWithHTTPClient(client *http.Client) *DomainKeysExpireParams {
	return &DomainKeysExpireParams{
		HTTPClient: client,
	}
}

/*
DomainKeysExpireParams contains all the parameters to send to the API endpoint

	for the domain keys expire operation.

	Typically these are written to a http.Request.
*/
type DomainKeysExpireParams struct {

	/* Domain.

	     Domain to expire keys for.
	Sample: example.org
	*/
	Domain *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the domain keys expire params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DomainKeysExpireParams) WithDefaults() *DomainKeysExpireParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the domain keys expire params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DomainKeysExpireParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the domain keys expire params
func (o *DomainKeysExpireParams) WithTimeout(timeout time.Duration) *DomainKeysExpireParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domain keys expire params
func (o *DomainKeysExpireParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domain keys expire params
func (o *DomainKeysExpireParams) WithContext(ctx context.Context) *DomainKeysExpireParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domain keys expire params
func (o *DomainKeysExpireParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domain keys expire params
func (o *DomainKeysExpireParams) WithHTTPClient(client *http.Client) *DomainKeysExpireParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domain keys expire params
func (o *DomainKeysExpireParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomain adds the domain to the domain keys expire params
func (o *DomainKeysExpireParams) WithDomain(domain *string) *DomainKeysExpireParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the domain keys expire params
func (o *DomainKeysExpireParams) SetDomain(domain *string) {
	o.Domain = domain
}

// WriteToRequest writes these params to a swagger request
func (o *DomainKeysExpireParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Domain != nil {

		// form param domain
		var frDomain string
		if o.Domain != nil {
			frDomain = *o.Domain
		}
		fDomain := frDomain
		if fDomain != "" {
			if err := r.SetFormParam("domain", fDomain); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
