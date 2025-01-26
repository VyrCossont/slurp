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

// AccountGetReader is a Reader for the AccountGet structure.
type AccountGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccountGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAccountGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccountGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewAccountGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAccountGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/accounts/{id}] accountGet", response, response.Code())
	}
}

// NewAccountGetOK creates a AccountGetOK with default headers values
func NewAccountGetOK() *AccountGetOK {
	return &AccountGetOK{}
}

/*
AccountGetOK describes a response with status code 200, with default header values.

The requested account.
*/
type AccountGetOK struct {
	Payload *models.Account
}

// IsSuccess returns true when this account get o k response has a 2xx status code
func (o *AccountGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account get o k response has a 3xx status code
func (o *AccountGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account get o k response has a 4xx status code
func (o *AccountGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account get o k response has a 5xx status code
func (o *AccountGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account get o k response a status code equal to that given
func (o *AccountGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the account get o k response
func (o *AccountGetOK) Code() int {
	return 200
}

func (o *AccountGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/accounts/{id}][%d] accountGetOK %s", 200, payload)
}

func (o *AccountGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/accounts/{id}][%d] accountGetOK %s", 200, payload)
}

func (o *AccountGetOK) GetPayload() *models.Account {
	return o.Payload
}

func (o *AccountGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Account)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountGetBadRequest creates a AccountGetBadRequest with default headers values
func NewAccountGetBadRequest() *AccountGetBadRequest {
	return &AccountGetBadRequest{}
}

/*
AccountGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type AccountGetBadRequest struct {
}

// IsSuccess returns true when this account get bad request response has a 2xx status code
func (o *AccountGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account get bad request response has a 3xx status code
func (o *AccountGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account get bad request response has a 4xx status code
func (o *AccountGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this account get bad request response has a 5xx status code
func (o *AccountGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this account get bad request response a status code equal to that given
func (o *AccountGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the account get bad request response
func (o *AccountGetBadRequest) Code() int {
	return 400
}

func (o *AccountGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}][%d] accountGetBadRequest", 400)
}

func (o *AccountGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}][%d] accountGetBadRequest", 400)
}

func (o *AccountGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountGetUnauthorized creates a AccountGetUnauthorized with default headers values
func NewAccountGetUnauthorized() *AccountGetUnauthorized {
	return &AccountGetUnauthorized{}
}

/*
AccountGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type AccountGetUnauthorized struct {
}

// IsSuccess returns true when this account get unauthorized response has a 2xx status code
func (o *AccountGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account get unauthorized response has a 3xx status code
func (o *AccountGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account get unauthorized response has a 4xx status code
func (o *AccountGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this account get unauthorized response has a 5xx status code
func (o *AccountGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this account get unauthorized response a status code equal to that given
func (o *AccountGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the account get unauthorized response
func (o *AccountGetUnauthorized) Code() int {
	return 401
}

func (o *AccountGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}][%d] accountGetUnauthorized", 401)
}

func (o *AccountGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}][%d] accountGetUnauthorized", 401)
}

func (o *AccountGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountGetNotFound creates a AccountGetNotFound with default headers values
func NewAccountGetNotFound() *AccountGetNotFound {
	return &AccountGetNotFound{}
}

/*
AccountGetNotFound describes a response with status code 404, with default header values.

not found
*/
type AccountGetNotFound struct {
}

// IsSuccess returns true when this account get not found response has a 2xx status code
func (o *AccountGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account get not found response has a 3xx status code
func (o *AccountGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account get not found response has a 4xx status code
func (o *AccountGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this account get not found response has a 5xx status code
func (o *AccountGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this account get not found response a status code equal to that given
func (o *AccountGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the account get not found response
func (o *AccountGetNotFound) Code() int {
	return 404
}

func (o *AccountGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}][%d] accountGetNotFound", 404)
}

func (o *AccountGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}][%d] accountGetNotFound", 404)
}

func (o *AccountGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountGetNotAcceptable creates a AccountGetNotAcceptable with default headers values
func NewAccountGetNotAcceptable() *AccountGetNotAcceptable {
	return &AccountGetNotAcceptable{}
}

/*
AccountGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type AccountGetNotAcceptable struct {
}

// IsSuccess returns true when this account get not acceptable response has a 2xx status code
func (o *AccountGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account get not acceptable response has a 3xx status code
func (o *AccountGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account get not acceptable response has a 4xx status code
func (o *AccountGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this account get not acceptable response has a 5xx status code
func (o *AccountGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this account get not acceptable response a status code equal to that given
func (o *AccountGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the account get not acceptable response
func (o *AccountGetNotAcceptable) Code() int {
	return 406
}

func (o *AccountGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}][%d] accountGetNotAcceptable", 406)
}

func (o *AccountGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}][%d] accountGetNotAcceptable", 406)
}

func (o *AccountGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountGetInternalServerError creates a AccountGetInternalServerError with default headers values
func NewAccountGetInternalServerError() *AccountGetInternalServerError {
	return &AccountGetInternalServerError{}
}

/*
AccountGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type AccountGetInternalServerError struct {
}

// IsSuccess returns true when this account get internal server error response has a 2xx status code
func (o *AccountGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account get internal server error response has a 3xx status code
func (o *AccountGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account get internal server error response has a 4xx status code
func (o *AccountGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this account get internal server error response has a 5xx status code
func (o *AccountGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this account get internal server error response a status code equal to that given
func (o *AccountGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the account get internal server error response
func (o *AccountGetInternalServerError) Code() int {
	return 500
}

func (o *AccountGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}][%d] accountGetInternalServerError", 500)
}

func (o *AccountGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}][%d] accountGetInternalServerError", 500)
}

func (o *AccountGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
