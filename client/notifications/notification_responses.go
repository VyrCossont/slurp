// Code generated by go-swagger; DO NOT EDIT.

package notifications

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

// NotificationReader is a Reader for the Notification structure.
type NotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNotificationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNotificationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNotificationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewNotificationNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNotificationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/notification/{id}] notification", response, response.Code())
	}
}

// NewNotificationOK creates a NotificationOK with default headers values
func NewNotificationOK() *NotificationOK {
	return &NotificationOK{}
}

/*
NotificationOK describes a response with status code 200, with default header values.

Requested notification.
*/
type NotificationOK struct {
	Payload *models.Notification
}

// IsSuccess returns true when this notification o k response has a 2xx status code
func (o *NotificationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this notification o k response has a 3xx status code
func (o *NotificationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notification o k response has a 4xx status code
func (o *NotificationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this notification o k response has a 5xx status code
func (o *NotificationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this notification o k response a status code equal to that given
func (o *NotificationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the notification o k response
func (o *NotificationOK) Code() int {
	return 200
}

func (o *NotificationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/notification/{id}][%d] notificationOK %s", 200, payload)
}

func (o *NotificationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/notification/{id}][%d] notificationOK %s", 200, payload)
}

func (o *NotificationOK) GetPayload() *models.Notification {
	return o.Payload
}

func (o *NotificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Notification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationBadRequest creates a NotificationBadRequest with default headers values
func NewNotificationBadRequest() *NotificationBadRequest {
	return &NotificationBadRequest{}
}

/*
NotificationBadRequest describes a response with status code 400, with default header values.

bad request
*/
type NotificationBadRequest struct {
}

// IsSuccess returns true when this notification bad request response has a 2xx status code
func (o *NotificationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notification bad request response has a 3xx status code
func (o *NotificationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notification bad request response has a 4xx status code
func (o *NotificationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this notification bad request response has a 5xx status code
func (o *NotificationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this notification bad request response a status code equal to that given
func (o *NotificationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the notification bad request response
func (o *NotificationBadRequest) Code() int {
	return 400
}

func (o *NotificationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/notification/{id}][%d] notificationBadRequest", 400)
}

func (o *NotificationBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/notification/{id}][%d] notificationBadRequest", 400)
}

func (o *NotificationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNotificationUnauthorized creates a NotificationUnauthorized with default headers values
func NewNotificationUnauthorized() *NotificationUnauthorized {
	return &NotificationUnauthorized{}
}

/*
NotificationUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type NotificationUnauthorized struct {
}

// IsSuccess returns true when this notification unauthorized response has a 2xx status code
func (o *NotificationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notification unauthorized response has a 3xx status code
func (o *NotificationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notification unauthorized response has a 4xx status code
func (o *NotificationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this notification unauthorized response has a 5xx status code
func (o *NotificationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this notification unauthorized response a status code equal to that given
func (o *NotificationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the notification unauthorized response
func (o *NotificationUnauthorized) Code() int {
	return 401
}

func (o *NotificationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/notification/{id}][%d] notificationUnauthorized", 401)
}

func (o *NotificationUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/notification/{id}][%d] notificationUnauthorized", 401)
}

func (o *NotificationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNotificationNotFound creates a NotificationNotFound with default headers values
func NewNotificationNotFound() *NotificationNotFound {
	return &NotificationNotFound{}
}

/*
NotificationNotFound describes a response with status code 404, with default header values.

not found
*/
type NotificationNotFound struct {
}

// IsSuccess returns true when this notification not found response has a 2xx status code
func (o *NotificationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notification not found response has a 3xx status code
func (o *NotificationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notification not found response has a 4xx status code
func (o *NotificationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this notification not found response has a 5xx status code
func (o *NotificationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this notification not found response a status code equal to that given
func (o *NotificationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the notification not found response
func (o *NotificationNotFound) Code() int {
	return 404
}

func (o *NotificationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/notification/{id}][%d] notificationNotFound", 404)
}

func (o *NotificationNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/notification/{id}][%d] notificationNotFound", 404)
}

func (o *NotificationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNotificationNotAcceptable creates a NotificationNotAcceptable with default headers values
func NewNotificationNotAcceptable() *NotificationNotAcceptable {
	return &NotificationNotAcceptable{}
}

/*
NotificationNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type NotificationNotAcceptable struct {
}

// IsSuccess returns true when this notification not acceptable response has a 2xx status code
func (o *NotificationNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notification not acceptable response has a 3xx status code
func (o *NotificationNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notification not acceptable response has a 4xx status code
func (o *NotificationNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this notification not acceptable response has a 5xx status code
func (o *NotificationNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this notification not acceptable response a status code equal to that given
func (o *NotificationNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the notification not acceptable response
func (o *NotificationNotAcceptable) Code() int {
	return 406
}

func (o *NotificationNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/notification/{id}][%d] notificationNotAcceptable", 406)
}

func (o *NotificationNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/notification/{id}][%d] notificationNotAcceptable", 406)
}

func (o *NotificationNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNotificationInternalServerError creates a NotificationInternalServerError with default headers values
func NewNotificationInternalServerError() *NotificationInternalServerError {
	return &NotificationInternalServerError{}
}

/*
NotificationInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type NotificationInternalServerError struct {
}

// IsSuccess returns true when this notification internal server error response has a 2xx status code
func (o *NotificationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notification internal server error response has a 3xx status code
func (o *NotificationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notification internal server error response has a 4xx status code
func (o *NotificationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this notification internal server error response has a 5xx status code
func (o *NotificationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this notification internal server error response a status code equal to that given
func (o *NotificationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the notification internal server error response
func (o *NotificationInternalServerError) Code() int {
	return 500
}

func (o *NotificationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/notification/{id}][%d] notificationInternalServerError", 500)
}

func (o *NotificationInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/notification/{id}][%d] notificationInternalServerError", 500)
}

func (o *NotificationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
