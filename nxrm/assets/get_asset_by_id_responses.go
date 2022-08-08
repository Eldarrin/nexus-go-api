// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"nexus-go-api/models"
)

// GetAssetByIDReader is a Reader for the GetAssetByID structure.
type GetAssetByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAssetByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAssetByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetAssetByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAssetByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetAssetByIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAssetByIDOK creates a GetAssetByIDOK with default headers values
func NewGetAssetByIDOK() *GetAssetByIDOK {
	return &GetAssetByIDOK{}
}

/* GetAssetByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAssetByIDOK struct {
	Payload *models.AssetXO
}

// IsSuccess returns true when this get asset by Id o k response has a 2xx status code
func (o *GetAssetByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get asset by Id o k response has a 3xx status code
func (o *GetAssetByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get asset by Id o k response has a 4xx status code
func (o *GetAssetByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get asset by Id o k response has a 5xx status code
func (o *GetAssetByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get asset by Id o k response a status code equal to that given
func (o *GetAssetByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAssetByIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/assets/{id}][%d] getAssetByIdOK  %+v", 200, o.Payload)
}

func (o *GetAssetByIDOK) String() string {
	return fmt.Sprintf("[GET /v1/assets/{id}][%d] getAssetByIdOK  %+v", 200, o.Payload)
}

func (o *GetAssetByIDOK) GetPayload() *models.AssetXO {
	return o.Payload
}

func (o *GetAssetByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssetXO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssetByIDForbidden creates a GetAssetByIDForbidden with default headers values
func NewGetAssetByIDForbidden() *GetAssetByIDForbidden {
	return &GetAssetByIDForbidden{}
}

/* GetAssetByIDForbidden describes a response with status code 403, with default header values.

Insufficient permissions to get asset
*/
type GetAssetByIDForbidden struct {
}

// IsSuccess returns true when this get asset by Id forbidden response has a 2xx status code
func (o *GetAssetByIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get asset by Id forbidden response has a 3xx status code
func (o *GetAssetByIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get asset by Id forbidden response has a 4xx status code
func (o *GetAssetByIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get asset by Id forbidden response has a 5xx status code
func (o *GetAssetByIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get asset by Id forbidden response a status code equal to that given
func (o *GetAssetByIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAssetByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/assets/{id}][%d] getAssetByIdForbidden ", 403)
}

func (o *GetAssetByIDForbidden) String() string {
	return fmt.Sprintf("[GET /v1/assets/{id}][%d] getAssetByIdForbidden ", 403)
}

func (o *GetAssetByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAssetByIDNotFound creates a GetAssetByIDNotFound with default headers values
func NewGetAssetByIDNotFound() *GetAssetByIDNotFound {
	return &GetAssetByIDNotFound{}
}

/* GetAssetByIDNotFound describes a response with status code 404, with default header values.

Asset not found
*/
type GetAssetByIDNotFound struct {
}

// IsSuccess returns true when this get asset by Id not found response has a 2xx status code
func (o *GetAssetByIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get asset by Id not found response has a 3xx status code
func (o *GetAssetByIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get asset by Id not found response has a 4xx status code
func (o *GetAssetByIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get asset by Id not found response has a 5xx status code
func (o *GetAssetByIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get asset by Id not found response a status code equal to that given
func (o *GetAssetByIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAssetByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/assets/{id}][%d] getAssetByIdNotFound ", 404)
}

func (o *GetAssetByIDNotFound) String() string {
	return fmt.Sprintf("[GET /v1/assets/{id}][%d] getAssetByIdNotFound ", 404)
}

func (o *GetAssetByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAssetByIDUnprocessableEntity creates a GetAssetByIDUnprocessableEntity with default headers values
func NewGetAssetByIDUnprocessableEntity() *GetAssetByIDUnprocessableEntity {
	return &GetAssetByIDUnprocessableEntity{}
}

/* GetAssetByIDUnprocessableEntity describes a response with status code 422, with default header values.

Malformed ID
*/
type GetAssetByIDUnprocessableEntity struct {
}

// IsSuccess returns true when this get asset by Id unprocessable entity response has a 2xx status code
func (o *GetAssetByIDUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get asset by Id unprocessable entity response has a 3xx status code
func (o *GetAssetByIDUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get asset by Id unprocessable entity response has a 4xx status code
func (o *GetAssetByIDUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get asset by Id unprocessable entity response has a 5xx status code
func (o *GetAssetByIDUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get asset by Id unprocessable entity response a status code equal to that given
func (o *GetAssetByIDUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *GetAssetByIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /v1/assets/{id}][%d] getAssetByIdUnprocessableEntity ", 422)
}

func (o *GetAssetByIDUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /v1/assets/{id}][%d] getAssetByIdUnprocessableEntity ", 422)
}

func (o *GetAssetByIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
