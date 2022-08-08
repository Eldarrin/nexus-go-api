// Code generated by go-swagger; DO NOT EDIT.

package support

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"Eldarrin/nexus-go-api/models"
)

// SupportzippathReader is a Reader for the Supportzippath structure.
type SupportzippathReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SupportzippathReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSupportzippathOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSupportzippathOK creates a SupportzippathOK with default headers values
func NewSupportzippathOK() *SupportzippathOK {
	return &SupportzippathOK{}
}

/* SupportzippathOK describes a response with status code 200, with default header values.

successful operation
*/
type SupportzippathOK struct {
	Payload *models.SupportZipXO
}

// IsSuccess returns true when this supportzippath o k response has a 2xx status code
func (o *SupportzippathOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this supportzippath o k response has a 3xx status code
func (o *SupportzippathOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this supportzippath o k response has a 4xx status code
func (o *SupportzippathOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this supportzippath o k response has a 5xx status code
func (o *SupportzippathOK) IsServerError() bool {
	return false
}

// IsCode returns true when this supportzippath o k response a status code equal to that given
func (o *SupportzippathOK) IsCode(code int) bool {
	return code == 200
}

func (o *SupportzippathOK) Error() string {
	return fmt.Sprintf("[POST /v1/support/supportzippath][%d] supportzippathOK  %+v", 200, o.Payload)
}

func (o *SupportzippathOK) String() string {
	return fmt.Sprintf("[POST /v1/support/supportzippath][%d] supportzippathOK  %+v", 200, o.Payload)
}

func (o *SupportzippathOK) GetPayload() *models.SupportZipXO {
	return o.Payload
}

func (o *SupportzippathOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SupportZipXO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
