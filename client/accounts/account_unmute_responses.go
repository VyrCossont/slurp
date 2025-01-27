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

// AccountUnmuteReader is a Reader for the AccountUnmute structure.
type AccountUnmuteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountUnmuteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountUnmuteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccountUnmuteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAccountUnmuteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccountUnmuteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewAccountUnmuteNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAccountUnmuteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/accounts/{id}/unmute] accountUnmute", response, response.Code())
	}
}

// NewAccountUnmuteOK creates a AccountUnmuteOK with default headers values
func NewAccountUnmuteOK() *AccountUnmuteOK {
	return &AccountUnmuteOK{}
}

/*
AccountUnmuteOK describes a response with status code 200, with default header values.

Your relationship to this account.
*/
type AccountUnmuteOK struct {
	Payload *models.Relationship
}

// IsSuccess returns true when this account unmute o k response has a 2xx status code
func (o *AccountUnmuteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account unmute o k response has a 3xx status code
func (o *AccountUnmuteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account unmute o k response has a 4xx status code
func (o *AccountUnmuteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account unmute o k response has a 5xx status code
func (o *AccountUnmuteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account unmute o k response a status code equal to that given
func (o *AccountUnmuteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the account unmute o k response
func (o *AccountUnmuteOK) Code() int {
	return 200
}

func (o *AccountUnmuteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/unmute][%d] accountUnmuteOK %s", 200, payload)
}

func (o *AccountUnmuteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/unmute][%d] accountUnmuteOK %s", 200, payload)
}

func (o *AccountUnmuteOK) GetPayload() *models.Relationship {
	return o.Payload
}

func (o *AccountUnmuteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Relationship)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountUnmuteBadRequest creates a AccountUnmuteBadRequest with default headers values
func NewAccountUnmuteBadRequest() *AccountUnmuteBadRequest {
	return &AccountUnmuteBadRequest{}
}

/*
AccountUnmuteBadRequest describes a response with status code 400, with default header values.

bad request
*/
type AccountUnmuteBadRequest struct {
}

// IsSuccess returns true when this account unmute bad request response has a 2xx status code
func (o *AccountUnmuteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account unmute bad request response has a 3xx status code
func (o *AccountUnmuteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account unmute bad request response has a 4xx status code
func (o *AccountUnmuteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this account unmute bad request response has a 5xx status code
func (o *AccountUnmuteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this account unmute bad request response a status code equal to that given
func (o *AccountUnmuteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the account unmute bad request response
func (o *AccountUnmuteBadRequest) Code() int {
	return 400
}

func (o *AccountUnmuteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/unmute][%d] accountUnmuteBadRequest", 400)
}

func (o *AccountUnmuteBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/unmute][%d] accountUnmuteBadRequest", 400)
}

func (o *AccountUnmuteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountUnmuteUnauthorized creates a AccountUnmuteUnauthorized with default headers values
func NewAccountUnmuteUnauthorized() *AccountUnmuteUnauthorized {
	return &AccountUnmuteUnauthorized{}
}

/*
AccountUnmuteUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type AccountUnmuteUnauthorized struct {
}

// IsSuccess returns true when this account unmute unauthorized response has a 2xx status code
func (o *AccountUnmuteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account unmute unauthorized response has a 3xx status code
func (o *AccountUnmuteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account unmute unauthorized response has a 4xx status code
func (o *AccountUnmuteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this account unmute unauthorized response has a 5xx status code
func (o *AccountUnmuteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this account unmute unauthorized response a status code equal to that given
func (o *AccountUnmuteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the account unmute unauthorized response
func (o *AccountUnmuteUnauthorized) Code() int {
	return 401
}

func (o *AccountUnmuteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/unmute][%d] accountUnmuteUnauthorized", 401)
}

func (o *AccountUnmuteUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/unmute][%d] accountUnmuteUnauthorized", 401)
}

func (o *AccountUnmuteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountUnmuteNotFound creates a AccountUnmuteNotFound with default headers values
func NewAccountUnmuteNotFound() *AccountUnmuteNotFound {
	return &AccountUnmuteNotFound{}
}

/*
AccountUnmuteNotFound describes a response with status code 404, with default header values.

not found
*/
type AccountUnmuteNotFound struct {
}

// IsSuccess returns true when this account unmute not found response has a 2xx status code
func (o *AccountUnmuteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account unmute not found response has a 3xx status code
func (o *AccountUnmuteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account unmute not found response has a 4xx status code
func (o *AccountUnmuteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this account unmute not found response has a 5xx status code
func (o *AccountUnmuteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this account unmute not found response a status code equal to that given
func (o *AccountUnmuteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the account unmute not found response
func (o *AccountUnmuteNotFound) Code() int {
	return 404
}

func (o *AccountUnmuteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/unmute][%d] accountUnmuteNotFound", 404)
}

func (o *AccountUnmuteNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/unmute][%d] accountUnmuteNotFound", 404)
}

func (o *AccountUnmuteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountUnmuteNotAcceptable creates a AccountUnmuteNotAcceptable with default headers values
func NewAccountUnmuteNotAcceptable() *AccountUnmuteNotAcceptable {
	return &AccountUnmuteNotAcceptable{}
}

/*
AccountUnmuteNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type AccountUnmuteNotAcceptable struct {
}

// IsSuccess returns true when this account unmute not acceptable response has a 2xx status code
func (o *AccountUnmuteNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account unmute not acceptable response has a 3xx status code
func (o *AccountUnmuteNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account unmute not acceptable response has a 4xx status code
func (o *AccountUnmuteNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this account unmute not acceptable response has a 5xx status code
func (o *AccountUnmuteNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this account unmute not acceptable response a status code equal to that given
func (o *AccountUnmuteNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the account unmute not acceptable response
func (o *AccountUnmuteNotAcceptable) Code() int {
	return 406
}

func (o *AccountUnmuteNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/unmute][%d] accountUnmuteNotAcceptable", 406)
}

func (o *AccountUnmuteNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/unmute][%d] accountUnmuteNotAcceptable", 406)
}

func (o *AccountUnmuteNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountUnmuteInternalServerError creates a AccountUnmuteInternalServerError with default headers values
func NewAccountUnmuteInternalServerError() *AccountUnmuteInternalServerError {
	return &AccountUnmuteInternalServerError{}
}

/*
AccountUnmuteInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type AccountUnmuteInternalServerError struct {
}

// IsSuccess returns true when this account unmute internal server error response has a 2xx status code
func (o *AccountUnmuteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account unmute internal server error response has a 3xx status code
func (o *AccountUnmuteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account unmute internal server error response has a 4xx status code
func (o *AccountUnmuteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this account unmute internal server error response has a 5xx status code
func (o *AccountUnmuteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this account unmute internal server error response a status code equal to that given
func (o *AccountUnmuteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the account unmute internal server error response
func (o *AccountUnmuteInternalServerError) Code() int {
	return 500
}

func (o *AccountUnmuteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/unmute][%d] accountUnmuteInternalServerError", 500)
}

func (o *AccountUnmuteInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/accounts/{id}/unmute][%d] accountUnmuteInternalServerError", 500)
}

func (o *AccountUnmuteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
