// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

// AccountAvatarDeleteReader is a Reader for the AccountAvatarDelete structure.
type AccountAvatarDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountAvatarDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountAvatarDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccountAvatarDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAccountAvatarDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAccountAvatarDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewAccountAvatarDeleteNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAccountAvatarDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /api/v1/profile/avatar] accountAvatarDelete", response, response.Code())
	}
}

// NewAccountAvatarDeleteOK creates a AccountAvatarDeleteOK with default headers values
func NewAccountAvatarDeleteOK() *AccountAvatarDeleteOK {
	return &AccountAvatarDeleteOK{}
}

/*
AccountAvatarDeleteOK describes a response with status code 200, with default header values.

The updated account, including profile source information.
*/
type AccountAvatarDeleteOK struct {
	Payload *models.Account
}

// IsSuccess returns true when this account avatar delete o k response has a 2xx status code
func (o *AccountAvatarDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account avatar delete o k response has a 3xx status code
func (o *AccountAvatarDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account avatar delete o k response has a 4xx status code
func (o *AccountAvatarDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account avatar delete o k response has a 5xx status code
func (o *AccountAvatarDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account avatar delete o k response a status code equal to that given
func (o *AccountAvatarDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the account avatar delete o k response
func (o *AccountAvatarDeleteOK) Code() int {
	return 200
}

func (o *AccountAvatarDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1/profile/avatar][%d] accountAvatarDeleteOK %s", 200, payload)
}

func (o *AccountAvatarDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1/profile/avatar][%d] accountAvatarDeleteOK %s", 200, payload)
}

func (o *AccountAvatarDeleteOK) GetPayload() *models.Account {
	return o.Payload
}

func (o *AccountAvatarDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Account)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountAvatarDeleteBadRequest creates a AccountAvatarDeleteBadRequest with default headers values
func NewAccountAvatarDeleteBadRequest() *AccountAvatarDeleteBadRequest {
	return &AccountAvatarDeleteBadRequest{}
}

/*
AccountAvatarDeleteBadRequest describes a response with status code 400, with default header values.

bad request
*/
type AccountAvatarDeleteBadRequest struct {
}

// IsSuccess returns true when this account avatar delete bad request response has a 2xx status code
func (o *AccountAvatarDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account avatar delete bad request response has a 3xx status code
func (o *AccountAvatarDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account avatar delete bad request response has a 4xx status code
func (o *AccountAvatarDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this account avatar delete bad request response has a 5xx status code
func (o *AccountAvatarDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this account avatar delete bad request response a status code equal to that given
func (o *AccountAvatarDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the account avatar delete bad request response
func (o *AccountAvatarDeleteBadRequest) Code() int {
	return 400
}

func (o *AccountAvatarDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/profile/avatar][%d] accountAvatarDeleteBadRequest", 400)
}

func (o *AccountAvatarDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v1/profile/avatar][%d] accountAvatarDeleteBadRequest", 400)
}

func (o *AccountAvatarDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountAvatarDeleteUnauthorized creates a AccountAvatarDeleteUnauthorized with default headers values
func NewAccountAvatarDeleteUnauthorized() *AccountAvatarDeleteUnauthorized {
	return &AccountAvatarDeleteUnauthorized{}
}

/*
AccountAvatarDeleteUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type AccountAvatarDeleteUnauthorized struct {
}

// IsSuccess returns true when this account avatar delete unauthorized response has a 2xx status code
func (o *AccountAvatarDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account avatar delete unauthorized response has a 3xx status code
func (o *AccountAvatarDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account avatar delete unauthorized response has a 4xx status code
func (o *AccountAvatarDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this account avatar delete unauthorized response has a 5xx status code
func (o *AccountAvatarDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this account avatar delete unauthorized response a status code equal to that given
func (o *AccountAvatarDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the account avatar delete unauthorized response
func (o *AccountAvatarDeleteUnauthorized) Code() int {
	return 401
}

func (o *AccountAvatarDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/profile/avatar][%d] accountAvatarDeleteUnauthorized", 401)
}

func (o *AccountAvatarDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v1/profile/avatar][%d] accountAvatarDeleteUnauthorized", 401)
}

func (o *AccountAvatarDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountAvatarDeleteForbidden creates a AccountAvatarDeleteForbidden with default headers values
func NewAccountAvatarDeleteForbidden() *AccountAvatarDeleteForbidden {
	return &AccountAvatarDeleteForbidden{}
}

/*
AccountAvatarDeleteForbidden describes a response with status code 403, with default header values.

forbidden
*/
type AccountAvatarDeleteForbidden struct {
}

// IsSuccess returns true when this account avatar delete forbidden response has a 2xx status code
func (o *AccountAvatarDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account avatar delete forbidden response has a 3xx status code
func (o *AccountAvatarDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account avatar delete forbidden response has a 4xx status code
func (o *AccountAvatarDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this account avatar delete forbidden response has a 5xx status code
func (o *AccountAvatarDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this account avatar delete forbidden response a status code equal to that given
func (o *AccountAvatarDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the account avatar delete forbidden response
func (o *AccountAvatarDeleteForbidden) Code() int {
	return 403
}

func (o *AccountAvatarDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/profile/avatar][%d] accountAvatarDeleteForbidden", 403)
}

func (o *AccountAvatarDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/profile/avatar][%d] accountAvatarDeleteForbidden", 403)
}

func (o *AccountAvatarDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountAvatarDeleteNotAcceptable creates a AccountAvatarDeleteNotAcceptable with default headers values
func NewAccountAvatarDeleteNotAcceptable() *AccountAvatarDeleteNotAcceptable {
	return &AccountAvatarDeleteNotAcceptable{}
}

/*
AccountAvatarDeleteNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type AccountAvatarDeleteNotAcceptable struct {
}

// IsSuccess returns true when this account avatar delete not acceptable response has a 2xx status code
func (o *AccountAvatarDeleteNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account avatar delete not acceptable response has a 3xx status code
func (o *AccountAvatarDeleteNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account avatar delete not acceptable response has a 4xx status code
func (o *AccountAvatarDeleteNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this account avatar delete not acceptable response has a 5xx status code
func (o *AccountAvatarDeleteNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this account avatar delete not acceptable response a status code equal to that given
func (o *AccountAvatarDeleteNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the account avatar delete not acceptable response
func (o *AccountAvatarDeleteNotAcceptable) Code() int {
	return 406
}

func (o *AccountAvatarDeleteNotAcceptable) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/profile/avatar][%d] accountAvatarDeleteNotAcceptable", 406)
}

func (o *AccountAvatarDeleteNotAcceptable) String() string {
	return fmt.Sprintf("[DELETE /api/v1/profile/avatar][%d] accountAvatarDeleteNotAcceptable", 406)
}

func (o *AccountAvatarDeleteNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountAvatarDeleteInternalServerError creates a AccountAvatarDeleteInternalServerError with default headers values
func NewAccountAvatarDeleteInternalServerError() *AccountAvatarDeleteInternalServerError {
	return &AccountAvatarDeleteInternalServerError{}
}

/*
AccountAvatarDeleteInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type AccountAvatarDeleteInternalServerError struct {
}

// IsSuccess returns true when this account avatar delete internal server error response has a 2xx status code
func (o *AccountAvatarDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account avatar delete internal server error response has a 3xx status code
func (o *AccountAvatarDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account avatar delete internal server error response has a 4xx status code
func (o *AccountAvatarDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this account avatar delete internal server error response has a 5xx status code
func (o *AccountAvatarDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this account avatar delete internal server error response a status code equal to that given
func (o *AccountAvatarDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the account avatar delete internal server error response
func (o *AccountAvatarDeleteInternalServerError) Code() int {
	return 500
}

func (o *AccountAvatarDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/profile/avatar][%d] accountAvatarDeleteInternalServerError", 500)
}

func (o *AccountAvatarDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v1/profile/avatar][%d] accountAvatarDeleteInternalServerError", 500)
}

func (o *AccountAvatarDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
