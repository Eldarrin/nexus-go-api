// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateBowerProxyRepositoryReader is a Reader for the UpdateBowerProxyRepository structure.
type UpdateBowerProxyRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBowerProxyRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateBowerProxyRepositoryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateBowerProxyRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateBowerProxyRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateBowerProxyRepositoryNoContent creates a UpdateBowerProxyRepositoryNoContent with default headers values
func NewUpdateBowerProxyRepositoryNoContent() *UpdateBowerProxyRepositoryNoContent {
	return &UpdateBowerProxyRepositoryNoContent{}
}

/*
	UpdateBowerProxyRepositoryNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateBowerProxyRepositoryNoContent struct {
}

// IsSuccess returns true when this update repository40 no content response has a 2xx status code
func (o *UpdateBowerProxyRepositoryNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository40 no content response has a 3xx status code
func (o *UpdateBowerProxyRepositoryNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository40 no content response has a 4xx status code
func (o *UpdateBowerProxyRepositoryNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository40 no content response has a 5xx status code
func (o *UpdateBowerProxyRepositoryNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository40 no content response a status code equal to that given
func (o *UpdateBowerProxyRepositoryNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateBowerProxyRepositoryNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/proxy/{repositoryName}][%d] updateRepository40NoContent ", 204)
}

func (o *UpdateBowerProxyRepositoryNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/proxy/{repositoryName}][%d] updateRepository40NoContent ", 204)
}

func (o *UpdateBowerProxyRepositoryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateBowerProxyRepositoryUnauthorized creates a UpdateBowerProxyRepositoryUnauthorized with default headers values
func NewUpdateBowerProxyRepositoryUnauthorized() *UpdateBowerProxyRepositoryUnauthorized {
	return &UpdateBowerProxyRepositoryUnauthorized{}
}

/*
	UpdateBowerProxyRepositoryUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateBowerProxyRepositoryUnauthorized struct {
}

// IsSuccess returns true when this update repository40 unauthorized response has a 2xx status code
func (o *UpdateBowerProxyRepositoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository40 unauthorized response has a 3xx status code
func (o *UpdateBowerProxyRepositoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository40 unauthorized response has a 4xx status code
func (o *UpdateBowerProxyRepositoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository40 unauthorized response has a 5xx status code
func (o *UpdateBowerProxyRepositoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository40 unauthorized response a status code equal to that given
func (o *UpdateBowerProxyRepositoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateBowerProxyRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/proxy/{repositoryName}][%d] updateRepository40Unauthorized ", 401)
}

func (o *UpdateBowerProxyRepositoryUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/proxy/{repositoryName}][%d] updateRepository40Unauthorized ", 401)
}

func (o *UpdateBowerProxyRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateBowerProxyRepositoryForbidden creates a UpdateBowerProxyRepositoryForbidden with default headers values
func NewUpdateBowerProxyRepositoryForbidden() *UpdateBowerProxyRepositoryForbidden {
	return &UpdateBowerProxyRepositoryForbidden{}
}

/*
	UpdateBowerProxyRepositoryForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateBowerProxyRepositoryForbidden struct {
}

// IsSuccess returns true when this update repository40 forbidden response has a 2xx status code
func (o *UpdateBowerProxyRepositoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository40 forbidden response has a 3xx status code
func (o *UpdateBowerProxyRepositoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository40 forbidden response has a 4xx status code
func (o *UpdateBowerProxyRepositoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository40 forbidden response has a 5xx status code
func (o *UpdateBowerProxyRepositoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository40 forbidden response a status code equal to that given
func (o *UpdateBowerProxyRepositoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateBowerProxyRepositoryForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/proxy/{repositoryName}][%d] updateRepository40Forbidden ", 403)
}

func (o *UpdateBowerProxyRepositoryForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/proxy/{repositoryName}][%d] updateRepository40Forbidden ", 403)
}

func (o *UpdateBowerProxyRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
