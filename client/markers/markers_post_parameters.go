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
)

// NewMarkersPostParams creates a new MarkersPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMarkersPostParams() *MarkersPostParams {
	return &MarkersPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMarkersPostParamsWithTimeout creates a new MarkersPostParams object
// with the ability to set a timeout on a request.
func NewMarkersPostParamsWithTimeout(timeout time.Duration) *MarkersPostParams {
	return &MarkersPostParams{
		timeout: timeout,
	}
}

// NewMarkersPostParamsWithContext creates a new MarkersPostParams object
// with the ability to set a context for a request.
func NewMarkersPostParamsWithContext(ctx context.Context) *MarkersPostParams {
	return &MarkersPostParams{
		Context: ctx,
	}
}

// NewMarkersPostParamsWithHTTPClient creates a new MarkersPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewMarkersPostParamsWithHTTPClient(client *http.Client) *MarkersPostParams {
	return &MarkersPostParams{
		HTTPClient: client,
	}
}

/*
MarkersPostParams contains all the parameters to send to the API endpoint

	for the markers post operation.

	Typically these are written to a http.Request.
*/
type MarkersPostParams struct {

	/* HomeLastReadID.

	   Last status ID read on the home timeline.
	*/
	HomeLastReadID *string

	/* NotificationsLastReadID.

	   Last notification ID read on the notifications timeline.
	*/
	NotificationsLastReadID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the markers post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MarkersPostParams) WithDefaults() *MarkersPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the markers post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MarkersPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the markers post params
func (o *MarkersPostParams) WithTimeout(timeout time.Duration) *MarkersPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the markers post params
func (o *MarkersPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the markers post params
func (o *MarkersPostParams) WithContext(ctx context.Context) *MarkersPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the markers post params
func (o *MarkersPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the markers post params
func (o *MarkersPostParams) WithHTTPClient(client *http.Client) *MarkersPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the markers post params
func (o *MarkersPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHomeLastReadID adds the homeLastReadID to the markers post params
func (o *MarkersPostParams) WithHomeLastReadID(homeLastReadID *string) *MarkersPostParams {
	o.SetHomeLastReadID(homeLastReadID)
	return o
}

// SetHomeLastReadID adds the homeLastReadId to the markers post params
func (o *MarkersPostParams) SetHomeLastReadID(homeLastReadID *string) {
	o.HomeLastReadID = homeLastReadID
}

// WithNotificationsLastReadID adds the notificationsLastReadID to the markers post params
func (o *MarkersPostParams) WithNotificationsLastReadID(notificationsLastReadID *string) *MarkersPostParams {
	o.SetNotificationsLastReadID(notificationsLastReadID)
	return o
}

// SetNotificationsLastReadID adds the notificationsLastReadId to the markers post params
func (o *MarkersPostParams) SetNotificationsLastReadID(notificationsLastReadID *string) {
	o.NotificationsLastReadID = notificationsLastReadID
}

// WriteToRequest writes these params to a swagger request
func (o *MarkersPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.HomeLastReadID != nil {

		// form param home[last_read_id]
		var frHomeLastReadID string
		if o.HomeLastReadID != nil {
			frHomeLastReadID = *o.HomeLastReadID
		}
		fHomeLastReadID := frHomeLastReadID
		if fHomeLastReadID != "" {
			if err := r.SetFormParam("home[last_read_id]", fHomeLastReadID); err != nil {
				return err
			}
		}
	}

	if o.NotificationsLastReadID != nil {

		// form param notifications[last_read_id]
		var frNotificationsLastReadID string
		if o.NotificationsLastReadID != nil {
			frNotificationsLastReadID = *o.NotificationsLastReadID
		}
		fNotificationsLastReadID := frNotificationsLastReadID
		if fNotificationsLastReadID != "" {
			if err := r.SetFormParam("notifications[last_read_id]", fNotificationsLastReadID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
