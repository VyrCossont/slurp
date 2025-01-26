// Code generated by go-swagger; DO NOT EDIT.

package push

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

// NewPushSubscriptionPutParams creates a new PushSubscriptionPutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPushSubscriptionPutParams() *PushSubscriptionPutParams {
	return &PushSubscriptionPutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPushSubscriptionPutParamsWithTimeout creates a new PushSubscriptionPutParams object
// with the ability to set a timeout on a request.
func NewPushSubscriptionPutParamsWithTimeout(timeout time.Duration) *PushSubscriptionPutParams {
	return &PushSubscriptionPutParams{
		timeout: timeout,
	}
}

// NewPushSubscriptionPutParamsWithContext creates a new PushSubscriptionPutParams object
// with the ability to set a context for a request.
func NewPushSubscriptionPutParamsWithContext(ctx context.Context) *PushSubscriptionPutParams {
	return &PushSubscriptionPutParams{
		Context: ctx,
	}
}

// NewPushSubscriptionPutParamsWithHTTPClient creates a new PushSubscriptionPutParams object
// with the ability to set a custom HTTPClient for a request.
func NewPushSubscriptionPutParamsWithHTTPClient(client *http.Client) *PushSubscriptionPutParams {
	return &PushSubscriptionPutParams{
		HTTPClient: client,
	}
}

