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

// ListBlobStoresReader is a Reader for the ListBlobStores structure.
type ListBlobStoresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListBlobStoresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListBlobStoresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListBlobStoresOK creates a ListBlobStoresOK with default headers values
func NewListBlobStoresOK() *ListBlobStoresOK {
	return &ListBlobStoresOK{}
}

/* ListBlobStoresOK describes a response with status code 200, with default header values.

successful operation
*/
type ListBlobStoresOK struct {
	Payload []*models.GenericBlobStoreAPIResponse
}

// IsSuccess returns true when this list blob stores o k response has a 2xx status code
func (o *ListBlobStoresOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list blob stores o k response has a 3xx status code
func (o *ListBlobStoresOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list blob stores o k response has a 4xx status code
func (o *ListBlobStoresOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list blob stores o k response has a 5xx status code
func (o *ListBlobStoresOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list blob stores o k response a status code equal to that given
func (o *ListBlobStoresOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListBlobStoresOK) Error() string {
	return fmt.Sprintf("[GET /v1/blobstores][%d] listBlobStoresOK  %+v", 200, o.Payload)
}

func (o *ListBlobStoresOK) String() string {
	return fmt.Sprintf("[GET /v1/blobstores][%d] listBlobStoresOK  %+v", 200, o.Payload)
}

func (o *ListBlobStoresOK) GetPayload() []*models.GenericBlobStoreAPIResponse {
	return o.Payload
}

func (o *ListBlobStoresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}