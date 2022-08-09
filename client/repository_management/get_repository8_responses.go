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

// GetRawProxyRepositoryReader is a Reader for the GetRawProxyRepository structure.
type GetRawProxyRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRawProxyRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRawProxyRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRawProxyRepositoryOK creates a GetRawProxyRepositoryOK with default headers values
func NewGetRawProxyRepositoryOK() *GetRawProxyRepositoryOK {
	return &GetRawProxyRepositoryOK{}
}

/*
	GetRawProxyRepositoryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRawProxyRepositoryOK struct {
	Payload *models.SimpleAPIProxyRepository
}

// IsSuccess returns true when this get repository8 o k response has a 2xx status code
func (o *GetRawProxyRepositoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repository8 o k response has a 3xx status code
func (o *GetRawProxyRepositoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository8 o k response has a 4xx status code
func (o *GetRawProxyRepositoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repository8 o k response has a 5xx status code
func (o *GetRawProxyRepositoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repository8 o k response a status code equal to that given
func (o *GetRawProxyRepositoryOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRawProxyRepositoryOK) Error() string {
	return fmt.Sprintf("[GET /v1/repositories/raw/proxy/{repositoryName}][%d] getRepository8OK  %+v", 200, o.Payload)
}

func (o *GetRawProxyRepositoryOK) String() string {
	return fmt.Sprintf("[GET /v1/repositories/raw/proxy/{repositoryName}][%d] getRepository8OK  %+v", 200, o.Payload)
}

func (o *GetRawProxyRepositoryOK) GetPayload() *models.SimpleAPIProxyRepository {
	return o.Payload
}

func (o *GetRawProxyRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIProxyRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
