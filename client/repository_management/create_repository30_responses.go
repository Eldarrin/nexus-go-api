// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateConanProxyRepositoryReader is a Reader for the CreateConanProxyRepository structure.
type CreateConanProxyRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateConanProxyRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateConanProxyRepositoryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateConanProxyRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateConanProxyRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateConanProxyRepositoryMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateConanProxyRepositoryCreated creates a CreateConanProxyRepositoryCreated with default headers values
func NewCreateConanProxyRepositoryCreated() *CreateConanProxyRepositoryCreated {
	return &CreateConanProxyRepositoryCreated{}
}

/*
	CreateConanProxyRepositoryCreated describes a response with status code 201, with default header values.

Repository created
*/
type CreateConanProxyRepositoryCreated struct {
}

// IsSuccess returns true when this create repository30 created response has a 2xx status code
func (o *CreateConanProxyRepositoryCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create repository30 created response has a 3xx status code
func (o *CreateConanProxyRepositoryCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository30 created response has a 4xx status code
func (o *CreateConanProxyRepositoryCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create repository30 created response has a 5xx status code
func (o *CreateConanProxyRepositoryCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository30 created response a status code equal to that given
func (o *CreateConanProxyRepositoryCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateConanProxyRepositoryCreated) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/conan/proxy][%d] createRepository30Created ", 201)
}

func (o *CreateConanProxyRepositoryCreated) String() string {
	return fmt.Sprintf("[POST /v1/repositories/conan/proxy][%d] createRepository30Created ", 201)
}

func (o *CreateConanProxyRepositoryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateConanProxyRepositoryUnauthorized creates a CreateConanProxyRepositoryUnauthorized with default headers values
func NewCreateConanProxyRepositoryUnauthorized() *CreateConanProxyRepositoryUnauthorized {
	return &CreateConanProxyRepositoryUnauthorized{}
}

/*
	CreateConanProxyRepositoryUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateConanProxyRepositoryUnauthorized struct {
}

// IsSuccess returns true when this create repository30 unauthorized response has a 2xx status code
func (o *CreateConanProxyRepositoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository30 unauthorized response has a 3xx status code
func (o *CreateConanProxyRepositoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository30 unauthorized response has a 4xx status code
func (o *CreateConanProxyRepositoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository30 unauthorized response has a 5xx status code
func (o *CreateConanProxyRepositoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository30 unauthorized response a status code equal to that given
func (o *CreateConanProxyRepositoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateConanProxyRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/conan/proxy][%d] createRepository30Unauthorized ", 401)
}

func (o *CreateConanProxyRepositoryUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/conan/proxy][%d] createRepository30Unauthorized ", 401)
}

func (o *CreateConanProxyRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateConanProxyRepositoryForbidden creates a CreateConanProxyRepositoryForbidden with default headers values
func NewCreateConanProxyRepositoryForbidden() *CreateConanProxyRepositoryForbidden {
	return &CreateConanProxyRepositoryForbidden{}
}

/*
	CreateConanProxyRepositoryForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateConanProxyRepositoryForbidden struct {
}

// IsSuccess returns true when this create repository30 forbidden response has a 2xx status code
func (o *CreateConanProxyRepositoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository30 forbidden response has a 3xx status code
func (o *CreateConanProxyRepositoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository30 forbidden response has a 4xx status code
func (o *CreateConanProxyRepositoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository30 forbidden response has a 5xx status code
func (o *CreateConanProxyRepositoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository30 forbidden response a status code equal to that given
func (o *CreateConanProxyRepositoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateConanProxyRepositoryForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/conan/proxy][%d] createRepository30Forbidden ", 403)
}

func (o *CreateConanProxyRepositoryForbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/conan/proxy][%d] createRepository30Forbidden ", 403)
}

func (o *CreateConanProxyRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateConanProxyRepositoryMethodNotAllowed creates a CreateConanProxyRepositoryMethodNotAllowed with default headers values
func NewCreateConanProxyRepositoryMethodNotAllowed() *CreateConanProxyRepositoryMethodNotAllowed {
	return &CreateConanProxyRepositoryMethodNotAllowed{}
}

/*
	CreateConanProxyRepositoryMethodNotAllowed describes a response with status code 405, with default header values.

Feature is disabled in High Availability
*/
type CreateConanProxyRepositoryMethodNotAllowed struct {
}

// IsSuccess returns true when this create repository30 method not allowed response has a 2xx status code
func (o *CreateConanProxyRepositoryMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository30 method not allowed response has a 3xx status code
func (o *CreateConanProxyRepositoryMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository30 method not allowed response has a 4xx status code
func (o *CreateConanProxyRepositoryMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository30 method not allowed response has a 5xx status code
func (o *CreateConanProxyRepositoryMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository30 method not allowed response a status code equal to that given
func (o *CreateConanProxyRepositoryMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *CreateConanProxyRepositoryMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/conan/proxy][%d] createRepository30MethodNotAllowed ", 405)
}

func (o *CreateConanProxyRepositoryMethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /v1/repositories/conan/proxy][%d] createRepository30MethodNotAllowed ", 405)
}

func (o *CreateConanProxyRepositoryMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
