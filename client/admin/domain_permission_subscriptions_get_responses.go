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

// DomainPermissionSubscriptionsGetReader is a Reader for the DomainPermissionSubscriptionsGet structure.
type DomainPermissionSubscriptionsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainPermissionSubscriptionsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainPermissionSubscriptionsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDomainPermissionSubscriptionsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDomainPermissionSubscriptionsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDomainPermissionSubscriptionsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDomainPermissionSubscriptionsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewDomainPermissionSubscriptionsGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDomainPermissionSubscriptionsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/admin/domain_permission_subscriptions] domainPermissionSubscriptionsGet", response, response.Code())
	}
}

// NewDomainPermissionSubscriptionsGetOK creates a DomainPermissionSubscriptionsGetOK with default headers values
func NewDomainPermissionSubscriptionsGetOK() *DomainPermissionSubscriptionsGetOK {
	return &DomainPermissionSubscriptionsGetOK{}
}

/*
DomainPermissionSubscriptionsGetOK describes a response with status code 200, with default header values.

Domain permission subscriptions.
*/
type DomainPermissionSubscriptionsGetOK struct {

	/* Links to the next and previous queries.
	 */
	Link string

	Payload []*models.DomainPermissionSubscription
}

// IsSuccess returns true when this domain permission subscriptions get o k response has a 2xx status code
func (o *DomainPermissionSubscriptionsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this domain permission subscriptions get o k response has a 3xx status code
func (o *DomainPermissionSubscriptionsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscriptions get o k response has a 4xx status code
func (o *DomainPermissionSubscriptionsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain permission subscriptions get o k response has a 5xx status code
func (o *DomainPermissionSubscriptionsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscriptions get o k response a status code equal to that given
func (o *DomainPermissionSubscriptionsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the domain permission subscriptions get o k response
func (o *DomainPermissionSubscriptionsGetOK) Code() int {
	return 200
}

func (o *DomainPermissionSubscriptionsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionsGetOK %s", 200, payload)
}

func (o *DomainPermissionSubscriptionsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionsGetOK %s", 200, payload)
}

func (o *DomainPermissionSubscriptionsGetOK) GetPayload() []*models.DomainPermissionSubscription {
	return o.Payload
}

func (o *DomainPermissionSubscriptionsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDomainPermissionSubscriptionsGetBadRequest creates a DomainPermissionSubscriptionsGetBadRequest with default headers values
func NewDomainPermissionSubscriptionsGetBadRequest() *DomainPermissionSubscriptionsGetBadRequest {
	return &DomainPermissionSubscriptionsGetBadRequest{}
}

/*
DomainPermissionSubscriptionsGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type DomainPermissionSubscriptionsGetBadRequest struct {
}

// IsSuccess returns true when this domain permission subscriptions get bad request response has a 2xx status code
func (o *DomainPermissionSubscriptionsGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscriptions get bad request response has a 3xx status code
func (o *DomainPermissionSubscriptionsGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscriptions get bad request response has a 4xx status code
func (o *DomainPermissionSubscriptionsGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission subscriptions get bad request response has a 5xx status code
func (o *DomainPermissionSubscriptionsGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscriptions get bad request response a status code equal to that given
func (o *DomainPermissionSubscriptionsGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the domain permission subscriptions get bad request response
func (o *DomainPermissionSubscriptionsGetBadRequest) Code() int {
	return 400
}

func (o *DomainPermissionSubscriptionsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionsGetBadRequest", 400)
}

func (o *DomainPermissionSubscriptionsGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionsGetBadRequest", 400)
}

func (o *DomainPermissionSubscriptionsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionSubscriptionsGetUnauthorized creates a DomainPermissionSubscriptionsGetUnauthorized with default headers values
func NewDomainPermissionSubscriptionsGetUnauthorized() *DomainPermissionSubscriptionsGetUnauthorized {
	return &DomainPermissionSubscriptionsGetUnauthorized{}
}

/*
DomainPermissionSubscriptionsGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type DomainPermissionSubscriptionsGetUnauthorized struct {
}

// IsSuccess returns true when this domain permission subscriptions get unauthorized response has a 2xx status code
func (o *DomainPermissionSubscriptionsGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscriptions get unauthorized response has a 3xx status code
func (o *DomainPermissionSubscriptionsGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscriptions get unauthorized response has a 4xx status code
func (o *DomainPermissionSubscriptionsGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission subscriptions get unauthorized response has a 5xx status code
func (o *DomainPermissionSubscriptionsGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscriptions get unauthorized response a status code equal to that given
func (o *DomainPermissionSubscriptionsGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the domain permission subscriptions get unauthorized response
func (o *DomainPermissionSubscriptionsGetUnauthorized) Code() int {
	return 401
}

func (o *DomainPermissionSubscriptionsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionsGetUnauthorized", 401)
}

func (o *DomainPermissionSubscriptionsGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionsGetUnauthorized", 401)
}

func (o *DomainPermissionSubscriptionsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionSubscriptionsGetForbidden creates a DomainPermissionSubscriptionsGetForbidden with default headers values
func NewDomainPermissionSubscriptionsGetForbidden() *DomainPermissionSubscriptionsGetForbidden {
	return &DomainPermissionSubscriptionsGetForbidden{}
}

/*
DomainPermissionSubscriptionsGetForbidden describes a response with status code 403, with default header values.

forbidden
*/
type DomainPermissionSubscriptionsGetForbidden struct {
}

// IsSuccess returns true when this domain permission subscriptions get forbidden response has a 2xx status code
func (o *DomainPermissionSubscriptionsGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscriptions get forbidden response has a 3xx status code
func (o *DomainPermissionSubscriptionsGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscriptions get forbidden response has a 4xx status code
func (o *DomainPermissionSubscriptionsGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission subscriptions get forbidden response has a 5xx status code
func (o *DomainPermissionSubscriptionsGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscriptions get forbidden response a status code equal to that given
func (o *DomainPermissionSubscriptionsGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the domain permission subscriptions get forbidden response
func (o *DomainPermissionSubscriptionsGetForbidden) Code() int {
	return 403
}

func (o *DomainPermissionSubscriptionsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionsGetForbidden", 403)
}

func (o *DomainPermissionSubscriptionsGetForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionsGetForbidden", 403)
}

func (o *DomainPermissionSubscriptionsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionSubscriptionsGetNotFound creates a DomainPermissionSubscriptionsGetNotFound with default headers values
func NewDomainPermissionSubscriptionsGetNotFound() *DomainPermissionSubscriptionsGetNotFound {
	return &DomainPermissionSubscriptionsGetNotFound{}
}

/*
DomainPermissionSubscriptionsGetNotFound describes a response with status code 404, with default header values.

not found
*/
type DomainPermissionSubscriptionsGetNotFound struct {
}

// IsSuccess returns true when this domain permission subscriptions get not found response has a 2xx status code
func (o *DomainPermissionSubscriptionsGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscriptions get not found response has a 3xx status code
func (o *DomainPermissionSubscriptionsGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscriptions get not found response has a 4xx status code
func (o *DomainPermissionSubscriptionsGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission subscriptions get not found response has a 5xx status code
func (o *DomainPermissionSubscriptionsGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscriptions get not found response a status code equal to that given
func (o *DomainPermissionSubscriptionsGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the domain permission subscriptions get not found response
func (o *DomainPermissionSubscriptionsGetNotFound) Code() int {
	return 404
}

func (o *DomainPermissionSubscriptionsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionsGetNotFound", 404)
}

func (o *DomainPermissionSubscriptionsGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionsGetNotFound", 404)
}

func (o *DomainPermissionSubscriptionsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionSubscriptionsGetNotAcceptable creates a DomainPermissionSubscriptionsGetNotAcceptable with default headers values
func NewDomainPermissionSubscriptionsGetNotAcceptable() *DomainPermissionSubscriptionsGetNotAcceptable {
	return &DomainPermissionSubscriptionsGetNotAcceptable{}
}

/*
DomainPermissionSubscriptionsGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type DomainPermissionSubscriptionsGetNotAcceptable struct {
}

// IsSuccess returns true when this domain permission subscriptions get not acceptable response has a 2xx status code
func (o *DomainPermissionSubscriptionsGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscriptions get not acceptable response has a 3xx status code
func (o *DomainPermissionSubscriptionsGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscriptions get not acceptable response has a 4xx status code
func (o *DomainPermissionSubscriptionsGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission subscriptions get not acceptable response has a 5xx status code
func (o *DomainPermissionSubscriptionsGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscriptions get not acceptable response a status code equal to that given
func (o *DomainPermissionSubscriptionsGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the domain permission subscriptions get not acceptable response
func (o *DomainPermissionSubscriptionsGetNotAcceptable) Code() int {
	return 406
}

func (o *DomainPermissionSubscriptionsGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionsGetNotAcceptable", 406)
}

func (o *DomainPermissionSubscriptionsGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionsGetNotAcceptable", 406)
}

func (o *DomainPermissionSubscriptionsGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionSubscriptionsGetInternalServerError creates a DomainPermissionSubscriptionsGetInternalServerError with default headers values
func NewDomainPermissionSubscriptionsGetInternalServerError() *DomainPermissionSubscriptionsGetInternalServerError {
	return &DomainPermissionSubscriptionsGetInternalServerError{}
}

/*
DomainPermissionSubscriptionsGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type DomainPermissionSubscriptionsGetInternalServerError struct {
}

// IsSuccess returns true when this domain permission subscriptions get internal server error response has a 2xx status code
func (o *DomainPermissionSubscriptionsGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscriptions get internal server error response has a 3xx status code
func (o *DomainPermissionSubscriptionsGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscriptions get internal server error response has a 4xx status code
func (o *DomainPermissionSubscriptionsGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain permission subscriptions get internal server error response has a 5xx status code
func (o *DomainPermissionSubscriptionsGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this domain permission subscriptions get internal server error response a status code equal to that given
func (o *DomainPermissionSubscriptionsGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the domain permission subscriptions get internal server error response
func (o *DomainPermissionSubscriptionsGetInternalServerError) Code() int {
	return 500
}

func (o *DomainPermissionSubscriptionsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionsGetInternalServerError", 500)
}

func (o *DomainPermissionSubscriptionsGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionsGetInternalServerError", 500)
}

func (o *DomainPermissionSubscriptionsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
