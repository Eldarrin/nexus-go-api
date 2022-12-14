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

// GetMavenProxyRepositoryReader is a Reader for the GetMavenProxyRepository structure.
type GetMavenProxyRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMavenProxyRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMavenProxyRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMavenProxyRepositoryOK creates a GetMavenProxyRepositoryOK with default headers values
func NewGetMavenProxyRepositoryOK() *GetMavenProxyRepositoryOK {
	return &GetMavenProxyRepositoryOK{}
}

/*
	GetMavenProxyRepositoryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetMavenProxyRepositoryOK struct {
	Payload *models.MavenProxyAPIRepository
}

// IsSuccess returns true when this get repository3 o k response has a 2xx status code
func (o *GetMavenProxyRepositoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repository3 o k response has a 3xx status code
func (o *GetMavenProxyRepositoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository3 o k response has a 4xx status code
func (o *GetMavenProxyRepositoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repository3 o k response has a 5xx status code
func (o *GetMavenProxyRepositoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repository3 o k response a status code equal to that given
func (o *GetMavenProxyRepositoryOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetMavenProxyRepositoryOK) Error() string {
	return fmt.Sprintf("[GET /v1/repositories/maven/proxy/{repositoryName}][%d] getRepository3OK  %+v", 200, o.Payload)
}

func (o *GetMavenProxyRepositoryOK) String() string {
	return fmt.Sprintf("[GET /v1/repositories/maven/proxy/{repositoryName}][%d] getRepository3OK  %+v", 200, o.Payload)
}

func (o *GetMavenProxyRepositoryOK) GetPayload() *models.MavenProxyAPIRepository {
	return o.Payload
}

func (o *GetMavenProxyRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MavenProxyAPIRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
