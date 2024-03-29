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
	"github.com/go-openapi/swag"
)

// NewEmojisGetParams creates a new EmojisGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEmojisGetParams() *EmojisGetParams {
	return &EmojisGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEmojisGetParamsWithTimeout creates a new EmojisGetParams object
// with the ability to set a timeout on a request.
func NewEmojisGetParamsWithTimeout(timeout time.Duration) *EmojisGetParams {
	return &EmojisGetParams{
		timeout: timeout,
	}
}

// NewEmojisGetParamsWithContext creates a new EmojisGetParams object
// with the ability to set a context for a request.
func NewEmojisGetParamsWithContext(ctx context.Context) *EmojisGetParams {
	return &EmojisGetParams{
		Context: ctx,
	}
}

// NewEmojisGetParamsWithHTTPClient creates a new EmojisGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewEmojisGetParamsWithHTTPClient(client *http.Client) *EmojisGetParams {
	return &EmojisGetParams{
		HTTPClient: client,
	}
}

/*
EmojisGetParams contains all the parameters to send to the API endpoint

	for the emojis get operation.

	Typically these are written to a http.Request.
*/
type EmojisGetParams struct {

	/* Filter.

	     Comma-separated list of filters to apply to results. Recognized filters are:

	`domain:[domain]` -- show emojis from the given domain, eg `?filter=domain:example.org` will show emojis from `example.org` only.
	Instead of giving a specific domain, you can also give either one of the key words `local` or `all` to show either local emojis only (`domain:local`) or show all emojis from all domains (`domain:all`).
	Note: `domain:*` is equivalent to `domain:all` (including local).
	If no domain filter is provided, `domain:all` will be assumed.

	`disabled` -- include emojis that have been disabled.

	`enabled` -- include emojis that are enabled.

	`shortcode:[shortcode]` -- show only emojis with the given shortcode, eg `?filter=shortcode:blob_cat_uwu` will show only emojis with the shortcode `blob_cat_uwu` (case sensitive).

	If neither `disabled` or `enabled` are provided, both disabled and enabled emojis will be shown.

	If no filter query string is provided, the default `domain:all` will be used, which will show all emojis from all domains.

	     Default: "domain:all"
	*/
	Filter *string

	/* Limit.

	   Number of emojis to return. Less than 1, or not set, means unlimited (all emojis).

	   Default: 50
	*/
	Limit *int64

	/* MaxShortcodeDomain.

	     Return only emojis with `[shortcode]@[domain]` *LOWER* (alphabetically) than given `[shortcode]@[domain]`. For example, if `max_shortcode_domain=beep@example.org`, then returned values might include emojis with `[shortcode]@[domain]`s like `car@example.org`, `debian@aaa.com`, `test@` (local emoji), etc.
	Emoji with the given `[shortcode]@[domain]` will not be included in the result set.
	*/
	MaxShortcodeDomain *string

	/* MinShortcodeDomain.

	     Return only emojis with `[shortcode]@[domain]` *HIGHER* (alphabetically) than given `[shortcode]@[domain]`. For example, if `max_shortcode_domain=beep@example.org`, then returned values might include emojis with `[shortcode]@[domain]`s like `arse@test.com`, `0101_binary@hackers.net`, `bee@` (local emoji), etc.
	Emoji with the given `[shortcode]@[domain]` will not be included in the result set.
	*/
	MinShortcodeDomain *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the emojis get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmojisGetParams) WithDefaults() *EmojisGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the emojis get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmojisGetParams) SetDefaults() {
	var (
		filterDefault = string("domain:all")

		limitDefault = int64(50)
	)

	val := EmojisGetParams{
		Filter: &filterDefault,
		Limit:  &limitDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the emojis get params
func (o *EmojisGetParams) WithTimeout(timeout time.Duration) *EmojisGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the emojis get params
func (o *EmojisGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the emojis get params
func (o *EmojisGetParams) WithContext(ctx context.Context) *EmojisGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the emojis get params
func (o *EmojisGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the emojis get params
func (o *EmojisGetParams) WithHTTPClient(client *http.Client) *EmojisGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the emojis get params
func (o *EmojisGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the emojis get params
func (o *EmojisGetParams) WithFilter(filter *string) *EmojisGetParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the emojis get params
func (o *EmojisGetParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the emojis get params
func (o *EmojisGetParams) WithLimit(limit *int64) *EmojisGetParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the emojis get params
func (o *EmojisGetParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMaxShortcodeDomain adds the maxShortcodeDomain to the emojis get params
func (o *EmojisGetParams) WithMaxShortcodeDomain(maxShortcodeDomain *string) *EmojisGetParams {
	o.SetMaxShortcodeDomain(maxShortcodeDomain)
	return o
}

// SetMaxShortcodeDomain adds the maxShortcodeDomain to the emojis get params
func (o *EmojisGetParams) SetMaxShortcodeDomain(maxShortcodeDomain *string) {
	o.MaxShortcodeDomain = maxShortcodeDomain
}

// WithMinShortcodeDomain adds the minShortcodeDomain to the emojis get params
func (o *EmojisGetParams) WithMinShortcodeDomain(minShortcodeDomain *string) *EmojisGetParams {
	o.SetMinShortcodeDomain(minShortcodeDomain)
	return o
}

// SetMinShortcodeDomain adds the minShortcodeDomain to the emojis get params
func (o *EmojisGetParams) SetMinShortcodeDomain(minShortcodeDomain *string) {
	o.MinShortcodeDomain = minShortcodeDomain
}

// WriteToRequest writes these params to a swagger request
func (o *EmojisGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

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

	if o.MaxShortcodeDomain != nil {

		// query param max_shortcode_domain
		var qrMaxShortcodeDomain string

		if o.MaxShortcodeDomain != nil {
			qrMaxShortcodeDomain = *o.MaxShortcodeDomain
		}
		qMaxShortcodeDomain := qrMaxShortcodeDomain
		if qMaxShortcodeDomain != "" {

			if err := r.SetQueryParam("max_shortcode_domain", qMaxShortcodeDomain); err != nil {
				return err
			}
		}
	}

	if o.MinShortcodeDomain != nil {

		// query param min_shortcode_domain
		var qrMinShortcodeDomain string

		if o.MinShortcodeDomain != nil {
			qrMinShortcodeDomain = *o.MinShortcodeDomain
		}
		qMinShortcodeDomain := qrMinShortcodeDomain
		if qMinShortcodeDomain != "" {

			if err := r.SetQueryParam("min_shortcode_domain", qMinShortcodeDomain); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
