// Code generated by go-swagger; DO NOT EDIT.

package filters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VyrCossont/slurp/models"
)

// FilterStatusGetReader is a Reader for the FilterStatusGet structure.
type FilterStatusGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FilterStatusGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFilterStatusGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFilterStatusGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFilterStatusGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFilterStatusGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewFilterStatusGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFilterStatusGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v2/filters/statuses/{id}] filterStatusGet", response, response.Code())
	}
}

// NewFilterStatusGetOK creates a FilterStatusGetOK with default headers values
func NewFilterStatusGetOK() *FilterStatusGetOK {
	return &FilterStatusGetOK{}
}

/*
FilterStatusGetOK describes a response with status code 200, with default header values.

Requested filter status.
*/
type FilterStatusGetOK struct {
	Payload *models.FilterStatus
}

// IsSuccess returns true when this filter status get o k response has a 2xx status code
func (o *FilterStatusGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this filter status get o k response has a 3xx status code
func (o *FilterStatusGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter status get o k response has a 4xx status code
func (o *FilterStatusGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this filter status get o k response has a 5xx status code
func (o *FilterStatusGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this filter status get o k response a status code equal to that given
func (o *FilterStatusGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the filter status get o k response
func (o *FilterStatusGetOK) Code() int {
	return 200
}

func (o *FilterStatusGetOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/filters/statuses/{id}][%d] filterStatusGetOK  %+v", 200, o.Payload)
}

func (o *FilterStatusGetOK) String() string {
	return fmt.Sprintf("[GET /api/v2/filters/statuses/{id}][%d] filterStatusGetOK  %+v", 200, o.Payload)
}

func (o *FilterStatusGetOK) GetPayload() *models.FilterStatus {
	return o.Payload
}

func (o *FilterStatusGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FilterStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFilterStatusGetBadRequest creates a FilterStatusGetBadRequest with default headers values
func NewFilterStatusGetBadRequest() *FilterStatusGetBadRequest {
	return &FilterStatusGetBadRequest{}
}

/*
FilterStatusGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type FilterStatusGetBadRequest struct {
}

// IsSuccess returns true when this filter status get bad request response has a 2xx status code
func (o *FilterStatusGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter status get bad request response has a 3xx status code
func (o *FilterStatusGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter status get bad request response has a 4xx status code
func (o *FilterStatusGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter status get bad request response has a 5xx status code
func (o *FilterStatusGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this filter status get bad request response a status code equal to that given
func (o *FilterStatusGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the filter status get bad request response
func (o *FilterStatusGetBadRequest) Code() int {
	return 400
}

func (o *FilterStatusGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/filters/statuses/{id}][%d] filterStatusGetBadRequest ", 400)
}

func (o *FilterStatusGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/filters/statuses/{id}][%d] filterStatusGetBadRequest ", 400)
}

func (o *FilterStatusGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterStatusGetUnauthorized creates a FilterStatusGetUnauthorized with default headers values
func NewFilterStatusGetUnauthorized() *FilterStatusGetUnauthorized {
	return &FilterStatusGetUnauthorized{}
}

/*
FilterStatusGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type FilterStatusGetUnauthorized struct {
}

// IsSuccess returns true when this filter status get unauthorized response has a 2xx status code
func (o *FilterStatusGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter status get unauthorized response has a 3xx status code
func (o *FilterStatusGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter status get unauthorized response has a 4xx status code
func (o *FilterStatusGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter status get unauthorized response has a 5xx status code
func (o *FilterStatusGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this filter status get unauthorized response a status code equal to that given
func (o *FilterStatusGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the filter status get unauthorized response
func (o *FilterStatusGetUnauthorized) Code() int {
	return 401
}

func (o *FilterStatusGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/filters/statuses/{id}][%d] filterStatusGetUnauthorized ", 401)
}

func (o *FilterStatusGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/filters/statuses/{id}][%d] filterStatusGetUnauthorized ", 401)
}

func (o *FilterStatusGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterStatusGetNotFound creates a FilterStatusGetNotFound with default headers values
func NewFilterStatusGetNotFound() *FilterStatusGetNotFound {
	return &FilterStatusGetNotFound{}
}

/*
FilterStatusGetNotFound describes a response with status code 404, with default header values.

not found
*/
type FilterStatusGetNotFound struct {
}

// IsSuccess returns true when this filter status get not found response has a 2xx status code
func (o *FilterStatusGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter status get not found response has a 3xx status code
func (o *FilterStatusGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter status get not found response has a 4xx status code
func (o *FilterStatusGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter status get not found response has a 5xx status code
func (o *FilterStatusGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this filter status get not found response a status code equal to that given
func (o *FilterStatusGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the filter status get not found response
func (o *FilterStatusGetNotFound) Code() int {
	return 404
}

func (o *FilterStatusGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/filters/statuses/{id}][%d] filterStatusGetNotFound ", 404)
}

func (o *FilterStatusGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/filters/statuses/{id}][%d] filterStatusGetNotFound ", 404)
}

func (o *FilterStatusGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterStatusGetNotAcceptable creates a FilterStatusGetNotAcceptable with default headers values
func NewFilterStatusGetNotAcceptable() *FilterStatusGetNotAcceptable {
	return &FilterStatusGetNotAcceptable{}
}

/*
FilterStatusGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type FilterStatusGetNotAcceptable struct {
}

// IsSuccess returns true when this filter status get not acceptable response has a 2xx status code
func (o *FilterStatusGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter status get not acceptable response has a 3xx status code
func (o *FilterStatusGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter status get not acceptable response has a 4xx status code
func (o *FilterStatusGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter status get not acceptable response has a 5xx status code
func (o *FilterStatusGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this filter status get not acceptable response a status code equal to that given
func (o *FilterStatusGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the filter status get not acceptable response
func (o *FilterStatusGetNotAcceptable) Code() int {
	return 406
}

func (o *FilterStatusGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v2/filters/statuses/{id}][%d] filterStatusGetNotAcceptable ", 406)
}

func (o *FilterStatusGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v2/filters/statuses/{id}][%d] filterStatusGetNotAcceptable ", 406)
}

func (o *FilterStatusGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterStatusGetInternalServerError creates a FilterStatusGetInternalServerError with default headers values
func NewFilterStatusGetInternalServerError() *FilterStatusGetInternalServerError {
	return &FilterStatusGetInternalServerError{}
}

/*
FilterStatusGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type FilterStatusGetInternalServerError struct {
}

// IsSuccess returns true when this filter status get internal server error response has a 2xx status code
func (o *FilterStatusGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter status get internal server error response has a 3xx status code
func (o *FilterStatusGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter status get internal server error response has a 4xx status code
func (o *FilterStatusGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this filter status get internal server error response has a 5xx status code
func (o *FilterStatusGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this filter status get internal server error response a status code equal to that given
func (o *FilterStatusGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the filter status get internal server error response
func (o *FilterStatusGetInternalServerError) Code() int {
	return 500
}

func (o *FilterStatusGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/filters/statuses/{id}][%d] filterStatusGetInternalServerError ", 500)
}

func (o *FilterStatusGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/filters/statuses/{id}][%d] filterStatusGetInternalServerError ", 500)
}

func (o *FilterStatusGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
