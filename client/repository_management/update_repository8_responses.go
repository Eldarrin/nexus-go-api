// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateNpmGroupRepositoryReader is a Reader for the UpdateNpmGroupRepository structure.
type UpdateNpmGroupRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNpmGroupRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateNpmGroupRepositoryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateNpmGroupRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateNpmGroupRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNpmGroupRepositoryNoContent creates a UpdateNpmGroupRepositoryNoContent with default headers values
func NewUpdateNpmGroupRepositoryNoContent() *UpdateNpmGroupRepositoryNoContent {
	return &UpdateNpmGroupRepositoryNoContent{}
}

/*
	UpdateNpmGroupRepositoryNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateNpmGroupRepositoryNoContent struct {
}

// IsSuccess returns true when this update repository8 no content response has a 2xx status code
func (o *UpdateNpmGroupRepositoryNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository8 no content response has a 3xx status code
func (o *UpdateNpmGroupRepositoryNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository8 no content response has a 4xx status code
func (o *UpdateNpmGroupRepositoryNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository8 no content response has a 5xx status code
func (o *UpdateNpmGroupRepositoryNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository8 no content response a status code equal to that given
func (o *UpdateNpmGroupRepositoryNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateNpmGroupRepositoryNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/npm/group/{repositoryName}][%d] updateRepository8NoContent ", 204)
}

func (o *UpdateNpmGroupRepositoryNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/npm/group/{repositoryName}][%d] updateRepository8NoContent ", 204)
}

func (o *UpdateNpmGroupRepositoryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateNpmGroupRepositoryUnauthorized creates a UpdateNpmGroupRepositoryUnauthorized with default headers values
func NewUpdateNpmGroupRepositoryUnauthorized() *UpdateNpmGroupRepositoryUnauthorized {
	return &UpdateNpmGroupRepositoryUnauthorized{}
}

/*
	UpdateNpmGroupRepositoryUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateNpmGroupRepositoryUnauthorized struct {
}

// IsSuccess returns true when this update repository8 unauthorized response has a 2xx status code
func (o *UpdateNpmGroupRepositoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository8 unauthorized response has a 3xx status code
func (o *UpdateNpmGroupRepositoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository8 unauthorized response has a 4xx status code
func (o *UpdateNpmGroupRepositoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository8 unauthorized response has a 5xx status code
func (o *UpdateNpmGroupRepositoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository8 unauthorized response a status code equal to that given
func (o *UpdateNpmGroupRepositoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateNpmGroupRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/npm/group/{repositoryName}][%d] updateRepository8Unauthorized ", 401)
}

func (o *UpdateNpmGroupRepositoryUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/npm/group/{repositoryName}][%d] updateRepository8Unauthorized ", 401)
}

func (o *UpdateNpmGroupRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateNpmGroupRepositoryForbidden creates a UpdateNpmGroupRepositoryForbidden with default headers values
func NewUpdateNpmGroupRepositoryForbidden() *UpdateNpmGroupRepositoryForbidden {
	return &UpdateNpmGroupRepositoryForbidden{}
}

/*
	UpdateNpmGroupRepositoryForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateNpmGroupRepositoryForbidden struct {
}

// IsSuccess returns true when this update repository8 forbidden response has a 2xx status code
func (o *UpdateNpmGroupRepositoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository8 forbidden response has a 3xx status code
func (o *UpdateNpmGroupRepositoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository8 forbidden response has a 4xx status code
func (o *UpdateNpmGroupRepositoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository8 forbidden response has a 5xx status code
func (o *UpdateNpmGroupRepositoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository8 forbidden response a status code equal to that given
func (o *UpdateNpmGroupRepositoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateNpmGroupRepositoryForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/npm/group/{repositoryName}][%d] updateRepository8Forbidden ", 403)
}

func (o *UpdateNpmGroupRepositoryForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/npm/group/{repositoryName}][%d] updateRepository8Forbidden ", 403)
}

func (o *UpdateNpmGroupRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
