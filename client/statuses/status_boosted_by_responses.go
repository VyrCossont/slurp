// Code generated by go-swagger; DO NOT EDIT.

package statuses

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

// StatusBoostedByReader is a Reader for the StatusBoostedBy structure.
type StatusBoostedByReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusBoostedByReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatusBoostedByOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStatusBoostedByBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStatusBoostedByUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStatusBoostedByForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStatusBoostedByNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/statuses/{id}/reblogged_by] statusBoostedBy", response, response.Code())
	}
}

// NewStatusBoostedByOK creates a StatusBoostedByOK with default headers values
func NewStatusBoostedByOK() *StatusBoostedByOK {
	return &StatusBoostedByOK{}
}

/*
StatusBoostedByOK describes a response with status code 200, with default header values.

StatusBoostedByOK status boosted by o k
*/
type StatusBoostedByOK struct {
	Payload []*models.Account
}

// IsSuccess returns true when this status boosted by o k response has a 2xx status code
func (o *StatusBoostedByOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this status boosted by o k response has a 3xx status code
func (o *StatusBoostedByOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status boosted by o k response has a 4xx status code
func (o *StatusBoostedByOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this status boosted by o k response has a 5xx status code
func (o *StatusBoostedByOK) IsServerError() bool {
	return false
}

// IsCode returns true when this status boosted by o k response a status code equal to that given
func (o *StatusBoostedByOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the status boosted by o k response
func (o *StatusBoostedByOK) Code() int {
	return 200
}

func (o *StatusBoostedByOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/reblogged_by][%d] statusBoostedByOK %s", 200, payload)
}

func (o *StatusBoostedByOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/reblogged_by][%d] statusBoostedByOK %s", 200, payload)
}

func (o *StatusBoostedByOK) GetPayload() []*models.Account {
	return o.Payload
}

func (o *StatusBoostedByOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatusBoostedByBadRequest creates a StatusBoostedByBadRequest with default headers values
func NewStatusBoostedByBadRequest() *StatusBoostedByBadRequest {
	return &StatusBoostedByBadRequest{}
}

/*
StatusBoostedByBadRequest describes a response with status code 400, with default header values.

bad request
*/
type StatusBoostedByBadRequest struct {
}

// IsSuccess returns true when this status boosted by bad request response has a 2xx status code
func (o *StatusBoostedByBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status boosted by bad request response has a 3xx status code
func (o *StatusBoostedByBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status boosted by bad request response has a 4xx status code
func (o *StatusBoostedByBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this status boosted by bad request response has a 5xx status code
func (o *StatusBoostedByBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this status boosted by bad request response a status code equal to that given
func (o *StatusBoostedByBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the status boosted by bad request response
func (o *StatusBoostedByBadRequest) Code() int {
	return 400
}

func (o *StatusBoostedByBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/reblogged_by][%d] statusBoostedByBadRequest", 400)
}

func (o *StatusBoostedByBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/reblogged_by][%d] statusBoostedByBadRequest", 400)
}

func (o *StatusBoostedByBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusBoostedByUnauthorized creates a StatusBoostedByUnauthorized with default headers values
func NewStatusBoostedByUnauthorized() *StatusBoostedByUnauthorized {
	return &StatusBoostedByUnauthorized{}
}

/*
StatusBoostedByUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type StatusBoostedByUnauthorized struct {
}

// IsSuccess returns true when this status boosted by unauthorized response has a 2xx status code
func (o *StatusBoostedByUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status boosted by unauthorized response has a 3xx status code
func (o *StatusBoostedByUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status boosted by unauthorized response has a 4xx status code
func (o *StatusBoostedByUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this status boosted by unauthorized response has a 5xx status code
func (o *StatusBoostedByUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this status boosted by unauthorized response a status code equal to that given
func (o *StatusBoostedByUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the status boosted by unauthorized response
func (o *StatusBoostedByUnauthorized) Code() int {
	return 401
}

func (o *StatusBoostedByUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/reblogged_by][%d] statusBoostedByUnauthorized", 401)
}

func (o *StatusBoostedByUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/reblogged_by][%d] statusBoostedByUnauthorized", 401)
}

func (o *StatusBoostedByUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusBoostedByForbidden creates a StatusBoostedByForbidden with default headers values
func NewStatusBoostedByForbidden() *StatusBoostedByForbidden {
	return &StatusBoostedByForbidden{}
}

/*
StatusBoostedByForbidden describes a response with status code 403, with default header values.

forbidden
*/
type StatusBoostedByForbidden struct {
}

// IsSuccess returns true when this status boosted by forbidden response has a 2xx status code
func (o *StatusBoostedByForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status boosted by forbidden response has a 3xx status code
func (o *StatusBoostedByForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status boosted by forbidden response has a 4xx status code
func (o *StatusBoostedByForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this status boosted by forbidden response has a 5xx status code
func (o *StatusBoostedByForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this status boosted by forbidden response a status code equal to that given
func (o *StatusBoostedByForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the status boosted by forbidden response
func (o *StatusBoostedByForbidden) Code() int {
	return 403
}

func (o *StatusBoostedByForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/reblogged_by][%d] statusBoostedByForbidden", 403)
}

func (o *StatusBoostedByForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/reblogged_by][%d] statusBoostedByForbidden", 403)
}

func (o *StatusBoostedByForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusBoostedByNotFound creates a StatusBoostedByNotFound with default headers values
func NewStatusBoostedByNotFound() *StatusBoostedByNotFound {
	return &StatusBoostedByNotFound{}
}

/*
StatusBoostedByNotFound describes a response with status code 404, with default header values.

not found
*/
type StatusBoostedByNotFound struct {
}

// IsSuccess returns true when this status boosted by not found response has a 2xx status code
func (o *StatusBoostedByNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status boosted by not found response has a 3xx status code
func (o *StatusBoostedByNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status boosted by not found response has a 4xx status code
func (o *StatusBoostedByNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this status boosted by not found response has a 5xx status code
func (o *StatusBoostedByNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this status boosted by not found response a status code equal to that given
func (o *StatusBoostedByNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the status boosted by not found response
func (o *StatusBoostedByNotFound) Code() int {
	return 404
}

func (o *StatusBoostedByNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/reblogged_by][%d] statusBoostedByNotFound", 404)
}

func (o *StatusBoostedByNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}/reblogged_by][%d] statusBoostedByNotFound", 404)
}

func (o *StatusBoostedByNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
