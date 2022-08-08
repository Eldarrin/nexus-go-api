// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository31Reader is a Reader for the UpdateRepository31 structure.
type UpdateRepository31Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository31Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository31NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRepository31Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository31Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepository31NoContent creates a UpdateRepository31NoContent with default headers values
func NewUpdateRepository31NoContent() *UpdateRepository31NoContent {
	return &UpdateRepository31NoContent{}
}

/* UpdateRepository31NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository31NoContent struct {
}

// IsSuccess returns true when this update repository31 no content response has a 2xx status code
func (o *UpdateRepository31NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository31 no content response has a 3xx status code
func (o *UpdateRepository31NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository31 no content response has a 4xx status code
func (o *UpdateRepository31NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository31 no content response has a 5xx status code
func (o *UpdateRepository31NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository31 no content response a status code equal to that given
func (o *UpdateRepository31NoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateRepository31NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/r/group/{repositoryName}][%d] updateRepository31NoContent ", 204)
}

func (o *UpdateRepository31NoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/r/group/{repositoryName}][%d] updateRepository31NoContent ", 204)
}

func (o *UpdateRepository31NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository31Unauthorized creates a UpdateRepository31Unauthorized with default headers values
func NewUpdateRepository31Unauthorized() *UpdateRepository31Unauthorized {
	return &UpdateRepository31Unauthorized{}
}

/* UpdateRepository31Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository31Unauthorized struct {
}

// IsSuccess returns true when this update repository31 unauthorized response has a 2xx status code
func (o *UpdateRepository31Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository31 unauthorized response has a 3xx status code
func (o *UpdateRepository31Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository31 unauthorized response has a 4xx status code
func (o *UpdateRepository31Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository31 unauthorized response has a 5xx status code
func (o *UpdateRepository31Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository31 unauthorized response a status code equal to that given
func (o *UpdateRepository31Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateRepository31Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/r/group/{repositoryName}][%d] updateRepository31Unauthorized ", 401)
}

func (o *UpdateRepository31Unauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/r/group/{repositoryName}][%d] updateRepository31Unauthorized ", 401)
}

func (o *UpdateRepository31Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository31Forbidden creates a UpdateRepository31Forbidden with default headers values
func NewUpdateRepository31Forbidden() *UpdateRepository31Forbidden {
	return &UpdateRepository31Forbidden{}
}

/* UpdateRepository31Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository31Forbidden struct {
}

// IsSuccess returns true when this update repository31 forbidden response has a 2xx status code
func (o *UpdateRepository31Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository31 forbidden response has a 3xx status code
func (o *UpdateRepository31Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository31 forbidden response has a 4xx status code
func (o *UpdateRepository31Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository31 forbidden response has a 5xx status code
func (o *UpdateRepository31Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository31 forbidden response a status code equal to that given
func (o *UpdateRepository31Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateRepository31Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/r/group/{repositoryName}][%d] updateRepository31Forbidden ", 403)
}

func (o *UpdateRepository31Forbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/r/group/{repositoryName}][%d] updateRepository31Forbidden ", 403)
}

func (o *UpdateRepository31Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}