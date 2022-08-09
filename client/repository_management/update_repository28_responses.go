// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository28Reader is a Reader for the UpdateRepository28 structure.
type UpdateRepository28Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository28Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository28NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRepository28Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository28Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepository28NoContent creates a UpdateRepository28NoContent with default headers values
func NewUpdateRepository28NoContent() *UpdateRepository28NoContent {
	return &UpdateRepository28NoContent{}
}

/* UpdateRepository28NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository28NoContent struct {
}

// IsSuccess returns true when this update repository28 no content response has a 2xx status code
func (o *UpdateRepository28NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository28 no content response has a 3xx status code
func (o *UpdateRepository28NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository28 no content response has a 4xx status code
func (o *UpdateRepository28NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository28 no content response has a 5xx status code
func (o *UpdateRepository28NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository28 no content response a status code equal to that given
func (o *UpdateRepository28NoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateRepository28NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/proxy/{repositoryName}][%d] updateRepository28NoContent ", 204)
}

func (o *UpdateRepository28NoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/proxy/{repositoryName}][%d] updateRepository28NoContent ", 204)
}

func (o *UpdateRepository28NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository28Unauthorized creates a UpdateRepository28Unauthorized with default headers values
func NewUpdateRepository28Unauthorized() *UpdateRepository28Unauthorized {
	return &UpdateRepository28Unauthorized{}
}

/* UpdateRepository28Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository28Unauthorized struct {
}

// IsSuccess returns true when this update repository28 unauthorized response has a 2xx status code
func (o *UpdateRepository28Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository28 unauthorized response has a 3xx status code
func (o *UpdateRepository28Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository28 unauthorized response has a 4xx status code
func (o *UpdateRepository28Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository28 unauthorized response has a 5xx status code
func (o *UpdateRepository28Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository28 unauthorized response a status code equal to that given
func (o *UpdateRepository28Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateRepository28Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/proxy/{repositoryName}][%d] updateRepository28Unauthorized ", 401)
}

func (o *UpdateRepository28Unauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/proxy/{repositoryName}][%d] updateRepository28Unauthorized ", 401)
}

func (o *UpdateRepository28Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository28Forbidden creates a UpdateRepository28Forbidden with default headers values
func NewUpdateRepository28Forbidden() *UpdateRepository28Forbidden {
	return &UpdateRepository28Forbidden{}
}

/* UpdateRepository28Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository28Forbidden struct {
}

// IsSuccess returns true when this update repository28 forbidden response has a 2xx status code
func (o *UpdateRepository28Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository28 forbidden response has a 3xx status code
func (o *UpdateRepository28Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository28 forbidden response has a 4xx status code
func (o *UpdateRepository28Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository28 forbidden response has a 5xx status code
func (o *UpdateRepository28Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository28 forbidden response a status code equal to that given
func (o *UpdateRepository28Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateRepository28Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/proxy/{repositoryName}][%d] updateRepository28Forbidden ", 403)
}

func (o *UpdateRepository28Forbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/proxy/{repositoryName}][%d] updateRepository28Forbidden ", 403)
}

func (o *UpdateRepository28Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}