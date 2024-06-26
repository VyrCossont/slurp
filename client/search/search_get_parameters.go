// Code generated by go-swagger; DO NOT EDIT.

package search

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

// NewSearchGetParams creates a new SearchGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchGetParams() *SearchGetParams {
	return &SearchGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchGetParamsWithTimeout creates a new SearchGetParams object
// with the ability to set a timeout on a request.
func NewSearchGetParamsWithTimeout(timeout time.Duration) *SearchGetParams {
	return &SearchGetParams{
		timeout: timeout,
	}
}

// NewSearchGetParamsWithContext creates a new SearchGetParams object
// with the ability to set a context for a request.
func NewSearchGetParamsWithContext(ctx context.Context) *SearchGetParams {
	return &SearchGetParams{
		Context: ctx,
	}
}

// NewSearchGetParamsWithHTTPClient creates a new SearchGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchGetParamsWithHTTPClient(client *http.Client) *SearchGetParams {
	return &SearchGetParams{
		HTTPClient: client,
	}
}

/*
SearchGetParams contains all the parameters to send to the API endpoint

	for the search get operation.

	Typically these are written to a http.Request.
*/
type SearchGetParams struct {

	/* AccountID.

	   Restrict results to statuses created by the specified account.
	*/
	AccountID *string

	/* APIVersion.

	   Version of the API to use. Must be either `v1` or `v2`. If v1 is used, Hashtag results will be a slice of strings. If v2 is used, Hashtag results will be a slice of apimodel tags.
	*/
	APIVersion string

	/* ExcludeUnreviewed.

	   If searching for hashtags, exclude those not yet approved by instance admin. Currently this parameter is unused.
	*/
	ExcludeUnreviewed *bool

	/* Following.

	   If search type includes accounts, and search query is an arbitrary string, show only accounts that the requesting account follows. If this is set to `true`, then the GoToSocial instance will enhance the search by also searching within account notes, not just in usernames and display names.
	*/
	Following *bool

	/* Limit.

	   Number of each type of item to return.

	   Default: 20
	*/
	Limit *int64

	/* MaxID.

	   Return only items *OLDER* than the given max ID. The item with the specified ID will not be included in the response. Currently only used if 'type' is set to a specific type.
	*/
	MaxID *string

	/* MinID.

	   Return only items *immediately newer* than the given min ID. The item with the specified ID will not be included in the response. Currently only used if 'type' is set to a specific type.
	*/
	MinID *string

	/* Offset.

	   Page number of results to return (starts at 0). This parameter is currently not used, page by selecting a specific query type and using maxID and minID instead.
	*/
	Offset *int64

	/* Q.

	     Query string to search for. This can be in the following forms:
	- `@[username]` -- search for an account with the given username on any domain. Can return multiple results.
	- @[username]@[domain]` -- search for a remote account with exact username and domain. Will only ever return 1 result at most.
	- `https://example.org/some/arbitrary/url` -- search for an account OR a status with the given URL. Will only ever return 1 result at most.
	- `#[hashtag_name]` -- search for a hashtag with the given hashtag name, or starting with the given hashtag name. Case insensitive. Can return multiple results.
	- any arbitrary string -- search for accounts or statuses containing the given string. Can return multiple results.

	Arbitrary string queries may include the following operators:
	- `from:localuser`, `from:remoteuser@instance.tld`: restrict results to statuses created by the specified account.
	*/
	Q string

	/* Resolve.

	   If searching query is for `@[username]@[domain]`, or a URL, allow the GoToSocial instance to resolve the search by making calls to remote instances (webfinger, ActivityPub, etc).
	*/
	Resolve *bool

	/* Type.

	     Type of item to return. One of:
	- `` -- empty string; return any/all results.
	- `accounts` -- return only account(s).
	- `statuses` -- return only status(es).
	- `hashtags` -- return only hashtag(s).
	If `type` is specified, paging can be performed using max_id and min_id parameters.
	If `type` is not specified, see the `offset` parameter for paging.
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchGetParams) WithDefaults() *SearchGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchGetParams) SetDefaults() {
	var (
		excludeUnreviewedDefault = bool(false)

		followingDefault = bool(false)

		limitDefault = int64(20)

		offsetDefault = int64(0)

		resolveDefault = bool(false)
	)

	val := SearchGetParams{
		ExcludeUnreviewed: &excludeUnreviewedDefault,
		Following:         &followingDefault,
		Limit:             &limitDefault,
		Offset:            &offsetDefault,
		Resolve:           &resolveDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the search get params
func (o *SearchGetParams) WithTimeout(timeout time.Duration) *SearchGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search get params
func (o *SearchGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search get params
func (o *SearchGetParams) WithContext(ctx context.Context) *SearchGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search get params
func (o *SearchGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search get params
func (o *SearchGetParams) WithHTTPClient(client *http.Client) *SearchGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search get params
func (o *SearchGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the search get params
func (o *SearchGetParams) WithAccountID(accountID *string) *SearchGetParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the search get params
func (o *SearchGetParams) SetAccountID(accountID *string) {
	o.AccountID = accountID
}

// WithAPIVersion adds the aPIVersion to the search get params
func (o *SearchGetParams) WithAPIVersion(aPIVersion string) *SearchGetParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the search get params
func (o *SearchGetParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithExcludeUnreviewed adds the excludeUnreviewed to the search get params
func (o *SearchGetParams) WithExcludeUnreviewed(excludeUnreviewed *bool) *SearchGetParams {
	o.SetExcludeUnreviewed(excludeUnreviewed)
	return o
}

// SetExcludeUnreviewed adds the excludeUnreviewed to the search get params
func (o *SearchGetParams) SetExcludeUnreviewed(excludeUnreviewed *bool) {
	o.ExcludeUnreviewed = excludeUnreviewed
}

// WithFollowing adds the following to the search get params
func (o *SearchGetParams) WithFollowing(following *bool) *SearchGetParams {
	o.SetFollowing(following)
	return o
}

// SetFollowing adds the following to the search get params
func (o *SearchGetParams) SetFollowing(following *bool) {
	o.Following = following
}

// WithLimit adds the limit to the search get params
func (o *SearchGetParams) WithLimit(limit *int64) *SearchGetParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search get params
func (o *SearchGetParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMaxID adds the maxID to the search get params
func (o *SearchGetParams) WithMaxID(maxID *string) *SearchGetParams {
	o.SetMaxID(maxID)
	return o
}

// SetMaxID adds the maxId to the search get params
func (o *SearchGetParams) SetMaxID(maxID *string) {
	o.MaxID = maxID
}

// WithMinID adds the minID to the search get params
func (o *SearchGetParams) WithMinID(minID *string) *SearchGetParams {
	o.SetMinID(minID)
	return o
}

// SetMinID adds the minId to the search get params
func (o *SearchGetParams) SetMinID(minID *string) {
	o.MinID = minID
}

// WithOffset adds the offset to the search get params
func (o *SearchGetParams) WithOffset(offset *int64) *SearchGetParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search get params
func (o *SearchGetParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the search get params
func (o *SearchGetParams) WithQ(q string) *SearchGetParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the search get params
func (o *SearchGetParams) SetQ(q string) {
	o.Q = q
}

// WithResolve adds the resolve to the search get params
func (o *SearchGetParams) WithResolve(resolve *bool) *SearchGetParams {
	o.SetResolve(resolve)
	return o
}

// SetResolve adds the resolve to the search get params
func (o *SearchGetParams) SetResolve(resolve *bool) {
	o.Resolve = resolve
}

// WithType adds the typeVar to the search get params
func (o *SearchGetParams) WithType(typeVar *string) *SearchGetParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the search get params
func (o *SearchGetParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *SearchGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountID != nil {

		// query param account_id
		var qrAccountID string

		if o.AccountID != nil {
			qrAccountID = *o.AccountID
		}
		qAccountID := qrAccountID
		if qAccountID != "" {

			if err := r.SetQueryParam("account_id", qAccountID); err != nil {
				return err
			}
		}
	}

	// path param api_version
	if err := r.SetPathParam("api_version", o.APIVersion); err != nil {
		return err
	}

	if o.ExcludeUnreviewed != nil {

		// query param exclude_unreviewed
		var qrExcludeUnreviewed bool

		if o.ExcludeUnreviewed != nil {
			qrExcludeUnreviewed = *o.ExcludeUnreviewed
		}
		qExcludeUnreviewed := swag.FormatBool(qrExcludeUnreviewed)
		if qExcludeUnreviewed != "" {

			if err := r.SetQueryParam("exclude_unreviewed", qExcludeUnreviewed); err != nil {
				return err
			}
		}
	}

	if o.Following != nil {

		// query param following
		var qrFollowing bool

		if o.Following != nil {
			qrFollowing = *o.Following
		}
		qFollowing := swag.FormatBool(qrFollowing)
		if qFollowing != "" {

			if err := r.SetQueryParam("following", qFollowing); err != nil {
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

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	// query param q
	qrQ := o.Q
	qQ := qrQ
	if qQ != "" {

		if err := r.SetQueryParam("q", qQ); err != nil {
			return err
		}
	}

	if o.Resolve != nil {

		// query param resolve
		var qrResolve bool

		if o.Resolve != nil {
			qrResolve = *o.Resolve
		}
		qResolve := swag.FormatBool(qrResolve)
		if qResolve != "" {

			if err := r.SetQueryParam("resolve", qResolve); err != nil {
				return err
			}
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
