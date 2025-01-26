// Code generated by go-swagger; DO NOT EDIT.

package filters

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

// FilterStatusesGetReader is a Reader for the FilterStatusesGet structure.
type FilterStatusesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FilterStatusesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFilterStatusesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFilterStatusesGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFilterStatusesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFilterStatusesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewFilterStatusesGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFilterStatusesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v2/filters/{id}/statuses] filterStatusesGet", response, response.Code())
	}
}

// NewFilterStatusesGetOK creates a FilterStatusesGetOK with default headers values
func NewFilterStatusesGetOK() *FilterStatusesGetOK {
	return &FilterStatusesGetOK{}
}

/*
FilterStatusesGetOK describes a response with status code 200, with default header values.

Requested filter statuses.
*/
type FilterStatusesGetOK struct {
	Payload []*models.FilterStatus
}

// IsSuccess returns true when this filter statuses get o k response has a 2xx status code
func (o *FilterStatusesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this filter statuses get o k response has a 3xx status code
func (o *FilterStatusesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter statuses get o k response has a 4xx status code
func (o *FilterStatusesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this filter statuses get o k response has a 5xx status code
func (o *FilterStatusesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this filter statuses get o k response a status code equal to that given
func (o *FilterStatusesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the filter statuses get o k response
func (o *FilterStatusesGetOK) Code() int {
	return 200
}

func (o *FilterStatusesGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/filters/{id}/statuses][%d] filterStatusesGetOK %s", 200, payload)
}

func (o *FilterStatusesGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/filters/{id}/statuses][%d] filterStatusesGetOK %s", 200, payload)
}

func (o *FilterStatusesGetOK) GetPayload() []*models.FilterStatus {
	return o.Payload
}

func (o *FilterStatusesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFilterStatusesGetBadRequest creates a FilterStatusesGetBadRequest with default headers values
func NewFilterStatusesGetBadRequest() *FilterStatusesGetBadRequest {
	return &FilterStatusesGetBadRequest{}
}

/*
FilterStatusesGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type FilterStatusesGetBadRequest struct {
}

// IsSuccess returns true when this filter statuses get bad request response has a 2xx status code
func (o *FilterStatusesGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter statuses get bad request response has a 3xx status code
func (o *FilterStatusesGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter statuses get bad request response has a 4xx status code
func (o *FilterStatusesGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter statuses get bad request response has a 5xx status code
func (o *FilterStatusesGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this filter statuses get bad request response a status code equal to that given
func (o *FilterStatusesGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the filter statuses get bad request response
func (o *FilterStatusesGetBadRequest) Code() int {
	return 400
}

func (o *FilterStatusesGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/filters/{id}/statuses][%d] filterStatusesGetBadRequest", 400)
}

func (o *FilterStatusesGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/filters/{id}/statuses][%d] filterStatusesGetBadRequest", 400)
}

func (o *FilterStatusesGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterStatusesGetUnauthorized creates a FilterStatusesGetUnauthorized with default headers values
func NewFilterStatusesGetUnauthorized() *FilterStatusesGetUnauthorized {
	return &FilterStatusesGetUnauthorized{}
}

/*
FilterStatusesGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type FilterStatusesGetUnauthorized struct {
}

// IsSuccess returns true when this filter statuses get unauthorized response has a 2xx status code
func (o *FilterStatusesGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter statuses get unauthorized response has a 3xx status code
func (o *FilterStatusesGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter statuses get unauthorized response has a 4xx status code
func (o *FilterStatusesGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter statuses get unauthorized response has a 5xx status code
func (o *FilterStatusesGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this filter statuses get unauthorized response a status code equal to that given
func (o *FilterStatusesGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the filter statuses get unauthorized response
func (o *FilterStatusesGetUnauthorized) Code() int {
	return 401
}

func (o *FilterStatusesGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/filters/{id}/statuses][%d] filterStatusesGetUnauthorized", 401)
}

func (o *FilterStatusesGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/filters/{id}/statuses][%d] filterStatusesGetUnauthorized", 401)
}

func (o *FilterStatusesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterStatusesGetNotFound creates a FilterStatusesGetNotFound with default headers values
func NewFilterStatusesGetNotFound() *FilterStatusesGetNotFound {
	return &FilterStatusesGetNotFound{}
}

/*
FilterStatusesGetNotFound describes a response with status code 404, with default header values.

not found
*/
type FilterStatusesGetNotFound struct {
}

// IsSuccess returns true when this filter statuses get not found response has a 2xx status code
func (o *FilterStatusesGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter statuses get not found response has a 3xx status code
func (o *FilterStatusesGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter statuses get not found response has a 4xx status code
func (o *FilterStatusesGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter statuses get not found response has a 5xx status code
func (o *FilterStatusesGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this filter statuses get not found response a status code equal to that given
func (o *FilterStatusesGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the filter statuses get not found response
func (o *FilterStatusesGetNotFound) Code() int {
	return 404
}

func (o *FilterStatusesGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/filters/{id}/statuses][%d] filterStatusesGetNotFound", 404)
}

func (o *FilterStatusesGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/filters/{id}/statuses][%d] filterStatusesGetNotFound", 404)
}

func (o *FilterStatusesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterStatusesGetNotAcceptable creates a FilterStatusesGetNotAcceptable with default headers values
func NewFilterStatusesGetNotAcceptable() *FilterStatusesGetNotAcceptable {
	return &FilterStatusesGetNotAcceptable{}
}

/*
FilterStatusesGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type FilterStatusesGetNotAcceptable struct {
}

// IsSuccess returns true when this filter statuses get not acceptable response has a 2xx status code
func (o *FilterStatusesGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter statuses get not acceptable response has a 3xx status code
func (o *FilterStatusesGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter statuses get not acceptable response has a 4xx status code
func (o *FilterStatusesGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter statuses get not acceptable response has a 5xx status code
func (o *FilterStatusesGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this filter statuses get not acceptable response a status code equal to that given
func (o *FilterStatusesGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the filter statuses get not acceptable response
func (o *FilterStatusesGetNotAcceptable) Code() int {
	return 406
}

func (o *FilterStatusesGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v2/filters/{id}/statuses][%d] filterStatusesGetNotAcceptable", 406)
}

func (o *FilterStatusesGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v2/filters/{id}/statuses][%d] filterStatusesGetNotAcceptable", 406)
}

func (o *FilterStatusesGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterStatusesGetInternalServerError creates a FilterStatusesGetInternalServerError with default headers values
func NewFilterStatusesGetInternalServerError() *FilterStatusesGetInternalServerError {
	return &FilterStatusesGetInternalServerError{}
}

/*
FilterStatusesGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type FilterStatusesGetInternalServerError struct {
}

// IsSuccess returns true when this filter statuses get internal server error response has a 2xx status code
func (o *FilterStatusesGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter statuses get internal server error response has a 3xx status code
func (o *FilterStatusesGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter statuses get internal server error response has a 4xx status code
func (o *FilterStatusesGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this filter statuses get internal server error response has a 5xx status code
func (o *FilterStatusesGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this filter statuses get internal server error response a status code equal to that given
func (o *FilterStatusesGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the filter statuses get internal server error response
func (o *FilterStatusesGetInternalServerError) Code() int {
	return 500
}

func (o *FilterStatusesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/filters/{id}/statuses][%d] filterStatusesGetInternalServerError", 500)
}

func (o *FilterStatusesGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/filters/{id}/statuses][%d] filterStatusesGetInternalServerError", 500)
}

func (o *FilterStatusesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
