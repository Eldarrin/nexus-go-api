// Code generated by go-swagger; DO NOT EDIT.

package script

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AddReader is a Reader for the Add structure.
type AddReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAddNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 410:
		result := NewAddGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddNoContent creates a AddNoContent with default headers values
func NewAddNoContent() *AddNoContent {
	return &AddNoContent{}
}

/* AddNoContent describes a response with status code 204, with default header values.

Script was added
*/
type AddNoContent struct {
}

// IsSuccess returns true when this add no content response has a 2xx status code
func (o *AddNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add no content response has a 3xx status code
func (o *AddNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add no content response has a 4xx status code
func (o *AddNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this add no content response has a 5xx status code
func (o *AddNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this add no content response a status code equal to that given
func (o *AddNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *AddNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/script][%d] addNoContent ", 204)
}

func (o *AddNoContent) String() string {
	return fmt.Sprintf("[POST /v1/script][%d] addNoContent ", 204)
}

func (o *AddNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddGone creates a AddGone with default headers values
func NewAddGone() *AddGone {
	return &AddGone{}
}

/* AddGone describes a response with status code 410, with default header values.

Script creation is disabled
*/
type AddGone struct {
}

// IsSuccess returns true when this add gone response has a 2xx status code
func (o *AddGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add gone response has a 3xx status code
func (o *AddGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add gone response has a 4xx status code
func (o *AddGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this add gone response has a 5xx status code
func (o *AddGone) IsServerError() bool {
	return false
}

// IsCode returns true when this add gone response a status code equal to that given
func (o *AddGone) IsCode(code int) bool {
	return code == 410
}

func (o *AddGone) Error() string {
	return fmt.Sprintf("[POST /v1/script][%d] addGone ", 410)
}

func (o *AddGone) String() string {
	return fmt.Sprintf("[POST /v1/script][%d] addGone ", 410)
}

func (o *AddGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}