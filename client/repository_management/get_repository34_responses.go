// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Eldarrin/nexus-go-api/models"
)

// GetRepository34Reader is a Reader for the GetRepository34 structure.
type GetRepository34Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepository34Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepository34OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepository34OK creates a GetRepository34OK with default headers values
func NewGetRepository34OK() *GetRepository34OK {
	return &GetRepository34OK{}
}

/* GetRepository34OK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepository34OK struct {
	Payload *models.SimpleAPIProxyRepository
}

// IsSuccess returns true when this get repository34 o k response has a 2xx status code
func (o *GetRepository34OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repository34 o k response has a 3xx status code
func (o *GetRepository34OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository34 o k response has a 4xx status code
func (o *GetRepository34OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repository34 o k response has a 5xx status code
func (o *GetRepository34OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repository34 o k response a status code equal to that given
func (o *GetRepository34OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRepository34OK) Error() string {
	return fmt.Sprintf("[GET /v1/repositories/r/proxy/{repositoryName}][%d] getRepository34OK  %+v", 200, o.Payload)
}

func (o *GetRepository34OK) String() string {
	return fmt.Sprintf("[GET /v1/repositories/r/proxy/{repositoryName}][%d] getRepository34OK  %+v", 200, o.Payload)
}

func (o *GetRepository34OK) GetPayload() *models.SimpleAPIProxyRepository {
	return o.Payload
}

func (o *GetRepository34OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIProxyRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
