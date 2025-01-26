// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebPushSubscriptionAlerts WebPushSubscriptionAlerts represents the specific events that this Web Push subscription will receive.
//
// swagger:model WebPushSubscriptionAlerts
type WebPushSubscriptionAlerts struct {

	// Receive a push notification when a new report has been filed?
	AdminReport bool `json:"admin.report,omitempty"`

	// Receive a push notification when a new user has signed up?
	AdminSignup bool `json:"admin.sign_up,omitempty"`

	// Receive a push notification when a status you created has been favourited by someone else?
	Favourite bool `json:"favourite,omitempty"`

	// Receive a push notification when someone has followed you?
	Follow bool `json:"follow,omitempty"`

	// Receive a push notification when someone has requested to follow you?
	FollowRequest bool `json:"follow_request,omitempty"`

	// Receive a push notification when someone else has mentioned you in a status?
	Mention bool `json:"mention,omitempty"`

	// Receive a push notification when a fave is pending?
	PendingFavourite bool `json:"pending.favourite,omitempty"`

	// Receive a push notification when a boost is pending?
	PendingReblog bool `json:"pending.reblog,omitempty"`

	// Receive a push notification when a reply is pending?
	PendingReply bool `json:"pending.reply,omitempty"`

	// Receive a push notification when a poll you voted in or created has ended?
	Poll bool `json:"poll,omitempty"`

	// Receive a push notification when a status you created has been boosted by someone else?
	Reblog bool `json:"reblog,omitempty"`

	// Receive a push notification when a subscribed account posts a status?
	Status bool `json:"status,omitempty"`

	// Receive a push notification when a status you interacted with has been edited?
	Update bool `json:"update,omitempty"`
}

// Validate validates this web push subscription alerts
func (m *WebPushSubscriptionAlerts) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this web push subscription alerts based on context it is used
func (m *WebPushSubscriptionAlerts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebPushSubscriptionAlerts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebPushSubscriptionAlerts) UnmarshalBinary(b []byte) error {
	var res WebPushSubscriptionAlerts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
