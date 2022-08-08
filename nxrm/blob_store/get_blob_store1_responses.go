// Code generated by go-swagger; DO NOT EDIT.

package blob_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"Eldarrin/nexus-go-api/models"
)

// GetBlobStore1Reader is a Reader for the GetBlobStore1 structure.
type GetBlobStore1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBlobStore1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBlobStore1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBlobStore1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetBlobStore1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBlobStore1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBlobStore1OK creates a GetBlobStore1OK with default headers values
func NewGetBlobStore1OK() *GetBlobStore1OK {
	return &GetBlobStore1OK{}
}

/* GetBlobStore1OK describes a response with status code 200, with default header values.

Success
*/
type GetBlobStore1OK struct {
	Payload *models.AzureBlobStoreAPIModel
}

// IsSuccess returns true when this get blob store1 o k response has a 2xx status code
func (o *GetBlobStore1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get blob store1 o k response has a 3xx status code
func (o *GetBlobStore1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get blob store1 o k response has a 4xx status code
func (o *GetBlobStore1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get blob store1 o k response has a 5xx status code
func (o *GetBlobStore1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get blob store1 o k response a status code equal to that given
func (o *GetBlobStore1OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetBlobStore1OK) Error() string {
	return fmt.Sprintf("[GET /v1/blobstores/azure/{name}][%d] getBlobStore1OK  %+v", 200, o.Payload)
}

func (o *GetBlobStore1OK) String() string {
	return fmt.Sprintf("[GET /v1/blobstores/azure/{name}][%d] getBlobStore1OK  %+v", 200, o.Payload)
}

func (o *GetBlobStore1OK) GetPayload() *models.AzureBlobStoreAPIModel {
	return o.Payload
}

func (o *GetBlobStore1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzureBlobStoreAPIModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBlobStore1BadRequest creates a GetBlobStore1BadRequest with default headers values
func NewGetBlobStore1BadRequest() *GetBlobStore1BadRequest {
	return &GetBlobStore1BadRequest{}
}

/* GetBlobStore1BadRequest describes a response with status code 400, with default header values.

Specified Azure blob store doesn't exist
*/
type GetBlobStore1BadRequest struct {
}

// IsSuccess returns true when this get blob store1 bad request response has a 2xx status code
func (o *GetBlobStore1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get blob store1 bad request response has a 3xx status code
func (o *GetBlobStore1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get blob store1 bad request response has a 4xx status code
func (o *GetBlobStore1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get blob store1 bad request response has a 5xx status code
func (o *GetBlobStore1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get blob store1 bad request response a status code equal to that given
func (o *GetBlobStore1BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetBlobStore1BadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/blobstores/azure/{name}][%d] getBlobStore1BadRequest ", 400)
}

func (o *GetBlobStore1BadRequest) String() string {
	return fmt.Sprintf("[GET /v1/blobstores/azure/{name}][%d] getBlobStore1BadRequest ", 400)
}

func (o *GetBlobStore1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBlobStore1Unauthorized creates a GetBlobStore1Unauthorized with default headers values
func NewGetBlobStore1Unauthorized() *GetBlobStore1Unauthorized {
	return &GetBlobStore1Unauthorized{}
}

/* GetBlobStore1Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type GetBlobStore1Unauthorized struct {
}

// IsSuccess returns true when this get blob store1 unauthorized response has a 2xx status code
func (o *GetBlobStore1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get blob store1 unauthorized response has a 3xx status code
func (o *GetBlobStore1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get blob store1 unauthorized response has a 4xx status code
func (o *GetBlobStore1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get blob store1 unauthorized response has a 5xx status code
func (o *GetBlobStore1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get blob store1 unauthorized response a status code equal to that given
func (o *GetBlobStore1Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetBlobStore1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/blobstores/azure/{name}][%d] getBlobStore1Unauthorized ", 401)
}

func (o *GetBlobStore1Unauthorized) String() string {
	return fmt.Sprintf("[GET /v1/blobstores/azure/{name}][%d] getBlobStore1Unauthorized ", 401)
}

func (o *GetBlobStore1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBlobStore1Forbidden creates a GetBlobStore1Forbidden with default headers values
func NewGetBlobStore1Forbidden() *GetBlobStore1Forbidden {
	return &GetBlobStore1Forbidden{}
}

/* GetBlobStore1Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type GetBlobStore1Forbidden struct {
}

// IsSuccess returns true when this get blob store1 forbidden response has a 2xx status code
func (o *GetBlobStore1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get blob store1 forbidden response has a 3xx status code
func (o *GetBlobStore1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get blob store1 forbidden response has a 4xx status code
func (o *GetBlobStore1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get blob store1 forbidden response has a 5xx status code
func (o *GetBlobStore1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get blob store1 forbidden response a status code equal to that given
func (o *GetBlobStore1Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetBlobStore1Forbidden) Error() string {
	return fmt.Sprintf("[GET /v1/blobstores/azure/{name}][%d] getBlobStore1Forbidden ", 403)
}

func (o *GetBlobStore1Forbidden) String() string {
	return fmt.Sprintf("[GET /v1/blobstores/azure/{name}][%d] getBlobStore1Forbidden ", 403)
}

func (o *GetBlobStore1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}