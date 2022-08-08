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

// GetRepository33Reader is a Reader for the GetRepository33 structure.
type GetRepository33Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepository33Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepository33OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepository33OK creates a GetRepository33OK with default headers values
func NewGetRepository33OK() *GetRepository33OK {
	return &GetRepository33OK{}
}

/* GetRepository33OK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepository33OK struct {
	Payload *models.SimpleAPIHostedRepository
}

// IsSuccess returns true when this get repository33 o k response has a 2xx status code
func (o *GetRepository33OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repository33 o k response has a 3xx status code
func (o *GetRepository33OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository33 o k response has a 4xx status code
func (o *GetRepository33OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repository33 o k response has a 5xx status code
func (o *GetRepository33OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repository33 o k response a status code equal to that given
func (o *GetRepository33OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRepository33OK) Error() string {
	return fmt.Sprintf("[GET /v1/repositories/r/hosted/{repositoryName}][%d] getRepository33OK  %+v", 200, o.Payload)
}

func (o *GetRepository33OK) String() string {
	return fmt.Sprintf("[GET /v1/repositories/r/hosted/{repositoryName}][%d] getRepository33OK  %+v", 200, o.Payload)
}

func (o *GetRepository33OK) GetPayload() *models.SimpleAPIHostedRepository {
	return o.Payload
}

func (o *GetRepository33OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIHostedRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}