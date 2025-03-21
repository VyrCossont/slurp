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

// ReportsReader is a Reader for the Reports structure.
type ReportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReportsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewReportsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReportsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewReportsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReportsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/reports] reports", response, response.Code())
	}
}

// NewReportsOK creates a ReportsOK with default headers values
func NewReportsOK() *ReportsOK {
	return &ReportsOK{}
}

/*
ReportsOK describes a response with status code 200, with default header values.

Array of reports.
*/
type ReportsOK struct {

	/* Links to the next and previous queries.
	 */
	Link string

	Payload []*models.Report
}

// IsSuccess returns true when this reports o k response has a 2xx status code
func (o *ReportsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this reports o k response has a 3xx status code
func (o *ReportsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reports o k response has a 4xx status code
func (o *ReportsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this reports o k response has a 5xx status code
func (o *ReportsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this reports o k response a status code equal to that given
func (o *ReportsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the reports o k response
func (o *ReportsOK) Code() int {
	return 200
}

func (o *ReportsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/reports][%d] reportsOK %s", 200, payload)
}

func (o *ReportsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/reports][%d] reportsOK %s", 200, payload)
}

func (o *ReportsOK) GetPayload() []*models.Report {
	return o.Payload
}

func (o *ReportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReportsBadRequest creates a ReportsBadRequest with default headers values
func NewReportsBadRequest() *ReportsBadRequest {
	return &ReportsBadRequest{}
}

/*
ReportsBadRequest describes a response with status code 400, with default header values.

bad request
*/
type ReportsBadRequest struct {
}

// IsSuccess returns true when this reports bad request response has a 2xx status code
func (o *ReportsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this reports bad request response has a 3xx status code
func (o *ReportsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reports bad request response has a 4xx status code
func (o *ReportsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this reports bad request response has a 5xx status code
func (o *ReportsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this reports bad request response a status code equal to that given
func (o *ReportsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the reports bad request response
func (o *ReportsBadRequest) Code() int {
	return 400
}

func (o *ReportsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/reports][%d] reportsBadRequest", 400)
}

func (o *ReportsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/reports][%d] reportsBadRequest", 400)
}

func (o *ReportsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReportsUnauthorized creates a ReportsUnauthorized with default headers values
func NewReportsUnauthorized() *ReportsUnauthorized {
	return &ReportsUnauthorized{}
}

/*
ReportsUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type ReportsUnauthorized struct {
}

// IsSuccess returns true when this reports unauthorized response has a 2xx status code
func (o *ReportsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this reports unauthorized response has a 3xx status code
func (o *ReportsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reports unauthorized response has a 4xx status code
func (o *ReportsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this reports unauthorized response has a 5xx status code
func (o *ReportsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this reports unauthorized response a status code equal to that given
func (o *ReportsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the reports unauthorized response
func (o *ReportsUnauthorized) Code() int {
	return 401
}

func (o *ReportsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/reports][%d] reportsUnauthorized", 401)
}

func (o *ReportsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/reports][%d] reportsUnauthorized", 401)
}

func (o *ReportsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReportsNotFound creates a ReportsNotFound with default headers values
func NewReportsNotFound() *ReportsNotFound {
	return &ReportsNotFound{}
}

/*
ReportsNotFound describes a response with status code 404, with default header values.

not found
*/
type ReportsNotFound struct {
}

// IsSuccess returns true when this reports not found response has a 2xx status code
func (o *ReportsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this reports not found response has a 3xx status code
func (o *ReportsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reports not found response has a 4xx status code
func (o *ReportsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this reports not found response has a 5xx status code
func (o *ReportsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this reports not found response a status code equal to that given
func (o *ReportsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the reports not found response
func (o *ReportsNotFound) Code() int {
	return 404
}

func (o *ReportsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/reports][%d] reportsNotFound", 404)
}

func (o *ReportsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/reports][%d] reportsNotFound", 404)
}

func (o *ReportsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReportsNotAcceptable creates a ReportsNotAcceptable with default headers values
func NewReportsNotAcceptable() *ReportsNotAcceptable {
	return &ReportsNotAcceptable{}
}

/*
ReportsNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type ReportsNotAcceptable struct {
}

// IsSuccess returns true when this reports not acceptable response has a 2xx status code
func (o *ReportsNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this reports not acceptable response has a 3xx status code
func (o *ReportsNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reports not acceptable response has a 4xx status code
func (o *ReportsNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this reports not acceptable response has a 5xx status code
func (o *ReportsNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this reports not acceptable response a status code equal to that given
func (o *ReportsNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the reports not acceptable response
func (o *ReportsNotAcceptable) Code() int {
	return 406
}

func (o *ReportsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/reports][%d] reportsNotAcceptable", 406)
}

func (o *ReportsNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/reports][%d] reportsNotAcceptable", 406)
}

func (o *ReportsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReportsInternalServerError creates a ReportsInternalServerError with default headers values
func NewReportsInternalServerError() *ReportsInternalServerError {
	return &ReportsInternalServerError{}
}

/*
ReportsInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type ReportsInternalServerError struct {
}

// IsSuccess returns true when this reports internal server error response has a 2xx status code
func (o *ReportsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this reports internal server error response has a 3xx status code
func (o *ReportsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reports internal server error response has a 4xx status code
func (o *ReportsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this reports internal server error response has a 5xx status code
func (o *ReportsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this reports internal server error response a status code equal to that given
func (o *ReportsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the reports internal server error response
func (o *ReportsInternalServerError) Code() int {
	return 500
}

func (o *ReportsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/reports][%d] reportsInternalServerError", 500)
}

func (o *ReportsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/reports][%d] reportsInternalServerError", 500)
}

func (o *ReportsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
