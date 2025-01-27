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

// EmojiCategoriesGetReader is a Reader for the EmojiCategoriesGet structure.
type EmojiCategoriesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmojiCategoriesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmojiCategoriesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEmojiCategoriesGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEmojiCategoriesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEmojiCategoriesGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEmojiCategoriesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewEmojiCategoriesGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEmojiCategoriesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/admin/custom_emojis/categories] emojiCategoriesGet", response, response.Code())
	}
}

// NewEmojiCategoriesGetOK creates a EmojiCategoriesGetOK with default headers values
func NewEmojiCategoriesGetOK() *EmojiCategoriesGetOK {
	return &EmojiCategoriesGetOK{}
}

/*
EmojiCategoriesGetOK describes a response with status code 200, with default header values.

Array of existing emoji categories.
*/
type EmojiCategoriesGetOK struct {
	Payload []*models.EmojiCategory
}

// IsSuccess returns true when this emoji categories get o k response has a 2xx status code
func (o *EmojiCategoriesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this emoji categories get o k response has a 3xx status code
func (o *EmojiCategoriesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emoji categories get o k response has a 4xx status code
func (o *EmojiCategoriesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this emoji categories get o k response has a 5xx status code
func (o *EmojiCategoriesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this emoji categories get o k response a status code equal to that given
func (o *EmojiCategoriesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the emoji categories get o k response
func (o *EmojiCategoriesGetOK) Code() int {
	return 200
}

func (o *EmojiCategoriesGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/admin/custom_emojis/categories][%d] emojiCategoriesGetOK %s", 200, payload)
}

func (o *EmojiCategoriesGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/admin/custom_emojis/categories][%d] emojiCategoriesGetOK %s", 200, payload)
}

func (o *EmojiCategoriesGetOK) GetPayload() []*models.EmojiCategory {
	return o.Payload
}

func (o *EmojiCategoriesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmojiCategoriesGetBadRequest creates a EmojiCategoriesGetBadRequest with default headers values
func NewEmojiCategoriesGetBadRequest() *EmojiCategoriesGetBadRequest {
	return &EmojiCategoriesGetBadRequest{}
}

/*
EmojiCategoriesGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type EmojiCategoriesGetBadRequest struct {
}

// IsSuccess returns true when this emoji categories get bad request response has a 2xx status code
func (o *EmojiCategoriesGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this emoji categories get bad request response has a 3xx status code
func (o *EmojiCategoriesGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emoji categories get bad request response has a 4xx status code
func (o *EmojiCategoriesGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this emoji categories get bad request response has a 5xx status code
func (o *EmojiCategoriesGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this emoji categories get bad request response a status code equal to that given
func (o *EmojiCategoriesGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the emoji categories get bad request response
func (o *EmojiCategoriesGetBadRequest) Code() int {
	return 400
}

func (o *EmojiCategoriesGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/custom_emojis/categories][%d] emojiCategoriesGetBadRequest", 400)
}

func (o *EmojiCategoriesGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/custom_emojis/categories][%d] emojiCategoriesGetBadRequest", 400)
}

func (o *EmojiCategoriesGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmojiCategoriesGetUnauthorized creates a EmojiCategoriesGetUnauthorized with default headers values
func NewEmojiCategoriesGetUnauthorized() *EmojiCategoriesGetUnauthorized {
	return &EmojiCategoriesGetUnauthorized{}
}

/*
EmojiCategoriesGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type EmojiCategoriesGetUnauthorized struct {
}

// IsSuccess returns true when this emoji categories get unauthorized response has a 2xx status code
func (o *EmojiCategoriesGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this emoji categories get unauthorized response has a 3xx status code
func (o *EmojiCategoriesGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emoji categories get unauthorized response has a 4xx status code
func (o *EmojiCategoriesGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this emoji categories get unauthorized response has a 5xx status code
func (o *EmojiCategoriesGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this emoji categories get unauthorized response a status code equal to that given
func (o *EmojiCategoriesGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the emoji categories get unauthorized response
func (o *EmojiCategoriesGetUnauthorized) Code() int {
	return 401
}

func (o *EmojiCategoriesGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/custom_emojis/categories][%d] emojiCategoriesGetUnauthorized", 401)
}

func (o *EmojiCategoriesGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/custom_emojis/categories][%d] emojiCategoriesGetUnauthorized", 401)
}

func (o *EmojiCategoriesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmojiCategoriesGetForbidden creates a EmojiCategoriesGetForbidden with default headers values
func NewEmojiCategoriesGetForbidden() *EmojiCategoriesGetForbidden {
	return &EmojiCategoriesGetForbidden{}
}

/*
EmojiCategoriesGetForbidden describes a response with status code 403, with default header values.

forbidden
*/
type EmojiCategoriesGetForbidden struct {
}

// IsSuccess returns true when this emoji categories get forbidden response has a 2xx status code
func (o *EmojiCategoriesGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this emoji categories get forbidden response has a 3xx status code
func (o *EmojiCategoriesGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emoji categories get forbidden response has a 4xx status code
func (o *EmojiCategoriesGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this emoji categories get forbidden response has a 5xx status code
func (o *EmojiCategoriesGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this emoji categories get forbidden response a status code equal to that given
func (o *EmojiCategoriesGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the emoji categories get forbidden response
func (o *EmojiCategoriesGetForbidden) Code() int {
	return 403
}

func (o *EmojiCategoriesGetForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/custom_emojis/categories][%d] emojiCategoriesGetForbidden", 403)
}

func (o *EmojiCategoriesGetForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/custom_emojis/categories][%d] emojiCategoriesGetForbidden", 403)
}

func (o *EmojiCategoriesGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmojiCategoriesGetNotFound creates a EmojiCategoriesGetNotFound with default headers values
func NewEmojiCategoriesGetNotFound() *EmojiCategoriesGetNotFound {
	return &EmojiCategoriesGetNotFound{}
}

/*
EmojiCategoriesGetNotFound describes a response with status code 404, with default header values.

not found
*/
type EmojiCategoriesGetNotFound struct {
}

// IsSuccess returns true when this emoji categories get not found response has a 2xx status code
func (o *EmojiCategoriesGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this emoji categories get not found response has a 3xx status code
func (o *EmojiCategoriesGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emoji categories get not found response has a 4xx status code
func (o *EmojiCategoriesGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this emoji categories get not found response has a 5xx status code
func (o *EmojiCategoriesGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this emoji categories get not found response a status code equal to that given
func (o *EmojiCategoriesGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the emoji categories get not found response
func (o *EmojiCategoriesGetNotFound) Code() int {
	return 404
}

func (o *EmojiCategoriesGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/custom_emojis/categories][%d] emojiCategoriesGetNotFound", 404)
}

func (o *EmojiCategoriesGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/custom_emojis/categories][%d] emojiCategoriesGetNotFound", 404)
}

func (o *EmojiCategoriesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmojiCategoriesGetNotAcceptable creates a EmojiCategoriesGetNotAcceptable with default headers values
func NewEmojiCategoriesGetNotAcceptable() *EmojiCategoriesGetNotAcceptable {
	return &EmojiCategoriesGetNotAcceptable{}
}

/*
EmojiCategoriesGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type EmojiCategoriesGetNotAcceptable struct {
}

// IsSuccess returns true when this emoji categories get not acceptable response has a 2xx status code
func (o *EmojiCategoriesGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this emoji categories get not acceptable response has a 3xx status code
func (o *EmojiCategoriesGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emoji categories get not acceptable response has a 4xx status code
func (o *EmojiCategoriesGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this emoji categories get not acceptable response has a 5xx status code
func (o *EmojiCategoriesGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this emoji categories get not acceptable response a status code equal to that given
func (o *EmojiCategoriesGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the emoji categories get not acceptable response
func (o *EmojiCategoriesGetNotAcceptable) Code() int {
	return 406
}

func (o *EmojiCategoriesGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/custom_emojis/categories][%d] emojiCategoriesGetNotAcceptable", 406)
}

func (o *EmojiCategoriesGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/custom_emojis/categories][%d] emojiCategoriesGetNotAcceptable", 406)
}

func (o *EmojiCategoriesGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmojiCategoriesGetInternalServerError creates a EmojiCategoriesGetInternalServerError with default headers values
func NewEmojiCategoriesGetInternalServerError() *EmojiCategoriesGetInternalServerError {
	return &EmojiCategoriesGetInternalServerError{}
}

/*
EmojiCategoriesGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type EmojiCategoriesGetInternalServerError struct {
}

// IsSuccess returns true when this emoji categories get internal server error response has a 2xx status code
func (o *EmojiCategoriesGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this emoji categories get internal server error response has a 3xx status code
func (o *EmojiCategoriesGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emoji categories get internal server error response has a 4xx status code
func (o *EmojiCategoriesGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this emoji categories get internal server error response has a 5xx status code
func (o *EmojiCategoriesGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this emoji categories get internal server error response a status code equal to that given
func (o *EmojiCategoriesGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the emoji categories get internal server error response
func (o *EmojiCategoriesGetInternalServerError) Code() int {
	return 500
}

func (o *EmojiCategoriesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/custom_emojis/categories][%d] emojiCategoriesGetInternalServerError", 500)
}

func (o *EmojiCategoriesGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/custom_emojis/categories][%d] emojiCategoriesGetInternalServerError", 500)
}

func (o *EmojiCategoriesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
