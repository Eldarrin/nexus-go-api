// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository36Reader is a Reader for the CreateRepository36 structure.
type CreateRepository36Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository36Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository36Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository36Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository36Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateRepository36MethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRepository36Created creates a CreateRepository36Created with default headers values
func NewCreateRepository36Created() *CreateRepository36Created {
	return &CreateRepository36Created{}
}

/* CreateRepository36Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository36Created struct {
}

// IsSuccess returns true when this create repository36 created response has a 2xx status code
func (o *CreateRepository36Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create repository36 created response has a 3xx status code
func (o *CreateRepository36Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository36 created response has a 4xx status code
func (o *CreateRepository36Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create repository36 created response has a 5xx status code
func (o *CreateRepository36Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository36 created response a status code equal to that given
func (o *CreateRepository36Created) IsCode(code int) bool {
	return code == 201
}

func (o *CreateRepository36Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] createRepository36Created ", 201)
}

func (o *CreateRepository36Created) String() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] createRepository36Created ", 201)
}

func (o *CreateRepository36Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository36Unauthorized creates a CreateRepository36Unauthorized with default headers values
func NewCreateRepository36Unauthorized() *CreateRepository36Unauthorized {
	return &CreateRepository36Unauthorized{}
}

/* CreateRepository36Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository36Unauthorized struct {
}

// IsSuccess returns true when this create repository36 unauthorized response has a 2xx status code
func (o *CreateRepository36Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository36 unauthorized response has a 3xx status code
func (o *CreateRepository36Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository36 unauthorized response has a 4xx status code
func (o *CreateRepository36Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository36 unauthorized response has a 5xx status code
func (o *CreateRepository36Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository36 unauthorized response a status code equal to that given
func (o *CreateRepository36Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateRepository36Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] createRepository36Unauthorized ", 401)
}

func (o *CreateRepository36Unauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] createRepository36Unauthorized ", 401)
}

func (o *CreateRepository36Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository36Forbidden creates a CreateRepository36Forbidden with default headers values
func NewCreateRepository36Forbidden() *CreateRepository36Forbidden {
	return &CreateRepository36Forbidden{}
}

/* CreateRepository36Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository36Forbidden struct {
}

// IsSuccess returns true when this create repository36 forbidden response has a 2xx status code
func (o *CreateRepository36Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository36 forbidden response has a 3xx status code
func (o *CreateRepository36Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository36 forbidden response has a 4xx status code
func (o *CreateRepository36Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository36 forbidden response has a 5xx status code
func (o *CreateRepository36Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository36 forbidden response a status code equal to that given
func (o *CreateRepository36Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateRepository36Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] createRepository36Forbidden ", 403)
}

func (o *CreateRepository36Forbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] createRepository36Forbidden ", 403)
}

func (o *CreateRepository36Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository36MethodNotAllowed creates a CreateRepository36MethodNotAllowed with default headers values
func NewCreateRepository36MethodNotAllowed() *CreateRepository36MethodNotAllowed {
	return &CreateRepository36MethodNotAllowed{}
}

/* CreateRepository36MethodNotAllowed describes a response with status code 405, with default header values.

Feature is disabled in High Availability
*/
type CreateRepository36MethodNotAllowed struct {
}

// IsSuccess returns true when this create repository36 method not allowed response has a 2xx status code
func (o *CreateRepository36MethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository36 method not allowed response has a 3xx status code
func (o *CreateRepository36MethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository36 method not allowed response has a 4xx status code
func (o *CreateRepository36MethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository36 method not allowed response has a 5xx status code
func (o *CreateRepository36MethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository36 method not allowed response a status code equal to that given
func (o *CreateRepository36MethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *CreateRepository36MethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] createRepository36MethodNotAllowed ", 405)
}

func (o *CreateRepository36MethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] createRepository36MethodNotAllowed ", 405)
}

func (o *CreateRepository36MethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
