// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateGitLfsHostedRepositoryReader is a Reader for the UpdateGitLfsHostedRepository structure.
type UpdateGitLfsHostedRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateGitLfsHostedRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateGitLfsHostedRepositoryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateGitLfsHostedRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateGitLfsHostedRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateGitLfsHostedRepositoryNoContent creates a UpdateGitLfsHostedRepositoryNoContent with default headers values
func NewUpdateGitLfsHostedRepositoryNoContent() *UpdateGitLfsHostedRepositoryNoContent {
	return &UpdateGitLfsHostedRepositoryNoContent{}
}

/*
	UpdateGitLfsHostedRepositoryNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateGitLfsHostedRepositoryNoContent struct {
}

// IsSuccess returns true when this update repository25 no content response has a 2xx status code
func (o *UpdateGitLfsHostedRepositoryNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository25 no content response has a 3xx status code
func (o *UpdateGitLfsHostedRepositoryNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository25 no content response has a 4xx status code
func (o *UpdateGitLfsHostedRepositoryNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository25 no content response has a 5xx status code
func (o *UpdateGitLfsHostedRepositoryNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository25 no content response a status code equal to that given
func (o *UpdateGitLfsHostedRepositoryNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateGitLfsHostedRepositoryNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}][%d] updateRepository25NoContent ", 204)
}

func (o *UpdateGitLfsHostedRepositoryNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}][%d] updateRepository25NoContent ", 204)
}

func (o *UpdateGitLfsHostedRepositoryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateGitLfsHostedRepositoryUnauthorized creates a UpdateGitLfsHostedRepositoryUnauthorized with default headers values
func NewUpdateGitLfsHostedRepositoryUnauthorized() *UpdateGitLfsHostedRepositoryUnauthorized {
	return &UpdateGitLfsHostedRepositoryUnauthorized{}
}

/*
	UpdateGitLfsHostedRepositoryUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateGitLfsHostedRepositoryUnauthorized struct {
}

// IsSuccess returns true when this update repository25 unauthorized response has a 2xx status code
func (o *UpdateGitLfsHostedRepositoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository25 unauthorized response has a 3xx status code
func (o *UpdateGitLfsHostedRepositoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository25 unauthorized response has a 4xx status code
func (o *UpdateGitLfsHostedRepositoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository25 unauthorized response has a 5xx status code
func (o *UpdateGitLfsHostedRepositoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository25 unauthorized response a status code equal to that given
func (o *UpdateGitLfsHostedRepositoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateGitLfsHostedRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}][%d] updateRepository25Unauthorized ", 401)
}

func (o *UpdateGitLfsHostedRepositoryUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}][%d] updateRepository25Unauthorized ", 401)
}

func (o *UpdateGitLfsHostedRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateGitLfsHostedRepositoryForbidden creates a UpdateGitLfsHostedRepositoryForbidden with default headers values
func NewUpdateGitLfsHostedRepositoryForbidden() *UpdateGitLfsHostedRepositoryForbidden {
	return &UpdateGitLfsHostedRepositoryForbidden{}
}

/*
	UpdateGitLfsHostedRepositoryForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateGitLfsHostedRepositoryForbidden struct {
}

// IsSuccess returns true when this update repository25 forbidden response has a 2xx status code
func (o *UpdateGitLfsHostedRepositoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository25 forbidden response has a 3xx status code
func (o *UpdateGitLfsHostedRepositoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository25 forbidden response has a 4xx status code
func (o *UpdateGitLfsHostedRepositoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository25 forbidden response has a 5xx status code
func (o *UpdateGitLfsHostedRepositoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository25 forbidden response a status code equal to that given
func (o *UpdateGitLfsHostedRepositoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateGitLfsHostedRepositoryForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}][%d] updateRepository25Forbidden ", 403)
}

func (o *UpdateGitLfsHostedRepositoryForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}][%d] updateRepository25Forbidden ", 403)
}

func (o *UpdateGitLfsHostedRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
