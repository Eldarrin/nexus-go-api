// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreatePrivilege3Reader is a Reader for the CreatePrivilege3 structure.
type CreatePrivilege3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePrivilege3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewCreatePrivilege3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreatePrivilege3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreatePrivilege3BadRequest creates a CreatePrivilege3BadRequest with default headers values
func NewCreatePrivilege3BadRequest() *CreatePrivilege3BadRequest {
	return &CreatePrivilege3BadRequest{}
}

/* CreatePrivilege3BadRequest describes a response with status code 400, with default header values.

Privilege object not configured properly.
*/
type CreatePrivilege3BadRequest struct {
}

// IsSuccess returns true when this create privilege3 bad request response has a 2xx status code
func (o *CreatePrivilege3BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create privilege3 bad request response has a 3xx status code
func (o *CreatePrivilege3BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create privilege3 bad request response has a 4xx status code
func (o *CreatePrivilege3BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create privilege3 bad request response has a 5xx status code
func (o *CreatePrivilege3BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create privilege3 bad request response a status code equal to that given
func (o *CreatePrivilege3BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreatePrivilege3BadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/security/privileges/repository-admin][%d] createPrivilege3BadRequest ", 400)
}

func (o *CreatePrivilege3BadRequest) String() string {
	return fmt.Sprintf("[POST /v1/security/privileges/repository-admin][%d] createPrivilege3BadRequest ", 400)
}

func (o *CreatePrivilege3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePrivilege3Forbidden creates a CreatePrivilege3Forbidden with default headers values
func NewCreatePrivilege3Forbidden() *CreatePrivilege3Forbidden {
	return &CreatePrivilege3Forbidden{}
}

/* CreatePrivilege3Forbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the operation.
*/
type CreatePrivilege3Forbidden struct {
}

// IsSuccess returns true when this create privilege3 forbidden response has a 2xx status code
func (o *CreatePrivilege3Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create privilege3 forbidden response has a 3xx status code
func (o *CreatePrivilege3Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create privilege3 forbidden response has a 4xx status code
func (o *CreatePrivilege3Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create privilege3 forbidden response has a 5xx status code
func (o *CreatePrivilege3Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create privilege3 forbidden response a status code equal to that given
func (o *CreatePrivilege3Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreatePrivilege3Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/security/privileges/repository-admin][%d] createPrivilege3Forbidden ", 403)
}

func (o *CreatePrivilege3Forbidden) String() string {
	return fmt.Sprintf("[POST /v1/security/privileges/repository-admin][%d] createPrivilege3Forbidden ", 403)
}

func (o *CreatePrivilege3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}