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

// EmojiDeleteReader is a Reader for the EmojiDelete structure.
type EmojiDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmojiDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmojiDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEmojiDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEmojiDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEmojiDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEmojiDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewEmojiDeleteNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEmojiDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /api/v1/admin/custom_emojis/{id}] emojiDelete", response, response.Code())
	}
}

// NewEmojiDeleteOK creates a EmojiDeleteOK with default headers values
func NewEmojiDeleteOK() *EmojiDeleteOK {
	return &EmojiDeleteOK{}
}

/*
EmojiDeleteOK describes a response with status code 200, with default header values.

The deleted emoji will be returned to the caller in case further processing is necessary.
*/
type EmojiDeleteOK struct {
	Payload *models.AdminEmoji
}

// IsSuccess returns true when this emoji delete o k response has a 2xx status code
func (o *EmojiDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this emoji delete o k response has a 3xx status code
func (o *EmojiDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emoji delete o k response has a 4xx status code
func (o *EmojiDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this emoji delete o k response has a 5xx status code
func (o *EmojiDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this emoji delete o k response a status code equal to that given
func (o *EmojiDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the emoji delete o k response
func (o *EmojiDeleteOK) Code() int {
	return 200
}

func (o *EmojiDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/custom_emojis/{id}][%d] emojiDeleteOK  %+v", 200, o.Payload)
}

func (o *EmojiDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/custom_emojis/{id}][%d] emojiDeleteOK  %+v", 200, o.Payload)
}

func (o *EmojiDeleteOK) GetPayload() *models.AdminEmoji {
	return o.Payload
}

func (o *EmojiDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdminEmoji)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmojiDeleteBadRequest creates a EmojiDeleteBadRequest with default headers values
func NewEmojiDeleteBadRequest() *EmojiDeleteBadRequest {
	return &EmojiDeleteBadRequest{}
}

/*
EmojiDeleteBadRequest describes a response with status code 400, with default header values.

bad request
*/
type EmojiDeleteBadRequest struct {
}

// IsSuccess returns true when this emoji delete bad request response has a 2xx status code
func (o *EmojiDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this emoji delete bad request response has a 3xx status code
func (o *EmojiDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emoji delete bad request response has a 4xx status code
func (o *EmojiDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this emoji delete bad request response has a 5xx status code
func (o *EmojiDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this emoji delete bad request response a status code equal to that given
func (o *EmojiDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the emoji delete bad request response
func (o *EmojiDeleteBadRequest) Code() int {
	return 400
}

func (o *EmojiDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/custom_emojis/{id}][%d] emojiDeleteBadRequest ", 400)
}

func (o *EmojiDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/custom_emojis/{id}][%d] emojiDeleteBadRequest ", 400)
}

func (o *EmojiDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmojiDeleteUnauthorized creates a EmojiDeleteUnauthorized with default headers values
func NewEmojiDeleteUnauthorized() *EmojiDeleteUnauthorized {
	return &EmojiDeleteUnauthorized{}
}

/*
EmojiDeleteUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type EmojiDeleteUnauthorized struct {
}

// IsSuccess returns true when this emoji delete unauthorized response has a 2xx status code
func (o *EmojiDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this emoji delete unauthorized response has a 3xx status code
func (o *EmojiDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emoji delete unauthorized response has a 4xx status code
func (o *EmojiDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this emoji delete unauthorized response has a 5xx status code
func (o *EmojiDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this emoji delete unauthorized response a status code equal to that given
func (o *EmojiDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the emoji delete unauthorized response
func (o *EmojiDeleteUnauthorized) Code() int {
	return 401
}

func (o *EmojiDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/custom_emojis/{id}][%d] emojiDeleteUnauthorized ", 401)
}

func (o *EmojiDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/custom_emojis/{id}][%d] emojiDeleteUnauthorized ", 401)
}

func (o *EmojiDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmojiDeleteForbidden creates a EmojiDeleteForbidden with default headers values
func NewEmojiDeleteForbidden() *EmojiDeleteForbidden {
	return &EmojiDeleteForbidden{}
}

/*
EmojiDeleteForbidden describes a response with status code 403, with default header values.

forbidden
*/
type EmojiDeleteForbidden struct {
}

// IsSuccess returns true when this emoji delete forbidden response has a 2xx status code
func (o *EmojiDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this emoji delete forbidden response has a 3xx status code
func (o *EmojiDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emoji delete forbidden response has a 4xx status code
func (o *EmojiDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this emoji delete forbidden response has a 5xx status code
func (o *EmojiDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this emoji delete forbidden response a status code equal to that given
func (o *EmojiDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the emoji delete forbidden response
func (o *EmojiDeleteForbidden) Code() int {
	return 403
}

func (o *EmojiDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/custom_emojis/{id}][%d] emojiDeleteForbidden ", 403)
}

func (o *EmojiDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/custom_emojis/{id}][%d] emojiDeleteForbidden ", 403)
}

func (o *EmojiDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmojiDeleteNotFound creates a EmojiDeleteNotFound with default headers values
func NewEmojiDeleteNotFound() *EmojiDeleteNotFound {
	return &EmojiDeleteNotFound{}
}

/*
EmojiDeleteNotFound describes a response with status code 404, with default header values.

not found
*/
type EmojiDeleteNotFound struct {
}

// IsSuccess returns true when this emoji delete not found response has a 2xx status code
func (o *EmojiDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this emoji delete not found response has a 3xx status code
func (o *EmojiDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emoji delete not found response has a 4xx status code
func (o *EmojiDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this emoji delete not found response has a 5xx status code
func (o *EmojiDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this emoji delete not found response a status code equal to that given
func (o *EmojiDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the emoji delete not found response
func (o *EmojiDeleteNotFound) Code() int {
	return 404
}

func (o *EmojiDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/custom_emojis/{id}][%d] emojiDeleteNotFound ", 404)
}

func (o *EmojiDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/custom_emojis/{id}][%d] emojiDeleteNotFound ", 404)
}

func (o *EmojiDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmojiDeleteNotAcceptable creates a EmojiDeleteNotAcceptable with default headers values
func NewEmojiDeleteNotAcceptable() *EmojiDeleteNotAcceptable {
	return &EmojiDeleteNotAcceptable{}
}

/*
EmojiDeleteNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type EmojiDeleteNotAcceptable struct {
}

// IsSuccess returns true when this emoji delete not acceptable response has a 2xx status code
func (o *EmojiDeleteNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this emoji delete not acceptable response has a 3xx status code
func (o *EmojiDeleteNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emoji delete not acceptable response has a 4xx status code
func (o *EmojiDeleteNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this emoji delete not acceptable response has a 5xx status code
func (o *EmojiDeleteNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this emoji delete not acceptable response a status code equal to that given
func (o *EmojiDeleteNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the emoji delete not acceptable response
func (o *EmojiDeleteNotAcceptable) Code() int {
	return 406
}

func (o *EmojiDeleteNotAcceptable) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/custom_emojis/{id}][%d] emojiDeleteNotAcceptable ", 406)
}

func (o *EmojiDeleteNotAcceptable) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/custom_emojis/{id}][%d] emojiDeleteNotAcceptable ", 406)
}

func (o *EmojiDeleteNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmojiDeleteInternalServerError creates a EmojiDeleteInternalServerError with default headers values
func NewEmojiDeleteInternalServerError() *EmojiDeleteInternalServerError {
	return &EmojiDeleteInternalServerError{}
}

/*
EmojiDeleteInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type EmojiDeleteInternalServerError struct {
}

// IsSuccess returns true when this emoji delete internal server error response has a 2xx status code
func (o *EmojiDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this emoji delete internal server error response has a 3xx status code
func (o *EmojiDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emoji delete internal server error response has a 4xx status code
func (o *EmojiDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this emoji delete internal server error response has a 5xx status code
func (o *EmojiDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this emoji delete internal server error response a status code equal to that given
func (o *EmojiDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the emoji delete internal server error response
func (o *EmojiDeleteInternalServerError) Code() int {
	return 500
}

func (o *EmojiDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/custom_emojis/{id}][%d] emojiDeleteInternalServerError ", 500)
}

func (o *EmojiDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/custom_emojis/{id}][%d] emojiDeleteInternalServerError ", 500)
}

func (o *EmojiDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
