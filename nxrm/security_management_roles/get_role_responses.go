// Code generated by go-swagger; DO NOT EDIT.

package security_management_roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"Eldarrin/nexus-go-api/models"
)

// GetRoleReader is a Reader for the GetRole structure.
type GetRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoleOK creates a GetRoleOK with default headers values
func NewGetRoleOK() *GetRoleOK {
	return &GetRoleOK{}
}

/* GetRoleOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoleOK struct {
	Payload *models.RoleXOResponse
}

// IsSuccess returns true when this get role o k response has a 2xx status code
func (o *GetRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get role o k response has a 3xx status code
func (o *GetRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get role o k response has a 4xx status code
func (o *GetRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get role o k response has a 5xx status code
func (o *GetRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get role o k response a status code equal to that given
func (o *GetRoleOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRoleOK) Error() string {
	return fmt.Sprintf("[GET /v1/security/roles/{id}][%d] getRoleOK  %+v", 200, o.Payload)
}

func (o *GetRoleOK) String() string {
	return fmt.Sprintf("[GET /v1/security/roles/{id}][%d] getRoleOK  %+v", 200, o.Payload)
}

func (o *GetRoleOK) GetPayload() *models.RoleXOResponse {
	return o.Payload
}

func (o *GetRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleXOResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleBadRequest creates a GetRoleBadRequest with default headers values
func NewGetRoleBadRequest() *GetRoleBadRequest {
	return &GetRoleBadRequest{}
}

/* GetRoleBadRequest describes a response with status code 400, with default header values.

The specified source does not exist
*/
type GetRoleBadRequest struct {
}

// IsSuccess returns true when this get role bad request response has a 2xx status code
func (o *GetRoleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get role bad request response has a 3xx status code
func (o *GetRoleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get role bad request response has a 4xx status code
func (o *GetRoleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get role bad request response has a 5xx status code
func (o *GetRoleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get role bad request response a status code equal to that given
func (o *GetRoleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRoleBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/security/roles/{id}][%d] getRoleBadRequest ", 400)
}

func (o *GetRoleBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/security/roles/{id}][%d] getRoleBadRequest ", 400)
}

func (o *GetRoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoleForbidden creates a GetRoleForbidden with default headers values
func NewGetRoleForbidden() *GetRoleForbidden {
	return &GetRoleForbidden{}
}

/* GetRoleForbidden describes a response with status code 403, with default header values.

Insufficient permissions to read roles
*/
type GetRoleForbidden struct {
}

// IsSuccess returns true when this get role forbidden response has a 2xx status code
func (o *GetRoleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get role forbidden response has a 3xx status code
func (o *GetRoleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get role forbidden response has a 4xx status code
func (o *GetRoleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get role forbidden response has a 5xx status code
func (o *GetRoleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get role forbidden response a status code equal to that given
func (o *GetRoleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRoleForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/security/roles/{id}][%d] getRoleForbidden ", 403)
}

func (o *GetRoleForbidden) String() string {
	return fmt.Sprintf("[GET /v1/security/roles/{id}][%d] getRoleForbidden ", 403)
}

func (o *GetRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoleNotFound creates a GetRoleNotFound with default headers values
func NewGetRoleNotFound() *GetRoleNotFound {
	return &GetRoleNotFound{}
}

/* GetRoleNotFound describes a response with status code 404, with default header values.

Role not found
*/
type GetRoleNotFound struct {
}

// IsSuccess returns true when this get role not found response has a 2xx status code
func (o *GetRoleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get role not found response has a 3xx status code
func (o *GetRoleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get role not found response has a 4xx status code
func (o *GetRoleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get role not found response has a 5xx status code
func (o *GetRoleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get role not found response a status code equal to that given
func (o *GetRoleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRoleNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/security/roles/{id}][%d] getRoleNotFound ", 404)
}

func (o *GetRoleNotFound) String() string {
	return fmt.Sprintf("[GET /v1/security/roles/{id}][%d] getRoleNotFound ", 404)
}

func (o *GetRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}