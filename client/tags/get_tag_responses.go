// Code generated by go-swagger; DO NOT EDIT.

package tags

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

// GetTagReader is a Reader for the GetTag structure.
type GetTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTagBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTagUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetTagNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTagInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/tags/{tag_name}] getTag", response, response.Code())
	}
}

// NewGetTagOK creates a GetTagOK with default headers values
func NewGetTagOK() *GetTagOK {
	return &GetTagOK{}
}

/*
GetTagOK describes a response with status code 200, with default header values.

Info about the tag.
*/
type GetTagOK struct {
	Payload *models.Tag
}

// IsSuccess returns true when this get tag o k response has a 2xx status code
func (o *GetTagOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get tag o k response has a 3xx status code
func (o *GetTagOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tag o k response has a 4xx status code
func (o *GetTagOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tag o k response has a 5xx status code
func (o *GetTagOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get tag o k response a status code equal to that given
func (o *GetTagOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get tag o k response
func (o *GetTagOK) Code() int {
	return 200
}

func (o *GetTagOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/tags/{tag_name}][%d] getTagOK %s", 200, payload)
}

func (o *GetTagOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/tags/{tag_name}][%d] getTagOK %s", 200, payload)
}

func (o *GetTagOK) GetPayload() *models.Tag {
	return o.Payload
}

func (o *GetTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTagBadRequest creates a GetTagBadRequest with default headers values
func NewGetTagBadRequest() *GetTagBadRequest {
	return &GetTagBadRequest{}
}

/*
GetTagBadRequest describes a response with status code 400, with default header values.

bad request
*/
type GetTagBadRequest struct {
}

// IsSuccess returns true when this get tag bad request response has a 2xx status code
func (o *GetTagBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tag bad request response has a 3xx status code
func (o *GetTagBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tag bad request response has a 4xx status code
func (o *GetTagBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tag bad request response has a 5xx status code
func (o *GetTagBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get tag bad request response a status code equal to that given
func (o *GetTagBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get tag bad request response
func (o *GetTagBadRequest) Code() int {
	return 400
}

func (o *GetTagBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/tags/{tag_name}][%d] getTagBadRequest", 400)
}

func (o *GetTagBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/tags/{tag_name}][%d] getTagBadRequest", 400)
}

func (o *GetTagBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTagUnauthorized creates a GetTagUnauthorized with default headers values
func NewGetTagUnauthorized() *GetTagUnauthorized {
	return &GetTagUnauthorized{}
}

/*
GetTagUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type GetTagUnauthorized struct {
}

// IsSuccess returns true when this get tag unauthorized response has a 2xx status code
func (o *GetTagUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tag unauthorized response has a 3xx status code
func (o *GetTagUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tag unauthorized response has a 4xx status code
func (o *GetTagUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tag unauthorized response has a 5xx status code
func (o *GetTagUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get tag unauthorized response a status code equal to that given
func (o *GetTagUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get tag unauthorized response
func (o *GetTagUnauthorized) Code() int {
	return 401
}

func (o *GetTagUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/tags/{tag_name}][%d] getTagUnauthorized", 401)
}

func (o *GetTagUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/tags/{tag_name}][%d] getTagUnauthorized", 401)
}

func (o *GetTagUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTagNotFound creates a GetTagNotFound with default headers values
func NewGetTagNotFound() *GetTagNotFound {
	return &GetTagNotFound{}
}

/*
GetTagNotFound describes a response with status code 404, with default header values.

not found
*/
type GetTagNotFound struct {
}

// IsSuccess returns true when this get tag not found response has a 2xx status code
func (o *GetTagNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tag not found response has a 3xx status code
func (o *GetTagNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tag not found response has a 4xx status code
func (o *GetTagNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tag not found response has a 5xx status code
func (o *GetTagNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get tag not found response a status code equal to that given
func (o *GetTagNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get tag not found response
func (o *GetTagNotFound) Code() int {
	return 404
}

func (o *GetTagNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/tags/{tag_name}][%d] getTagNotFound", 404)
}

func (o *GetTagNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/tags/{tag_name}][%d] getTagNotFound", 404)
}

func (o *GetTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTagNotAcceptable creates a GetTagNotAcceptable with default headers values
func NewGetTagNotAcceptable() *GetTagNotAcceptable {
	return &GetTagNotAcceptable{}
}

/*
GetTagNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type GetTagNotAcceptable struct {
}

// IsSuccess returns true when this get tag not acceptable response has a 2xx status code
func (o *GetTagNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tag not acceptable response has a 3xx status code
func (o *GetTagNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tag not acceptable response has a 4xx status code
func (o *GetTagNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tag not acceptable response has a 5xx status code
func (o *GetTagNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this get tag not acceptable response a status code equal to that given
func (o *GetTagNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the get tag not acceptable response
func (o *GetTagNotAcceptable) Code() int {
	return 406
}

func (o *GetTagNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/tags/{tag_name}][%d] getTagNotAcceptable", 406)
}

func (o *GetTagNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/tags/{tag_name}][%d] getTagNotAcceptable", 406)
}

func (o *GetTagNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTagInternalServerError creates a GetTagInternalServerError with default headers values
func NewGetTagInternalServerError() *GetTagInternalServerError {
	return &GetTagInternalServerError{}
}

/*
GetTagInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type GetTagInternalServerError struct {
}

// IsSuccess returns true when this get tag internal server error response has a 2xx status code
func (o *GetTagInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tag internal server error response has a 3xx status code
func (o *GetTagInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tag internal server error response has a 4xx status code
func (o *GetTagInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tag internal server error response has a 5xx status code
func (o *GetTagInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get tag internal server error response a status code equal to that given
func (o *GetTagInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get tag internal server error response
func (o *GetTagInternalServerError) Code() int {
	return 500
}

func (o *GetTagInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/tags/{tag_name}][%d] getTagInternalServerError", 500)
}

func (o *GetTagInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/tags/{tag_name}][%d] getTagInternalServerError", 500)
}

func (o *GetTagInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
