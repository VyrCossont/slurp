// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VyrCossont/slurp/models"
)

// DomainAllowCreateReader is a Reader for the DomainAllowCreate structure.
type DomainAllowCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainAllowCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainAllowCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDomainAllowCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDomainAllowCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDomainAllowCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDomainAllowCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewDomainAllowCreateNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDomainAllowCreateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDomainAllowCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/admin/domain_allows] domainAllowCreate", response, response.Code())
	}
}

// NewDomainAllowCreateOK creates a DomainAllowCreateOK with default headers values
func NewDomainAllowCreateOK() *DomainAllowCreateOK {
	return &DomainAllowCreateOK{}
}

/*
DomainAllowCreateOK describes a response with status code 200, with default header values.

The newly created domain allow, if `import` != `true`. If a list has been imported, then an `array` of newly created domain allows will be returned instead.
*/
type DomainAllowCreateOK struct {
	Payload *models.DomainPermission
}

// IsSuccess returns true when this domain allow create o k response has a 2xx status code
func (o *DomainAllowCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this domain allow create o k response has a 3xx status code
func (o *DomainAllowCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain allow create o k response has a 4xx status code
func (o *DomainAllowCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain allow create o k response has a 5xx status code
func (o *DomainAllowCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this domain allow create o k response a status code equal to that given
func (o *DomainAllowCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the domain allow create o k response
func (o *DomainAllowCreateOK) Code() int {
	return 200
}

func (o *DomainAllowCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_allows][%d] domainAllowCreateOK  %+v", 200, o.Payload)
}

func (o *DomainAllowCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_allows][%d] domainAllowCreateOK  %+v", 200, o.Payload)
}

func (o *DomainAllowCreateOK) GetPayload() *models.DomainPermission {
	return o.Payload
}

func (o *DomainAllowCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDomainAllowCreateBadRequest creates a DomainAllowCreateBadRequest with default headers values
func NewDomainAllowCreateBadRequest() *DomainAllowCreateBadRequest {
	return &DomainAllowCreateBadRequest{}
}

/*
DomainAllowCreateBadRequest describes a response with status code 400, with default header values.

bad request
*/
type DomainAllowCreateBadRequest struct {
}

// IsSuccess returns true when this domain allow create bad request response has a 2xx status code
func (o *DomainAllowCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain allow create bad request response has a 3xx status code
func (o *DomainAllowCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain allow create bad request response has a 4xx status code
func (o *DomainAllowCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain allow create bad request response has a 5xx status code
func (o *DomainAllowCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this domain allow create bad request response a status code equal to that given
func (o *DomainAllowCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the domain allow create bad request response
func (o *DomainAllowCreateBadRequest) Code() int {
	return 400
}

func (o *DomainAllowCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_allows][%d] domainAllowCreateBadRequest ", 400)
}

func (o *DomainAllowCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_allows][%d] domainAllowCreateBadRequest ", 400)
}

func (o *DomainAllowCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainAllowCreateUnauthorized creates a DomainAllowCreateUnauthorized with default headers values
func NewDomainAllowCreateUnauthorized() *DomainAllowCreateUnauthorized {
	return &DomainAllowCreateUnauthorized{}
}

/*
DomainAllowCreateUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type DomainAllowCreateUnauthorized struct {
}

// IsSuccess returns true when this domain allow create unauthorized response has a 2xx status code
func (o *DomainAllowCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain allow create unauthorized response has a 3xx status code
func (o *DomainAllowCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain allow create unauthorized response has a 4xx status code
func (o *DomainAllowCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain allow create unauthorized response has a 5xx status code
func (o *DomainAllowCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this domain allow create unauthorized response a status code equal to that given
func (o *DomainAllowCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the domain allow create unauthorized response
func (o *DomainAllowCreateUnauthorized) Code() int {
	return 401
}

func (o *DomainAllowCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_allows][%d] domainAllowCreateUnauthorized ", 401)
}

func (o *DomainAllowCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_allows][%d] domainAllowCreateUnauthorized ", 401)
}

func (o *DomainAllowCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainAllowCreateForbidden creates a DomainAllowCreateForbidden with default headers values
func NewDomainAllowCreateForbidden() *DomainAllowCreateForbidden {
	return &DomainAllowCreateForbidden{}
}

/*
DomainAllowCreateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type DomainAllowCreateForbidden struct {
}

// IsSuccess returns true when this domain allow create forbidden response has a 2xx status code
func (o *DomainAllowCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain allow create forbidden response has a 3xx status code
func (o *DomainAllowCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain allow create forbidden response has a 4xx status code
func (o *DomainAllowCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain allow create forbidden response has a 5xx status code
func (o *DomainAllowCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this domain allow create forbidden response a status code equal to that given
func (o *DomainAllowCreateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the domain allow create forbidden response
func (o *DomainAllowCreateForbidden) Code() int {
	return 403
}

func (o *DomainAllowCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_allows][%d] domainAllowCreateForbidden ", 403)
}

func (o *DomainAllowCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_allows][%d] domainAllowCreateForbidden ", 403)
}

func (o *DomainAllowCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainAllowCreateNotFound creates a DomainAllowCreateNotFound with default headers values
func NewDomainAllowCreateNotFound() *DomainAllowCreateNotFound {
	return &DomainAllowCreateNotFound{}
}

/*
DomainAllowCreateNotFound describes a response with status code 404, with default header values.

not found
*/
type DomainAllowCreateNotFound struct {
}

// IsSuccess returns true when this domain allow create not found response has a 2xx status code
func (o *DomainAllowCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain allow create not found response has a 3xx status code
func (o *DomainAllowCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain allow create not found response has a 4xx status code
func (o *DomainAllowCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain allow create not found response has a 5xx status code
func (o *DomainAllowCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this domain allow create not found response a status code equal to that given
func (o *DomainAllowCreateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the domain allow create not found response
func (o *DomainAllowCreateNotFound) Code() int {
	return 404
}

func (o *DomainAllowCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_allows][%d] domainAllowCreateNotFound ", 404)
}

func (o *DomainAllowCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_allows][%d] domainAllowCreateNotFound ", 404)
}

func (o *DomainAllowCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainAllowCreateNotAcceptable creates a DomainAllowCreateNotAcceptable with default headers values
func NewDomainAllowCreateNotAcceptable() *DomainAllowCreateNotAcceptable {
	return &DomainAllowCreateNotAcceptable{}
}

/*
DomainAllowCreateNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type DomainAllowCreateNotAcceptable struct {
}

// IsSuccess returns true when this domain allow create not acceptable response has a 2xx status code
func (o *DomainAllowCreateNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain allow create not acceptable response has a 3xx status code
func (o *DomainAllowCreateNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain allow create not acceptable response has a 4xx status code
func (o *DomainAllowCreateNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain allow create not acceptable response has a 5xx status code
func (o *DomainAllowCreateNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this domain allow create not acceptable response a status code equal to that given
func (o *DomainAllowCreateNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the domain allow create not acceptable response
func (o *DomainAllowCreateNotAcceptable) Code() int {
	return 406
}

func (o *DomainAllowCreateNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_allows][%d] domainAllowCreateNotAcceptable ", 406)
}

func (o *DomainAllowCreateNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_allows][%d] domainAllowCreateNotAcceptable ", 406)
}

func (o *DomainAllowCreateNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainAllowCreateConflict creates a DomainAllowCreateConflict with default headers values
func NewDomainAllowCreateConflict() *DomainAllowCreateConflict {
	return &DomainAllowCreateConflict{}
}

/*
DomainAllowCreateConflict describes a response with status code 409, with default header values.

Conflict: There is already an admin action running that conflicts with this action. Check the error message in the response body for more information. This is a temporary error; it should be possible to process this action if you try again in a bit.
*/
type DomainAllowCreateConflict struct {
}

// IsSuccess returns true when this domain allow create conflict response has a 2xx status code
func (o *DomainAllowCreateConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain allow create conflict response has a 3xx status code
func (o *DomainAllowCreateConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain allow create conflict response has a 4xx status code
func (o *DomainAllowCreateConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain allow create conflict response has a 5xx status code
func (o *DomainAllowCreateConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this domain allow create conflict response a status code equal to that given
func (o *DomainAllowCreateConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the domain allow create conflict response
func (o *DomainAllowCreateConflict) Code() int {
	return 409
}

func (o *DomainAllowCreateConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_allows][%d] domainAllowCreateConflict ", 409)
}

func (o *DomainAllowCreateConflict) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_allows][%d] domainAllowCreateConflict ", 409)
}

func (o *DomainAllowCreateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainAllowCreateInternalServerError creates a DomainAllowCreateInternalServerError with default headers values
func NewDomainAllowCreateInternalServerError() *DomainAllowCreateInternalServerError {
	return &DomainAllowCreateInternalServerError{}
}

/*
DomainAllowCreateInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type DomainAllowCreateInternalServerError struct {
}

// IsSuccess returns true when this domain allow create internal server error response has a 2xx status code
func (o *DomainAllowCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain allow create internal server error response has a 3xx status code
func (o *DomainAllowCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain allow create internal server error response has a 4xx status code
func (o *DomainAllowCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain allow create internal server error response has a 5xx status code
func (o *DomainAllowCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this domain allow create internal server error response a status code equal to that given
func (o *DomainAllowCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the domain allow create internal server error response
func (o *DomainAllowCreateInternalServerError) Code() int {
	return 500
}

func (o *DomainAllowCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_allows][%d] domainAllowCreateInternalServerError ", 500)
}

func (o *DomainAllowCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_allows][%d] domainAllowCreateInternalServerError ", 500)
}

func (o *DomainAllowCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
