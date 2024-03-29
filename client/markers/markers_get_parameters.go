// Code generated by go-swagger; DO NOT EDIT.

package markers

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

// NewMarkersGetParams creates a new MarkersGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMarkersGetParams() *MarkersGetParams {
	return &MarkersGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMarkersGetParamsWithTimeout creates a new MarkersGetParams object
// with the ability to set a timeout on a request.
func NewMarkersGetParamsWithTimeout(timeout time.Duration) *MarkersGetParams {
	return &MarkersGetParams{
		timeout: timeout,
	}
}

// NewMarkersGetParamsWithContext creates a new MarkersGetParams object
// with the ability to set a context for a request.
func NewMarkersGetParamsWithContext(ctx context.Context) *MarkersGetParams {
	return &MarkersGetParams{
		Context: ctx,
	}
}

// NewMarkersGetParamsWithHTTPClient creates a new MarkersGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewMarkersGetParamsWithHTTPClient(client *http.Client) *MarkersGetParams {
	return &MarkersGetParams{
		HTTPClient: client,
	}
}

/*
MarkersGetParams contains all the parameters to send to the API endpoint

	for the markers get operation.

	Typically these are written to a http.Request.
*/
type MarkersGetParams struct {

	/* Timeline.

	   Timelines to retrieve.
	*/
	Timeline []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the markers get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MarkersGetParams) WithDefaults() *MarkersGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the markers get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MarkersGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the markers get params
func (o *MarkersGetParams) WithTimeout(timeout time.Duration) *MarkersGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the markers get params
func (o *MarkersGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the markers get params
func (o *MarkersGetParams) WithContext(ctx context.Context) *MarkersGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the markers get params
func (o *MarkersGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the markers get params
func (o *MarkersGetParams) WithHTTPClient(client *http.Client) *MarkersGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the markers get params
func (o *MarkersGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTimeline adds the timeline to the markers get params
func (o *MarkersGetParams) WithTimeline(timeline []string) *MarkersGetParams {
	o.SetTimeline(timeline)
	return o
}

// SetTimeline adds the timeline to the markers get params
func (o *MarkersGetParams) SetTimeline(timeline []string) {
	o.Timeline = timeline
}

// WriteToRequest writes these params to a swagger request
func (o *MarkersGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Timeline != nil {

		// binding items for timeline
		joinedTimeline := o.bindParamTimeline(reg)

		// query array param timeline
		if err := r.SetQueryParam("timeline", joinedTimeline...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamMarkersGet binds the parameter timeline
func (o *MarkersGetParams) bindParamTimeline(formats strfmt.Registry) []string {
	timelineIR := o.Timeline

	var timelineIC []string
	for _, timelineIIR := range timelineIR { // explode []string

		timelineIIV := timelineIIR // string as string
		timelineIC = append(timelineIC, timelineIIV)
	}

	// items.CollectionFormat: ""
	timelineIS := swag.JoinByFormat(timelineIC, "")

	return timelineIS
}
