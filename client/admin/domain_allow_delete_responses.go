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

// DomainAllowDeleteReader is a Reader for the DomainAllowDelete structure.
type DomainAllowDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainAllowDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainAllowDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDomainAllowDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDomainAllowDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDomainAllowDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDomainAllowDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewDomainAllowDeleteNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDomainAllowDeleteConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDomainAllowDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /api/v1/admin/domain_allows/{id}] domainAllowDelete", response, response.Code())
	}
}

// NewDomainAllowDeleteOK creates a DomainAllowDeleteOK with default headers values
func NewDomainAllowDeleteOK() *DomainAllowDeleteOK {
	return &DomainAllowDeleteOK{}
}

/*
DomainAllowDeleteOK describes a response with status code 200, with default header values.

The domain allow that was just deleted.
*/
type DomainAllowDeleteOK struct {
	Payload *models.DomainPermission
}

// IsSuccess returns true when this domain allow delete o k response has a 2xx status code
func (o *DomainAllowDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this domain allow delete o k response has a 3xx status code
func (o *DomainAllowDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain allow delete o k response has a 4xx status code
func (o *DomainAllowDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain allow delete o k response has a 5xx status code
func (o *DomainAllowDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this domain allow delete o k response a status code equal to that given
func (o *DomainAllowDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the domain allow delete o k response
func (o *DomainAllowDeleteOK) Code() int {
	return 200
}

func (o *DomainAllowDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_allows/{id}][%d] domainAllowDeleteOK  %+v", 200, o.Payload)
}

func (o *DomainAllowDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_allows/{id}][%d] domainAllowDeleteOK  %+v", 200, o.Payload)
}

func (o *DomainAllowDeleteOK) GetPayload() *models.DomainPermission {
	return o.Payload
}

func (o *DomainAllowDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDomainAllowDeleteBadRequest creates a DomainAllowDeleteBadRequest with default headers values
func NewDomainAllowDeleteBadRequest() *DomainAllowDeleteBadRequest {
	return &DomainAllowDeleteBadRequest{}
}

/*
DomainAllowDeleteBadRequest describes a response with status code 400, with default header values.

bad request
*/
type DomainAllowDeleteBadRequest struct {
}

// IsSuccess returns true when this domain allow delete bad request response has a 2xx status code
func (o *DomainAllowDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain allow delete bad request response has a 3xx status code
func (o *DomainAllowDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain allow delete bad request response has a 4xx status code
func (o *DomainAllowDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain allow delete bad request response has a 5xx status code
func (o *DomainAllowDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this domain allow delete bad request response a status code equal to that given
func (o *DomainAllowDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the domain allow delete bad request response
func (o *DomainAllowDeleteBadRequest) Code() int {
	return 400
}

func (o *DomainAllowDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_allows/{id}][%d] domainAllowDeleteBadRequest ", 400)
}

func (o *DomainAllowDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_allows/{id}][%d] domainAllowDeleteBadRequest ", 400)
}

func (o *DomainAllowDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainAllowDeleteUnauthorized creates a DomainAllowDeleteUnauthorized with default headers values
func NewDomainAllowDeleteUnauthorized() *DomainAllowDeleteUnauthorized {
	return &DomainAllowDeleteUnauthorized{}
}

/*
DomainAllowDeleteUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type DomainAllowDeleteUnauthorized struct {
}

// IsSuccess returns true when this domain allow delete unauthorized response has a 2xx status code
func (o *DomainAllowDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain allow delete unauthorized response has a 3xx status code
func (o *DomainAllowDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain allow delete unauthorized response has a 4xx status code
func (o *DomainAllowDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain allow delete unauthorized response has a 5xx status code
func (o *DomainAllowDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this domain allow delete unauthorized response a status code equal to that given
func (o *DomainAllowDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the domain allow delete unauthorized response
func (o *DomainAllowDeleteUnauthorized) Code() int {
	return 401
}

func (o *DomainAllowDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_allows/{id}][%d] domainAllowDeleteUnauthorized ", 401)
}

func (o *DomainAllowDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_allows/{id}][%d] domainAllowDeleteUnauthorized ", 401)
}

func (o *DomainAllowDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainAllowDeleteForbidden creates a DomainAllowDeleteForbidden with default headers values
func NewDomainAllowDeleteForbidden() *DomainAllowDeleteForbidden {
	return &DomainAllowDeleteForbidden{}
}

/*
DomainAllowDeleteForbidden describes a response with status code 403, with default header values.

forbidden
*/
type DomainAllowDeleteForbidden struct {
}

// IsSuccess returns true when this domain allow delete forbidden response has a 2xx status code
func (o *DomainAllowDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain allow delete forbidden response has a 3xx status code
func (o *DomainAllowDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain allow delete forbidden response has a 4xx status code
func (o *DomainAllowDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain allow delete forbidden response has a 5xx status code
func (o *DomainAllowDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this domain allow delete forbidden response a status code equal to that given
func (o *DomainAllowDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the domain allow delete forbidden response
func (o *DomainAllowDeleteForbidden) Code() int {
	return 403
}

func (o *DomainAllowDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_allows/{id}][%d] domainAllowDeleteForbidden ", 403)
}

func (o *DomainAllowDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_allows/{id}][%d] domainAllowDeleteForbidden ", 403)
}

func (o *DomainAllowDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainAllowDeleteNotFound creates a DomainAllowDeleteNotFound with default headers values
func NewDomainAllowDeleteNotFound() *DomainAllowDeleteNotFound {
	return &DomainAllowDeleteNotFound{}
}

/*
DomainAllowDeleteNotFound describes a response with status code 404, with default header values.

not found
*/
type DomainAllowDeleteNotFound struct {
}

// IsSuccess returns true when this domain allow delete not found response has a 2xx status code
func (o *DomainAllowDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain allow delete not found response has a 3xx status code
func (o *DomainAllowDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain allow delete not found response has a 4xx status code
func (o *DomainAllowDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain allow delete not found response has a 5xx status code
func (o *DomainAllowDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this domain allow delete not found response a status code equal to that given
func (o *DomainAllowDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the domain allow delete not found response
func (o *DomainAllowDeleteNotFound) Code() int {
	return 404
}

func (o *DomainAllowDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_allows/{id}][%d] domainAllowDeleteNotFound ", 404)
}

func (o *DomainAllowDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_allows/{id}][%d] domainAllowDeleteNotFound ", 404)
}

func (o *DomainAllowDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainAllowDeleteNotAcceptable creates a DomainAllowDeleteNotAcceptable with default headers values
func NewDomainAllowDeleteNotAcceptable() *DomainAllowDeleteNotAcceptable {
	return &DomainAllowDeleteNotAcceptable{}
}

/*
DomainAllowDeleteNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type DomainAllowDeleteNotAcceptable struct {
}

// IsSuccess returns true when this domain allow delete not acceptable response has a 2xx status code
func (o *DomainAllowDeleteNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain allow delete not acceptable response has a 3xx status code
func (o *DomainAllowDeleteNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain allow delete not acceptable response has a 4xx status code
func (o *DomainAllowDeleteNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain allow delete not acceptable response has a 5xx status code
func (o *DomainAllowDeleteNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this domain allow delete not acceptable response a status code equal to that given
func (o *DomainAllowDeleteNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the domain allow delete not acceptable response
func (o *DomainAllowDeleteNotAcceptable) Code() int {
	return 406
}

func (o *DomainAllowDeleteNotAcceptable) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_allows/{id}][%d] domainAllowDeleteNotAcceptable ", 406)
}

func (o *DomainAllowDeleteNotAcceptable) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_allows/{id}][%d] domainAllowDeleteNotAcceptable ", 406)
}

func (o *DomainAllowDeleteNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainAllowDeleteConflict creates a DomainAllowDeleteConflict with default headers values
func NewDomainAllowDeleteConflict() *DomainAllowDeleteConflict {
	return &DomainAllowDeleteConflict{}
}

/*
DomainAllowDeleteConflict describes a response with status code 409, with default header values.

Conflict: There is already an admin action running that conflicts with this action. Check the error message in the response body for more information. This is a temporary error; it should be possible to process this action if you try again in a bit.
*/
type DomainAllowDeleteConflict struct {
}

// IsSuccess returns true when this domain allow delete conflict response has a 2xx status code
func (o *DomainAllowDeleteConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain allow delete conflict response has a 3xx status code
func (o *DomainAllowDeleteConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain allow delete conflict response has a 4xx status code
func (o *DomainAllowDeleteConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain allow delete conflict response has a 5xx status code
func (o *DomainAllowDeleteConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this domain allow delete conflict response a status code equal to that given
func (o *DomainAllowDeleteConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the domain allow delete conflict response
func (o *DomainAllowDeleteConflict) Code() int {
	return 409
}

func (o *DomainAllowDeleteConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_allows/{id}][%d] domainAllowDeleteConflict ", 409)
}

func (o *DomainAllowDeleteConflict) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_allows/{id}][%d] domainAllowDeleteConflict ", 409)
}

func (o *DomainAllowDeleteConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainAllowDeleteInternalServerError creates a DomainAllowDeleteInternalServerError with default headers values
func NewDomainAllowDeleteInternalServerError() *DomainAllowDeleteInternalServerError {
	return &DomainAllowDeleteInternalServerError{}
}

/*
DomainAllowDeleteInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type DomainAllowDeleteInternalServerError struct {
}

// IsSuccess returns true when this domain allow delete internal server error response has a 2xx status code
func (o *DomainAllowDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain allow delete internal server error response has a 3xx status code
func (o *DomainAllowDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain allow delete internal server error response has a 4xx status code
func (o *DomainAllowDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain allow delete internal server error response has a 5xx status code
func (o *DomainAllowDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this domain allow delete internal server error response a status code equal to that given
func (o *DomainAllowDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the domain allow delete internal server error response
func (o *DomainAllowDeleteInternalServerError) Code() int {
	return 500
}

func (o *DomainAllowDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_allows/{id}][%d] domainAllowDeleteInternalServerError ", 500)
}

func (o *DomainAllowDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_allows/{id}][%d] domainAllowDeleteInternalServerError ", 500)
}

func (o *DomainAllowDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
