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

// FilterV1PostReader is a Reader for the FilterV1Post structure.
type FilterV1PostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FilterV1PostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFilterV1PostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFilterV1PostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFilterV1PostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFilterV1PostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFilterV1PostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewFilterV1PostNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewFilterV1PostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewFilterV1PostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFilterV1PostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/filters] filterV1Post", response, response.Code())
	}
}

// NewFilterV1PostOK creates a FilterV1PostOK with default headers values
func NewFilterV1PostOK() *FilterV1PostOK {
	return &FilterV1PostOK{}
}

/*
FilterV1PostOK describes a response with status code 200, with default header values.

New filter.
*/
type FilterV1PostOK struct {
	Payload *models.FilterV1
}

// IsSuccess returns true when this filter v1 post o k response has a 2xx status code
func (o *FilterV1PostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this filter v1 post o k response has a 3xx status code
func (o *FilterV1PostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v1 post o k response has a 4xx status code
func (o *FilterV1PostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this filter v1 post o k response has a 5xx status code
func (o *FilterV1PostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this filter v1 post o k response a status code equal to that given
func (o *FilterV1PostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the filter v1 post o k response
func (o *FilterV1PostOK) Code() int {
	return 200
}

func (o *FilterV1PostOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/filters][%d] filterV1PostOK %s", 200, payload)
}

func (o *FilterV1PostOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/filters][%d] filterV1PostOK %s", 200, payload)
}

func (o *FilterV1PostOK) GetPayload() *models.FilterV1 {
	return o.Payload
}

func (o *FilterV1PostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FilterV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFilterV1PostBadRequest creates a FilterV1PostBadRequest with default headers values
func NewFilterV1PostBadRequest() *FilterV1PostBadRequest {
	return &FilterV1PostBadRequest{}
}

/*
FilterV1PostBadRequest describes a response with status code 400, with default header values.

bad request
*/
type FilterV1PostBadRequest struct {
}

// IsSuccess returns true when this filter v1 post bad request response has a 2xx status code
func (o *FilterV1PostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter v1 post bad request response has a 3xx status code
func (o *FilterV1PostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v1 post bad request response has a 4xx status code
func (o *FilterV1PostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter v1 post bad request response has a 5xx status code
func (o *FilterV1PostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this filter v1 post bad request response a status code equal to that given
func (o *FilterV1PostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the filter v1 post bad request response
func (o *FilterV1PostBadRequest) Code() int {
	return 400
}

func (o *FilterV1PostBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/filters][%d] filterV1PostBadRequest", 400)
}

func (o *FilterV1PostBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/filters][%d] filterV1PostBadRequest", 400)
}

func (o *FilterV1PostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterV1PostUnauthorized creates a FilterV1PostUnauthorized with default headers values
func NewFilterV1PostUnauthorized() *FilterV1PostUnauthorized {
	return &FilterV1PostUnauthorized{}
}

/*
FilterV1PostUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type FilterV1PostUnauthorized struct {
}

// IsSuccess returns true when this filter v1 post unauthorized response has a 2xx status code
func (o *FilterV1PostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter v1 post unauthorized response has a 3xx status code
func (o *FilterV1PostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v1 post unauthorized response has a 4xx status code
func (o *FilterV1PostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter v1 post unauthorized response has a 5xx status code
func (o *FilterV1PostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this filter v1 post unauthorized response a status code equal to that given
func (o *FilterV1PostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the filter v1 post unauthorized response
func (o *FilterV1PostUnauthorized) Code() int {
	return 401
}

func (o *FilterV1PostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/filters][%d] filterV1PostUnauthorized", 401)
}

func (o *FilterV1PostUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/filters][%d] filterV1PostUnauthorized", 401)
}

func (o *FilterV1PostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterV1PostForbidden creates a FilterV1PostForbidden with default headers values
func NewFilterV1PostForbidden() *FilterV1PostForbidden {
	return &FilterV1PostForbidden{}
}

/*
FilterV1PostForbidden describes a response with status code 403, with default header values.

forbidden to moved accounts
*/
type FilterV1PostForbidden struct {
}

// IsSuccess returns true when this filter v1 post forbidden response has a 2xx status code
func (o *FilterV1PostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter v1 post forbidden response has a 3xx status code
func (o *FilterV1PostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v1 post forbidden response has a 4xx status code
func (o *FilterV1PostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter v1 post forbidden response has a 5xx status code
func (o *FilterV1PostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this filter v1 post forbidden response a status code equal to that given
func (o *FilterV1PostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the filter v1 post forbidden response
func (o *FilterV1PostForbidden) Code() int {
	return 403
}

func (o *FilterV1PostForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/filters][%d] filterV1PostForbidden", 403)
}

func (o *FilterV1PostForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/filters][%d] filterV1PostForbidden", 403)
}

func (o *FilterV1PostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterV1PostNotFound creates a FilterV1PostNotFound with default headers values
func NewFilterV1PostNotFound() *FilterV1PostNotFound {
	return &FilterV1PostNotFound{}
}

/*
FilterV1PostNotFound describes a response with status code 404, with default header values.

not found
*/
type FilterV1PostNotFound struct {
}

// IsSuccess returns true when this filter v1 post not found response has a 2xx status code
func (o *FilterV1PostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter v1 post not found response has a 3xx status code
func (o *FilterV1PostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v1 post not found response has a 4xx status code
func (o *FilterV1PostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter v1 post not found response has a 5xx status code
func (o *FilterV1PostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this filter v1 post not found response a status code equal to that given
func (o *FilterV1PostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the filter v1 post not found response
func (o *FilterV1PostNotFound) Code() int {
	return 404
}

func (o *FilterV1PostNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/filters][%d] filterV1PostNotFound", 404)
}

func (o *FilterV1PostNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/filters][%d] filterV1PostNotFound", 404)
}

func (o *FilterV1PostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterV1PostNotAcceptable creates a FilterV1PostNotAcceptable with default headers values
func NewFilterV1PostNotAcceptable() *FilterV1PostNotAcceptable {
	return &FilterV1PostNotAcceptable{}
}

/*
FilterV1PostNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type FilterV1PostNotAcceptable struct {
}

// IsSuccess returns true when this filter v1 post not acceptable response has a 2xx status code
func (o *FilterV1PostNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter v1 post not acceptable response has a 3xx status code
func (o *FilterV1PostNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v1 post not acceptable response has a 4xx status code
func (o *FilterV1PostNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter v1 post not acceptable response has a 5xx status code
func (o *FilterV1PostNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this filter v1 post not acceptable response a status code equal to that given
func (o *FilterV1PostNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the filter v1 post not acceptable response
func (o *FilterV1PostNotAcceptable) Code() int {
	return 406
}

func (o *FilterV1PostNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/filters][%d] filterV1PostNotAcceptable", 406)
}

func (o *FilterV1PostNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v1/filters][%d] filterV1PostNotAcceptable", 406)
}

func (o *FilterV1PostNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterV1PostConflict creates a FilterV1PostConflict with default headers values
func NewFilterV1PostConflict() *FilterV1PostConflict {
	return &FilterV1PostConflict{}
}

/*
FilterV1PostConflict describes a response with status code 409, with default header values.

conflict (duplicate keyword)
*/
type FilterV1PostConflict struct {
}

// IsSuccess returns true when this filter v1 post conflict response has a 2xx status code
func (o *FilterV1PostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter v1 post conflict response has a 3xx status code
func (o *FilterV1PostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v1 post conflict response has a 4xx status code
func (o *FilterV1PostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter v1 post conflict response has a 5xx status code
func (o *FilterV1PostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this filter v1 post conflict response a status code equal to that given
func (o *FilterV1PostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the filter v1 post conflict response
func (o *FilterV1PostConflict) Code() int {
	return 409
}

func (o *FilterV1PostConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/filters][%d] filterV1PostConflict", 409)
}

func (o *FilterV1PostConflict) String() string {
	return fmt.Sprintf("[POST /api/v1/filters][%d] filterV1PostConflict", 409)
}

func (o *FilterV1PostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterV1PostUnprocessableEntity creates a FilterV1PostUnprocessableEntity with default headers values
func NewFilterV1PostUnprocessableEntity() *FilterV1PostUnprocessableEntity {
	return &FilterV1PostUnprocessableEntity{}
}

/*
FilterV1PostUnprocessableEntity describes a response with status code 422, with default header values.

unprocessable content
*/
type FilterV1PostUnprocessableEntity struct {
}

// IsSuccess returns true when this filter v1 post unprocessable entity response has a 2xx status code
func (o *FilterV1PostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter v1 post unprocessable entity response has a 3xx status code
func (o *FilterV1PostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v1 post unprocessable entity response has a 4xx status code
func (o *FilterV1PostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter v1 post unprocessable entity response has a 5xx status code
func (o *FilterV1PostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this filter v1 post unprocessable entity response a status code equal to that given
func (o *FilterV1PostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the filter v1 post unprocessable entity response
func (o *FilterV1PostUnprocessableEntity) Code() int {
	return 422
}

func (o *FilterV1PostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v1/filters][%d] filterV1PostUnprocessableEntity", 422)
}

func (o *FilterV1PostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/v1/filters][%d] filterV1PostUnprocessableEntity", 422)
}

func (o *FilterV1PostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterV1PostInternalServerError creates a FilterV1PostInternalServerError with default headers values
func NewFilterV1PostInternalServerError() *FilterV1PostInternalServerError {
	return &FilterV1PostInternalServerError{}
}

/*
FilterV1PostInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type FilterV1PostInternalServerError struct {
}

// IsSuccess returns true when this filter v1 post internal server error response has a 2xx status code
func (o *FilterV1PostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter v1 post internal server error response has a 3xx status code
func (o *FilterV1PostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v1 post internal server error response has a 4xx status code
func (o *FilterV1PostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this filter v1 post internal server error response has a 5xx status code
func (o *FilterV1PostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this filter v1 post internal server error response a status code equal to that given
func (o *FilterV1PostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the filter v1 post internal server error response
func (o *FilterV1PostInternalServerError) Code() int {
	return 500
}

func (o *FilterV1PostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/filters][%d] filterV1PostInternalServerError", 500)
}

func (o *FilterV1PostInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/filters][%d] filterV1PostInternalServerError", 500)
}

func (o *FilterV1PostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
