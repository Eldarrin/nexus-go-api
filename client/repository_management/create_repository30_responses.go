// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository30Reader is a Reader for the CreateRepository30 structure.
type CreateRepository30Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository30Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository30Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository30Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository30Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateRepository30MethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRepository30Created creates a CreateRepository30Created with default headers values
func NewCreateRepository30Created() *CreateRepository30Created {
	return &CreateRepository30Created{}
}

/* CreateRepository30Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository30Created struct {
}

// IsSuccess returns true when this create repository30 created response has a 2xx status code
func (o *CreateRepository30Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create repository30 created response has a 3xx status code
func (o *CreateRepository30Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository30 created response has a 4xx status code
func (o *CreateRepository30Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create repository30 created response has a 5xx status code
func (o *CreateRepository30Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository30 created response a status code equal to that given
func (o *CreateRepository30Created) IsCode(code int) bool {
	return code == 201
}

func (o *CreateRepository30Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/conan/proxy][%d] createRepository30Created ", 201)
}

func (o *CreateRepository30Created) String() string {
	return fmt.Sprintf("[POST /v1/repositories/conan/proxy][%d] createRepository30Created ", 201)
}

func (o *CreateRepository30Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository30Unauthorized creates a CreateRepository30Unauthorized with default headers values
func NewCreateRepository30Unauthorized() *CreateRepository30Unauthorized {
	return &CreateRepository30Unauthorized{}
}

/* CreateRepository30Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository30Unauthorized struct {
}

// IsSuccess returns true when this create repository30 unauthorized response has a 2xx status code
func (o *CreateRepository30Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository30 unauthorized response has a 3xx status code
func (o *CreateRepository30Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository30 unauthorized response has a 4xx status code
func (o *CreateRepository30Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository30 unauthorized response has a 5xx status code
func (o *CreateRepository30Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository30 unauthorized response a status code equal to that given
func (o *CreateRepository30Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateRepository30Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/conan/proxy][%d] createRepository30Unauthorized ", 401)
}

func (o *CreateRepository30Unauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/conan/proxy][%d] createRepository30Unauthorized ", 401)
}

func (o *CreateRepository30Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository30Forbidden creates a CreateRepository30Forbidden with default headers values
func NewCreateRepository30Forbidden() *CreateRepository30Forbidden {
	return &CreateRepository30Forbidden{}
}

/* CreateRepository30Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository30Forbidden struct {
}

// IsSuccess returns true when this create repository30 forbidden response has a 2xx status code
func (o *CreateRepository30Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository30 forbidden response has a 3xx status code
func (o *CreateRepository30Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository30 forbidden response has a 4xx status code
func (o *CreateRepository30Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository30 forbidden response has a 5xx status code
func (o *CreateRepository30Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository30 forbidden response a status code equal to that given
func (o *CreateRepository30Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateRepository30Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/conan/proxy][%d] createRepository30Forbidden ", 403)
}

func (o *CreateRepository30Forbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/conan/proxy][%d] createRepository30Forbidden ", 403)
}

func (o *CreateRepository30Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository30MethodNotAllowed creates a CreateRepository30MethodNotAllowed with default headers values
func NewCreateRepository30MethodNotAllowed() *CreateRepository30MethodNotAllowed {
	return &CreateRepository30MethodNotAllowed{}
}

/* CreateRepository30MethodNotAllowed describes a response with status code 405, with default header values.

Feature is disabled in High Availability
*/
type CreateRepository30MethodNotAllowed struct {
}

// IsSuccess returns true when this create repository30 method not allowed response has a 2xx status code
func (o *CreateRepository30MethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository30 method not allowed response has a 3xx status code
func (o *CreateRepository30MethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository30 method not allowed response has a 4xx status code
func (o *CreateRepository30MethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository30 method not allowed response has a 5xx status code
func (o *CreateRepository30MethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository30 method not allowed response a status code equal to that given
func (o *CreateRepository30MethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *CreateRepository30MethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/conan/proxy][%d] createRepository30MethodNotAllowed ", 405)
}

func (o *CreateRepository30MethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /v1/repositories/conan/proxy][%d] createRepository30MethodNotAllowed ", 405)
}

func (o *CreateRepository30MethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}