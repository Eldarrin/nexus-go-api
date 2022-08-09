// Code generated by go-swagger; DO NOT EDIT.

package blob_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateBlobStore1Reader is a Reader for the CreateBlobStore1 structure.
type CreateBlobStore1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBlobStore1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateBlobStore1Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateBlobStore1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateBlobStore1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateBlobStore1Created creates a CreateBlobStore1Created with default headers values
func NewCreateBlobStore1Created() *CreateBlobStore1Created {
	return &CreateBlobStore1Created{}
}

/*
	CreateBlobStore1Created describes a response with status code 201, with default header values.

Azure blob store created
*/
type CreateBlobStore1Created struct {
}

// IsSuccess returns true when this create blob store1 created response has a 2xx status code
func (o *CreateBlobStore1Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create blob store1 created response has a 3xx status code
func (o *CreateBlobStore1Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create blob store1 created response has a 4xx status code
func (o *CreateBlobStore1Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create blob store1 created response has a 5xx status code
func (o *CreateBlobStore1Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create blob store1 created response a status code equal to that given
func (o *CreateBlobStore1Created) IsCode(code int) bool {
	return code == 201
}

func (o *CreateBlobStore1Created) Error() string {
	return fmt.Sprintf("[POST /v1/blobstores/azure][%d] createBlobStore1Created ", 201)
}

func (o *CreateBlobStore1Created) String() string {
	return fmt.Sprintf("[POST /v1/blobstores/azure][%d] createBlobStore1Created ", 201)
}

func (o *CreateBlobStore1Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateBlobStore1Unauthorized creates a CreateBlobStore1Unauthorized with default headers values
func NewCreateBlobStore1Unauthorized() *CreateBlobStore1Unauthorized {
	return &CreateBlobStore1Unauthorized{}
}

/*
	CreateBlobStore1Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateBlobStore1Unauthorized struct {
}

// IsSuccess returns true when this create blob store1 unauthorized response has a 2xx status code
func (o *CreateBlobStore1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create blob store1 unauthorized response has a 3xx status code
func (o *CreateBlobStore1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create blob store1 unauthorized response has a 4xx status code
func (o *CreateBlobStore1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create blob store1 unauthorized response has a 5xx status code
func (o *CreateBlobStore1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create blob store1 unauthorized response a status code equal to that given
func (o *CreateBlobStore1Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateBlobStore1Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/blobstores/azure][%d] createBlobStore1Unauthorized ", 401)
}

func (o *CreateBlobStore1Unauthorized) String() string {
	return fmt.Sprintf("[POST /v1/blobstores/azure][%d] createBlobStore1Unauthorized ", 401)
}

func (o *CreateBlobStore1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateBlobStore1Forbidden creates a CreateBlobStore1Forbidden with default headers values
func NewCreateBlobStore1Forbidden() *CreateBlobStore1Forbidden {
	return &CreateBlobStore1Forbidden{}
}

/*
	CreateBlobStore1Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateBlobStore1Forbidden struct {
}

// IsSuccess returns true when this create blob store1 forbidden response has a 2xx status code
func (o *CreateBlobStore1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create blob store1 forbidden response has a 3xx status code
func (o *CreateBlobStore1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create blob store1 forbidden response has a 4xx status code
func (o *CreateBlobStore1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create blob store1 forbidden response has a 5xx status code
func (o *CreateBlobStore1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create blob store1 forbidden response a status code equal to that given
func (o *CreateBlobStore1Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateBlobStore1Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/blobstores/azure][%d] createBlobStore1Forbidden ", 403)
}

func (o *CreateBlobStore1Forbidden) String() string {
	return fmt.Sprintf("[POST /v1/blobstores/azure][%d] createBlobStore1Forbidden ", 403)
}

func (o *CreateBlobStore1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
