// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreatePrivilege2Reader is a Reader for the CreatePrivilege2 structure.
type CreatePrivilege2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePrivilege2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewCreatePrivilege2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreatePrivilege2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreatePrivilege2BadRequest creates a CreatePrivilege2BadRequest with default headers values
func NewCreatePrivilege2BadRequest() *CreatePrivilege2BadRequest {
	return &CreatePrivilege2BadRequest{}
}

/* CreatePrivilege2BadRequest describes a response with status code 400, with default header values.

Privilege object not configured properly.
*/
type CreatePrivilege2BadRequest struct {
}

// IsSuccess returns true when this create privilege2 bad request response has a 2xx status code
func (o *CreatePrivilege2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create privilege2 bad request response has a 3xx status code
func (o *CreatePrivilege2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create privilege2 bad request response has a 4xx status code
func (o *CreatePrivilege2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create privilege2 bad request response has a 5xx status code
func (o *CreatePrivilege2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create privilege2 bad request response a status code equal to that given
func (o *CreatePrivilege2BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreatePrivilege2BadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/security/privileges/repository-content-selector][%d] createPrivilege2BadRequest ", 400)
}

func (o *CreatePrivilege2BadRequest) String() string {
	return fmt.Sprintf("[POST /v1/security/privileges/repository-content-selector][%d] createPrivilege2BadRequest ", 400)
}

func (o *CreatePrivilege2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePrivilege2Forbidden creates a CreatePrivilege2Forbidden with default headers values
func NewCreatePrivilege2Forbidden() *CreatePrivilege2Forbidden {
	return &CreatePrivilege2Forbidden{}
}

/* CreatePrivilege2Forbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the operation.
*/
type CreatePrivilege2Forbidden struct {
}

// IsSuccess returns true when this create privilege2 forbidden response has a 2xx status code
func (o *CreatePrivilege2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create privilege2 forbidden response has a 3xx status code
func (o *CreatePrivilege2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create privilege2 forbidden response has a 4xx status code
func (o *CreatePrivilege2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create privilege2 forbidden response has a 5xx status code
func (o *CreatePrivilege2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create privilege2 forbidden response a status code equal to that given
func (o *CreatePrivilege2Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreatePrivilege2Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/security/privileges/repository-content-selector][%d] createPrivilege2Forbidden ", 403)
}

func (o *CreatePrivilege2Forbidden) String() string {
	return fmt.Sprintf("[POST /v1/security/privileges/repository-content-selector][%d] createPrivilege2Forbidden ", 403)
}

func (o *CreatePrivilege2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
