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

// HeaderFilterBlockCreateReader is a Reader for the HeaderFilterBlockCreate structure.
type HeaderFilterBlockCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeaderFilterBlockCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeaderFilterBlockCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHeaderFilterBlockCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewHeaderFilterBlockCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHeaderFilterBlockCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHeaderFilterBlockCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/admin/header_blocks] headerFilterBlockCreate", response, response.Code())
	}
}

// NewHeaderFilterBlockCreateOK creates a HeaderFilterBlockCreateOK with default headers values
func NewHeaderFilterBlockCreateOK() *HeaderFilterBlockCreateOK {
	return &HeaderFilterBlockCreateOK{}
}

/*
HeaderFilterBlockCreateOK describes a response with status code 200, with default header values.

The newly created "block" header filter.
*/
type HeaderFilterBlockCreateOK struct {
	Payload *models.HeaderFilter
}

// IsSuccess returns true when this header filter block create o k response has a 2xx status code
func (o *HeaderFilterBlockCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this header filter block create o k response has a 3xx status code
func (o *HeaderFilterBlockCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter block create o k response has a 4xx status code
func (o *HeaderFilterBlockCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this header filter block create o k response has a 5xx status code
func (o *HeaderFilterBlockCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter block create o k response a status code equal to that given
func (o *HeaderFilterBlockCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the header filter block create o k response
func (o *HeaderFilterBlockCreateOK) Code() int {
	return 200
}

func (o *HeaderFilterBlockCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/admin/header_blocks][%d] headerFilterBlockCreateOK %s", 200, payload)
}

func (o *HeaderFilterBlockCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/admin/header_blocks][%d] headerFilterBlockCreateOK %s", 200, payload)
}

func (o *HeaderFilterBlockCreateOK) GetPayload() *models.HeaderFilter {
	return o.Payload
}

func (o *HeaderFilterBlockCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HeaderFilter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeaderFilterBlockCreateBadRequest creates a HeaderFilterBlockCreateBadRequest with default headers values
func NewHeaderFilterBlockCreateBadRequest() *HeaderFilterBlockCreateBadRequest {
	return &HeaderFilterBlockCreateBadRequest{}
}

/*
HeaderFilterBlockCreateBadRequest describes a response with status code 400, with default header values.

bad request
*/
type HeaderFilterBlockCreateBadRequest struct {
}

// IsSuccess returns true when this header filter block create bad request response has a 2xx status code
func (o *HeaderFilterBlockCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter block create bad request response has a 3xx status code
func (o *HeaderFilterBlockCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter block create bad request response has a 4xx status code
func (o *HeaderFilterBlockCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this header filter block create bad request response has a 5xx status code
func (o *HeaderFilterBlockCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter block create bad request response a status code equal to that given
func (o *HeaderFilterBlockCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the header filter block create bad request response
func (o *HeaderFilterBlockCreateBadRequest) Code() int {
	return 400
}

func (o *HeaderFilterBlockCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/header_blocks][%d] headerFilterBlockCreateBadRequest", 400)
}

func (o *HeaderFilterBlockCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/header_blocks][%d] headerFilterBlockCreateBadRequest", 400)
}

func (o *HeaderFilterBlockCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeaderFilterBlockCreateUnauthorized creates a HeaderFilterBlockCreateUnauthorized with default headers values
func NewHeaderFilterBlockCreateUnauthorized() *HeaderFilterBlockCreateUnauthorized {
	return &HeaderFilterBlockCreateUnauthorized{}
}

/*
HeaderFilterBlockCreateUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type HeaderFilterBlockCreateUnauthorized struct {
}

// IsSuccess returns true when this header filter block create unauthorized response has a 2xx status code
func (o *HeaderFilterBlockCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter block create unauthorized response has a 3xx status code
func (o *HeaderFilterBlockCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter block create unauthorized response has a 4xx status code
func (o *HeaderFilterBlockCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this header filter block create unauthorized response has a 5xx status code
func (o *HeaderFilterBlockCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter block create unauthorized response a status code equal to that given
func (o *HeaderFilterBlockCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the header filter block create unauthorized response
func (o *HeaderFilterBlockCreateUnauthorized) Code() int {
	return 401
}

func (o *HeaderFilterBlockCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/header_blocks][%d] headerFilterBlockCreateUnauthorized", 401)
}

func (o *HeaderFilterBlockCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/header_blocks][%d] headerFilterBlockCreateUnauthorized", 401)
}

func (o *HeaderFilterBlockCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeaderFilterBlockCreateForbidden creates a HeaderFilterBlockCreateForbidden with default headers values
func NewHeaderFilterBlockCreateForbidden() *HeaderFilterBlockCreateForbidden {
	return &HeaderFilterBlockCreateForbidden{}
}

/*
HeaderFilterBlockCreateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type HeaderFilterBlockCreateForbidden struct {
}

// IsSuccess returns true when this header filter block create forbidden response has a 2xx status code
func (o *HeaderFilterBlockCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter block create forbidden response has a 3xx status code
func (o *HeaderFilterBlockCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter block create forbidden response has a 4xx status code
func (o *HeaderFilterBlockCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this header filter block create forbidden response has a 5xx status code
func (o *HeaderFilterBlockCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter block create forbidden response a status code equal to that given
func (o *HeaderFilterBlockCreateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the header filter block create forbidden response
func (o *HeaderFilterBlockCreateForbidden) Code() int {
	return 403
}

func (o *HeaderFilterBlockCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/header_blocks][%d] headerFilterBlockCreateForbidden", 403)
}

func (o *HeaderFilterBlockCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/header_blocks][%d] headerFilterBlockCreateForbidden", 403)
}

func (o *HeaderFilterBlockCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeaderFilterBlockCreateInternalServerError creates a HeaderFilterBlockCreateInternalServerError with default headers values
func NewHeaderFilterBlockCreateInternalServerError() *HeaderFilterBlockCreateInternalServerError {
	return &HeaderFilterBlockCreateInternalServerError{}
}

/*
HeaderFilterBlockCreateInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type HeaderFilterBlockCreateInternalServerError struct {
}

// IsSuccess returns true when this header filter block create internal server error response has a 2xx status code
func (o *HeaderFilterBlockCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter block create internal server error response has a 3xx status code
func (o *HeaderFilterBlockCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter block create internal server error response has a 4xx status code
func (o *HeaderFilterBlockCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this header filter block create internal server error response has a 5xx status code
func (o *HeaderFilterBlockCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this header filter block create internal server error response a status code equal to that given
func (o *HeaderFilterBlockCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the header filter block create internal server error response
func (o *HeaderFilterBlockCreateInternalServerError) Code() int {
	return 500
}

func (o *HeaderFilterBlockCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/header_blocks][%d] headerFilterBlockCreateInternalServerError", 500)
}

func (o *HeaderFilterBlockCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/header_blocks][%d] headerFilterBlockCreateInternalServerError", 500)
}

func (o *HeaderFilterBlockCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
