// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// DomainBlockDeleteReader is a Reader for the DomainBlockDelete structure.
type DomainBlockDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainBlockDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainBlockDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDomainBlockDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDomainBlockDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDomainBlockDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDomainBlockDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewDomainBlockDeleteNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDomainBlockDeleteConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDomainBlockDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /api/v1/admin/domain_blocks/{id}] domainBlockDelete", response, response.Code())
	}
}

// NewDomainBlockDeleteOK creates a DomainBlockDeleteOK with default headers values
func NewDomainBlockDeleteOK() *DomainBlockDeleteOK {
	return &DomainBlockDeleteOK{}
}

/*
DomainBlockDeleteOK describes a response with status code 200, with default header values.

The domain block that was just deleted.
*/
type DomainBlockDeleteOK struct {
	Payload *models.DomainPermission
}

// IsSuccess returns true when this domain block delete o k response has a 2xx status code
func (o *DomainBlockDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this domain block delete o k response has a 3xx status code
func (o *DomainBlockDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain block delete o k response has a 4xx status code
func (o *DomainBlockDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain block delete o k response has a 5xx status code
func (o *DomainBlockDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this domain block delete o k response a status code equal to that given
func (o *DomainBlockDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the domain block delete o k response
func (o *DomainBlockDeleteOK) Code() int {
	return 200
}

func (o *DomainBlockDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_blocks/{id}][%d] domainBlockDeleteOK %s", 200, payload)
}

func (o *DomainBlockDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_blocks/{id}][%d] domainBlockDeleteOK %s", 200, payload)
}

func (o *DomainBlockDeleteOK) GetPayload() *models.DomainPermission {
	return o.Payload
}

func (o *DomainBlockDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDomainBlockDeleteBadRequest creates a DomainBlockDeleteBadRequest with default headers values
func NewDomainBlockDeleteBadRequest() *DomainBlockDeleteBadRequest {
	return &DomainBlockDeleteBadRequest{}
}

/*
DomainBlockDeleteBadRequest describes a response with status code 400, with default header values.

bad request
*/
type DomainBlockDeleteBadRequest struct {
}

// IsSuccess returns true when this domain block delete bad request response has a 2xx status code
func (o *DomainBlockDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain block delete bad request response has a 3xx status code
func (o *DomainBlockDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain block delete bad request response has a 4xx status code
func (o *DomainBlockDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain block delete bad request response has a 5xx status code
func (o *DomainBlockDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this domain block delete bad request response a status code equal to that given
func (o *DomainBlockDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the domain block delete bad request response
func (o *DomainBlockDeleteBadRequest) Code() int {
	return 400
}

func (o *DomainBlockDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_blocks/{id}][%d] domainBlockDeleteBadRequest", 400)
}

func (o *DomainBlockDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_blocks/{id}][%d] domainBlockDeleteBadRequest", 400)
}

func (o *DomainBlockDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainBlockDeleteUnauthorized creates a DomainBlockDeleteUnauthorized with default headers values
func NewDomainBlockDeleteUnauthorized() *DomainBlockDeleteUnauthorized {
	return &DomainBlockDeleteUnauthorized{}
}

/*
DomainBlockDeleteUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type DomainBlockDeleteUnauthorized struct {
}

// IsSuccess returns true when this domain block delete unauthorized response has a 2xx status code
func (o *DomainBlockDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain block delete unauthorized response has a 3xx status code
func (o *DomainBlockDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain block delete unauthorized response has a 4xx status code
func (o *DomainBlockDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain block delete unauthorized response has a 5xx status code
func (o *DomainBlockDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this domain block delete unauthorized response a status code equal to that given
func (o *DomainBlockDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the domain block delete unauthorized response
func (o *DomainBlockDeleteUnauthorized) Code() int {
	return 401
}

func (o *DomainBlockDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_blocks/{id}][%d] domainBlockDeleteUnauthorized", 401)
}

func (o *DomainBlockDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_blocks/{id}][%d] domainBlockDeleteUnauthorized", 401)
}

func (o *DomainBlockDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainBlockDeleteForbidden creates a DomainBlockDeleteForbidden with default headers values
func NewDomainBlockDeleteForbidden() *DomainBlockDeleteForbidden {
	return &DomainBlockDeleteForbidden{}
}

/*
DomainBlockDeleteForbidden describes a response with status code 403, with default header values.

forbidden
*/
type DomainBlockDeleteForbidden struct {
}

// IsSuccess returns true when this domain block delete forbidden response has a 2xx status code
func (o *DomainBlockDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain block delete forbidden response has a 3xx status code
func (o *DomainBlockDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain block delete forbidden response has a 4xx status code
func (o *DomainBlockDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain block delete forbidden response has a 5xx status code
func (o *DomainBlockDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this domain block delete forbidden response a status code equal to that given
func (o *DomainBlockDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the domain block delete forbidden response
func (o *DomainBlockDeleteForbidden) Code() int {
	return 403
}

func (o *DomainBlockDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_blocks/{id}][%d] domainBlockDeleteForbidden", 403)
}

func (o *DomainBlockDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_blocks/{id}][%d] domainBlockDeleteForbidden", 403)
}

func (o *DomainBlockDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainBlockDeleteNotFound creates a DomainBlockDeleteNotFound with default headers values
func NewDomainBlockDeleteNotFound() *DomainBlockDeleteNotFound {
	return &DomainBlockDeleteNotFound{}
}

/*
DomainBlockDeleteNotFound describes a response with status code 404, with default header values.

not found
*/
type DomainBlockDeleteNotFound struct {
}

// IsSuccess returns true when this domain block delete not found response has a 2xx status code
func (o *DomainBlockDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain block delete not found response has a 3xx status code
func (o *DomainBlockDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain block delete not found response has a 4xx status code
func (o *DomainBlockDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain block delete not found response has a 5xx status code
func (o *DomainBlockDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this domain block delete not found response a status code equal to that given
func (o *DomainBlockDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the domain block delete not found response
func (o *DomainBlockDeleteNotFound) Code() int {
	return 404
}

func (o *DomainBlockDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_blocks/{id}][%d] domainBlockDeleteNotFound", 404)
}

func (o *DomainBlockDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_blocks/{id}][%d] domainBlockDeleteNotFound", 404)
}

func (o *DomainBlockDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainBlockDeleteNotAcceptable creates a DomainBlockDeleteNotAcceptable with default headers values
func NewDomainBlockDeleteNotAcceptable() *DomainBlockDeleteNotAcceptable {
	return &DomainBlockDeleteNotAcceptable{}
}

/*
DomainBlockDeleteNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type DomainBlockDeleteNotAcceptable struct {
}

// IsSuccess returns true when this domain block delete not acceptable response has a 2xx status code
func (o *DomainBlockDeleteNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain block delete not acceptable response has a 3xx status code
func (o *DomainBlockDeleteNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain block delete not acceptable response has a 4xx status code
func (o *DomainBlockDeleteNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain block delete not acceptable response has a 5xx status code
func (o *DomainBlockDeleteNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this domain block delete not acceptable response a status code equal to that given
func (o *DomainBlockDeleteNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the domain block delete not acceptable response
func (o *DomainBlockDeleteNotAcceptable) Code() int {
	return 406
}

func (o *DomainBlockDeleteNotAcceptable) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_blocks/{id}][%d] domainBlockDeleteNotAcceptable", 406)
}

func (o *DomainBlockDeleteNotAcceptable) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_blocks/{id}][%d] domainBlockDeleteNotAcceptable", 406)
}

func (o *DomainBlockDeleteNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainBlockDeleteConflict creates a DomainBlockDeleteConflict with default headers values
func NewDomainBlockDeleteConflict() *DomainBlockDeleteConflict {
	return &DomainBlockDeleteConflict{}
}

/*
DomainBlockDeleteConflict describes a response with status code 409, with default header values.

Conflict: There is already an admin action running that conflicts with this action. Check the error message in the response body for more information. This is a temporary error; it should be possible to process this action if you try again in a bit.
*/
type DomainBlockDeleteConflict struct {
}

// IsSuccess returns true when this domain block delete conflict response has a 2xx status code
func (o *DomainBlockDeleteConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain block delete conflict response has a 3xx status code
func (o *DomainBlockDeleteConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain block delete conflict response has a 4xx status code
func (o *DomainBlockDeleteConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain block delete conflict response has a 5xx status code
func (o *DomainBlockDeleteConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this domain block delete conflict response a status code equal to that given
func (o *DomainBlockDeleteConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the domain block delete conflict response
func (o *DomainBlockDeleteConflict) Code() int {
	return 409
}

func (o *DomainBlockDeleteConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_blocks/{id}][%d] domainBlockDeleteConflict", 409)
}

func (o *DomainBlockDeleteConflict) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_blocks/{id}][%d] domainBlockDeleteConflict", 409)
}

func (o *DomainBlockDeleteConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainBlockDeleteInternalServerError creates a DomainBlockDeleteInternalServerError with default headers values
func NewDomainBlockDeleteInternalServerError() *DomainBlockDeleteInternalServerError {
	return &DomainBlockDeleteInternalServerError{}
}

/*
DomainBlockDeleteInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type DomainBlockDeleteInternalServerError struct {
}

// IsSuccess returns true when this domain block delete internal server error response has a 2xx status code
func (o *DomainBlockDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain block delete internal server error response has a 3xx status code
func (o *DomainBlockDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain block delete internal server error response has a 4xx status code
func (o *DomainBlockDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain block delete internal server error response has a 5xx status code
func (o *DomainBlockDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this domain block delete internal server error response a status code equal to that given
func (o *DomainBlockDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the domain block delete internal server error response
func (o *DomainBlockDeleteInternalServerError) Code() int {
	return 500
}

func (o *DomainBlockDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_blocks/{id}][%d] domainBlockDeleteInternalServerError", 500)
}

func (o *DomainBlockDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/domain_blocks/{id}][%d] domainBlockDeleteInternalServerError", 500)
}

func (o *DomainBlockDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
