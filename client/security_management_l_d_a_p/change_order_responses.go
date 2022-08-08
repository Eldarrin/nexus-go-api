// Code generated by go-swagger; DO NOT EDIT.

package security_management_l_d_a_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ChangeOrderReader is a Reader for the ChangeOrder structure.
type ChangeOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewChangeOrderNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewChangeOrderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewChangeOrderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewChangeOrderNoContent creates a ChangeOrderNoContent with default headers values
func NewChangeOrderNoContent() *ChangeOrderNoContent {
	return &ChangeOrderNoContent{}
}

/* ChangeOrderNoContent describes a response with status code 204, with default header values.

LDAP server order changed
*/
type ChangeOrderNoContent struct {
}

// IsSuccess returns true when this change order no content response has a 2xx status code
func (o *ChangeOrderNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this change order no content response has a 3xx status code
func (o *ChangeOrderNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change order no content response has a 4xx status code
func (o *ChangeOrderNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this change order no content response has a 5xx status code
func (o *ChangeOrderNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this change order no content response a status code equal to that given
func (o *ChangeOrderNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ChangeOrderNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/security/ldap/change-order][%d] changeOrderNoContent ", 204)
}

func (o *ChangeOrderNoContent) String() string {
	return fmt.Sprintf("[POST /v1/security/ldap/change-order][%d] changeOrderNoContent ", 204)
}

func (o *ChangeOrderNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeOrderUnauthorized creates a ChangeOrderUnauthorized with default headers values
func NewChangeOrderUnauthorized() *ChangeOrderUnauthorized {
	return &ChangeOrderUnauthorized{}
}

/* ChangeOrderUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type ChangeOrderUnauthorized struct {
}

// IsSuccess returns true when this change order unauthorized response has a 2xx status code
func (o *ChangeOrderUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this change order unauthorized response has a 3xx status code
func (o *ChangeOrderUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change order unauthorized response has a 4xx status code
func (o *ChangeOrderUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this change order unauthorized response has a 5xx status code
func (o *ChangeOrderUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this change order unauthorized response a status code equal to that given
func (o *ChangeOrderUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ChangeOrderUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/security/ldap/change-order][%d] changeOrderUnauthorized ", 401)
}

func (o *ChangeOrderUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/security/ldap/change-order][%d] changeOrderUnauthorized ", 401)
}

func (o *ChangeOrderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeOrderForbidden creates a ChangeOrderForbidden with default headers values
func NewChangeOrderForbidden() *ChangeOrderForbidden {
	return &ChangeOrderForbidden{}
}

/* ChangeOrderForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type ChangeOrderForbidden struct {
}

// IsSuccess returns true when this change order forbidden response has a 2xx status code
func (o *ChangeOrderForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this change order forbidden response has a 3xx status code
func (o *ChangeOrderForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change order forbidden response has a 4xx status code
func (o *ChangeOrderForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this change order forbidden response has a 5xx status code
func (o *ChangeOrderForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this change order forbidden response a status code equal to that given
func (o *ChangeOrderForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ChangeOrderForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/security/ldap/change-order][%d] changeOrderForbidden ", 403)
}

func (o *ChangeOrderForbidden) String() string {
	return fmt.Sprintf("[POST /v1/security/ldap/change-order][%d] changeOrderForbidden ", 403)
}

func (o *ChangeOrderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
