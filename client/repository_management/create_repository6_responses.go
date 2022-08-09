// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository6Reader is a Reader for the CreateRepository6 structure.
type CreateRepository6Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository6Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository6Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository6Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository6Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRepository6Created creates a CreateRepository6Created with default headers values
func NewCreateRepository6Created() *CreateRepository6Created {
	return &CreateRepository6Created{}
}

/* CreateRepository6Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository6Created struct {
}

// IsSuccess returns true when this create repository6 created response has a 2xx status code
func (o *CreateRepository6Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create repository6 created response has a 3xx status code
func (o *CreateRepository6Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository6 created response has a 4xx status code
func (o *CreateRepository6Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create repository6 created response has a 5xx status code
func (o *CreateRepository6Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository6 created response a status code equal to that given
func (o *CreateRepository6Created) IsCode(code int) bool {
	return code == 201
}

func (o *CreateRepository6Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/raw/hosted][%d] createRepository6Created ", 201)
}

func (o *CreateRepository6Created) String() string {
	return fmt.Sprintf("[POST /v1/repositories/raw/hosted][%d] createRepository6Created ", 201)
}

func (o *CreateRepository6Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository6Unauthorized creates a CreateRepository6Unauthorized with default headers values
func NewCreateRepository6Unauthorized() *CreateRepository6Unauthorized {
	return &CreateRepository6Unauthorized{}
}

/* CreateRepository6Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository6Unauthorized struct {
}

// IsSuccess returns true when this create repository6 unauthorized response has a 2xx status code
func (o *CreateRepository6Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository6 unauthorized response has a 3xx status code
func (o *CreateRepository6Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository6 unauthorized response has a 4xx status code
func (o *CreateRepository6Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository6 unauthorized response has a 5xx status code
func (o *CreateRepository6Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository6 unauthorized response a status code equal to that given
func (o *CreateRepository6Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateRepository6Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/raw/hosted][%d] createRepository6Unauthorized ", 401)
}

func (o *CreateRepository6Unauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/raw/hosted][%d] createRepository6Unauthorized ", 401)
}

func (o *CreateRepository6Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository6Forbidden creates a CreateRepository6Forbidden with default headers values
func NewCreateRepository6Forbidden() *CreateRepository6Forbidden {
	return &CreateRepository6Forbidden{}
}

/* CreateRepository6Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository6Forbidden struct {
}

// IsSuccess returns true when this create repository6 forbidden response has a 2xx status code
func (o *CreateRepository6Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository6 forbidden response has a 3xx status code
func (o *CreateRepository6Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository6 forbidden response has a 4xx status code
func (o *CreateRepository6Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository6 forbidden response has a 5xx status code
func (o *CreateRepository6Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository6 forbidden response a status code equal to that given
func (o *CreateRepository6Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateRepository6Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/raw/hosted][%d] createRepository6Forbidden ", 403)
}

func (o *CreateRepository6Forbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/raw/hosted][%d] createRepository6Forbidden ", 403)
}

func (o *CreateRepository6Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}