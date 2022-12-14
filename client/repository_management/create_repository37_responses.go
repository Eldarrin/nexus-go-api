// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateP2ProxyRepositoryReader is a Reader for the CreateP2ProxyRepository structure.
type CreateP2ProxyRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateP2ProxyRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateP2ProxyRepositoryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateP2ProxyRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateP2ProxyRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateP2ProxyRepositoryMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateP2ProxyRepositoryCreated creates a CreateP2ProxyRepositoryCreated with default headers values
func NewCreateP2ProxyRepositoryCreated() *CreateP2ProxyRepositoryCreated {
	return &CreateP2ProxyRepositoryCreated{}
}

/*
	CreateP2ProxyRepositoryCreated describes a response with status code 201, with default header values.

Repository created
*/
type CreateP2ProxyRepositoryCreated struct {
}

// IsSuccess returns true when this create repository37 created response has a 2xx status code
func (o *CreateP2ProxyRepositoryCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create repository37 created response has a 3xx status code
func (o *CreateP2ProxyRepositoryCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository37 created response has a 4xx status code
func (o *CreateP2ProxyRepositoryCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create repository37 created response has a 5xx status code
func (o *CreateP2ProxyRepositoryCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository37 created response a status code equal to that given
func (o *CreateP2ProxyRepositoryCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateP2ProxyRepositoryCreated) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/p2/proxy][%d] createRepository37Created ", 201)
}

func (o *CreateP2ProxyRepositoryCreated) String() string {
	return fmt.Sprintf("[POST /v1/repositories/p2/proxy][%d] createRepository37Created ", 201)
}

func (o *CreateP2ProxyRepositoryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateP2ProxyRepositoryUnauthorized creates a CreateP2ProxyRepositoryUnauthorized with default headers values
func NewCreateP2ProxyRepositoryUnauthorized() *CreateP2ProxyRepositoryUnauthorized {
	return &CreateP2ProxyRepositoryUnauthorized{}
}

/*
	CreateP2ProxyRepositoryUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateP2ProxyRepositoryUnauthorized struct {
}

// IsSuccess returns true when this create repository37 unauthorized response has a 2xx status code
func (o *CreateP2ProxyRepositoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository37 unauthorized response has a 3xx status code
func (o *CreateP2ProxyRepositoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository37 unauthorized response has a 4xx status code
func (o *CreateP2ProxyRepositoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository37 unauthorized response has a 5xx status code
func (o *CreateP2ProxyRepositoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository37 unauthorized response a status code equal to that given
func (o *CreateP2ProxyRepositoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateP2ProxyRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/p2/proxy][%d] createRepository37Unauthorized ", 401)
}

func (o *CreateP2ProxyRepositoryUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/p2/proxy][%d] createRepository37Unauthorized ", 401)
}

func (o *CreateP2ProxyRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateP2ProxyRepositoryForbidden creates a CreateP2ProxyRepositoryForbidden with default headers values
func NewCreateP2ProxyRepositoryForbidden() *CreateP2ProxyRepositoryForbidden {
	return &CreateP2ProxyRepositoryForbidden{}
}

/*
	CreateP2ProxyRepositoryForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateP2ProxyRepositoryForbidden struct {
}

// IsSuccess returns true when this create repository37 forbidden response has a 2xx status code
func (o *CreateP2ProxyRepositoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository37 forbidden response has a 3xx status code
func (o *CreateP2ProxyRepositoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository37 forbidden response has a 4xx status code
func (o *CreateP2ProxyRepositoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository37 forbidden response has a 5xx status code
func (o *CreateP2ProxyRepositoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository37 forbidden response a status code equal to that given
func (o *CreateP2ProxyRepositoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateP2ProxyRepositoryForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/p2/proxy][%d] createRepository37Forbidden ", 403)
}

func (o *CreateP2ProxyRepositoryForbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/p2/proxy][%d] createRepository37Forbidden ", 403)
}

func (o *CreateP2ProxyRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateP2ProxyRepositoryMethodNotAllowed creates a CreateP2ProxyRepositoryMethodNotAllowed with default headers values
func NewCreateP2ProxyRepositoryMethodNotAllowed() *CreateP2ProxyRepositoryMethodNotAllowed {
	return &CreateP2ProxyRepositoryMethodNotAllowed{}
}

/*
	CreateP2ProxyRepositoryMethodNotAllowed describes a response with status code 405, with default header values.

Feature is disabled in High Availability
*/
type CreateP2ProxyRepositoryMethodNotAllowed struct {
}

// IsSuccess returns true when this create repository37 method not allowed response has a 2xx status code
func (o *CreateP2ProxyRepositoryMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository37 method not allowed response has a 3xx status code
func (o *CreateP2ProxyRepositoryMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository37 method not allowed response has a 4xx status code
func (o *CreateP2ProxyRepositoryMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository37 method not allowed response has a 5xx status code
func (o *CreateP2ProxyRepositoryMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository37 method not allowed response a status code equal to that given
func (o *CreateP2ProxyRepositoryMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *CreateP2ProxyRepositoryMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/p2/proxy][%d] createRepository37MethodNotAllowed ", 405)
}

func (o *CreateP2ProxyRepositoryMethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /v1/repositories/p2/proxy][%d] createRepository37MethodNotAllowed ", 405)
}

func (o *CreateP2ProxyRepositoryMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
