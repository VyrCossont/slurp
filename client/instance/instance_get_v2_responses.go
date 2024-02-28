// Code generated by go-swagger; DO NOT EDIT.

package instance

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

// InstanceGetV2Reader is a Reader for the InstanceGetV2 structure.
type InstanceGetV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstanceGetV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInstanceGetV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 406:
		result := NewInstanceGetV2NotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInstanceGetV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v2/instance] instanceGetV2", response, response.Code())
	}
}

// NewInstanceGetV2OK creates a InstanceGetV2OK with default headers values
func NewInstanceGetV2OK() *InstanceGetV2OK {
	return &InstanceGetV2OK{}
}

/*
InstanceGetV2OK describes a response with status code 200, with default header values.

Instance information.
*/
type InstanceGetV2OK struct {
	Payload *models.InstanceV2
}

// IsSuccess returns true when this instance get v2 o k response has a 2xx status code
func (o *InstanceGetV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this instance get v2 o k response has a 3xx status code
func (o *InstanceGetV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instance get v2 o k response has a 4xx status code
func (o *InstanceGetV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this instance get v2 o k response has a 5xx status code
func (o *InstanceGetV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this instance get v2 o k response a status code equal to that given
func (o *InstanceGetV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the instance get v2 o k response
func (o *InstanceGetV2OK) Code() int {
	return 200
}

func (o *InstanceGetV2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/instance][%d] instanceGetV2OK %s", 200, payload)
}

func (o *InstanceGetV2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/instance][%d] instanceGetV2OK %s", 200, payload)
}

func (o *InstanceGetV2OK) GetPayload() *models.InstanceV2 {
	return o.Payload
}

func (o *InstanceGetV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InstanceV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstanceGetV2NotAcceptable creates a InstanceGetV2NotAcceptable with default headers values
func NewInstanceGetV2NotAcceptable() *InstanceGetV2NotAcceptable {
	return &InstanceGetV2NotAcceptable{}
}

/*
InstanceGetV2NotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type InstanceGetV2NotAcceptable struct {
}

// IsSuccess returns true when this instance get v2 not acceptable response has a 2xx status code
func (o *InstanceGetV2NotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this instance get v2 not acceptable response has a 3xx status code
func (o *InstanceGetV2NotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instance get v2 not acceptable response has a 4xx status code
func (o *InstanceGetV2NotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this instance get v2 not acceptable response has a 5xx status code
func (o *InstanceGetV2NotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this instance get v2 not acceptable response a status code equal to that given
func (o *InstanceGetV2NotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the instance get v2 not acceptable response
func (o *InstanceGetV2NotAcceptable) Code() int {
	return 406
}

func (o *InstanceGetV2NotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v2/instance][%d] instanceGetV2NotAcceptable", 406)
}

func (o *InstanceGetV2NotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v2/instance][%d] instanceGetV2NotAcceptable", 406)
}

func (o *InstanceGetV2NotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstanceGetV2InternalServerError creates a InstanceGetV2InternalServerError with default headers values
func NewInstanceGetV2InternalServerError() *InstanceGetV2InternalServerError {
	return &InstanceGetV2InternalServerError{}
}

/*
InstanceGetV2InternalServerError describes a response with status code 500, with default header values.

internal error
*/
type InstanceGetV2InternalServerError struct {
}

// IsSuccess returns true when this instance get v2 internal server error response has a 2xx status code
func (o *InstanceGetV2InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this instance get v2 internal server error response has a 3xx status code
func (o *InstanceGetV2InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instance get v2 internal server error response has a 4xx status code
func (o *InstanceGetV2InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this instance get v2 internal server error response has a 5xx status code
func (o *InstanceGetV2InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this instance get v2 internal server error response a status code equal to that given
func (o *InstanceGetV2InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the instance get v2 internal server error response
func (o *InstanceGetV2InternalServerError) Code() int {
	return 500
}

func (o *InstanceGetV2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/instance][%d] instanceGetV2InternalServerError", 500)
}

func (o *InstanceGetV2InternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/instance][%d] instanceGetV2InternalServerError", 500)
}

func (o *InstanceGetV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
