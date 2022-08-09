// Code generated by go-swagger; DO NOT EDIT.

package read_only

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ForceReleaseReader is a Reader for the ForceRelease structure.
type ForceReleaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ForceReleaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewForceReleaseNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewForceReleaseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewForceReleaseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewForceReleaseNoContent creates a ForceReleaseNoContent with default headers values
func NewForceReleaseNoContent() *ForceReleaseNoContent {
	return &ForceReleaseNoContent{}
}

/*
	ForceReleaseNoContent describes a response with status code 204, with default header values.

System is no longer read-only
*/
type ForceReleaseNoContent struct {
}

// IsSuccess returns true when this force release no content response has a 2xx status code
func (o *ForceReleaseNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this force release no content response has a 3xx status code
func (o *ForceReleaseNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this force release no content response has a 4xx status code
func (o *ForceReleaseNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this force release no content response has a 5xx status code
func (o *ForceReleaseNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this force release no content response a status code equal to that given
func (o *ForceReleaseNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ForceReleaseNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/read-only/force-release][%d] forceReleaseNoContent ", 204)
}

func (o *ForceReleaseNoContent) String() string {
	return fmt.Sprintf("[POST /v1/read-only/force-release][%d] forceReleaseNoContent ", 204)
}

func (o *ForceReleaseNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewForceReleaseForbidden creates a ForceReleaseForbidden with default headers values
func NewForceReleaseForbidden() *ForceReleaseForbidden {
	return &ForceReleaseForbidden{}
}

/*
	ForceReleaseForbidden describes a response with status code 403, with default header values.

Authentication required
*/
type ForceReleaseForbidden struct {
}

// IsSuccess returns true when this force release forbidden response has a 2xx status code
func (o *ForceReleaseForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this force release forbidden response has a 3xx status code
func (o *ForceReleaseForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this force release forbidden response has a 4xx status code
func (o *ForceReleaseForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this force release forbidden response has a 5xx status code
func (o *ForceReleaseForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this force release forbidden response a status code equal to that given
func (o *ForceReleaseForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ForceReleaseForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/read-only/force-release][%d] forceReleaseForbidden ", 403)
}

func (o *ForceReleaseForbidden) String() string {
	return fmt.Sprintf("[POST /v1/read-only/force-release][%d] forceReleaseForbidden ", 403)
}

func (o *ForceReleaseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewForceReleaseNotFound creates a ForceReleaseNotFound with default headers values
func NewForceReleaseNotFound() *ForceReleaseNotFound {
	return &ForceReleaseNotFound{}
}

/*
	ForceReleaseNotFound describes a response with status code 404, with default header values.

No change to read-only state
*/
type ForceReleaseNotFound struct {
}

// IsSuccess returns true when this force release not found response has a 2xx status code
func (o *ForceReleaseNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this force release not found response has a 3xx status code
func (o *ForceReleaseNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this force release not found response has a 4xx status code
func (o *ForceReleaseNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this force release not found response has a 5xx status code
func (o *ForceReleaseNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this force release not found response a status code equal to that given
func (o *ForceReleaseNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ForceReleaseNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/read-only/force-release][%d] forceReleaseNotFound ", 404)
}

func (o *ForceReleaseNotFound) String() string {
	return fmt.Sprintf("[POST /v1/read-only/force-release][%d] forceReleaseNotFound ", 404)
}

func (o *ForceReleaseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
