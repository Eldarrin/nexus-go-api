// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdatePrivilege4Reader is a Reader for the UpdatePrivilege4 structure.
type UpdatePrivilege4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePrivilege4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewUpdatePrivilege4BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdatePrivilege4Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdatePrivilege4NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdatePrivilege4BadRequest creates a UpdatePrivilege4BadRequest with default headers values
func NewUpdatePrivilege4BadRequest() *UpdatePrivilege4BadRequest {
	return &UpdatePrivilege4BadRequest{}
}

/* UpdatePrivilege4BadRequest describes a response with status code 400, with default header values.

Privilege object not configured properly.
*/
type UpdatePrivilege4BadRequest struct {
}

// IsSuccess returns true when this update privilege4 bad request response has a 2xx status code
func (o *UpdatePrivilege4BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update privilege4 bad request response has a 3xx status code
func (o *UpdatePrivilege4BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update privilege4 bad request response has a 4xx status code
func (o *UpdatePrivilege4BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update privilege4 bad request response has a 5xx status code
func (o *UpdatePrivilege4BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update privilege4 bad request response a status code equal to that given
func (o *UpdatePrivilege4BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdatePrivilege4BadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/repository-admin/{privilegeName}][%d] updatePrivilege4BadRequest ", 400)
}

func (o *UpdatePrivilege4BadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/repository-admin/{privilegeName}][%d] updatePrivilege4BadRequest ", 400)
}

func (o *UpdatePrivilege4BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePrivilege4Forbidden creates a UpdatePrivilege4Forbidden with default headers values
func NewUpdatePrivilege4Forbidden() *UpdatePrivilege4Forbidden {
	return &UpdatePrivilege4Forbidden{}
}

/* UpdatePrivilege4Forbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the operation.
*/
type UpdatePrivilege4Forbidden struct {
}

// IsSuccess returns true when this update privilege4 forbidden response has a 2xx status code
func (o *UpdatePrivilege4Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update privilege4 forbidden response has a 3xx status code
func (o *UpdatePrivilege4Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update privilege4 forbidden response has a 4xx status code
func (o *UpdatePrivilege4Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update privilege4 forbidden response has a 5xx status code
func (o *UpdatePrivilege4Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update privilege4 forbidden response a status code equal to that given
func (o *UpdatePrivilege4Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdatePrivilege4Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/repository-admin/{privilegeName}][%d] updatePrivilege4Forbidden ", 403)
}

func (o *UpdatePrivilege4Forbidden) String() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/repository-admin/{privilegeName}][%d] updatePrivilege4Forbidden ", 403)
}

func (o *UpdatePrivilege4Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePrivilege4NotFound creates a UpdatePrivilege4NotFound with default headers values
func NewUpdatePrivilege4NotFound() *UpdatePrivilege4NotFound {
	return &UpdatePrivilege4NotFound{}
}

/* UpdatePrivilege4NotFound describes a response with status code 404, with default header values.

Privilege not found in the system.
*/
type UpdatePrivilege4NotFound struct {
}

// IsSuccess returns true when this update privilege4 not found response has a 2xx status code
func (o *UpdatePrivilege4NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update privilege4 not found response has a 3xx status code
func (o *UpdatePrivilege4NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update privilege4 not found response has a 4xx status code
func (o *UpdatePrivilege4NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update privilege4 not found response has a 5xx status code
func (o *UpdatePrivilege4NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update privilege4 not found response a status code equal to that given
func (o *UpdatePrivilege4NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdatePrivilege4NotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/repository-admin/{privilegeName}][%d] updatePrivilege4NotFound ", 404)
}

func (o *UpdatePrivilege4NotFound) String() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/repository-admin/{privilegeName}][%d] updatePrivilege4NotFound ", 404)
}

func (o *UpdatePrivilege4NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
