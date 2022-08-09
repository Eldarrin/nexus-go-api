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

// GetGitLfsHostedRepositoryReader is a Reader for the GetGitLfsHostedRepository structure.
type GetGitLfsHostedRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGitLfsHostedRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGitLfsHostedRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGitLfsHostedRepositoryOK creates a GetGitLfsHostedRepositoryOK with default headers values
func NewGetGitLfsHostedRepositoryOK() *GetGitLfsHostedRepositoryOK {
	return &GetGitLfsHostedRepositoryOK{}
}

/*
	GetGitLfsHostedRepositoryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetGitLfsHostedRepositoryOK struct {
	Payload *models.SimpleAPIHostedRepository
}

// IsSuccess returns true when this get repository26 o k response has a 2xx status code
func (o *GetGitLfsHostedRepositoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repository26 o k response has a 3xx status code
func (o *GetGitLfsHostedRepositoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository26 o k response has a 4xx status code
func (o *GetGitLfsHostedRepositoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repository26 o k response has a 5xx status code
func (o *GetGitLfsHostedRepositoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repository26 o k response a status code equal to that given
func (o *GetGitLfsHostedRepositoryOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetGitLfsHostedRepositoryOK) Error() string {
	return fmt.Sprintf("[GET /v1/repositories/gitlfs/hosted/{repositoryName}][%d] getRepository26OK  %+v", 200, o.Payload)
}

func (o *GetGitLfsHostedRepositoryOK) String() string {
	return fmt.Sprintf("[GET /v1/repositories/gitlfs/hosted/{repositoryName}][%d] getRepository26OK  %+v", 200, o.Payload)
}

func (o *GetGitLfsHostedRepositoryOK) GetPayload() *models.SimpleAPIHostedRepository {
	return o.Payload
}

func (o *GetGitLfsHostedRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIHostedRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
