// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository40Reader is a Reader for the UpdateRepository40 structure.
type UpdateRepository40Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository40Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository40NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRepository40Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository40Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepository40NoContent creates a UpdateRepository40NoContent with default headers values
func NewUpdateRepository40NoContent() *UpdateRepository40NoContent {
	return &UpdateRepository40NoContent{}
}

/* UpdateRepository40NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository40NoContent struct {
}

// IsSuccess returns true when this update repository40 no content response has a 2xx status code
func (o *UpdateRepository40NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository40 no content response has a 3xx status code
func (o *UpdateRepository40NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository40 no content response has a 4xx status code
func (o *UpdateRepository40NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository40 no content response has a 5xx status code
func (o *UpdateRepository40NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository40 no content response a status code equal to that given
func (o *UpdateRepository40NoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateRepository40NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/proxy/{repositoryName}][%d] updateRepository40NoContent ", 204)
}

func (o *UpdateRepository40NoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/proxy/{repositoryName}][%d] updateRepository40NoContent ", 204)
}

func (o *UpdateRepository40NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository40Unauthorized creates a UpdateRepository40Unauthorized with default headers values
func NewUpdateRepository40Unauthorized() *UpdateRepository40Unauthorized {
	return &UpdateRepository40Unauthorized{}
}

/* UpdateRepository40Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository40Unauthorized struct {
}

// IsSuccess returns true when this update repository40 unauthorized response has a 2xx status code
func (o *UpdateRepository40Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository40 unauthorized response has a 3xx status code
func (o *UpdateRepository40Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository40 unauthorized response has a 4xx status code
func (o *UpdateRepository40Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository40 unauthorized response has a 5xx status code
func (o *UpdateRepository40Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository40 unauthorized response a status code equal to that given
func (o *UpdateRepository40Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateRepository40Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/proxy/{repositoryName}][%d] updateRepository40Unauthorized ", 401)
}

func (o *UpdateRepository40Unauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/proxy/{repositoryName}][%d] updateRepository40Unauthorized ", 401)
}

func (o *UpdateRepository40Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository40Forbidden creates a UpdateRepository40Forbidden with default headers values
func NewUpdateRepository40Forbidden() *UpdateRepository40Forbidden {
	return &UpdateRepository40Forbidden{}
}

/* UpdateRepository40Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository40Forbidden struct {
}

// IsSuccess returns true when this update repository40 forbidden response has a 2xx status code
func (o *UpdateRepository40Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository40 forbidden response has a 3xx status code
func (o *UpdateRepository40Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository40 forbidden response has a 4xx status code
func (o *UpdateRepository40Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository40 forbidden response has a 5xx status code
func (o *UpdateRepository40Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository40 forbidden response a status code equal to that given
func (o *UpdateRepository40Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateRepository40Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/proxy/{repositoryName}][%d] updateRepository40Forbidden ", 403)
}

func (o *UpdateRepository40Forbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/proxy/{repositoryName}][%d] updateRepository40Forbidden ", 403)
}

func (o *UpdateRepository40Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
