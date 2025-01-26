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

// StatusUnbookmarkReader is a Reader for the StatusUnbookmark structure.
type StatusUnbookmarkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusUnbookmarkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatusUnbookmarkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStatusUnbookmarkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStatusUnbookmarkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStatusUnbookmarkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStatusUnbookmarkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewStatusUnbookmarkNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStatusUnbookmarkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/statuses/{id}/unbookmark] statusUnbookmark", response, response.Code())
	}
}

// NewStatusUnbookmarkOK creates a StatusUnbookmarkOK with default headers values
func NewStatusUnbookmarkOK() *StatusUnbookmarkOK {
	return &StatusUnbookmarkOK{}
}

/*
StatusUnbookmarkOK describes a response with status code 200, with default header values.

The status.
*/
type StatusUnbookmarkOK struct {
	Payload *models.Status
}

// IsSuccess returns true when this status unbookmark o k response has a 2xx status code
func (o *StatusUnbookmarkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this status unbookmark o k response has a 3xx status code
func (o *StatusUnbookmarkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status unbookmark o k response has a 4xx status code
func (o *StatusUnbookmarkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this status unbookmark o k response has a 5xx status code
func (o *StatusUnbookmarkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this status unbookmark o k response a status code equal to that given
func (o *StatusUnbookmarkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the status unbookmark o k response
func (o *StatusUnbookmarkOK) Code() int {
	return 200
}

func (o *StatusUnbookmarkOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/unbookmark][%d] statusUnbookmarkOK %s", 200, payload)
}

func (o *StatusUnbookmarkOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/unbookmark][%d] statusUnbookmarkOK %s", 200, payload)
}

func (o *StatusUnbookmarkOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *StatusUnbookmarkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatusUnbookmarkBadRequest creates a StatusUnbookmarkBadRequest with default headers values
func NewStatusUnbookmarkBadRequest() *StatusUnbookmarkBadRequest {
	return &StatusUnbookmarkBadRequest{}
}

/*
StatusUnbookmarkBadRequest describes a response with status code 400, with default header values.

bad request
*/
type StatusUnbookmarkBadRequest struct {
}

// IsSuccess returns true when this status unbookmark bad request response has a 2xx status code
func (o *StatusUnbookmarkBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status unbookmark bad request response has a 3xx status code
func (o *StatusUnbookmarkBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status unbookmark bad request response has a 4xx status code
func (o *StatusUnbookmarkBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this status unbookmark bad request response has a 5xx status code
func (o *StatusUnbookmarkBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this status unbookmark bad request response a status code equal to that given
func (o *StatusUnbookmarkBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the status unbookmark bad request response
func (o *StatusUnbookmarkBadRequest) Code() int {
	return 400
}

func (o *StatusUnbookmarkBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/unbookmark][%d] statusUnbookmarkBadRequest", 400)
}

func (o *StatusUnbookmarkBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/unbookmark][%d] statusUnbookmarkBadRequest", 400)
}

func (o *StatusUnbookmarkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusUnbookmarkUnauthorized creates a StatusUnbookmarkUnauthorized with default headers values
func NewStatusUnbookmarkUnauthorized() *StatusUnbookmarkUnauthorized {
	return &StatusUnbookmarkUnauthorized{}
}

/*
StatusUnbookmarkUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type StatusUnbookmarkUnauthorized struct {
}

// IsSuccess returns true when this status unbookmark unauthorized response has a 2xx status code
func (o *StatusUnbookmarkUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status unbookmark unauthorized response has a 3xx status code
func (o *StatusUnbookmarkUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status unbookmark unauthorized response has a 4xx status code
func (o *StatusUnbookmarkUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this status unbookmark unauthorized response has a 5xx status code
func (o *StatusUnbookmarkUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this status unbookmark unauthorized response a status code equal to that given
func (o *StatusUnbookmarkUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the status unbookmark unauthorized response
func (o *StatusUnbookmarkUnauthorized) Code() int {
	return 401
}

func (o *StatusUnbookmarkUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/unbookmark][%d] statusUnbookmarkUnauthorized", 401)
}

func (o *StatusUnbookmarkUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/unbookmark][%d] statusUnbookmarkUnauthorized", 401)
}

func (o *StatusUnbookmarkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusUnbookmarkForbidden creates a StatusUnbookmarkForbidden with default headers values
func NewStatusUnbookmarkForbidden() *StatusUnbookmarkForbidden {
	return &StatusUnbookmarkForbidden{}
}

/*
StatusUnbookmarkForbidden describes a response with status code 403, with default header values.

forbidden
*/
type StatusUnbookmarkForbidden struct {
}

// IsSuccess returns true when this status unbookmark forbidden response has a 2xx status code
func (o *StatusUnbookmarkForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status unbookmark forbidden response has a 3xx status code
func (o *StatusUnbookmarkForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status unbookmark forbidden response has a 4xx status code
func (o *StatusUnbookmarkForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this status unbookmark forbidden response has a 5xx status code
func (o *StatusUnbookmarkForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this status unbookmark forbidden response a status code equal to that given
func (o *StatusUnbookmarkForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the status unbookmark forbidden response
func (o *StatusUnbookmarkForbidden) Code() int {
	return 403
}

func (o *StatusUnbookmarkForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/unbookmark][%d] statusUnbookmarkForbidden", 403)
}

func (o *StatusUnbookmarkForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/unbookmark][%d] statusUnbookmarkForbidden", 403)
}

func (o *StatusUnbookmarkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusUnbookmarkNotFound creates a StatusUnbookmarkNotFound with default headers values
func NewStatusUnbookmarkNotFound() *StatusUnbookmarkNotFound {
	return &StatusUnbookmarkNotFound{}
}

/*
StatusUnbookmarkNotFound describes a response with status code 404, with default header values.

not found
*/
type StatusUnbookmarkNotFound struct {
}

// IsSuccess returns true when this status unbookmark not found response has a 2xx status code
func (o *StatusUnbookmarkNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status unbookmark not found response has a 3xx status code
func (o *StatusUnbookmarkNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status unbookmark not found response has a 4xx status code
func (o *StatusUnbookmarkNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this status unbookmark not found response has a 5xx status code
func (o *StatusUnbookmarkNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this status unbookmark not found response a status code equal to that given
func (o *StatusUnbookmarkNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the status unbookmark not found response
func (o *StatusUnbookmarkNotFound) Code() int {
	return 404
}

func (o *StatusUnbookmarkNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/unbookmark][%d] statusUnbookmarkNotFound", 404)
}

func (o *StatusUnbookmarkNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/unbookmark][%d] statusUnbookmarkNotFound", 404)
}

func (o *StatusUnbookmarkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusUnbookmarkNotAcceptable creates a StatusUnbookmarkNotAcceptable with default headers values
func NewStatusUnbookmarkNotAcceptable() *StatusUnbookmarkNotAcceptable {
	return &StatusUnbookmarkNotAcceptable{}
}

/*
StatusUnbookmarkNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type StatusUnbookmarkNotAcceptable struct {
}

// IsSuccess returns true when this status unbookmark not acceptable response has a 2xx status code
func (o *StatusUnbookmarkNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status unbookmark not acceptable response has a 3xx status code
func (o *StatusUnbookmarkNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status unbookmark not acceptable response has a 4xx status code
func (o *StatusUnbookmarkNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this status unbookmark not acceptable response has a 5xx status code
func (o *StatusUnbookmarkNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this status unbookmark not acceptable response a status code equal to that given
func (o *StatusUnbookmarkNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the status unbookmark not acceptable response
func (o *StatusUnbookmarkNotAcceptable) Code() int {
	return 406
}

func (o *StatusUnbookmarkNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/unbookmark][%d] statusUnbookmarkNotAcceptable", 406)
}

func (o *StatusUnbookmarkNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/unbookmark][%d] statusUnbookmarkNotAcceptable", 406)
}

func (o *StatusUnbookmarkNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusUnbookmarkInternalServerError creates a StatusUnbookmarkInternalServerError with default headers values
func NewStatusUnbookmarkInternalServerError() *StatusUnbookmarkInternalServerError {
	return &StatusUnbookmarkInternalServerError{}
}

/*
StatusUnbookmarkInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type StatusUnbookmarkInternalServerError struct {
}

// IsSuccess returns true when this status unbookmark internal server error response has a 2xx status code
func (o *StatusUnbookmarkInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status unbookmark internal server error response has a 3xx status code
func (o *StatusUnbookmarkInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status unbookmark internal server error response has a 4xx status code
func (o *StatusUnbookmarkInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this status unbookmark internal server error response has a 5xx status code
func (o *StatusUnbookmarkInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this status unbookmark internal server error response a status code equal to that given
func (o *StatusUnbookmarkInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the status unbookmark internal server error response
func (o *StatusUnbookmarkInternalServerError) Code() int {
	return 500
}

func (o *StatusUnbookmarkInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/unbookmark][%d] statusUnbookmarkInternalServerError", 500)
}

func (o *StatusUnbookmarkInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/statuses/{id}/unbookmark][%d] statusUnbookmarkInternalServerError", 500)
}

func (o *StatusUnbookmarkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