/*
PushSubscriptionPutParams contains all the parameters to send to the API endpoint

	for the push subscription put operation.

	Typically these are written to a http.Request.
*/
type PushSubscriptionPutParams struct {

	/* DataAlertsAdminReport.

	   Receive a push notification when a new report has been filed?
	*/
	DataAlertsAdminReport *bool

	/* DataAlertsAdminSignUp.

	   Receive a push notification when a new user has signed up?
	*/
	DataAlertsAdminSignUp *bool

	/* DataAlertsFavourite.

	   Receive a push notification when a status you created has been favourited by someone else?
	*/
	DataAlertsFavourite *bool

	/* DataAlertsFollow.

	   Receive a push notification when someone has followed you?
	*/
	DataAlertsFollow *bool

	/* DataAlertsFollowRequest.

	   Receive a push notification when someone has requested to follow you?
	*/
	DataAlertsFollowRequest *bool

	/* DataAlertsMention.

	   Receive a push notification when someone else has mentioned you in a status?
	*/
	DataAlertsMention *bool

	/* DataAlertsPendingFavourite.

	   Receive a push notification when a fave is pending?
	*/
	DataAlertsPendingFavourite *bool

	/* DataAlertsPendingReblog.

	   Receive a push notification when a boost is pending?
	*/
	DataAlertsPendingReblog *bool

	/* DataAlertsPendingReply.

	   Receive a push notification when a reply is pending?
	*/
	DataAlertsPendingReply *bool

	/* DataAlertsPoll.

	   Receive a push notification when a poll you voted in or created has ended?
	*/
	DataAlertsPoll *bool

	/* DataAlertsReblog.

	   Receive a push notification when a status you created has been boosted by someone else?
	*/
	DataAlertsReblog *bool

	/* DataAlertsStatus.

	   Receive a push notification when a subscribed account posts a status?
	*/
	DataAlertsStatus *bool

	/* DataAlertsUpdate.

	   Receive a push notification when a status you interacted with has been edited?
	*/
	DataAlertsUpdate *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the push subscription put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PushSubscriptionPutParams) WithDefaults() *PushSubscriptionPutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the push subscription put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PushSubscriptionPutParams) SetDefaults() {
	var (
		dataAlertsAdminReportDefault = bool(false)

		dataAlertsAdminSignUpDefault = bool(false)

		dataAlertsFavouriteDefault = bool(false)

		dataAlertsFollowDefault = bool(false)

		dataAlertsFollowRequestDefault = bool(false)

		dataAlertsMentionDefault = bool(false)

		dataAlertsPendingFavouriteDefault = bool(false)

		dataAlertsPendingReblogDefault = bool(false)

		dataAlertsPendingReplyDefault = bool(false)

		dataAlertsPollDefault = bool(false)

		dataAlertsReblogDefault = bool(false)

		dataAlertsStatusDefault = bool(false)

		dataAlertsUpdateDefault = bool(false)
	)

	val := PushSubscriptionPutParams{
		DataAlertsAdminReport:      &dataAlertsAdminReportDefault,
		DataAlertsAdminSignUp:      &dataAlertsAdminSignUpDefault,
		DataAlertsFavourite:        &dataAlertsFavouriteDefault,
		DataAlertsFollow:           &dataAlertsFollowDefault,
		DataAlertsFollowRequest:    &dataAlertsFollowRequestDefault,
		DataAlertsMention:          &dataAlertsMentionDefault,
		DataAlertsPendingFavourite: &dataAlertsPendingFavouriteDefault,
		DataAlertsPendingReblog:    &dataAlertsPendingReblogDefault,
		DataAlertsPendingReply:     &dataAlertsPendingReplyDefault,
		DataAlertsPoll:             &dataAlertsPollDefault,
		DataAlertsReblog:           &dataAlertsReblogDefault,
		DataAlertsStatus:           &dataAlertsStatusDefault,
		DataAlertsUpdate:           &dataAlertsUpdateDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the push subscription put params
func (o *PushSubscriptionPutParams) WithTimeout(timeout time.Duration) *PushSubscriptionPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the push subscription put params
func (o *PushSubscriptionPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the push subscription put params
func (o *PushSubscriptionPutParams) WithContext(ctx context.Context) *PushSubscriptionPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the push subscription put params
func (o *PushSubscriptionPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the push subscription put params
func (o *PushSubscriptionPutParams) WithHTTPClient(client *http.Client) *PushSubscriptionPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the push subscription put params
func (o *PushSubscriptionPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDataAlertsAdminReport adds the dataAlertsAdminReport to the push subscription put params
func (o *PushSubscriptionPutParams) WithDataAlertsAdminReport(dataAlertsAdminReport *bool) *PushSubscriptionPutParams {
	o.SetDataAlertsAdminReport(dataAlertsAdminReport)
	return o
}

// SetDataAlertsAdminReport adds the dataAlertsAdminReport to the push subscription put params
func (o *PushSubscriptionPutParams) SetDataAlertsAdminReport(dataAlertsAdminReport *bool) {
	o.DataAlertsAdminReport = dataAlertsAdminReport
}

// WithDataAlertsAdminSignUp adds the dataAlertsAdminSignUp to the push subscription put params
func (o *PushSubscriptionPutParams) WithDataAlertsAdminSignUp(dataAlertsAdminSignUp *bool) *PushSubscriptionPutParams {
	o.SetDataAlertsAdminSignUp(dataAlertsAdminSignUp)
	return o
}

// SetDataAlertsAdminSignUp adds the dataAlertsAdminSignUp to the push subscription put params
func (o *PushSubscriptionPutParams) SetDataAlertsAdminSignUp(dataAlertsAdminSignUp *bool) {
	o.DataAlertsAdminSignUp = dataAlertsAdminSignUp
}

// WithDataAlertsFavourite adds the dataAlertsFavourite to the push subscription put params
func (o *PushSubscriptionPutParams) WithDataAlertsFavourite(dataAlertsFavourite *bool) *PushSubscriptionPutParams {
	o.SetDataAlertsFavourite(dataAlertsFavourite)
	return o
}

// SetDataAlertsFavourite adds the dataAlertsFavourite to the push subscription put params
func (o *PushSubscriptionPutParams) SetDataAlertsFavourite(dataAlertsFavourite *bool) {
	o.DataAlertsFavourite = dataAlertsFavourite
}

// WithDataAlertsFollow adds the dataAlertsFollow to the push subscription put params
func (o *PushSubscriptionPutParams) WithDataAlertsFollow(dataAlertsFollow *bool) *PushSubscriptionPutParams {
	o.SetDataAlertsFollow(dataAlertsFollow)
	return o
}

// SetDataAlertsFollow adds the dataAlertsFollow to the push subscription put params
func (o *PushSubscriptionPutParams) SetDataAlertsFollow(dataAlertsFollow *bool) {
	o.DataAlertsFollow = dataAlertsFollow
}

// WithDataAlertsFollowRequest adds the dataAlertsFollowRequest to the push subscription put params
func (o *PushSubscriptionPutParams) WithDataAlertsFollowRequest(dataAlertsFollowRequest *bool) *PushSubscriptionPutParams {
	o.SetDataAlertsFollowRequest(dataAlertsFollowRequest)
	return o
}

// SetDataAlertsFollowRequest adds the dataAlertsFollowRequest to the push subscription put params
func (o *PushSubscriptionPutParams) SetDataAlertsFollowRequest(dataAlertsFollowRequest *bool) {
	o.DataAlertsFollowRequest = dataAlertsFollowRequest
}

// WithDataAlertsMention adds the dataAlertsMention to the push subscription put params
func (o *PushSubscriptionPutParams) WithDataAlertsMention(dataAlertsMention *bool) *PushSubscriptionPutParams {
	o.SetDataAlertsMention(dataAlertsMention)
	return o
}

// SetDataAlertsMention adds the dataAlertsMention to the push subscription put params
func (o *PushSubscriptionPutParams) SetDataAlertsMention(dataAlertsMention *bool) {
	o.DataAlertsMention = dataAlertsMention
}

// WithDataAlertsPendingFavourite adds the dataAlertsPendingFavourite to the push subscription put params
func (o *PushSubscriptionPutParams) WithDataAlertsPendingFavourite(dataAlertsPendingFavourite *bool) *PushSubscriptionPutParams {
	o.SetDataAlertsPendingFavourite(dataAlertsPendingFavourite)
	return o
}

// SetDataAlertsPendingFavourite adds the dataAlertsPendingFavourite to the push subscription put params
func (o *PushSubscriptionPutParams) SetDataAlertsPendingFavourite(dataAlertsPendingFavourite *bool) {
	o.DataAlertsPendingFavourite = dataAlertsPendingFavourite
}

// WithDataAlertsPendingReblog adds the dataAlertsPendingReblog to the push subscription put params
func (o *PushSubscriptionPutParams) WithDataAlertsPendingReblog(dataAlertsPendingReblog *bool) *PushSubscriptionPutParams {
	o.SetDataAlertsPendingReblog(dataAlertsPendingReblog)
	return o
}

// SetDataAlertsPendingReblog adds the dataAlertsPendingReblog to the push subscription put params
func (o *PushSubscriptionPutParams) SetDataAlertsPendingReblog(dataAlertsPendingReblog *bool) {
	o.DataAlertsPendingReblog = dataAlertsPendingReblog
}

// WithDataAlertsPendingReply adds the dataAlertsPendingReply to the push subscription put params
func (o *PushSubscriptionPutParams) WithDataAlertsPendingReply(dataAlertsPendingReply *bool) *PushSubscriptionPutParams {
	o.SetDataAlertsPendingReply(dataAlertsPendingReply)
	return o
}

// SetDataAlertsPendingReply adds the dataAlertsPendingReply to the push subscription put params
func (o *PushSubscriptionPutParams) SetDataAlertsPendingReply(dataAlertsPendingReply *bool) {
	o.DataAlertsPendingReply = dataAlertsPendingReply
}

// WithDataAlertsPoll adds the dataAlertsPoll to the push subscription put params
func (o *PushSubscriptionPutParams) WithDataAlertsPoll(dataAlertsPoll *bool) *PushSubscriptionPutParams {
	o.SetDataAlertsPoll(dataAlertsPoll)
	return o
}

// SetDataAlertsPoll adds the dataAlertsPoll to the push subscription put params
func (o *PushSubscriptionPutParams) SetDataAlertsPoll(dataAlertsPoll *bool) {
	o.DataAlertsPoll = dataAlertsPoll
}

// WithDataAlertsReblog adds the dataAlertsReblog to the push subscription put params
func (o *PushSubscriptionPutParams) WithDataAlertsReblog(dataAlertsReblog *bool) *PushSubscriptionPutParams {
	o.SetDataAlertsReblog(dataAlertsReblog)
	return o
}

// SetDataAlertsReblog adds the dataAlertsReblog to the push subscription put params
func (o *PushSubscriptionPutParams) SetDataAlertsReblog(dataAlertsReblog *bool) {
	o.DataAlertsReblog = dataAlertsReblog
}

// WithDataAlertsStatus adds the dataAlertsStatus to the push subscription put params
func (o *PushSubscriptionPutParams) WithDataAlertsStatus(dataAlertsStatus *bool) *PushSubscriptionPutParams {
	o.SetDataAlertsStatus(dataAlertsStatus)
	return o
}

// SetDataAlertsStatus adds the dataAlertsStatus to the push subscription put params
func (o *PushSubscriptionPutParams) SetDataAlertsStatus(dataAlertsStatus *bool) {
	o.DataAlertsStatus = dataAlertsStatus
}

// WithDataAlertsUpdate adds the dataAlertsUpdate to the push subscription put params
func (o *PushSubscriptionPutParams) WithDataAlertsUpdate(dataAlertsUpdate *bool) *PushSubscriptionPutParams {
	o.SetDataAlertsUpdate(dataAlertsUpdate)
	return o
}

// SetDataAlertsUpdate adds the dataAlertsUpdate to the push subscription put params
func (o *PushSubscriptionPutParams) SetDataAlertsUpdate(dataAlertsUpdate *bool) {
	o.DataAlertsUpdate = dataAlertsUpdate
}

// WriteToRequest writes these params to a swagger request
func (o *PushSubscriptionPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DataAlertsAdminReport != nil {

		// form param data[alerts][admin.report]
		var frDataAlertsAdminReport bool
		if o.DataAlertsAdminReport != nil {
			frDataAlertsAdminReport = *o.DataAlertsAdminReport
		}
		fDataAlertsAdminReport := swag.FormatBool(frDataAlertsAdminReport)
		if fDataAlertsAdminReport != "" {
			if err := r.SetFormParam("data[alerts][admin.report]", fDataAlertsAdminReport); err != nil {
				return err
			}
		}
	}

	if o.DataAlertsAdminSignUp != nil {

		// form param data[alerts][admin.sign_up]
		var frDataAlertsAdminSignUp bool
		if o.DataAlertsAdminSignUp != nil {
			frDataAlertsAdminSignUp = *o.DataAlertsAdminSignUp
		}
		fDataAlertsAdminSignUp := swag.FormatBool(frDataAlertsAdminSignUp)
		if fDataAlertsAdminSignUp != "" {
			if err := r.SetFormParam("data[alerts][admin.sign_up]", fDataAlertsAdminSignUp); err != nil {
				return err
			}
		}
	}

	if o.DataAlertsFavourite != nil {

		// form param data[alerts][favourite]
		var frDataAlertsFavourite bool
		if o.DataAlertsFavourite != nil {
			frDataAlertsFavourite = *o.DataAlertsFavourite
		}
		fDataAlertsFavourite := swag.FormatBool(frDataAlertsFavourite)
		if fDataAlertsFavourite != "" {
			if err := r.SetFormParam("data[alerts][favourite]", fDataAlertsFavourite); err != nil {
				return err
			}
		}
	}

	if o.DataAlertsFollow != nil {

		// form param data[alerts][follow]
		var frDataAlertsFollow bool
		if o.DataAlertsFollow != nil {
			frDataAlertsFollow = *o.DataAlertsFollow
		}
		fDataAlertsFollow := swag.FormatBool(frDataAlertsFollow)
		if fDataAlertsFollow != "" {
			if err := r.SetFormParam("data[alerts][follow]", fDataAlertsFollow); err != nil {
				return err
			}
		}
	}

	if o.DataAlertsFollowRequest != nil {

		// form param data[alerts][follow_request]
		var frDataAlertsFollowRequest bool
		if o.DataAlertsFollowRequest != nil {
			frDataAlertsFollowRequest = *o.DataAlertsFollowRequest
		}
		fDataAlertsFollowRequest := swag.FormatBool(frDataAlertsFollowRequest)
		if fDataAlertsFollowRequest != "" {
			if err := r.SetFormParam("data[alerts][follow_request]", fDataAlertsFollowRequest); err != nil {
				return err
			}
		}
	}

	if o.DataAlertsMention != nil {

		// form param data[alerts][mention]
		var frDataAlertsMention bool
		if o.DataAlertsMention != nil {
			frDataAlertsMention = *o.DataAlertsMention
		}
		fDataAlertsMention := swag.FormatBool(frDataAlertsMention)
		if fDataAlertsMention != "" {
			if err := r.SetFormParam("data[alerts][mention]", fDataAlertsMention); err != nil {
				return err
			}
		}
	}

	if o.DataAlertsPendingFavourite != nil {

		// form param data[alerts][pending.favourite]
		var frDataAlertsPendingFavourite bool
		if o.DataAlertsPendingFavourite != nil {
			frDataAlertsPendingFavourite = *o.DataAlertsPendingFavourite
		}
		fDataAlertsPendingFavourite := swag.FormatBool(frDataAlertsPendingFavourite)
		if fDataAlertsPendingFavourite != "" {
			if err := r.SetFormParam("data[alerts][pending.favourite]", fDataAlertsPendingFavourite); err != nil {
				return err
			}
		}
	}

	if o.DataAlertsPendingReblog != nil {

		// form param data[alerts][pending.reblog]
		var frDataAlertsPendingReblog bool
		if o.DataAlertsPendingReblog != nil {
			frDataAlertsPendingReblog = *o.DataAlertsPendingReblog
		}
		fDataAlertsPendingReblog := swag.FormatBool(frDataAlertsPendingReblog)
		if fDataAlertsPendingReblog != "" {
			if err := r.SetFormParam("data[alerts][pending.reblog]", fDataAlertsPendingReblog); err != nil {
				return err
			}
		}
	}

	if o.DataAlertsPendingReply != nil {

		// form param data[alerts][pending.reply]
		var frDataAlertsPendingReply bool
		if o.DataAlertsPendingReply != nil {
			frDataAlertsPendingReply = *o.DataAlertsPendingReply
		}
		fDataAlertsPendingReply := swag.FormatBool(frDataAlertsPendingReply)
		if fDataAlertsPendingReply != "" {
			if err := r.SetFormParam("data[alerts][pending.reply]", fDataAlertsPendingReply); err != nil {
				return err
			}
		}
	}

	if o.DataAlertsPoll != nil {

		// form param data[alerts][poll]
		var frDataAlertsPoll bool
		if o.DataAlertsPoll != nil {
			frDataAlertsPoll = *o.DataAlertsPoll
		}
		fDataAlertsPoll := swag.FormatBool(frDataAlertsPoll)
		if fDataAlertsPoll != "" {
			if err := r.SetFormParam("data[alerts][poll]", fDataAlertsPoll); err != nil {
				return err
			}
		}
	}

	if o.DataAlertsReblog != nil {

		// form param data[alerts][reblog]
		var frDataAlertsReblog bool
		if o.DataAlertsReblog != nil {
			frDataAlertsReblog = *o.DataAlertsReblog
		}
		fDataAlertsReblog := swag.FormatBool(frDataAlertsReblog)
		if fDataAlertsReblog != "" {
			if err := r.SetFormParam("data[alerts][reblog]", fDataAlertsReblog); err != nil {
				return err
			}
		}
	}

	if o.DataAlertsStatus != nil {

		// form param data[alerts][status]
		var frDataAlertsStatus bool
		if o.DataAlertsStatus != nil {
			frDataAlertsStatus = *o.DataAlertsStatus
		}
		fDataAlertsStatus := swag.FormatBool(frDataAlertsStatus)
		if fDataAlertsStatus != "" {
			if err := r.SetFormParam("data[alerts][status]", fDataAlertsStatus); err != nil {
				return err
			}
		}
	}

	if o.DataAlertsUpdate != nil {

		// form param data[alerts][update]
		var frDataAlertsUpdate bool
		if o.DataAlertsUpdate != nil {
			frDataAlertsUpdate = *o.DataAlertsUpdate
		}
		fDataAlertsUpdate := swag.FormatBool(frDataAlertsUpdate)
		if fDataAlertsUpdate != "" {
			if err := r.SetFormParam("data[alerts][update]", fDataAlertsUpdate); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}