// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRawProxyRepositoryReader is a Reader for the CreateRawProxyRepository structure.
type CreateRawProxyRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRawProxyRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRawProxyRepositoryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRawProxyRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRawProxyRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRawProxyRepositoryCreated creates a CreateRawProxyRepositoryCreated with default headers values
func NewCreateRawProxyRepositoryCreated() *CreateRawProxyRepositoryCreated {
	return &CreateRawProxyRepositoryCreated{}
}

/*
	CreateRawProxyRepositoryCreated describes a response with status code 201, with default header values.

Repository created
*/
type CreateRawProxyRepositoryCreated struct {
}

// IsSuccess returns true when this create repository7 created response has a 2xx status code
func (o *CreateRawProxyRepositoryCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create repository7 created response has a 3xx status code
func (o *CreateRawProxyRepositoryCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository7 created response has a 4xx status code
func (o *CreateRawProxyRepositoryCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create repository7 created response has a 5xx status code
func (o *CreateRawProxyRepositoryCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository7 created response a status code equal to that given
func (o *CreateRawProxyRepositoryCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateRawProxyRepositoryCreated) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/raw/proxy][%d] createRepository7Created ", 201)
}

func (o *CreateRawProxyRepositoryCreated) String() string {
	return fmt.Sprintf("[POST /v1/repositories/raw/proxy][%d] createRepository7Created ", 201)
}

func (o *CreateRawProxyRepositoryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRawProxyRepositoryUnauthorized creates a CreateRawProxyRepositoryUnauthorized with default headers values
func NewCreateRawProxyRepositoryUnauthorized() *CreateRawProxyRepositoryUnauthorized {
	return &CreateRawProxyRepositoryUnauthorized{}
}

/*
	CreateRawProxyRepositoryUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRawProxyRepositoryUnauthorized struct {
}

// IsSuccess returns true when this create repository7 unauthorized response has a 2xx status code
func (o *CreateRawProxyRepositoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository7 unauthorized response has a 3xx status code
func (o *CreateRawProxyRepositoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository7 unauthorized response has a 4xx status code
func (o *CreateRawProxyRepositoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository7 unauthorized response has a 5xx status code
func (o *CreateRawProxyRepositoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository7 unauthorized response a status code equal to that given
func (o *CreateRawProxyRepositoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateRawProxyRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/raw/proxy][%d] createRepository7Unauthorized ", 401)
}

func (o *CreateRawProxyRepositoryUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/raw/proxy][%d] createRepository7Unauthorized ", 401)
}

func (o *CreateRawProxyRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRawProxyRepositoryForbidden creates a CreateRawProxyRepositoryForbidden with default headers values
func NewCreateRawProxyRepositoryForbidden() *CreateRawProxyRepositoryForbidden {
	return &CreateRawProxyRepositoryForbidden{}
}

/*
	CreateRawProxyRepositoryForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRawProxyRepositoryForbidden struct {
}

// IsSuccess returns true when this create repository7 forbidden response has a 2xx status code
func (o *CreateRawProxyRepositoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository7 forbidden response has a 3xx status code
func (o *CreateRawProxyRepositoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository7 forbidden response has a 4xx status code
func (o *CreateRawProxyRepositoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository7 forbidden response has a 5xx status code
func (o *CreateRawProxyRepositoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository7 forbidden response a status code equal to that given
func (o *CreateRawProxyRepositoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateRawProxyRepositoryForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/raw/proxy][%d] createRepository7Forbidden ", 403)
}

func (o *CreateRawProxyRepositoryForbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/raw/proxy][%d] createRepository7Forbidden ", 403)
}

func (o *CreateRawProxyRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
