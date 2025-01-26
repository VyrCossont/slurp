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

// DomainPermissionDraftRemoveReader is a Reader for the DomainPermissionDraftRemove structure.
type DomainPermissionDraftRemoveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainPermissionDraftRemoveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainPermissionDraftRemoveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDomainPermissionDraftRemoveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDomainPermissionDraftRemoveUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDomainPermissionDraftRemoveForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewDomainPermissionDraftRemoveNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDomainPermissionDraftRemoveConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDomainPermissionDraftRemoveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/admin/domain_permission_drafts/{id}/remove] domainPermissionDraftRemove", response, response.Code())
	}
}

// NewDomainPermissionDraftRemoveOK creates a DomainPermissionDraftRemoveOK with default headers values
func NewDomainPermissionDraftRemoveOK() *DomainPermissionDraftRemoveOK {
	return &DomainPermissionDraftRemoveOK{}
}

/*
DomainPermissionDraftRemoveOK describes a response with status code 200, with default header values.

The removed domain permission draft.
*/
type DomainPermissionDraftRemoveOK struct {
	Payload *models.DomainPermission
}

// IsSuccess returns true when this domain permission draft remove o k response has a 2xx status code
func (o *DomainPermissionDraftRemoveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this domain permission draft remove o k response has a 3xx status code
func (o *DomainPermissionDraftRemoveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission draft remove o k response has a 4xx status code
func (o *DomainPermissionDraftRemoveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain permission draft remove o k response has a 5xx status code
func (o *DomainPermissionDraftRemoveOK) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission draft remove o k response a status code equal to that given
func (o *DomainPermissionDraftRemoveOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the domain permission draft remove o k response
func (o *DomainPermissionDraftRemoveOK) Code() int {
	return 200
}

func (o *DomainPermissionDraftRemoveOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts/{id}/remove][%d] domainPermissionDraftRemoveOK %s", 200, payload)
}

func (o *DomainPermissionDraftRemoveOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts/{id}/remove][%d] domainPermissionDraftRemoveOK %s", 200, payload)
}

func (o *DomainPermissionDraftRemoveOK) GetPayload() *models.DomainPermission {
	return o.Payload
}

func (o *DomainPermissionDraftRemoveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDomainPermissionDraftRemoveBadRequest creates a DomainPermissionDraftRemoveBadRequest with default headers values
func NewDomainPermissionDraftRemoveBadRequest() *DomainPermissionDraftRemoveBadRequest {
	return &DomainPermissionDraftRemoveBadRequest{}
}

/*
DomainPermissionDraftRemoveBadRequest describes a response with status code 400, with default header values.

bad request
*/
type DomainPermissionDraftRemoveBadRequest struct {
}

// IsSuccess returns true when this domain permission draft remove bad request response has a 2xx status code
func (o *DomainPermissionDraftRemoveBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission draft remove bad request response has a 3xx status code
func (o *DomainPermissionDraftRemoveBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission draft remove bad request response has a 4xx status code
func (o *DomainPermissionDraftRemoveBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission draft remove bad request response has a 5xx status code
func (o *DomainPermissionDraftRemoveBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission draft remove bad request response a status code equal to that given
func (o *DomainPermissionDraftRemoveBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the domain permission draft remove bad request response
func (o *DomainPermissionDraftRemoveBadRequest) Code() int {
	return 400
}

func (o *DomainPermissionDraftRemoveBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts/{id}/remove][%d] domainPermissionDraftRemoveBadRequest", 400)
}

func (o *DomainPermissionDraftRemoveBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts/{id}/remove][%d] domainPermissionDraftRemoveBadRequest", 400)
}

func (o *DomainPermissionDraftRemoveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionDraftRemoveUnauthorized creates a DomainPermissionDraftRemoveUnauthorized with default headers values
func NewDomainPermissionDraftRemoveUnauthorized() *DomainPermissionDraftRemoveUnauthorized {
	return &DomainPermissionDraftRemoveUnauthorized{}
}

/*
DomainPermissionDraftRemoveUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type DomainPermissionDraftRemoveUnauthorized struct {
}

// IsSuccess returns true when this domain permission draft remove unauthorized response has a 2xx status code
func (o *DomainPermissionDraftRemoveUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission draft remove unauthorized response has a 3xx status code
func (o *DomainPermissionDraftRemoveUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission draft remove unauthorized response has a 4xx status code
func (o *DomainPermissionDraftRemoveUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission draft remove unauthorized response has a 5xx status code
func (o *DomainPermissionDraftRemoveUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission draft remove unauthorized response a status code equal to that given
func (o *DomainPermissionDraftRemoveUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the domain permission draft remove unauthorized response
func (o *DomainPermissionDraftRemoveUnauthorized) Code() int {
	return 401
}

func (o *DomainPermissionDraftRemoveUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts/{id}/remove][%d] domainPermissionDraftRemoveUnauthorized", 401)
}

func (o *DomainPermissionDraftRemoveUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts/{id}/remove][%d] domainPermissionDraftRemoveUnauthorized", 401)
}

func (o *DomainPermissionDraftRemoveUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionDraftRemoveForbidden creates a DomainPermissionDraftRemoveForbidden with default headers values
func NewDomainPermissionDraftRemoveForbidden() *DomainPermissionDraftRemoveForbidden {
	return &DomainPermissionDraftRemoveForbidden{}
}

/*
DomainPermissionDraftRemoveForbidden describes a response with status code 403, with default header values.

forbidden
*/
type DomainPermissionDraftRemoveForbidden struct {
}

// IsSuccess returns true when this domain permission draft remove forbidden response has a 2xx status code
func (o *DomainPermissionDraftRemoveForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission draft remove forbidden response has a 3xx status code
func (o *DomainPermissionDraftRemoveForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission draft remove forbidden response has a 4xx status code
func (o *DomainPermissionDraftRemoveForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission draft remove forbidden response has a 5xx status code
func (o *DomainPermissionDraftRemoveForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission draft remove forbidden response a status code equal to that given
func (o *DomainPermissionDraftRemoveForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the domain permission draft remove forbidden response
func (o *DomainPermissionDraftRemoveForbidden) Code() int {
	return 403
}

func (o *DomainPermissionDraftRemoveForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts/{id}/remove][%d] domainPermissionDraftRemoveForbidden", 403)
}

func (o *DomainPermissionDraftRemoveForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts/{id}/remove][%d] domainPermissionDraftRemoveForbidden", 403)
}

func (o *DomainPermissionDraftRemoveForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionDraftRemoveNotAcceptable creates a DomainPermissionDraftRemoveNotAcceptable with default headers values
func NewDomainPermissionDraftRemoveNotAcceptable() *DomainPermissionDraftRemoveNotAcceptable {
	return &DomainPermissionDraftRemoveNotAcceptable{}
}

/*
DomainPermissionDraftRemoveNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type DomainPermissionDraftRemoveNotAcceptable struct {
}

// IsSuccess returns true when this domain permission draft remove not acceptable response has a 2xx status code
func (o *DomainPermissionDraftRemoveNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission draft remove not acceptable response has a 3xx status code
func (o *DomainPermissionDraftRemoveNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission draft remove not acceptable response has a 4xx status code
func (o *DomainPermissionDraftRemoveNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission draft remove not acceptable response has a 5xx status code
func (o *DomainPermissionDraftRemoveNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission draft remove not acceptable response a status code equal to that given
func (o *DomainPermissionDraftRemoveNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the domain permission draft remove not acceptable response
func (o *DomainPermissionDraftRemoveNotAcceptable) Code() int {
	return 406
}

func (o *DomainPermissionDraftRemoveNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts/{id}/remove][%d] domainPermissionDraftRemoveNotAcceptable", 406)
}

func (o *DomainPermissionDraftRemoveNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts/{id}/remove][%d] domainPermissionDraftRemoveNotAcceptable", 406)
}

func (o *DomainPermissionDraftRemoveNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionDraftRemoveConflict creates a DomainPermissionDraftRemoveConflict with default headers values
func NewDomainPermissionDraftRemoveConflict() *DomainPermissionDraftRemoveConflict {
	return &DomainPermissionDraftRemoveConflict{}
}

/*
DomainPermissionDraftRemoveConflict describes a response with status code 409, with default header values.

conflict
*/
type DomainPermissionDraftRemoveConflict struct {
}

// IsSuccess returns true when this domain permission draft remove conflict response has a 2xx status code
func (o *DomainPermissionDraftRemoveConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission draft remove conflict response has a 3xx status code
func (o *DomainPermissionDraftRemoveConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission draft remove conflict response has a 4xx status code
func (o *DomainPermissionDraftRemoveConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission draft remove conflict response has a 5xx status code
func (o *DomainPermissionDraftRemoveConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission draft remove conflict response a status code equal to that given
func (o *DomainPermissionDraftRemoveConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the domain permission draft remove conflict response
func (o *DomainPermissionDraftRemoveConflict) Code() int {
	return 409
}

func (o *DomainPermissionDraftRemoveConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts/{id}/remove][%d] domainPermissionDraftRemoveConflict", 409)
}

func (o *DomainPermissionDraftRemoveConflict) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts/{id}/remove][%d] domainPermissionDraftRemoveConflict", 409)
}

func (o *DomainPermissionDraftRemoveConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionDraftRemoveInternalServerError creates a DomainPermissionDraftRemoveInternalServerError with default headers values
func NewDomainPermissionDraftRemoveInternalServerError() *DomainPermissionDraftRemoveInternalServerError {
	return &DomainPermissionDraftRemoveInternalServerError{}
}

/*
DomainPermissionDraftRemoveInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type DomainPermissionDraftRemoveInternalServerError struct {
}

// IsSuccess returns true when this domain permission draft remove internal server error response has a 2xx status code
func (o *DomainPermissionDraftRemoveInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission draft remove internal server error response has a 3xx status code
func (o *DomainPermissionDraftRemoveInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission draft remove internal server error response has a 4xx status code
func (o *DomainPermissionDraftRemoveInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain permission draft remove internal server error response has a 5xx status code
func (o *DomainPermissionDraftRemoveInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this domain permission draft remove internal server error response a status code equal to that given
func (o *DomainPermissionDraftRemoveInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the domain permission draft remove internal server error response
func (o *DomainPermissionDraftRemoveInternalServerError) Code() int {
	return 500
}

func (o *DomainPermissionDraftRemoveInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts/{id}/remove][%d] domainPermissionDraftRemoveInternalServerError", 500)
}

func (o *DomainPermissionDraftRemoveInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts/{id}/remove][%d] domainPermissionDraftRemoveInternalServerError", 500)
}

func (o *DomainPermissionDraftRemoveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}