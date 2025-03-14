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

// NotificationsReader is a Reader for the Notifications structure.
type NotificationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotificationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNotificationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNotificationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNotificationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewNotificationsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNotificationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/notifications] notifications", response, response.Code())
	}
}

// NewNotificationsOK creates a NotificationsOK with default headers values
func NewNotificationsOK() *NotificationsOK {
	return &NotificationsOK{}
}

/*
NotificationsOK describes a response with status code 200, with default header values.

Array of notifications.
*/
type NotificationsOK struct {

	/* Links to the next and previous queries.
	 */
	Link string

	Payload []*models.Notification
}

// IsSuccess returns true when this notifications o k response has a 2xx status code
func (o *NotificationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this notifications o k response has a 3xx status code
func (o *NotificationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications o k response has a 4xx status code
func (o *NotificationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this notifications o k response has a 5xx status code
func (o *NotificationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications o k response a status code equal to that given
func (o *NotificationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the notifications o k response
func (o *NotificationsOK) Code() int {
	return 200
}

func (o *NotificationsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/notifications][%d] notificationsOK %s", 200, payload)
}

func (o *NotificationsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/notifications][%d] notificationsOK %s", 200, payload)
}

func (o *NotificationsOK) GetPayload() []*models.Notification {
	return o.Payload
}

func (o *NotificationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsBadRequest creates a NotificationsBadRequest with default headers values
func NewNotificationsBadRequest() *NotificationsBadRequest {
	return &NotificationsBadRequest{}
}

/*
NotificationsBadRequest describes a response with status code 400, with default header values.

bad request
*/
type NotificationsBadRequest struct {
}

// IsSuccess returns true when this notifications bad request response has a 2xx status code
func (o *NotificationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications bad request response has a 3xx status code
func (o *NotificationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications bad request response has a 4xx status code
func (o *NotificationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this notifications bad request response has a 5xx status code
func (o *NotificationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications bad request response a status code equal to that given
func (o *NotificationsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the notifications bad request response
func (o *NotificationsBadRequest) Code() int {
	return 400
}

func (o *NotificationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/notifications][%d] notificationsBadRequest", 400)
}

func (o *NotificationsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/notifications][%d] notificationsBadRequest", 400)
}

func (o *NotificationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNotificationsUnauthorized creates a NotificationsUnauthorized with default headers values
func NewNotificationsUnauthorized() *NotificationsUnauthorized {
	return &NotificationsUnauthorized{}
}

/*
NotificationsUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type NotificationsUnauthorized struct {
}

// IsSuccess returns true when this notifications unauthorized response has a 2xx status code
func (o *NotificationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications unauthorized response has a 3xx status code
func (o *NotificationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications unauthorized response has a 4xx status code
func (o *NotificationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this notifications unauthorized response has a 5xx status code
func (o *NotificationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications unauthorized response a status code equal to that given
func (o *NotificationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the notifications unauthorized response
func (o *NotificationsUnauthorized) Code() int {
	return 401
}

func (o *NotificationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/notifications][%d] notificationsUnauthorized", 401)
}

func (o *NotificationsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/notifications][%d] notificationsUnauthorized", 401)
}

func (o *NotificationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNotificationsNotFound creates a NotificationsNotFound with default headers values
func NewNotificationsNotFound() *NotificationsNotFound {
	return &NotificationsNotFound{}
}

/*
NotificationsNotFound describes a response with status code 404, with default header values.

not found
*/
type NotificationsNotFound struct {
}

// IsSuccess returns true when this notifications not found response has a 2xx status code
func (o *NotificationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications not found response has a 3xx status code
func (o *NotificationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications not found response has a 4xx status code
func (o *NotificationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this notifications not found response has a 5xx status code
func (o *NotificationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications not found response a status code equal to that given
func (o *NotificationsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the notifications not found response
func (o *NotificationsNotFound) Code() int {
	return 404
}

func (o *NotificationsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/notifications][%d] notificationsNotFound", 404)
}

func (o *NotificationsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/notifications][%d] notificationsNotFound", 404)
}

func (o *NotificationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNotificationsNotAcceptable creates a NotificationsNotAcceptable with default headers values
func NewNotificationsNotAcceptable() *NotificationsNotAcceptable {
	return &NotificationsNotAcceptable{}
}

/*
NotificationsNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type NotificationsNotAcceptable struct {
}

// IsSuccess returns true when this notifications not acceptable response has a 2xx status code
func (o *NotificationsNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications not acceptable response has a 3xx status code
func (o *NotificationsNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications not acceptable response has a 4xx status code
func (o *NotificationsNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this notifications not acceptable response has a 5xx status code
func (o *NotificationsNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications not acceptable response a status code equal to that given
func (o *NotificationsNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the notifications not acceptable response
func (o *NotificationsNotAcceptable) Code() int {
	return 406
}

func (o *NotificationsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/notifications][%d] notificationsNotAcceptable", 406)
}

func (o *NotificationsNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/notifications][%d] notificationsNotAcceptable", 406)
}

func (o *NotificationsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNotificationsInternalServerError creates a NotificationsInternalServerError with default headers values
func NewNotificationsInternalServerError() *NotificationsInternalServerError {
	return &NotificationsInternalServerError{}
}

/*
NotificationsInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type NotificationsInternalServerError struct {
}

// IsSuccess returns true when this notifications internal server error response has a 2xx status code
func (o *NotificationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications internal server error response has a 3xx status code
func (o *NotificationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications internal server error response has a 4xx status code
func (o *NotificationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this notifications internal server error response has a 5xx status code
func (o *NotificationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this notifications internal server error response a status code equal to that given
func (o *NotificationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the notifications internal server error response
func (o *NotificationsInternalServerError) Code() int {
	return 500
}

func (o *NotificationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/notifications][%d] notificationsInternalServerError", 500)
}

func (o *NotificationsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/notifications][%d] notificationsInternalServerError", 500)
}

func (o *NotificationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
