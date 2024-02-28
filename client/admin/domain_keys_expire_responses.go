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

// DomainKeysExpireReader is a Reader for the DomainKeysExpire structure.
type DomainKeysExpireReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainKeysExpireReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDomainKeysExpireAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDomainKeysExpireBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDomainKeysExpireUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDomainKeysExpireForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDomainKeysExpireNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewDomainKeysExpireNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDomainKeysExpireConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDomainKeysExpireInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/admin/domain_keys_expire] domainKeysExpire", response, response.Code())
	}
}

// NewDomainKeysExpireAccepted creates a DomainKeysExpireAccepted with default headers values
func NewDomainKeysExpireAccepted() *DomainKeysExpireAccepted {
	return &DomainKeysExpireAccepted{}
}

/*
DomainKeysExpireAccepted describes a response with status code 202, with default header values.

Request accepted and will be processed. Check the logs for progress / errors.
*/
type DomainKeysExpireAccepted struct {
	Payload *models.AdminActionResponse
}

// IsSuccess returns true when this domain keys expire accepted response has a 2xx status code
func (o *DomainKeysExpireAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this domain keys expire accepted response has a 3xx status code
func (o *DomainKeysExpireAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain keys expire accepted response has a 4xx status code
func (o *DomainKeysExpireAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain keys expire accepted response has a 5xx status code
func (o *DomainKeysExpireAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this domain keys expire accepted response a status code equal to that given
func (o *DomainKeysExpireAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the domain keys expire accepted response
func (o *DomainKeysExpireAccepted) Code() int {
	return 202
}

func (o *DomainKeysExpireAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/admin/domain_keys_expire][%d] domainKeysExpireAccepted %s", 202, payload)
}

func (o *DomainKeysExpireAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/admin/domain_keys_expire][%d] domainKeysExpireAccepted %s", 202, payload)
}

func (o *DomainKeysExpireAccepted) GetPayload() *models.AdminActionResponse {
	return o.Payload
}

func (o *DomainKeysExpireAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdminActionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDomainKeysExpireBadRequest creates a DomainKeysExpireBadRequest with default headers values
func NewDomainKeysExpireBadRequest() *DomainKeysExpireBadRequest {
	return &DomainKeysExpireBadRequest{}
}

/*
DomainKeysExpireBadRequest describes a response with status code 400, with default header values.

bad request
*/
type DomainKeysExpireBadRequest struct {
}

// IsSuccess returns true when this domain keys expire bad request response has a 2xx status code
func (o *DomainKeysExpireBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain keys expire bad request response has a 3xx status code
func (o *DomainKeysExpireBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain keys expire bad request response has a 4xx status code
func (o *DomainKeysExpireBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain keys expire bad request response has a 5xx status code
func (o *DomainKeysExpireBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this domain keys expire bad request response a status code equal to that given
func (o *DomainKeysExpireBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the domain keys expire bad request response
func (o *DomainKeysExpireBadRequest) Code() int {
	return 400
}

func (o *DomainKeysExpireBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_keys_expire][%d] domainKeysExpireBadRequest", 400)
}

func (o *DomainKeysExpireBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_keys_expire][%d] domainKeysExpireBadRequest", 400)
}

func (o *DomainKeysExpireBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainKeysExpireUnauthorized creates a DomainKeysExpireUnauthorized with default headers values
func NewDomainKeysExpireUnauthorized() *DomainKeysExpireUnauthorized {
	return &DomainKeysExpireUnauthorized{}
}

/*
DomainKeysExpireUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type DomainKeysExpireUnauthorized struct {
}

// IsSuccess returns true when this domain keys expire unauthorized response has a 2xx status code
func (o *DomainKeysExpireUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain keys expire unauthorized response has a 3xx status code
func (o *DomainKeysExpireUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain keys expire unauthorized response has a 4xx status code
func (o *DomainKeysExpireUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain keys expire unauthorized response has a 5xx status code
func (o *DomainKeysExpireUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this domain keys expire unauthorized response a status code equal to that given
func (o *DomainKeysExpireUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the domain keys expire unauthorized response
func (o *DomainKeysExpireUnauthorized) Code() int {
	return 401
}

func (o *DomainKeysExpireUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_keys_expire][%d] domainKeysExpireUnauthorized", 401)
}

func (o *DomainKeysExpireUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_keys_expire][%d] domainKeysExpireUnauthorized", 401)
}

func (o *DomainKeysExpireUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainKeysExpireForbidden creates a DomainKeysExpireForbidden with default headers values
func NewDomainKeysExpireForbidden() *DomainKeysExpireForbidden {
	return &DomainKeysExpireForbidden{}
}

/*
DomainKeysExpireForbidden describes a response with status code 403, with default header values.

forbidden
*/
type DomainKeysExpireForbidden struct {
}

// IsSuccess returns true when this domain keys expire forbidden response has a 2xx status code
func (o *DomainKeysExpireForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain keys expire forbidden response has a 3xx status code
func (o *DomainKeysExpireForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain keys expire forbidden response has a 4xx status code
func (o *DomainKeysExpireForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain keys expire forbidden response has a 5xx status code
func (o *DomainKeysExpireForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this domain keys expire forbidden response a status code equal to that given
func (o *DomainKeysExpireForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the domain keys expire forbidden response
func (o *DomainKeysExpireForbidden) Code() int {
	return 403
}

func (o *DomainKeysExpireForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_keys_expire][%d] domainKeysExpireForbidden", 403)
}

func (o *DomainKeysExpireForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_keys_expire][%d] domainKeysExpireForbidden", 403)
}

func (o *DomainKeysExpireForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainKeysExpireNotFound creates a DomainKeysExpireNotFound with default headers values
func NewDomainKeysExpireNotFound() *DomainKeysExpireNotFound {
	return &DomainKeysExpireNotFound{}
}

/*
DomainKeysExpireNotFound describes a response with status code 404, with default header values.

not found
*/
type DomainKeysExpireNotFound struct {
}

// IsSuccess returns true when this domain keys expire not found response has a 2xx status code
func (o *DomainKeysExpireNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain keys expire not found response has a 3xx status code
func (o *DomainKeysExpireNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain keys expire not found response has a 4xx status code
func (o *DomainKeysExpireNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain keys expire not found response has a 5xx status code
func (o *DomainKeysExpireNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this domain keys expire not found response a status code equal to that given
func (o *DomainKeysExpireNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the domain keys expire not found response
func (o *DomainKeysExpireNotFound) Code() int {
	return 404
}

func (o *DomainKeysExpireNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_keys_expire][%d] domainKeysExpireNotFound", 404)
}

func (o *DomainKeysExpireNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_keys_expire][%d] domainKeysExpireNotFound", 404)
}

func (o *DomainKeysExpireNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainKeysExpireNotAcceptable creates a DomainKeysExpireNotAcceptable with default headers values
func NewDomainKeysExpireNotAcceptable() *DomainKeysExpireNotAcceptable {
	return &DomainKeysExpireNotAcceptable{}
}

/*
DomainKeysExpireNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type DomainKeysExpireNotAcceptable struct {
}

// IsSuccess returns true when this domain keys expire not acceptable response has a 2xx status code
func (o *DomainKeysExpireNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain keys expire not acceptable response has a 3xx status code
func (o *DomainKeysExpireNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain keys expire not acceptable response has a 4xx status code
func (o *DomainKeysExpireNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain keys expire not acceptable response has a 5xx status code
func (o *DomainKeysExpireNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this domain keys expire not acceptable response a status code equal to that given
func (o *DomainKeysExpireNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the domain keys expire not acceptable response
func (o *DomainKeysExpireNotAcceptable) Code() int {
	return 406
}

func (o *DomainKeysExpireNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_keys_expire][%d] domainKeysExpireNotAcceptable", 406)
}

func (o *DomainKeysExpireNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_keys_expire][%d] domainKeysExpireNotAcceptable", 406)
}

func (o *DomainKeysExpireNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainKeysExpireConflict creates a DomainKeysExpireConflict with default headers values
func NewDomainKeysExpireConflict() *DomainKeysExpireConflict {
	return &DomainKeysExpireConflict{}
}

/*
DomainKeysExpireConflict describes a response with status code 409, with default header values.

Conflict: There is already an admin action running that conflicts with this action. Check the error message in the response body for more information. This is a temporary error; it should be possible to process this action if you try again in a bit.
*/
type DomainKeysExpireConflict struct {
}

// IsSuccess returns true when this domain keys expire conflict response has a 2xx status code
func (o *DomainKeysExpireConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain keys expire conflict response has a 3xx status code
func (o *DomainKeysExpireConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain keys expire conflict response has a 4xx status code
func (o *DomainKeysExpireConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain keys expire conflict response has a 5xx status code
func (o *DomainKeysExpireConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this domain keys expire conflict response a status code equal to that given
func (o *DomainKeysExpireConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the domain keys expire conflict response
func (o *DomainKeysExpireConflict) Code() int {
	return 409
}

func (o *DomainKeysExpireConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_keys_expire][%d] domainKeysExpireConflict", 409)
}

func (o *DomainKeysExpireConflict) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_keys_expire][%d] domainKeysExpireConflict", 409)
}

func (o *DomainKeysExpireConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainKeysExpireInternalServerError creates a DomainKeysExpireInternalServerError with default headers values
func NewDomainKeysExpireInternalServerError() *DomainKeysExpireInternalServerError {
	return &DomainKeysExpireInternalServerError{}
}

/*
DomainKeysExpireInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type DomainKeysExpireInternalServerError struct {
}

// IsSuccess returns true when this domain keys expire internal server error response has a 2xx status code
func (o *DomainKeysExpireInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain keys expire internal server error response has a 3xx status code
func (o *DomainKeysExpireInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain keys expire internal server error response has a 4xx status code
func (o *DomainKeysExpireInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain keys expire internal server error response has a 5xx status code
func (o *DomainKeysExpireInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this domain keys expire internal server error response a status code equal to that given
func (o *DomainKeysExpireInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the domain keys expire internal server error response
func (o *DomainKeysExpireInternalServerError) Code() int {
	return 500
}

func (o *DomainKeysExpireInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_keys_expire][%d] domainKeysExpireInternalServerError", 500)
}

func (o *DomainKeysExpireInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_keys_expire][%d] domainKeysExpireInternalServerError", 500)
}

func (o *DomainKeysExpireInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
