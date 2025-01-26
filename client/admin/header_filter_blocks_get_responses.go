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

// HeaderFilterBlocksGetReader is a Reader for the HeaderFilterBlocksGet structure.
type HeaderFilterBlocksGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeaderFilterBlocksGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeaderFilterBlocksGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHeaderFilterBlocksGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewHeaderFilterBlocksGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHeaderFilterBlocksGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHeaderFilterBlocksGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHeaderFilterBlocksGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/admin/header_blocks] headerFilterBlocksGet", response, response.Code())
	}
}

// NewHeaderFilterBlocksGetOK creates a HeaderFilterBlocksGetOK with default headers values
func NewHeaderFilterBlocksGetOK() *HeaderFilterBlocksGetOK {
	return &HeaderFilterBlocksGetOK{}
}

/*
HeaderFilterBlocksGetOK describes a response with status code 200, with default header values.

All "block" header filters currently in place.
*/
type HeaderFilterBlocksGetOK struct {
	Payload []*models.HeaderFilter
}

// IsSuccess returns true when this header filter blocks get o k response has a 2xx status code
func (o *HeaderFilterBlocksGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this header filter blocks get o k response has a 3xx status code
func (o *HeaderFilterBlocksGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter blocks get o k response has a 4xx status code
func (o *HeaderFilterBlocksGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this header filter blocks get o k response has a 5xx status code
func (o *HeaderFilterBlocksGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter blocks get o k response a status code equal to that given
func (o *HeaderFilterBlocksGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the header filter blocks get o k response
func (o *HeaderFilterBlocksGetOK) Code() int {
	return 200
}

func (o *HeaderFilterBlocksGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks][%d] headerFilterBlocksGetOK %s", 200, payload)
}

func (o *HeaderFilterBlocksGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks][%d] headerFilterBlocksGetOK %s", 200, payload)
}

func (o *HeaderFilterBlocksGetOK) GetPayload() []*models.HeaderFilter {
	return o.Payload
}

func (o *HeaderFilterBlocksGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeaderFilterBlocksGetBadRequest creates a HeaderFilterBlocksGetBadRequest with default headers values
func NewHeaderFilterBlocksGetBadRequest() *HeaderFilterBlocksGetBadRequest {
	return &HeaderFilterBlocksGetBadRequest{}
}

/*
HeaderFilterBlocksGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type HeaderFilterBlocksGetBadRequest struct {
}

// IsSuccess returns true when this header filter blocks get bad request response has a 2xx status code
func (o *HeaderFilterBlocksGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter blocks get bad request response has a 3xx status code
func (o *HeaderFilterBlocksGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter blocks get bad request response has a 4xx status code
func (o *HeaderFilterBlocksGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this header filter blocks get bad request response has a 5xx status code
func (o *HeaderFilterBlocksGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter blocks get bad request response a status code equal to that given
func (o *HeaderFilterBlocksGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the header filter blocks get bad request response
func (o *HeaderFilterBlocksGetBadRequest) Code() int {
	return 400
}

func (o *HeaderFilterBlocksGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks][%d] headerFilterBlocksGetBadRequest", 400)
}

func (o *HeaderFilterBlocksGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks][%d] headerFilterBlocksGetBadRequest", 400)
}

func (o *HeaderFilterBlocksGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeaderFilterBlocksGetUnauthorized creates a HeaderFilterBlocksGetUnauthorized with default headers values
func NewHeaderFilterBlocksGetUnauthorized() *HeaderFilterBlocksGetUnauthorized {
	return &HeaderFilterBlocksGetUnauthorized{}
}

/*
HeaderFilterBlocksGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type HeaderFilterBlocksGetUnauthorized struct {
}

// IsSuccess returns true when this header filter blocks get unauthorized response has a 2xx status code
func (o *HeaderFilterBlocksGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter blocks get unauthorized response has a 3xx status code
func (o *HeaderFilterBlocksGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter blocks get unauthorized response has a 4xx status code
func (o *HeaderFilterBlocksGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this header filter blocks get unauthorized response has a 5xx status code
func (o *HeaderFilterBlocksGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter blocks get unauthorized response a status code equal to that given
func (o *HeaderFilterBlocksGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the header filter blocks get unauthorized response
func (o *HeaderFilterBlocksGetUnauthorized) Code() int {
	return 401
}

func (o *HeaderFilterBlocksGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks][%d] headerFilterBlocksGetUnauthorized", 401)
}

func (o *HeaderFilterBlocksGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks][%d] headerFilterBlocksGetUnauthorized", 401)
}

func (o *HeaderFilterBlocksGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeaderFilterBlocksGetForbidden creates a HeaderFilterBlocksGetForbidden with default headers values
func NewHeaderFilterBlocksGetForbidden() *HeaderFilterBlocksGetForbidden {
	return &HeaderFilterBlocksGetForbidden{}
}

/*
HeaderFilterBlocksGetForbidden describes a response with status code 403, with default header values.

forbidden
*/
type HeaderFilterBlocksGetForbidden struct {
}

// IsSuccess returns true when this header filter blocks get forbidden response has a 2xx status code
func (o *HeaderFilterBlocksGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter blocks get forbidden response has a 3xx status code
func (o *HeaderFilterBlocksGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter blocks get forbidden response has a 4xx status code
func (o *HeaderFilterBlocksGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this header filter blocks get forbidden response has a 5xx status code
func (o *HeaderFilterBlocksGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter blocks get forbidden response a status code equal to that given
func (o *HeaderFilterBlocksGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the header filter blocks get forbidden response
func (o *HeaderFilterBlocksGetForbidden) Code() int {
	return 403
}

func (o *HeaderFilterBlocksGetForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks][%d] headerFilterBlocksGetForbidden", 403)
}

func (o *HeaderFilterBlocksGetForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks][%d] headerFilterBlocksGetForbidden", 403)
}

func (o *HeaderFilterBlocksGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeaderFilterBlocksGetNotFound creates a HeaderFilterBlocksGetNotFound with default headers values
func NewHeaderFilterBlocksGetNotFound() *HeaderFilterBlocksGetNotFound {
	return &HeaderFilterBlocksGetNotFound{}
}

/*
HeaderFilterBlocksGetNotFound describes a response with status code 404, with default header values.

not found
*/
type HeaderFilterBlocksGetNotFound struct {
}

// IsSuccess returns true when this header filter blocks get not found response has a 2xx status code
func (o *HeaderFilterBlocksGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter blocks get not found response has a 3xx status code
func (o *HeaderFilterBlocksGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter blocks get not found response has a 4xx status code
func (o *HeaderFilterBlocksGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this header filter blocks get not found response has a 5xx status code
func (o *HeaderFilterBlocksGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter blocks get not found response a status code equal to that given
func (o *HeaderFilterBlocksGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the header filter blocks get not found response
func (o *HeaderFilterBlocksGetNotFound) Code() int {
	return 404
}

func (o *HeaderFilterBlocksGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks][%d] headerFilterBlocksGetNotFound", 404)
}

func (o *HeaderFilterBlocksGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks][%d] headerFilterBlocksGetNotFound", 404)
}

func (o *HeaderFilterBlocksGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeaderFilterBlocksGetInternalServerError creates a HeaderFilterBlocksGetInternalServerError with default headers values
func NewHeaderFilterBlocksGetInternalServerError() *HeaderFilterBlocksGetInternalServerError {
	return &HeaderFilterBlocksGetInternalServerError{}
}

/*
HeaderFilterBlocksGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type HeaderFilterBlocksGetInternalServerError struct {
}

// IsSuccess returns true when this header filter blocks get internal server error response has a 2xx status code
func (o *HeaderFilterBlocksGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter blocks get internal server error response has a 3xx status code
func (o *HeaderFilterBlocksGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter blocks get internal server error response has a 4xx status code
func (o *HeaderFilterBlocksGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this header filter blocks get internal server error response has a 5xx status code
func (o *HeaderFilterBlocksGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this header filter blocks get internal server error response a status code equal to that given
func (o *HeaderFilterBlocksGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the header filter blocks get internal server error response
func (o *HeaderFilterBlocksGetInternalServerError) Code() int {
	return 500
}

func (o *HeaderFilterBlocksGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks][%d] headerFilterBlocksGetInternalServerError", 500)
}

func (o *HeaderFilterBlocksGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/header_blocks][%d] headerFilterBlocksGetInternalServerError", 500)
}

func (o *HeaderFilterBlocksGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
