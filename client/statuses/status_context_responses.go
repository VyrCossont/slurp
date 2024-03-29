// Code generated by go-swagger; DO NOT EDIT.

package statuses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VyrCossont/slurp/models"
)

// StatusContextReader is a Reader for the StatusContext structure.
type StatusContextReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusContextReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatusContextOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStatusContextBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStatusContextUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStatusContextForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStatusContextNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewStatusContextNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStatusContextInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/statuses/{id}/context] statusContext", response, response.Code())
	}
}

// NewStatusContextOK creates a StatusContextOK with default headers values
func NewStatusContextOK() *StatusContextOK {
	return &StatusContextOK{}
}

/*
StatusContextOK describes a response with status code 200, with default header values.

Status context object.
*/
type StatusContextOK struct {
	Payload *models.Context
}

// IsSuccess returns true when this status context o k response has a 2xx status code
func (o *StatusContextOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this status context o k response has a 3xx status code
func (o *StatusContextOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status context o k response has a 4xx status code
func (o *StatusContextOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this status context o k response has a 5xx status code
func (o *StatusContextOK) IsServerError() bool {
	return false
}

// IsCode returns true when this status context o k response a status code equal to that given
func (o *StatusContextOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the status context o k response
func (o *StatusContextOK) Code() int {
	return 200
}

func (o *StatusContextOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/context][%d] statusContextOK  %+v", 200, o.Payload)
}

func (o *StatusContextOK) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/context][%d] statusContextOK  %+v", 200, o.Payload)
}

func (o *StatusContextOK) GetPayload() *models.Context {
	return o.Payload
}

func (o *StatusContextOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Context)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatusContextBadRequest creates a StatusContextBadRequest with default headers values
func NewStatusContextBadRequest() *StatusContextBadRequest {
	return &StatusContextBadRequest{}
}

/*
StatusContextBadRequest describes a response with status code 400, with default header values.

bad request
*/
type StatusContextBadRequest struct {
}

// IsSuccess returns true when this status context bad request response has a 2xx status code
func (o *StatusContextBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status context bad request response has a 3xx status code
func (o *StatusContextBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status context bad request response has a 4xx status code
func (o *StatusContextBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this status context bad request response has a 5xx status code
func (o *StatusContextBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this status context bad request response a status code equal to that given
func (o *StatusContextBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the status context bad request response
func (o *StatusContextBadRequest) Code() int {
	return 400
}

func (o *StatusContextBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/context][%d] statusContextBadRequest ", 400)
}

func (o *StatusContextBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/context][%d] statusContextBadRequest ", 400)
}

func (o *StatusContextBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusContextUnauthorized creates a StatusContextUnauthorized with default headers values
func NewStatusContextUnauthorized() *StatusContextUnauthorized {
	return &StatusContextUnauthorized{}
}

/*
StatusContextUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type StatusContextUnauthorized struct {
}

// IsSuccess returns true when this status context unauthorized response has a 2xx status code
func (o *StatusContextUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status context unauthorized response has a 3xx status code
func (o *StatusContextUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status context unauthorized response has a 4xx status code
func (o *StatusContextUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this status context unauthorized response has a 5xx status code
func (o *StatusContextUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this status context unauthorized response a status code equal to that given
func (o *StatusContextUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the status context unauthorized response
func (o *StatusContextUnauthorized) Code() int {
	return 401
}

func (o *StatusContextUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/context][%d] statusContextUnauthorized ", 401)
}

func (o *StatusContextUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/context][%d] statusContextUnauthorized ", 401)
}

func (o *StatusContextUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusContextForbidden creates a StatusContextForbidden with default headers values
func NewStatusContextForbidden() *StatusContextForbidden {
	return &StatusContextForbidden{}
}

/*
StatusContextForbidden describes a response with status code 403, with default header values.

forbidden
*/
type StatusContextForbidden struct {
}

// IsSuccess returns true when this status context forbidden response has a 2xx status code
func (o *StatusContextForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status context forbidden response has a 3xx status code
func (o *StatusContextForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status context forbidden response has a 4xx status code
func (o *StatusContextForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this status context forbidden response has a 5xx status code
func (o *StatusContextForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this status context forbidden response a status code equal to that given
func (o *StatusContextForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the status context forbidden response
func (o *StatusContextForbidden) Code() int {
	return 403
}

func (o *StatusContextForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/context][%d] statusContextForbidden ", 403)
}

func (o *StatusContextForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/context][%d] statusContextForbidden ", 403)
}

func (o *StatusContextForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusContextNotFound creates a StatusContextNotFound with default headers values
func NewStatusContextNotFound() *StatusContextNotFound {
	return &StatusContextNotFound{}
}

/*
StatusContextNotFound describes a response with status code 404, with default header values.

not found
*/
type StatusContextNotFound struct {
}

// IsSuccess returns true when this status context not found response has a 2xx status code
func (o *StatusContextNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status context not found response has a 3xx status code
func (o *StatusContextNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status context not found response has a 4xx status code
func (o *StatusContextNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this status context not found response has a 5xx status code
func (o *StatusContextNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this status context not found response a status code equal to that given
func (o *StatusContextNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the status context not found response
func (o *StatusContextNotFound) Code() int {
	return 404
}

func (o *StatusContextNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/context][%d] statusContextNotFound ", 404)
}

func (o *StatusContextNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/context][%d] statusContextNotFound ", 404)
}

func (o *StatusContextNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusContextNotAcceptable creates a StatusContextNotAcceptable with default headers values
func NewStatusContextNotAcceptable() *StatusContextNotAcceptable {
	return &StatusContextNotAcceptable{}
}

/*
StatusContextNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type StatusContextNotAcceptable struct {
}

// IsSuccess returns true when this status context not acceptable response has a 2xx status code
func (o *StatusContextNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status context not acceptable response has a 3xx status code
func (o *StatusContextNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status context not acceptable response has a 4xx status code
func (o *StatusContextNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this status context not acceptable response has a 5xx status code
func (o *StatusContextNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this status context not acceptable response a status code equal to that given
func (o *StatusContextNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the status context not acceptable response
func (o *StatusContextNotAcceptable) Code() int {
	return 406
}

func (o *StatusContextNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/context][%d] statusContextNotAcceptable ", 406)
}

func (o *StatusContextNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/context][%d] statusContextNotAcceptable ", 406)
}

func (o *StatusContextNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusContextInternalServerError creates a StatusContextInternalServerError with default headers values
func NewStatusContextInternalServerError() *StatusContextInternalServerError {
	return &StatusContextInternalServerError{}
}

/*
StatusContextInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type StatusContextInternalServerError struct {
}

// IsSuccess returns true when this status context internal server error response has a 2xx status code
func (o *StatusContextInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status context internal server error response has a 3xx status code
func (o *StatusContextInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status context internal server error response has a 4xx status code
func (o *StatusContextInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this status context internal server error response has a 5xx status code
func (o *StatusContextInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this status context internal server error response a status code equal to that given
func (o *StatusContextInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the status context internal server error response
func (o *StatusContextInternalServerError) Code() int {
	return 500
}

func (o *StatusContextInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/context][%d] statusContextInternalServerError ", 500)
}

func (o *StatusContextInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/context][%d] statusContextInternalServerError ", 500)
}

func (o *StatusContextInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
