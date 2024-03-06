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

// ListReader is a Reader for the List structure.
type ListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewListNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/lists/{id}] list", response, response.Code())
	}
}

// NewListOK creates a ListOK with default headers values
func NewListOK() *ListOK {
	return &ListOK{}
}

/*
ListOK describes a response with status code 200, with default header values.

Requested list.
*/
type ListOK struct {
	Payload *models.List
}

// IsSuccess returns true when this list o k response has a 2xx status code
func (o *ListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list o k response has a 3xx status code
func (o *ListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list o k response has a 4xx status code
func (o *ListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list o k response has a 5xx status code
func (o *ListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list o k response a status code equal to that given
func (o *ListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list o k response
func (o *ListOK) Code() int {
	return 200
}

func (o *ListOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/lists/{id}][%d] listOK  %+v", 200, o.Payload)
}

func (o *ListOK) String() string {
	return fmt.Sprintf("[GET /api/v1/lists/{id}][%d] listOK  %+v", 200, o.Payload)
}

func (o *ListOK) GetPayload() *models.List {
	return o.Payload
}

func (o *ListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.List)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBadRequest creates a ListBadRequest with default headers values
func NewListBadRequest() *ListBadRequest {
	return &ListBadRequest{}
}

/*
ListBadRequest describes a response with status code 400, with default header values.

bad request
*/
type ListBadRequest struct {
}

// IsSuccess returns true when this list bad request response has a 2xx status code
func (o *ListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list bad request response has a 3xx status code
func (o *ListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list bad request response has a 4xx status code
func (o *ListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list bad request response has a 5xx status code
func (o *ListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list bad request response a status code equal to that given
func (o *ListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list bad request response
func (o *ListBadRequest) Code() int {
	return 400
}

func (o *ListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/lists/{id}][%d] listBadRequest ", 400)
}

func (o *ListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/lists/{id}][%d] listBadRequest ", 400)
}

func (o *ListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListUnauthorized creates a ListUnauthorized with default headers values
func NewListUnauthorized() *ListUnauthorized {
	return &ListUnauthorized{}
}

/*
ListUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type ListUnauthorized struct {
}

// IsSuccess returns true when this list unauthorized response has a 2xx status code
func (o *ListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list unauthorized response has a 3xx status code
func (o *ListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list unauthorized response has a 4xx status code
func (o *ListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list unauthorized response has a 5xx status code
func (o *ListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list unauthorized response a status code equal to that given
func (o *ListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list unauthorized response
func (o *ListUnauthorized) Code() int {
	return 401
}

func (o *ListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/lists/{id}][%d] listUnauthorized ", 401)
}

func (o *ListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/lists/{id}][%d] listUnauthorized ", 401)
}

func (o *ListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListNotFound creates a ListNotFound with default headers values
func NewListNotFound() *ListNotFound {
	return &ListNotFound{}
}

/*
ListNotFound describes a response with status code 404, with default header values.

not found
*/
type ListNotFound struct {
}

// IsSuccess returns true when this list not found response has a 2xx status code
func (o *ListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list not found response has a 3xx status code
func (o *ListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list not found response has a 4xx status code
func (o *ListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list not found response has a 5xx status code
func (o *ListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list not found response a status code equal to that given
func (o *ListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list not found response
func (o *ListNotFound) Code() int {
	return 404
}

func (o *ListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/lists/{id}][%d] listNotFound ", 404)
}

func (o *ListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/lists/{id}][%d] listNotFound ", 404)
}

func (o *ListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListNotAcceptable creates a ListNotAcceptable with default headers values
func NewListNotAcceptable() *ListNotAcceptable {
	return &ListNotAcceptable{}
}

/*
ListNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type ListNotAcceptable struct {
}

// IsSuccess returns true when this list not acceptable response has a 2xx status code
func (o *ListNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list not acceptable response has a 3xx status code
func (o *ListNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list not acceptable response has a 4xx status code
func (o *ListNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this list not acceptable response has a 5xx status code
func (o *ListNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this list not acceptable response a status code equal to that given
func (o *ListNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the list not acceptable response
func (o *ListNotAcceptable) Code() int {
	return 406
}

func (o *ListNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/lists/{id}][%d] listNotAcceptable ", 406)
}

func (o *ListNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/lists/{id}][%d] listNotAcceptable ", 406)
}

func (o *ListNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListInternalServerError creates a ListInternalServerError with default headers values
func NewListInternalServerError() *ListInternalServerError {
	return &ListInternalServerError{}
}

/*
ListInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type ListInternalServerError struct {
}

// IsSuccess returns true when this list internal server error response has a 2xx status code
func (o *ListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list internal server error response has a 3xx status code
func (o *ListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list internal server error response has a 4xx status code
func (o *ListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list internal server error response has a 5xx status code
func (o *ListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list internal server error response a status code equal to that given
func (o *ListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list internal server error response
func (o *ListInternalServerError) Code() int {
	return 500
}

func (o *ListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/lists/{id}][%d] listInternalServerError ", 500)
}

func (o *ListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/lists/{id}][%d] listInternalServerError ", 500)
}

func (o *ListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
