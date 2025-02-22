// Code generated by go-swagger; DO NOT EDIT.

package interaction_requests

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

// GetInteractionRequestsReader is a Reader for the GetInteractionRequests structure.
type GetInteractionRequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInteractionRequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInteractionRequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInteractionRequestsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInteractionRequestsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInteractionRequestsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetInteractionRequestsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInteractionRequestsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/interaction_requests] getInteractionRequests", response, response.Code())
	}
}

// NewGetInteractionRequestsOK creates a GetInteractionRequestsOK with default headers values
func NewGetInteractionRequestsOK() *GetInteractionRequestsOK {
	return &GetInteractionRequestsOK{}
}

/*
GetInteractionRequestsOK describes a response with status code 200, with default header values.

GetInteractionRequestsOK get interaction requests o k
*/
type GetInteractionRequestsOK struct {

	/* Links to the next and previous queries.
	 */
	Link string

	Payload []*models.InteractionRequest
}

// IsSuccess returns true when this get interaction requests o k response has a 2xx status code
func (o *GetInteractionRequestsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get interaction requests o k response has a 3xx status code
func (o *GetInteractionRequestsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get interaction requests o k response has a 4xx status code
func (o *GetInteractionRequestsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get interaction requests o k response has a 5xx status code
func (o *GetInteractionRequestsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get interaction requests o k response a status code equal to that given
func (o *GetInteractionRequestsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get interaction requests o k response
func (o *GetInteractionRequestsOK) Code() int {
	return 200
}

func (o *GetInteractionRequestsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/interaction_requests][%d] getInteractionRequestsOK %s", 200, payload)
}

func (o *GetInteractionRequestsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/interaction_requests][%d] getInteractionRequestsOK %s", 200, payload)
}

func (o *GetInteractionRequestsOK) GetPayload() []*models.InteractionRequest {
	return o.Payload
}

func (o *GetInteractionRequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetInteractionRequestsBadRequest creates a GetInteractionRequestsBadRequest with default headers values
func NewGetInteractionRequestsBadRequest() *GetInteractionRequestsBadRequest {
	return &GetInteractionRequestsBadRequest{}
}

/*
GetInteractionRequestsBadRequest describes a response with status code 400, with default header values.

bad request
*/
type GetInteractionRequestsBadRequest struct {
}

// IsSuccess returns true when this get interaction requests bad request response has a 2xx status code
func (o *GetInteractionRequestsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get interaction requests bad request response has a 3xx status code
func (o *GetInteractionRequestsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get interaction requests bad request response has a 4xx status code
func (o *GetInteractionRequestsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get interaction requests bad request response has a 5xx status code
func (o *GetInteractionRequestsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get interaction requests bad request response a status code equal to that given
func (o *GetInteractionRequestsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get interaction requests bad request response
func (o *GetInteractionRequestsBadRequest) Code() int {
	return 400
}

func (o *GetInteractionRequestsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/interaction_requests][%d] getInteractionRequestsBadRequest", 400)
}

func (o *GetInteractionRequestsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/interaction_requests][%d] getInteractionRequestsBadRequest", 400)
}

func (o *GetInteractionRequestsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInteractionRequestsUnauthorized creates a GetInteractionRequestsUnauthorized with default headers values
func NewGetInteractionRequestsUnauthorized() *GetInteractionRequestsUnauthorized {
	return &GetInteractionRequestsUnauthorized{}
}

/*
GetInteractionRequestsUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type GetInteractionRequestsUnauthorized struct {
}

// IsSuccess returns true when this get interaction requests unauthorized response has a 2xx status code
func (o *GetInteractionRequestsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get interaction requests unauthorized response has a 3xx status code
func (o *GetInteractionRequestsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get interaction requests unauthorized response has a 4xx status code
func (o *GetInteractionRequestsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get interaction requests unauthorized response has a 5xx status code
func (o *GetInteractionRequestsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get interaction requests unauthorized response a status code equal to that given
func (o *GetInteractionRequestsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get interaction requests unauthorized response
func (o *GetInteractionRequestsUnauthorized) Code() int {
	return 401
}

func (o *GetInteractionRequestsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/interaction_requests][%d] getInteractionRequestsUnauthorized", 401)
}

func (o *GetInteractionRequestsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/interaction_requests][%d] getInteractionRequestsUnauthorized", 401)
}

func (o *GetInteractionRequestsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInteractionRequestsNotFound creates a GetInteractionRequestsNotFound with default headers values
func NewGetInteractionRequestsNotFound() *GetInteractionRequestsNotFound {
	return &GetInteractionRequestsNotFound{}
}

/*
GetInteractionRequestsNotFound describes a response with status code 404, with default header values.

not found
*/
type GetInteractionRequestsNotFound struct {
}

// IsSuccess returns true when this get interaction requests not found response has a 2xx status code
func (o *GetInteractionRequestsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get interaction requests not found response has a 3xx status code
func (o *GetInteractionRequestsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get interaction requests not found response has a 4xx status code
func (o *GetInteractionRequestsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get interaction requests not found response has a 5xx status code
func (o *GetInteractionRequestsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get interaction requests not found response a status code equal to that given
func (o *GetInteractionRequestsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get interaction requests not found response
func (o *GetInteractionRequestsNotFound) Code() int {
	return 404
}

func (o *GetInteractionRequestsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/interaction_requests][%d] getInteractionRequestsNotFound", 404)
}

func (o *GetInteractionRequestsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/interaction_requests][%d] getInteractionRequestsNotFound", 404)
}

func (o *GetInteractionRequestsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInteractionRequestsNotAcceptable creates a GetInteractionRequestsNotAcceptable with default headers values
func NewGetInteractionRequestsNotAcceptable() *GetInteractionRequestsNotAcceptable {
	return &GetInteractionRequestsNotAcceptable{}
}

/*
GetInteractionRequestsNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type GetInteractionRequestsNotAcceptable struct {
}

// IsSuccess returns true when this get interaction requests not acceptable response has a 2xx status code
func (o *GetInteractionRequestsNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get interaction requests not acceptable response has a 3xx status code
func (o *GetInteractionRequestsNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get interaction requests not acceptable response has a 4xx status code
func (o *GetInteractionRequestsNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this get interaction requests not acceptable response has a 5xx status code
func (o *GetInteractionRequestsNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this get interaction requests not acceptable response a status code equal to that given
func (o *GetInteractionRequestsNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the get interaction requests not acceptable response
func (o *GetInteractionRequestsNotAcceptable) Code() int {
	return 406
}

func (o *GetInteractionRequestsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/interaction_requests][%d] getInteractionRequestsNotAcceptable", 406)
}

func (o *GetInteractionRequestsNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/interaction_requests][%d] getInteractionRequestsNotAcceptable", 406)
}

func (o *GetInteractionRequestsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInteractionRequestsInternalServerError creates a GetInteractionRequestsInternalServerError with default headers values
func NewGetInteractionRequestsInternalServerError() *GetInteractionRequestsInternalServerError {
	return &GetInteractionRequestsInternalServerError{}
}

/*
GetInteractionRequestsInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type GetInteractionRequestsInternalServerError struct {
}

// IsSuccess returns true when this get interaction requests internal server error response has a 2xx status code
func (o *GetInteractionRequestsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get interaction requests internal server error response has a 3xx status code
func (o *GetInteractionRequestsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get interaction requests internal server error response has a 4xx status code
func (o *GetInteractionRequestsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get interaction requests internal server error response has a 5xx status code
func (o *GetInteractionRequestsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get interaction requests internal server error response a status code equal to that given
func (o *GetInteractionRequestsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get interaction requests internal server error response
func (o *GetInteractionRequestsInternalServerError) Code() int {
	return 500
}

func (o *GetInteractionRequestsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/interaction_requests][%d] getInteractionRequestsInternalServerError", 500)
}

func (o *GetInteractionRequestsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/interaction_requests][%d] getInteractionRequestsInternalServerError", 500)
}

func (o *GetInteractionRequestsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
