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

// GetRepositories1Reader is a Reader for the GetRepositories1 structure.
type GetRepositories1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositories1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositories1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepositories1OK creates a GetRepositories1OK with default headers values
func NewGetRepositories1OK() *GetRepositories1OK {
	return &GetRepositories1OK{}
}

/* GetRepositories1OK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepositories1OK struct {
	Payload []*models.RepositoryXO
}

// IsSuccess returns true when this get repositories1 o k response has a 2xx status code
func (o *GetRepositories1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repositories1 o k response has a 3xx status code
func (o *GetRepositories1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repositories1 o k response has a 4xx status code
func (o *GetRepositories1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repositories1 o k response has a 5xx status code
func (o *GetRepositories1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repositories1 o k response a status code equal to that given
func (o *GetRepositories1OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRepositories1OK) Error() string {
	return fmt.Sprintf("[GET /v1/repositories][%d] getRepositories1OK  %+v", 200, o.Payload)
}

func (o *GetRepositories1OK) String() string {
	return fmt.Sprintf("[GET /v1/repositories][%d] getRepositories1OK  %+v", 200, o.Payload)
}

func (o *GetRepositories1OK) GetPayload() []*models.RepositoryXO {
	return o.Payload
}

func (o *GetRepositories1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}