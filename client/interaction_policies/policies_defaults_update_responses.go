// Code generated by go-swagger; DO NOT EDIT.

package interaction_policies

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

// PoliciesDefaultsUpdateReader is a Reader for the PoliciesDefaultsUpdate structure.
type PoliciesDefaultsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PoliciesDefaultsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPoliciesDefaultsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPoliciesDefaultsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPoliciesDefaultsUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewPoliciesDefaultsUpdateNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPoliciesDefaultsUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPoliciesDefaultsUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /api/v1/interaction_policies/defaults] policiesDefaultsUpdate", response, response.Code())
	}
}

// NewPoliciesDefaultsUpdateOK creates a PoliciesDefaultsUpdateOK with default headers values
func NewPoliciesDefaultsUpdateOK() *PoliciesDefaultsUpdateOK {
	return &PoliciesDefaultsUpdateOK{}
}

/*
PoliciesDefaultsUpdateOK describes a response with status code 200, with default header values.

Updated default policies object containing a policy for each status visibility.
*/
type PoliciesDefaultsUpdateOK struct {
	Payload *models.DefaultPolicies
}

// IsSuccess returns true when this policies defaults update o k response has a 2xx status code
func (o *PoliciesDefaultsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this policies defaults update o k response has a 3xx status code
func (o *PoliciesDefaultsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this policies defaults update o k response has a 4xx status code
func (o *PoliciesDefaultsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this policies defaults update o k response has a 5xx status code
func (o *PoliciesDefaultsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this policies defaults update o k response a status code equal to that given
func (o *PoliciesDefaultsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the policies defaults update o k response
func (o *PoliciesDefaultsUpdateOK) Code() int {
	return 200
}

func (o *PoliciesDefaultsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1/interaction_policies/defaults][%d] policiesDefaultsUpdateOK %s", 200, payload)
}

func (o *PoliciesDefaultsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1/interaction_policies/defaults][%d] policiesDefaultsUpdateOK %s", 200, payload)
}

func (o *PoliciesDefaultsUpdateOK) GetPayload() *models.DefaultPolicies {
	return o.Payload
}

func (o *PoliciesDefaultsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultPolicies)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPoliciesDefaultsUpdateBadRequest creates a PoliciesDefaultsUpdateBadRequest with default headers values
func NewPoliciesDefaultsUpdateBadRequest() *PoliciesDefaultsUpdateBadRequest {
	return &PoliciesDefaultsUpdateBadRequest{}
}

/*
PoliciesDefaultsUpdateBadRequest describes a response with status code 400, with default header values.

bad request
*/
type PoliciesDefaultsUpdateBadRequest struct {
}

// IsSuccess returns true when this policies defaults update bad request response has a 2xx status code
func (o *PoliciesDefaultsUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this policies defaults update bad request response has a 3xx status code
func (o *PoliciesDefaultsUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this policies defaults update bad request response has a 4xx status code
func (o *PoliciesDefaultsUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this policies defaults update bad request response has a 5xx status code
func (o *PoliciesDefaultsUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this policies defaults update bad request response a status code equal to that given
func (o *PoliciesDefaultsUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the policies defaults update bad request response
func (o *PoliciesDefaultsUpdateBadRequest) Code() int {
	return 400
}

func (o *PoliciesDefaultsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/interaction_policies/defaults][%d] policiesDefaultsUpdateBadRequest", 400)
}

func (o *PoliciesDefaultsUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v1/interaction_policies/defaults][%d] policiesDefaultsUpdateBadRequest", 400)
}

func (o *PoliciesDefaultsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPoliciesDefaultsUpdateUnauthorized creates a PoliciesDefaultsUpdateUnauthorized with default headers values
func NewPoliciesDefaultsUpdateUnauthorized() *PoliciesDefaultsUpdateUnauthorized {
	return &PoliciesDefaultsUpdateUnauthorized{}
}

/*
PoliciesDefaultsUpdateUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type PoliciesDefaultsUpdateUnauthorized struct {
}

// IsSuccess returns true when this policies defaults update unauthorized response has a 2xx status code
func (o *PoliciesDefaultsUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this policies defaults update unauthorized response has a 3xx status code
func (o *PoliciesDefaultsUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this policies defaults update unauthorized response has a 4xx status code
func (o *PoliciesDefaultsUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this policies defaults update unauthorized response has a 5xx status code
func (o *PoliciesDefaultsUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this policies defaults update unauthorized response a status code equal to that given
func (o *PoliciesDefaultsUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the policies defaults update unauthorized response
func (o *PoliciesDefaultsUpdateUnauthorized) Code() int {
	return 401
}

func (o *PoliciesDefaultsUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/interaction_policies/defaults][%d] policiesDefaultsUpdateUnauthorized", 401)
}

func (o *PoliciesDefaultsUpdateUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v1/interaction_policies/defaults][%d] policiesDefaultsUpdateUnauthorized", 401)
}

func (o *PoliciesDefaultsUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPoliciesDefaultsUpdateNotAcceptable creates a PoliciesDefaultsUpdateNotAcceptable with default headers values
func NewPoliciesDefaultsUpdateNotAcceptable() *PoliciesDefaultsUpdateNotAcceptable {
	return &PoliciesDefaultsUpdateNotAcceptable{}
}

/*
PoliciesDefaultsUpdateNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type PoliciesDefaultsUpdateNotAcceptable struct {
}

// IsSuccess returns true when this policies defaults update not acceptable response has a 2xx status code
func (o *PoliciesDefaultsUpdateNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this policies defaults update not acceptable response has a 3xx status code
func (o *PoliciesDefaultsUpdateNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this policies defaults update not acceptable response has a 4xx status code
func (o *PoliciesDefaultsUpdateNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this policies defaults update not acceptable response has a 5xx status code
func (o *PoliciesDefaultsUpdateNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this policies defaults update not acceptable response a status code equal to that given
func (o *PoliciesDefaultsUpdateNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the policies defaults update not acceptable response
func (o *PoliciesDefaultsUpdateNotAcceptable) Code() int {
	return 406
}

func (o *PoliciesDefaultsUpdateNotAcceptable) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/interaction_policies/defaults][%d] policiesDefaultsUpdateNotAcceptable", 406)
}

func (o *PoliciesDefaultsUpdateNotAcceptable) String() string {
	return fmt.Sprintf("[PATCH /api/v1/interaction_policies/defaults][%d] policiesDefaultsUpdateNotAcceptable", 406)
}

func (o *PoliciesDefaultsUpdateNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPoliciesDefaultsUpdateUnprocessableEntity creates a PoliciesDefaultsUpdateUnprocessableEntity with default headers values
func NewPoliciesDefaultsUpdateUnprocessableEntity() *PoliciesDefaultsUpdateUnprocessableEntity {
	return &PoliciesDefaultsUpdateUnprocessableEntity{}
}

/*
PoliciesDefaultsUpdateUnprocessableEntity describes a response with status code 422, with default header values.

unprocessable
*/
type PoliciesDefaultsUpdateUnprocessableEntity struct {
}

// IsSuccess returns true when this policies defaults update unprocessable entity response has a 2xx status code
func (o *PoliciesDefaultsUpdateUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this policies defaults update unprocessable entity response has a 3xx status code
func (o *PoliciesDefaultsUpdateUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this policies defaults update unprocessable entity response has a 4xx status code
func (o *PoliciesDefaultsUpdateUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this policies defaults update unprocessable entity response has a 5xx status code
func (o *PoliciesDefaultsUpdateUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this policies defaults update unprocessable entity response a status code equal to that given
func (o *PoliciesDefaultsUpdateUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the policies defaults update unprocessable entity response
func (o *PoliciesDefaultsUpdateUnprocessableEntity) Code() int {
	return 422
}

func (o *PoliciesDefaultsUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/interaction_policies/defaults][%d] policiesDefaultsUpdateUnprocessableEntity", 422)
}

func (o *PoliciesDefaultsUpdateUnprocessableEntity) String() string {
	return fmt.Sprintf("[PATCH /api/v1/interaction_policies/defaults][%d] policiesDefaultsUpdateUnprocessableEntity", 422)
}

func (o *PoliciesDefaultsUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPoliciesDefaultsUpdateInternalServerError creates a PoliciesDefaultsUpdateInternalServerError with default headers values
func NewPoliciesDefaultsUpdateInternalServerError() *PoliciesDefaultsUpdateInternalServerError {
	return &PoliciesDefaultsUpdateInternalServerError{}
}

/*
PoliciesDefaultsUpdateInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type PoliciesDefaultsUpdateInternalServerError struct {
}

// IsSuccess returns true when this policies defaults update internal server error response has a 2xx status code
func (o *PoliciesDefaultsUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this policies defaults update internal server error response has a 3xx status code
func (o *PoliciesDefaultsUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this policies defaults update internal server error response has a 4xx status code
func (o *PoliciesDefaultsUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this policies defaults update internal server error response has a 5xx status code
func (o *PoliciesDefaultsUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this policies defaults update internal server error response a status code equal to that given
func (o *PoliciesDefaultsUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the policies defaults update internal server error response
func (o *PoliciesDefaultsUpdateInternalServerError) Code() int {
	return 500
}

func (o *PoliciesDefaultsUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/interaction_policies/defaults][%d] policiesDefaultsUpdateInternalServerError", 500)
}

func (o *PoliciesDefaultsUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v1/interaction_policies/defaults][%d] policiesDefaultsUpdateInternalServerError", 500)
}

func (o *PoliciesDefaultsUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
