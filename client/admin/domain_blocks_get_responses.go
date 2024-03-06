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

// DomainBlocksGetReader is a Reader for the DomainBlocksGet structure.
type DomainBlocksGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainBlocksGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainBlocksGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDomainBlocksGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDomainBlocksGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDomainBlocksGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDomainBlocksGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewDomainBlocksGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDomainBlocksGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/admin/domain_blocks] domainBlocksGet", response, response.Code())
	}
}

// NewDomainBlocksGetOK creates a DomainBlocksGetOK with default headers values
func NewDomainBlocksGetOK() *DomainBlocksGetOK {
	return &DomainBlocksGetOK{}
}

/*
DomainBlocksGetOK describes a response with status code 200, with default header values.

All domain blocks currently in place.
*/
type DomainBlocksGetOK struct {
	Payload []*models.DomainPermission
}

// IsSuccess returns true when this domain blocks get o k response has a 2xx status code
func (o *DomainBlocksGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this domain blocks get o k response has a 3xx status code
func (o *DomainBlocksGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain blocks get o k response has a 4xx status code
func (o *DomainBlocksGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain blocks get o k response has a 5xx status code
func (o *DomainBlocksGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this domain blocks get o k response a status code equal to that given
func (o *DomainBlocksGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the domain blocks get o k response
func (o *DomainBlocksGetOK) Code() int {
	return 200
}

func (o *DomainBlocksGetOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks][%d] domainBlocksGetOK  %+v", 200, o.Payload)
}

func (o *DomainBlocksGetOK) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks][%d] domainBlocksGetOK  %+v", 200, o.Payload)
}

func (o *DomainBlocksGetOK) GetPayload() []*models.DomainPermission {
	return o.Payload
}

func (o *DomainBlocksGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDomainBlocksGetBadRequest creates a DomainBlocksGetBadRequest with default headers values
func NewDomainBlocksGetBadRequest() *DomainBlocksGetBadRequest {
	return &DomainBlocksGetBadRequest{}
}

/*
DomainBlocksGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type DomainBlocksGetBadRequest struct {
}

// IsSuccess returns true when this domain blocks get bad request response has a 2xx status code
func (o *DomainBlocksGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain blocks get bad request response has a 3xx status code
func (o *DomainBlocksGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain blocks get bad request response has a 4xx status code
func (o *DomainBlocksGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain blocks get bad request response has a 5xx status code
func (o *DomainBlocksGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this domain blocks get bad request response a status code equal to that given
func (o *DomainBlocksGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the domain blocks get bad request response
func (o *DomainBlocksGetBadRequest) Code() int {
	return 400
}

func (o *DomainBlocksGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks][%d] domainBlocksGetBadRequest ", 400)
}

func (o *DomainBlocksGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks][%d] domainBlocksGetBadRequest ", 400)
}

func (o *DomainBlocksGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainBlocksGetUnauthorized creates a DomainBlocksGetUnauthorized with default headers values
func NewDomainBlocksGetUnauthorized() *DomainBlocksGetUnauthorized {
	return &DomainBlocksGetUnauthorized{}
}

/*
DomainBlocksGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type DomainBlocksGetUnauthorized struct {
}

// IsSuccess returns true when this domain blocks get unauthorized response has a 2xx status code
func (o *DomainBlocksGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain blocks get unauthorized response has a 3xx status code
func (o *DomainBlocksGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain blocks get unauthorized response has a 4xx status code
func (o *DomainBlocksGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain blocks get unauthorized response has a 5xx status code
func (o *DomainBlocksGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this domain blocks get unauthorized response a status code equal to that given
func (o *DomainBlocksGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the domain blocks get unauthorized response
func (o *DomainBlocksGetUnauthorized) Code() int {
	return 401
}

func (o *DomainBlocksGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks][%d] domainBlocksGetUnauthorized ", 401)
}

func (o *DomainBlocksGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks][%d] domainBlocksGetUnauthorized ", 401)
}

func (o *DomainBlocksGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainBlocksGetForbidden creates a DomainBlocksGetForbidden with default headers values
func NewDomainBlocksGetForbidden() *DomainBlocksGetForbidden {
	return &DomainBlocksGetForbidden{}
}

/*
DomainBlocksGetForbidden describes a response with status code 403, with default header values.

forbidden
*/
type DomainBlocksGetForbidden struct {
}

// IsSuccess returns true when this domain blocks get forbidden response has a 2xx status code
func (o *DomainBlocksGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain blocks get forbidden response has a 3xx status code
func (o *DomainBlocksGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain blocks get forbidden response has a 4xx status code
func (o *DomainBlocksGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain blocks get forbidden response has a 5xx status code
func (o *DomainBlocksGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this domain blocks get forbidden response a status code equal to that given
func (o *DomainBlocksGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the domain blocks get forbidden response
func (o *DomainBlocksGetForbidden) Code() int {
	return 403
}

func (o *DomainBlocksGetForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks][%d] domainBlocksGetForbidden ", 403)
}

func (o *DomainBlocksGetForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks][%d] domainBlocksGetForbidden ", 403)
}

func (o *DomainBlocksGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainBlocksGetNotFound creates a DomainBlocksGetNotFound with default headers values
func NewDomainBlocksGetNotFound() *DomainBlocksGetNotFound {
	return &DomainBlocksGetNotFound{}
}

/*
DomainBlocksGetNotFound describes a response with status code 404, with default header values.

not found
*/
type DomainBlocksGetNotFound struct {
}

// IsSuccess returns true when this domain blocks get not found response has a 2xx status code
func (o *DomainBlocksGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain blocks get not found response has a 3xx status code
func (o *DomainBlocksGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain blocks get not found response has a 4xx status code
func (o *DomainBlocksGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain blocks get not found response has a 5xx status code
func (o *DomainBlocksGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this domain blocks get not found response a status code equal to that given
func (o *DomainBlocksGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the domain blocks get not found response
func (o *DomainBlocksGetNotFound) Code() int {
	return 404
}

func (o *DomainBlocksGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks][%d] domainBlocksGetNotFound ", 404)
}

func (o *DomainBlocksGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks][%d] domainBlocksGetNotFound ", 404)
}

func (o *DomainBlocksGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainBlocksGetNotAcceptable creates a DomainBlocksGetNotAcceptable with default headers values
func NewDomainBlocksGetNotAcceptable() *DomainBlocksGetNotAcceptable {
	return &DomainBlocksGetNotAcceptable{}
}

/*
DomainBlocksGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type DomainBlocksGetNotAcceptable struct {
}

// IsSuccess returns true when this domain blocks get not acceptable response has a 2xx status code
func (o *DomainBlocksGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain blocks get not acceptable response has a 3xx status code
func (o *DomainBlocksGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain blocks get not acceptable response has a 4xx status code
func (o *DomainBlocksGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain blocks get not acceptable response has a 5xx status code
func (o *DomainBlocksGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this domain blocks get not acceptable response a status code equal to that given
func (o *DomainBlocksGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the domain blocks get not acceptable response
func (o *DomainBlocksGetNotAcceptable) Code() int {
	return 406
}

func (o *DomainBlocksGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks][%d] domainBlocksGetNotAcceptable ", 406)
}

func (o *DomainBlocksGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks][%d] domainBlocksGetNotAcceptable ", 406)
}

func (o *DomainBlocksGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainBlocksGetInternalServerError creates a DomainBlocksGetInternalServerError with default headers values
func NewDomainBlocksGetInternalServerError() *DomainBlocksGetInternalServerError {
	return &DomainBlocksGetInternalServerError{}
}

/*
DomainBlocksGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type DomainBlocksGetInternalServerError struct {
}

// IsSuccess returns true when this domain blocks get internal server error response has a 2xx status code
func (o *DomainBlocksGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain blocks get internal server error response has a 3xx status code
func (o *DomainBlocksGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain blocks get internal server error response has a 4xx status code
func (o *DomainBlocksGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain blocks get internal server error response has a 5xx status code
func (o *DomainBlocksGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this domain blocks get internal server error response a status code equal to that given
func (o *DomainBlocksGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the domain blocks get internal server error response
func (o *DomainBlocksGetInternalServerError) Code() int {
	return 500
}

func (o *DomainBlocksGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks][%d] domainBlocksGetInternalServerError ", 500)
}

func (o *DomainBlocksGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_blocks][%d] domainBlocksGetInternalServerError ", 500)
}

func (o *DomainBlocksGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
