// Code generated by go-swagger; DO NOT EDIT.

package markers

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

// MarkersGetReader is a Reader for the MarkersGet structure.
type MarkersGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MarkersGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMarkersGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMarkersGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewMarkersGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewMarkersGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/markers] markersGet", response, response.Code())
	}
}

// NewMarkersGetOK creates a MarkersGetOK with default headers values
func NewMarkersGetOK() *MarkersGetOK {
	return &MarkersGetOK{}
}

/*
MarkersGetOK describes a response with status code 200, with default header values.

Requested markers
*/
type MarkersGetOK struct {
	Payload *models.Marker
}

// IsSuccess returns true when this markers get o k response has a 2xx status code
func (o *MarkersGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this markers get o k response has a 3xx status code
func (o *MarkersGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this markers get o k response has a 4xx status code
func (o *MarkersGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this markers get o k response has a 5xx status code
func (o *MarkersGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this markers get o k response a status code equal to that given
func (o *MarkersGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the markers get o k response
func (o *MarkersGetOK) Code() int {
	return 200
}

func (o *MarkersGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/markers][%d] markersGetOK %s", 200, payload)
}

func (o *MarkersGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/markers][%d] markersGetOK %s", 200, payload)
}

func (o *MarkersGetOK) GetPayload() *models.Marker {
	return o.Payload
}

func (o *MarkersGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Marker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMarkersGetBadRequest creates a MarkersGetBadRequest with default headers values
func NewMarkersGetBadRequest() *MarkersGetBadRequest {
	return &MarkersGetBadRequest{}
}

/*
MarkersGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type MarkersGetBadRequest struct {
}

// IsSuccess returns true when this markers get bad request response has a 2xx status code
func (o *MarkersGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this markers get bad request response has a 3xx status code
func (o *MarkersGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this markers get bad request response has a 4xx status code
func (o *MarkersGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this markers get bad request response has a 5xx status code
func (o *MarkersGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this markers get bad request response a status code equal to that given
func (o *MarkersGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the markers get bad request response
func (o *MarkersGetBadRequest) Code() int {
	return 400
}

func (o *MarkersGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/markers][%d] markersGetBadRequest", 400)
}

func (o *MarkersGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/markers][%d] markersGetBadRequest", 400)
}

func (o *MarkersGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMarkersGetUnauthorized creates a MarkersGetUnauthorized with default headers values
func NewMarkersGetUnauthorized() *MarkersGetUnauthorized {
	return &MarkersGetUnauthorized{}
}

/*
MarkersGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type MarkersGetUnauthorized struct {
}

// IsSuccess returns true when this markers get unauthorized response has a 2xx status code
func (o *MarkersGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this markers get unauthorized response has a 3xx status code
func (o *MarkersGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this markers get unauthorized response has a 4xx status code
func (o *MarkersGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this markers get unauthorized response has a 5xx status code
func (o *MarkersGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this markers get unauthorized response a status code equal to that given
func (o *MarkersGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the markers get unauthorized response
func (o *MarkersGetUnauthorized) Code() int {
	return 401
}

func (o *MarkersGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/markers][%d] markersGetUnauthorized", 401)
}

func (o *MarkersGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/markers][%d] markersGetUnauthorized", 401)
}

func (o *MarkersGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMarkersGetInternalServerError creates a MarkersGetInternalServerError with default headers values
func NewMarkersGetInternalServerError() *MarkersGetInternalServerError {
	return &MarkersGetInternalServerError{}
}

/*
MarkersGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type MarkersGetInternalServerError struct {
}

// IsSuccess returns true when this markers get internal server error response has a 2xx status code
func (o *MarkersGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this markers get internal server error response has a 3xx status code
func (o *MarkersGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this markers get internal server error response has a 4xx status code
func (o *MarkersGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this markers get internal server error response has a 5xx status code
func (o *MarkersGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this markers get internal server error response a status code equal to that given
func (o *MarkersGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the markers get internal server error response
func (o *MarkersGetInternalServerError) Code() int {
	return 500
}

func (o *MarkersGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/markers][%d] markersGetInternalServerError", 500)
}

func (o *MarkersGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/markers][%d] markersGetInternalServerError", 500)
}

func (o *MarkersGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
