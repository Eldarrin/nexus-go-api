// Code generated by go-swagger; DO NOT EDIT.

package security_management_realms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"Eldarrin/nexus-go-api/models"
)

// GetRealmsReader is a Reader for the GetRealms structure.
type GetRealmsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRealmsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRealmsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRealmsOK creates a GetRealmsOK with default headers values
func NewGetRealmsOK() *GetRealmsOK {
	return &GetRealmsOK{}
}

/* GetRealmsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRealmsOK struct {
	Payload []*models.RealmAPIXO
}

// IsSuccess returns true when this get realms o k response has a 2xx status code
func (o *GetRealmsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get realms o k response has a 3xx status code
func (o *GetRealmsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get realms o k response has a 4xx status code
func (o *GetRealmsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get realms o k response has a 5xx status code
func (o *GetRealmsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get realms o k response a status code equal to that given
func (o *GetRealmsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRealmsOK) Error() string {
	return fmt.Sprintf("[GET /v1/security/realms/available][%d] getRealmsOK  %+v", 200, o.Payload)
}

func (o *GetRealmsOK) String() string {
	return fmt.Sprintf("[GET /v1/security/realms/available][%d] getRealmsOK  %+v", 200, o.Payload)
}

func (o *GetRealmsOK) GetPayload() []*models.RealmAPIXO {
	return o.Payload
}

func (o *GetRealmsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
