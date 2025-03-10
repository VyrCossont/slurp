// Code generated by go-swagger; DO NOT EDIT.

package user

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

// UserEmailChangeReader is a Reader for the UserEmailChange structure.
type UserEmailChangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserEmailChangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewUserEmailChangeAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserEmailChangeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserEmailChangeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserEmailChangeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewUserEmailChangeNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUserEmailChangeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserEmailChangeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/user/email_change] userEmailChange", response, response.Code())
	}
}

// NewUserEmailChangeAccepted creates a UserEmailChangeAccepted with default headers values
func NewUserEmailChangeAccepted() *UserEmailChangeAccepted {
	return &UserEmailChangeAccepted{}
}

/*
UserEmailChangeAccepted describes a response with status code 202, with default header values.

Accepted: email change is processing; check your inbox to confirm new address.
*/
type UserEmailChangeAccepted struct {
	Payload *models.User
}

// IsSuccess returns true when this user email change accepted response has a 2xx status code
func (o *UserEmailChangeAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user email change accepted response has a 3xx status code
func (o *UserEmailChangeAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user email change accepted response has a 4xx status code
func (o *UserEmailChangeAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this user email change accepted response has a 5xx status code
func (o *UserEmailChangeAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this user email change accepted response a status code equal to that given
func (o *UserEmailChangeAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the user email change accepted response
func (o *UserEmailChangeAccepted) Code() int {
	return 202
}

func (o *UserEmailChangeAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/user/email_change][%d] userEmailChangeAccepted %s", 202, payload)
}

func (o *UserEmailChangeAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/user/email_change][%d] userEmailChangeAccepted %s", 202, payload)
}

func (o *UserEmailChangeAccepted) GetPayload() *models.User {
	return o.Payload
}

func (o *UserEmailChangeAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserEmailChangeBadRequest creates a UserEmailChangeBadRequest with default headers values
func NewUserEmailChangeBadRequest() *UserEmailChangeBadRequest {
	return &UserEmailChangeBadRequest{}
}

/*
UserEmailChangeBadRequest describes a response with status code 400, with default header values.

bad request
*/
type UserEmailChangeBadRequest struct {
}

// IsSuccess returns true when this user email change bad request response has a 2xx status code
func (o *UserEmailChangeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user email change bad request response has a 3xx status code
func (o *UserEmailChangeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user email change bad request response has a 4xx status code
func (o *UserEmailChangeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user email change bad request response has a 5xx status code
func (o *UserEmailChangeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user email change bad request response a status code equal to that given
func (o *UserEmailChangeBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the user email change bad request response
func (o *UserEmailChangeBadRequest) Code() int {
	return 400
}

func (o *UserEmailChangeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/user/email_change][%d] userEmailChangeBadRequest", 400)
}

func (o *UserEmailChangeBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/user/email_change][%d] userEmailChangeBadRequest", 400)
}

func (o *UserEmailChangeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserEmailChangeUnauthorized creates a UserEmailChangeUnauthorized with default headers values
func NewUserEmailChangeUnauthorized() *UserEmailChangeUnauthorized {
	return &UserEmailChangeUnauthorized{}
}

/*
UserEmailChangeUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type UserEmailChangeUnauthorized struct {
}

// IsSuccess returns true when this user email change unauthorized response has a 2xx status code
func (o *UserEmailChangeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user email change unauthorized response has a 3xx status code
func (o *UserEmailChangeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user email change unauthorized response has a 4xx status code
func (o *UserEmailChangeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this user email change unauthorized response has a 5xx status code
func (o *UserEmailChangeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this user email change unauthorized response a status code equal to that given
func (o *UserEmailChangeUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the user email change unauthorized response
func (o *UserEmailChangeUnauthorized) Code() int {
	return 401
}

func (o *UserEmailChangeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/user/email_change][%d] userEmailChangeUnauthorized", 401)
}

func (o *UserEmailChangeUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/user/email_change][%d] userEmailChangeUnauthorized", 401)
}

func (o *UserEmailChangeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserEmailChangeForbidden creates a UserEmailChangeForbidden with default headers values
func NewUserEmailChangeForbidden() *UserEmailChangeForbidden {
	return &UserEmailChangeForbidden{}
}

/*
UserEmailChangeForbidden describes a response with status code 403, with default header values.

forbidden
*/
type UserEmailChangeForbidden struct {
}

// IsSuccess returns true when this user email change forbidden response has a 2xx status code
func (o *UserEmailChangeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user email change forbidden response has a 3xx status code
func (o *UserEmailChangeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user email change forbidden response has a 4xx status code
func (o *UserEmailChangeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user email change forbidden response has a 5xx status code
func (o *UserEmailChangeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user email change forbidden response a status code equal to that given
func (o *UserEmailChangeForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the user email change forbidden response
func (o *UserEmailChangeForbidden) Code() int {
	return 403
}

func (o *UserEmailChangeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/user/email_change][%d] userEmailChangeForbidden", 403)
}

func (o *UserEmailChangeForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/user/email_change][%d] userEmailChangeForbidden", 403)
}

func (o *UserEmailChangeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserEmailChangeNotAcceptable creates a UserEmailChangeNotAcceptable with default headers values
func NewUserEmailChangeNotAcceptable() *UserEmailChangeNotAcceptable {
	return &UserEmailChangeNotAcceptable{}
}

/*
UserEmailChangeNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type UserEmailChangeNotAcceptable struct {
}

// IsSuccess returns true when this user email change not acceptable response has a 2xx status code
func (o *UserEmailChangeNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user email change not acceptable response has a 3xx status code
func (o *UserEmailChangeNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user email change not acceptable response has a 4xx status code
func (o *UserEmailChangeNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this user email change not acceptable response has a 5xx status code
func (o *UserEmailChangeNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this user email change not acceptable response a status code equal to that given
func (o *UserEmailChangeNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the user email change not acceptable response
func (o *UserEmailChangeNotAcceptable) Code() int {
	return 406
}

func (o *UserEmailChangeNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/user/email_change][%d] userEmailChangeNotAcceptable", 406)
}

func (o *UserEmailChangeNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v1/user/email_change][%d] userEmailChangeNotAcceptable", 406)
}

func (o *UserEmailChangeNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserEmailChangeConflict creates a UserEmailChangeConflict with default headers values
func NewUserEmailChangeConflict() *UserEmailChangeConflict {
	return &UserEmailChangeConflict{}
}

/*
UserEmailChangeConflict describes a response with status code 409, with default header values.

Conflict: desired email address already in use
*/
type UserEmailChangeConflict struct {
}

// IsSuccess returns true when this user email change conflict response has a 2xx status code
func (o *UserEmailChangeConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user email change conflict response has a 3xx status code
func (o *UserEmailChangeConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user email change conflict response has a 4xx status code
func (o *UserEmailChangeConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this user email change conflict response has a 5xx status code
func (o *UserEmailChangeConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this user email change conflict response a status code equal to that given
func (o *UserEmailChangeConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the user email change conflict response
func (o *UserEmailChangeConflict) Code() int {
	return 409
}

func (o *UserEmailChangeConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/user/email_change][%d] userEmailChangeConflict", 409)
}

func (o *UserEmailChangeConflict) String() string {
	return fmt.Sprintf("[POST /api/v1/user/email_change][%d] userEmailChangeConflict", 409)
}

func (o *UserEmailChangeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserEmailChangeInternalServerError creates a UserEmailChangeInternalServerError with default headers values
func NewUserEmailChangeInternalServerError() *UserEmailChangeInternalServerError {
	return &UserEmailChangeInternalServerError{}
}

/*
UserEmailChangeInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type UserEmailChangeInternalServerError struct {
}

// IsSuccess returns true when this user email change internal server error response has a 2xx status code
func (o *UserEmailChangeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user email change internal server error response has a 3xx status code
func (o *UserEmailChangeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user email change internal server error response has a 4xx status code
func (o *UserEmailChangeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user email change internal server error response has a 5xx status code
func (o *UserEmailChangeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user email change internal server error response a status code equal to that given
func (o *UserEmailChangeInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the user email change internal server error response
func (o *UserEmailChangeInternalServerError) Code() int {
	return 500
}

func (o *UserEmailChangeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/user/email_change][%d] userEmailChangeInternalServerError", 500)
}

func (o *UserEmailChangeInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/user/email_change][%d] userEmailChangeInternalServerError", 500)
}

func (o *UserEmailChangeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
