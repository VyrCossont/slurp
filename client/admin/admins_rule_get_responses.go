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

// AdminsRuleGetReader is a Reader for the AdminsRuleGet structure.
type AdminsRuleGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminsRuleGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminsRuleGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminsRuleGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminsRuleGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminsRuleGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewAdminsRuleGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminsRuleGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/admin/rules] adminsRuleGet", response, response.Code())
	}
}

// NewAdminsRuleGetOK creates a AdminsRuleGetOK with default headers values
func NewAdminsRuleGetOK() *AdminsRuleGetOK {
	return &AdminsRuleGetOK{}
}

/*
AdminsRuleGetOK describes a response with status code 200, with default header values.

An array with all the rules for the local instance.
*/
type AdminsRuleGetOK struct {
	Payload []*models.InstanceRule
}

// IsSuccess returns true when this admins rule get o k response has a 2xx status code
func (o *AdminsRuleGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admins rule get o k response has a 3xx status code
func (o *AdminsRuleGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admins rule get o k response has a 4xx status code
func (o *AdminsRuleGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admins rule get o k response has a 5xx status code
func (o *AdminsRuleGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admins rule get o k response a status code equal to that given
func (o *AdminsRuleGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the admins rule get o k response
func (o *AdminsRuleGetOK) Code() int {
	return 200
}

func (o *AdminsRuleGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/admin/rules][%d] adminsRuleGetOK %s", 200, payload)
}

func (o *AdminsRuleGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/admin/rules][%d] adminsRuleGetOK %s", 200, payload)
}

func (o *AdminsRuleGetOK) GetPayload() []*models.InstanceRule {
	return o.Payload
}

func (o *AdminsRuleGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminsRuleGetBadRequest creates a AdminsRuleGetBadRequest with default headers values
func NewAdminsRuleGetBadRequest() *AdminsRuleGetBadRequest {
	return &AdminsRuleGetBadRequest{}
}

/*
AdminsRuleGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type AdminsRuleGetBadRequest struct {
}

// IsSuccess returns true when this admins rule get bad request response has a 2xx status code
func (o *AdminsRuleGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admins rule get bad request response has a 3xx status code
func (o *AdminsRuleGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admins rule get bad request response has a 4xx status code
func (o *AdminsRuleGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this admins rule get bad request response has a 5xx status code
func (o *AdminsRuleGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this admins rule get bad request response a status code equal to that given
func (o *AdminsRuleGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the admins rule get bad request response
func (o *AdminsRuleGetBadRequest) Code() int {
	return 400
}

func (o *AdminsRuleGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/rules][%d] adminsRuleGetBadRequest", 400)
}

func (o *AdminsRuleGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/rules][%d] adminsRuleGetBadRequest", 400)
}

func (o *AdminsRuleGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminsRuleGetUnauthorized creates a AdminsRuleGetUnauthorized with default headers values
func NewAdminsRuleGetUnauthorized() *AdminsRuleGetUnauthorized {
	return &AdminsRuleGetUnauthorized{}
}

/*
AdminsRuleGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type AdminsRuleGetUnauthorized struct {
}

// IsSuccess returns true when this admins rule get unauthorized response has a 2xx status code
func (o *AdminsRuleGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admins rule get unauthorized response has a 3xx status code
func (o *AdminsRuleGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admins rule get unauthorized response has a 4xx status code
func (o *AdminsRuleGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this admins rule get unauthorized response has a 5xx status code
func (o *AdminsRuleGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this admins rule get unauthorized response a status code equal to that given
func (o *AdminsRuleGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the admins rule get unauthorized response
func (o *AdminsRuleGetUnauthorized) Code() int {
	return 401
}

func (o *AdminsRuleGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/rules][%d] adminsRuleGetUnauthorized", 401)
}

func (o *AdminsRuleGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/rules][%d] adminsRuleGetUnauthorized", 401)
}

func (o *AdminsRuleGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminsRuleGetNotFound creates a AdminsRuleGetNotFound with default headers values
func NewAdminsRuleGetNotFound() *AdminsRuleGetNotFound {
	return &AdminsRuleGetNotFound{}
}

/*
AdminsRuleGetNotFound describes a response with status code 404, with default header values.

not found
*/
type AdminsRuleGetNotFound struct {
}

// IsSuccess returns true when this admins rule get not found response has a 2xx status code
func (o *AdminsRuleGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admins rule get not found response has a 3xx status code
func (o *AdminsRuleGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admins rule get not found response has a 4xx status code
func (o *AdminsRuleGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this admins rule get not found response has a 5xx status code
func (o *AdminsRuleGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this admins rule get not found response a status code equal to that given
func (o *AdminsRuleGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the admins rule get not found response
func (o *AdminsRuleGetNotFound) Code() int {
	return 404
}

func (o *AdminsRuleGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/rules][%d] adminsRuleGetNotFound", 404)
}

func (o *AdminsRuleGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/rules][%d] adminsRuleGetNotFound", 404)
}

func (o *AdminsRuleGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminsRuleGetNotAcceptable creates a AdminsRuleGetNotAcceptable with default headers values
func NewAdminsRuleGetNotAcceptable() *AdminsRuleGetNotAcceptable {
	return &AdminsRuleGetNotAcceptable{}
}

/*
AdminsRuleGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type AdminsRuleGetNotAcceptable struct {
}

// IsSuccess returns true when this admins rule get not acceptable response has a 2xx status code
func (o *AdminsRuleGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admins rule get not acceptable response has a 3xx status code
func (o *AdminsRuleGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admins rule get not acceptable response has a 4xx status code
func (o *AdminsRuleGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this admins rule get not acceptable response has a 5xx status code
func (o *AdminsRuleGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this admins rule get not acceptable response a status code equal to that given
func (o *AdminsRuleGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the admins rule get not acceptable response
func (o *AdminsRuleGetNotAcceptable) Code() int {
	return 406
}

func (o *AdminsRuleGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/rules][%d] adminsRuleGetNotAcceptable", 406)
}

func (o *AdminsRuleGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/rules][%d] adminsRuleGetNotAcceptable", 406)
}

func (o *AdminsRuleGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminsRuleGetInternalServerError creates a AdminsRuleGetInternalServerError with default headers values
func NewAdminsRuleGetInternalServerError() *AdminsRuleGetInternalServerError {
	return &AdminsRuleGetInternalServerError{}
}

/*
AdminsRuleGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type AdminsRuleGetInternalServerError struct {
}

// IsSuccess returns true when this admins rule get internal server error response has a 2xx status code
func (o *AdminsRuleGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admins rule get internal server error response has a 3xx status code
func (o *AdminsRuleGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admins rule get internal server error response has a 4xx status code
func (o *AdminsRuleGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this admins rule get internal server error response has a 5xx status code
func (o *AdminsRuleGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this admins rule get internal server error response a status code equal to that given
func (o *AdminsRuleGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the admins rule get internal server error response
func (o *AdminsRuleGetInternalServerError) Code() int {
	return 500
}

func (o *AdminsRuleGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/rules][%d] adminsRuleGetInternalServerError", 500)
}

func (o *AdminsRuleGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/rules][%d] adminsRuleGetInternalServerError", 500)
}

func (o *AdminsRuleGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
