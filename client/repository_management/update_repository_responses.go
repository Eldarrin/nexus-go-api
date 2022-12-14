// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateMavenGroupRepositoryReader is a Reader for the UpdateMavenGroupRepository structure.
type UpdateMavenGroupRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMavenGroupRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateMavenGroupRepositoryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateMavenGroupRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateMavenGroupRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateMavenGroupRepositoryNoContent creates a UpdateMavenGroupRepositoryNoContent with default headers values
func NewUpdateMavenGroupRepositoryNoContent() *UpdateMavenGroupRepositoryNoContent {
	return &UpdateMavenGroupRepositoryNoContent{}
}

/*
	UpdateMavenGroupRepositoryNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateMavenGroupRepositoryNoContent struct {
}

// IsSuccess returns true when this update repository no content response has a 2xx status code
func (o *UpdateMavenGroupRepositoryNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository no content response has a 3xx status code
func (o *UpdateMavenGroupRepositoryNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository no content response has a 4xx status code
func (o *UpdateMavenGroupRepositoryNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository no content response has a 5xx status code
func (o *UpdateMavenGroupRepositoryNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository no content response a status code equal to that given
func (o *UpdateMavenGroupRepositoryNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateMavenGroupRepositoryNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/group/{repositoryName}][%d] updateRepositoryNoContent ", 204)
}

func (o *UpdateMavenGroupRepositoryNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/group/{repositoryName}][%d] updateRepositoryNoContent ", 204)
}

func (o *UpdateMavenGroupRepositoryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMavenGroupRepositoryUnauthorized creates a UpdateMavenGroupRepositoryUnauthorized with default headers values
func NewUpdateMavenGroupRepositoryUnauthorized() *UpdateMavenGroupRepositoryUnauthorized {
	return &UpdateMavenGroupRepositoryUnauthorized{}
}

/*
	UpdateMavenGroupRepositoryUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateMavenGroupRepositoryUnauthorized struct {
}

// IsSuccess returns true when this update repository unauthorized response has a 2xx status code
func (o *UpdateMavenGroupRepositoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository unauthorized response has a 3xx status code
func (o *UpdateMavenGroupRepositoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository unauthorized response has a 4xx status code
func (o *UpdateMavenGroupRepositoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository unauthorized response has a 5xx status code
func (o *UpdateMavenGroupRepositoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository unauthorized response a status code equal to that given
func (o *UpdateMavenGroupRepositoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateMavenGroupRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/group/{repositoryName}][%d] updateRepositoryUnauthorized ", 401)
}

func (o *UpdateMavenGroupRepositoryUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/group/{repositoryName}][%d] updateRepositoryUnauthorized ", 401)
}

func (o *UpdateMavenGroupRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMavenGroupRepositoryForbidden creates a UpdateMavenGroupRepositoryForbidden with default headers values
func NewUpdateMavenGroupRepositoryForbidden() *UpdateMavenGroupRepositoryForbidden {
	return &UpdateMavenGroupRepositoryForbidden{}
}

/*
	UpdateMavenGroupRepositoryForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateMavenGroupRepositoryForbidden struct {
}

// IsSuccess returns true when this update repository forbidden response has a 2xx status code
func (o *UpdateMavenGroupRepositoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository forbidden response has a 3xx status code
func (o *UpdateMavenGroupRepositoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository forbidden response has a 4xx status code
func (o *UpdateMavenGroupRepositoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository forbidden response has a 5xx status code
func (o *UpdateMavenGroupRepositoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository forbidden response a status code equal to that given
func (o *UpdateMavenGroupRepositoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateMavenGroupRepositoryForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/group/{repositoryName}][%d] updateRepositoryForbidden ", 403)
}

func (o *UpdateMavenGroupRepositoryForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/group/{repositoryName}][%d] updateRepositoryForbidden ", 403)
}

func (o *UpdateMavenGroupRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
