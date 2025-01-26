// Code generated by go-swagger; DO NOT EDIT.

package filters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// FilterV2DeleteReader is a Reader for the FilterV2Delete structure.
type FilterV2DeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FilterV2DeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFilterV2DeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFilterV2DeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFilterV2DeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFilterV2DeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewFilterV2DeleteNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFilterV2DeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /api/v2/filters/{id}] filterV2Delete", response, response.Code())
	}
}

// NewFilterV2DeleteOK creates a FilterV2DeleteOK with default headers values
func NewFilterV2DeleteOK() *FilterV2DeleteOK {
	return &FilterV2DeleteOK{}
}

/*
FilterV2DeleteOK describes a response with status code 200, with default header values.

filter deleted
*/
type FilterV2DeleteOK struct {
}

// IsSuccess returns true when this filter v2 delete o k response has a 2xx status code
func (o *FilterV2DeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this filter v2 delete o k response has a 3xx status code
func (o *FilterV2DeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v2 delete o k response has a 4xx status code
func (o *FilterV2DeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this filter v2 delete o k response has a 5xx status code
func (o *FilterV2DeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this filter v2 delete o k response a status code equal to that given
func (o *FilterV2DeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the filter v2 delete o k response
func (o *FilterV2DeleteOK) Code() int {
	return 200
}

func (o *FilterV2DeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/{id}][%d] filterV2DeleteOK", 200)
}

func (o *FilterV2DeleteOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/{id}][%d] filterV2DeleteOK", 200)
}

func (o *FilterV2DeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterV2DeleteBadRequest creates a FilterV2DeleteBadRequest with default headers values
func NewFilterV2DeleteBadRequest() *FilterV2DeleteBadRequest {
	return &FilterV2DeleteBadRequest{}
}

/*
FilterV2DeleteBadRequest describes a response with status code 400, with default header values.

bad request
*/
type FilterV2DeleteBadRequest struct {
}

// IsSuccess returns true when this filter v2 delete bad request response has a 2xx status code
func (o *FilterV2DeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter v2 delete bad request response has a 3xx status code
func (o *FilterV2DeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v2 delete bad request response has a 4xx status code
func (o *FilterV2DeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter v2 delete bad request response has a 5xx status code
func (o *FilterV2DeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this filter v2 delete bad request response a status code equal to that given
func (o *FilterV2DeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the filter v2 delete bad request response
func (o *FilterV2DeleteBadRequest) Code() int {
	return 400
}

func (o *FilterV2DeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/{id}][%d] filterV2DeleteBadRequest", 400)
}

func (o *FilterV2DeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/{id}][%d] filterV2DeleteBadRequest", 400)
}

func (o *FilterV2DeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterV2DeleteUnauthorized creates a FilterV2DeleteUnauthorized with default headers values
func NewFilterV2DeleteUnauthorized() *FilterV2DeleteUnauthorized {
	return &FilterV2DeleteUnauthorized{}
}

/*
FilterV2DeleteUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type FilterV2DeleteUnauthorized struct {
}

// IsSuccess returns true when this filter v2 delete unauthorized response has a 2xx status code
func (o *FilterV2DeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter v2 delete unauthorized response has a 3xx status code
func (o *FilterV2DeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v2 delete unauthorized response has a 4xx status code
func (o *FilterV2DeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter v2 delete unauthorized response has a 5xx status code
func (o *FilterV2DeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this filter v2 delete unauthorized response a status code equal to that given
func (o *FilterV2DeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the filter v2 delete unauthorized response
func (o *FilterV2DeleteUnauthorized) Code() int {
	return 401
}

func (o *FilterV2DeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/{id}][%d] filterV2DeleteUnauthorized", 401)
}

func (o *FilterV2DeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/{id}][%d] filterV2DeleteUnauthorized", 401)
}

func (o *FilterV2DeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterV2DeleteNotFound creates a FilterV2DeleteNotFound with default headers values
func NewFilterV2DeleteNotFound() *FilterV2DeleteNotFound {
	return &FilterV2DeleteNotFound{}
}

/*
FilterV2DeleteNotFound describes a response with status code 404, with default header values.

not found
*/
type FilterV2DeleteNotFound struct {
}

// IsSuccess returns true when this filter v2 delete not found response has a 2xx status code
func (o *FilterV2DeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter v2 delete not found response has a 3xx status code
func (o *FilterV2DeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v2 delete not found response has a 4xx status code
func (o *FilterV2DeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter v2 delete not found response has a 5xx status code
func (o *FilterV2DeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this filter v2 delete not found response a status code equal to that given
func (o *FilterV2DeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the filter v2 delete not found response
func (o *FilterV2DeleteNotFound) Code() int {
	return 404
}

func (o *FilterV2DeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/{id}][%d] filterV2DeleteNotFound", 404)
}

func (o *FilterV2DeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/{id}][%d] filterV2DeleteNotFound", 404)
}

func (o *FilterV2DeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterV2DeleteNotAcceptable creates a FilterV2DeleteNotAcceptable with default headers values
func NewFilterV2DeleteNotAcceptable() *FilterV2DeleteNotAcceptable {
	return &FilterV2DeleteNotAcceptable{}
}

/*
FilterV2DeleteNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type FilterV2DeleteNotAcceptable struct {
}

// IsSuccess returns true when this filter v2 delete not acceptable response has a 2xx status code
func (o *FilterV2DeleteNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter v2 delete not acceptable response has a 3xx status code
func (o *FilterV2DeleteNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v2 delete not acceptable response has a 4xx status code
func (o *FilterV2DeleteNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter v2 delete not acceptable response has a 5xx status code
func (o *FilterV2DeleteNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this filter v2 delete not acceptable response a status code equal to that given
func (o *FilterV2DeleteNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the filter v2 delete not acceptable response
func (o *FilterV2DeleteNotAcceptable) Code() int {
	return 406
}

func (o *FilterV2DeleteNotAcceptable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/{id}][%d] filterV2DeleteNotAcceptable", 406)
}

func (o *FilterV2DeleteNotAcceptable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/{id}][%d] filterV2DeleteNotAcceptable", 406)
}

func (o *FilterV2DeleteNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterV2DeleteInternalServerError creates a FilterV2DeleteInternalServerError with default headers values
func NewFilterV2DeleteInternalServerError() *FilterV2DeleteInternalServerError {
	return &FilterV2DeleteInternalServerError{}
}

/*
FilterV2DeleteInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type FilterV2DeleteInternalServerError struct {
}

// IsSuccess returns true when this filter v2 delete internal server error response has a 2xx status code
func (o *FilterV2DeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter v2 delete internal server error response has a 3xx status code
func (o *FilterV2DeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v2 delete internal server error response has a 4xx status code
func (o *FilterV2DeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this filter v2 delete internal server error response has a 5xx status code
func (o *FilterV2DeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this filter v2 delete internal server error response a status code equal to that given
func (o *FilterV2DeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the filter v2 delete internal server error response
func (o *FilterV2DeleteInternalServerError) Code() int {
	return 500
}

func (o *FilterV2DeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/{id}][%d] filterV2DeleteInternalServerError", 500)
}

func (o *FilterV2DeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/{id}][%d] filterV2DeleteInternalServerError", 500)
}

func (o *FilterV2DeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
