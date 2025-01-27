// Code generated by go-swagger; DO NOT EDIT.

package reports

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

// ReportGetReader is a Reader for the ReportGet structure.
type ReportGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReportGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReportGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReportGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewReportGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReportGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewReportGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReportGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/reports/{id}] reportGet", response, response.Code())
	}
}

// NewReportGetOK creates a ReportGetOK with default headers values
func NewReportGetOK() *ReportGetOK {
	return &ReportGetOK{}
}

/*
ReportGetOK describes a response with status code 200, with default header values.

The requested report.
*/
type ReportGetOK struct {
	Payload *models.Report
}

// IsSuccess returns true when this report get o k response has a 2xx status code
func (o *ReportGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this report get o k response has a 3xx status code
func (o *ReportGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report get o k response has a 4xx status code
func (o *ReportGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this report get o k response has a 5xx status code
func (o *ReportGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this report get o k response a status code equal to that given
func (o *ReportGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the report get o k response
func (o *ReportGetOK) Code() int {
	return 200
}

func (o *ReportGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/reports/{id}][%d] reportGetOK %s", 200, payload)
}

func (o *ReportGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/reports/{id}][%d] reportGetOK %s", 200, payload)
}

func (o *ReportGetOK) GetPayload() *models.Report {
	return o.Payload
}

func (o *ReportGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Report)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReportGetBadRequest creates a ReportGetBadRequest with default headers values
func NewReportGetBadRequest() *ReportGetBadRequest {
	return &ReportGetBadRequest{}
}

/*
ReportGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type ReportGetBadRequest struct {
}

// IsSuccess returns true when this report get bad request response has a 2xx status code
func (o *ReportGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this report get bad request response has a 3xx status code
func (o *ReportGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report get bad request response has a 4xx status code
func (o *ReportGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this report get bad request response has a 5xx status code
func (o *ReportGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this report get bad request response a status code equal to that given
func (o *ReportGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the report get bad request response
func (o *ReportGetBadRequest) Code() int {
	return 400
}

func (o *ReportGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/reports/{id}][%d] reportGetBadRequest", 400)
}

func (o *ReportGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/reports/{id}][%d] reportGetBadRequest", 400)
}

func (o *ReportGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReportGetUnauthorized creates a ReportGetUnauthorized with default headers values
func NewReportGetUnauthorized() *ReportGetUnauthorized {
	return &ReportGetUnauthorized{}
}

/*
ReportGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type ReportGetUnauthorized struct {
}

// IsSuccess returns true when this report get unauthorized response has a 2xx status code
func (o *ReportGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this report get unauthorized response has a 3xx status code
func (o *ReportGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report get unauthorized response has a 4xx status code
func (o *ReportGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this report get unauthorized response has a 5xx status code
func (o *ReportGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this report get unauthorized response a status code equal to that given
func (o *ReportGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the report get unauthorized response
func (o *ReportGetUnauthorized) Code() int {
	return 401
}

func (o *ReportGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/reports/{id}][%d] reportGetUnauthorized", 401)
}

func (o *ReportGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/reports/{id}][%d] reportGetUnauthorized", 401)
}

func (o *ReportGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReportGetNotFound creates a ReportGetNotFound with default headers values
func NewReportGetNotFound() *ReportGetNotFound {
	return &ReportGetNotFound{}
}

/*
ReportGetNotFound describes a response with status code 404, with default header values.

not found
*/
type ReportGetNotFound struct {
}

// IsSuccess returns true when this report get not found response has a 2xx status code
func (o *ReportGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this report get not found response has a 3xx status code
func (o *ReportGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report get not found response has a 4xx status code
func (o *ReportGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this report get not found response has a 5xx status code
func (o *ReportGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this report get not found response a status code equal to that given
func (o *ReportGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the report get not found response
func (o *ReportGetNotFound) Code() int {
	return 404
}

func (o *ReportGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/reports/{id}][%d] reportGetNotFound", 404)
}

func (o *ReportGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/reports/{id}][%d] reportGetNotFound", 404)
}

func (o *ReportGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReportGetNotAcceptable creates a ReportGetNotAcceptable with default headers values
func NewReportGetNotAcceptable() *ReportGetNotAcceptable {
	return &ReportGetNotAcceptable{}
}

/*
ReportGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type ReportGetNotAcceptable struct {
}

// IsSuccess returns true when this report get not acceptable response has a 2xx status code
func (o *ReportGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this report get not acceptable response has a 3xx status code
func (o *ReportGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report get not acceptable response has a 4xx status code
func (o *ReportGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this report get not acceptable response has a 5xx status code
func (o *ReportGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this report get not acceptable response a status code equal to that given
func (o *ReportGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the report get not acceptable response
func (o *ReportGetNotAcceptable) Code() int {
	return 406
}

func (o *ReportGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/reports/{id}][%d] reportGetNotAcceptable", 406)
}

func (o *ReportGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/reports/{id}][%d] reportGetNotAcceptable", 406)
}

func (o *ReportGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReportGetInternalServerError creates a ReportGetInternalServerError with default headers values
func NewReportGetInternalServerError() *ReportGetInternalServerError {
	return &ReportGetInternalServerError{}
}

/*
ReportGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type ReportGetInternalServerError struct {
}

// IsSuccess returns true when this report get internal server error response has a 2xx status code
func (o *ReportGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this report get internal server error response has a 3xx status code
func (o *ReportGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report get internal server error response has a 4xx status code
func (o *ReportGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this report get internal server error response has a 5xx status code
func (o *ReportGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this report get internal server error response a status code equal to that given
func (o *ReportGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the report get internal server error response
func (o *ReportGetInternalServerError) Code() int {
	return 500
}

func (o *ReportGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/reports/{id}][%d] reportGetInternalServerError", 500)
}

func (o *ReportGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/reports/{id}][%d] reportGetInternalServerError", 500)
}

func (o *ReportGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
