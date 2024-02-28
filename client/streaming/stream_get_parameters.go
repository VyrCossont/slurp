// Code generated by go-swagger; DO NOT EDIT.

package streaming

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

// NewStreamGetParams creates a new StreamGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStreamGetParams() *StreamGetParams {
	return &StreamGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStreamGetParamsWithTimeout creates a new StreamGetParams object
// with the ability to set a timeout on a request.
func NewStreamGetParamsWithTimeout(timeout time.Duration) *StreamGetParams {
	return &StreamGetParams{
		timeout: timeout,
	}
}

// NewStreamGetParamsWithContext creates a new StreamGetParams object
// with the ability to set a context for a request.
func NewStreamGetParamsWithContext(ctx context.Context) *StreamGetParams {
	return &StreamGetParams{
		Context: ctx,
	}
}

// NewStreamGetParamsWithHTTPClient creates a new StreamGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewStreamGetParamsWithHTTPClient(client *http.Client) *StreamGetParams {
	return &StreamGetParams{
		HTTPClient: client,
	}
}

/*
StreamGetParams contains all the parameters to send to the API endpoint

	for the stream get operation.

	Typically these are written to a http.Request.
*/
type StreamGetParams struct {

	/* AccessToken.

	   Access token for the requesting account.
	*/
	AccessToken string

	/* List.

	     ID of the list to subscribe to.
	Only used if stream type is 'list'.
	*/
	List *string

	/* Stream.

	     Type of stream to request.

	Options are:

	`user`: receive updates for the account's home timeline.
	`public`: receive updates for the public timeline.
	`public:local`: receive updates for the local timeline.
	`hashtag`: receive updates for a given hashtag.
	`hashtag:local`: receive local updates for a given hashtag.
	`list`: receive updates for a certain list of accounts.
	`direct`: receive updates for direct messages.
	*/
	Stream string

	/* Tag.

	     Name of the tag to subscribe to.
	Only used if stream type is 'hashtag' or 'hashtag:local'.
	*/
	Tag *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stream get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StreamGetParams) WithDefaults() *StreamGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stream get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StreamGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stream get params
func (o *StreamGetParams) WithTimeout(timeout time.Duration) *StreamGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stream get params
func (o *StreamGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stream get params
func (o *StreamGetParams) WithContext(ctx context.Context) *StreamGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stream get params
func (o *StreamGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stream get params
func (o *StreamGetParams) WithHTTPClient(client *http.Client) *StreamGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stream get params
func (o *StreamGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessToken adds the accessToken to the stream get params
func (o *StreamGetParams) WithAccessToken(accessToken string) *StreamGetParams {
	o.SetAccessToken(accessToken)
	return o
}

// SetAccessToken adds the accessToken to the stream get params
func (o *StreamGetParams) SetAccessToken(accessToken string) {
	o.AccessToken = accessToken
}

// WithList adds the list to the stream get params
func (o *StreamGetParams) WithList(list *string) *StreamGetParams {
	o.SetList(list)
	return o
}

// SetList adds the list to the stream get params
func (o *StreamGetParams) SetList(list *string) {
	o.List = list
}

// WithStream adds the stream to the stream get params
func (o *StreamGetParams) WithStream(stream string) *StreamGetParams {
	o.SetStream(stream)
	return o
}

// SetStream adds the stream to the stream get params
func (o *StreamGetParams) SetStream(stream string) {
	o.Stream = stream
}

// WithTag adds the tag to the stream get params
func (o *StreamGetParams) WithTag(tag *string) *StreamGetParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the stream get params
func (o *StreamGetParams) SetTag(tag *string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *StreamGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param access_token
	qrAccessToken := o.AccessToken
	qAccessToken := qrAccessToken
	if qAccessToken != "" {

		if err := r.SetQueryParam("access_token", qAccessToken); err != nil {
			return err
		}
	}

	if o.List != nil {

		// query param list
		var qrList string

		if o.List != nil {
			qrList = *o.List
		}
		qList := qrList
		if qList != "" {

			if err := r.SetQueryParam("list", qList); err != nil {
				return err
			}
		}
	}

	// query param stream
	qrStream := o.Stream
	qStream := qrStream
	if qStream != "" {

		if err := r.SetQueryParam("stream", qStream); err != nil {
			return err
		}
	}

	if o.Tag != nil {

		// query param tag
		var qrTag string

		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {

			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
