// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// HeaderFilterBlockGetReader is a Reader for the HeaderFilterBlockGet structure.
type HeaderFilterBlockGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeaderFilterBlockGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeaderFilterBlockGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHeaderFilterBlockGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewHeaderFilterBlockGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHeaderFilterBlockGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHeaderFilterBlockGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHeaderFilterBlockGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/admin/header_blocks/{id}] headerFilterBlockGet", response, response.Code())
	}
}

// NewHeaderFilterBlockGetOK creates a HeaderFilterBlockGetOK with default headers values
func NewHeaderFilterBlockGetOK() *HeaderFilterBlockGetOK {
	return &HeaderFilterBlockGetOK{}
}

/*
HeaderFilterBlockGetOK describes a response with status code 200, with default header values.

The requested "block" header filter.
*/
type HeaderFilterBlockGetOK struct {
	Payload *models.HeaderFilter
}

// IsSuccess returns true when this header filter block get o k response has a 2xx status code
func (o *HeaderFilterBlockGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this header filter block get o k response has a 3xx status code
func (o *HeaderFilterBlockGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter block get o k response has a 4xx status code
func (o *HeaderFilterBlockGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this header filter block get o k response has a 5xx status code
func (o *HeaderFilterBlockGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter block get o k response a status code equal to that given
func (o *HeaderFilterBlockGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the header filter block get o k response
func (o *HeaderFilterBlockGetOK) Code() int {
	return 200
}

func (o *HeaderFilterBlockGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockGetOK %s", 200, payload)
}

func (o *HeaderFilterBlockGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockGetOK %s", 200, payload)
}

func (o *HeaderFilterBlockGetOK) GetPayload() *models.HeaderFilter {
	return o.Payload
}

func (o *HeaderFilterBlockGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HeaderFilter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeaderFilterBlockGetBadRequest creates a HeaderFilterBlockGetBadRequest with default headers values
func NewHeaderFilterBlockGetBadRequest() *HeaderFilterBlockGetBadRequest {
	return &HeaderFilterBlockGetBadRequest{}
}

/*
HeaderFilterBlockGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type HeaderFilterBlockGetBadRequest struct {
}

// IsSuccess returns true when this header filter block get bad request response has a 2xx status code
func (o *HeaderFilterBlockGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter block get bad request response has a 3xx status code
func (o *HeaderFilterBlockGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter block get bad request response has a 4xx status code
func (o *HeaderFilterBlockGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this header filter block get bad request response has a 5xx status code
func (o *HeaderFilterBlockGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter block get bad request response a status code equal to that given
func (o *HeaderFilterBlockGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the header filter block get bad request response
func (o *HeaderFilterBlockGetBadRequest) Code() int {
	return 400
}

func (o *HeaderFilterBlockGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockGetBadRequest", 400)
}

func (o *HeaderFilterBlockGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockGetBadRequest", 400)
}

func (o *HeaderFilterBlockGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeaderFilterBlockGetUnauthorized creates a HeaderFilterBlockGetUnauthorized with default headers values
func NewHeaderFilterBlockGetUnauthorized() *HeaderFilterBlockGetUnauthorized {
	return &HeaderFilterBlockGetUnauthorized{}
}

/*
HeaderFilterBlockGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type HeaderFilterBlockGetUnauthorized struct {
}

// IsSuccess returns true when this header filter block get unauthorized response has a 2xx status code
func (o *HeaderFilterBlockGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter block get unauthorized response has a 3xx status code
func (o *HeaderFilterBlockGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter block get unauthorized response has a 4xx status code
func (o *HeaderFilterBlockGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this header filter block get unauthorized response has a 5xx status code
func (o *HeaderFilterBlockGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter block get unauthorized response a status code equal to that given
func (o *HeaderFilterBlockGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the header filter block get unauthorized response
func (o *HeaderFilterBlockGetUnauthorized) Code() int {
	return 401
}

func (o *HeaderFilterBlockGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockGetUnauthorized", 401)
}

func (o *HeaderFilterBlockGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockGetUnauthorized", 401)
}

func (o *HeaderFilterBlockGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeaderFilterBlockGetForbidden creates a HeaderFilterBlockGetForbidden with default headers values
func NewHeaderFilterBlockGetForbidden() *HeaderFilterBlockGetForbidden {
	return &HeaderFilterBlockGetForbidden{}
}

/*
HeaderFilterBlockGetForbidden describes a response with status code 403, with default header values.

forbidden
*/
type HeaderFilterBlockGetForbidden struct {
}

// IsSuccess returns true when this header filter block get forbidden response has a 2xx status code
func (o *HeaderFilterBlockGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter block get forbidden response has a 3xx status code
func (o *HeaderFilterBlockGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter block get forbidden response has a 4xx status code
func (o *HeaderFilterBlockGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this header filter block get forbidden response has a 5xx status code
func (o *HeaderFilterBlockGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter block get forbidden response a status code equal to that given
func (o *HeaderFilterBlockGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the header filter block get forbidden response
func (o *HeaderFilterBlockGetForbidden) Code() int {
	return 403
}

func (o *HeaderFilterBlockGetForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockGetForbidden", 403)
}

func (o *HeaderFilterBlockGetForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockGetForbidden", 403)
}

func (o *HeaderFilterBlockGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeaderFilterBlockGetNotFound creates a HeaderFilterBlockGetNotFound with default headers values
func NewHeaderFilterBlockGetNotFound() *HeaderFilterBlockGetNotFound {
	return &HeaderFilterBlockGetNotFound{}
}

/*
HeaderFilterBlockGetNotFound describes a response with status code 404, with default header values.

not found
*/
type HeaderFilterBlockGetNotFound struct {
}

// IsSuccess returns true when this header filter block get not found response has a 2xx status code
func (o *HeaderFilterBlockGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter block get not found response has a 3xx status code
func (o *HeaderFilterBlockGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter block get not found response has a 4xx status code
func (o *HeaderFilterBlockGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this header filter block get not found response has a 5xx status code
func (o *HeaderFilterBlockGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter block get not found response a status code equal to that given
func (o *HeaderFilterBlockGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the header filter block get not found response
func (o *HeaderFilterBlockGetNotFound) Code() int {
	return 404
}

func (o *HeaderFilterBlockGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockGetNotFound", 404)
}

func (o *HeaderFilterBlockGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockGetNotFound", 404)
}

func (o *HeaderFilterBlockGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeaderFilterBlockGetInternalServerError creates a HeaderFilterBlockGetInternalServerError with default headers values
func NewHeaderFilterBlockGetInternalServerError() *HeaderFilterBlockGetInternalServerError {
	return &HeaderFilterBlockGetInternalServerError{}
}

/*
HeaderFilterBlockGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type HeaderFilterBlockGetInternalServerError struct {
}

// IsSuccess returns true when this header filter block get internal server error response has a 2xx status code
func (o *HeaderFilterBlockGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter block get internal server error response has a 3xx status code
func (o *HeaderFilterBlockGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter block get internal server error response has a 4xx status code
func (o *HeaderFilterBlockGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this header filter block get internal server error response has a 5xx status code
func (o *HeaderFilterBlockGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this header filter block get internal server error response a status code equal to that given
func (o *HeaderFilterBlockGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the header filter block get internal server error response
func (o *HeaderFilterBlockGetInternalServerError) Code() int {
	return 500
}

func (o *HeaderFilterBlockGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockGetInternalServerError", 500)
}

func (o *HeaderFilterBlockGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks/{id}][%d] headerFilterBlockGetInternalServerError", 500)
}

func (o *HeaderFilterBlockGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
