// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRProxyRepositoryReader is a Reader for the UpdateRProxyRepository structure.
type UpdateRProxyRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRProxyRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRProxyRepositoryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRProxyRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRProxyRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRProxyRepositoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRProxyRepositoryNoContent creates a UpdateRProxyRepositoryNoContent with default headers values
func NewUpdateRProxyRepositoryNoContent() *UpdateRProxyRepositoryNoContent {
	return &UpdateRProxyRepositoryNoContent{}
}

/*
	UpdateRProxyRepositoryNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRProxyRepositoryNoContent struct {
}

// IsSuccess returns true when this update repository33 no content response has a 2xx status code
func (o *UpdateRProxyRepositoryNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository33 no content response has a 3xx status code
func (o *UpdateRProxyRepositoryNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository33 no content response has a 4xx status code
func (o *UpdateRProxyRepositoryNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository33 no content response has a 5xx status code
func (o *UpdateRProxyRepositoryNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository33 no content response a status code equal to that given
func (o *UpdateRProxyRepositoryNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateRProxyRepositoryNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/r/proxy/{repositoryName}][%d] updateRepository33NoContent ", 204)
}

func (o *UpdateRProxyRepositoryNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/r/proxy/{repositoryName}][%d] updateRepository33NoContent ", 204)
}

func (o *UpdateRProxyRepositoryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRProxyRepositoryUnauthorized creates a UpdateRProxyRepositoryUnauthorized with default headers values
func NewUpdateRProxyRepositoryUnauthorized() *UpdateRProxyRepositoryUnauthorized {
	return &UpdateRProxyRepositoryUnauthorized{}
}

/*
	UpdateRProxyRepositoryUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRProxyRepositoryUnauthorized struct {
}

// IsSuccess returns true when this update repository33 unauthorized response has a 2xx status code
func (o *UpdateRProxyRepositoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository33 unauthorized response has a 3xx status code
func (o *UpdateRProxyRepositoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository33 unauthorized response has a 4xx status code
func (o *UpdateRProxyRepositoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository33 unauthorized response has a 5xx status code
func (o *UpdateRProxyRepositoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository33 unauthorized response a status code equal to that given
func (o *UpdateRProxyRepositoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateRProxyRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/r/proxy/{repositoryName}][%d] updateRepository33Unauthorized ", 401)
}

func (o *UpdateRProxyRepositoryUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/r/proxy/{repositoryName}][%d] updateRepository33Unauthorized ", 401)
}

func (o *UpdateRProxyRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRProxyRepositoryForbidden creates a UpdateRProxyRepositoryForbidden with default headers values
func NewUpdateRProxyRepositoryForbidden() *UpdateRProxyRepositoryForbidden {
	return &UpdateRProxyRepositoryForbidden{}
}

/*
	UpdateRProxyRepositoryForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRProxyRepositoryForbidden struct {
}

// IsSuccess returns true when this update repository33 forbidden response has a 2xx status code
func (o *UpdateRProxyRepositoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository33 forbidden response has a 3xx status code
func (o *UpdateRProxyRepositoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository33 forbidden response has a 4xx status code
func (o *UpdateRProxyRepositoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository33 forbidden response has a 5xx status code
func (o *UpdateRProxyRepositoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository33 forbidden response a status code equal to that given
func (o *UpdateRProxyRepositoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateRProxyRepositoryForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/r/proxy/{repositoryName}][%d] updateRepository33Forbidden ", 403)
}

func (o *UpdateRProxyRepositoryForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/r/proxy/{repositoryName}][%d] updateRepository33Forbidden ", 403)
}

func (o *UpdateRProxyRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRProxyRepositoryNotFound creates a UpdateRProxyRepositoryNotFound with default headers values
func NewUpdateRProxyRepositoryNotFound() *UpdateRProxyRepositoryNotFound {
	return &UpdateRProxyRepositoryNotFound{}
}

/*
	UpdateRProxyRepositoryNotFound describes a response with status code 404, with default header values.

Repository not found
*/
type UpdateRProxyRepositoryNotFound struct {
}

// IsSuccess returns true when this update repository33 not found response has a 2xx status code
func (o *UpdateRProxyRepositoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository33 not found response has a 3xx status code
func (o *UpdateRProxyRepositoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository33 not found response has a 4xx status code
func (o *UpdateRProxyRepositoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository33 not found response has a 5xx status code
func (o *UpdateRProxyRepositoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository33 not found response a status code equal to that given
func (o *UpdateRProxyRepositoryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateRProxyRepositoryNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/r/proxy/{repositoryName}][%d] updateRepository33NotFound ", 404)
}

func (o *UpdateRProxyRepositoryNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/r/proxy/{repositoryName}][%d] updateRepository33NotFound ", 404)
}

func (o *UpdateRProxyRepositoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
