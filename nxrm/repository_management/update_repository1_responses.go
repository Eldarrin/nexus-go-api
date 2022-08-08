// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository1Reader is a Reader for the UpdateRepository1 structure.
type UpdateRepository1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository1NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRepository1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepository1NoContent creates a UpdateRepository1NoContent with default headers values
func NewUpdateRepository1NoContent() *UpdateRepository1NoContent {
	return &UpdateRepository1NoContent{}
}

/* UpdateRepository1NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository1NoContent struct {
}

// IsSuccess returns true when this update repository1 no content response has a 2xx status code
func (o *UpdateRepository1NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository1 no content response has a 3xx status code
func (o *UpdateRepository1NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository1 no content response has a 4xx status code
func (o *UpdateRepository1NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository1 no content response has a 5xx status code
func (o *UpdateRepository1NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository1 no content response a status code equal to that given
func (o *UpdateRepository1NoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateRepository1NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/hosted/{repositoryName}][%d] updateRepository1NoContent ", 204)
}

func (o *UpdateRepository1NoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/hosted/{repositoryName}][%d] updateRepository1NoContent ", 204)
}

func (o *UpdateRepository1NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository1Unauthorized creates a UpdateRepository1Unauthorized with default headers values
func NewUpdateRepository1Unauthorized() *UpdateRepository1Unauthorized {
	return &UpdateRepository1Unauthorized{}
}

/* UpdateRepository1Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository1Unauthorized struct {
}

// IsSuccess returns true when this update repository1 unauthorized response has a 2xx status code
func (o *UpdateRepository1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository1 unauthorized response has a 3xx status code
func (o *UpdateRepository1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository1 unauthorized response has a 4xx status code
func (o *UpdateRepository1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository1 unauthorized response has a 5xx status code
func (o *UpdateRepository1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository1 unauthorized response a status code equal to that given
func (o *UpdateRepository1Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateRepository1Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/hosted/{repositoryName}][%d] updateRepository1Unauthorized ", 401)
}

func (o *UpdateRepository1Unauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/hosted/{repositoryName}][%d] updateRepository1Unauthorized ", 401)
}

func (o *UpdateRepository1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository1Forbidden creates a UpdateRepository1Forbidden with default headers values
func NewUpdateRepository1Forbidden() *UpdateRepository1Forbidden {
	return &UpdateRepository1Forbidden{}
}

/* UpdateRepository1Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository1Forbidden struct {
}

// IsSuccess returns true when this update repository1 forbidden response has a 2xx status code
func (o *UpdateRepository1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository1 forbidden response has a 3xx status code
func (o *UpdateRepository1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository1 forbidden response has a 4xx status code
func (o *UpdateRepository1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository1 forbidden response has a 5xx status code
func (o *UpdateRepository1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository1 forbidden response a status code equal to that given
func (o *UpdateRepository1Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateRepository1Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/hosted/{repositoryName}][%d] updateRepository1Forbidden ", 403)
}

func (o *UpdateRepository1Forbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/hosted/{repositoryName}][%d] updateRepository1Forbidden ", 403)
}

func (o *UpdateRepository1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}