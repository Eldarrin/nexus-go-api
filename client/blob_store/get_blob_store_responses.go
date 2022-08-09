// Code generated by go-swagger; DO NOT EDIT.

package blob_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Eldarrin/nexus-go-api/models"
)

// GetBlobStoreReader is a Reader for the GetBlobStore structure.
type GetBlobStoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBlobStoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBlobStoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBlobStoreBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetBlobStoreUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBlobStoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBlobStoreOK creates a GetBlobStoreOK with default headers values
func NewGetBlobStoreOK() *GetBlobStoreOK {
	return &GetBlobStoreOK{}
}

/*
	GetBlobStoreOK describes a response with status code 200, with default header values.

Success
*/
type GetBlobStoreOK struct {
	Payload *models.S3BlobStoreAPIModel
}

// IsSuccess returns true when this get blob store o k response has a 2xx status code
func (o *GetBlobStoreOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get blob store o k response has a 3xx status code
func (o *GetBlobStoreOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get blob store o k response has a 4xx status code
func (o *GetBlobStoreOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get blob store o k response has a 5xx status code
func (o *GetBlobStoreOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get blob store o k response a status code equal to that given
func (o *GetBlobStoreOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetBlobStoreOK) Error() string {
	return fmt.Sprintf("[GET /v1/blobstores/s3/{name}][%d] getBlobStoreOK  %+v", 200, o.Payload)
}

func (o *GetBlobStoreOK) String() string {
	return fmt.Sprintf("[GET /v1/blobstores/s3/{name}][%d] getBlobStoreOK  %+v", 200, o.Payload)
}

func (o *GetBlobStoreOK) GetPayload() *models.S3BlobStoreAPIModel {
	return o.Payload
}

func (o *GetBlobStoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3BlobStoreAPIModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBlobStoreBadRequest creates a GetBlobStoreBadRequest with default headers values
func NewGetBlobStoreBadRequest() *GetBlobStoreBadRequest {
	return &GetBlobStoreBadRequest{}
}

/*
	GetBlobStoreBadRequest describes a response with status code 400, with default header values.

Specified S3 blob store doesn't exist
*/
type GetBlobStoreBadRequest struct {
}

// IsSuccess returns true when this get blob store bad request response has a 2xx status code
func (o *GetBlobStoreBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get blob store bad request response has a 3xx status code
func (o *GetBlobStoreBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get blob store bad request response has a 4xx status code
func (o *GetBlobStoreBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get blob store bad request response has a 5xx status code
func (o *GetBlobStoreBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get blob store bad request response a status code equal to that given
func (o *GetBlobStoreBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetBlobStoreBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/blobstores/s3/{name}][%d] getBlobStoreBadRequest ", 400)
}

func (o *GetBlobStoreBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/blobstores/s3/{name}][%d] getBlobStoreBadRequest ", 400)
}

func (o *GetBlobStoreBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBlobStoreUnauthorized creates a GetBlobStoreUnauthorized with default headers values
func NewGetBlobStoreUnauthorized() *GetBlobStoreUnauthorized {
	return &GetBlobStoreUnauthorized{}
}

/*
	GetBlobStoreUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type GetBlobStoreUnauthorized struct {
}

// IsSuccess returns true when this get blob store unauthorized response has a 2xx status code
func (o *GetBlobStoreUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get blob store unauthorized response has a 3xx status code
func (o *GetBlobStoreUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get blob store unauthorized response has a 4xx status code
func (o *GetBlobStoreUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get blob store unauthorized response has a 5xx status code
func (o *GetBlobStoreUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get blob store unauthorized response a status code equal to that given
func (o *GetBlobStoreUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetBlobStoreUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/blobstores/s3/{name}][%d] getBlobStoreUnauthorized ", 401)
}

func (o *GetBlobStoreUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/blobstores/s3/{name}][%d] getBlobStoreUnauthorized ", 401)
}

func (o *GetBlobStoreUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBlobStoreForbidden creates a GetBlobStoreForbidden with default headers values
func NewGetBlobStoreForbidden() *GetBlobStoreForbidden {
	return &GetBlobStoreForbidden{}
}

/*
	GetBlobStoreForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type GetBlobStoreForbidden struct {
}

// IsSuccess returns true when this get blob store forbidden response has a 2xx status code
func (o *GetBlobStoreForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get blob store forbidden response has a 3xx status code
func (o *GetBlobStoreForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get blob store forbidden response has a 4xx status code
func (o *GetBlobStoreForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get blob store forbidden response has a 5xx status code
func (o *GetBlobStoreForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get blob store forbidden response a status code equal to that given
func (o *GetBlobStoreForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetBlobStoreForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/blobstores/s3/{name}][%d] getBlobStoreForbidden ", 403)
}

func (o *GetBlobStoreForbidden) String() string {
	return fmt.Sprintf("[GET /v1/blobstores/s3/{name}][%d] getBlobStoreForbidden ", 403)
}

func (o *GetBlobStoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
