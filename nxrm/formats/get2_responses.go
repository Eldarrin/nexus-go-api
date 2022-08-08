// Code generated by go-swagger; DO NOT EDIT.

package formats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"Eldarrin/nexus-go-api/models"
)

// Get2Reader is a Reader for the Get2 structure.
type Get2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Get2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGet2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGet2OK creates a Get2OK with default headers values
func NewGet2OK() *Get2OK {
	return &Get2OK{}
}

/* Get2OK describes a response with status code 200, with default header values.

successful operation
*/
type Get2OK struct {
	Payload []*models.UploadDefinitionXO
}

// IsSuccess returns true when this get2 o k response has a 2xx status code
func (o *Get2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get2 o k response has a 3xx status code
func (o *Get2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get2 o k response has a 4xx status code
func (o *Get2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get2 o k response has a 5xx status code
func (o *Get2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get2 o k response a status code equal to that given
func (o *Get2OK) IsCode(code int) bool {
	return code == 200
}

func (o *Get2OK) Error() string {
	return fmt.Sprintf("[GET /v1/formats/upload-specs][%d] get2OK  %+v", 200, o.Payload)
}

func (o *Get2OK) String() string {
	return fmt.Sprintf("[GET /v1/formats/upload-specs][%d] get2OK  %+v", 200, o.Payload)
}

func (o *Get2OK) GetPayload() []*models.UploadDefinitionXO {
	return o.Payload
}

func (o *Get2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
