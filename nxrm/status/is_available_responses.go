// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IsAvailableReader is a Reader for the IsAvailable structure.
type IsAvailableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IsAvailableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIsAvailableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 503:
		result := NewIsAvailableServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIsAvailableOK creates a IsAvailableOK with default headers values
func NewIsAvailableOK() *IsAvailableOK {
	return &IsAvailableOK{}
}

/* IsAvailableOK describes a response with status code 200, with default header values.

Available to service requests
*/
type IsAvailableOK struct {
}

// IsSuccess returns true when this is available o k response has a 2xx status code
func (o *IsAvailableOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this is available o k response has a 3xx status code
func (o *IsAvailableOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this is available o k response has a 4xx status code
func (o *IsAvailableOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this is available o k response has a 5xx status code
func (o *IsAvailableOK) IsServerError() bool {
	return false
}

// IsCode returns true when this is available o k response a status code equal to that given
func (o *IsAvailableOK) IsCode(code int) bool {
	return code == 200
}

func (o *IsAvailableOK) Error() string {
	return fmt.Sprintf("[GET /v1/status][%d] isAvailableOK ", 200)
}

func (o *IsAvailableOK) String() string {
	return fmt.Sprintf("[GET /v1/status][%d] isAvailableOK ", 200)
}

func (o *IsAvailableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIsAvailableServiceUnavailable creates a IsAvailableServiceUnavailable with default headers values
func NewIsAvailableServiceUnavailable() *IsAvailableServiceUnavailable {
	return &IsAvailableServiceUnavailable{}
}

/* IsAvailableServiceUnavailable describes a response with status code 503, with default header values.

Unavailable to service requests
*/
type IsAvailableServiceUnavailable struct {
}

// IsSuccess returns true when this is available service unavailable response has a 2xx status code
func (o *IsAvailableServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this is available service unavailable response has a 3xx status code
func (o *IsAvailableServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this is available service unavailable response has a 4xx status code
func (o *IsAvailableServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this is available service unavailable response has a 5xx status code
func (o *IsAvailableServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this is available service unavailable response a status code equal to that given
func (o *IsAvailableServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *IsAvailableServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/status][%d] isAvailableServiceUnavailable ", 503)
}

func (o *IsAvailableServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/status][%d] isAvailableServiceUnavailable ", 503)
}

func (o *IsAvailableServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
