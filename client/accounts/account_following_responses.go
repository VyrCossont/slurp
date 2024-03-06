// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AccountFollowingReader is a Reader for the AccountFollowing structure.
type AccountFollowingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountFollowingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountFollowingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccountFollowingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAccountFollowingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccountFollowingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewAccountFollowingNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAccountFollowingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/accounts/{id}/following] accountFollowing", response, response.Code())
	}
}

// NewAccountFollowingOK creates a AccountFollowingOK with default headers values
func NewAccountFollowingOK() *AccountFollowingOK {
	return &AccountFollowingOK{}
}

/*
AccountFollowingOK describes a response with status code 200, with default header values.

Array of accounts that are followed by this account.
*/
type AccountFollowingOK struct {
}

// IsSuccess returns true when this account following o k response has a 2xx status code
func (o *AccountFollowingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account following o k response has a 3xx status code
func (o *AccountFollowingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account following o k response has a 4xx status code
func (o *AccountFollowingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account following o k response has a 5xx status code
func (o *AccountFollowingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account following o k response a status code equal to that given
func (o *AccountFollowingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the account following o k response
func (o *AccountFollowingOK) Code() int {
	return 200
}

func (o *AccountFollowingOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/following][%d] accountFollowingOK ", 200)
}

func (o *AccountFollowingOK) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/following][%d] accountFollowingOK ", 200)
}

func (o *AccountFollowingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountFollowingBadRequest creates a AccountFollowingBadRequest with default headers values
func NewAccountFollowingBadRequest() *AccountFollowingBadRequest {
	return &AccountFollowingBadRequest{}
}

/*
AccountFollowingBadRequest describes a response with status code 400, with default header values.

bad request
*/
type AccountFollowingBadRequest struct {
}

// IsSuccess returns true when this account following bad request response has a 2xx status code
func (o *AccountFollowingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account following bad request response has a 3xx status code
func (o *AccountFollowingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account following bad request response has a 4xx status code
func (o *AccountFollowingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this account following bad request response has a 5xx status code
func (o *AccountFollowingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this account following bad request response a status code equal to that given
func (o *AccountFollowingBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the account following bad request response
func (o *AccountFollowingBadRequest) Code() int {
	return 400
}

func (o *AccountFollowingBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/following][%d] accountFollowingBadRequest ", 400)
}

func (o *AccountFollowingBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/following][%d] accountFollowingBadRequest ", 400)
}

func (o *AccountFollowingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountFollowingUnauthorized creates a AccountFollowingUnauthorized with default headers values
func NewAccountFollowingUnauthorized() *AccountFollowingUnauthorized {
	return &AccountFollowingUnauthorized{}
}

/*
AccountFollowingUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type AccountFollowingUnauthorized struct {
}

// IsSuccess returns true when this account following unauthorized response has a 2xx status code
func (o *AccountFollowingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account following unauthorized response has a 3xx status code
func (o *AccountFollowingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account following unauthorized response has a 4xx status code
func (o *AccountFollowingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this account following unauthorized response has a 5xx status code
func (o *AccountFollowingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this account following unauthorized response a status code equal to that given
func (o *AccountFollowingUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the account following unauthorized response
func (o *AccountFollowingUnauthorized) Code() int {
	return 401
}

func (o *AccountFollowingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/following][%d] accountFollowingUnauthorized ", 401)
}

func (o *AccountFollowingUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/following][%d] accountFollowingUnauthorized ", 401)
}

func (o *AccountFollowingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountFollowingNotFound creates a AccountFollowingNotFound with default headers values
func NewAccountFollowingNotFound() *AccountFollowingNotFound {
	return &AccountFollowingNotFound{}
}

/*
AccountFollowingNotFound describes a response with status code 404, with default header values.

not found
*/
type AccountFollowingNotFound struct {
}

// IsSuccess returns true when this account following not found response has a 2xx status code
func (o *AccountFollowingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account following not found response has a 3xx status code
func (o *AccountFollowingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account following not found response has a 4xx status code
func (o *AccountFollowingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this account following not found response has a 5xx status code
func (o *AccountFollowingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this account following not found response a status code equal to that given
func (o *AccountFollowingNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the account following not found response
func (o *AccountFollowingNotFound) Code() int {
	return 404
}

func (o *AccountFollowingNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/following][%d] accountFollowingNotFound ", 404)
}

func (o *AccountFollowingNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/following][%d] accountFollowingNotFound ", 404)
}

func (o *AccountFollowingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountFollowingNotAcceptable creates a AccountFollowingNotAcceptable with default headers values
func NewAccountFollowingNotAcceptable() *AccountFollowingNotAcceptable {
	return &AccountFollowingNotAcceptable{}
}

/*
AccountFollowingNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type AccountFollowingNotAcceptable struct {
}

// IsSuccess returns true when this account following not acceptable response has a 2xx status code
func (o *AccountFollowingNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account following not acceptable response has a 3xx status code
func (o *AccountFollowingNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account following not acceptable response has a 4xx status code
func (o *AccountFollowingNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this account following not acceptable response has a 5xx status code
func (o *AccountFollowingNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this account following not acceptable response a status code equal to that given
func (o *AccountFollowingNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the account following not acceptable response
func (o *AccountFollowingNotAcceptable) Code() int {
	return 406
}

func (o *AccountFollowingNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/following][%d] accountFollowingNotAcceptable ", 406)
}

func (o *AccountFollowingNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/following][%d] accountFollowingNotAcceptable ", 406)
}

func (o *AccountFollowingNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountFollowingInternalServerError creates a AccountFollowingInternalServerError with default headers values
func NewAccountFollowingInternalServerError() *AccountFollowingInternalServerError {
	return &AccountFollowingInternalServerError{}
}

/*
AccountFollowingInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type AccountFollowingInternalServerError struct {
}

// IsSuccess returns true when this account following internal server error response has a 2xx status code
func (o *AccountFollowingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account following internal server error response has a 3xx status code
func (o *AccountFollowingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account following internal server error response has a 4xx status code
func (o *AccountFollowingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this account following internal server error response has a 5xx status code
func (o *AccountFollowingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this account following internal server error response a status code equal to that given
func (o *AccountFollowingInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the account following internal server error response
func (o *AccountFollowingInternalServerError) Code() int {
	return 500
}

func (o *AccountFollowingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/following][%d] accountFollowingInternalServerError ", 500)
}

func (o *AccountFollowingInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/following][%d] accountFollowingInternalServerError ", 500)
}

func (o *AccountFollowingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
