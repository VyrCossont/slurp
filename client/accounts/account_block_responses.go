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

// AccountBlockReader is a Reader for the AccountBlock structure.
type AccountBlockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountBlockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountBlockOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccountBlockBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAccountBlockUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccountBlockNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewAccountBlockNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAccountBlockInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/accounts/{id}/block] accountBlock", response, response.Code())
	}
}

// NewAccountBlockOK creates a AccountBlockOK with default headers values
func NewAccountBlockOK() *AccountBlockOK {
	return &AccountBlockOK{}
}

/*
AccountBlockOK describes a response with status code 200, with default header values.

Your relationship to the account.
*/
type AccountBlockOK struct {
	Payload *models.Relationship
}

// IsSuccess returns true when this account block o k response has a 2xx status code
func (o *AccountBlockOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account block o k response has a 3xx status code
func (o *AccountBlockOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account block o k response has a 4xx status code
func (o *AccountBlockOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account block o k response has a 5xx status code
func (o *AccountBlockOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account block o k response a status code equal to that given
func (o *AccountBlockOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the account block o k response
func (o *AccountBlockOK) Code() int {
	return 200
}

func (o *AccountBlockOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/block][%d] accountBlockOK  %+v", 200, o.Payload)
}

func (o *AccountBlockOK) String() string {
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/block][%d] accountBlockOK  %+v", 200, o.Payload)
}

func (o *AccountBlockOK) GetPayload() *models.Relationship {
	return o.Payload
}

func (o *AccountBlockOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Relationship)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountBlockBadRequest creates a AccountBlockBadRequest with default headers values
func NewAccountBlockBadRequest() *AccountBlockBadRequest {
	return &AccountBlockBadRequest{}
}

/*
AccountBlockBadRequest describes a response with status code 400, with default header values.

bad request
*/
type AccountBlockBadRequest struct {
}

// IsSuccess returns true when this account block bad request response has a 2xx status code
func (o *AccountBlockBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account block bad request response has a 3xx status code
func (o *AccountBlockBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account block bad request response has a 4xx status code
func (o *AccountBlockBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this account block bad request response has a 5xx status code
func (o *AccountBlockBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this account block bad request response a status code equal to that given
func (o *AccountBlockBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the account block bad request response
func (o *AccountBlockBadRequest) Code() int {
	return 400
}

func (o *AccountBlockBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/block][%d] accountBlockBadRequest ", 400)
}

func (o *AccountBlockBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/block][%d] accountBlockBadRequest ", 400)
}

func (o *AccountBlockBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountBlockUnauthorized creates a AccountBlockUnauthorized with default headers values
func NewAccountBlockUnauthorized() *AccountBlockUnauthorized {
	return &AccountBlockUnauthorized{}
}

/*
AccountBlockUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type AccountBlockUnauthorized struct {
}

// IsSuccess returns true when this account block unauthorized response has a 2xx status code
func (o *AccountBlockUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account block unauthorized response has a 3xx status code
func (o *AccountBlockUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account block unauthorized response has a 4xx status code
func (o *AccountBlockUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this account block unauthorized response has a 5xx status code
func (o *AccountBlockUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this account block unauthorized response a status code equal to that given
func (o *AccountBlockUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the account block unauthorized response
func (o *AccountBlockUnauthorized) Code() int {
	return 401
}

func (o *AccountBlockUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/block][%d] accountBlockUnauthorized ", 401)
}

func (o *AccountBlockUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/block][%d] accountBlockUnauthorized ", 401)
}

func (o *AccountBlockUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountBlockNotFound creates a AccountBlockNotFound with default headers values
func NewAccountBlockNotFound() *AccountBlockNotFound {
	return &AccountBlockNotFound{}
}

/*
AccountBlockNotFound describes a response with status code 404, with default header values.

not found
*/
type AccountBlockNotFound struct {
}

// IsSuccess returns true when this account block not found response has a 2xx status code
func (o *AccountBlockNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account block not found response has a 3xx status code
func (o *AccountBlockNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account block not found response has a 4xx status code
func (o *AccountBlockNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this account block not found response has a 5xx status code
func (o *AccountBlockNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this account block not found response a status code equal to that given
func (o *AccountBlockNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the account block not found response
func (o *AccountBlockNotFound) Code() int {
	return 404
}

func (o *AccountBlockNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/block][%d] accountBlockNotFound ", 404)
}

func (o *AccountBlockNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/block][%d] accountBlockNotFound ", 404)
}

func (o *AccountBlockNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountBlockNotAcceptable creates a AccountBlockNotAcceptable with default headers values
func NewAccountBlockNotAcceptable() *AccountBlockNotAcceptable {
	return &AccountBlockNotAcceptable{}
}

/*
AccountBlockNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type AccountBlockNotAcceptable struct {
}

// IsSuccess returns true when this account block not acceptable response has a 2xx status code
func (o *AccountBlockNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account block not acceptable response has a 3xx status code
func (o *AccountBlockNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account block not acceptable response has a 4xx status code
func (o *AccountBlockNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this account block not acceptable response has a 5xx status code
func (o *AccountBlockNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this account block not acceptable response a status code equal to that given
func (o *AccountBlockNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the account block not acceptable response
func (o *AccountBlockNotAcceptable) Code() int {
	return 406
}

func (o *AccountBlockNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/block][%d] accountBlockNotAcceptable ", 406)
}

func (o *AccountBlockNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/block][%d] accountBlockNotAcceptable ", 406)
}

func (o *AccountBlockNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountBlockInternalServerError creates a AccountBlockInternalServerError with default headers values
func NewAccountBlockInternalServerError() *AccountBlockInternalServerError {
	return &AccountBlockInternalServerError{}
}

/*
AccountBlockInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type AccountBlockInternalServerError struct {
}

// IsSuccess returns true when this account block internal server error response has a 2xx status code
func (o *AccountBlockInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account block internal server error response has a 3xx status code
func (o *AccountBlockInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account block internal server error response has a 4xx status code
func (o *AccountBlockInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this account block internal server error response has a 5xx status code
func (o *AccountBlockInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this account block internal server error response a status code equal to that given
func (o *AccountBlockInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the account block internal server error response
func (o *AccountBlockInternalServerError) Code() int {
	return 500
}

func (o *AccountBlockInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/block][%d] accountBlockInternalServerError ", 500)
}

func (o *AccountBlockInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/block][%d] accountBlockInternalServerError ", 500)
}

func (o *AccountBlockInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
