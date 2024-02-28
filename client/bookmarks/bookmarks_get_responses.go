// Code generated by go-swagger; DO NOT EDIT.

package bookmarks

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

// BookmarksGetReader is a Reader for the BookmarksGet structure.
type BookmarksGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BookmarksGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBookmarksGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewBookmarksGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewBookmarksGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBookmarksGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/bookmarks] bookmarksGet", response, response.Code())
	}
}

// NewBookmarksGetOK creates a BookmarksGetOK with default headers values
func NewBookmarksGetOK() *BookmarksGetOK {
	return &BookmarksGetOK{}
}

/*
BookmarksGetOK describes a response with status code 200, with default header values.

Array of bookmarked statuses
*/
type BookmarksGetOK struct {

	/* Links to the next and previous queries.
	 */
	Link string

	Payload []*models.Status
}

// IsSuccess returns true when this bookmarks get o k response has a 2xx status code
func (o *BookmarksGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this bookmarks get o k response has a 3xx status code
func (o *BookmarksGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bookmarks get o k response has a 4xx status code
func (o *BookmarksGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this bookmarks get o k response has a 5xx status code
func (o *BookmarksGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this bookmarks get o k response a status code equal to that given
func (o *BookmarksGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the bookmarks get o k response
func (o *BookmarksGetOK) Code() int {
	return 200
}

func (o *BookmarksGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/bookmarks][%d] bookmarksGetOK %s", 200, payload)
}

func (o *BookmarksGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/bookmarks][%d] bookmarksGetOK %s", 200, payload)
}

func (o *BookmarksGetOK) GetPayload() []*models.Status {
	return o.Payload
}

func (o *BookmarksGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookmarksGetUnauthorized creates a BookmarksGetUnauthorized with default headers values
func NewBookmarksGetUnauthorized() *BookmarksGetUnauthorized {
	return &BookmarksGetUnauthorized{}
}

/*
BookmarksGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type BookmarksGetUnauthorized struct {
}

// IsSuccess returns true when this bookmarks get unauthorized response has a 2xx status code
func (o *BookmarksGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bookmarks get unauthorized response has a 3xx status code
func (o *BookmarksGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bookmarks get unauthorized response has a 4xx status code
func (o *BookmarksGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this bookmarks get unauthorized response has a 5xx status code
func (o *BookmarksGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this bookmarks get unauthorized response a status code equal to that given
func (o *BookmarksGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the bookmarks get unauthorized response
func (o *BookmarksGetUnauthorized) Code() int {
	return 401
}

func (o *BookmarksGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/bookmarks][%d] bookmarksGetUnauthorized", 401)
}

func (o *BookmarksGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/bookmarks][%d] bookmarksGetUnauthorized", 401)
}

func (o *BookmarksGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBookmarksGetNotAcceptable creates a BookmarksGetNotAcceptable with default headers values
func NewBookmarksGetNotAcceptable() *BookmarksGetNotAcceptable {
	return &BookmarksGetNotAcceptable{}
}

/*
BookmarksGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type BookmarksGetNotAcceptable struct {
}

// IsSuccess returns true when this bookmarks get not acceptable response has a 2xx status code
func (o *BookmarksGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bookmarks get not acceptable response has a 3xx status code
func (o *BookmarksGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bookmarks get not acceptable response has a 4xx status code
func (o *BookmarksGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this bookmarks get not acceptable response has a 5xx status code
func (o *BookmarksGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this bookmarks get not acceptable response a status code equal to that given
func (o *BookmarksGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the bookmarks get not acceptable response
func (o *BookmarksGetNotAcceptable) Code() int {
	return 406
}

func (o *BookmarksGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/bookmarks][%d] bookmarksGetNotAcceptable", 406)
}

func (o *BookmarksGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/bookmarks][%d] bookmarksGetNotAcceptable", 406)
}

func (o *BookmarksGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBookmarksGetInternalServerError creates a BookmarksGetInternalServerError with default headers values
func NewBookmarksGetInternalServerError() *BookmarksGetInternalServerError {
	return &BookmarksGetInternalServerError{}
}

/*
BookmarksGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type BookmarksGetInternalServerError struct {
}

// IsSuccess returns true when this bookmarks get internal server error response has a 2xx status code
func (o *BookmarksGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bookmarks get internal server error response has a 3xx status code
func (o *BookmarksGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bookmarks get internal server error response has a 4xx status code
func (o *BookmarksGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this bookmarks get internal server error response has a 5xx status code
func (o *BookmarksGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this bookmarks get internal server error response a status code equal to that given
func (o *BookmarksGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the bookmarks get internal server error response
func (o *BookmarksGetInternalServerError) Code() int {
	return 500
}

func (o *BookmarksGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/bookmarks][%d] bookmarksGetInternalServerError", 500)
}

func (o *BookmarksGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/bookmarks][%d] bookmarksGetInternalServerError", 500)
}

func (o *BookmarksGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
