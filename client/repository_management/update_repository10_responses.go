// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateNpmProxyRepositoryReader is a Reader for the UpdateNpmProxyRepository structure.
type UpdateNpmProxyRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNpmProxyRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateNpmProxyRepositoryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateNpmProxyRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateNpmProxyRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNpmProxyRepositoryNoContent creates a UpdateNpmProxyRepositoryNoContent with default headers values
func NewUpdateNpmProxyRepositoryNoContent() *UpdateNpmProxyRepositoryNoContent {
	return &UpdateNpmProxyRepositoryNoContent{}
}

/*
	UpdateNpmProxyRepositoryNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateNpmProxyRepositoryNoContent struct {
}

// IsSuccess returns true when this update repository10 no content response has a 2xx status code
func (o *UpdateNpmProxyRepositoryNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository10 no content response has a 3xx status code
func (o *UpdateNpmProxyRepositoryNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository10 no content response has a 4xx status code
func (o *UpdateNpmProxyRepositoryNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository10 no content response has a 5xx status code
func (o *UpdateNpmProxyRepositoryNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository10 no content response a status code equal to that given
func (o *UpdateNpmProxyRepositoryNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateNpmProxyRepositoryNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/npm/proxy/{repositoryName}][%d] updateRepository10NoContent ", 204)
}

func (o *UpdateNpmProxyRepositoryNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/npm/proxy/{repositoryName}][%d] updateRepository10NoContent ", 204)
}

func (o *UpdateNpmProxyRepositoryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateNpmProxyRepositoryUnauthorized creates a UpdateNpmProxyRepositoryUnauthorized with default headers values
func NewUpdateNpmProxyRepositoryUnauthorized() *UpdateNpmProxyRepositoryUnauthorized {
	return &UpdateNpmProxyRepositoryUnauthorized{}
}

/*
	UpdateNpmProxyRepositoryUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateNpmProxyRepositoryUnauthorized struct {
}

// IsSuccess returns true when this update repository10 unauthorized response has a 2xx status code
func (o *UpdateNpmProxyRepositoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository10 unauthorized response has a 3xx status code
func (o *UpdateNpmProxyRepositoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository10 unauthorized response has a 4xx status code
func (o *UpdateNpmProxyRepositoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository10 unauthorized response has a 5xx status code
func (o *UpdateNpmProxyRepositoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository10 unauthorized response a status code equal to that given
func (o *UpdateNpmProxyRepositoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateNpmProxyRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/npm/proxy/{repositoryName}][%d] updateRepository10Unauthorized ", 401)
}

func (o *UpdateNpmProxyRepositoryUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/npm/proxy/{repositoryName}][%d] updateRepository10Unauthorized ", 401)
}

func (o *UpdateNpmProxyRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateNpmProxyRepositoryForbidden creates a UpdateNpmProxyRepositoryForbidden with default headers values
func NewUpdateNpmProxyRepositoryForbidden() *UpdateNpmProxyRepositoryForbidden {
	return &UpdateNpmProxyRepositoryForbidden{}
}

/*
	UpdateNpmProxyRepositoryForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateNpmProxyRepositoryForbidden struct {
}

// IsSuccess returns true when this update repository10 forbidden response has a 2xx status code
func (o *UpdateNpmProxyRepositoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository10 forbidden response has a 3xx status code
func (o *UpdateNpmProxyRepositoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository10 forbidden response has a 4xx status code
func (o *UpdateNpmProxyRepositoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository10 forbidden response has a 5xx status code
func (o *UpdateNpmProxyRepositoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository10 forbidden response a status code equal to that given
func (o *UpdateNpmProxyRepositoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateNpmProxyRepositoryForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/npm/proxy/{repositoryName}][%d] updateRepository10Forbidden ", 403)
}

func (o *UpdateNpmProxyRepositoryForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/npm/proxy/{repositoryName}][%d] updateRepository10Forbidden ", 403)
}

func (o *UpdateNpmProxyRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
