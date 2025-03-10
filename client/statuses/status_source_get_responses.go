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

// StatusSourceGetReader is a Reader for the StatusSourceGet structure.
type StatusSourceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusSourceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatusSourceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStatusSourceGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStatusSourceGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStatusSourceGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStatusSourceGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewStatusSourceGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStatusSourceGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/statuses/{id}/source] statusSourceGet", response, response.Code())
	}
}

// NewStatusSourceGetOK creates a StatusSourceGetOK with default headers values
func NewStatusSourceGetOK() *StatusSourceGetOK {
	return &StatusSourceGetOK{}
}

/*
StatusSourceGetOK describes a response with status code 200, with default header values.

StatusSourceGetOK status source get o k
*/
type StatusSourceGetOK struct {
	Payload []*models.StatusSource
}

// IsSuccess returns true when this status source get o k response has a 2xx status code
func (o *StatusSourceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this status source get o k response has a 3xx status code
func (o *StatusSourceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status source get o k response has a 4xx status code
func (o *StatusSourceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this status source get o k response has a 5xx status code
func (o *StatusSourceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this status source get o k response a status code equal to that given
func (o *StatusSourceGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the status source get o k response
func (o *StatusSourceGetOK) Code() int {
	return 200
}

func (o *StatusSourceGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/source][%d] statusSourceGetOK %s", 200, payload)
}

func (o *StatusSourceGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/source][%d] statusSourceGetOK %s", 200, payload)
}

func (o *StatusSourceGetOK) GetPayload() []*models.StatusSource {
	return o.Payload
}

func (o *StatusSourceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatusSourceGetBadRequest creates a StatusSourceGetBadRequest with default headers values
func NewStatusSourceGetBadRequest() *StatusSourceGetBadRequest {
	return &StatusSourceGetBadRequest{}
}

/*
StatusSourceGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type StatusSourceGetBadRequest struct {
}

// IsSuccess returns true when this status source get bad request response has a 2xx status code
func (o *StatusSourceGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status source get bad request response has a 3xx status code
func (o *StatusSourceGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status source get bad request response has a 4xx status code
func (o *StatusSourceGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this status source get bad request response has a 5xx status code
func (o *StatusSourceGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this status source get bad request response a status code equal to that given
func (o *StatusSourceGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the status source get bad request response
func (o *StatusSourceGetBadRequest) Code() int {
	return 400
}

func (o *StatusSourceGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/source][%d] statusSourceGetBadRequest", 400)
}

func (o *StatusSourceGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/source][%d] statusSourceGetBadRequest", 400)
}

func (o *StatusSourceGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusSourceGetUnauthorized creates a StatusSourceGetUnauthorized with default headers values
func NewStatusSourceGetUnauthorized() *StatusSourceGetUnauthorized {
	return &StatusSourceGetUnauthorized{}
}

/*
StatusSourceGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type StatusSourceGetUnauthorized struct {
}

// IsSuccess returns true when this status source get unauthorized response has a 2xx status code
func (o *StatusSourceGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status source get unauthorized response has a 3xx status code
func (o *StatusSourceGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status source get unauthorized response has a 4xx status code
func (o *StatusSourceGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this status source get unauthorized response has a 5xx status code
func (o *StatusSourceGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this status source get unauthorized response a status code equal to that given
func (o *StatusSourceGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the status source get unauthorized response
func (o *StatusSourceGetUnauthorized) Code() int {
	return 401
}

func (o *StatusSourceGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/source][%d] statusSourceGetUnauthorized", 401)
}

func (o *StatusSourceGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/source][%d] statusSourceGetUnauthorized", 401)
}

func (o *StatusSourceGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusSourceGetForbidden creates a StatusSourceGetForbidden with default headers values
func NewStatusSourceGetForbidden() *StatusSourceGetForbidden {
	return &StatusSourceGetForbidden{}
}

/*
StatusSourceGetForbidden describes a response with status code 403, with default header values.

forbidden
*/
type StatusSourceGetForbidden struct {
}

// IsSuccess returns true when this status source get forbidden response has a 2xx status code
func (o *StatusSourceGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status source get forbidden response has a 3xx status code
func (o *StatusSourceGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status source get forbidden response has a 4xx status code
func (o *StatusSourceGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this status source get forbidden response has a 5xx status code
func (o *StatusSourceGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this status source get forbidden response a status code equal to that given
func (o *StatusSourceGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the status source get forbidden response
func (o *StatusSourceGetForbidden) Code() int {
	return 403
}

func (o *StatusSourceGetForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/source][%d] statusSourceGetForbidden", 403)
}

func (o *StatusSourceGetForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/source][%d] statusSourceGetForbidden", 403)
}

func (o *StatusSourceGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusSourceGetNotFound creates a StatusSourceGetNotFound with default headers values
func NewStatusSourceGetNotFound() *StatusSourceGetNotFound {
	return &StatusSourceGetNotFound{}
}

/*
StatusSourceGetNotFound describes a response with status code 404, with default header values.

not found
*/
type StatusSourceGetNotFound struct {
}

// IsSuccess returns true when this status source get not found response has a 2xx status code
func (o *StatusSourceGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status source get not found response has a 3xx status code
func (o *StatusSourceGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status source get not found response has a 4xx status code
func (o *StatusSourceGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this status source get not found response has a 5xx status code
func (o *StatusSourceGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this status source get not found response a status code equal to that given
func (o *StatusSourceGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the status source get not found response
func (o *StatusSourceGetNotFound) Code() int {
	return 404
}

func (o *StatusSourceGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/source][%d] statusSourceGetNotFound", 404)
}

func (o *StatusSourceGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/source][%d] statusSourceGetNotFound", 404)
}

func (o *StatusSourceGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusSourceGetNotAcceptable creates a StatusSourceGetNotAcceptable with default headers values
func NewStatusSourceGetNotAcceptable() *StatusSourceGetNotAcceptable {
	return &StatusSourceGetNotAcceptable{}
}

/*
StatusSourceGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type StatusSourceGetNotAcceptable struct {
}

// IsSuccess returns true when this status source get not acceptable response has a 2xx status code
func (o *StatusSourceGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status source get not acceptable response has a 3xx status code
func (o *StatusSourceGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status source get not acceptable response has a 4xx status code
func (o *StatusSourceGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this status source get not acceptable response has a 5xx status code
func (o *StatusSourceGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this status source get not acceptable response a status code equal to that given
func (o *StatusSourceGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the status source get not acceptable response
func (o *StatusSourceGetNotAcceptable) Code() int {
	return 406
}

func (o *StatusSourceGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/source][%d] statusSourceGetNotAcceptable", 406)
}

func (o *StatusSourceGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/source][%d] statusSourceGetNotAcceptable", 406)
}

func (o *StatusSourceGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusSourceGetInternalServerError creates a StatusSourceGetInternalServerError with default headers values
func NewStatusSourceGetInternalServerError() *StatusSourceGetInternalServerError {
	return &StatusSourceGetInternalServerError{}
}

/*
StatusSourceGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type StatusSourceGetInternalServerError struct {
}

// IsSuccess returns true when this status source get internal server error response has a 2xx status code
func (o *StatusSourceGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status source get internal server error response has a 3xx status code
func (o *StatusSourceGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status source get internal server error response has a 4xx status code
func (o *StatusSourceGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this status source get internal server error response has a 5xx status code
func (o *StatusSourceGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this status source get internal server error response a status code equal to that given
func (o *StatusSourceGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the status source get internal server error response
func (o *StatusSourceGetInternalServerError) Code() int {
	return 500
}

func (o *StatusSourceGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/source][%d] statusSourceGetInternalServerError", 500)
}

func (o *StatusSourceGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/source][%d] statusSourceGetInternalServerError", 500)
}

func (o *StatusSourceGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
