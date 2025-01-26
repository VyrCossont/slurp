// Code generated by go-swagger; DO NOT EDIT.

package mutes

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
	"github.com/go-openapi/swag"
)

// NewMutesGetParams creates a new MutesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMutesGetParams() *MutesGetParams {
	return &MutesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMutesGetParamsWithTimeout creates a new MutesGetParams object
// with the ability to set a timeout on a request.
func NewMutesGetParamsWithTimeout(timeout time.Duration) *MutesGetParams {
	return &MutesGetParams{
		timeout: timeout,
	}
}

// NewMutesGetParamsWithContext creates a new MutesGetParams object
// with the ability to set a context for a request.
func NewMutesGetParamsWithContext(ctx context.Context) *MutesGetParams {
	return &MutesGetParams{
		Context: ctx,
	}
}

// NewMutesGetParamsWithHTTPClient creates a new MutesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewMutesGetParamsWithHTTPClient(client *http.Client) *MutesGetParams {
	return &MutesGetParams{
		HTTPClient: client,
	}
}

/*
MutesGetParams contains all the parameters to send to the API endpoint

	for the mutes get operation.

	Typically these are written to a http.Request.
*/
type MutesGetParams struct {

	/* Limit.

	   Number of muted accounts to return.

	   Default: 40
	*/
	Limit *int64

	/* MaxID.

	   Return only muted accounts *OLDER* than the given max ID. The muted account with the specified ID will not be included in the response. NOTE: the ID is of the internal mute, NOT any of the returned accounts.
	*/
	MaxID *string

	/* MinID.

	   Return only muted accounts *IMMEDIATELY NEWER* than the given min ID. The muted account with the specified ID will not be included in the response. NOTE: the ID is of the internal mute, NOT any of the returned accounts.
	*/
	MinID *string

	/* SinceID.

	   Return only muted accounts *NEWER* than the given since ID. The muted account with the specified ID will not be included in the response. NOTE: the ID is of the internal mute, NOT any of the returned accounts.
	*/
	SinceID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the mutes get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MutesGetParams) WithDefaults() *MutesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the mutes get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MutesGetParams) SetDefaults() {
	var (
		limitDefault = int64(40)
	)

	val := MutesGetParams{
		Limit: &limitDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the mutes get params
func (o *MutesGetParams) WithTimeout(timeout time.Duration) *MutesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the mutes get params
func (o *MutesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the mutes get params
func (o *MutesGetParams) WithContext(ctx context.Context) *MutesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the mutes get params
func (o *MutesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the mutes get params
func (o *MutesGetParams) WithHTTPClient(client *http.Client) *MutesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the mutes get params
func (o *MutesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the mutes get params
func (o *MutesGetParams) WithLimit(limit *int64) *MutesGetParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the mutes get params
func (o *MutesGetParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMaxID adds the maxID to the mutes get params
func (o *MutesGetParams) WithMaxID(maxID *string) *MutesGetParams {
	o.SetMaxID(maxID)
	return o
}

// SetMaxID adds the maxId to the mutes get params
func (o *MutesGetParams) SetMaxID(maxID *string) {
	o.MaxID = maxID
}

// WithMinID adds the minID to the mutes get params
func (o *MutesGetParams) WithMinID(minID *string) *MutesGetParams {
	o.SetMinID(minID)
	return o
}

// SetMinID adds the minId to the mutes get params
func (o *MutesGetParams) SetMinID(minID *string) {
	o.MinID = minID
}

// WithSinceID adds the sinceID to the mutes get params
func (o *MutesGetParams) WithSinceID(sinceID *string) *MutesGetParams {
	o.SetSinceID(sinceID)
	return o
}

// SetSinceID adds the sinceId to the mutes get params
func (o *MutesGetParams) SetSinceID(sinceID *string) {
	o.SinceID = sinceID
}

// WriteToRequest writes these params to a swagger request
func (o *MutesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.MaxID != nil {

		// query param max_id
		var qrMaxID string

		if o.MaxID != nil {
			qrMaxID = *o.MaxID
		}
		qMaxID := qrMaxID
		if qMaxID != "" {

			if err := r.SetQueryParam("max_id", qMaxID); err != nil {
				return err
			}
		}
	}

	if o.MinID != nil {

		// query param min_id
		var qrMinID string

		if o.MinID != nil {
			qrMinID = *o.MinID
		}
		qMinID := qrMinID
		if qMinID != "" {

			if err := r.SetQueryParam("min_id", qMinID); err != nil {
				return err
			}
		}
	}

	if o.SinceID != nil {

		// query param since_id
		var qrSinceID string

		if o.SinceID != nil {
			qrSinceID = *o.SinceID
		}
		qSinceID := qrSinceID
		if qSinceID != "" {

			if err := r.SetQueryParam("since_id", qSinceID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}