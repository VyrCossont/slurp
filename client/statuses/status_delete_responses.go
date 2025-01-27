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

// StatusDeleteReader is a Reader for the StatusDelete structure.
type StatusDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatusDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStatusDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStatusDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStatusDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStatusDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewStatusDeleteNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStatusDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /api/v1/statuses/{id}] statusDelete", response, response.Code())
	}
}

// NewStatusDeleteOK creates a StatusDeleteOK with default headers values
func NewStatusDeleteOK() *StatusDeleteOK {
	return &StatusDeleteOK{}
}

/*
StatusDeleteOK describes a response with status code 200, with default header values.

The status that was just deleted.
*/
type StatusDeleteOK struct {
	Payload *models.Status
}

// IsSuccess returns true when this status delete o k response has a 2xx status code
func (o *StatusDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this status delete o k response has a 3xx status code
func (o *StatusDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status delete o k response has a 4xx status code
func (o *StatusDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this status delete o k response has a 5xx status code
func (o *StatusDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this status delete o k response a status code equal to that given
func (o *StatusDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the status delete o k response
func (o *StatusDeleteOK) Code() int {
	return 200
}

func (o *StatusDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1/statuses/{id}][%d] statusDeleteOK %s", 200, payload)
}

func (o *StatusDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1/statuses/{id}][%d] statusDeleteOK %s", 200, payload)
}

func (o *StatusDeleteOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *StatusDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatusDeleteBadRequest creates a StatusDeleteBadRequest with default headers values
func NewStatusDeleteBadRequest() *StatusDeleteBadRequest {
	return &StatusDeleteBadRequest{}
}

/*
StatusDeleteBadRequest describes a response with status code 400, with default header values.

bad request
*/
type StatusDeleteBadRequest struct {
}

// IsSuccess returns true when this status delete bad request response has a 2xx status code
func (o *StatusDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status delete bad request response has a 3xx status code
func (o *StatusDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status delete bad request response has a 4xx status code
func (o *StatusDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this status delete bad request response has a 5xx status code
func (o *StatusDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this status delete bad request response a status code equal to that given
func (o *StatusDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the status delete bad request response
func (o *StatusDeleteBadRequest) Code() int {
	return 400
}

func (o *StatusDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/statuses/{id}][%d] statusDeleteBadRequest", 400)
}

func (o *StatusDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v1/statuses/{id}][%d] statusDeleteBadRequest", 400)
}

func (o *StatusDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusDeleteUnauthorized creates a StatusDeleteUnauthorized with default headers values
func NewStatusDeleteUnauthorized() *StatusDeleteUnauthorized {
	return &StatusDeleteUnauthorized{}
}

/*
StatusDeleteUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type StatusDeleteUnauthorized struct {
}

// IsSuccess returns true when this status delete unauthorized response has a 2xx status code
func (o *StatusDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status delete unauthorized response has a 3xx status code
func (o *StatusDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status delete unauthorized response has a 4xx status code
func (o *StatusDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this status delete unauthorized response has a 5xx status code
func (o *StatusDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this status delete unauthorized response a status code equal to that given
func (o *StatusDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the status delete unauthorized response
func (o *StatusDeleteUnauthorized) Code() int {
	return 401
}

func (o *StatusDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/statuses/{id}][%d] statusDeleteUnauthorized", 401)
}

func (o *StatusDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v1/statuses/{id}][%d] statusDeleteUnauthorized", 401)
}

func (o *StatusDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusDeleteForbidden creates a StatusDeleteForbidden with default headers values
func NewStatusDeleteForbidden() *StatusDeleteForbidden {
	return &StatusDeleteForbidden{}
}

/*
StatusDeleteForbidden describes a response with status code 403, with default header values.

forbidden
*/
type StatusDeleteForbidden struct {
}

// IsSuccess returns true when this status delete forbidden response has a 2xx status code
func (o *StatusDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status delete forbidden response has a 3xx status code
func (o *StatusDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status delete forbidden response has a 4xx status code
func (o *StatusDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this status delete forbidden response has a 5xx status code
func (o *StatusDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this status delete forbidden response a status code equal to that given
func (o *StatusDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the status delete forbidden response
func (o *StatusDeleteForbidden) Code() int {
	return 403
}

func (o *StatusDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/statuses/{id}][%d] statusDeleteForbidden", 403)
}

func (o *StatusDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/statuses/{id}][%d] statusDeleteForbidden", 403)
}

func (o *StatusDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusDeleteNotFound creates a StatusDeleteNotFound with default headers values
func NewStatusDeleteNotFound() *StatusDeleteNotFound {
	return &StatusDeleteNotFound{}
}

/*
StatusDeleteNotFound describes a response with status code 404, with default header values.

not found
*/
type StatusDeleteNotFound struct {
}

// IsSuccess returns true when this status delete not found response has a 2xx status code
func (o *StatusDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status delete not found response has a 3xx status code
func (o *StatusDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status delete not found response has a 4xx status code
func (o *StatusDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this status delete not found response has a 5xx status code
func (o *StatusDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this status delete not found response a status code equal to that given
func (o *StatusDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the status delete not found response
func (o *StatusDeleteNotFound) Code() int {
	return 404
}

func (o *StatusDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/statuses/{id}][%d] statusDeleteNotFound", 404)
}

func (o *StatusDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/statuses/{id}][%d] statusDeleteNotFound", 404)
}

func (o *StatusDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusDeleteNotAcceptable creates a StatusDeleteNotAcceptable with default headers values
func NewStatusDeleteNotAcceptable() *StatusDeleteNotAcceptable {
	return &StatusDeleteNotAcceptable{}
}

/*
StatusDeleteNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type StatusDeleteNotAcceptable struct {
}

// IsSuccess returns true when this status delete not acceptable response has a 2xx status code
func (o *StatusDeleteNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status delete not acceptable response has a 3xx status code
func (o *StatusDeleteNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status delete not acceptable response has a 4xx status code
func (o *StatusDeleteNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this status delete not acceptable response has a 5xx status code
func (o *StatusDeleteNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this status delete not acceptable response a status code equal to that given
func (o *StatusDeleteNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the status delete not acceptable response
func (o *StatusDeleteNotAcceptable) Code() int {
	return 406
}

func (o *StatusDeleteNotAcceptable) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/statuses/{id}][%d] statusDeleteNotAcceptable", 406)
}

func (o *StatusDeleteNotAcceptable) String() string {
	return fmt.Sprintf("[DELETE /api/v1/statuses/{id}][%d] statusDeleteNotAcceptable", 406)
}

func (o *StatusDeleteNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusDeleteInternalServerError creates a StatusDeleteInternalServerError with default headers values
func NewStatusDeleteInternalServerError() *StatusDeleteInternalServerError {
	return &StatusDeleteInternalServerError{}
}

/*
StatusDeleteInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type StatusDeleteInternalServerError struct {
}

// IsSuccess returns true when this status delete internal server error response has a 2xx status code
func (o *StatusDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status delete internal server error response has a 3xx status code
func (o *StatusDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status delete internal server error response has a 4xx status code
func (o *StatusDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this status delete internal server error response has a 5xx status code
func (o *StatusDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this status delete internal server error response a status code equal to that given
func (o *StatusDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the status delete internal server error response
func (o *StatusDeleteInternalServerError) Code() int {
	return 500
}

func (o *StatusDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/statuses/{id}][%d] statusDeleteInternalServerError", 500)
}

func (o *StatusDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v1/statuses/{id}][%d] statusDeleteInternalServerError", 500)
}

func (o *StatusDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
