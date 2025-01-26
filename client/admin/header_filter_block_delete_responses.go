// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// HeaderFilterBlockDeleteReader is a Reader for the HeaderFilterBlockDelete structure.
type HeaderFilterBlockDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeaderFilterBlockDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewHeaderFilterBlockDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHeaderFilterBlockDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewHeaderFilterBlockDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHeaderFilterBlockDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHeaderFilterBlockDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHeaderFilterBlockDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /api/v1/admin/header_blocks/{id}] headerFilterBlockDelete", response, response.Code())
	}
}

// NewHeaderFilterBlockDeleteAccepted creates a HeaderFilterBlockDeleteAccepted with default headers values
func NewHeaderFilterBlockDeleteAccepted() *HeaderFilterBlockDeleteAccepted {
	return &HeaderFilterBlockDeleteAccepted{}
}

/*
HeaderFilterBlockDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type HeaderFilterBlockDeleteAccepted struct {
}

// IsSuccess returns true when this header filter block delete accepted response has a 2xx status code
func (o *HeaderFilterBlockDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this header filter block delete accepted response has a 3xx status code
func (o *HeaderFilterBlockDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter block delete accepted response has a 4xx status code
func (o *HeaderFilterBlockDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this header filter block delete accepted response has a 5xx status code
func (o *HeaderFilterBlockDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter block delete accepted response a status code equal to that given
func (o *HeaderFilterBlockDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the header filter block delete accepted response
func (o *HeaderFilterBlockDeleteAccepted) Code() int {
	return 202
}

func (o *HeaderFilterBlockDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockDeleteAccepted", 202)
}

func (o *HeaderFilterBlockDeleteAccepted) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockDeleteAccepted", 202)
}

func (o *HeaderFilterBlockDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeaderFilterBlockDeleteBadRequest creates a HeaderFilterBlockDeleteBadRequest with default headers values
func NewHeaderFilterBlockDeleteBadRequest() *HeaderFilterBlockDeleteBadRequest {
	return &HeaderFilterBlockDeleteBadRequest{}
}

/*
HeaderFilterBlockDeleteBadRequest describes a response with status code 400, with default header values.

bad request
*/
type HeaderFilterBlockDeleteBadRequest struct {
}

