// Code generated by go-swagger; DO NOT EDIT.

package lists

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

// ListsReader is a Reader for the Lists structure.
type ListsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewListsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/lists] lists", response, response.Code())
	}
}

// NewListsOK creates a ListsOK with default headers values
func NewListsOK() *ListsOK {
	return &ListsOK{}
}

/*
ListsOK describes a response with status code 200, with default header values.

Array of all lists owned by the requesting user.
*/
type ListsOK struct {
	Payload []*models.List
}

// IsSuccess returns true when this lists o k response has a 2xx status code
func (o *ListsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lists o k response has a 3xx status code
func (o *ListsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lists o k response has a 4xx status code
func (o *ListsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lists o k response has a 5xx status code
func (o *ListsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this lists o k response a status code equal to that given
func (o *ListsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the lists o k response
func (o *ListsOK) Code() int {
	return 200
}

func (o *ListsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/lists][%d] listsOK %s", 200, payload)
}

func (o *ListsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/lists][%d] listsOK %s", 200, payload)
}

func (o *ListsOK) GetPayload() []*models.List {
	return o.Payload
}

func (o *ListsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsBadRequest creates a ListsBadRequest with default headers values
func NewListsBadRequest() *ListsBadRequest {
	return &ListsBadRequest{}
}

/*
ListsBadRequest describes a response with status code 400, with default header values.

bad request
*/
type ListsBadRequest struct {
}

// IsSuccess returns true when this lists bad request response has a 2xx status code
func (o *ListsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this lists bad request response has a 3xx status code
func (o *ListsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lists bad request response has a 4xx status code
func (o *ListsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this lists bad request response has a 5xx status code
func (o *ListsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this lists bad request response a status code equal to that given
func (o *ListsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the lists bad request response
func (o *ListsBadRequest) Code() int {
	return 400
}

func (o *ListsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/lists][%d] listsBadRequest", 400)
}

func (o *ListsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/lists][%d] listsBadRequest", 400)
}

func (o *ListsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListsUnauthorized creates a ListsUnauthorized with default headers values
func NewListsUnauthorized() *ListsUnauthorized {
	return &ListsUnauthorized{}
}

/*
ListsUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type ListsUnauthorized struct {
}

// IsSuccess returns true when this lists unauthorized response has a 2xx status code
func (o *ListsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this lists unauthorized response has a 3xx status code
func (o *ListsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lists unauthorized response has a 4xx status code
func (o *ListsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this lists unauthorized response has a 5xx status code
func (o *ListsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this lists unauthorized response a status code equal to that given
func (o *ListsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the lists unauthorized response
func (o *ListsUnauthorized) Code() int {
	return 401
}

func (o *ListsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/lists][%d] listsUnauthorized", 401)
}

func (o *ListsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/lists][%d] listsUnauthorized", 401)
}

func (o *ListsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListsNotFound creates a ListsNotFound with default headers values
func NewListsNotFound() *ListsNotFound {
	return &ListsNotFound{}
}

/*
ListsNotFound describes a response with status code 404, with default header values.

not found
*/
type ListsNotFound struct {
}

// IsSuccess returns true when this lists not found response has a 2xx status code
func (o *ListsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this lists not found response has a 3xx status code
func (o *ListsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lists not found response has a 4xx status code
func (o *ListsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this lists not found response has a 5xx status code
func (o *ListsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this lists not found response a status code equal to that given
func (o *ListsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the lists not found response
func (o *ListsNotFound) Code() int {
	return 404
}

func (o *ListsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/lists][%d] listsNotFound", 404)
}

func (o *ListsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/lists][%d] listsNotFound", 404)
}

func (o *ListsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListsNotAcceptable creates a ListsNotAcceptable with default headers values
func NewListsNotAcceptable() *ListsNotAcceptable {
	return &ListsNotAcceptable{}
}

/*
ListsNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type ListsNotAcceptable struct {
}

// IsSuccess returns true when this lists not acceptable response has a 2xx status code
func (o *ListsNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this lists not acceptable response has a 3xx status code
func (o *ListsNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lists not acceptable response has a 4xx status code
func (o *ListsNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this lists not acceptable response has a 5xx status code
func (o *ListsNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this lists not acceptable response a status code equal to that given
func (o *ListsNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the lists not acceptable response
func (o *ListsNotAcceptable) Code() int {
	return 406
}

func (o *ListsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/lists][%d] listsNotAcceptable", 406)
}

func (o *ListsNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/lists][%d] listsNotAcceptable", 406)
}

func (o *ListsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListsInternalServerError creates a ListsInternalServerError with default headers values
func NewListsInternalServerError() *ListsInternalServerError {
	return &ListsInternalServerError{}
}

/*
ListsInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type ListsInternalServerError struct {
}

// IsSuccess returns true when this lists internal server error response has a 2xx status code
func (o *ListsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this lists internal server error response has a 3xx status code
func (o *ListsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lists internal server error response has a 4xx status code
func (o *ListsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this lists internal server error response has a 5xx status code
func (o *ListsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this lists internal server error response a status code equal to that given
func (o *ListsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the lists internal server error response
func (o *ListsInternalServerError) Code() int {
	return 500
}

func (o *ListsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/lists][%d] listsInternalServerError", 500)
}

func (o *ListsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/lists][%d] listsInternalServerError", 500)
}

func (o *ListsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
