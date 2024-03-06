// Code generated by go-swagger; DO NOT EDIT.

package lists

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VyrCossont/slurp/models"
)

// ListCreateReader is a Reader for the ListCreate structure.
type ListCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewListCreateNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/lists] listCreate", response, response.Code())
	}
}

// NewListCreateOK creates a ListCreateOK with default headers values
func NewListCreateOK() *ListCreateOK {
	return &ListCreateOK{}
}

/*
ListCreateOK describes a response with status code 200, with default header values.

The newly created list.
*/
type ListCreateOK struct {
	Payload *models.List
}

// IsSuccess returns true when this list create o k response has a 2xx status code
func (o *ListCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list create o k response has a 3xx status code
func (o *ListCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list create o k response has a 4xx status code
func (o *ListCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list create o k response has a 5xx status code
func (o *ListCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list create o k response a status code equal to that given
func (o *ListCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list create o k response
func (o *ListCreateOK) Code() int {
	return 200
}

func (o *ListCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/lists][%d] listCreateOK  %+v", 200, o.Payload)
}

func (o *ListCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v1/lists][%d] listCreateOK  %+v", 200, o.Payload)
}

func (o *ListCreateOK) GetPayload() *models.List {
	return o.Payload
}

func (o *ListCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.List)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCreateBadRequest creates a ListCreateBadRequest with default headers values
func NewListCreateBadRequest() *ListCreateBadRequest {
	return &ListCreateBadRequest{}
}

/*
ListCreateBadRequest describes a response with status code 400, with default header values.

bad request
*/
type ListCreateBadRequest struct {
}

// IsSuccess returns true when this list create bad request response has a 2xx status code
func (o *ListCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list create bad request response has a 3xx status code
func (o *ListCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list create bad request response has a 4xx status code
func (o *ListCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list create bad request response has a 5xx status code
func (o *ListCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list create bad request response a status code equal to that given
func (o *ListCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list create bad request response
func (o *ListCreateBadRequest) Code() int {
	return 400
}

func (o *ListCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/lists][%d] listCreateBadRequest ", 400)
}

func (o *ListCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/lists][%d] listCreateBadRequest ", 400)
}

func (o *ListCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListCreateUnauthorized creates a ListCreateUnauthorized with default headers values
func NewListCreateUnauthorized() *ListCreateUnauthorized {
	return &ListCreateUnauthorized{}
}

/*
ListCreateUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type ListCreateUnauthorized struct {
}

// IsSuccess returns true when this list create unauthorized response has a 2xx status code
func (o *ListCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list create unauthorized response has a 3xx status code
func (o *ListCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list create unauthorized response has a 4xx status code
func (o *ListCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list create unauthorized response has a 5xx status code
func (o *ListCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list create unauthorized response a status code equal to that given
func (o *ListCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list create unauthorized response
func (o *ListCreateUnauthorized) Code() int {
	return 401
}

func (o *ListCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/lists][%d] listCreateUnauthorized ", 401)
}

func (o *ListCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/lists][%d] listCreateUnauthorized ", 401)
}

func (o *ListCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListCreateForbidden creates a ListCreateForbidden with default headers values
func NewListCreateForbidden() *ListCreateForbidden {
	return &ListCreateForbidden{}
}

/*
ListCreateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type ListCreateForbidden struct {
}

// IsSuccess returns true when this list create forbidden response has a 2xx status code
func (o *ListCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list create forbidden response has a 3xx status code
func (o *ListCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list create forbidden response has a 4xx status code
func (o *ListCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list create forbidden response has a 5xx status code
func (o *ListCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list create forbidden response a status code equal to that given
func (o *ListCreateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list create forbidden response
func (o *ListCreateForbidden) Code() int {
	return 403
}

func (o *ListCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/lists][%d] listCreateForbidden ", 403)
}

func (o *ListCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/lists][%d] listCreateForbidden ", 403)
}

func (o *ListCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListCreateNotFound creates a ListCreateNotFound with default headers values
func NewListCreateNotFound() *ListCreateNotFound {
	return &ListCreateNotFound{}
}

/*
ListCreateNotFound describes a response with status code 404, with default header values.

not found
*/
type ListCreateNotFound struct {
}

// IsSuccess returns true when this list create not found response has a 2xx status code
func (o *ListCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list create not found response has a 3xx status code
func (o *ListCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list create not found response has a 4xx status code
func (o *ListCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list create not found response has a 5xx status code
func (o *ListCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list create not found response a status code equal to that given
func (o *ListCreateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list create not found response
func (o *ListCreateNotFound) Code() int {
	return 404
}

func (o *ListCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/lists][%d] listCreateNotFound ", 404)
}

func (o *ListCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/lists][%d] listCreateNotFound ", 404)
}

func (o *ListCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListCreateNotAcceptable creates a ListCreateNotAcceptable with default headers values
func NewListCreateNotAcceptable() *ListCreateNotAcceptable {
	return &ListCreateNotAcceptable{}
}

/*
ListCreateNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type ListCreateNotAcceptable struct {
}

// IsSuccess returns true when this list create not acceptable response has a 2xx status code
func (o *ListCreateNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list create not acceptable response has a 3xx status code
func (o *ListCreateNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list create not acceptable response has a 4xx status code
func (o *ListCreateNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this list create not acceptable response has a 5xx status code
func (o *ListCreateNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this list create not acceptable response a status code equal to that given
func (o *ListCreateNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the list create not acceptable response
func (o *ListCreateNotAcceptable) Code() int {
	return 406
}

func (o *ListCreateNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/lists][%d] listCreateNotAcceptable ", 406)
}

func (o *ListCreateNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v1/lists][%d] listCreateNotAcceptable ", 406)
}

func (o *ListCreateNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListCreateInternalServerError creates a ListCreateInternalServerError with default headers values
func NewListCreateInternalServerError() *ListCreateInternalServerError {
	return &ListCreateInternalServerError{}
}

/*
ListCreateInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type ListCreateInternalServerError struct {
}

// IsSuccess returns true when this list create internal server error response has a 2xx status code
func (o *ListCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list create internal server error response has a 3xx status code
func (o *ListCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list create internal server error response has a 4xx status code
func (o *ListCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list create internal server error response has a 5xx status code
func (o *ListCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list create internal server error response a status code equal to that given
func (o *ListCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list create internal server error response
func (o *ListCreateInternalServerError) Code() int {
	return 500
}

func (o *ListCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/lists][%d] listCreateInternalServerError ", 500)
}

func (o *ListCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/lists][%d] listCreateInternalServerError ", 500)
}

func (o *ListCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
