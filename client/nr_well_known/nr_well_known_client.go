// Code generated by go-swagger; DO NOT EDIT.

package nr_well_known

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new nr well known API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for nr well known API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	HostMetaGet(params *HostMetaGetParams, opts ...ClientOption) (*HostMetaGetOK, error)

	NodeInfoWellKnownGet(params *NodeInfoWellKnownGetParams, opts ...ClientOption) (*NodeInfoWellKnownGetOK, error)

	WebfingerGet(params *WebfingerGetParams, opts ...ClientOption) (*WebfingerGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
HostMetaGet returns a compliant hostmeta response to web host metadata queries

See: https://www.rfc-editor.org/rfc/rfc6415.html
*/
func (a *Client) HostMetaGet(params *HostMetaGetParams, opts ...ClientOption) (*HostMetaGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHostMetaGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "hostMetaGet",
		Method:             "GET",
		PathPattern:        "/.well-known/host-meta",
		ProducesMediaTypes: []string{"application/xrd+xml\""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &HostMetaGetReader{formats: a.formats},
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
	success, ok := result.(*HostMetaGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for hostMetaGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	NodeInfoWellKnownGet returns a well known response which redirects callers to nodeinfo 2 0

	eg. `{"links":[{"rel":"http://nodeinfo.diaspora.software/ns/schema/2.0","href":"http://example.org/nodeinfo/2.0"}]}`

See: https://nodeinfo.diaspora.software/protocol.html
*/
func (a *Client) NodeInfoWellKnownGet(params *NodeInfoWellKnownGetParams, opts ...ClientOption) (*NodeInfoWellKnownGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNodeInfoWellKnownGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "nodeInfoWellKnownGet",
		Method:             "GET",
		PathPattern:        "/.well-known/nodeinfo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &NodeInfoWellKnownGetReader{formats: a.formats},
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
	success, ok := result.(*NodeInfoWellKnownGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for nodeInfoWellKnownGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	WebfingerGet handles webfinger account lookup requests

	For example, a GET to `https://goblin.technology/.well-known/webfinger?resource=acct:tobi@goblin.technology` would return:

```

{"subject":"acct:tobi@goblin.technology","aliases":["https://goblin.technology/users/tobi","https://goblin.technology/@tobi"],"links":[{"rel":"http://webfinger.net/rel/profile-page","type":"text/html","href":"https://goblin.technology/@tobi"},{"rel":"self","type":"application/activity+json","href":"https://goblin.technology/users/tobi"}]}

```

See: https://webfinger.net/
*/
func (a *Client) WebfingerGet(params *WebfingerGetParams, opts ...ClientOption) (*WebfingerGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWebfingerGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "webfingerGet",
		Method:             "GET",
		PathPattern:        "/.well-known/webfinger",
		ProducesMediaTypes: []string{"application/jrd+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &WebfingerGetReader{formats: a.formats},
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
	success, ok := result.(*WebfingerGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for webfingerGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
