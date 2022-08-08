// Code generated by go-swagger; DO NOT EDIT.

package security_management_l_d_a_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateLdapServerReader is a Reader for the UpdateLdapServer structure.
type UpdateLdapServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateLdapServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateLdapServerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateLdapServerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateLdapServerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateLdapServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateLdapServerNoContent creates a UpdateLdapServerNoContent with default headers values
func NewUpdateLdapServerNoContent() *UpdateLdapServerNoContent {
	return &UpdateLdapServerNoContent{}
}

/* UpdateLdapServerNoContent describes a response with status code 204, with default header values.

LDAP server updated
*/
type UpdateLdapServerNoContent struct {
}

// IsSuccess returns true when this update ldap server no content response has a 2xx status code
func (o *UpdateLdapServerNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update ldap server no content response has a 3xx status code
func (o *UpdateLdapServerNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update ldap server no content response has a 4xx status code
func (o *UpdateLdapServerNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update ldap server no content response has a 5xx status code
func (o *UpdateLdapServerNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update ldap server no content response a status code equal to that given
func (o *UpdateLdapServerNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateLdapServerNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/security/ldap/{name}][%d] updateLdapServerNoContent ", 204)
}

func (o *UpdateLdapServerNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/security/ldap/{name}][%d] updateLdapServerNoContent ", 204)
}

func (o *UpdateLdapServerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateLdapServerUnauthorized creates a UpdateLdapServerUnauthorized with default headers values
func NewUpdateLdapServerUnauthorized() *UpdateLdapServerUnauthorized {
	return &UpdateLdapServerUnauthorized{}
}

/* UpdateLdapServerUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateLdapServerUnauthorized struct {
}

// IsSuccess returns true when this update ldap server unauthorized response has a 2xx status code
func (o *UpdateLdapServerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update ldap server unauthorized response has a 3xx status code
func (o *UpdateLdapServerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update ldap server unauthorized response has a 4xx status code
func (o *UpdateLdapServerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update ldap server unauthorized response has a 5xx status code
func (o *UpdateLdapServerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update ldap server unauthorized response a status code equal to that given
func (o *UpdateLdapServerUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateLdapServerUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/security/ldap/{name}][%d] updateLdapServerUnauthorized ", 401)
}

func (o *UpdateLdapServerUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/security/ldap/{name}][%d] updateLdapServerUnauthorized ", 401)
}

func (o *UpdateLdapServerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateLdapServerForbidden creates a UpdateLdapServerForbidden with default headers values
func NewUpdateLdapServerForbidden() *UpdateLdapServerForbidden {
	return &UpdateLdapServerForbidden{}
}

/* UpdateLdapServerForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateLdapServerForbidden struct {
}

// IsSuccess returns true when this update ldap server forbidden response has a 2xx status code
func (o *UpdateLdapServerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update ldap server forbidden response has a 3xx status code
func (o *UpdateLdapServerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update ldap server forbidden response has a 4xx status code
func (o *UpdateLdapServerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update ldap server forbidden response has a 5xx status code
func (o *UpdateLdapServerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update ldap server forbidden response a status code equal to that given
func (o *UpdateLdapServerForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateLdapServerForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/security/ldap/{name}][%d] updateLdapServerForbidden ", 403)
}

func (o *UpdateLdapServerForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/security/ldap/{name}][%d] updateLdapServerForbidden ", 403)
}

func (o *UpdateLdapServerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateLdapServerNotFound creates a UpdateLdapServerNotFound with default headers values
func NewUpdateLdapServerNotFound() *UpdateLdapServerNotFound {
	return &UpdateLdapServerNotFound{}
}

/* UpdateLdapServerNotFound describes a response with status code 404, with default header values.

LDAP server not found
*/
type UpdateLdapServerNotFound struct {
}

// IsSuccess returns true when this update ldap server not found response has a 2xx status code
func (o *UpdateLdapServerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update ldap server not found response has a 3xx status code
func (o *UpdateLdapServerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update ldap server not found response has a 4xx status code
func (o *UpdateLdapServerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update ldap server not found response has a 5xx status code
func (o *UpdateLdapServerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update ldap server not found response a status code equal to that given
func (o *UpdateLdapServerNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateLdapServerNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/security/ldap/{name}][%d] updateLdapServerNotFound ", 404)
}

func (o *UpdateLdapServerNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/security/ldap/{name}][%d] updateLdapServerNotFound ", 404)
}

func (o *UpdateLdapServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}