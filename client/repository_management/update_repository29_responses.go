// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateCondaProxyRepositoryReader is a Reader for the UpdateCondaProxyRepository structure.
type UpdateCondaProxyRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCondaProxyRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateCondaProxyRepositoryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateCondaProxyRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateCondaProxyRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateCondaProxyRepositoryNoContent creates a UpdateCondaProxyRepositoryNoContent with default headers values
func NewUpdateCondaProxyRepositoryNoContent() *UpdateCondaProxyRepositoryNoContent {
	return &UpdateCondaProxyRepositoryNoContent{}
}

/*
	UpdateCondaProxyRepositoryNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateCondaProxyRepositoryNoContent struct {
}

// IsSuccess returns true when this update repository29 no content response has a 2xx status code
func (o *UpdateCondaProxyRepositoryNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository29 no content response has a 3xx status code
func (o *UpdateCondaProxyRepositoryNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository29 no content response has a 4xx status code
func (o *UpdateCondaProxyRepositoryNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository29 no content response has a 5xx status code
func (o *UpdateCondaProxyRepositoryNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository29 no content response a status code equal to that given
func (o *UpdateCondaProxyRepositoryNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateCondaProxyRepositoryNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/conda/proxy/{repositoryName}][%d] updateRepository29NoContent ", 204)
}

func (o *UpdateCondaProxyRepositoryNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/conda/proxy/{repositoryName}][%d] updateRepository29NoContent ", 204)
}

func (o *UpdateCondaProxyRepositoryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateCondaProxyRepositoryUnauthorized creates a UpdateCondaProxyRepositoryUnauthorized with default headers values
func NewUpdateCondaProxyRepositoryUnauthorized() *UpdateCondaProxyRepositoryUnauthorized {
	return &UpdateCondaProxyRepositoryUnauthorized{}
}

/*
	UpdateCondaProxyRepositoryUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateCondaProxyRepositoryUnauthorized struct {
}

// IsSuccess returns true when this update repository29 unauthorized response has a 2xx status code
func (o *UpdateCondaProxyRepositoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository29 unauthorized response has a 3xx status code
func (o *UpdateCondaProxyRepositoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository29 unauthorized response has a 4xx status code
func (o *UpdateCondaProxyRepositoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository29 unauthorized response has a 5xx status code
func (o *UpdateCondaProxyRepositoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository29 unauthorized response a status code equal to that given
func (o *UpdateCondaProxyRepositoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateCondaProxyRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/conda/proxy/{repositoryName}][%d] updateRepository29Unauthorized ", 401)
}

func (o *UpdateCondaProxyRepositoryUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/conda/proxy/{repositoryName}][%d] updateRepository29Unauthorized ", 401)
}

func (o *UpdateCondaProxyRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateCondaProxyRepositoryForbidden creates a UpdateCondaProxyRepositoryForbidden with default headers values
func NewUpdateCondaProxyRepositoryForbidden() *UpdateCondaProxyRepositoryForbidden {
	return &UpdateCondaProxyRepositoryForbidden{}
}

/*
	UpdateCondaProxyRepositoryForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateCondaProxyRepositoryForbidden struct {
}

// IsSuccess returns true when this update repository29 forbidden response has a 2xx status code
func (o *UpdateCondaProxyRepositoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository29 forbidden response has a 3xx status code
func (o *UpdateCondaProxyRepositoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository29 forbidden response has a 4xx status code
func (o *UpdateCondaProxyRepositoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository29 forbidden response has a 5xx status code
func (o *UpdateCondaProxyRepositoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository29 forbidden response a status code equal to that given
func (o *UpdateCondaProxyRepositoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateCondaProxyRepositoryForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/conda/proxy/{repositoryName}][%d] updateRepository29Forbidden ", 403)
}

func (o *UpdateCondaProxyRepositoryForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/conda/proxy/{repositoryName}][%d] updateRepository29Forbidden ", 403)
}

func (o *UpdateCondaProxyRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
