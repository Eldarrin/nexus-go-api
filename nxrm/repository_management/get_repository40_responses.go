// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"nexus-go-api/models"
)

// GetRepository40Reader is a Reader for the GetRepository40 structure.
type GetRepository40Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepository40Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepository40OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepository40OK creates a GetRepository40OK with default headers values
func NewGetRepository40OK() *GetRepository40OK {
	return &GetRepository40OK{}
}

/* GetRepository40OK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepository40OK struct {
	Payload *models.SimpleAPIHostedRepository
}

// IsSuccess returns true when this get repository40 o k response has a 2xx status code
func (o *GetRepository40OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repository40 o k response has a 3xx status code
func (o *GetRepository40OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository40 o k response has a 4xx status code
func (o *GetRepository40OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repository40 o k response has a 5xx status code
func (o *GetRepository40OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repository40 o k response a status code equal to that given
func (o *GetRepository40OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRepository40OK) Error() string {
	return fmt.Sprintf("[GET /v1/repositories/bower/hosted/{repositoryName}][%d] getRepository40OK  %+v", 200, o.Payload)
}

func (o *GetRepository40OK) String() string {
	return fmt.Sprintf("[GET /v1/repositories/bower/hosted/{repositoryName}][%d] getRepository40OK  %+v", 200, o.Payload)
}

func (o *GetRepository40OK) GetPayload() *models.SimpleAPIHostedRepository {
	return o.Payload
}

func (o *GetRepository40OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIHostedRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
