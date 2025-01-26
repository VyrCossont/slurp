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

// AdminAccountApproveReader is a Reader for the AdminAccountApprove structure.
type AdminAccountApproveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminAccountApproveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminAccountApproveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminAccountApproveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminAccountApproveUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminAccountApproveForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminAccountApproveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewAdminAccountApproveNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminAccountApproveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/admin/accounts/{id}/approve] adminAccountApprove", response, response.Code())
	}
}

// NewAdminAccountApproveOK creates a AdminAccountApproveOK with default headers values
func NewAdminAccountApproveOK() *AdminAccountApproveOK {
	return &AdminAccountApproveOK{}
}

/*
AdminAccountApproveOK describes a response with status code 200, with default header values.

The now-approved account.
*/
type AdminAccountApproveOK struct {
	Payload *models.AdminAccountInfo
}

// IsSuccess returns true when this admin account approve o k response has a 2xx status code
func (o *AdminAccountApproveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin account approve o k response has a 3xx status code
func (o *AdminAccountApproveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin account approve o k response has a 4xx status code
func (o *AdminAccountApproveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin account approve o k response has a 5xx status code
func (o *AdminAccountApproveOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin account approve o k response a status code equal to that given
func (o *AdminAccountApproveOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the admin account approve o k response
func (o *AdminAccountApproveOK) Code() int {
	return 200
}

func (o *AdminAccountApproveOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/approve][%d] adminAccountApproveOK %s", 200, payload)
}

func (o *AdminAccountApproveOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/approve][%d] adminAccountApproveOK %s", 200, payload)
}

func (o *AdminAccountApproveOK) GetPayload() *models.AdminAccountInfo {
	return o.Payload
}

func (o *AdminAccountApproveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdminAccountInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminAccountApproveBadRequest creates a AdminAccountApproveBadRequest with default headers values
func NewAdminAccountApproveBadRequest() *AdminAccountApproveBadRequest {
	return &AdminAccountApproveBadRequest{}
}

/*
AdminAccountApproveBadRequest describes a response with status code 400, with default header values.

bad request
*/
type AdminAccountApproveBadRequest struct {
}

// IsSuccess returns true when this admin account approve bad request response has a 2xx status code
func (o *AdminAccountApproveBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin account approve bad request response has a 3xx status code
func (o *AdminAccountApproveBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin account approve bad request response has a 4xx status code
func (o *AdminAccountApproveBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin account approve bad request response has a 5xx status code
func (o *AdminAccountApproveBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this admin account approve bad request response a status code equal to that given
func (o *AdminAccountApproveBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the admin account approve bad request response
func (o *AdminAccountApproveBadRequest) Code() int {
	return 400
}

func (o *AdminAccountApproveBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/approve][%d] adminAccountApproveBadRequest", 400)
}

func (o *AdminAccountApproveBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/approve][%d] adminAccountApproveBadRequest", 400)
}

func (o *AdminAccountApproveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminAccountApproveUnauthorized creates a AdminAccountApproveUnauthorized with default headers values
func NewAdminAccountApproveUnauthorized() *AdminAccountApproveUnauthorized {
	return &AdminAccountApproveUnauthorized{}
}

/*
AdminAccountApproveUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type AdminAccountApproveUnauthorized struct {
}

// IsSuccess returns true when this admin account approve unauthorized response has a 2xx status code
func (o *AdminAccountApproveUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin account approve unauthorized response has a 3xx status code
func (o *AdminAccountApproveUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin account approve unauthorized response has a 4xx status code
func (o *AdminAccountApproveUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin account approve unauthorized response has a 5xx status code
func (o *AdminAccountApproveUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this admin account approve unauthorized response a status code equal to that given
func (o *AdminAccountApproveUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the admin account approve unauthorized response
func (o *AdminAccountApproveUnauthorized) Code() int {
	return 401
}

func (o *AdminAccountApproveUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/approve][%d] adminAccountApproveUnauthorized", 401)
}

func (o *AdminAccountApproveUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/approve][%d] adminAccountApproveUnauthorized", 401)
}

func (o *AdminAccountApproveUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminAccountApproveForbidden creates a AdminAccountApproveForbidden with default headers values
func NewAdminAccountApproveForbidden() *AdminAccountApproveForbidden {
	return &AdminAccountApproveForbidden{}
}

/*
AdminAccountApproveForbidden describes a response with status code 403, with default header values.

forbidden
*/
type AdminAccountApproveForbidden struct {
}

// IsSuccess returns true when this admin account approve forbidden response has a 2xx status code
func (o *AdminAccountApproveForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin account approve forbidden response has a 3xx status code
func (o *AdminAccountApproveForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin account approve forbidden response has a 4xx status code
func (o *AdminAccountApproveForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin account approve forbidden response has a 5xx status code
func (o *AdminAccountApproveForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin account approve forbidden response a status code equal to that given
func (o *AdminAccountApproveForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the admin account approve forbidden response
func (o *AdminAccountApproveForbidden) Code() int {
	return 403
}

func (o *AdminAccountApproveForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/approve][%d] adminAccountApproveForbidden", 403)
}

func (o *AdminAccountApproveForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/approve][%d] adminAccountApproveForbidden", 403)
}

func (o *AdminAccountApproveForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminAccountApproveNotFound creates a AdminAccountApproveNotFound with default headers values
func NewAdminAccountApproveNotFound() *AdminAccountApproveNotFound {
	return &AdminAccountApproveNotFound{}
}

/*
AdminAccountApproveNotFound describes a response with status code 404, with default header values.

not found
*/
type AdminAccountApproveNotFound struct {
}

// IsSuccess returns true when this admin account approve not found response has a 2xx status code
func (o *AdminAccountApproveNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin account approve not found response has a 3xx status code
func (o *AdminAccountApproveNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin account approve not found response has a 4xx status code
func (o *AdminAccountApproveNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin account approve not found response has a 5xx status code
func (o *AdminAccountApproveNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this admin account approve not found response a status code equal to that given
func (o *AdminAccountApproveNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the admin account approve not found response
func (o *AdminAccountApproveNotFound) Code() int {
	return 404
}

func (o *AdminAccountApproveNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/approve][%d] adminAccountApproveNotFound", 404)
}

func (o *AdminAccountApproveNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/approve][%d] adminAccountApproveNotFound", 404)
}

func (o *AdminAccountApproveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminAccountApproveNotAcceptable creates a AdminAccountApproveNotAcceptable with default headers values
func NewAdminAccountApproveNotAcceptable() *AdminAccountApproveNotAcceptable {
	return &AdminAccountApproveNotAcceptable{}
}

/*
AdminAccountApproveNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type AdminAccountApproveNotAcceptable struct {
}

// IsSuccess returns true when this admin account approve not acceptable response has a 2xx status code
func (o *AdminAccountApproveNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin account approve not acceptable response has a 3xx status code
func (o *AdminAccountApproveNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin account approve not acceptable response has a 4xx status code
func (o *AdminAccountApproveNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin account approve not acceptable response has a 5xx status code
func (o *AdminAccountApproveNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this admin account approve not acceptable response a status code equal to that given
func (o *AdminAccountApproveNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the admin account approve not acceptable response
func (o *AdminAccountApproveNotAcceptable) Code() int {
	return 406
}

func (o *AdminAccountApproveNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/approve][%d] adminAccountApproveNotAcceptable", 406)
}

func (o *AdminAccountApproveNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/approve][%d] adminAccountApproveNotAcceptable", 406)
}

func (o *AdminAccountApproveNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminAccountApproveInternalServerError creates a AdminAccountApproveInternalServerError with default headers values
func NewAdminAccountApproveInternalServerError() *AdminAccountApproveInternalServerError {
	return &AdminAccountApproveInternalServerError{}
}

/*
AdminAccountApproveInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type AdminAccountApproveInternalServerError struct {
}

// IsSuccess returns true when this admin account approve internal server error response has a 2xx status code
func (o *AdminAccountApproveInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin account approve internal server error response has a 3xx status code
func (o *AdminAccountApproveInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin account approve internal server error response has a 4xx status code
func (o *AdminAccountApproveInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin account approve internal server error response has a 5xx status code
func (o *AdminAccountApproveInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this admin account approve internal server error response a status code equal to that given
func (o *AdminAccountApproveInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the admin account approve internal server error response
func (o *AdminAccountApproveInternalServerError) Code() int {
	return 500
}

func (o *AdminAccountApproveInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/approve][%d] adminAccountApproveInternalServerError", 500)
}

func (o *AdminAccountApproveInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/approve][%d] adminAccountApproveInternalServerError", 500)
}

func (o *AdminAccountApproveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
