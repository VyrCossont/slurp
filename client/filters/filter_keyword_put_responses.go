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

// FilterKeywordPutReader is a Reader for the FilterKeywordPut structure.
type FilterKeywordPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FilterKeywordPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFilterKeywordPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFilterKeywordPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFilterKeywordPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFilterKeywordPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFilterKeywordPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewFilterKeywordPutNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewFilterKeywordPutConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewFilterKeywordPutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFilterKeywordPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/v2/filters/keywords{id}] filterKeywordPut", response, response.Code())
	}
}

// NewFilterKeywordPutOK creates a FilterKeywordPutOK with default headers values
func NewFilterKeywordPutOK() *FilterKeywordPutOK {
	return &FilterKeywordPutOK{}
}

/*
FilterKeywordPutOK describes a response with status code 200, with default header values.

Updated filter keyword.
*/
type FilterKeywordPutOK struct {
	Payload *models.FilterKeyword
}

// IsSuccess returns true when this filter keyword put o k response has a 2xx status code
func (o *FilterKeywordPutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this filter keyword put o k response has a 3xx status code
func (o *FilterKeywordPutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter keyword put o k response has a 4xx status code
func (o *FilterKeywordPutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this filter keyword put o k response has a 5xx status code
func (o *FilterKeywordPutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this filter keyword put o k response a status code equal to that given
func (o *FilterKeywordPutOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the filter keyword put o k response
func (o *FilterKeywordPutOK) Code() int {
	return 200
}

func (o *FilterKeywordPutOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v2/filters/keywords{id}][%d] filterKeywordPutOK %s", 200, payload)
}

func (o *FilterKeywordPutOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v2/filters/keywords{id}][%d] filterKeywordPutOK %s", 200, payload)
}

func (o *FilterKeywordPutOK) GetPayload() *models.FilterKeyword {
	return o.Payload
}

func (o *FilterKeywordPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FilterKeyword)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFilterKeywordPutBadRequest creates a FilterKeywordPutBadRequest with default headers values
func NewFilterKeywordPutBadRequest() *FilterKeywordPutBadRequest {
	return &FilterKeywordPutBadRequest{}
}

/*
FilterKeywordPutBadRequest describes a response with status code 400, with default header values.

bad request
*/
type FilterKeywordPutBadRequest struct {
}

// IsSuccess returns true when this filter keyword put bad request response has a 2xx status code
func (o *FilterKeywordPutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter keyword put bad request response has a 3xx status code
func (o *FilterKeywordPutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter keyword put bad request response has a 4xx status code
func (o *FilterKeywordPutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter keyword put bad request response has a 5xx status code
func (o *FilterKeywordPutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this filter keyword put bad request response a status code equal to that given
func (o *FilterKeywordPutBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the filter keyword put bad request response
func (o *FilterKeywordPutBadRequest) Code() int {
	return 400
}

func (o *FilterKeywordPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/filters/keywords{id}][%d] filterKeywordPutBadRequest", 400)
}

func (o *FilterKeywordPutBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/filters/keywords{id}][%d] filterKeywordPutBadRequest", 400)
}

func (o *FilterKeywordPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterKeywordPutUnauthorized creates a FilterKeywordPutUnauthorized with default headers values
func NewFilterKeywordPutUnauthorized() *FilterKeywordPutUnauthorized {
	return &FilterKeywordPutUnauthorized{}
}

/*
FilterKeywordPutUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type FilterKeywordPutUnauthorized struct {
}

// IsSuccess returns true when this filter keyword put unauthorized response has a 2xx status code
func (o *FilterKeywordPutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter keyword put unauthorized response has a 3xx status code
func (o *FilterKeywordPutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter keyword put unauthorized response has a 4xx status code
func (o *FilterKeywordPutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter keyword put unauthorized response has a 5xx status code
func (o *FilterKeywordPutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this filter keyword put unauthorized response a status code equal to that given
func (o *FilterKeywordPutUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the filter keyword put unauthorized response
func (o *FilterKeywordPutUnauthorized) Code() int {
	return 401
}

func (o *FilterKeywordPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/filters/keywords{id}][%d] filterKeywordPutUnauthorized", 401)
}

func (o *FilterKeywordPutUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/filters/keywords{id}][%d] filterKeywordPutUnauthorized", 401)
}

func (o *FilterKeywordPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterKeywordPutForbidden creates a FilterKeywordPutForbidden with default headers values
func NewFilterKeywordPutForbidden() *FilterKeywordPutForbidden {
	return &FilterKeywordPutForbidden{}
}

/*
FilterKeywordPutForbidden describes a response with status code 403, with default header values.

forbidden to moved accounts
*/
type FilterKeywordPutForbidden struct {
}

// IsSuccess returns true when this filter keyword put forbidden response has a 2xx status code
func (o *FilterKeywordPutForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter keyword put forbidden response has a 3xx status code
func (o *FilterKeywordPutForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter keyword put forbidden response has a 4xx status code
func (o *FilterKeywordPutForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter keyword put forbidden response has a 5xx status code
func (o *FilterKeywordPutForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this filter keyword put forbidden response a status code equal to that given
func (o *FilterKeywordPutForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the filter keyword put forbidden response
func (o *FilterKeywordPutForbidden) Code() int {
	return 403
}

func (o *FilterKeywordPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/filters/keywords{id}][%d] filterKeywordPutForbidden", 403)
}

func (o *FilterKeywordPutForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/filters/keywords{id}][%d] filterKeywordPutForbidden", 403)
}

func (o *FilterKeywordPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterKeywordPutNotFound creates a FilterKeywordPutNotFound with default headers values
func NewFilterKeywordPutNotFound() *FilterKeywordPutNotFound {
	return &FilterKeywordPutNotFound{}
}

/*
FilterKeywordPutNotFound describes a response with status code 404, with default header values.

not found
*/
type FilterKeywordPutNotFound struct {
}

// IsSuccess returns true when this filter keyword put not found response has a 2xx status code
func (o *FilterKeywordPutNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter keyword put not found response has a 3xx status code
func (o *FilterKeywordPutNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter keyword put not found response has a 4xx status code
func (o *FilterKeywordPutNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter keyword put not found response has a 5xx status code
func (o *FilterKeywordPutNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this filter keyword put not found response a status code equal to that given
func (o *FilterKeywordPutNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the filter keyword put not found response
func (o *FilterKeywordPutNotFound) Code() int {
	return 404
}

func (o *FilterKeywordPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/filters/keywords{id}][%d] filterKeywordPutNotFound", 404)
}

func (o *FilterKeywordPutNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/filters/keywords{id}][%d] filterKeywordPutNotFound", 404)
}

func (o *FilterKeywordPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterKeywordPutNotAcceptable creates a FilterKeywordPutNotAcceptable with default headers values
func NewFilterKeywordPutNotAcceptable() *FilterKeywordPutNotAcceptable {
	return &FilterKeywordPutNotAcceptable{}
}

/*
FilterKeywordPutNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type FilterKeywordPutNotAcceptable struct {
}

// IsSuccess returns true when this filter keyword put not acceptable response has a 2xx status code
func (o *FilterKeywordPutNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter keyword put not acceptable response has a 3xx status code
func (o *FilterKeywordPutNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter keyword put not acceptable response has a 4xx status code
func (o *FilterKeywordPutNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter keyword put not acceptable response has a 5xx status code
func (o *FilterKeywordPutNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this filter keyword put not acceptable response a status code equal to that given
func (o *FilterKeywordPutNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the filter keyword put not acceptable response
func (o *FilterKeywordPutNotAcceptable) Code() int {
	return 406
}

func (o *FilterKeywordPutNotAcceptable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/filters/keywords{id}][%d] filterKeywordPutNotAcceptable", 406)
}

func (o *FilterKeywordPutNotAcceptable) String() string {
	return fmt.Sprintf("[PUT /api/v2/filters/keywords{id}][%d] filterKeywordPutNotAcceptable", 406)
}

func (o *FilterKeywordPutNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterKeywordPutConflict creates a FilterKeywordPutConflict with default headers values
func NewFilterKeywordPutConflict() *FilterKeywordPutConflict {
	return &FilterKeywordPutConflict{}
}

/*
FilterKeywordPutConflict describes a response with status code 409, with default header values.

conflict (duplicate keyword)
*/
type FilterKeywordPutConflict struct {
}

// IsSuccess returns true when this filter keyword put conflict response has a 2xx status code
func (o *FilterKeywordPutConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter keyword put conflict response has a 3xx status code
func (o *FilterKeywordPutConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter keyword put conflict response has a 4xx status code
func (o *FilterKeywordPutConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter keyword put conflict response has a 5xx status code
func (o *FilterKeywordPutConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this filter keyword put conflict response a status code equal to that given
func (o *FilterKeywordPutConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the filter keyword put conflict response
func (o *FilterKeywordPutConflict) Code() int {
	return 409
}

func (o *FilterKeywordPutConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/filters/keywords{id}][%d] filterKeywordPutConflict", 409)
}

func (o *FilterKeywordPutConflict) String() string {
	return fmt.Sprintf("[PUT /api/v2/filters/keywords{id}][%d] filterKeywordPutConflict", 409)
}

func (o *FilterKeywordPutConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterKeywordPutUnprocessableEntity creates a FilterKeywordPutUnprocessableEntity with default headers values
func NewFilterKeywordPutUnprocessableEntity() *FilterKeywordPutUnprocessableEntity {
	return &FilterKeywordPutUnprocessableEntity{}
}

/*
FilterKeywordPutUnprocessableEntity describes a response with status code 422, with default header values.

unprocessable content
*/
type FilterKeywordPutUnprocessableEntity struct {
}

// IsSuccess returns true when this filter keyword put unprocessable entity response has a 2xx status code
func (o *FilterKeywordPutUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter keyword put unprocessable entity response has a 3xx status code
func (o *FilterKeywordPutUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter keyword put unprocessable entity response has a 4xx status code
func (o *FilterKeywordPutUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter keyword put unprocessable entity response has a 5xx status code
func (o *FilterKeywordPutUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this filter keyword put unprocessable entity response a status code equal to that given
func (o *FilterKeywordPutUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the filter keyword put unprocessable entity response
func (o *FilterKeywordPutUnprocessableEntity) Code() int {
	return 422
}

func (o *FilterKeywordPutUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v2/filters/keywords{id}][%d] filterKeywordPutUnprocessableEntity", 422)
}

func (o *FilterKeywordPutUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /api/v2/filters/keywords{id}][%d] filterKeywordPutUnprocessableEntity", 422)
}

func (o *FilterKeywordPutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterKeywordPutInternalServerError creates a FilterKeywordPutInternalServerError with default headers values
func NewFilterKeywordPutInternalServerError() *FilterKeywordPutInternalServerError {
	return &FilterKeywordPutInternalServerError{}
}

/*
FilterKeywordPutInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type FilterKeywordPutInternalServerError struct {
}

// IsSuccess returns true when this filter keyword put internal server error response has a 2xx status code
func (o *FilterKeywordPutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter keyword put internal server error response has a 3xx status code
func (o *FilterKeywordPutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter keyword put internal server error response has a 4xx status code
func (o *FilterKeywordPutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this filter keyword put internal server error response has a 5xx status code
func (o *FilterKeywordPutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this filter keyword put internal server error response a status code equal to that given
func (o *FilterKeywordPutInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the filter keyword put internal server error response
func (o *FilterKeywordPutInternalServerError) Code() int {
	return 500
}

func (o *FilterKeywordPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/filters/keywords{id}][%d] filterKeywordPutInternalServerError", 500)
}

func (o *FilterKeywordPutInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/filters/keywords{id}][%d] filterKeywordPutInternalServerError", 500)
}

func (o *FilterKeywordPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
