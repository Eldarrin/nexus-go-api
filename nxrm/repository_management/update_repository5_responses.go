// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository5Reader is a Reader for the UpdateRepository5 structure.
type UpdateRepository5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository5NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRepository5Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository5Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepository5NoContent creates a UpdateRepository5NoContent with default headers values
func NewUpdateRepository5NoContent() *UpdateRepository5NoContent {
	return &UpdateRepository5NoContent{}
}

/* UpdateRepository5NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository5NoContent struct {
}

// IsSuccess returns true when this update repository5 no content response has a 2xx status code
func (o *UpdateRepository5NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository5 no content response has a 3xx status code
func (o *UpdateRepository5NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository5 no content response has a 4xx status code
func (o *UpdateRepository5NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository5 no content response has a 5xx status code
func (o *UpdateRepository5NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository5 no content response a status code equal to that given
func (o *UpdateRepository5NoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateRepository5NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/raw/group/{repositoryName}][%d] updateRepository5NoContent ", 204)
}

func (o *UpdateRepository5NoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/raw/group/{repositoryName}][%d] updateRepository5NoContent ", 204)
}

func (o *UpdateRepository5NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository5Unauthorized creates a UpdateRepository5Unauthorized with default headers values
func NewUpdateRepository5Unauthorized() *UpdateRepository5Unauthorized {
	return &UpdateRepository5Unauthorized{}
}

/* UpdateRepository5Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository5Unauthorized struct {
}

// IsSuccess returns true when this update repository5 unauthorized response has a 2xx status code
func (o *UpdateRepository5Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository5 unauthorized response has a 3xx status code
func (o *UpdateRepository5Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository5 unauthorized response has a 4xx status code
func (o *UpdateRepository5Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository5 unauthorized response has a 5xx status code
func (o *UpdateRepository5Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository5 unauthorized response a status code equal to that given
func (o *UpdateRepository5Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateRepository5Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/raw/group/{repositoryName}][%d] updateRepository5Unauthorized ", 401)
}

func (o *UpdateRepository5Unauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/raw/group/{repositoryName}][%d] updateRepository5Unauthorized ", 401)
}

func (o *UpdateRepository5Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository5Forbidden creates a UpdateRepository5Forbidden with default headers values
func NewUpdateRepository5Forbidden() *UpdateRepository5Forbidden {
	return &UpdateRepository5Forbidden{}
}

/* UpdateRepository5Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository5Forbidden struct {
}

// IsSuccess returns true when this update repository5 forbidden response has a 2xx status code
func (o *UpdateRepository5Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository5 forbidden response has a 3xx status code
func (o *UpdateRepository5Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository5 forbidden response has a 4xx status code
func (o *UpdateRepository5Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository5 forbidden response has a 5xx status code
func (o *UpdateRepository5Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository5 forbidden response a status code equal to that given
func (o *UpdateRepository5Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateRepository5Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/raw/group/{repositoryName}][%d] updateRepository5Forbidden ", 403)
}

func (o *UpdateRepository5Forbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/raw/group/{repositoryName}][%d] updateRepository5Forbidden ", 403)
}

func (o *UpdateRepository5Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}