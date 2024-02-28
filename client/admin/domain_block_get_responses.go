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

// DomainBlockGetReader is a Reader for the DomainBlockGet structure.
type DomainBlockGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainBlockGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainBlockGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDomainBlockGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDomainBlockGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDomainBlockGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDomainBlockGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewDomainBlockGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDomainBlockGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/admin/domain_blocks/{id}] domainBlockGet", response, response.Code())
	}
}

// NewDomainBlockGetOK creates a DomainBlockGetOK with default headers values
func NewDomainBlockGetOK() *DomainBlockGetOK {
	return &DomainBlockGetOK{}
}

/*
DomainBlockGetOK describes a response with status code 200, with default header values.

The requested domain block.
*/
type DomainBlockGetOK struct {
	Payload *models.DomainPermission
}

// IsSuccess returns true when this domain block get o k response has a 2xx status code
func (o *DomainBlockGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this domain block get o k response has a 3xx status code
func (o *DomainBlockGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain block get o k response has a 4xx status code
func (o *DomainBlockGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain block get o k response has a 5xx status code
func (o *DomainBlockGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this domain block get o k response a status code equal to that given
func (o *DomainBlockGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the domain block get o k response
func (o *DomainBlockGetOK) Code() int {
	return 200
}

func (o *DomainBlockGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks/{id}][%d] domainBlockGetOK %s", 200, payload)
}

func (o *DomainBlockGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks/{id}][%d] domainBlockGetOK %s", 200, payload)
}

func (o *DomainBlockGetOK) GetPayload() *models.DomainPermission {
	return o.Payload
}

func (o *DomainBlockGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDomainBlockGetBadRequest creates a DomainBlockGetBadRequest with default headers values
func NewDomainBlockGetBadRequest() *DomainBlockGetBadRequest {
	return &DomainBlockGetBadRequest{}
}

/*
DomainBlockGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type DomainBlockGetBadRequest struct {
}

// IsSuccess returns true when this domain block get bad request response has a 2xx status code
func (o *DomainBlockGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain block get bad request response has a 3xx status code
func (o *DomainBlockGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain block get bad request response has a 4xx status code
func (o *DomainBlockGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain block get bad request response has a 5xx status code
func (o *DomainBlockGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this domain block get bad request response a status code equal to that given
func (o *DomainBlockGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the domain block get bad request response
func (o *DomainBlockGetBadRequest) Code() int {
	return 400
}

func (o *DomainBlockGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks/{id}][%d] domainBlockGetBadRequest", 400)
}

func (o *DomainBlockGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks/{id}][%d] domainBlockGetBadRequest", 400)
}

func (o *DomainBlockGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainBlockGetUnauthorized creates a DomainBlockGetUnauthorized with default headers values
func NewDomainBlockGetUnauthorized() *DomainBlockGetUnauthorized {
	return &DomainBlockGetUnauthorized{}
}

/*
DomainBlockGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type DomainBlockGetUnauthorized struct {
}

// IsSuccess returns true when this domain block get unauthorized response has a 2xx status code
func (o *DomainBlockGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain block get unauthorized response has a 3xx status code
func (o *DomainBlockGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain block get unauthorized response has a 4xx status code
func (o *DomainBlockGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain block get unauthorized response has a 5xx status code
func (o *DomainBlockGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this domain block get unauthorized response a status code equal to that given
func (o *DomainBlockGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the domain block get unauthorized response
func (o *DomainBlockGetUnauthorized) Code() int {
	return 401
}

func (o *DomainBlockGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks/{id}][%d] domainBlockGetUnauthorized", 401)
}

func (o *DomainBlockGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks/{id}][%d] domainBlockGetUnauthorized", 401)
}

func (o *DomainBlockGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainBlockGetForbidden creates a DomainBlockGetForbidden with default headers values
func NewDomainBlockGetForbidden() *DomainBlockGetForbidden {
	return &DomainBlockGetForbidden{}
}

/*
DomainBlockGetForbidden describes a response with status code 403, with default header values.

forbidden
*/
type DomainBlockGetForbidden struct {
}

// IsSuccess returns true when this domain block get forbidden response has a 2xx status code
func (o *DomainBlockGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain block get forbidden response has a 3xx status code
func (o *DomainBlockGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain block get forbidden response has a 4xx status code
func (o *DomainBlockGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain block get forbidden response has a 5xx status code
func (o *DomainBlockGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this domain block get forbidden response a status code equal to that given
func (o *DomainBlockGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the domain block get forbidden response
func (o *DomainBlockGetForbidden) Code() int {
	return 403
}

func (o *DomainBlockGetForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks/{id}][%d] domainBlockGetForbidden", 403)
}

func (o *DomainBlockGetForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks/{id}][%d] domainBlockGetForbidden", 403)
}

func (o *DomainBlockGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainBlockGetNotFound creates a DomainBlockGetNotFound with default headers values
func NewDomainBlockGetNotFound() *DomainBlockGetNotFound {
	return &DomainBlockGetNotFound{}
}

/*
DomainBlockGetNotFound describes a response with status code 404, with default header values.

not found
*/
type DomainBlockGetNotFound struct {
}

// IsSuccess returns true when this domain block get not found response has a 2xx status code
func (o *DomainBlockGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain block get not found response has a 3xx status code
func (o *DomainBlockGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain block get not found response has a 4xx status code
func (o *DomainBlockGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain block get not found response has a 5xx status code
func (o *DomainBlockGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this domain block get not found response a status code equal to that given
func (o *DomainBlockGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the domain block get not found response
func (o *DomainBlockGetNotFound) Code() int {
	return 404
}

func (o *DomainBlockGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks/{id}][%d] domainBlockGetNotFound", 404)
}

func (o *DomainBlockGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks/{id}][%d] domainBlockGetNotFound", 404)
}

func (o *DomainBlockGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainBlockGetNotAcceptable creates a DomainBlockGetNotAcceptable with default headers values
func NewDomainBlockGetNotAcceptable() *DomainBlockGetNotAcceptable {
	return &DomainBlockGetNotAcceptable{}
}

/*
DomainBlockGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type DomainBlockGetNotAcceptable struct {
}

// IsSuccess returns true when this domain block get not acceptable response has a 2xx status code
func (o *DomainBlockGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain block get not acceptable response has a 3xx status code
func (o *DomainBlockGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain block get not acceptable response has a 4xx status code
func (o *DomainBlockGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain block get not acceptable response has a 5xx status code
func (o *DomainBlockGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this domain block get not acceptable response a status code equal to that given
func (o *DomainBlockGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the domain block get not acceptable response
func (o *DomainBlockGetNotAcceptable) Code() int {
	return 406
}

func (o *DomainBlockGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks/{id}][%d] domainBlockGetNotAcceptable", 406)
}

func (o *DomainBlockGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks/{id}][%d] domainBlockGetNotAcceptable", 406)
}

func (o *DomainBlockGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainBlockGetInternalServerError creates a DomainBlockGetInternalServerError with default headers values
func NewDomainBlockGetInternalServerError() *DomainBlockGetInternalServerError {
	return &DomainBlockGetInternalServerError{}
}

/*
DomainBlockGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type DomainBlockGetInternalServerError struct {
}

// IsSuccess returns true when this domain block get internal server error response has a 2xx status code
func (o *DomainBlockGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain block get internal server error response has a 3xx status code
func (o *DomainBlockGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain block get internal server error response has a 4xx status code
func (o *DomainBlockGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain block get internal server error response has a 5xx status code
func (o *DomainBlockGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this domain block get internal server error response a status code equal to that given
func (o *DomainBlockGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the domain block get internal server error response
func (o *DomainBlockGetInternalServerError) Code() int {
	return 500
}

func (o *DomainBlockGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks/{id}][%d] domainBlockGetInternalServerError", 500)
}

func (o *DomainBlockGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks/{id}][%d] domainBlockGetInternalServerError", 500)
}

func (o *DomainBlockGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
