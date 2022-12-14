// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateDockerHostedRepositoryReader is a Reader for the UpdateDockerHostedRepository structure.
type UpdateDockerHostedRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDockerHostedRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateDockerHostedRepositoryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateDockerHostedRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateDockerHostedRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateDockerHostedRepositoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateDockerHostedRepositoryNoContent creates a UpdateDockerHostedRepositoryNoContent with default headers values
func NewUpdateDockerHostedRepositoryNoContent() *UpdateDockerHostedRepositoryNoContent {
	return &UpdateDockerHostedRepositoryNoContent{}
}

/*
	UpdateDockerHostedRepositoryNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateDockerHostedRepositoryNoContent struct {
}

// IsSuccess returns true when this update repository18 no content response has a 2xx status code
func (o *UpdateDockerHostedRepositoryNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository18 no content response has a 3xx status code
func (o *UpdateDockerHostedRepositoryNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository18 no content response has a 4xx status code
func (o *UpdateDockerHostedRepositoryNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository18 no content response has a 5xx status code
func (o *UpdateDockerHostedRepositoryNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository18 no content response a status code equal to that given
func (o *UpdateDockerHostedRepositoryNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateDockerHostedRepositoryNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/hosted/{repositoryName}][%d] updateRepository18NoContent ", 204)
}

func (o *UpdateDockerHostedRepositoryNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/hosted/{repositoryName}][%d] updateRepository18NoContent ", 204)
}

func (o *UpdateDockerHostedRepositoryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDockerHostedRepositoryUnauthorized creates a UpdateDockerHostedRepositoryUnauthorized with default headers values
func NewUpdateDockerHostedRepositoryUnauthorized() *UpdateDockerHostedRepositoryUnauthorized {
	return &UpdateDockerHostedRepositoryUnauthorized{}
}

/*
	UpdateDockerHostedRepositoryUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateDockerHostedRepositoryUnauthorized struct {
}

// IsSuccess returns true when this update repository18 unauthorized response has a 2xx status code
func (o *UpdateDockerHostedRepositoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository18 unauthorized response has a 3xx status code
func (o *UpdateDockerHostedRepositoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository18 unauthorized response has a 4xx status code
func (o *UpdateDockerHostedRepositoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository18 unauthorized response has a 5xx status code
func (o *UpdateDockerHostedRepositoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository18 unauthorized response a status code equal to that given
func (o *UpdateDockerHostedRepositoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateDockerHostedRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/hosted/{repositoryName}][%d] updateRepository18Unauthorized ", 401)
}

func (o *UpdateDockerHostedRepositoryUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/hosted/{repositoryName}][%d] updateRepository18Unauthorized ", 401)
}

func (o *UpdateDockerHostedRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDockerHostedRepositoryForbidden creates a UpdateDockerHostedRepositoryForbidden with default headers values
func NewUpdateDockerHostedRepositoryForbidden() *UpdateDockerHostedRepositoryForbidden {
	return &UpdateDockerHostedRepositoryForbidden{}
}

/*
	UpdateDockerHostedRepositoryForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateDockerHostedRepositoryForbidden struct {
}

// IsSuccess returns true when this update repository18 forbidden response has a 2xx status code
func (o *UpdateDockerHostedRepositoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository18 forbidden response has a 3xx status code
func (o *UpdateDockerHostedRepositoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository18 forbidden response has a 4xx status code
func (o *UpdateDockerHostedRepositoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository18 forbidden response has a 5xx status code
func (o *UpdateDockerHostedRepositoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository18 forbidden response a status code equal to that given
func (o *UpdateDockerHostedRepositoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateDockerHostedRepositoryForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/hosted/{repositoryName}][%d] updateRepository18Forbidden ", 403)
}

func (o *UpdateDockerHostedRepositoryForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/hosted/{repositoryName}][%d] updateRepository18Forbidden ", 403)
}

func (o *UpdateDockerHostedRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDockerHostedRepositoryNotFound creates a UpdateDockerHostedRepositoryNotFound with default headers values
func NewUpdateDockerHostedRepositoryNotFound() *UpdateDockerHostedRepositoryNotFound {
	return &UpdateDockerHostedRepositoryNotFound{}
}

/*
	UpdateDockerHostedRepositoryNotFound describes a response with status code 404, with default header values.

Repository not found
*/
type UpdateDockerHostedRepositoryNotFound struct {
}

// IsSuccess returns true when this update repository18 not found response has a 2xx status code
func (o *UpdateDockerHostedRepositoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository18 not found response has a 3xx status code
func (o *UpdateDockerHostedRepositoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository18 not found response has a 4xx status code
func (o *UpdateDockerHostedRepositoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository18 not found response has a 5xx status code
func (o *UpdateDockerHostedRepositoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository18 not found response a status code equal to that given
func (o *UpdateDockerHostedRepositoryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateDockerHostedRepositoryNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/hosted/{repositoryName}][%d] updateRepository18NotFound ", 404)
}

func (o *UpdateDockerHostedRepositoryNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/hosted/{repositoryName}][%d] updateRepository18NotFound ", 404)
}

func (o *UpdateDockerHostedRepositoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
