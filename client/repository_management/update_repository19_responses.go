// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateDockerProxyRepositoryReader is a Reader for the UpdateDockerProxyRepository structure.
type UpdateDockerProxyRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDockerProxyRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateDockerProxyRepositoryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateDockerProxyRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateDockerProxyRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateDockerProxyRepositoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateDockerProxyRepositoryNoContent creates a UpdateDockerProxyRepositoryNoContent with default headers values
func NewUpdateDockerProxyRepositoryNoContent() *UpdateDockerProxyRepositoryNoContent {
	return &UpdateDockerProxyRepositoryNoContent{}
}

/*
	UpdateDockerProxyRepositoryNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateDockerProxyRepositoryNoContent struct {
}

// IsSuccess returns true when this update repository19 no content response has a 2xx status code
func (o *UpdateDockerProxyRepositoryNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository19 no content response has a 3xx status code
func (o *UpdateDockerProxyRepositoryNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository19 no content response has a 4xx status code
func (o *UpdateDockerProxyRepositoryNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository19 no content response has a 5xx status code
func (o *UpdateDockerProxyRepositoryNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository19 no content response a status code equal to that given
func (o *UpdateDockerProxyRepositoryNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateDockerProxyRepositoryNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/proxy/{repositoryName}][%d] updateRepository19NoContent ", 204)
}

func (o *UpdateDockerProxyRepositoryNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/proxy/{repositoryName}][%d] updateRepository19NoContent ", 204)
}

func (o *UpdateDockerProxyRepositoryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDockerProxyRepositoryUnauthorized creates a UpdateDockerProxyRepositoryUnauthorized with default headers values
func NewUpdateDockerProxyRepositoryUnauthorized() *UpdateDockerProxyRepositoryUnauthorized {
	return &UpdateDockerProxyRepositoryUnauthorized{}
}

/*
	UpdateDockerProxyRepositoryUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateDockerProxyRepositoryUnauthorized struct {
}

// IsSuccess returns true when this update repository19 unauthorized response has a 2xx status code
func (o *UpdateDockerProxyRepositoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository19 unauthorized response has a 3xx status code
func (o *UpdateDockerProxyRepositoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository19 unauthorized response has a 4xx status code
func (o *UpdateDockerProxyRepositoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository19 unauthorized response has a 5xx status code
func (o *UpdateDockerProxyRepositoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository19 unauthorized response a status code equal to that given
func (o *UpdateDockerProxyRepositoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateDockerProxyRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/proxy/{repositoryName}][%d] updateRepository19Unauthorized ", 401)
}

func (o *UpdateDockerProxyRepositoryUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/proxy/{repositoryName}][%d] updateRepository19Unauthorized ", 401)
}

func (o *UpdateDockerProxyRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDockerProxyRepositoryForbidden creates a UpdateDockerProxyRepositoryForbidden with default headers values
func NewUpdateDockerProxyRepositoryForbidden() *UpdateDockerProxyRepositoryForbidden {
	return &UpdateDockerProxyRepositoryForbidden{}
}

/*
	UpdateDockerProxyRepositoryForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateDockerProxyRepositoryForbidden struct {
}

// IsSuccess returns true when this update repository19 forbidden response has a 2xx status code
func (o *UpdateDockerProxyRepositoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository19 forbidden response has a 3xx status code
func (o *UpdateDockerProxyRepositoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository19 forbidden response has a 4xx status code
func (o *UpdateDockerProxyRepositoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository19 forbidden response has a 5xx status code
func (o *UpdateDockerProxyRepositoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository19 forbidden response a status code equal to that given
func (o *UpdateDockerProxyRepositoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateDockerProxyRepositoryForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/proxy/{repositoryName}][%d] updateRepository19Forbidden ", 403)
}

func (o *UpdateDockerProxyRepositoryForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/proxy/{repositoryName}][%d] updateRepository19Forbidden ", 403)
}

func (o *UpdateDockerProxyRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDockerProxyRepositoryNotFound creates a UpdateDockerProxyRepositoryNotFound with default headers values
func NewUpdateDockerProxyRepositoryNotFound() *UpdateDockerProxyRepositoryNotFound {
	return &UpdateDockerProxyRepositoryNotFound{}
}

/*
	UpdateDockerProxyRepositoryNotFound describes a response with status code 404, with default header values.

Repository not found
*/
type UpdateDockerProxyRepositoryNotFound struct {
}

// IsSuccess returns true when this update repository19 not found response has a 2xx status code
func (o *UpdateDockerProxyRepositoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository19 not found response has a 3xx status code
func (o *UpdateDockerProxyRepositoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository19 not found response has a 4xx status code
func (o *UpdateDockerProxyRepositoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository19 not found response has a 5xx status code
func (o *UpdateDockerProxyRepositoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository19 not found response a status code equal to that given
func (o *UpdateDockerProxyRepositoryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateDockerProxyRepositoryNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/proxy/{repositoryName}][%d] updateRepository19NotFound ", 404)
}

func (o *UpdateDockerProxyRepositoryNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/proxy/{repositoryName}][%d] updateRepository19NotFound ", 404)
}

func (o *UpdateDockerProxyRepositoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
