// Code generated by go-swagger; DO NOT EDIT.

package federation

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

// S2sFeaturedCollectionGetReader is a Reader for the S2sFeaturedCollectionGet structure.
type S2sFeaturedCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S2sFeaturedCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS2sFeaturedCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewS2sFeaturedCollectionGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewS2sFeaturedCollectionGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewS2sFeaturedCollectionGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewS2sFeaturedCollectionGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /users/{username}/collections/featured] s2sFeaturedCollectionGet", response, response.Code())
	}
}

// NewS2sFeaturedCollectionGetOK creates a S2sFeaturedCollectionGetOK with default headers values
func NewS2sFeaturedCollectionGetOK() *S2sFeaturedCollectionGetOK {
	return &S2sFeaturedCollectionGetOK{}
}

/*
S2sFeaturedCollectionGetOK describes a response with status code 200, with default header values.

S2sFeaturedCollectionGetOK s2s featured collection get o k
*/
type S2sFeaturedCollectionGetOK struct {
	Payload *models.SwaggerFeaturedCollection
}

// IsSuccess returns true when this s2s featured collection get o k response has a 2xx status code
func (o *S2sFeaturedCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s2s featured collection get o k response has a 3xx status code
func (o *S2sFeaturedCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s2s featured collection get o k response has a 4xx status code
func (o *S2sFeaturedCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s2s featured collection get o k response has a 5xx status code
func (o *S2sFeaturedCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s2s featured collection get o k response a status code equal to that given
func (o *S2sFeaturedCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the s2s featured collection get o k response
func (o *S2sFeaturedCollectionGetOK) Code() int {
	return 200
}

func (o *S2sFeaturedCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{username}/collections/featured][%d] s2sFeaturedCollectionGetOK %s", 200, payload)
}

func (o *S2sFeaturedCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{username}/collections/featured][%d] s2sFeaturedCollectionGetOK %s", 200, payload)
}

func (o *S2sFeaturedCollectionGetOK) GetPayload() *models.SwaggerFeaturedCollection {
	return o.Payload
}

func (o *S2sFeaturedCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SwaggerFeaturedCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS2sFeaturedCollectionGetBadRequest creates a S2sFeaturedCollectionGetBadRequest with default headers values
func NewS2sFeaturedCollectionGetBadRequest() *S2sFeaturedCollectionGetBadRequest {
	return &S2sFeaturedCollectionGetBadRequest{}
}

/*
S2sFeaturedCollectionGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type S2sFeaturedCollectionGetBadRequest struct {
}

// IsSuccess returns true when this s2s featured collection get bad request response has a 2xx status code
func (o *S2sFeaturedCollectionGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s2s featured collection get bad request response has a 3xx status code
func (o *S2sFeaturedCollectionGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s2s featured collection get bad request response has a 4xx status code
func (o *S2sFeaturedCollectionGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this s2s featured collection get bad request response has a 5xx status code
func (o *S2sFeaturedCollectionGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this s2s featured collection get bad request response a status code equal to that given
func (o *S2sFeaturedCollectionGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the s2s featured collection get bad request response
func (o *S2sFeaturedCollectionGetBadRequest) Code() int {
	return 400
}

func (o *S2sFeaturedCollectionGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{username}/collections/featured][%d] s2sFeaturedCollectionGetBadRequest", 400)
}

func (o *S2sFeaturedCollectionGetBadRequest) String() string {
	return fmt.Sprintf("[GET /users/{username}/collections/featured][%d] s2sFeaturedCollectionGetBadRequest", 400)
}

func (o *S2sFeaturedCollectionGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewS2sFeaturedCollectionGetUnauthorized creates a S2sFeaturedCollectionGetUnauthorized with default headers values
func NewS2sFeaturedCollectionGetUnauthorized() *S2sFeaturedCollectionGetUnauthorized {
	return &S2sFeaturedCollectionGetUnauthorized{}
}

/*
S2sFeaturedCollectionGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type S2sFeaturedCollectionGetUnauthorized struct {
}

// IsSuccess returns true when this s2s featured collection get unauthorized response has a 2xx status code
func (o *S2sFeaturedCollectionGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s2s featured collection get unauthorized response has a 3xx status code
func (o *S2sFeaturedCollectionGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s2s featured collection get unauthorized response has a 4xx status code
func (o *S2sFeaturedCollectionGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this s2s featured collection get unauthorized response has a 5xx status code
func (o *S2sFeaturedCollectionGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this s2s featured collection get unauthorized response a status code equal to that given
func (o *S2sFeaturedCollectionGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the s2s featured collection get unauthorized response
func (o *S2sFeaturedCollectionGetUnauthorized) Code() int {
	return 401
}

func (o *S2sFeaturedCollectionGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{username}/collections/featured][%d] s2sFeaturedCollectionGetUnauthorized", 401)
}

func (o *S2sFeaturedCollectionGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /users/{username}/collections/featured][%d] s2sFeaturedCollectionGetUnauthorized", 401)
}

func (o *S2sFeaturedCollectionGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewS2sFeaturedCollectionGetForbidden creates a S2sFeaturedCollectionGetForbidden with default headers values
func NewS2sFeaturedCollectionGetForbidden() *S2sFeaturedCollectionGetForbidden {
	return &S2sFeaturedCollectionGetForbidden{}
}

/*
S2sFeaturedCollectionGetForbidden describes a response with status code 403, with default header values.

forbidden
*/
type S2sFeaturedCollectionGetForbidden struct {
}

// IsSuccess returns true when this s2s featured collection get forbidden response has a 2xx status code
func (o *S2sFeaturedCollectionGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s2s featured collection get forbidden response has a 3xx status code
func (o *S2sFeaturedCollectionGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s2s featured collection get forbidden response has a 4xx status code
func (o *S2sFeaturedCollectionGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this s2s featured collection get forbidden response has a 5xx status code
func (o *S2sFeaturedCollectionGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this s2s featured collection get forbidden response a status code equal to that given
func (o *S2sFeaturedCollectionGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the s2s featured collection get forbidden response
func (o *S2sFeaturedCollectionGetForbidden) Code() int {
	return 403
}

func (o *S2sFeaturedCollectionGetForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{username}/collections/featured][%d] s2sFeaturedCollectionGetForbidden", 403)
}

func (o *S2sFeaturedCollectionGetForbidden) String() string {
	return fmt.Sprintf("[GET /users/{username}/collections/featured][%d] s2sFeaturedCollectionGetForbidden", 403)
}

func (o *S2sFeaturedCollectionGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewS2sFeaturedCollectionGetNotFound creates a S2sFeaturedCollectionGetNotFound with default headers values
func NewS2sFeaturedCollectionGetNotFound() *S2sFeaturedCollectionGetNotFound {
	return &S2sFeaturedCollectionGetNotFound{}
}

/*
S2sFeaturedCollectionGetNotFound describes a response with status code 404, with default header values.

not found
*/
type S2sFeaturedCollectionGetNotFound struct {
}

// IsSuccess returns true when this s2s featured collection get not found response has a 2xx status code
func (o *S2sFeaturedCollectionGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s2s featured collection get not found response has a 3xx status code
func (o *S2sFeaturedCollectionGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s2s featured collection get not found response has a 4xx status code
func (o *S2sFeaturedCollectionGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this s2s featured collection get not found response has a 5xx status code
func (o *S2sFeaturedCollectionGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this s2s featured collection get not found response a status code equal to that given
func (o *S2sFeaturedCollectionGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the s2s featured collection get not found response
func (o *S2sFeaturedCollectionGetNotFound) Code() int {
	return 404
}

func (o *S2sFeaturedCollectionGetNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{username}/collections/featured][%d] s2sFeaturedCollectionGetNotFound", 404)
}

func (o *S2sFeaturedCollectionGetNotFound) String() string {
	return fmt.Sprintf("[GET /users/{username}/collections/featured][%d] s2sFeaturedCollectionGetNotFound", 404)
}

func (o *S2sFeaturedCollectionGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
