// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetUser(params *GetUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserOK, error)

	UserEmailChange(params *UserEmailChangeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserEmailChangeAccepted, error)

	UserPasswordChange(params *UserPasswordChangeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserPasswordChangeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetUser gets your own user model
*/
func (a *Client) GetUser(params *GetUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUser",
		Method:             "GET",
		PathPattern:        "/api/v1/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UserEmailChange requests changing the email address of authenticated user
*/
func (a *Client) UserEmailChange(params *UserEmailChangeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserEmailChangeAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserEmailChangeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "userEmailChange",
		Method:             "POST",
		PathPattern:        "/api/v1/user/email_change",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded", "application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserEmailChangeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UserEmailChangeAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for userEmailChange: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	UserPasswordChange changes the password of authenticated user

	The parameters can also be given in the body of the request, as JSON, if the content-type is set to 'application/json'.

The parameters can also be given in the body of the request, as XML, if the content-type is set to 'application/xml'.
*/
func (a *Client) UserPasswordChange(params *UserPasswordChangeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserPasswordChangeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserPasswordChangeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "userPasswordChange",
		Method:             "POST",
		PathPattern:        "/api/v1/user/password_change",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded", "application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserPasswordChangeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UserPasswordChangeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for userPasswordChange: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
