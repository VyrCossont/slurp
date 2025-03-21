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

// FilterV2PutReader is a Reader for the FilterV2Put structure.
type FilterV2PutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FilterV2PutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFilterV2PutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFilterV2PutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFilterV2PutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFilterV2PutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFilterV2PutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewFilterV2PutNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewFilterV2PutConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewFilterV2PutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFilterV2PutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/v2/filters/{id}] filterV2Put", response, response.Code())
	}
}

// NewFilterV2PutOK creates a FilterV2PutOK with default headers values
func NewFilterV2PutOK() *FilterV2PutOK {
	return &FilterV2PutOK{}
}

/*
FilterV2PutOK describes a response with status code 200, with default header values.

Updated filter.
*/
type FilterV2PutOK struct {
	Payload *models.FilterV2
}

// IsSuccess returns true when this filter v2 put o k response has a 2xx status code
func (o *FilterV2PutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this filter v2 put o k response has a 3xx status code
func (o *FilterV2PutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v2 put o k response has a 4xx status code
func (o *FilterV2PutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this filter v2 put o k response has a 5xx status code
func (o *FilterV2PutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this filter v2 put o k response a status code equal to that given
func (o *FilterV2PutOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the filter v2 put o k response
func (o *FilterV2PutOK) Code() int {
	return 200
}

func (o *FilterV2PutOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v2/filters/{id}][%d] filterV2PutOK %s", 200, payload)
}

func (o *FilterV2PutOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v2/filters/{id}][%d] filterV2PutOK %s", 200, payload)
}

func (o *FilterV2PutOK) GetPayload() *models.FilterV2 {
	return o.Payload
}

func (o *FilterV2PutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FilterV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFilterV2PutBadRequest creates a FilterV2PutBadRequest with default headers values
func NewFilterV2PutBadRequest() *FilterV2PutBadRequest {
	return &FilterV2PutBadRequest{}
}

/*
FilterV2PutBadRequest describes a response with status code 400, with default header values.

bad request
*/
type FilterV2PutBadRequest struct {
}

// IsSuccess returns true when this filter v2 put bad request response has a 2xx status code
func (o *FilterV2PutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter v2 put bad request response has a 3xx status code
func (o *FilterV2PutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v2 put bad request response has a 4xx status code
func (o *FilterV2PutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter v2 put bad request response has a 5xx status code
func (o *FilterV2PutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this filter v2 put bad request response a status code equal to that given
func (o *FilterV2PutBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the filter v2 put bad request response
func (o *FilterV2PutBadRequest) Code() int {
	return 400
}

func (o *FilterV2PutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/filters/{id}][%d] filterV2PutBadRequest", 400)
}

func (o *FilterV2PutBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/filters/{id}][%d] filterV2PutBadRequest", 400)
}

func (o *FilterV2PutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterV2PutUnauthorized creates a FilterV2PutUnauthorized with default headers values
func NewFilterV2PutUnauthorized() *FilterV2PutUnauthorized {
	return &FilterV2PutUnauthorized{}
}

/*
FilterV2PutUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type FilterV2PutUnauthorized struct {
}

// IsSuccess returns true when this filter v2 put unauthorized response has a 2xx status code
func (o *FilterV2PutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter v2 put unauthorized response has a 3xx status code
func (o *FilterV2PutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v2 put unauthorized response has a 4xx status code
func (o *FilterV2PutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter v2 put unauthorized response has a 5xx status code
func (o *FilterV2PutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this filter v2 put unauthorized response a status code equal to that given
func (o *FilterV2PutUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the filter v2 put unauthorized response
func (o *FilterV2PutUnauthorized) Code() int {
	return 401
}

func (o *FilterV2PutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/filters/{id}][%d] filterV2PutUnauthorized", 401)
}

func (o *FilterV2PutUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/filters/{id}][%d] filterV2PutUnauthorized", 401)
}

func (o *FilterV2PutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterV2PutForbidden creates a FilterV2PutForbidden with default headers values
func NewFilterV2PutForbidden() *FilterV2PutForbidden {
	return &FilterV2PutForbidden{}
}

/*
FilterV2PutForbidden describes a response with status code 403, with default header values.

forbidden to moved accounts
*/
type FilterV2PutForbidden struct {
}

// IsSuccess returns true when this filter v2 put forbidden response has a 2xx status code
func (o *FilterV2PutForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter v2 put forbidden response has a 3xx status code
func (o *FilterV2PutForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v2 put forbidden response has a 4xx status code
func (o *FilterV2PutForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter v2 put forbidden response has a 5xx status code
func (o *FilterV2PutForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this filter v2 put forbidden response a status code equal to that given
func (o *FilterV2PutForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the filter v2 put forbidden response
func (o *FilterV2PutForbidden) Code() int {
	return 403
}

func (o *FilterV2PutForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/filters/{id}][%d] filterV2PutForbidden", 403)
}

func (o *FilterV2PutForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/filters/{id}][%d] filterV2PutForbidden", 403)
}

func (o *FilterV2PutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterV2PutNotFound creates a FilterV2PutNotFound with default headers values
func NewFilterV2PutNotFound() *FilterV2PutNotFound {
	return &FilterV2PutNotFound{}
}

/*
FilterV2PutNotFound describes a response with status code 404, with default header values.

not found
*/
type FilterV2PutNotFound struct {
}

// IsSuccess returns true when this filter v2 put not found response has a 2xx status code
func (o *FilterV2PutNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter v2 put not found response has a 3xx status code
func (o *FilterV2PutNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v2 put not found response has a 4xx status code
func (o *FilterV2PutNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter v2 put not found response has a 5xx status code
func (o *FilterV2PutNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this filter v2 put not found response a status code equal to that given
func (o *FilterV2PutNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the filter v2 put not found response
func (o *FilterV2PutNotFound) Code() int {
	return 404
}

func (o *FilterV2PutNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/filters/{id}][%d] filterV2PutNotFound", 404)
}

func (o *FilterV2PutNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/filters/{id}][%d] filterV2PutNotFound", 404)
}

func (o *FilterV2PutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterV2PutNotAcceptable creates a FilterV2PutNotAcceptable with default headers values
func NewFilterV2PutNotAcceptable() *FilterV2PutNotAcceptable {
	return &FilterV2PutNotAcceptable{}
}

/*
FilterV2PutNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type FilterV2PutNotAcceptable struct {
}

// IsSuccess returns true when this filter v2 put not acceptable response has a 2xx status code
func (o *FilterV2PutNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter v2 put not acceptable response has a 3xx status code
func (o *FilterV2PutNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v2 put not acceptable response has a 4xx status code
func (o *FilterV2PutNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter v2 put not acceptable response has a 5xx status code
func (o *FilterV2PutNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this filter v2 put not acceptable response a status code equal to that given
func (o *FilterV2PutNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the filter v2 put not acceptable response
func (o *FilterV2PutNotAcceptable) Code() int {
	return 406
}

func (o *FilterV2PutNotAcceptable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/filters/{id}][%d] filterV2PutNotAcceptable", 406)
}

func (o *FilterV2PutNotAcceptable) String() string {
	return fmt.Sprintf("[PUT /api/v2/filters/{id}][%d] filterV2PutNotAcceptable", 406)
}

func (o *FilterV2PutNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterV2PutConflict creates a FilterV2PutConflict with default headers values
func NewFilterV2PutConflict() *FilterV2PutConflict {
	return &FilterV2PutConflict{}
}

/*
FilterV2PutConflict describes a response with status code 409, with default header values.

conflict (duplicate title, keyword, or status)
*/
type FilterV2PutConflict struct {
}

// IsSuccess returns true when this filter v2 put conflict response has a 2xx status code
func (o *FilterV2PutConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter v2 put conflict response has a 3xx status code
func (o *FilterV2PutConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v2 put conflict response has a 4xx status code
func (o *FilterV2PutConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter v2 put conflict response has a 5xx status code
func (o *FilterV2PutConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this filter v2 put conflict response a status code equal to that given
func (o *FilterV2PutConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the filter v2 put conflict response
func (o *FilterV2PutConflict) Code() int {
	return 409
}

func (o *FilterV2PutConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/filters/{id}][%d] filterV2PutConflict", 409)
}

func (o *FilterV2PutConflict) String() string {
	return fmt.Sprintf("[PUT /api/v2/filters/{id}][%d] filterV2PutConflict", 409)
}

func (o *FilterV2PutConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterV2PutUnprocessableEntity creates a FilterV2PutUnprocessableEntity with default headers values
func NewFilterV2PutUnprocessableEntity() *FilterV2PutUnprocessableEntity {
	return &FilterV2PutUnprocessableEntity{}
}

/*
FilterV2PutUnprocessableEntity describes a response with status code 422, with default header values.

unprocessable content
*/
type FilterV2PutUnprocessableEntity struct {
}

// IsSuccess returns true when this filter v2 put unprocessable entity response has a 2xx status code
func (o *FilterV2PutUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter v2 put unprocessable entity response has a 3xx status code
func (o *FilterV2PutUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v2 put unprocessable entity response has a 4xx status code
func (o *FilterV2PutUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter v2 put unprocessable entity response has a 5xx status code
func (o *FilterV2PutUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this filter v2 put unprocessable entity response a status code equal to that given
func (o *FilterV2PutUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the filter v2 put unprocessable entity response
func (o *FilterV2PutUnprocessableEntity) Code() int {
	return 422
}

func (o *FilterV2PutUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v2/filters/{id}][%d] filterV2PutUnprocessableEntity", 422)
}

func (o *FilterV2PutUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /api/v2/filters/{id}][%d] filterV2PutUnprocessableEntity", 422)
}

func (o *FilterV2PutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterV2PutInternalServerError creates a FilterV2PutInternalServerError with default headers values
func NewFilterV2PutInternalServerError() *FilterV2PutInternalServerError {
	return &FilterV2PutInternalServerError{}
}

/*
FilterV2PutInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type FilterV2PutInternalServerError struct {
}

// IsSuccess returns true when this filter v2 put internal server error response has a 2xx status code
func (o *FilterV2PutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter v2 put internal server error response has a 3xx status code
func (o *FilterV2PutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter v2 put internal server error response has a 4xx status code
func (o *FilterV2PutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this filter v2 put internal server error response has a 5xx status code
func (o *FilterV2PutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this filter v2 put internal server error response a status code equal to that given
func (o *FilterV2PutInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the filter v2 put internal server error response
func (o *FilterV2PutInternalServerError) Code() int {
	return 500
}

func (o *FilterV2PutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/filters/{id}][%d] filterV2PutInternalServerError", 500)
}

func (o *FilterV2PutInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/filters/{id}][%d] filterV2PutInternalServerError", 500)
}

func (o *FilterV2PutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
