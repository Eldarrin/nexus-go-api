// Code generated by go-swagger; DO NOT EDIT.

package blob_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteBlobStoreReader is a Reader for the DeleteBlobStore structure.
type DeleteBlobStoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBlobStoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDeleteBlobStoreDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDeleteBlobStoreDefault creates a DeleteBlobStoreDefault with default headers values
func NewDeleteBlobStoreDefault(code int) *DeleteBlobStoreDefault {
	return &DeleteBlobStoreDefault{
		_statusCode: code,
	}
}

/* DeleteBlobStoreDefault describes a response with status code -1, with default header values.

successful operation
*/
type DeleteBlobStoreDefault struct {
	_statusCode int
}

// Code gets the status code for the delete blob store default response
func (o *DeleteBlobStoreDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete blob store default response has a 2xx status code
func (o *DeleteBlobStoreDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete blob store default response has a 3xx status code
func (o *DeleteBlobStoreDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete blob store default response has a 4xx status code
func (o *DeleteBlobStoreDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete blob store default response has a 5xx status code
func (o *DeleteBlobStoreDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete blob store default response a status code equal to that given
func (o *DeleteBlobStoreDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteBlobStoreDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/blobstores/{name}][%d] deleteBlobStore default ", o._statusCode)
}

func (o *DeleteBlobStoreDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/blobstores/{name}][%d] deleteBlobStore default ", o._statusCode)
}

func (o *DeleteBlobStoreDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
