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

// StatusCreateReader is a Reader for the StatusCreate structure.
type StatusCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatusCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStatusCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStatusCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStatusCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStatusCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewStatusCreateNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStatusCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewStatusCreateNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/statuses] statusCreate", response, response.Code())
	}
}

// NewStatusCreateOK creates a StatusCreateOK with default headers values
func NewStatusCreateOK() *StatusCreateOK {
	return &StatusCreateOK{}
}

/*
StatusCreateOK describes a response with status code 200, with default header values.

The newly created status.
*/
type StatusCreateOK struct {
	Payload *models.Status
}

// IsSuccess returns true when this status create o k response has a 2xx status code
func (o *StatusCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this status create o k response has a 3xx status code
func (o *StatusCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status create o k response has a 4xx status code
func (o *StatusCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this status create o k response has a 5xx status code
func (o *StatusCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this status create o k response a status code equal to that given
func (o *StatusCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the status create o k response
func (o *StatusCreateOK) Code() int {
	return 200
}

func (o *StatusCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/statuses][%d] statusCreateOK %s", 200, payload)
}

func (o *StatusCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/statuses][%d] statusCreateOK %s", 200, payload)
}

func (o *StatusCreateOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *StatusCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatusCreateBadRequest creates a StatusCreateBadRequest with default headers values
func NewStatusCreateBadRequest() *StatusCreateBadRequest {
	return &StatusCreateBadRequest{}
}

/*
StatusCreateBadRequest describes a response with status code 400, with default header values.

bad request
*/
type StatusCreateBadRequest struct {
}

// IsSuccess returns true when this status create bad request response has a 2xx status code
func (o *StatusCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status create bad request response has a 3xx status code
func (o *StatusCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status create bad request response has a 4xx status code
func (o *StatusCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this status create bad request response has a 5xx status code
func (o *StatusCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this status create bad request response a status code equal to that given
func (o *StatusCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the status create bad request response
func (o *StatusCreateBadRequest) Code() int {
	return 400
}

func (o *StatusCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/statuses][%d] statusCreateBadRequest", 400)
}

func (o *StatusCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/statuses][%d] statusCreateBadRequest", 400)
}

func (o *StatusCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusCreateUnauthorized creates a StatusCreateUnauthorized with default headers values
func NewStatusCreateUnauthorized() *StatusCreateUnauthorized {
	return &StatusCreateUnauthorized{}
}

/*
StatusCreateUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type StatusCreateUnauthorized struct {
}

// IsSuccess returns true when this status create unauthorized response has a 2xx status code
func (o *StatusCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status create unauthorized response has a 3xx status code
func (o *StatusCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status create unauthorized response has a 4xx status code
func (o *StatusCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this status create unauthorized response has a 5xx status code
func (o *StatusCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this status create unauthorized response a status code equal to that given
func (o *StatusCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the status create unauthorized response
func (o *StatusCreateUnauthorized) Code() int {
	return 401
}

func (o *StatusCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/statuses][%d] statusCreateUnauthorized", 401)
}

func (o *StatusCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/statuses][%d] statusCreateUnauthorized", 401)
}

func (o *StatusCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusCreateForbidden creates a StatusCreateForbidden with default headers values
func NewStatusCreateForbidden() *StatusCreateForbidden {
	return &StatusCreateForbidden{}
}

/*
StatusCreateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type StatusCreateForbidden struct {
}

// IsSuccess returns true when this status create forbidden response has a 2xx status code
func (o *StatusCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status create forbidden response has a 3xx status code
func (o *StatusCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status create forbidden response has a 4xx status code
func (o *StatusCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this status create forbidden response has a 5xx status code
func (o *StatusCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this status create forbidden response a status code equal to that given
func (o *StatusCreateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the status create forbidden response
func (o *StatusCreateForbidden) Code() int {
	return 403
}

func (o *StatusCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/statuses][%d] statusCreateForbidden", 403)
}

func (o *StatusCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/statuses][%d] statusCreateForbidden", 403)
}

func (o *StatusCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusCreateNotFound creates a StatusCreateNotFound with default headers values
func NewStatusCreateNotFound() *StatusCreateNotFound {
	return &StatusCreateNotFound{}
}

/*
StatusCreateNotFound describes a response with status code 404, with default header values.

not found
*/
type StatusCreateNotFound struct {
}

// IsSuccess returns true when this status create not found response has a 2xx status code
func (o *StatusCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status create not found response has a 3xx status code
func (o *StatusCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status create not found response has a 4xx status code
func (o *StatusCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this status create not found response has a 5xx status code
func (o *StatusCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this status create not found response a status code equal to that given
func (o *StatusCreateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the status create not found response
func (o *StatusCreateNotFound) Code() int {
	return 404
}

func (o *StatusCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/statuses][%d] statusCreateNotFound", 404)
}

func (o *StatusCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/statuses][%d] statusCreateNotFound", 404)
}

func (o *StatusCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusCreateNotAcceptable creates a StatusCreateNotAcceptable with default headers values
func NewStatusCreateNotAcceptable() *StatusCreateNotAcceptable {
	return &StatusCreateNotAcceptable{}
}

/*
StatusCreateNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type StatusCreateNotAcceptable struct {
}

// IsSuccess returns true when this status create not acceptable response has a 2xx status code
func (o *StatusCreateNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status create not acceptable response has a 3xx status code
func (o *StatusCreateNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status create not acceptable response has a 4xx status code
func (o *StatusCreateNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this status create not acceptable response has a 5xx status code
func (o *StatusCreateNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this status create not acceptable response a status code equal to that given
func (o *StatusCreateNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the status create not acceptable response
func (o *StatusCreateNotAcceptable) Code() int {
	return 406
}

func (o *StatusCreateNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/statuses][%d] statusCreateNotAcceptable", 406)
}

func (o *StatusCreateNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v1/statuses][%d] statusCreateNotAcceptable", 406)
}

func (o *StatusCreateNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusCreateInternalServerError creates a StatusCreateInternalServerError with default headers values
func NewStatusCreateInternalServerError() *StatusCreateInternalServerError {
	return &StatusCreateInternalServerError{}
}

/*
StatusCreateInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type StatusCreateInternalServerError struct {
}

// IsSuccess returns true when this status create internal server error response has a 2xx status code
func (o *StatusCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status create internal server error response has a 3xx status code
func (o *StatusCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status create internal server error response has a 4xx status code
func (o *StatusCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this status create internal server error response has a 5xx status code
func (o *StatusCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this status create internal server error response a status code equal to that given
func (o *StatusCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the status create internal server error response
func (o *StatusCreateInternalServerError) Code() int {
	return 500
}

func (o *StatusCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/statuses][%d] statusCreateInternalServerError", 500)
}

func (o *StatusCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/statuses][%d] statusCreateInternalServerError", 500)
}

func (o *StatusCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusCreateNotImplemented creates a StatusCreateNotImplemented with default headers values
func NewStatusCreateNotImplemented() *StatusCreateNotImplemented {
	return &StatusCreateNotImplemented{}
}

/*
StatusCreateNotImplemented describes a response with status code 501, with default header values.

scheduled_at was set, but this feature is not yet implemented
*/
type StatusCreateNotImplemented struct {
}

// IsSuccess returns true when this status create not implemented response has a 2xx status code
func (o *StatusCreateNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status create not implemented response has a 3xx status code
func (o *StatusCreateNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status create not implemented response has a 4xx status code
func (o *StatusCreateNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this status create not implemented response has a 5xx status code
func (o *StatusCreateNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this status create not implemented response a status code equal to that given
func (o *StatusCreateNotImplemented) IsCode(code int) bool {
	return code == 501
}

// Code gets the status code for the status create not implemented response
func (o *StatusCreateNotImplemented) Code() int {
	return 501
}

func (o *StatusCreateNotImplemented) Error() string {
	return fmt.Sprintf("[POST /api/v1/statuses][%d] statusCreateNotImplemented", 501)
}

func (o *StatusCreateNotImplemented) String() string {
	return fmt.Sprintf("[POST /api/v1/statuses][%d] statusCreateNotImplemented", 501)
}

func (o *StatusCreateNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
