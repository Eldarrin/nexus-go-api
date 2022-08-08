// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Eldarrin/nexus-go-api/models"
)

// GetPrivilegesReader is a Reader for the GetPrivileges structure.
type GetPrivilegesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivilegesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrivilegesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetPrivilegesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPrivilegesOK creates a GetPrivilegesOK with default headers values
func NewGetPrivilegesOK() *GetPrivilegesOK {
	return &GetPrivilegesOK{}
}

/* GetPrivilegesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetPrivilegesOK struct {
	Payload []*models.APIPrivilege
}

// IsSuccess returns true when this get privileges o k response has a 2xx status code
func (o *GetPrivilegesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get privileges o k response has a 3xx status code
func (o *GetPrivilegesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get privileges o k response has a 4xx status code
func (o *GetPrivilegesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get privileges o k response has a 5xx status code
func (o *GetPrivilegesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get privileges o k response a status code equal to that given
func (o *GetPrivilegesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPrivilegesOK) Error() string {
	return fmt.Sprintf("[GET /v1/security/privileges][%d] getPrivilegesOK  %+v", 200, o.Payload)
}

func (o *GetPrivilegesOK) String() string {
	return fmt.Sprintf("[GET /v1/security/privileges][%d] getPrivilegesOK  %+v", 200, o.Payload)
}

func (o *GetPrivilegesOK) GetPayload() []*models.APIPrivilege {
	return o.Payload
}

func (o *GetPrivilegesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrivilegesForbidden creates a GetPrivilegesForbidden with default headers values
func NewGetPrivilegesForbidden() *GetPrivilegesForbidden {
	return &GetPrivilegesForbidden{}
}

/* GetPrivilegesForbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the operation.
*/
type GetPrivilegesForbidden struct {
}

// IsSuccess returns true when this get privileges forbidden response has a 2xx status code
func (o *GetPrivilegesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get privileges forbidden response has a 3xx status code
func (o *GetPrivilegesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get privileges forbidden response has a 4xx status code
func (o *GetPrivilegesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get privileges forbidden response has a 5xx status code
func (o *GetPrivilegesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get privileges forbidden response a status code equal to that given
func (o *GetPrivilegesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetPrivilegesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/security/privileges][%d] getPrivilegesForbidden ", 403)
}

func (o *GetPrivilegesForbidden) String() string {
	return fmt.Sprintf("[GET /v1/security/privileges][%d] getPrivilegesForbidden ", 403)
}

func (o *GetPrivilegesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
