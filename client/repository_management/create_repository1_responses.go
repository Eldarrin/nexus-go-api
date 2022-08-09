// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateMavenHostedRepositoryReader is a Reader for the CreateMavenHostedRepository structure.
type CreateMavenHostedRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMavenHostedRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateMavenHostedRepositoryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateMavenHostedRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateMavenHostedRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateMavenHostedRepositoryCreated creates a CreateMavenHostedRepositoryCreated with default headers values
func NewCreateMavenHostedRepositoryCreated() *CreateMavenHostedRepositoryCreated {
	return &CreateMavenHostedRepositoryCreated{}
}

/*
	CreateMavenHostedRepositoryCreated describes a response with status code 201, with default header values.

Repository created
*/
type CreateMavenHostedRepositoryCreated struct {
}

// IsSuccess returns true when this create repository1 created response has a 2xx status code
func (o *CreateMavenHostedRepositoryCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create repository1 created response has a 3xx status code
func (o *CreateMavenHostedRepositoryCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository1 created response has a 4xx status code
func (o *CreateMavenHostedRepositoryCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create repository1 created response has a 5xx status code
func (o *CreateMavenHostedRepositoryCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository1 created response a status code equal to that given
func (o *CreateMavenHostedRepositoryCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateMavenHostedRepositoryCreated) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/hosted][%d] createRepository1Created ", 201)
}

func (o *CreateMavenHostedRepositoryCreated) String() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/hosted][%d] createRepository1Created ", 201)
}

func (o *CreateMavenHostedRepositoryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateMavenHostedRepositoryUnauthorized creates a CreateMavenHostedRepositoryUnauthorized with default headers values
func NewCreateMavenHostedRepositoryUnauthorized() *CreateMavenHostedRepositoryUnauthorized {
	return &CreateMavenHostedRepositoryUnauthorized{}
}

/*
	CreateMavenHostedRepositoryUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateMavenHostedRepositoryUnauthorized struct {
}

// IsSuccess returns true when this create repository1 unauthorized response has a 2xx status code
func (o *CreateMavenHostedRepositoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository1 unauthorized response has a 3xx status code
func (o *CreateMavenHostedRepositoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository1 unauthorized response has a 4xx status code
func (o *CreateMavenHostedRepositoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository1 unauthorized response has a 5xx status code
func (o *CreateMavenHostedRepositoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository1 unauthorized response a status code equal to that given
func (o *CreateMavenHostedRepositoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateMavenHostedRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/hosted][%d] createRepository1Unauthorized ", 401)
}

func (o *CreateMavenHostedRepositoryUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/hosted][%d] createRepository1Unauthorized ", 401)
}

func (o *CreateMavenHostedRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateMavenHostedRepositoryForbidden creates a CreateMavenHostedRepositoryForbidden with default headers values
func NewCreateMavenHostedRepositoryForbidden() *CreateMavenHostedRepositoryForbidden {
	return &CreateMavenHostedRepositoryForbidden{}
}

/*
	CreateMavenHostedRepositoryForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateMavenHostedRepositoryForbidden struct {
}

// IsSuccess returns true when this create repository1 forbidden response has a 2xx status code
func (o *CreateMavenHostedRepositoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository1 forbidden response has a 3xx status code
func (o *CreateMavenHostedRepositoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository1 forbidden response has a 4xx status code
func (o *CreateMavenHostedRepositoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository1 forbidden response has a 5xx status code
func (o *CreateMavenHostedRepositoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository1 forbidden response a status code equal to that given
func (o *CreateMavenHostedRepositoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateMavenHostedRepositoryForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/hosted][%d] createRepository1Forbidden ", 403)
}

func (o *CreateMavenHostedRepositoryForbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/hosted][%d] createRepository1Forbidden ", 403)
}

func (o *CreateMavenHostedRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
