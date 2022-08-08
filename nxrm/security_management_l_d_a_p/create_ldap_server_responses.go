// Code generated by go-swagger; DO NOT EDIT.

package security_management_l_d_a_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateLdapServerReader is a Reader for the CreateLdapServer structure.
type CreateLdapServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLdapServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateLdapServerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateLdapServerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateLdapServerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateLdapServerCreated creates a CreateLdapServerCreated with default headers values
func NewCreateLdapServerCreated() *CreateLdapServerCreated {
	return &CreateLdapServerCreated{}
}

/* CreateLdapServerCreated describes a response with status code 201, with default header values.

LDAP server created
*/
type CreateLdapServerCreated struct {
}

// IsSuccess returns true when this create ldap server created response has a 2xx status code
func (o *CreateLdapServerCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create ldap server created response has a 3xx status code
func (o *CreateLdapServerCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create ldap server created response has a 4xx status code
func (o *CreateLdapServerCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create ldap server created response has a 5xx status code
func (o *CreateLdapServerCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create ldap server created response a status code equal to that given
func (o *CreateLdapServerCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateLdapServerCreated) Error() string {
	return fmt.Sprintf("[POST /v1/security/ldap][%d] createLdapServerCreated ", 201)
}

func (o *CreateLdapServerCreated) String() string {
	return fmt.Sprintf("[POST /v1/security/ldap][%d] createLdapServerCreated ", 201)
}

func (o *CreateLdapServerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateLdapServerUnauthorized creates a CreateLdapServerUnauthorized with default headers values
func NewCreateLdapServerUnauthorized() *CreateLdapServerUnauthorized {
	return &CreateLdapServerUnauthorized{}
}

/* CreateLdapServerUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateLdapServerUnauthorized struct {
}

// IsSuccess returns true when this create ldap server unauthorized response has a 2xx status code
func (o *CreateLdapServerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create ldap server unauthorized response has a 3xx status code
func (o *CreateLdapServerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create ldap server unauthorized response has a 4xx status code
func (o *CreateLdapServerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create ldap server unauthorized response has a 5xx status code
func (o *CreateLdapServerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create ldap server unauthorized response a status code equal to that given
func (o *CreateLdapServerUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateLdapServerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/security/ldap][%d] createLdapServerUnauthorized ", 401)
}

func (o *CreateLdapServerUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/security/ldap][%d] createLdapServerUnauthorized ", 401)
}

func (o *CreateLdapServerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateLdapServerForbidden creates a CreateLdapServerForbidden with default headers values
func NewCreateLdapServerForbidden() *CreateLdapServerForbidden {
	return &CreateLdapServerForbidden{}
}

/* CreateLdapServerForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateLdapServerForbidden struct {
}

// IsSuccess returns true when this create ldap server forbidden response has a 2xx status code
func (o *CreateLdapServerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create ldap server forbidden response has a 3xx status code
func (o *CreateLdapServerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create ldap server forbidden response has a 4xx status code
func (o *CreateLdapServerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create ldap server forbidden response has a 5xx status code
func (o *CreateLdapServerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create ldap server forbidden response a status code equal to that given
func (o *CreateLdapServerForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateLdapServerForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/security/ldap][%d] createLdapServerForbidden ", 403)
}

func (o *CreateLdapServerForbidden) String() string {
	return fmt.Sprintf("[POST /v1/security/ldap][%d] createLdapServerForbidden ", 403)
}

func (o *CreateLdapServerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}