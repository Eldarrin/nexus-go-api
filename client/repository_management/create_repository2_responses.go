// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateMavenProxyRepositoryReader is a Reader for the CreateMavenProxyRepository structure.
type CreateMavenProxyRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMavenProxyRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateMavenProxyRepositoryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateMavenProxyRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateMavenProxyRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateMavenProxyRepositoryCreated creates a CreateMavenProxyRepositoryCreated with default headers values
func NewCreateMavenProxyRepositoryCreated() *CreateMavenProxyRepositoryCreated {
	return &CreateMavenProxyRepositoryCreated{}
}

/*
	CreateMavenProxyRepositoryCreated describes a response with status code 201, with default header values.

Repository created
*/
type CreateMavenProxyRepositoryCreated struct {
}

// IsSuccess returns true when this create repository2 created response has a 2xx status code
func (o *CreateMavenProxyRepositoryCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create repository2 created response has a 3xx status code
func (o *CreateMavenProxyRepositoryCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository2 created response has a 4xx status code
func (o *CreateMavenProxyRepositoryCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create repository2 created response has a 5xx status code
func (o *CreateMavenProxyRepositoryCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository2 created response a status code equal to that given
func (o *CreateMavenProxyRepositoryCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateMavenProxyRepositoryCreated) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/proxy][%d] createRepository2Created ", 201)
}

func (o *CreateMavenProxyRepositoryCreated) String() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/proxy][%d] createRepository2Created ", 201)
}

func (o *CreateMavenProxyRepositoryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateMavenProxyRepositoryUnauthorized creates a CreateMavenProxyRepositoryUnauthorized with default headers values
func NewCreateMavenProxyRepositoryUnauthorized() *CreateMavenProxyRepositoryUnauthorized {
	return &CreateMavenProxyRepositoryUnauthorized{}
}

/*
	CreateMavenProxyRepositoryUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateMavenProxyRepositoryUnauthorized struct {
}

// IsSuccess returns true when this create repository2 unauthorized response has a 2xx status code
func (o *CreateMavenProxyRepositoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository2 unauthorized response has a 3xx status code
func (o *CreateMavenProxyRepositoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository2 unauthorized response has a 4xx status code
func (o *CreateMavenProxyRepositoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository2 unauthorized response has a 5xx status code
func (o *CreateMavenProxyRepositoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository2 unauthorized response a status code equal to that given
func (o *CreateMavenProxyRepositoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateMavenProxyRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/proxy][%d] createRepository2Unauthorized ", 401)
}

func (o *CreateMavenProxyRepositoryUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/proxy][%d] createRepository2Unauthorized ", 401)
}

func (o *CreateMavenProxyRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateMavenProxyRepositoryForbidden creates a CreateMavenProxyRepositoryForbidden with default headers values
func NewCreateMavenProxyRepositoryForbidden() *CreateMavenProxyRepositoryForbidden {
	return &CreateMavenProxyRepositoryForbidden{}
}

/*
	CreateMavenProxyRepositoryForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateMavenProxyRepositoryForbidden struct {
}

// IsSuccess returns true when this create repository2 forbidden response has a 2xx status code
func (o *CreateMavenProxyRepositoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository2 forbidden response has a 3xx status code
func (o *CreateMavenProxyRepositoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository2 forbidden response has a 4xx status code
func (o *CreateMavenProxyRepositoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository2 forbidden response has a 5xx status code
func (o *CreateMavenProxyRepositoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository2 forbidden response a status code equal to that given
func (o *CreateMavenProxyRepositoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateMavenProxyRepositoryForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/proxy][%d] createRepository2Forbidden ", 403)
}

func (o *CreateMavenProxyRepositoryForbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/proxy][%d] createRepository2Forbidden ", 403)
}

func (o *CreateMavenProxyRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
