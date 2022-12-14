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

// GetYumProxyRepositoryReader is a Reader for the GetYumProxyRepository structure.
type GetYumProxyRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetYumProxyRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetYumProxyRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetYumProxyRepositoryOK creates a GetYumProxyRepositoryOK with default headers values
func NewGetYumProxyRepositoryOK() *GetYumProxyRepositoryOK {
	return &GetYumProxyRepositoryOK{}
}

/*
	GetYumProxyRepositoryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetYumProxyRepositoryOK struct {
	Payload *models.SimpleAPIProxyRepository
}

// IsSuccess returns true when this get repository23 o k response has a 2xx status code
func (o *GetYumProxyRepositoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repository23 o k response has a 3xx status code
func (o *GetYumProxyRepositoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository23 o k response has a 4xx status code
func (o *GetYumProxyRepositoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repository23 o k response has a 5xx status code
func (o *GetYumProxyRepositoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repository23 o k response a status code equal to that given
func (o *GetYumProxyRepositoryOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetYumProxyRepositoryOK) Error() string {
	return fmt.Sprintf("[GET /v1/repositories/yum/proxy/{repositoryName}][%d] getRepository23OK  %+v", 200, o.Payload)
}

func (o *GetYumProxyRepositoryOK) String() string {
	return fmt.Sprintf("[GET /v1/repositories/yum/proxy/{repositoryName}][%d] getRepository23OK  %+v", 200, o.Payload)
}

func (o *GetYumProxyRepositoryOK) GetPayload() *models.SimpleAPIProxyRepository {
	return o.Payload
}

func (o *GetYumProxyRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIProxyRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
