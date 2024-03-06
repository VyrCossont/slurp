// Code generated by go-swagger; DO NOT EDIT.

package blocks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VyrCossont/slurp/models"
)

// BlocksGetReader is a Reader for the BlocksGet structure.
type BlocksGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BlocksGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBlocksGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBlocksGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBlocksGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBlocksGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewBlocksGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBlocksGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/blocks] blocksGet", response, response.Code())
	}
}

// NewBlocksGetOK creates a BlocksGetOK with default headers values
func NewBlocksGetOK() *BlocksGetOK {
	return &BlocksGetOK{}
}

/*
BlocksGetOK describes a response with status code 200, with default header values.

BlocksGetOK blocks get o k
*/
type BlocksGetOK struct {

	/* Links to the next and previous queries.
	 */
	Link string

	Payload []*models.Account
}

// IsSuccess returns true when this blocks get o k response has a 2xx status code
func (o *BlocksGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this blocks get o k response has a 3xx status code
func (o *BlocksGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this blocks get o k response has a 4xx status code
func (o *BlocksGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this blocks get o k response has a 5xx status code
func (o *BlocksGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this blocks get o k response a status code equal to that given
func (o *BlocksGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the blocks get o k response
func (o *BlocksGetOK) Code() int {
	return 200
}

func (o *BlocksGetOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/blocks][%d] blocksGetOK  %+v", 200, o.Payload)
}

func (o *BlocksGetOK) String() string {
	return fmt.Sprintf("[GET /api/v1/blocks][%d] blocksGetOK  %+v", 200, o.Payload)
}

func (o *BlocksGetOK) GetPayload() []*models.Account {
	return o.Payload
}

func (o *BlocksGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewBlocksGetBadRequest creates a BlocksGetBadRequest with default headers values
func NewBlocksGetBadRequest() *BlocksGetBadRequest {
	return &BlocksGetBadRequest{}
}

/*
BlocksGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type BlocksGetBadRequest struct {
}

// IsSuccess returns true when this blocks get bad request response has a 2xx status code
func (o *BlocksGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this blocks get bad request response has a 3xx status code
func (o *BlocksGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this blocks get bad request response has a 4xx status code
func (o *BlocksGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this blocks get bad request response has a 5xx status code
func (o *BlocksGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this blocks get bad request response a status code equal to that given
func (o *BlocksGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the blocks get bad request response
func (o *BlocksGetBadRequest) Code() int {
	return 400
}

func (o *BlocksGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/blocks][%d] blocksGetBadRequest ", 400)
}

func (o *BlocksGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/blocks][%d] blocksGetBadRequest ", 400)
}

func (o *BlocksGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBlocksGetUnauthorized creates a BlocksGetUnauthorized with default headers values
func NewBlocksGetUnauthorized() *BlocksGetUnauthorized {
	return &BlocksGetUnauthorized{}
}

/*
BlocksGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type BlocksGetUnauthorized struct {
}

// IsSuccess returns true when this blocks get unauthorized response has a 2xx status code
func (o *BlocksGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this blocks get unauthorized response has a 3xx status code
func (o *BlocksGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this blocks get unauthorized response has a 4xx status code
func (o *BlocksGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this blocks get unauthorized response has a 5xx status code
func (o *BlocksGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this blocks get unauthorized response a status code equal to that given
func (o *BlocksGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the blocks get unauthorized response
func (o *BlocksGetUnauthorized) Code() int {
	return 401
}

func (o *BlocksGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/blocks][%d] blocksGetUnauthorized ", 401)
}

func (o *BlocksGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/blocks][%d] blocksGetUnauthorized ", 401)
}

func (o *BlocksGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBlocksGetNotFound creates a BlocksGetNotFound with default headers values
func NewBlocksGetNotFound() *BlocksGetNotFound {
	return &BlocksGetNotFound{}
}

/*
BlocksGetNotFound describes a response with status code 404, with default header values.

not found
*/
type BlocksGetNotFound struct {
}

// IsSuccess returns true when this blocks get not found response has a 2xx status code
func (o *BlocksGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this blocks get not found response has a 3xx status code
func (o *BlocksGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this blocks get not found response has a 4xx status code
func (o *BlocksGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this blocks get not found response has a 5xx status code
func (o *BlocksGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this blocks get not found response a status code equal to that given
func (o *BlocksGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the blocks get not found response
func (o *BlocksGetNotFound) Code() int {
	return 404
}

func (o *BlocksGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/blocks][%d] blocksGetNotFound ", 404)
}

func (o *BlocksGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/blocks][%d] blocksGetNotFound ", 404)
}

func (o *BlocksGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBlocksGetNotAcceptable creates a BlocksGetNotAcceptable with default headers values
func NewBlocksGetNotAcceptable() *BlocksGetNotAcceptable {
	return &BlocksGetNotAcceptable{}
}

/*
BlocksGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type BlocksGetNotAcceptable struct {
}

// IsSuccess returns true when this blocks get not acceptable response has a 2xx status code
func (o *BlocksGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this blocks get not acceptable response has a 3xx status code
func (o *BlocksGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this blocks get not acceptable response has a 4xx status code
func (o *BlocksGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this blocks get not acceptable response has a 5xx status code
func (o *BlocksGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this blocks get not acceptable response a status code equal to that given
func (o *BlocksGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the blocks get not acceptable response
func (o *BlocksGetNotAcceptable) Code() int {
	return 406
}

func (o *BlocksGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/blocks][%d] blocksGetNotAcceptable ", 406)
}

func (o *BlocksGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/blocks][%d] blocksGetNotAcceptable ", 406)
}

func (o *BlocksGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBlocksGetInternalServerError creates a BlocksGetInternalServerError with default headers values
func NewBlocksGetInternalServerError() *BlocksGetInternalServerError {
	return &BlocksGetInternalServerError{}
}

/*
BlocksGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type BlocksGetInternalServerError struct {
}

// IsSuccess returns true when this blocks get internal server error response has a 2xx status code
func (o *BlocksGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this blocks get internal server error response has a 3xx status code
func (o *BlocksGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this blocks get internal server error response has a 4xx status code
func (o *BlocksGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this blocks get internal server error response has a 5xx status code
func (o *BlocksGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this blocks get internal server error response a status code equal to that given
func (o *BlocksGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the blocks get internal server error response
func (o *BlocksGetInternalServerError) Code() int {
	return 500
}

func (o *BlocksGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/blocks][%d] blocksGetInternalServerError ", 500)
}

func (o *BlocksGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/blocks][%d] blocksGetInternalServerError ", 500)
}

func (o *BlocksGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
