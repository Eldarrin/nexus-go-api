// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateBowerHostedRepositoryReader is a Reader for the UpdateBowerHostedRepository structure.
type UpdateBowerHostedRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBowerHostedRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateBowerHostedRepositoryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateBowerHostedRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateBowerHostedRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateBowerHostedRepositoryNoContent creates a UpdateBowerHostedRepositoryNoContent with default headers values
func NewUpdateBowerHostedRepositoryNoContent() *UpdateBowerHostedRepositoryNoContent {
	return &UpdateBowerHostedRepositoryNoContent{}
}

/*
	UpdateBowerHostedRepositoryNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateBowerHostedRepositoryNoContent struct {
}

// IsSuccess returns true when this update repository39 no content response has a 2xx status code
func (o *UpdateBowerHostedRepositoryNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository39 no content response has a 3xx status code
func (o *UpdateBowerHostedRepositoryNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository39 no content response has a 4xx status code
func (o *UpdateBowerHostedRepositoryNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository39 no content response has a 5xx status code
func (o *UpdateBowerHostedRepositoryNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository39 no content response a status code equal to that given
func (o *UpdateBowerHostedRepositoryNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateBowerHostedRepositoryNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/hosted/{repositoryName}][%d] updateRepository39NoContent ", 204)
}

func (o *UpdateBowerHostedRepositoryNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/hosted/{repositoryName}][%d] updateRepository39NoContent ", 204)
}

func (o *UpdateBowerHostedRepositoryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateBowerHostedRepositoryUnauthorized creates a UpdateBowerHostedRepositoryUnauthorized with default headers values
func NewUpdateBowerHostedRepositoryUnauthorized() *UpdateBowerHostedRepositoryUnauthorized {
	return &UpdateBowerHostedRepositoryUnauthorized{}
}

/*
	UpdateBowerHostedRepositoryUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateBowerHostedRepositoryUnauthorized struct {
}

// IsSuccess returns true when this update repository39 unauthorized response has a 2xx status code
func (o *UpdateBowerHostedRepositoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository39 unauthorized response has a 3xx status code
func (o *UpdateBowerHostedRepositoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository39 unauthorized response has a 4xx status code
func (o *UpdateBowerHostedRepositoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository39 unauthorized response has a 5xx status code
func (o *UpdateBowerHostedRepositoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository39 unauthorized response a status code equal to that given
func (o *UpdateBowerHostedRepositoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateBowerHostedRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/hosted/{repositoryName}][%d] updateRepository39Unauthorized ", 401)
}

func (o *UpdateBowerHostedRepositoryUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/hosted/{repositoryName}][%d] updateRepository39Unauthorized ", 401)
}

func (o *UpdateBowerHostedRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateBowerHostedRepositoryForbidden creates a UpdateBowerHostedRepositoryForbidden with default headers values
func NewUpdateBowerHostedRepositoryForbidden() *UpdateBowerHostedRepositoryForbidden {
	return &UpdateBowerHostedRepositoryForbidden{}
}

/*
	UpdateBowerHostedRepositoryForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateBowerHostedRepositoryForbidden struct {
}

// IsSuccess returns true when this update repository39 forbidden response has a 2xx status code
func (o *UpdateBowerHostedRepositoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository39 forbidden response has a 3xx status code
func (o *UpdateBowerHostedRepositoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository39 forbidden response has a 4xx status code
func (o *UpdateBowerHostedRepositoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository39 forbidden response has a 5xx status code
func (o *UpdateBowerHostedRepositoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository39 forbidden response a status code equal to that given
func (o *UpdateBowerHostedRepositoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateBowerHostedRepositoryForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/hosted/{repositoryName}][%d] updateRepository39Forbidden ", 403)
}

func (o *UpdateBowerHostedRepositoryForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/hosted/{repositoryName}][%d] updateRepository39Forbidden ", 403)
}

func (o *UpdateBowerHostedRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
