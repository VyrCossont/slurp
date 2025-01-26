// Code generated by go-swagger; DO NOT EDIT.

package conversations

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

// ConversationReadReader is a Reader for the ConversationRead structure.
type ConversationReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConversationReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConversationReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewConversationReadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewConversationReadUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewConversationReadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewConversationReadNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewConversationReadUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewConversationReadInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/conversation/{id}/read] conversationRead", response, response.Code())
	}
}

// NewConversationReadOK creates a ConversationReadOK with default headers values
func NewConversationReadOK() *ConversationReadOK {
	return &ConversationReadOK{}
}

/*
ConversationReadOK describes a response with status code 200, with default header values.

Updated conversation.
*/
type ConversationReadOK struct {
	Payload *models.Conversation
}

// IsSuccess returns true when this conversation read o k response has a 2xx status code
func (o *ConversationReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this conversation read o k response has a 3xx status code
func (o *ConversationReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this conversation read o k response has a 4xx status code
func (o *ConversationReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this conversation read o k response has a 5xx status code
func (o *ConversationReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this conversation read o k response a status code equal to that given
func (o *ConversationReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the conversation read o k response
func (o *ConversationReadOK) Code() int {
	return 200
}

func (o *ConversationReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/conversation/{id}/read][%d] conversationReadOK %s", 200, payload)
}

func (o *ConversationReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/conversation/{id}/read][%d] conversationReadOK %s", 200, payload)
}

func (o *ConversationReadOK) GetPayload() *models.Conversation {
	return o.Payload
}

func (o *ConversationReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Conversation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConversationReadBadRequest creates a ConversationReadBadRequest with default headers values
func NewConversationReadBadRequest() *ConversationReadBadRequest {
	return &ConversationReadBadRequest{}
}

/*
ConversationReadBadRequest describes a response with status code 400, with default header values.

bad request
*/
type ConversationReadBadRequest struct {
}

// IsSuccess returns true when this conversation read bad request response has a 2xx status code
func (o *ConversationReadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this conversation read bad request response has a 3xx status code
func (o *ConversationReadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this conversation read bad request response has a 4xx status code
func (o *ConversationReadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this conversation read bad request response has a 5xx status code
func (o *ConversationReadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this conversation read bad request response a status code equal to that given
func (o *ConversationReadBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the conversation read bad request response
func (o *ConversationReadBadRequest) Code() int {
	return 400
}

func (o *ConversationReadBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/conversation/{id}/read][%d] conversationReadBadRequest", 400)
}

func (o *ConversationReadBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/conversation/{id}/read][%d] conversationReadBadRequest", 400)
}

func (o *ConversationReadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConversationReadUnauthorized creates a ConversationReadUnauthorized with default headers values
func NewConversationReadUnauthorized() *ConversationReadUnauthorized {
	return &ConversationReadUnauthorized{}
}

/*
ConversationReadUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type ConversationReadUnauthorized struct {
}

// IsSuccess returns true when this conversation read unauthorized response has a 2xx status code
func (o *ConversationReadUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this conversation read unauthorized response has a 3xx status code
func (o *ConversationReadUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this conversation read unauthorized response has a 4xx status code
func (o *ConversationReadUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this conversation read unauthorized response has a 5xx status code
func (o *ConversationReadUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this conversation read unauthorized response a status code equal to that given
func (o *ConversationReadUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the conversation read unauthorized response
func (o *ConversationReadUnauthorized) Code() int {
	return 401
}

func (o *ConversationReadUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/conversation/{id}/read][%d] conversationReadUnauthorized", 401)
}

func (o *ConversationReadUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/conversation/{id}/read][%d] conversationReadUnauthorized", 401)
}

func (o *ConversationReadUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConversationReadNotFound creates a ConversationReadNotFound with default headers values
func NewConversationReadNotFound() *ConversationReadNotFound {
	return &ConversationReadNotFound{}
}

/*
ConversationReadNotFound describes a response with status code 404, with default header values.

not found
*/
type ConversationReadNotFound struct {
}

// IsSuccess returns true when this conversation read not found response has a 2xx status code
func (o *ConversationReadNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this conversation read not found response has a 3xx status code
func (o *ConversationReadNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this conversation read not found response has a 4xx status code
func (o *ConversationReadNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this conversation read not found response has a 5xx status code
func (o *ConversationReadNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this conversation read not found response a status code equal to that given
func (o *ConversationReadNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the conversation read not found response
func (o *ConversationReadNotFound) Code() int {
	return 404
}

func (o *ConversationReadNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/conversation/{id}/read][%d] conversationReadNotFound", 404)
}

func (o *ConversationReadNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/conversation/{id}/read][%d] conversationReadNotFound", 404)
}

func (o *ConversationReadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConversationReadNotAcceptable creates a ConversationReadNotAcceptable with default headers values
func NewConversationReadNotAcceptable() *ConversationReadNotAcceptable {
	return &ConversationReadNotAcceptable{}
}

/*
ConversationReadNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type ConversationReadNotAcceptable struct {
}

// IsSuccess returns true when this conversation read not acceptable response has a 2xx status code
func (o *ConversationReadNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this conversation read not acceptable response has a 3xx status code
func (o *ConversationReadNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this conversation read not acceptable response has a 4xx status code
func (o *ConversationReadNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this conversation read not acceptable response has a 5xx status code
func (o *ConversationReadNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this conversation read not acceptable response a status code equal to that given
func (o *ConversationReadNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the conversation read not acceptable response
func (o *ConversationReadNotAcceptable) Code() int {
	return 406
}

func (o *ConversationReadNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/conversation/{id}/read][%d] conversationReadNotAcceptable", 406)
}

func (o *ConversationReadNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v1/conversation/{id}/read][%d] conversationReadNotAcceptable", 406)
}

func (o *ConversationReadNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConversationReadUnprocessableEntity creates a ConversationReadUnprocessableEntity with default headers values
func NewConversationReadUnprocessableEntity() *ConversationReadUnprocessableEntity {
	return &ConversationReadUnprocessableEntity{}
}

/*
ConversationReadUnprocessableEntity describes a response with status code 422, with default header values.

unprocessable content
*/
type ConversationReadUnprocessableEntity struct {
}

// IsSuccess returns true when this conversation read unprocessable entity response has a 2xx status code
func (o *ConversationReadUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this conversation read unprocessable entity response has a 3xx status code
func (o *ConversationReadUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this conversation read unprocessable entity response has a 4xx status code
func (o *ConversationReadUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this conversation read unprocessable entity response has a 5xx status code
func (o *ConversationReadUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this conversation read unprocessable entity response a status code equal to that given
func (o *ConversationReadUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the conversation read unprocessable entity response
func (o *ConversationReadUnprocessableEntity) Code() int {
	return 422
}

func (o *ConversationReadUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v1/conversation/{id}/read][%d] conversationReadUnprocessableEntity", 422)
}

func (o *ConversationReadUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/v1/conversation/{id}/read][%d] conversationReadUnprocessableEntity", 422)
}

func (o *ConversationReadUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConversationReadInternalServerError creates a ConversationReadInternalServerError with default headers values
func NewConversationReadInternalServerError() *ConversationReadInternalServerError {
	return &ConversationReadInternalServerError{}
}

/*
ConversationReadInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type ConversationReadInternalServerError struct {
}

// IsSuccess returns true when this conversation read internal server error response has a 2xx status code
func (o *ConversationReadInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this conversation read internal server error response has a 3xx status code
func (o *ConversationReadInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this conversation read internal server error response has a 4xx status code
func (o *ConversationReadInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this conversation read internal server error response has a 5xx status code
func (o *ConversationReadInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this conversation read internal server error response a status code equal to that given
func (o *ConversationReadInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the conversation read internal server error response
func (o *ConversationReadInternalServerError) Code() int {
	return 500
}

func (o *ConversationReadInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/conversation/{id}/read][%d] conversationReadInternalServerError", 500)
}

func (o *ConversationReadInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/conversation/{id}/read][%d] conversationReadInternalServerError", 500)
}

func (o *ConversationReadInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
