// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

// AccountLookupGetReader is a Reader for the AccountLookupGet structure.
type AccountLookupGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountLookupGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountLookupGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccountLookupGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAccountLookupGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccountLookupGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewAccountLookupGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAccountLookupGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/accounts/lookup] accountLookupGet", response, response.Code())
	}
}

// NewAccountLookupGetOK creates a AccountLookupGetOK with default headers values
func NewAccountLookupGetOK() *AccountLookupGetOK {
	return &AccountLookupGetOK{}
}

/*
AccountLookupGetOK describes a response with status code 200, with default header values.

Result of the lookup.
*/
type AccountLookupGetOK struct {
	Payload *models.Account
}

// IsSuccess returns true when this account lookup get o k response has a 2xx status code
func (o *AccountLookupGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account lookup get o k response has a 3xx status code
func (o *AccountLookupGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account lookup get o k response has a 4xx status code
func (o *AccountLookupGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account lookup get o k response has a 5xx status code
func (o *AccountLookupGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account lookup get o k response a status code equal to that given
func (o *AccountLookupGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the account lookup get o k response
func (o *AccountLookupGetOK) Code() int {
	return 200
}

func (o *AccountLookupGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/accounts/lookup][%d] accountLookupGetOK %s", 200, payload)
}

func (o *AccountLookupGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/accounts/lookup][%d] accountLookupGetOK %s", 200, payload)
}

func (o *AccountLookupGetOK) GetPayload() *models.Account {
	return o.Payload
}

func (o *AccountLookupGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Account)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountLookupGetBadRequest creates a AccountLookupGetBadRequest with default headers values
func NewAccountLookupGetBadRequest() *AccountLookupGetBadRequest {
	return &AccountLookupGetBadRequest{}
}

/*
AccountLookupGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type AccountLookupGetBadRequest struct {
}

// IsSuccess returns true when this account lookup get bad request response has a 2xx status code
func (o *AccountLookupGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account lookup get bad request response has a 3xx status code
func (o *AccountLookupGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account lookup get bad request response has a 4xx status code
func (o *AccountLookupGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this account lookup get bad request response has a 5xx status code
func (o *AccountLookupGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this account lookup get bad request response a status code equal to that given
func (o *AccountLookupGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the account lookup get bad request response
func (o *AccountLookupGetBadRequest) Code() int {
	return 400
}

func (o *AccountLookupGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/lookup][%d] accountLookupGetBadRequest", 400)
}

func (o *AccountLookupGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/lookup][%d] accountLookupGetBadRequest", 400)
}

func (o *AccountLookupGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountLookupGetUnauthorized creates a AccountLookupGetUnauthorized with default headers values
func NewAccountLookupGetUnauthorized() *AccountLookupGetUnauthorized {
	return &AccountLookupGetUnauthorized{}
}

/*
AccountLookupGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type AccountLookupGetUnauthorized struct {
}

// IsSuccess returns true when this account lookup get unauthorized response has a 2xx status code
func (o *AccountLookupGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account lookup get unauthorized response has a 3xx status code
func (o *AccountLookupGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account lookup get unauthorized response has a 4xx status code
func (o *AccountLookupGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this account lookup get unauthorized response has a 5xx status code
func (o *AccountLookupGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this account lookup get unauthorized response a status code equal to that given
func (o *AccountLookupGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the account lookup get unauthorized response
func (o *AccountLookupGetUnauthorized) Code() int {
	return 401
}

func (o *AccountLookupGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/lookup][%d] accountLookupGetUnauthorized", 401)
}

func (o *AccountLookupGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/lookup][%d] accountLookupGetUnauthorized", 401)
}

func (o *AccountLookupGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountLookupGetNotFound creates a AccountLookupGetNotFound with default headers values
func NewAccountLookupGetNotFound() *AccountLookupGetNotFound {
	return &AccountLookupGetNotFound{}
}

/*
AccountLookupGetNotFound describes a response with status code 404, with default header values.

not found
*/
type AccountLookupGetNotFound struct {
}

// IsSuccess returns true when this account lookup get not found response has a 2xx status code
func (o *AccountLookupGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account lookup get not found response has a 3xx status code
func (o *AccountLookupGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account lookup get not found response has a 4xx status code
func (o *AccountLookupGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this account lookup get not found response has a 5xx status code
func (o *AccountLookupGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this account lookup get not found response a status code equal to that given
func (o *AccountLookupGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the account lookup get not found response
func (o *AccountLookupGetNotFound) Code() int {
	return 404
}

func (o *AccountLookupGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/lookup][%d] accountLookupGetNotFound", 404)
}

func (o *AccountLookupGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/lookup][%d] accountLookupGetNotFound", 404)
}

func (o *AccountLookupGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountLookupGetNotAcceptable creates a AccountLookupGetNotAcceptable with default headers values
func NewAccountLookupGetNotAcceptable() *AccountLookupGetNotAcceptable {
	return &AccountLookupGetNotAcceptable{}
}

/*
AccountLookupGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type AccountLookupGetNotAcceptable struct {
}

// IsSuccess returns true when this account lookup get not acceptable response has a 2xx status code
func (o *AccountLookupGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account lookup get not acceptable response has a 3xx status code
func (o *AccountLookupGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account lookup get not acceptable response has a 4xx status code
func (o *AccountLookupGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this account lookup get not acceptable response has a 5xx status code
func (o *AccountLookupGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this account lookup get not acceptable response a status code equal to that given
func (o *AccountLookupGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the account lookup get not acceptable response
func (o *AccountLookupGetNotAcceptable) Code() int {
	return 406
}

func (o *AccountLookupGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/lookup][%d] accountLookupGetNotAcceptable", 406)
}

func (o *AccountLookupGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/lookup][%d] accountLookupGetNotAcceptable", 406)
}

func (o *AccountLookupGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountLookupGetInternalServerError creates a AccountLookupGetInternalServerError with default headers values
func NewAccountLookupGetInternalServerError() *AccountLookupGetInternalServerError {
	return &AccountLookupGetInternalServerError{}
}

/*
AccountLookupGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type AccountLookupGetInternalServerError struct {
}

// IsSuccess returns true when this account lookup get internal server error response has a 2xx status code
func (o *AccountLookupGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account lookup get internal server error response has a 3xx status code
func (o *AccountLookupGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account lookup get internal server error response has a 4xx status code
func (o *AccountLookupGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this account lookup get internal server error response has a 5xx status code
func (o *AccountLookupGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this account lookup get internal server error response a status code equal to that given
func (o *AccountLookupGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the account lookup get internal server error response
func (o *AccountLookupGetInternalServerError) Code() int {
	return 500
}

func (o *AccountLookupGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/lookup][%d] accountLookupGetInternalServerError", 500)
}

func (o *AccountLookupGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/lookup][%d] accountLookupGetInternalServerError", 500)
}

func (o *AccountLookupGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
