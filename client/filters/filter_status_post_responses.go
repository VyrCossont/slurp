// Code generated by go-swagger; DO NOT EDIT.

package filters

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

// FilterStatusPostReader is a Reader for the FilterStatusPost structure.
type FilterStatusPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FilterStatusPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFilterStatusPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFilterStatusPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFilterStatusPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFilterStatusPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFilterStatusPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewFilterStatusPostNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewFilterStatusPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewFilterStatusPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFilterStatusPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v2/filters/{id}/statuses] filterStatusPost", response, response.Code())
	}
}

// NewFilterStatusPostOK creates a FilterStatusPostOK with default headers values
func NewFilterStatusPostOK() *FilterStatusPostOK {
	return &FilterStatusPostOK{}
}

/*
FilterStatusPostOK describes a response with status code 200, with default header values.

New filter status.
*/
type FilterStatusPostOK struct {
	Payload *models.FilterStatus
}

// IsSuccess returns true when this filter status post o k response has a 2xx status code
func (o *FilterStatusPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this filter status post o k response has a 3xx status code
func (o *FilterStatusPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter status post o k response has a 4xx status code
func (o *FilterStatusPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this filter status post o k response has a 5xx status code
func (o *FilterStatusPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this filter status post o k response a status code equal to that given
func (o *FilterStatusPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the filter status post o k response
func (o *FilterStatusPostOK) Code() int {
	return 200
}

func (o *FilterStatusPostOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v2/filters/{id}/statuses][%d] filterStatusPostOK %s", 200, payload)
}

func (o *FilterStatusPostOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v2/filters/{id}/statuses][%d] filterStatusPostOK %s", 200, payload)
}

func (o *FilterStatusPostOK) GetPayload() *models.FilterStatus {
	return o.Payload
}

func (o *FilterStatusPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FilterStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFilterStatusPostBadRequest creates a FilterStatusPostBadRequest with default headers values
func NewFilterStatusPostBadRequest() *FilterStatusPostBadRequest {
	return &FilterStatusPostBadRequest{}
}

/*
FilterStatusPostBadRequest describes a response with status code 400, with default header values.

bad request
*/
type FilterStatusPostBadRequest struct {
}

// IsSuccess returns true when this filter status post bad request response has a 2xx status code
func (o *FilterStatusPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter status post bad request response has a 3xx status code
func (o *FilterStatusPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter status post bad request response has a 4xx status code
func (o *FilterStatusPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter status post bad request response has a 5xx status code
func (o *FilterStatusPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this filter status post bad request response a status code equal to that given
func (o *FilterStatusPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the filter status post bad request response
func (o *FilterStatusPostBadRequest) Code() int {
	return 400
}

func (o *FilterStatusPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/filters/{id}/statuses][%d] filterStatusPostBadRequest", 400)
}

func (o *FilterStatusPostBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/filters/{id}/statuses][%d] filterStatusPostBadRequest", 400)
}

func (o *FilterStatusPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterStatusPostUnauthorized creates a FilterStatusPostUnauthorized with default headers values
func NewFilterStatusPostUnauthorized() *FilterStatusPostUnauthorized {
	return &FilterStatusPostUnauthorized{}
}

/*
FilterStatusPostUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type FilterStatusPostUnauthorized struct {
}

// IsSuccess returns true when this filter status post unauthorized response has a 2xx status code
func (o *FilterStatusPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter status post unauthorized response has a 3xx status code
func (o *FilterStatusPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter status post unauthorized response has a 4xx status code
func (o *FilterStatusPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter status post unauthorized response has a 5xx status code
func (o *FilterStatusPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this filter status post unauthorized response a status code equal to that given
func (o *FilterStatusPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the filter status post unauthorized response
func (o *FilterStatusPostUnauthorized) Code() int {
	return 401
}

func (o *FilterStatusPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/filters/{id}/statuses][%d] filterStatusPostUnauthorized", 401)
}

func (o *FilterStatusPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/filters/{id}/statuses][%d] filterStatusPostUnauthorized", 401)
}

func (o *FilterStatusPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterStatusPostForbidden creates a FilterStatusPostForbidden with default headers values
func NewFilterStatusPostForbidden() *FilterStatusPostForbidden {
	return &FilterStatusPostForbidden{}
}

/*
FilterStatusPostForbidden describes a response with status code 403, with default header values.

forbidden to moved accounts
*/
type FilterStatusPostForbidden struct {
}

// IsSuccess returns true when this filter status post forbidden response has a 2xx status code
func (o *FilterStatusPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter status post forbidden response has a 3xx status code
func (o *FilterStatusPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter status post forbidden response has a 4xx status code
func (o *FilterStatusPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter status post forbidden response has a 5xx status code
func (o *FilterStatusPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this filter status post forbidden response a status code equal to that given
func (o *FilterStatusPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the filter status post forbidden response
func (o *FilterStatusPostForbidden) Code() int {
	return 403
}

func (o *FilterStatusPostForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/filters/{id}/statuses][%d] filterStatusPostForbidden", 403)
}

func (o *FilterStatusPostForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/filters/{id}/statuses][%d] filterStatusPostForbidden", 403)
}

func (o *FilterStatusPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterStatusPostNotFound creates a FilterStatusPostNotFound with default headers values
func NewFilterStatusPostNotFound() *FilterStatusPostNotFound {
	return &FilterStatusPostNotFound{}
}

/*
FilterStatusPostNotFound describes a response with status code 404, with default header values.

not found
*/
type FilterStatusPostNotFound struct {
}

// IsSuccess returns true when this filter status post not found response has a 2xx status code
func (o *FilterStatusPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter status post not found response has a 3xx status code
func (o *FilterStatusPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter status post not found response has a 4xx status code
func (o *FilterStatusPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter status post not found response has a 5xx status code
func (o *FilterStatusPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this filter status post not found response a status code equal to that given
func (o *FilterStatusPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the filter status post not found response
func (o *FilterStatusPostNotFound) Code() int {
	return 404
}

func (o *FilterStatusPostNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/filters/{id}/statuses][%d] filterStatusPostNotFound", 404)
}

func (o *FilterStatusPostNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/filters/{id}/statuses][%d] filterStatusPostNotFound", 404)
}

func (o *FilterStatusPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterStatusPostNotAcceptable creates a FilterStatusPostNotAcceptable with default headers values
func NewFilterStatusPostNotAcceptable() *FilterStatusPostNotAcceptable {
	return &FilterStatusPostNotAcceptable{}
}

/*
FilterStatusPostNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type FilterStatusPostNotAcceptable struct {
}

// IsSuccess returns true when this filter status post not acceptable response has a 2xx status code
func (o *FilterStatusPostNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter status post not acceptable response has a 3xx status code
func (o *FilterStatusPostNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter status post not acceptable response has a 4xx status code
func (o *FilterStatusPostNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter status post not acceptable response has a 5xx status code
func (o *FilterStatusPostNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this filter status post not acceptable response a status code equal to that given
func (o *FilterStatusPostNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the filter status post not acceptable response
func (o *FilterStatusPostNotAcceptable) Code() int {
	return 406
}

func (o *FilterStatusPostNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v2/filters/{id}/statuses][%d] filterStatusPostNotAcceptable", 406)
}

func (o *FilterStatusPostNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v2/filters/{id}/statuses][%d] filterStatusPostNotAcceptable", 406)
}

func (o *FilterStatusPostNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterStatusPostConflict creates a FilterStatusPostConflict with default headers values
func NewFilterStatusPostConflict() *FilterStatusPostConflict {
	return &FilterStatusPostConflict{}
}

/*
FilterStatusPostConflict describes a response with status code 409, with default header values.

conflict (duplicate status)
*/
type FilterStatusPostConflict struct {
}

// IsSuccess returns true when this filter status post conflict response has a 2xx status code
func (o *FilterStatusPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter status post conflict response has a 3xx status code
func (o *FilterStatusPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter status post conflict response has a 4xx status code
func (o *FilterStatusPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter status post conflict response has a 5xx status code
func (o *FilterStatusPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this filter status post conflict response a status code equal to that given
func (o *FilterStatusPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the filter status post conflict response
func (o *FilterStatusPostConflict) Code() int {
	return 409
}

func (o *FilterStatusPostConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/filters/{id}/statuses][%d] filterStatusPostConflict", 409)
}

func (o *FilterStatusPostConflict) String() string {
	return fmt.Sprintf("[POST /api/v2/filters/{id}/statuses][%d] filterStatusPostConflict", 409)
}

func (o *FilterStatusPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterStatusPostUnprocessableEntity creates a FilterStatusPostUnprocessableEntity with default headers values
func NewFilterStatusPostUnprocessableEntity() *FilterStatusPostUnprocessableEntity {
	return &FilterStatusPostUnprocessableEntity{}
}

/*
FilterStatusPostUnprocessableEntity describes a response with status code 422, with default header values.

unprocessable content
*/
type FilterStatusPostUnprocessableEntity struct {
}

// IsSuccess returns true when this filter status post unprocessable entity response has a 2xx status code
func (o *FilterStatusPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter status post unprocessable entity response has a 3xx status code
func (o *FilterStatusPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter status post unprocessable entity response has a 4xx status code
func (o *FilterStatusPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter status post unprocessable entity response has a 5xx status code
func (o *FilterStatusPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this filter status post unprocessable entity response a status code equal to that given
func (o *FilterStatusPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the filter status post unprocessable entity response
func (o *FilterStatusPostUnprocessableEntity) Code() int {
	return 422
}

func (o *FilterStatusPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v2/filters/{id}/statuses][%d] filterStatusPostUnprocessableEntity", 422)
}

func (o *FilterStatusPostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/v2/filters/{id}/statuses][%d] filterStatusPostUnprocessableEntity", 422)
}

func (o *FilterStatusPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterStatusPostInternalServerError creates a FilterStatusPostInternalServerError with default headers values
func NewFilterStatusPostInternalServerError() *FilterStatusPostInternalServerError {
	return &FilterStatusPostInternalServerError{}
}

/*
FilterStatusPostInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type FilterStatusPostInternalServerError struct {
}

// IsSuccess returns true when this filter status post internal server error response has a 2xx status code
func (o *FilterStatusPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter status post internal server error response has a 3xx status code
func (o *FilterStatusPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter status post internal server error response has a 4xx status code
func (o *FilterStatusPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this filter status post internal server error response has a 5xx status code
func (o *FilterStatusPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this filter status post internal server error response a status code equal to that given
func (o *FilterStatusPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the filter status post internal server error response
func (o *FilterStatusPostInternalServerError) Code() int {
	return 500
}

func (o *FilterStatusPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/filters/{id}/statuses][%d] filterStatusPostInternalServerError", 500)
}

func (o *FilterStatusPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/filters/{id}/statuses][%d] filterStatusPostInternalServerError", 500)
}

func (o *FilterStatusPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
