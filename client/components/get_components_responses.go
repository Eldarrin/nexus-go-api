// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Eldarrin/nexus-go-api/models"
)

// GetComponentsReader is a Reader for the GetComponents structure.
type GetComponentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComponentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComponentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetComponentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetComponentsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetComponentsOK creates a GetComponentsOK with default headers values
func NewGetComponentsOK() *GetComponentsOK {
	return &GetComponentsOK{}
}

/*
	GetComponentsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetComponentsOK struct {
	Payload *models.PageComponentXO
}

// IsSuccess returns true when this get components o k response has a 2xx status code
func (o *GetComponentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get components o k response has a 3xx status code
func (o *GetComponentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get components o k response has a 4xx status code
func (o *GetComponentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get components o k response has a 5xx status code
func (o *GetComponentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get components o k response a status code equal to that given
func (o *GetComponentsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetComponentsOK) Error() string {
	return fmt.Sprintf("[GET /v1/components][%d] getComponentsOK  %+v", 200, o.Payload)
}

func (o *GetComponentsOK) String() string {
	return fmt.Sprintf("[GET /v1/components][%d] getComponentsOK  %+v", 200, o.Payload)
}

func (o *GetComponentsOK) GetPayload() *models.PageComponentXO {
	return o.Payload
}

func (o *GetComponentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageComponentXO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComponentsForbidden creates a GetComponentsForbidden with default headers values
func NewGetComponentsForbidden() *GetComponentsForbidden {
	return &GetComponentsForbidden{}
}

/*
	GetComponentsForbidden describes a response with status code 403, with default header values.

Insufficient permissions to list components
*/
type GetComponentsForbidden struct {
}

// IsSuccess returns true when this get components forbidden response has a 2xx status code
func (o *GetComponentsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get components forbidden response has a 3xx status code
func (o *GetComponentsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get components forbidden response has a 4xx status code
func (o *GetComponentsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get components forbidden response has a 5xx status code
func (o *GetComponentsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get components forbidden response a status code equal to that given
func (o *GetComponentsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetComponentsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/components][%d] getComponentsForbidden ", 403)
}

func (o *GetComponentsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/components][%d] getComponentsForbidden ", 403)
}

func (o *GetComponentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetComponentsUnprocessableEntity creates a GetComponentsUnprocessableEntity with default headers values
func NewGetComponentsUnprocessableEntity() *GetComponentsUnprocessableEntity {
	return &GetComponentsUnprocessableEntity{}
}

/*
	GetComponentsUnprocessableEntity describes a response with status code 422, with default header values.

Parameter 'repository' is required
*/
type GetComponentsUnprocessableEntity struct {
}

// IsSuccess returns true when this get components unprocessable entity response has a 2xx status code
func (o *GetComponentsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get components unprocessable entity response has a 3xx status code
func (o *GetComponentsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get components unprocessable entity response has a 4xx status code
func (o *GetComponentsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get components unprocessable entity response has a 5xx status code
func (o *GetComponentsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get components unprocessable entity response a status code equal to that given
func (o *GetComponentsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *GetComponentsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /v1/components][%d] getComponentsUnprocessableEntity ", 422)
}

func (o *GetComponentsUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /v1/components][%d] getComponentsUnprocessableEntity ", 422)
}

func (o *GetComponentsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
