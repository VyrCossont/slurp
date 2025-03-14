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

// RuleUpdateReader is a Reader for the RuleUpdate structure.
type RuleUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RuleUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRuleUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRuleUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRuleUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRuleUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRuleUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewRuleUpdateNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRuleUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /api/v1/admin/instance/rules/{id}] ruleUpdate", response, response.Code())
	}
}

// NewRuleUpdateOK creates a RuleUpdateOK with default headers values
func NewRuleUpdateOK() *RuleUpdateOK {
	return &RuleUpdateOK{}
}

/*
RuleUpdateOK describes a response with status code 200, with default header values.

The updated instance rule.
*/
type RuleUpdateOK struct {
	Payload *models.InstanceRule
}

// IsSuccess returns true when this rule update o k response has a 2xx status code
func (o *RuleUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rule update o k response has a 3xx status code
func (o *RuleUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rule update o k response has a 4xx status code
func (o *RuleUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this rule update o k response has a 5xx status code
func (o *RuleUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this rule update o k response a status code equal to that given
func (o *RuleUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the rule update o k response
func (o *RuleUpdateOK) Code() int {
	return 200
}

func (o *RuleUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1/admin/instance/rules/{id}][%d] ruleUpdateOK %s", 200, payload)
}

func (o *RuleUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1/admin/instance/rules/{id}][%d] ruleUpdateOK %s", 200, payload)
}

func (o *RuleUpdateOK) GetPayload() *models.InstanceRule {
	return o.Payload
}

func (o *RuleUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InstanceRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRuleUpdateBadRequest creates a RuleUpdateBadRequest with default headers values
func NewRuleUpdateBadRequest() *RuleUpdateBadRequest {
	return &RuleUpdateBadRequest{}
}

/*
RuleUpdateBadRequest describes a response with status code 400, with default header values.

bad request
*/
type RuleUpdateBadRequest struct {
}

// IsSuccess returns true when this rule update bad request response has a 2xx status code
func (o *RuleUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rule update bad request response has a 3xx status code
func (o *RuleUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rule update bad request response has a 4xx status code
func (o *RuleUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this rule update bad request response has a 5xx status code
func (o *RuleUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this rule update bad request response a status code equal to that given
func (o *RuleUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the rule update bad request response
func (o *RuleUpdateBadRequest) Code() int {
	return 400
}

func (o *RuleUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/instance/rules/{id}][%d] ruleUpdateBadRequest", 400)
}

func (o *RuleUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/instance/rules/{id}][%d] ruleUpdateBadRequest", 400)
}

func (o *RuleUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRuleUpdateUnauthorized creates a RuleUpdateUnauthorized with default headers values
func NewRuleUpdateUnauthorized() *RuleUpdateUnauthorized {
	return &RuleUpdateUnauthorized{}
}

/*
RuleUpdateUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type RuleUpdateUnauthorized struct {
}

// IsSuccess returns true when this rule update unauthorized response has a 2xx status code
func (o *RuleUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rule update unauthorized response has a 3xx status code
func (o *RuleUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rule update unauthorized response has a 4xx status code
func (o *RuleUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this rule update unauthorized response has a 5xx status code
func (o *RuleUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this rule update unauthorized response a status code equal to that given
func (o *RuleUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the rule update unauthorized response
func (o *RuleUpdateUnauthorized) Code() int {
	return 401
}

func (o *RuleUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/instance/rules/{id}][%d] ruleUpdateUnauthorized", 401)
}

func (o *RuleUpdateUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/instance/rules/{id}][%d] ruleUpdateUnauthorized", 401)
}

func (o *RuleUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRuleUpdateForbidden creates a RuleUpdateForbidden with default headers values
func NewRuleUpdateForbidden() *RuleUpdateForbidden {
	return &RuleUpdateForbidden{}
}

/*
RuleUpdateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type RuleUpdateForbidden struct {
}

// IsSuccess returns true when this rule update forbidden response has a 2xx status code
func (o *RuleUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rule update forbidden response has a 3xx status code
func (o *RuleUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rule update forbidden response has a 4xx status code
func (o *RuleUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this rule update forbidden response has a 5xx status code
func (o *RuleUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this rule update forbidden response a status code equal to that given
func (o *RuleUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the rule update forbidden response
func (o *RuleUpdateForbidden) Code() int {
	return 403
}

func (o *RuleUpdateForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/instance/rules/{id}][%d] ruleUpdateForbidden", 403)
}

func (o *RuleUpdateForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/instance/rules/{id}][%d] ruleUpdateForbidden", 403)
}

func (o *RuleUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRuleUpdateNotFound creates a RuleUpdateNotFound with default headers values
func NewRuleUpdateNotFound() *RuleUpdateNotFound {
	return &RuleUpdateNotFound{}
}

/*
RuleUpdateNotFound describes a response with status code 404, with default header values.

not found
*/
type RuleUpdateNotFound struct {
}

// IsSuccess returns true when this rule update not found response has a 2xx status code
func (o *RuleUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rule update not found response has a 3xx status code
func (o *RuleUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rule update not found response has a 4xx status code
func (o *RuleUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this rule update not found response has a 5xx status code
func (o *RuleUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this rule update not found response a status code equal to that given
func (o *RuleUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the rule update not found response
func (o *RuleUpdateNotFound) Code() int {
	return 404
}

func (o *RuleUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/instance/rules/{id}][%d] ruleUpdateNotFound", 404)
}

func (o *RuleUpdateNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/instance/rules/{id}][%d] ruleUpdateNotFound", 404)
}

func (o *RuleUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRuleUpdateNotAcceptable creates a RuleUpdateNotAcceptable with default headers values
func NewRuleUpdateNotAcceptable() *RuleUpdateNotAcceptable {
	return &RuleUpdateNotAcceptable{}
}

/*
RuleUpdateNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type RuleUpdateNotAcceptable struct {
}

// IsSuccess returns true when this rule update not acceptable response has a 2xx status code
func (o *RuleUpdateNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rule update not acceptable response has a 3xx status code
func (o *RuleUpdateNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rule update not acceptable response has a 4xx status code
func (o *RuleUpdateNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this rule update not acceptable response has a 5xx status code
func (o *RuleUpdateNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this rule update not acceptable response a status code equal to that given
func (o *RuleUpdateNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the rule update not acceptable response
func (o *RuleUpdateNotAcceptable) Code() int {
	return 406
}

func (o *RuleUpdateNotAcceptable) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/instance/rules/{id}][%d] ruleUpdateNotAcceptable", 406)
}

func (o *RuleUpdateNotAcceptable) String() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/instance/rules/{id}][%d] ruleUpdateNotAcceptable", 406)
}

func (o *RuleUpdateNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRuleUpdateInternalServerError creates a RuleUpdateInternalServerError with default headers values
func NewRuleUpdateInternalServerError() *RuleUpdateInternalServerError {
	return &RuleUpdateInternalServerError{}
}

/*
RuleUpdateInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type RuleUpdateInternalServerError struct {
}

// IsSuccess returns true when this rule update internal server error response has a 2xx status code
func (o *RuleUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rule update internal server error response has a 3xx status code
func (o *RuleUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rule update internal server error response has a 4xx status code
func (o *RuleUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this rule update internal server error response has a 5xx status code
func (o *RuleUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this rule update internal server error response a status code equal to that given
func (o *RuleUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the rule update internal server error response
func (o *RuleUpdateInternalServerError) Code() int {
	return 500
}

func (o *RuleUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/instance/rules/{id}][%d] ruleUpdateInternalServerError", 500)
}

func (o *RuleUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/instance/rules/{id}][%d] ruleUpdateInternalServerError", 500)
}

func (o *RuleUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
