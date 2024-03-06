// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VyrCossont/slurp/models"
)

// AccountListsReader is a Reader for the AccountLists structure.
type AccountListsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountListsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountListsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccountListsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAccountListsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccountListsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewAccountListsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAccountListsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/accounts/{id}/lists] accountLists", response, response.Code())
	}
}

// NewAccountListsOK creates a AccountListsOK with default headers values
func NewAccountListsOK() *AccountListsOK {
	return &AccountListsOK{}
}

/*
AccountListsOK describes a response with status code 200, with default header values.

Array of all lists containing this account.
*/
type AccountListsOK struct {
	Payload []*models.List
}

// IsSuccess returns true when this account lists o k response has a 2xx status code
func (o *AccountListsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account lists o k response has a 3xx status code
func (o *AccountListsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account lists o k response has a 4xx status code
func (o *AccountListsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account lists o k response has a 5xx status code
func (o *AccountListsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account lists o k response a status code equal to that given
func (o *AccountListsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the account lists o k response
func (o *AccountListsOK) Code() int {
	return 200
}

func (o *AccountListsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/lists][%d] accountListsOK  %+v", 200, o.Payload)
}

func (o *AccountListsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/lists][%d] accountListsOK  %+v", 200, o.Payload)
}

func (o *AccountListsOK) GetPayload() []*models.List {
	return o.Payload
}

func (o *AccountListsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountListsBadRequest creates a AccountListsBadRequest with default headers values
func NewAccountListsBadRequest() *AccountListsBadRequest {
	return &AccountListsBadRequest{}
}

/*
AccountListsBadRequest describes a response with status code 400, with default header values.

bad request
*/
type AccountListsBadRequest struct {
}

// IsSuccess returns true when this account lists bad request response has a 2xx status code
func (o *AccountListsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account lists bad request response has a 3xx status code
func (o *AccountListsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account lists bad request response has a 4xx status code
func (o *AccountListsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this account lists bad request response has a 5xx status code
func (o *AccountListsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this account lists bad request response a status code equal to that given
func (o *AccountListsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the account lists bad request response
func (o *AccountListsBadRequest) Code() int {
	return 400
}

func (o *AccountListsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/lists][%d] accountListsBadRequest ", 400)
}

func (o *AccountListsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/lists][%d] accountListsBadRequest ", 400)
}

func (o *AccountListsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountListsUnauthorized creates a AccountListsUnauthorized with default headers values
func NewAccountListsUnauthorized() *AccountListsUnauthorized {
	return &AccountListsUnauthorized{}
}

/*
AccountListsUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type AccountListsUnauthorized struct {
}

// IsSuccess returns true when this account lists unauthorized response has a 2xx status code
func (o *AccountListsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account lists unauthorized response has a 3xx status code
func (o *AccountListsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account lists unauthorized response has a 4xx status code
func (o *AccountListsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this account lists unauthorized response has a 5xx status code
func (o *AccountListsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this account lists unauthorized response a status code equal to that given
func (o *AccountListsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the account lists unauthorized response
func (o *AccountListsUnauthorized) Code() int {
	return 401
}

func (o *AccountListsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/lists][%d] accountListsUnauthorized ", 401)
}

func (o *AccountListsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/lists][%d] accountListsUnauthorized ", 401)
}

func (o *AccountListsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountListsNotFound creates a AccountListsNotFound with default headers values
func NewAccountListsNotFound() *AccountListsNotFound {
	return &AccountListsNotFound{}
}

/*
AccountListsNotFound describes a response with status code 404, with default header values.

not found
*/
type AccountListsNotFound struct {
}

// IsSuccess returns true when this account lists not found response has a 2xx status code
func (o *AccountListsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account lists not found response has a 3xx status code
func (o *AccountListsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account lists not found response has a 4xx status code
func (o *AccountListsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this account lists not found response has a 5xx status code
func (o *AccountListsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this account lists not found response a status code equal to that given
func (o *AccountListsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the account lists not found response
func (o *AccountListsNotFound) Code() int {
	return 404
}

func (o *AccountListsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/lists][%d] accountListsNotFound ", 404)
}

func (o *AccountListsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/lists][%d] accountListsNotFound ", 404)
}

func (o *AccountListsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountListsNotAcceptable creates a AccountListsNotAcceptable with default headers values
func NewAccountListsNotAcceptable() *AccountListsNotAcceptable {
	return &AccountListsNotAcceptable{}
}

/*
AccountListsNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type AccountListsNotAcceptable struct {
}

// IsSuccess returns true when this account lists not acceptable response has a 2xx status code
func (o *AccountListsNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account lists not acceptable response has a 3xx status code
func (o *AccountListsNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account lists not acceptable response has a 4xx status code
func (o *AccountListsNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this account lists not acceptable response has a 5xx status code
func (o *AccountListsNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this account lists not acceptable response a status code equal to that given
func (o *AccountListsNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the account lists not acceptable response
func (o *AccountListsNotAcceptable) Code() int {
	return 406
}

func (o *AccountListsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/lists][%d] accountListsNotAcceptable ", 406)
}

func (o *AccountListsNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/lists][%d] accountListsNotAcceptable ", 406)
}

func (o *AccountListsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountListsInternalServerError creates a AccountListsInternalServerError with default headers values
func NewAccountListsInternalServerError() *AccountListsInternalServerError {
	return &AccountListsInternalServerError{}
}

/*
AccountListsInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type AccountListsInternalServerError struct {
}

// IsSuccess returns true when this account lists internal server error response has a 2xx status code
func (o *AccountListsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account lists internal server error response has a 3xx status code
func (o *AccountListsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account lists internal server error response has a 4xx status code
func (o *AccountListsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this account lists internal server error response has a 5xx status code
func (o *AccountListsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this account lists internal server error response a status code equal to that given
func (o *AccountListsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the account lists internal server error response
func (o *AccountListsInternalServerError) Code() int {
	return 500
}

func (o *AccountListsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/lists][%d] accountListsInternalServerError ", 500)
}

func (o *AccountListsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/lists][%d] accountListsInternalServerError ", 500)
}

func (o *AccountListsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
