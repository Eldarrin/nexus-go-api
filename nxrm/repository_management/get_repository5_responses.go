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

// GetRepository5Reader is a Reader for the GetRepository5 structure.
type GetRepository5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepository5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepository5OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepository5OK creates a GetRepository5OK with default headers values
func NewGetRepository5OK() *GetRepository5OK {
	return &GetRepository5OK{}
}

/* GetRepository5OK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepository5OK struct {
	Payload *models.AptProxyAPIRepository
}

// IsSuccess returns true when this get repository5 o k response has a 2xx status code
func (o *GetRepository5OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repository5 o k response has a 3xx status code
func (o *GetRepository5OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository5 o k response has a 4xx status code
func (o *GetRepository5OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repository5 o k response has a 5xx status code
func (o *GetRepository5OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repository5 o k response a status code equal to that given
func (o *GetRepository5OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRepository5OK) Error() string {
	return fmt.Sprintf("[GET /v1/repositories/apt/proxy/{repositoryName}][%d] getRepository5OK  %+v", 200, o.Payload)
}

func (o *GetRepository5OK) String() string {
	return fmt.Sprintf("[GET /v1/repositories/apt/proxy/{repositoryName}][%d] getRepository5OK  %+v", 200, o.Payload)
}

func (o *GetRepository5OK) GetPayload() *models.AptProxyAPIRepository {
	return o.Payload
}

func (o *GetRepository5OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AptProxyAPIRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
