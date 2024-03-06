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

// EmojiUpdateReader is a Reader for the EmojiUpdate structure.
type EmojiUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmojiUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmojiUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEmojiUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEmojiUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEmojiUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEmojiUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewEmojiUpdateNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEmojiUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /api/v1/admin/custom_emojis/{id}] emojiUpdate", response, response.Code())
	}
}

// NewEmojiUpdateOK creates a EmojiUpdateOK with default headers values
func NewEmojiUpdateOK() *EmojiUpdateOK {
	return &EmojiUpdateOK{}
}

/*
EmojiUpdateOK describes a response with status code 200, with default header values.

The updated emoji.
*/
type EmojiUpdateOK struct {
	Payload *models.AdminEmoji
}

// IsSuccess returns true when this emoji update o k response has a 2xx status code
func (o *EmojiUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this emoji update o k response has a 3xx status code
func (o *EmojiUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emoji update o k response has a 4xx status code
func (o *EmojiUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this emoji update o k response has a 5xx status code
func (o *EmojiUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this emoji update o k response a status code equal to that given
func (o *EmojiUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the emoji update o k response
func (o *EmojiUpdateOK) Code() int {
	return 200
}

func (o *EmojiUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/custom_emojis/{id}][%d] emojiUpdateOK  %+v", 200, o.Payload)
}

func (o *EmojiUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/custom_emojis/{id}][%d] emojiUpdateOK  %+v", 200, o.Payload)
}

func (o *EmojiUpdateOK) GetPayload() *models.AdminEmoji {
	return o.Payload
}

func (o *EmojiUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdminEmoji)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmojiUpdateBadRequest creates a EmojiUpdateBadRequest with default headers values
func NewEmojiUpdateBadRequest() *EmojiUpdateBadRequest {
	return &EmojiUpdateBadRequest{}
}

/*
EmojiUpdateBadRequest describes a response with status code 400, with default header values.

bad request
*/
type EmojiUpdateBadRequest struct {
}

// IsSuccess returns true when this emoji update bad request response has a 2xx status code
func (o *EmojiUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this emoji update bad request response has a 3xx status code
func (o *EmojiUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emoji update bad request response has a 4xx status code
func (o *EmojiUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this emoji update bad request response has a 5xx status code
func (o *EmojiUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this emoji update bad request response a status code equal to that given
func (o *EmojiUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the emoji update bad request response
func (o *EmojiUpdateBadRequest) Code() int {
	return 400
}

func (o *EmojiUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/custom_emojis/{id}][%d] emojiUpdateBadRequest ", 400)
}

func (o *EmojiUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/custom_emojis/{id}][%d] emojiUpdateBadRequest ", 400)
}

func (o *EmojiUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmojiUpdateUnauthorized creates a EmojiUpdateUnauthorized with default headers values
func NewEmojiUpdateUnauthorized() *EmojiUpdateUnauthorized {
	return &EmojiUpdateUnauthorized{}
}

/*
EmojiUpdateUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type EmojiUpdateUnauthorized struct {
}

// IsSuccess returns true when this emoji update unauthorized response has a 2xx status code
func (o *EmojiUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this emoji update unauthorized response has a 3xx status code
func (o *EmojiUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emoji update unauthorized response has a 4xx status code
func (o *EmojiUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this emoji update unauthorized response has a 5xx status code
func (o *EmojiUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this emoji update unauthorized response a status code equal to that given
func (o *EmojiUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the emoji update unauthorized response
func (o *EmojiUpdateUnauthorized) Code() int {
	return 401
}

func (o *EmojiUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/custom_emojis/{id}][%d] emojiUpdateUnauthorized ", 401)
}

func (o *EmojiUpdateUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/custom_emojis/{id}][%d] emojiUpdateUnauthorized ", 401)
}

func (o *EmojiUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmojiUpdateForbidden creates a EmojiUpdateForbidden with default headers values
func NewEmojiUpdateForbidden() *EmojiUpdateForbidden {
	return &EmojiUpdateForbidden{}
}

/*
EmojiUpdateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type EmojiUpdateForbidden struct {
}

// IsSuccess returns true when this emoji update forbidden response has a 2xx status code
func (o *EmojiUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this emoji update forbidden response has a 3xx status code
func (o *EmojiUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emoji update forbidden response has a 4xx status code
func (o *EmojiUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this emoji update forbidden response has a 5xx status code
func (o *EmojiUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this emoji update forbidden response a status code equal to that given
func (o *EmojiUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the emoji update forbidden response
func (o *EmojiUpdateForbidden) Code() int {
	return 403
}

func (o *EmojiUpdateForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/custom_emojis/{id}][%d] emojiUpdateForbidden ", 403)
}

func (o *EmojiUpdateForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/custom_emojis/{id}][%d] emojiUpdateForbidden ", 403)
}

func (o *EmojiUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmojiUpdateNotFound creates a EmojiUpdateNotFound with default headers values
func NewEmojiUpdateNotFound() *EmojiUpdateNotFound {
	return &EmojiUpdateNotFound{}
}

/*
EmojiUpdateNotFound describes a response with status code 404, with default header values.

not found
*/
type EmojiUpdateNotFound struct {
}

// IsSuccess returns true when this emoji update not found response has a 2xx status code
func (o *EmojiUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this emoji update not found response has a 3xx status code
func (o *EmojiUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emoji update not found response has a 4xx status code
func (o *EmojiUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this emoji update not found response has a 5xx status code
func (o *EmojiUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this emoji update not found response a status code equal to that given
func (o *EmojiUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the emoji update not found response
func (o *EmojiUpdateNotFound) Code() int {
	return 404
}

func (o *EmojiUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/custom_emojis/{id}][%d] emojiUpdateNotFound ", 404)
}

func (o *EmojiUpdateNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/custom_emojis/{id}][%d] emojiUpdateNotFound ", 404)
}

func (o *EmojiUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmojiUpdateNotAcceptable creates a EmojiUpdateNotAcceptable with default headers values
func NewEmojiUpdateNotAcceptable() *EmojiUpdateNotAcceptable {
	return &EmojiUpdateNotAcceptable{}
}

/*
EmojiUpdateNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type EmojiUpdateNotAcceptable struct {
}

// IsSuccess returns true when this emoji update not acceptable response has a 2xx status code
func (o *EmojiUpdateNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this emoji update not acceptable response has a 3xx status code
func (o *EmojiUpdateNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emoji update not acceptable response has a 4xx status code
func (o *EmojiUpdateNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this emoji update not acceptable response has a 5xx status code
func (o *EmojiUpdateNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this emoji update not acceptable response a status code equal to that given
func (o *EmojiUpdateNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the emoji update not acceptable response
func (o *EmojiUpdateNotAcceptable) Code() int {
	return 406
}

func (o *EmojiUpdateNotAcceptable) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/custom_emojis/{id}][%d] emojiUpdateNotAcceptable ", 406)
}

func (o *EmojiUpdateNotAcceptable) String() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/custom_emojis/{id}][%d] emojiUpdateNotAcceptable ", 406)
}

func (o *EmojiUpdateNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmojiUpdateInternalServerError creates a EmojiUpdateInternalServerError with default headers values
func NewEmojiUpdateInternalServerError() *EmojiUpdateInternalServerError {
	return &EmojiUpdateInternalServerError{}
}

/*
EmojiUpdateInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type EmojiUpdateInternalServerError struct {
}

// IsSuccess returns true when this emoji update internal server error response has a 2xx status code
func (o *EmojiUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this emoji update internal server error response has a 3xx status code
func (o *EmojiUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emoji update internal server error response has a 4xx status code
func (o *EmojiUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this emoji update internal server error response has a 5xx status code
func (o *EmojiUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this emoji update internal server error response a status code equal to that given
func (o *EmojiUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the emoji update internal server error response
func (o *EmojiUpdateInternalServerError) Code() int {
	return 500
}

func (o *EmojiUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/custom_emojis/{id}][%d] emojiUpdateInternalServerError ", 500)
}

func (o *EmojiUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/custom_emojis/{id}][%d] emojiUpdateInternalServerError ", 500)
}

func (o *EmojiUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
