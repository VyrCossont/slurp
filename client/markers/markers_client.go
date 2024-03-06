// Code generated by go-swagger; DO NOT EDIT.

package markers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new markers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for markers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	MarkersGet(params *MarkersGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MarkersGetOK, error)

	MarkersPost(params *MarkersPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MarkersPostOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
MarkersGet Get timeline markers by name
*/
func (a *Client) MarkersGet(params *MarkersGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MarkersGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMarkersGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "markersGet",
		Method:             "GET",
		PathPattern:        "/api/v1/markers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &MarkersGetReader{formats: a.formats},
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
	success, ok := result.(*MarkersGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for markersGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
MarkersPost Update timeline markers by name
*/
func (a *Client) MarkersPost(params *MarkersPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MarkersPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMarkersPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "markersPost",
		Method:             "POST",
		PathPattern:        "/api/v1/markers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &MarkersPostReader{formats: a.formats},
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
	success, ok := result.(*MarkersPostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for markersPost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
