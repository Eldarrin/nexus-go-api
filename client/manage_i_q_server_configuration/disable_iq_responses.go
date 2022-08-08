// Code generated by go-swagger; DO NOT EDIT.

package manage_i_q_server_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DisableIqReader is a Reader for the DisableIq structure.
type DisableIqReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisableIqReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDisableIqNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDisableIqBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDisableIqNoContent creates a DisableIqNoContent with default headers values
func NewDisableIqNoContent() *DisableIqNoContent {
	return &DisableIqNoContent{}
}

/* DisableIqNoContent describes a response with status code 204, with default header values.

IQ server has been disabled
*/
type DisableIqNoContent struct {
}

// IsSuccess returns true when this disable iq no content response has a 2xx status code
func (o *DisableIqNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this disable iq no content response has a 3xx status code
func (o *DisableIqNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable iq no content response has a 4xx status code
func (o *DisableIqNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this disable iq no content response has a 5xx status code
func (o *DisableIqNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this disable iq no content response a status code equal to that given
func (o *DisableIqNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DisableIqNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/iq/disable][%d] disableIqNoContent ", 204)
}

func (o *DisableIqNoContent) String() string {
	return fmt.Sprintf("[POST /v1/iq/disable][%d] disableIqNoContent ", 204)
}

func (o *DisableIqNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisableIqBadRequest creates a DisableIqBadRequest with default headers values
func NewDisableIqBadRequest() *DisableIqBadRequest {
	return &DisableIqBadRequest{}
}

/* DisableIqBadRequest describes a response with status code 400, with default header values.

IQ server connection not configured
*/
type DisableIqBadRequest struct {
}

// IsSuccess returns true when this disable iq bad request response has a 2xx status code
func (o *DisableIqBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable iq bad request response has a 3xx status code
func (o *DisableIqBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable iq bad request response has a 4xx status code
func (o *DisableIqBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this disable iq bad request response has a 5xx status code
func (o *DisableIqBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this disable iq bad request response a status code equal to that given
func (o *DisableIqBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DisableIqBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/iq/disable][%d] disableIqBadRequest ", 400)
}

func (o *DisableIqBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/iq/disable][%d] disableIqBadRequest ", 400)
}

func (o *DisableIqBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
