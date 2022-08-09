// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdatePypiGroupRepositoryReader is a Reader for the UpdatePypiGroupRepository structure.
type UpdatePypiGroupRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePypiGroupRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdatePypiGroupRepositoryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdatePypiGroupRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdatePypiGroupRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdatePypiGroupRepositoryNoContent creates a UpdatePypiGroupRepositoryNoContent with default headers values
func NewUpdatePypiGroupRepositoryNoContent() *UpdatePypiGroupRepositoryNoContent {
	return &UpdatePypiGroupRepositoryNoContent{}
}

/*
	UpdatePypiGroupRepositoryNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdatePypiGroupRepositoryNoContent struct {
}

// IsSuccess returns true when this update repository26 no content response has a 2xx status code
func (o *UpdatePypiGroupRepositoryNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository26 no content response has a 3xx status code
func (o *UpdatePypiGroupRepositoryNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository26 no content response has a 4xx status code
func (o *UpdatePypiGroupRepositoryNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository26 no content response has a 5xx status code
func (o *UpdatePypiGroupRepositoryNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository26 no content response a status code equal to that given
func (o *UpdatePypiGroupRepositoryNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdatePypiGroupRepositoryNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/group/{repositoryName}][%d] updateRepository26NoContent ", 204)
}

func (o *UpdatePypiGroupRepositoryNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/group/{repositoryName}][%d] updateRepository26NoContent ", 204)
}

func (o *UpdatePypiGroupRepositoryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePypiGroupRepositoryUnauthorized creates a UpdatePypiGroupRepositoryUnauthorized with default headers values
func NewUpdatePypiGroupRepositoryUnauthorized() *UpdatePypiGroupRepositoryUnauthorized {
	return &UpdatePypiGroupRepositoryUnauthorized{}
}

/*
	UpdatePypiGroupRepositoryUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdatePypiGroupRepositoryUnauthorized struct {
}

// IsSuccess returns true when this update repository26 unauthorized response has a 2xx status code
func (o *UpdatePypiGroupRepositoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository26 unauthorized response has a 3xx status code
func (o *UpdatePypiGroupRepositoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository26 unauthorized response has a 4xx status code
func (o *UpdatePypiGroupRepositoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository26 unauthorized response has a 5xx status code
func (o *UpdatePypiGroupRepositoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository26 unauthorized response a status code equal to that given
func (o *UpdatePypiGroupRepositoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdatePypiGroupRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/group/{repositoryName}][%d] updateRepository26Unauthorized ", 401)
}

func (o *UpdatePypiGroupRepositoryUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/group/{repositoryName}][%d] updateRepository26Unauthorized ", 401)
}

func (o *UpdatePypiGroupRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePypiGroupRepositoryForbidden creates a UpdatePypiGroupRepositoryForbidden with default headers values
func NewUpdatePypiGroupRepositoryForbidden() *UpdatePypiGroupRepositoryForbidden {
	return &UpdatePypiGroupRepositoryForbidden{}
}

/*
	UpdatePypiGroupRepositoryForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdatePypiGroupRepositoryForbidden struct {
}

// IsSuccess returns true when this update repository26 forbidden response has a 2xx status code
func (o *UpdatePypiGroupRepositoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository26 forbidden response has a 3xx status code
func (o *UpdatePypiGroupRepositoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository26 forbidden response has a 4xx status code
func (o *UpdatePypiGroupRepositoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository26 forbidden response has a 5xx status code
func (o *UpdatePypiGroupRepositoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository26 forbidden response a status code equal to that given
func (o *UpdatePypiGroupRepositoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdatePypiGroupRepositoryForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/group/{repositoryName}][%d] updateRepository26Forbidden ", 403)
}

func (o *UpdatePypiGroupRepositoryForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/group/{repositoryName}][%d] updateRepository26Forbidden ", 403)
}

func (o *UpdatePypiGroupRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
