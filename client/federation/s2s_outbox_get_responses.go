// Code generated by go-swagger; DO NOT EDIT.

package federation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VyrCossont/slurp/models"
)

// S2sOutboxGetReader is a Reader for the S2sOutboxGet structure.
type S2sOutboxGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S2sOutboxGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS2sOutboxGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewS2sOutboxGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewS2sOutboxGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewS2sOutboxGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewS2sOutboxGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /users/{username}/outbox] s2sOutboxGet", response, response.Code())
	}
}

// NewS2sOutboxGetOK creates a S2sOutboxGetOK with default headers values
func NewS2sOutboxGetOK() *S2sOutboxGetOK {
	return &S2sOutboxGetOK{}
}

/*
S2sOutboxGetOK describes a response with status code 200, with default header values.

S2sOutboxGetOK s2s outbox get o k
*/
type S2sOutboxGetOK struct {
	Payload *models.SwaggerCollection
}

// IsSuccess returns true when this s2s outbox get o k response has a 2xx status code
func (o *S2sOutboxGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s2s outbox get o k response has a 3xx status code
func (o *S2sOutboxGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s2s outbox get o k response has a 4xx status code
func (o *S2sOutboxGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s2s outbox get o k response has a 5xx status code
func (o *S2sOutboxGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s2s outbox get o k response a status code equal to that given
func (o *S2sOutboxGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the s2s outbox get o k response
func (o *S2sOutboxGetOK) Code() int {
	return 200
}

func (o *S2sOutboxGetOK) Error() string {
	return fmt.Sprintf("[GET /users/{username}/outbox][%d] s2sOutboxGetOK  %+v", 200, o.Payload)
}

func (o *S2sOutboxGetOK) String() string {
	return fmt.Sprintf("[GET /users/{username}/outbox][%d] s2sOutboxGetOK  %+v", 200, o.Payload)
}

func (o *S2sOutboxGetOK) GetPayload() *models.SwaggerCollection {
	return o.Payload
}

func (o *S2sOutboxGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SwaggerCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS2sOutboxGetBadRequest creates a S2sOutboxGetBadRequest with default headers values
func NewS2sOutboxGetBadRequest() *S2sOutboxGetBadRequest {
	return &S2sOutboxGetBadRequest{}
}

/*
S2sOutboxGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type S2sOutboxGetBadRequest struct {
}

// IsSuccess returns true when this s2s outbox get bad request response has a 2xx status code
func (o *S2sOutboxGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s2s outbox get bad request response has a 3xx status code
func (o *S2sOutboxGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s2s outbox get bad request response has a 4xx status code
func (o *S2sOutboxGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this s2s outbox get bad request response has a 5xx status code
func (o *S2sOutboxGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this s2s outbox get bad request response a status code equal to that given
func (o *S2sOutboxGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the s2s outbox get bad request response
func (o *S2sOutboxGetBadRequest) Code() int {
	return 400
}

func (o *S2sOutboxGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{username}/outbox][%d] s2sOutboxGetBadRequest ", 400)
}

func (o *S2sOutboxGetBadRequest) String() string {
	return fmt.Sprintf("[GET /users/{username}/outbox][%d] s2sOutboxGetBadRequest ", 400)
}

func (o *S2sOutboxGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewS2sOutboxGetUnauthorized creates a S2sOutboxGetUnauthorized with default headers values
func NewS2sOutboxGetUnauthorized() *S2sOutboxGetUnauthorized {
	return &S2sOutboxGetUnauthorized{}
}

/*
S2sOutboxGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type S2sOutboxGetUnauthorized struct {
}

// IsSuccess returns true when this s2s outbox get unauthorized response has a 2xx status code
func (o *S2sOutboxGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s2s outbox get unauthorized response has a 3xx status code
func (o *S2sOutboxGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s2s outbox get unauthorized response has a 4xx status code
func (o *S2sOutboxGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this s2s outbox get unauthorized response has a 5xx status code
func (o *S2sOutboxGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this s2s outbox get unauthorized response a status code equal to that given
func (o *S2sOutboxGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the s2s outbox get unauthorized response
func (o *S2sOutboxGetUnauthorized) Code() int {
	return 401
}

func (o *S2sOutboxGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{username}/outbox][%d] s2sOutboxGetUnauthorized ", 401)
}

func (o *S2sOutboxGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /users/{username}/outbox][%d] s2sOutboxGetUnauthorized ", 401)
}

func (o *S2sOutboxGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewS2sOutboxGetForbidden creates a S2sOutboxGetForbidden with default headers values
func NewS2sOutboxGetForbidden() *S2sOutboxGetForbidden {
	return &S2sOutboxGetForbidden{}
}

/*
S2sOutboxGetForbidden describes a response with status code 403, with default header values.

forbidden
*/
type S2sOutboxGetForbidden struct {
}

// IsSuccess returns true when this s2s outbox get forbidden response has a 2xx status code
func (o *S2sOutboxGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s2s outbox get forbidden response has a 3xx status code
func (o *S2sOutboxGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s2s outbox get forbidden response has a 4xx status code
func (o *S2sOutboxGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this s2s outbox get forbidden response has a 5xx status code
func (o *S2sOutboxGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this s2s outbox get forbidden response a status code equal to that given
func (o *S2sOutboxGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the s2s outbox get forbidden response
func (o *S2sOutboxGetForbidden) Code() int {
	return 403
}

func (o *S2sOutboxGetForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{username}/outbox][%d] s2sOutboxGetForbidden ", 403)
}

func (o *S2sOutboxGetForbidden) String() string {
	return fmt.Sprintf("[GET /users/{username}/outbox][%d] s2sOutboxGetForbidden ", 403)
}

func (o *S2sOutboxGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewS2sOutboxGetNotFound creates a S2sOutboxGetNotFound with default headers values
func NewS2sOutboxGetNotFound() *S2sOutboxGetNotFound {
	return &S2sOutboxGetNotFound{}
}

/*
S2sOutboxGetNotFound describes a response with status code 404, with default header values.

not found
*/
type S2sOutboxGetNotFound struct {
}

// IsSuccess returns true when this s2s outbox get not found response has a 2xx status code
func (o *S2sOutboxGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s2s outbox get not found response has a 3xx status code
func (o *S2sOutboxGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s2s outbox get not found response has a 4xx status code
func (o *S2sOutboxGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this s2s outbox get not found response has a 5xx status code
func (o *S2sOutboxGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this s2s outbox get not found response a status code equal to that given
func (o *S2sOutboxGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the s2s outbox get not found response
func (o *S2sOutboxGetNotFound) Code() int {
	return 404
}

func (o *S2sOutboxGetNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{username}/outbox][%d] s2sOutboxGetNotFound ", 404)
}

func (o *S2sOutboxGetNotFound) String() string {
	return fmt.Sprintf("[GET /users/{username}/outbox][%d] s2sOutboxGetNotFound ", 404)
}

func (o *S2sOutboxGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
