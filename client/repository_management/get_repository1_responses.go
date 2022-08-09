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

// GetMavenGroupRepositoryReader is a Reader for the GetMavenGroupRepository structure.
type GetMavenGroupRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMavenGroupRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMavenGroupRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMavenGroupRepositoryOK creates a GetMavenGroupRepositoryOK with default headers values
func NewGetMavenGroupRepositoryOK() *GetMavenGroupRepositoryOK {
	return &GetMavenGroupRepositoryOK{}
}

/*
	GetMavenGroupRepositoryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetMavenGroupRepositoryOK struct {
	Payload *models.SimpleAPIGroupRepository
}

// IsSuccess returns true when this get repository1 o k response has a 2xx status code
func (o *GetMavenGroupRepositoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repository1 o k response has a 3xx status code
func (o *GetMavenGroupRepositoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository1 o k response has a 4xx status code
func (o *GetMavenGroupRepositoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repository1 o k response has a 5xx status code
func (o *GetMavenGroupRepositoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repository1 o k response a status code equal to that given
func (o *GetMavenGroupRepositoryOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetMavenGroupRepositoryOK) Error() string {
	return fmt.Sprintf("[GET /v1/repositories/maven/group/{repositoryName}][%d] getRepository1OK  %+v", 200, o.Payload)
}

func (o *GetMavenGroupRepositoryOK) String() string {
	return fmt.Sprintf("[GET /v1/repositories/maven/group/{repositoryName}][%d] getRepository1OK  %+v", 200, o.Payload)
}

func (o *GetMavenGroupRepositoryOK) GetPayload() *models.SimpleAPIGroupRepository {
	return o.Payload
}

func (o *GetMavenGroupRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIGroupRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
