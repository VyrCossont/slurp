// Code generated by go-swagger; DO NOT EDIT.

package streaming

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new streaming API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for streaming API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	StreamGet(params *StreamGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) error

	SetTransport(transport runtime.ClientTransport)
}

/*
	StreamGet initiates a websocket connection for live streaming of statuses and notifications

	The scheme used should *always* be `wss`. The streaming basepath can be viewed at `/api/v1/instance`.

On a successful connection, a code `101` will be returned, which indicates that the connection is being upgraded to a secure websocket connection.

As long as the connection is open, various message types will be streamed into it.

GoToSocial will ping the connection every 30 seconds to check whether the client is still receiving.

If the ping fails, or something else goes wrong during transmission, then the connection will be dropped, and the client will be expected to start it again.
*/
func (a *Client) StreamGet(params *StreamGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStreamGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "streamGet",
		Method:             "GET",
		PathPattern:        "/api/v1/streaming",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &StreamGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
