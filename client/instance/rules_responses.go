// Code generated by go-swagger; DO NOT EDIT.

package instance

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

// RulesReader is a Reader for the Rules structure.
type RulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewRulesNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/instance/rules] rules", response, response.Code())
	}
}

// NewRulesOK creates a RulesOK with default headers values
func NewRulesOK() *RulesOK {
	return &RulesOK{}
}

/*
RulesOK describes a response with status code 200, with default header values.

An array with all the rules for the local instance.
*/
type RulesOK struct {
	Payload []*models.InstanceRule
}

// IsSuccess returns true when this rules o k response has a 2xx status code
func (o *RulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rules o k response has a 3xx status code
func (o *RulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rules o k response has a 4xx status code
func (o *RulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this rules o k response has a 5xx status code
func (o *RulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this rules o k response a status code equal to that given
func (o *RulesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the rules o k response
func (o *RulesOK) Code() int {
	return 200
}

func (o *RulesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/instance/rules][%d] rulesOK %s", 200, payload)
}

func (o *RulesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/instance/rules][%d] rulesOK %s", 200, payload)
}

func (o *RulesOK) GetPayload() []*models.InstanceRule {
	return o.Payload
}

func (o *RulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRulesBadRequest creates a RulesBadRequest with default headers values
func NewRulesBadRequest() *RulesBadRequest {
	return &RulesBadRequest{}
}

/*
RulesBadRequest describes a response with status code 400, with default header values.

bad request
*/
type RulesBadRequest struct {
}

// IsSuccess returns true when this rules bad request response has a 2xx status code
func (o *RulesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rules bad request response has a 3xx status code
func (o *RulesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rules bad request response has a 4xx status code
func (o *RulesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this rules bad request response has a 5xx status code
func (o *RulesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this rules bad request response a status code equal to that given
func (o *RulesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the rules bad request response
func (o *RulesBadRequest) Code() int {
	return 400
}

func (o *RulesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/instance/rules][%d] rulesBadRequest", 400)
}

func (o *RulesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/instance/rules][%d] rulesBadRequest", 400)
}

func (o *RulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRulesNotFound creates a RulesNotFound with default headers values
func NewRulesNotFound() *RulesNotFound {
	return &RulesNotFound{}
}

/*
RulesNotFound describes a response with status code 404, with default header values.

not found
*/
type RulesNotFound struct {
}

// IsSuccess returns true when this rules not found response has a 2xx status code
func (o *RulesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rules not found response has a 3xx status code
func (o *RulesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rules not found response has a 4xx status code
func (o *RulesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this rules not found response has a 5xx status code
func (o *RulesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this rules not found response a status code equal to that given
func (o *RulesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the rules not found response
func (o *RulesNotFound) Code() int {
	return 404
}

func (o *RulesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/instance/rules][%d] rulesNotFound", 404)
}

func (o *RulesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/instance/rules][%d] rulesNotFound", 404)
}

func (o *RulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRulesNotAcceptable creates a RulesNotAcceptable with default headers values
func NewRulesNotAcceptable() *RulesNotAcceptable {
	return &RulesNotAcceptable{}
}

/*
RulesNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type RulesNotAcceptable struct {
}

// IsSuccess returns true when this rules not acceptable response has a 2xx status code
func (o *RulesNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rules not acceptable response has a 3xx status code
func (o *RulesNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rules not acceptable response has a 4xx status code
func (o *RulesNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this rules not acceptable response has a 5xx status code
func (o *RulesNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this rules not acceptable response a status code equal to that given
func (o *RulesNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the rules not acceptable response
func (o *RulesNotAcceptable) Code() int {
	return 406
}

func (o *RulesNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/instance/rules][%d] rulesNotAcceptable", 406)
}

func (o *RulesNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/instance/rules][%d] rulesNotAcceptable", 406)
}

func (o *RulesNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRulesInternalServerError creates a RulesInternalServerError with default headers values
func NewRulesInternalServerError() *RulesInternalServerError {
	return &RulesInternalServerError{}
}

/*
RulesInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type RulesInternalServerError struct {
}

// IsSuccess returns true when this rules internal server error response has a 2xx status code
func (o *RulesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rules internal server error response has a 3xx status code
func (o *RulesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rules internal server error response has a 4xx status code
func (o *RulesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this rules internal server error response has a 5xx status code
func (o *RulesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this rules internal server error response a status code equal to that given
func (o *RulesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the rules internal server error response
func (o *RulesInternalServerError) Code() int {
	return 500
}

func (o *RulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/instance/rules][%d] rulesInternalServerError", 500)
}

func (o *RulesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/instance/rules][%d] rulesInternalServerError", 500)
}

func (o *RulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
