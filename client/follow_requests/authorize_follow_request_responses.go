// Code generated by go-swagger; DO NOT EDIT.

package follow_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VyrCossont/slurp/models"
)

// AuthorizeFollowRequestReader is a Reader for the AuthorizeFollowRequest structure.
type AuthorizeFollowRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthorizeFollowRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthorizeFollowRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAuthorizeFollowRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAuthorizeFollowRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAuthorizeFollowRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewAuthorizeFollowRequestNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAuthorizeFollowRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/follow_requests/{account_id}/authorize] authorizeFollowRequest", response, response.Code())
	}
}

// NewAuthorizeFollowRequestOK creates a AuthorizeFollowRequestOK with default headers values
func NewAuthorizeFollowRequestOK() *AuthorizeFollowRequestOK {
	return &AuthorizeFollowRequestOK{}
}

/*
AuthorizeFollowRequestOK describes a response with status code 200, with default header values.

Your relationship to this account.
*/
type AuthorizeFollowRequestOK struct {
	Payload *models.Relationship
}

// IsSuccess returns true when this authorize follow request o k response has a 2xx status code
func (o *AuthorizeFollowRequestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this authorize follow request o k response has a 3xx status code
func (o *AuthorizeFollowRequestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authorize follow request o k response has a 4xx status code
func (o *AuthorizeFollowRequestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this authorize follow request o k response has a 5xx status code
func (o *AuthorizeFollowRequestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this authorize follow request o k response a status code equal to that given
func (o *AuthorizeFollowRequestOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the authorize follow request o k response
func (o *AuthorizeFollowRequestOK) Code() int {
	return 200
}

func (o *AuthorizeFollowRequestOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/follow_requests/{account_id}/authorize][%d] authorizeFollowRequestOK %s", 200, payload)
}

func (o *AuthorizeFollowRequestOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/follow_requests/{account_id}/authorize][%d] authorizeFollowRequestOK %s", 200, payload)
}

func (o *AuthorizeFollowRequestOK) GetPayload() *models.Relationship {
	return o.Payload
}

func (o *AuthorizeFollowRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Relationship)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthorizeFollowRequestBadRequest creates a AuthorizeFollowRequestBadRequest with default headers values
func NewAuthorizeFollowRequestBadRequest() *AuthorizeFollowRequestBadRequest {
	return &AuthorizeFollowRequestBadRequest{}
}

/*
AuthorizeFollowRequestBadRequest describes a response with status code 400, with default header values.

bad request
*/
type AuthorizeFollowRequestBadRequest struct {
}

// IsSuccess returns true when this authorize follow request bad request response has a 2xx status code
func (o *AuthorizeFollowRequestBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this authorize follow request bad request response has a 3xx status code
func (o *AuthorizeFollowRequestBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authorize follow request bad request response has a 4xx status code
func (o *AuthorizeFollowRequestBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this authorize follow request bad request response has a 5xx status code
func (o *AuthorizeFollowRequestBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this authorize follow request bad request response a status code equal to that given
func (o *AuthorizeFollowRequestBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the authorize follow request bad request response
func (o *AuthorizeFollowRequestBadRequest) Code() int {
	return 400
}

func (o *AuthorizeFollowRequestBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/follow_requests/{account_id}/authorize][%d] authorizeFollowRequestBadRequest", 400)
}

func (o *AuthorizeFollowRequestBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/follow_requests/{account_id}/authorize][%d] authorizeFollowRequestBadRequest", 400)
}

func (o *AuthorizeFollowRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAuthorizeFollowRequestUnauthorized creates a AuthorizeFollowRequestUnauthorized with default headers values
func NewAuthorizeFollowRequestUnauthorized() *AuthorizeFollowRequestUnauthorized {
	return &AuthorizeFollowRequestUnauthorized{}
}

/*
AuthorizeFollowRequestUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type AuthorizeFollowRequestUnauthorized struct {
}

// IsSuccess returns true when this authorize follow request unauthorized response has a 2xx status code
func (o *AuthorizeFollowRequestUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this authorize follow request unauthorized response has a 3xx status code
func (o *AuthorizeFollowRequestUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authorize follow request unauthorized response has a 4xx status code
func (o *AuthorizeFollowRequestUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this authorize follow request unauthorized response has a 5xx status code
func (o *AuthorizeFollowRequestUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this authorize follow request unauthorized response a status code equal to that given
func (o *AuthorizeFollowRequestUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the authorize follow request unauthorized response
func (o *AuthorizeFollowRequestUnauthorized) Code() int {
	return 401
}

func (o *AuthorizeFollowRequestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/follow_requests/{account_id}/authorize][%d] authorizeFollowRequestUnauthorized", 401)
}

func (o *AuthorizeFollowRequestUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/follow_requests/{account_id}/authorize][%d] authorizeFollowRequestUnauthorized", 401)
}

func (o *AuthorizeFollowRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAuthorizeFollowRequestNotFound creates a AuthorizeFollowRequestNotFound with default headers values
func NewAuthorizeFollowRequestNotFound() *AuthorizeFollowRequestNotFound {
	return &AuthorizeFollowRequestNotFound{}
}

/*
AuthorizeFollowRequestNotFound describes a response with status code 404, with default header values.

not found
*/
type AuthorizeFollowRequestNotFound struct {
}

// IsSuccess returns true when this authorize follow request not found response has a 2xx status code
func (o *AuthorizeFollowRequestNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this authorize follow request not found response has a 3xx status code
func (o *AuthorizeFollowRequestNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authorize follow request not found response has a 4xx status code
func (o *AuthorizeFollowRequestNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this authorize follow request not found response has a 5xx status code
func (o *AuthorizeFollowRequestNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this authorize follow request not found response a status code equal to that given
func (o *AuthorizeFollowRequestNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the authorize follow request not found response
func (o *AuthorizeFollowRequestNotFound) Code() int {
	return 404
}

func (o *AuthorizeFollowRequestNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/follow_requests/{account_id}/authorize][%d] authorizeFollowRequestNotFound", 404)
}

func (o *AuthorizeFollowRequestNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/follow_requests/{account_id}/authorize][%d] authorizeFollowRequestNotFound", 404)
}

func (o *AuthorizeFollowRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAuthorizeFollowRequestNotAcceptable creates a AuthorizeFollowRequestNotAcceptable with default headers values
func NewAuthorizeFollowRequestNotAcceptable() *AuthorizeFollowRequestNotAcceptable {
	return &AuthorizeFollowRequestNotAcceptable{}
}

/*
AuthorizeFollowRequestNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type AuthorizeFollowRequestNotAcceptable struct {
}

// IsSuccess returns true when this authorize follow request not acceptable response has a 2xx status code
func (o *AuthorizeFollowRequestNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this authorize follow request not acceptable response has a 3xx status code
func (o *AuthorizeFollowRequestNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authorize follow request not acceptable response has a 4xx status code
func (o *AuthorizeFollowRequestNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this authorize follow request not acceptable response has a 5xx status code
func (o *AuthorizeFollowRequestNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this authorize follow request not acceptable response a status code equal to that given
func (o *AuthorizeFollowRequestNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the authorize follow request not acceptable response
func (o *AuthorizeFollowRequestNotAcceptable) Code() int {
	return 406
}

func (o *AuthorizeFollowRequestNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/follow_requests/{account_id}/authorize][%d] authorizeFollowRequestNotAcceptable", 406)
}

func (o *AuthorizeFollowRequestNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v1/follow_requests/{account_id}/authorize][%d] authorizeFollowRequestNotAcceptable", 406)
}

func (o *AuthorizeFollowRequestNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAuthorizeFollowRequestInternalServerError creates a AuthorizeFollowRequestInternalServerError with default headers values
func NewAuthorizeFollowRequestInternalServerError() *AuthorizeFollowRequestInternalServerError {
	return &AuthorizeFollowRequestInternalServerError{}
}

/*
AuthorizeFollowRequestInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type AuthorizeFollowRequestInternalServerError struct {
}

// IsSuccess returns true when this authorize follow request internal server error response has a 2xx status code
func (o *AuthorizeFollowRequestInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this authorize follow request internal server error response has a 3xx status code
func (o *AuthorizeFollowRequestInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authorize follow request internal server error response has a 4xx status code
func (o *AuthorizeFollowRequestInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this authorize follow request internal server error response has a 5xx status code
func (o *AuthorizeFollowRequestInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this authorize follow request internal server error response a status code equal to that given
func (o *AuthorizeFollowRequestInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the authorize follow request internal server error response
func (o *AuthorizeFollowRequestInternalServerError) Code() int {
	return 500
}

func (o *AuthorizeFollowRequestInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/follow_requests/{account_id}/authorize][%d] authorizeFollowRequestInternalServerError", 500)
}

func (o *AuthorizeFollowRequestInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/follow_requests/{account_id}/authorize][%d] authorizeFollowRequestInternalServerError", 500)
}

func (o *AuthorizeFollowRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
