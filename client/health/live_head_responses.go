// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// LiveHeadReader is a Reader for the LiveHead structure.
type LiveHeadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LiveHeadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLiveHeadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[HEAD /livez] liveHead", response, response.Code())
	}
}

// NewLiveHeadOK creates a LiveHeadOK with default headers values
func NewLiveHeadOK() *LiveHeadOK {
	return &LiveHeadOK{}
}

/*
LiveHeadOK describes a response with status code 200, with default header values.

OK
*/
type LiveHeadOK struct {
}

// IsSuccess returns true when this live head o k response has a 2xx status code
func (o *LiveHeadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this live head o k response has a 3xx status code
func (o *LiveHeadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this live head o k response has a 4xx status code
func (o *LiveHeadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this live head o k response has a 5xx status code
func (o *LiveHeadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this live head o k response a status code equal to that given
func (o *LiveHeadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the live head o k response
func (o *LiveHeadOK) Code() int {
	return 200
}

func (o *LiveHeadOK) Error() string {
	return fmt.Sprintf("[HEAD /livez][%d] liveHeadOK ", 200)
}

func (o *LiveHeadOK) String() string {
	return fmt.Sprintf("[HEAD /livez][%d] liveHeadOK ", 200)
}

func (o *LiveHeadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
