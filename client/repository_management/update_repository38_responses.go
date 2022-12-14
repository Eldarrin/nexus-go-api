// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateBowerGroupRepositoryReader is a Reader for the UpdateBowerGroupRepository structure.
type UpdateBowerGroupRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBowerGroupRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateBowerGroupRepositoryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateBowerGroupRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateBowerGroupRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateBowerGroupRepositoryNoContent creates a UpdateBowerGroupRepositoryNoContent with default headers values
func NewUpdateBowerGroupRepositoryNoContent() *UpdateBowerGroupRepositoryNoContent {
	return &UpdateBowerGroupRepositoryNoContent{}
}

/*
	UpdateBowerGroupRepositoryNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateBowerGroupRepositoryNoContent struct {
}

// IsSuccess returns true when this update repository38 no content response has a 2xx status code
func (o *UpdateBowerGroupRepositoryNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository38 no content response has a 3xx status code
func (o *UpdateBowerGroupRepositoryNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository38 no content response has a 4xx status code
func (o *UpdateBowerGroupRepositoryNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository38 no content response has a 5xx status code
func (o *UpdateBowerGroupRepositoryNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository38 no content response a status code equal to that given
func (o *UpdateBowerGroupRepositoryNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateBowerGroupRepositoryNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/group/{repositoryName}][%d] updateRepository38NoContent ", 204)
}

func (o *UpdateBowerGroupRepositoryNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/group/{repositoryName}][%d] updateRepository38NoContent ", 204)
}

func (o *UpdateBowerGroupRepositoryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateBowerGroupRepositoryUnauthorized creates a UpdateBowerGroupRepositoryUnauthorized with default headers values
func NewUpdateBowerGroupRepositoryUnauthorized() *UpdateBowerGroupRepositoryUnauthorized {
	return &UpdateBowerGroupRepositoryUnauthorized{}
}

/*
	UpdateBowerGroupRepositoryUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateBowerGroupRepositoryUnauthorized struct {
}

// IsSuccess returns true when this update repository38 unauthorized response has a 2xx status code
func (o *UpdateBowerGroupRepositoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository38 unauthorized response has a 3xx status code
func (o *UpdateBowerGroupRepositoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository38 unauthorized response has a 4xx status code
func (o *UpdateBowerGroupRepositoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository38 unauthorized response has a 5xx status code
func (o *UpdateBowerGroupRepositoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository38 unauthorized response a status code equal to that given
func (o *UpdateBowerGroupRepositoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateBowerGroupRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/group/{repositoryName}][%d] updateRepository38Unauthorized ", 401)
}

func (o *UpdateBowerGroupRepositoryUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/group/{repositoryName}][%d] updateRepository38Unauthorized ", 401)
}

func (o *UpdateBowerGroupRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateBowerGroupRepositoryForbidden creates a UpdateBowerGroupRepositoryForbidden with default headers values
func NewUpdateBowerGroupRepositoryForbidden() *UpdateBowerGroupRepositoryForbidden {
	return &UpdateBowerGroupRepositoryForbidden{}
}

/*
	UpdateBowerGroupRepositoryForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateBowerGroupRepositoryForbidden struct {
}

// IsSuccess returns true when this update repository38 forbidden response has a 2xx status code
func (o *UpdateBowerGroupRepositoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository38 forbidden response has a 3xx status code
func (o *UpdateBowerGroupRepositoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository38 forbidden response has a 4xx status code
func (o *UpdateBowerGroupRepositoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository38 forbidden response has a 5xx status code
func (o *UpdateBowerGroupRepositoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository38 forbidden response a status code equal to that given
func (o *UpdateBowerGroupRepositoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateBowerGroupRepositoryForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/group/{repositoryName}][%d] updateRepository38Forbidden ", 403)
}

func (o *UpdateBowerGroupRepositoryForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/group/{repositoryName}][%d] updateRepository38Forbidden ", 403)
}

func (o *UpdateBowerGroupRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
