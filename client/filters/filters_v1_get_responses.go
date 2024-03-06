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

// FiltersV1GetReader is a Reader for the FiltersV1Get structure.
type FiltersV1GetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FiltersV1GetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFiltersV1GetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFiltersV1GetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFiltersV1GetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFiltersV1GetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewFiltersV1GetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFiltersV1GetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/filters] filtersV1Get", response, response.Code())
	}
}

// NewFiltersV1GetOK creates a FiltersV1GetOK with default headers values
func NewFiltersV1GetOK() *FiltersV1GetOK {
	return &FiltersV1GetOK{}
}

/*
FiltersV1GetOK describes a response with status code 200, with default header values.

Requested filters.
*/
type FiltersV1GetOK struct {
	Payload []*models.FilterV1
}

// IsSuccess returns true when this filters v1 get o k response has a 2xx status code
func (o *FiltersV1GetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this filters v1 get o k response has a 3xx status code
func (o *FiltersV1GetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filters v1 get o k response has a 4xx status code
func (o *FiltersV1GetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this filters v1 get o k response has a 5xx status code
func (o *FiltersV1GetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this filters v1 get o k response a status code equal to that given
func (o *FiltersV1GetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the filters v1 get o k response
func (o *FiltersV1GetOK) Code() int {
	return 200
}

func (o *FiltersV1GetOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/filters][%d] filtersV1GetOK  %+v", 200, o.Payload)
}

func (o *FiltersV1GetOK) String() string {
	return fmt.Sprintf("[GET /api/v1/filters][%d] filtersV1GetOK  %+v", 200, o.Payload)
}

func (o *FiltersV1GetOK) GetPayload() []*models.FilterV1 {
	return o.Payload
}

func (o *FiltersV1GetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFiltersV1GetBadRequest creates a FiltersV1GetBadRequest with default headers values
func NewFiltersV1GetBadRequest() *FiltersV1GetBadRequest {
	return &FiltersV1GetBadRequest{}
}

/*
FiltersV1GetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type FiltersV1GetBadRequest struct {
}

// IsSuccess returns true when this filters v1 get bad request response has a 2xx status code
func (o *FiltersV1GetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filters v1 get bad request response has a 3xx status code
func (o *FiltersV1GetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filters v1 get bad request response has a 4xx status code
func (o *FiltersV1GetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this filters v1 get bad request response has a 5xx status code
func (o *FiltersV1GetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this filters v1 get bad request response a status code equal to that given
func (o *FiltersV1GetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the filters v1 get bad request response
func (o *FiltersV1GetBadRequest) Code() int {
	return 400
}

func (o *FiltersV1GetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/filters][%d] filtersV1GetBadRequest ", 400)
}

func (o *FiltersV1GetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/filters][%d] filtersV1GetBadRequest ", 400)
}

func (o *FiltersV1GetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFiltersV1GetUnauthorized creates a FiltersV1GetUnauthorized with default headers values
func NewFiltersV1GetUnauthorized() *FiltersV1GetUnauthorized {
	return &FiltersV1GetUnauthorized{}
}

/*
FiltersV1GetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type FiltersV1GetUnauthorized struct {
}

// IsSuccess returns true when this filters v1 get unauthorized response has a 2xx status code
func (o *FiltersV1GetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filters v1 get unauthorized response has a 3xx status code
func (o *FiltersV1GetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filters v1 get unauthorized response has a 4xx status code
func (o *FiltersV1GetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this filters v1 get unauthorized response has a 5xx status code
func (o *FiltersV1GetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this filters v1 get unauthorized response a status code equal to that given
func (o *FiltersV1GetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the filters v1 get unauthorized response
func (o *FiltersV1GetUnauthorized) Code() int {
	return 401
}

func (o *FiltersV1GetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/filters][%d] filtersV1GetUnauthorized ", 401)
}

func (o *FiltersV1GetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/filters][%d] filtersV1GetUnauthorized ", 401)
}

func (o *FiltersV1GetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFiltersV1GetNotFound creates a FiltersV1GetNotFound with default headers values
func NewFiltersV1GetNotFound() *FiltersV1GetNotFound {
	return &FiltersV1GetNotFound{}
}

/*
FiltersV1GetNotFound describes a response with status code 404, with default header values.

not found
*/
type FiltersV1GetNotFound struct {
}

// IsSuccess returns true when this filters v1 get not found response has a 2xx status code
func (o *FiltersV1GetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filters v1 get not found response has a 3xx status code
func (o *FiltersV1GetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filters v1 get not found response has a 4xx status code
func (o *FiltersV1GetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this filters v1 get not found response has a 5xx status code
func (o *FiltersV1GetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this filters v1 get not found response a status code equal to that given
func (o *FiltersV1GetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the filters v1 get not found response
func (o *FiltersV1GetNotFound) Code() int {
	return 404
}

func (o *FiltersV1GetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/filters][%d] filtersV1GetNotFound ", 404)
}

func (o *FiltersV1GetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/filters][%d] filtersV1GetNotFound ", 404)
}

func (o *FiltersV1GetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFiltersV1GetNotAcceptable creates a FiltersV1GetNotAcceptable with default headers values
func NewFiltersV1GetNotAcceptable() *FiltersV1GetNotAcceptable {
	return &FiltersV1GetNotAcceptable{}
}

/*
FiltersV1GetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type FiltersV1GetNotAcceptable struct {
}

// IsSuccess returns true when this filters v1 get not acceptable response has a 2xx status code
func (o *FiltersV1GetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filters v1 get not acceptable response has a 3xx status code
func (o *FiltersV1GetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filters v1 get not acceptable response has a 4xx status code
func (o *FiltersV1GetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this filters v1 get not acceptable response has a 5xx status code
func (o *FiltersV1GetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this filters v1 get not acceptable response a status code equal to that given
func (o *FiltersV1GetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the filters v1 get not acceptable response
func (o *FiltersV1GetNotAcceptable) Code() int {
	return 406
}

func (o *FiltersV1GetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/filters][%d] filtersV1GetNotAcceptable ", 406)
}

func (o *FiltersV1GetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/filters][%d] filtersV1GetNotAcceptable ", 406)
}

func (o *FiltersV1GetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFiltersV1GetInternalServerError creates a FiltersV1GetInternalServerError with default headers values
func NewFiltersV1GetInternalServerError() *FiltersV1GetInternalServerError {
	return &FiltersV1GetInternalServerError{}
}

/*
FiltersV1GetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type FiltersV1GetInternalServerError struct {
}

// IsSuccess returns true when this filters v1 get internal server error response has a 2xx status code
func (o *FiltersV1GetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filters v1 get internal server error response has a 3xx status code
func (o *FiltersV1GetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filters v1 get internal server error response has a 4xx status code
func (o *FiltersV1GetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this filters v1 get internal server error response has a 5xx status code
func (o *FiltersV1GetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this filters v1 get internal server error response a status code equal to that given
func (o *FiltersV1GetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the filters v1 get internal server error response
func (o *FiltersV1GetInternalServerError) Code() int {
	return 500
}

func (o *FiltersV1GetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/filters][%d] filtersV1GetInternalServerError ", 500)
}

func (o *FiltersV1GetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/filters][%d] filtersV1GetInternalServerError ", 500)
}

func (o *FiltersV1GetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