// IsSuccess returns true when this header filter block delete bad request response has a 2xx status code
func (o *HeaderFilterBlockDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter block delete bad request response has a 3xx status code
func (o *HeaderFilterBlockDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter block delete bad request response has a 4xx status code
func (o *HeaderFilterBlockDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this header filter block delete bad request response has a 5xx status code
func (o *HeaderFilterBlockDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter block delete bad request response a status code equal to that given
func (o *HeaderFilterBlockDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the header filter block delete bad request response
func (o *HeaderFilterBlockDeleteBadRequest) Code() int {
	return 400
}

func (o *HeaderFilterBlockDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockDeleteBadRequest", 400)
}

func (o *HeaderFilterBlockDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockDeleteBadRequest", 400)
}

func (o *HeaderFilterBlockDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeaderFilterBlockDeleteUnauthorized creates a HeaderFilterBlockDeleteUnauthorized with default headers values
func NewHeaderFilterBlockDeleteUnauthorized() *HeaderFilterBlockDeleteUnauthorized {
	return &HeaderFilterBlockDeleteUnauthorized{}
}

/*
HeaderFilterBlockDeleteUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type HeaderFilterBlockDeleteUnauthorized struct {
}

// IsSuccess returns true when this header filter block delete unauthorized response has a 2xx status code
func (o *HeaderFilterBlockDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter block delete unauthorized response has a 3xx status code
func (o *HeaderFilterBlockDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter block delete unauthorized response has a 4xx status code
func (o *HeaderFilterBlockDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this header filter block delete unauthorized response has a 5xx status code
func (o *HeaderFilterBlockDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter block delete unauthorized response a status code equal to that given
func (o *HeaderFilterBlockDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the header filter block delete unauthorized response
func (o *HeaderFilterBlockDeleteUnauthorized) Code() int {
	return 401
}

func (o *HeaderFilterBlockDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockDeleteUnauthorized", 401)
}

func (o *HeaderFilterBlockDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockDeleteUnauthorized", 401)
}

func (o *HeaderFilterBlockDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeaderFilterBlockDeleteForbidden creates a HeaderFilterBlockDeleteForbidden with default headers values
func NewHeaderFilterBlockDeleteForbidden() *HeaderFilterBlockDeleteForbidden {
	return &HeaderFilterBlockDeleteForbidden{}
}

/*
HeaderFilterBlockDeleteForbidden describes a response with status code 403, with default header values.

forbidden
*/
type HeaderFilterBlockDeleteForbidden struct {
}

// IsSuccess returns true when this header filter block delete forbidden response has a 2xx status code
func (o *HeaderFilterBlockDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter block delete forbidden response has a 3xx status code
func (o *HeaderFilterBlockDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter block delete forbidden response has a 4xx status code
func (o *HeaderFilterBlockDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this header filter block delete forbidden response has a 5xx status code
func (o *HeaderFilterBlockDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter block delete forbidden response a status code equal to that given
func (o *HeaderFilterBlockDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the header filter block delete forbidden response
func (o *HeaderFilterBlockDeleteForbidden) Code() int {
	return 403
}

func (o *HeaderFilterBlockDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockDeleteForbidden", 403)
}

func (o *HeaderFilterBlockDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockDeleteForbidden", 403)
}

func (o *HeaderFilterBlockDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeaderFilterBlockDeleteNotFound creates a HeaderFilterBlockDeleteNotFound with default headers values
func NewHeaderFilterBlockDeleteNotFound() *HeaderFilterBlockDeleteNotFound {
	return &HeaderFilterBlockDeleteNotFound{}
}

/*
HeaderFilterBlockDeleteNotFound describes a response with status code 404, with default header values.

not found
*/
type HeaderFilterBlockDeleteNotFound struct {
}

// IsSuccess returns true when this header filter block delete not found response has a 2xx status code
func (o *HeaderFilterBlockDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter block delete not found response has a 3xx status code
func (o *HeaderFilterBlockDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter block delete not found response has a 4xx status code
func (o *HeaderFilterBlockDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this header filter block delete not found response has a 5xx status code
func (o *HeaderFilterBlockDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter block delete not found response a status code equal to that given
func (o *HeaderFilterBlockDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the header filter block delete not found response
func (o *HeaderFilterBlockDeleteNotFound) Code() int {
	return 404
}

func (o *HeaderFilterBlockDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockDeleteNotFound", 404)
}

func (o *HeaderFilterBlockDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockDeleteNotFound", 404)
}

func (o *HeaderFilterBlockDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeaderFilterBlockDeleteInternalServerError creates a HeaderFilterBlockDeleteInternalServerError with default headers values
func NewHeaderFilterBlockDeleteInternalServerError() *HeaderFilterBlockDeleteInternalServerError {
	return &HeaderFilterBlockDeleteInternalServerError{}
}

/*
HeaderFilterBlockDeleteInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type HeaderFilterBlockDeleteInternalServerError struct {
}

// IsSuccess returns true when this header filter block delete internal server error response has a 2xx status code
func (o *HeaderFilterBlockDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter block delete internal server error response has a 3xx status code
func (o *HeaderFilterBlockDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter block delete internal server error response has a 4xx status code
func (o *HeaderFilterBlockDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this header filter block delete internal server error response has a 5xx status code
func (o *HeaderFilterBlockDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this header filter block delete internal server error response a status code equal to that given
func (o *HeaderFilterBlockDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the header filter block delete internal server error response
func (o *HeaderFilterBlockDeleteInternalServerError) Code() int {
	return 500
}

func (o *HeaderFilterBlockDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockDeleteInternalServerError", 500)
}

func (o *HeaderFilterBlockDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockDeleteInternalServerError", 500)
}

func (o *HeaderFilterBlockDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
