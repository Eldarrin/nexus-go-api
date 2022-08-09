// Code generated by go-swagger; DO NOT EDIT.

package script

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Eldarrin/nexus-go-api/models"
)

// Read1Reader is a Reader for the Read1 structure.
type Read1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Read1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRead1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRead1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRead1OK creates a Read1OK with default headers values
func NewRead1OK() *Read1OK {
	return &Read1OK{}
}

/*
	Read1OK describes a response with status code 200, with default header values.

successful operation
*/
type Read1OK struct {
	Payload *models.ScriptXO
}

// IsSuccess returns true when this read1 o k response has a 2xx status code
func (o *Read1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read1 o k response has a 3xx status code
func (o *Read1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read1 o k response has a 4xx status code
func (o *Read1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read1 o k response has a 5xx status code
func (o *Read1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this read1 o k response a status code equal to that given
func (o *Read1OK) IsCode(code int) bool {
	return code == 200
}

func (o *Read1OK) Error() string {
	return fmt.Sprintf("[GET /v1/script/{name}][%d] read1OK  %+v", 200, o.Payload)
}

func (o *Read1OK) String() string {
	return fmt.Sprintf("[GET /v1/script/{name}][%d] read1OK  %+v", 200, o.Payload)
}

func (o *Read1OK) GetPayload() *models.ScriptXO {
	return o.Payload
}

func (o *Read1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScriptXO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRead1NotFound creates a Read1NotFound with default headers values
func NewRead1NotFound() *Read1NotFound {
	return &Read1NotFound{}
}

/*
	Read1NotFound describes a response with status code 404, with default header values.

No script with the specified name
*/
type Read1NotFound struct {
}

// IsSuccess returns true when this read1 not found response has a 2xx status code
func (o *Read1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read1 not found response has a 3xx status code
func (o *Read1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read1 not found response has a 4xx status code
func (o *Read1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this read1 not found response has a 5xx status code
func (o *Read1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this read1 not found response a status code equal to that given
func (o *Read1NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *Read1NotFound) Error() string {
	return fmt.Sprintf("[GET /v1/script/{name}][%d] read1NotFound ", 404)
}

func (o *Read1NotFound) String() string {
	return fmt.Sprintf("[GET /v1/script/{name}][%d] read1NotFound ", 404)
}

func (o *Read1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
