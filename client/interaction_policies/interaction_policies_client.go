// Code generated by go-swagger; DO NOT EDIT.

package interaction_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new interaction policies API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new interaction policies API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new interaction policies API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for interaction policies API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// This client is generated with a few options you might find useful for your swagger spec.
//
// Feel free to add you own set of options.

// WithContentType allows the client to force the Content-Type header
// to negotiate a specific Consumer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithContentType(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ConsumesMediaTypes = []string{mime}
	}
}

// WithContentTypeApplicationJSON sets the Content-Type header to "application/json".
func WithContentTypeApplicationJSON(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/json"}
}

// WithContentTypeApplicationxWwwFormUrlencoded sets the Content-Type header to "application/x-www-form-urlencoded".
func WithContentTypeApplicationxWwwFormUrlencoded(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/x-www-form-urlencoded"}
}

// WithContentTypeMultipartFormData sets the Content-Type header to "multipart/form-data".
func WithContentTypeMultipartFormData(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"multipart/form-data"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	PoliciesDefaultsGet(params *PoliciesDefaultsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PoliciesDefaultsGetOK, error)

	PoliciesDefaultsUpdate(params *PoliciesDefaultsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PoliciesDefaultsUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
PoliciesDefaultsGet gets default interaction policies for new statuses created by you
*/
func (a *Client) PoliciesDefaultsGet(params *PoliciesDefaultsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PoliciesDefaultsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPoliciesDefaultsGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "policiesDefaultsGet",
		Method:             "GET",
		PathPattern:        "/api/v1/interaction_policies/defaults",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PoliciesDefaultsGetReader{formats: a.formats},
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
	success, ok := result.(*PoliciesDefaultsGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for policiesDefaultsGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	PoliciesDefaultsUpdate updates default interaction policies per visibility level for new statuses created by you

	If submitting using form data, use the following pattern:

`VISIBILITY[INTERACTION_TYPE][CONDITION][INDEX]=Value`

For example: `public[can_reply][always][0]=author`

Using `curl` this might look something like:

`curl -F 'public[can_reply][always][0]=author' -F 'public[can_reply][always][1]=followers'`

The JSON equivalent would be:

`curl -H 'Content-Type: application/json' -d '{"public":{"can_reply":{"always":["author","followers"]}}}'`

Any visibility level left unspecified in the request body will be returned to the default.

Ie., in the example above, "public" would be updated, but "unlisted", "private", and "direct" would be reset to defaults.

The server will perform some normalization on submitted policies so that you can't submit totally invalid policies.
*/
func (a *Client) PoliciesDefaultsUpdate(params *PoliciesDefaultsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PoliciesDefaultsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPoliciesDefaultsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "policiesDefaultsUpdate",
		Method:             "PATCH",
		PathPattern:        "/api/v1/interaction_policies/defaults",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data", "application/x-www-form-urlencoded", "application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PoliciesDefaultsUpdateReader{formats: a.formats},
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
	success, ok := result.(*PoliciesDefaultsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for policiesDefaultsUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
