// Code generated by go-swagger; DO NOT EDIT.

package statuses

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

// StatusPinReader is a Reader for the StatusPin structure.
type StatusPinReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusPinReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatusPinOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStatusPinBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStatusPinUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStatusPinForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStatusPinNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewStatusPinNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStatusPinInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/statuses/{id}/pin] statusPin", response, response.Code())
	}
}

// NewStatusPinOK creates a StatusPinOK with default headers values
func NewStatusPinOK() *StatusPinOK {
	return &StatusPinOK{}
}

/*
StatusPinOK describes a response with status code 200, with default header values.

The status.
*/
type StatusPinOK struct {
	Payload *models.Status
}

// IsSuccess returns true when this status pin o k response has a 2xx status code
func (o *StatusPinOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this status pin o k response has a 3xx status code
func (o *StatusPinOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status pin o k response has a 4xx status code
func (o *StatusPinOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this status pin o k response has a 5xx status code
func (o *StatusPinOK) IsServerError() bool {
	return false
}

// IsCode returns true when this status pin o k response a status code equal to that given
func (o *StatusPinOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the status pin o k response
func (o *StatusPinOK) Code() int {
	return 200
}

func (o *StatusPinOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/pin][%d] statusPinOK %s", 200, payload)
}

func (o *StatusPinOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/pin][%d] statusPinOK %s", 200, payload)
}

func (o *StatusPinOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *StatusPinOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatusPinBadRequest creates a StatusPinBadRequest with default headers values
func NewStatusPinBadRequest() *StatusPinBadRequest {
	return &StatusPinBadRequest{}
}

/*
StatusPinBadRequest describes a response with status code 400, with default header values.

bad request
*/
type StatusPinBadRequest struct {
}

// IsSuccess returns true when this status pin bad request response has a 2xx status code
func (o *StatusPinBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status pin bad request response has a 3xx status code
func (o *StatusPinBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status pin bad request response has a 4xx status code
func (o *StatusPinBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this status pin bad request response has a 5xx status code
func (o *StatusPinBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this status pin bad request response a status code equal to that given
func (o *StatusPinBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the status pin bad request response
func (o *StatusPinBadRequest) Code() int {
	return 400
}

func (o *StatusPinBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/pin][%d] statusPinBadRequest", 400)
}

func (o *StatusPinBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/pin][%d] statusPinBadRequest", 400)
}

func (o *StatusPinBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusPinUnauthorized creates a StatusPinUnauthorized with default headers values
func NewStatusPinUnauthorized() *StatusPinUnauthorized {
	return &StatusPinUnauthorized{}
}

/*
StatusPinUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type StatusPinUnauthorized struct {
}

// IsSuccess returns true when this status pin unauthorized response has a 2xx status code
func (o *StatusPinUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status pin unauthorized response has a 3xx status code
func (o *StatusPinUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status pin unauthorized response has a 4xx status code
func (o *StatusPinUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this status pin unauthorized response has a 5xx status code
func (o *StatusPinUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this status pin unauthorized response a status code equal to that given
func (o *StatusPinUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the status pin unauthorized response
func (o *StatusPinUnauthorized) Code() int {
	return 401
}

func (o *StatusPinUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/pin][%d] statusPinUnauthorized", 401)
}

func (o *StatusPinUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/pin][%d] statusPinUnauthorized", 401)
}

func (o *StatusPinUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusPinForbidden creates a StatusPinForbidden with default headers values
func NewStatusPinForbidden() *StatusPinForbidden {
	return &StatusPinForbidden{}
}

/*
StatusPinForbidden describes a response with status code 403, with default header values.

forbidden
*/
type StatusPinForbidden struct {
}

// IsSuccess returns true when this status pin forbidden response has a 2xx status code
func (o *StatusPinForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status pin forbidden response has a 3xx status code
func (o *StatusPinForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status pin forbidden response has a 4xx status code
func (o *StatusPinForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this status pin forbidden response has a 5xx status code
func (o *StatusPinForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this status pin forbidden response a status code equal to that given
func (o *StatusPinForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the status pin forbidden response
func (o *StatusPinForbidden) Code() int {
	return 403
}

func (o *StatusPinForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/pin][%d] statusPinForbidden", 403)
}

func (o *StatusPinForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/pin][%d] statusPinForbidden", 403)
}

func (o *StatusPinForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusPinNotFound creates a StatusPinNotFound with default headers values
func NewStatusPinNotFound() *StatusPinNotFound {
	return &StatusPinNotFound{}
}

/*
StatusPinNotFound describes a response with status code 404, with default header values.

not found
*/
type StatusPinNotFound struct {
}

// IsSuccess returns true when this status pin not found response has a 2xx status code
func (o *StatusPinNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status pin not found response has a 3xx status code
func (o *StatusPinNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status pin not found response has a 4xx status code
func (o *StatusPinNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this status pin not found response has a 5xx status code
func (o *StatusPinNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this status pin not found response a status code equal to that given
func (o *StatusPinNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the status pin not found response
func (o *StatusPinNotFound) Code() int {
	return 404
}

func (o *StatusPinNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/pin][%d] statusPinNotFound", 404)
}

func (o *StatusPinNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/pin][%d] statusPinNotFound", 404)
}

func (o *StatusPinNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusPinNotAcceptable creates a StatusPinNotAcceptable with default headers values
func NewStatusPinNotAcceptable() *StatusPinNotAcceptable {
	return &StatusPinNotAcceptable{}
}

/*
StatusPinNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type StatusPinNotAcceptable struct {
}

// IsSuccess returns true when this status pin not acceptable response has a 2xx status code
func (o *StatusPinNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status pin not acceptable response has a 3xx status code
func (o *StatusPinNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status pin not acceptable response has a 4xx status code
func (o *StatusPinNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this status pin not acceptable response has a 5xx status code
func (o *StatusPinNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this status pin not acceptable response a status code equal to that given
func (o *StatusPinNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the status pin not acceptable response
func (o *StatusPinNotAcceptable) Code() int {
	return 406
}

func (o *StatusPinNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/pin][%d] statusPinNotAcceptable", 406)
}

func (o *StatusPinNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/pin][%d] statusPinNotAcceptable", 406)
}

func (o *StatusPinNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusPinInternalServerError creates a StatusPinInternalServerError with default headers values
func NewStatusPinInternalServerError() *StatusPinInternalServerError {
	return &StatusPinInternalServerError{}
}

/*
StatusPinInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type StatusPinInternalServerError struct {
}

// IsSuccess returns true when this status pin internal server error response has a 2xx status code
func (o *StatusPinInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status pin internal server error response has a 3xx status code
func (o *StatusPinInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status pin internal server error response has a 4xx status code
func (o *StatusPinInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this status pin internal server error response has a 5xx status code
func (o *StatusPinInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this status pin internal server error response a status code equal to that given
func (o *StatusPinInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the status pin internal server error response
func (o *StatusPinInternalServerError) Code() int {
	return 500
}

func (o *StatusPinInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/pin][%d] statusPinInternalServerError", 500)
}

func (o *StatusPinInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/pin][%d] statusPinInternalServerError", 500)
}

func (o *StatusPinInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
