// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository6Reader is a Reader for the UpdateRepository6 structure.
type UpdateRepository6Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository6Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository6NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRepository6Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository6Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepository6NoContent creates a UpdateRepository6NoContent with default headers values
func NewUpdateRepository6NoContent() *UpdateRepository6NoContent {
	return &UpdateRepository6NoContent{}
}

/* UpdateRepository6NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository6NoContent struct {
}

// IsSuccess returns true when this update repository6 no content response has a 2xx status code
func (o *UpdateRepository6NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository6 no content response has a 3xx status code
func (o *UpdateRepository6NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository6 no content response has a 4xx status code
func (o *UpdateRepository6NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository6 no content response has a 5xx status code
func (o *UpdateRepository6NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository6 no content response a status code equal to that given
func (o *UpdateRepository6NoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateRepository6NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/raw/hosted/{repositoryName}][%d] updateRepository6NoContent ", 204)
}

func (o *UpdateRepository6NoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/raw/hosted/{repositoryName}][%d] updateRepository6NoContent ", 204)
}

func (o *UpdateRepository6NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository6Unauthorized creates a UpdateRepository6Unauthorized with default headers values
func NewUpdateRepository6Unauthorized() *UpdateRepository6Unauthorized {
	return &UpdateRepository6Unauthorized{}
}

/* UpdateRepository6Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository6Unauthorized struct {
}

// IsSuccess returns true when this update repository6 unauthorized response has a 2xx status code
func (o *UpdateRepository6Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository6 unauthorized response has a 3xx status code
func (o *UpdateRepository6Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository6 unauthorized response has a 4xx status code
func (o *UpdateRepository6Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository6 unauthorized response has a 5xx status code
func (o *UpdateRepository6Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository6 unauthorized response a status code equal to that given
func (o *UpdateRepository6Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateRepository6Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/raw/hosted/{repositoryName}][%d] updateRepository6Unauthorized ", 401)
}

func (o *UpdateRepository6Unauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/raw/hosted/{repositoryName}][%d] updateRepository6Unauthorized ", 401)
}

func (o *UpdateRepository6Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository6Forbidden creates a UpdateRepository6Forbidden with default headers values
func NewUpdateRepository6Forbidden() *UpdateRepository6Forbidden {
	return &UpdateRepository6Forbidden{}
}

/* UpdateRepository6Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository6Forbidden struct {
}

// IsSuccess returns true when this update repository6 forbidden response has a 2xx status code
func (o *UpdateRepository6Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository6 forbidden response has a 3xx status code
func (o *UpdateRepository6Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository6 forbidden response has a 4xx status code
func (o *UpdateRepository6Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository6 forbidden response has a 5xx status code
func (o *UpdateRepository6Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository6 forbidden response a status code equal to that given
func (o *UpdateRepository6Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateRepository6Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/raw/hosted/{repositoryName}][%d] updateRepository6Forbidden ", 403)
}

func (o *UpdateRepository6Forbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/raw/hosted/{repositoryName}][%d] updateRepository6Forbidden ", 403)
}

func (o *UpdateRepository6Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}