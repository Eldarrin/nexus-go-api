// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"Eldarrin/nexus-go-api/models"
)

// GetRepository25Reader is a Reader for the GetRepository25 structure.
type GetRepository25Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepository25Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepository25OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepository25OK creates a GetRepository25OK with default headers values
func NewGetRepository25OK() *GetRepository25OK {
	return &GetRepository25OK{}
}

/* GetRepository25OK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepository25OK struct {
	Payload *models.SimpleAPIProxyRepository
}

// IsSuccess returns true when this get repository25 o k response has a 2xx status code
func (o *GetRepository25OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repository25 o k response has a 3xx status code
func (o *GetRepository25OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository25 o k response has a 4xx status code
func (o *GetRepository25OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repository25 o k response has a 5xx status code
func (o *GetRepository25OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repository25 o k response a status code equal to that given
func (o *GetRepository25OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRepository25OK) Error() string {
	return fmt.Sprintf("[GET /v1/repositories/helm/proxy/{repositoryName}][%d] getRepository25OK  %+v", 200, o.Payload)
}

func (o *GetRepository25OK) String() string {
	return fmt.Sprintf("[GET /v1/repositories/helm/proxy/{repositoryName}][%d] getRepository25OK  %+v", 200, o.Payload)
}

func (o *GetRepository25OK) GetPayload() *models.SimpleAPIProxyRepository {
	return o.Payload
}

func (o *GetRepository25OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIProxyRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
