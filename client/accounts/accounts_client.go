// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new accounts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for accounts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AccountAlias(params *AccountAliasParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountAliasOK, error)

	AccountBlock(params *AccountBlockParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountBlockOK, error)

	AccountCreate(params *AccountCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountCreateOK, error)

	AccountDelete(params *AccountDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountDeleteAccepted, error)

	AccountFollow(params *AccountFollowParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountFollowOK, error)

	AccountFollowers(params *AccountFollowersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountFollowersOK, error)

	AccountFollowing(params *AccountFollowingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountFollowingOK, error)

	AccountGet(params *AccountGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountGetOK, error)

	AccountLists(params *AccountListsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountListsOK, error)

	AccountLookupGet(params *AccountLookupGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountLookupGetOK, error)

	AccountMove(params *AccountMoveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountMoveAccepted, error)

	AccountNote(params *AccountNoteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountNoteOK, error)

	AccountRelationships(params *AccountRelationshipsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountRelationshipsOK, error)

	AccountSearchGet(params *AccountSearchGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountSearchGetOK, error)

	AccountStatuses(params *AccountStatusesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountStatusesOK, error)

	AccountUnblock(params *AccountUnblockParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountUnblockOK, error)

	AccountUnfollow(params *AccountUnfollowParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountUnfollowOK, error)

	AccountUpdate(params *AccountUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountUpdateOK, error)

	AccountVerify(params *AccountVerifyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountVerifyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	AccountAlias aliases your account to another account by setting also known as to the given URI

	This is useful when you want to move from another account this this account.

In such cases, you should set the alsoKnownAs of this account to the URI of
the account you want to move from.
*/
func (a *Client) AccountAlias(params *AccountAliasParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountAliasOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountAliasParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountAlias",
		Method:             "POST",
		PathPattern:        "/api/v1/accounts/alias",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AccountAliasReader{formats: a.formats},
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
	success, ok := result.(*AccountAliasOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for accountAlias: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AccountBlock blocks account with id
*/
func (a *Client) AccountBlock(params *AccountBlockParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountBlockOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountBlockParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountBlock",
		Method:             "POST",
		PathPattern:        "/api/v1/accounts/{id}/block",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AccountBlockReader{formats: a.formats},
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
	success, ok := result.(*AccountBlockOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for accountBlock: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	AccountCreate creates a new account using an application token

	The parameters can also be given in the body of the request, as JSON, if the content-type is set to 'application/json'.

The parameters can also be given in the body of the request, as XML, if the content-type is set to 'application/xml'.
*/
func (a *Client) AccountCreate(params *AccountCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountCreate",
		Method:             "POST",
		PathPattern:        "/api/v1/accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded", "application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AccountCreateReader{formats: a.formats},
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
	success, ok := result.(*AccountCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for accountCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AccountDelete deletes your account
*/
func (a *Client) AccountDelete(params *AccountDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountDeleteAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountDelete",
		Method:             "POST",
		PathPattern:        "/api/v1/accounts/delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AccountDeleteReader{formats: a.formats},
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
	success, ok := result.(*AccountDeleteAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for accountDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	AccountFollow follows account with id

	The parameters can also be given in the body of the request, as JSON, if the content-type is set to 'application/json'.

The parameters can also be given in the body of the request, as XML, if the content-type is set to 'application/xml'.

If you already follow (request) the given account, then the follow (request) will be updated instead using the
`reblogs` and `notify` parameters.
*/
func (a *Client) AccountFollow(params *AccountFollowParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountFollowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountFollowParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountFollow",
		Method:             "POST",
		PathPattern:        "/api/v1/accounts/{id}/follow",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded", "application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AccountFollowReader{formats: a.formats},
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
	success, ok := result.(*AccountFollowOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for accountFollow: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	AccountFollowers sees followers of account with given id

	The next and previous queries can be parsed from the returned Link header.

Example:

```
<https://example.org/api/v1/accounts/0657WMDEC3KQDTD6NZ4XJZBK4M/followers?limit=80&max_id=01FC0SKA48HNSVR6YKZCQGS2V8>; rel="next", <https://example.org/api/v1/accounts/0657WMDEC3KQDTD6NZ4XJZBK4M/followers?limit=80&min_id=01FC0SKW5JK2Q4EVAV2B462YY0>; rel="prev"
````
*/
func (a *Client) AccountFollowers(params *AccountFollowersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountFollowersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountFollowersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountFollowers",
		Method:             "GET",
		PathPattern:        "/api/v1/accounts/{id}/followers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AccountFollowersReader{formats: a.formats},
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
	success, ok := result.(*AccountFollowersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for accountFollowers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	AccountFollowing sees accounts followed by given account id

	The next and previous queries can be parsed from the returned Link header.

Example:

```
<https://example.org/api/v1/accounts/0657WMDEC3KQDTD6NZ4XJZBK4M/following?limit=80&max_id=01FC0SKA48HNSVR6YKZCQGS2V8>; rel="next", <https://example.org/api/v1/accounts/0657WMDEC3KQDTD6NZ4XJZBK4M/following?limit=80&min_id=01FC0SKW5JK2Q4EVAV2B462YY0>; rel="prev"
````
*/
func (a *Client) AccountFollowing(params *AccountFollowingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountFollowingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountFollowingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountFollowing",
		Method:             "GET",
		PathPattern:        "/api/v1/accounts/{id}/following",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AccountFollowingReader{formats: a.formats},
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
	success, ok := result.(*AccountFollowingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for accountFollowing: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AccountGet gets information about an account with the given ID
*/
func (a *Client) AccountGet(params *AccountGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountGet",
		Method:             "GET",
		PathPattern:        "/api/v1/accounts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AccountGetReader{formats: a.formats},
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
	success, ok := result.(*AccountGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for accountGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AccountLists sees all lists of yours that contain requested account
*/
func (a *Client) AccountLists(params *AccountListsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountListsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountListsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountLists",
		Method:             "GET",
		PathPattern:        "/api/v1/accounts/{id}/lists",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AccountListsReader{formats: a.formats},
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
	success, ok := result.(*AccountListsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for accountLists: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AccountLookupGet quicklies lookup a username to see if it is available skipping web finger resolution
*/
func (a *Client) AccountLookupGet(params *AccountLookupGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountLookupGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountLookupGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountLookupGet",
		Method:             "GET",
		PathPattern:        "/api/v1/accounts/lookup",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AccountLookupGetReader{formats: a.formats},
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
	success, ok := result.(*AccountLookupGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for accountLookupGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AccountMove moves your account to another account

NOT IMPLEMENTED YET!
*/
func (a *Client) AccountMove(params *AccountMoveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountMoveAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountMoveParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountMove",
		Method:             "POST",
		PathPattern:        "/api/v1/accounts/move",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AccountMoveReader{formats: a.formats},
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
	success, ok := result.(*AccountMoveAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for accountMove: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AccountNote sets a private note for an account with the given id
*/
func (a *Client) AccountNote(params *AccountNoteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountNoteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountNoteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountNote",
		Method:             "POST",
		PathPattern:        "/api/v1/accounts/{id}/note",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AccountNoteReader{formats: a.formats},
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
	success, ok := result.(*AccountNoteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for accountNote: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AccountRelationships sees your account s relationships with the given account i ds
*/
func (a *Client) AccountRelationships(params *AccountRelationshipsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountRelationshipsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountRelationshipsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountRelationships",
		Method:             "GET",
		PathPattern:        "/api/v1/accounts/relationships",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AccountRelationshipsReader{formats: a.formats},
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
	success, ok := result.(*AccountRelationshipsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for accountRelationships: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AccountSearchGet searches for accounts by username and or display name
*/
func (a *Client) AccountSearchGet(params *AccountSearchGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountSearchGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountSearchGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountSearchGet",
		Method:             "GET",
		PathPattern:        "/api/v1/accounts/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AccountSearchGetReader{formats: a.formats},
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
	success, ok := result.(*AccountSearchGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for accountSearchGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AccountStatuses sees statuses posted by the requested account

The statuses will be returned in descending chronological order (newest first), with sequential IDs (bigger = newer).
*/
func (a *Client) AccountStatuses(params *AccountStatusesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountStatusesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountStatusesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountStatuses",
		Method:             "GET",
		PathPattern:        "/api/v1/accounts/{id}/statuses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AccountStatusesReader{formats: a.formats},
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
	success, ok := result.(*AccountStatusesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for accountStatuses: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AccountUnblock unblocks account with ID
*/
func (a *Client) AccountUnblock(params *AccountUnblockParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountUnblockOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountUnblockParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountUnblock",
		Method:             "POST",
		PathPattern:        "/api/v1/accounts/{id}/unblock",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AccountUnblockReader{formats: a.formats},
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
	success, ok := result.(*AccountUnblockOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for accountUnblock: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AccountUnfollow unfollows account with id
*/
func (a *Client) AccountUnfollow(params *AccountUnfollowParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountUnfollowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountUnfollowParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountUnfollow",
		Method:             "POST",
		PathPattern:        "/api/v1/accounts/{id}/unfollow",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AccountUnfollowReader{formats: a.formats},
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
	success, ok := result.(*AccountUnfollowOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for accountUnfollow: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AccountUpdate updates your account
*/
func (a *Client) AccountUpdate(params *AccountUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountUpdate",
		Method:             "PATCH",
		PathPattern:        "/api/v1/accounts/update_credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AccountUpdateReader{formats: a.formats},
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
	success, ok := result.(*AccountUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for accountUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AccountVerify verifies a token by returning account details pertaining to it
*/
func (a *Client) AccountVerify(params *AccountVerifyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountVerifyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountVerifyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountVerify",
		Method:             "GET",
		PathPattern:        "/api/v1/accounts/verify_credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AccountVerifyReader{formats: a.formats},
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
	success, ok := result.(*AccountVerifyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for accountVerify: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
