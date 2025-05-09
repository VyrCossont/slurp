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

// AdminAccountGetReader is a Reader for the AdminAccountGet structure.
type AdminAccountGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminAccountGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminAccountGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminAccountGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminAccountGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminAccountGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminAccountGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewAdminAccountGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminAccountGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/admin/accounts/{id}] adminAccountGet", response, response.Code())
	}
}

// NewAdminAccountGetOK creates a AdminAccountGetOK with default headers values
func NewAdminAccountGetOK() *AdminAccountGetOK {
	return &AdminAccountGetOK{}
}

/*
AdminAccountGetOK describes a response with status code 200, with default header values.

OK
*/
type AdminAccountGetOK struct {
	Payload *models.AdminAccountInfo
}

// IsSuccess returns true when this admin account get o k response has a 2xx status code
func (o *AdminAccountGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin account get o k response has a 3xx status code
func (o *AdminAccountGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin account get o k response has a 4xx status code
func (o *AdminAccountGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin account get o k response has a 5xx status code
func (o *AdminAccountGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin account get o k response a status code equal to that given
func (o *AdminAccountGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the admin account get o k response
func (o *AdminAccountGetOK) Code() int {
	return 200
}

func (o *AdminAccountGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/admin/accounts/{id}][%d] adminAccountGetOK %s", 200, payload)
}

func (o *AdminAccountGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/admin/accounts/{id}][%d] adminAccountGetOK %s", 200, payload)
}

func (o *AdminAccountGetOK) GetPayload() *models.AdminAccountInfo {
	return o.Payload
}

func (o *AdminAccountGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdminAccountInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminAccountGetBadRequest creates a AdminAccountGetBadRequest with default headers values
func NewAdminAccountGetBadRequest() *AdminAccountGetBadRequest {
	return &AdminAccountGetBadRequest{}
}

/*
AdminAccountGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type AdminAccountGetBadRequest struct {
}

// IsSuccess returns true when this admin account get bad request response has a 2xx status code
func (o *AdminAccountGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin account get bad request response has a 3xx status code
func (o *AdminAccountGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin account get bad request response has a 4xx status code
func (o *AdminAccountGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin account get bad request response has a 5xx status code
func (o *AdminAccountGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this admin account get bad request response a status code equal to that given
func (o *AdminAccountGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the admin account get bad request response
func (o *AdminAccountGetBadRequest) Code() int {
	return 400
}

func (o *AdminAccountGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/accounts/{id}][%d] adminAccountGetBadRequest", 400)
}

func (o *AdminAccountGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/accounts/{id}][%d] adminAccountGetBadRequest", 400)
}

func (o *AdminAccountGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminAccountGetUnauthorized creates a AdminAccountGetUnauthorized with default headers values
func NewAdminAccountGetUnauthorized() *AdminAccountGetUnauthorized {
	return &AdminAccountGetUnauthorized{}
}

/*
AdminAccountGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type AdminAccountGetUnauthorized struct {
}

// IsSuccess returns true when this admin account get unauthorized response has a 2xx status code
func (o *AdminAccountGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin account get unauthorized response has a 3xx status code
func (o *AdminAccountGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin account get unauthorized response has a 4xx status code
func (o *AdminAccountGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin account get unauthorized response has a 5xx status code
func (o *AdminAccountGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this admin account get unauthorized response a status code equal to that given
func (o *AdminAccountGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the admin account get unauthorized response
func (o *AdminAccountGetUnauthorized) Code() int {
	return 401
}

func (o *AdminAccountGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/accounts/{id}][%d] adminAccountGetUnauthorized", 401)
}

func (o *AdminAccountGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/accounts/{id}][%d] adminAccountGetUnauthorized", 401)
}

func (o *AdminAccountGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminAccountGetForbidden creates a AdminAccountGetForbidden with default headers values
func NewAdminAccountGetForbidden() *AdminAccountGetForbidden {
	return &AdminAccountGetForbidden{}
}

/*
AdminAccountGetForbidden describes a response with status code 403, with default header values.

forbidden
*/
type AdminAccountGetForbidden struct {
}

// IsSuccess returns true when this admin account get forbidden response has a 2xx status code
func (o *AdminAccountGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin account get forbidden response has a 3xx status code
func (o *AdminAccountGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin account get forbidden response has a 4xx status code
func (o *AdminAccountGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin account get forbidden response has a 5xx status code
func (o *AdminAccountGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin account get forbidden response a status code equal to that given
func (o *AdminAccountGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the admin account get forbidden response
func (o *AdminAccountGetForbidden) Code() int {
	return 403
}

func (o *AdminAccountGetForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/accounts/{id}][%d] adminAccountGetForbidden", 403)
}

func (o *AdminAccountGetForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/accounts/{id}][%d] adminAccountGetForbidden", 403)
}

func (o *AdminAccountGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminAccountGetNotFound creates a AdminAccountGetNotFound with default headers values
func NewAdminAccountGetNotFound() *AdminAccountGetNotFound {
	return &AdminAccountGetNotFound{}
}

/*
AdminAccountGetNotFound describes a response with status code 404, with default header values.

not found
*/
type AdminAccountGetNotFound struct {
}

// IsSuccess returns true when this admin account get not found response has a 2xx status code
func (o *AdminAccountGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin account get not found response has a 3xx status code
func (o *AdminAccountGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin account get not found response has a 4xx status code
func (o *AdminAccountGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin account get not found response has a 5xx status code
func (o *AdminAccountGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this admin account get not found response a status code equal to that given
func (o *AdminAccountGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the admin account get not found response
func (o *AdminAccountGetNotFound) Code() int {
	return 404
}

func (o *AdminAccountGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/accounts/{id}][%d] adminAccountGetNotFound", 404)
}

func (o *AdminAccountGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/accounts/{id}][%d] adminAccountGetNotFound", 404)
}

func (o *AdminAccountGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminAccountGetNotAcceptable creates a AdminAccountGetNotAcceptable with default headers values
func NewAdminAccountGetNotAcceptable() *AdminAccountGetNotAcceptable {
	return &AdminAccountGetNotAcceptable{}
}

/*
AdminAccountGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type AdminAccountGetNotAcceptable struct {
}

// IsSuccess returns true when this admin account get not acceptable response has a 2xx status code
func (o *AdminAccountGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin account get not acceptable response has a 3xx status code
func (o *AdminAccountGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin account get not acceptable response has a 4xx status code
func (o *AdminAccountGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin account get not acceptable response has a 5xx status code
func (o *AdminAccountGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this admin account get not acceptable response a status code equal to that given
func (o *AdminAccountGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the admin account get not acceptable response
func (o *AdminAccountGetNotAcceptable) Code() int {
	return 406
}

func (o *AdminAccountGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/accounts/{id}][%d] adminAccountGetNotAcceptable", 406)
}

func (o *AdminAccountGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/accounts/{id}][%d] adminAccountGetNotAcceptable", 406)
}

func (o *AdminAccountGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminAccountGetInternalServerError creates a AdminAccountGetInternalServerError with default headers values
func NewAdminAccountGetInternalServerError() *AdminAccountGetInternalServerError {
	return &AdminAccountGetInternalServerError{}
}

/*
AdminAccountGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type AdminAccountGetInternalServerError struct {
}

// IsSuccess returns true when this admin account get internal server error response has a 2xx status code
func (o *AdminAccountGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin account get internal server error response has a 3xx status code
func (o *AdminAccountGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin account get internal server error response has a 4xx status code
func (o *AdminAccountGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin account get internal server error response has a 5xx status code
func (o *AdminAccountGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this admin account get internal server error response a status code equal to that given
func (o *AdminAccountGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the admin account get internal server error response
func (o *AdminAccountGetInternalServerError) Code() int {
	return 500
}

func (o *AdminAccountGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/accounts/{id}][%d] adminAccountGetInternalServerError", 500)
}

func (o *AdminAccountGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/accounts/{id}][%d] adminAccountGetInternalServerError", 500)
}

func (o *AdminAccountGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
