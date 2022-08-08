// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository16Reader is a Reader for the UpdateRepository16 structure.
type UpdateRepository16Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository16Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository16NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRepository16Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository16Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepository16NoContent creates a UpdateRepository16NoContent with default headers values
func NewUpdateRepository16NoContent() *UpdateRepository16NoContent {
	return &UpdateRepository16NoContent{}
}

/* UpdateRepository16NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository16NoContent struct {
}

// IsSuccess returns true when this update repository16 no content response has a 2xx status code
func (o *UpdateRepository16NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository16 no content response has a 3xx status code
func (o *UpdateRepository16NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository16 no content response has a 4xx status code
func (o *UpdateRepository16NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository16 no content response has a 5xx status code
func (o *UpdateRepository16NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository16 no content response a status code equal to that given
func (o *UpdateRepository16NoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateRepository16NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/rubygems/proxy/{repositoryName}][%d] updateRepository16NoContent ", 204)
}

func (o *UpdateRepository16NoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/rubygems/proxy/{repositoryName}][%d] updateRepository16NoContent ", 204)
}

func (o *UpdateRepository16NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository16Unauthorized creates a UpdateRepository16Unauthorized with default headers values
func NewUpdateRepository16Unauthorized() *UpdateRepository16Unauthorized {
	return &UpdateRepository16Unauthorized{}
}

/* UpdateRepository16Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository16Unauthorized struct {
}

// IsSuccess returns true when this update repository16 unauthorized response has a 2xx status code
func (o *UpdateRepository16Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository16 unauthorized response has a 3xx status code
func (o *UpdateRepository16Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository16 unauthorized response has a 4xx status code
func (o *UpdateRepository16Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository16 unauthorized response has a 5xx status code
func (o *UpdateRepository16Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository16 unauthorized response a status code equal to that given
func (o *UpdateRepository16Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateRepository16Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/rubygems/proxy/{repositoryName}][%d] updateRepository16Unauthorized ", 401)
}

func (o *UpdateRepository16Unauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/rubygems/proxy/{repositoryName}][%d] updateRepository16Unauthorized ", 401)
}

func (o *UpdateRepository16Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository16Forbidden creates a UpdateRepository16Forbidden with default headers values
func NewUpdateRepository16Forbidden() *UpdateRepository16Forbidden {
	return &UpdateRepository16Forbidden{}
}

/* UpdateRepository16Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository16Forbidden struct {
}

// IsSuccess returns true when this update repository16 forbidden response has a 2xx status code
func (o *UpdateRepository16Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository16 forbidden response has a 3xx status code
func (o *UpdateRepository16Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository16 forbidden response has a 4xx status code
func (o *UpdateRepository16Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository16 forbidden response has a 5xx status code
func (o *UpdateRepository16Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository16 forbidden response a status code equal to that given
func (o *UpdateRepository16Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateRepository16Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/rubygems/proxy/{repositoryName}][%d] updateRepository16Forbidden ", 403)
}

func (o *UpdateRepository16Forbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/rubygems/proxy/{repositoryName}][%d] updateRepository16Forbidden ", 403)
}

func (o *UpdateRepository16Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
