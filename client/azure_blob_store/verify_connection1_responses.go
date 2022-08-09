// Code generated by go-swagger; DO NOT EDIT.

package azure_blob_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VerifyConnection1Reader is a Reader for the VerifyConnection1 structure.
type VerifyConnection1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VerifyConnection1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVerifyConnection1NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVerifyConnection1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewVerifyConnection1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewVerifyConnection1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVerifyConnection1NoContent creates a VerifyConnection1NoContent with default headers values
func NewVerifyConnection1NoContent() *VerifyConnection1NoContent {
	return &VerifyConnection1NoContent{}
}

/*
	VerifyConnection1NoContent describes a response with status code 204, with default header values.

Azure Blob Store connection was successful
*/
type VerifyConnection1NoContent struct {
}

// IsSuccess returns true when this verify connection1 no content response has a 2xx status code
func (o *VerifyConnection1NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this verify connection1 no content response has a 3xx status code
func (o *VerifyConnection1NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this verify connection1 no content response has a 4xx status code
func (o *VerifyConnection1NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this verify connection1 no content response has a 5xx status code
func (o *VerifyConnection1NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this verify connection1 no content response a status code equal to that given
func (o *VerifyConnection1NoContent) IsCode(code int) bool {
	return code == 204
}

func (o *VerifyConnection1NoContent) Error() string {
	return fmt.Sprintf("[POST /v1/azureblobstore/test-connection][%d] verifyConnection1NoContent ", 204)
}

func (o *VerifyConnection1NoContent) String() string {
	return fmt.Sprintf("[POST /v1/azureblobstore/test-connection][%d] verifyConnection1NoContent ", 204)
}

func (o *VerifyConnection1NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVerifyConnection1BadRequest creates a VerifyConnection1BadRequest with default headers values
func NewVerifyConnection1BadRequest() *VerifyConnection1BadRequest {
	return &VerifyConnection1BadRequest{}
}

/*
	VerifyConnection1BadRequest describes a response with status code 400, with default header values.

Azure Blob Store connection failed
*/
type VerifyConnection1BadRequest struct {
}

// IsSuccess returns true when this verify connection1 bad request response has a 2xx status code
func (o *VerifyConnection1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this verify connection1 bad request response has a 3xx status code
func (o *VerifyConnection1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this verify connection1 bad request response has a 4xx status code
func (o *VerifyConnection1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this verify connection1 bad request response has a 5xx status code
func (o *VerifyConnection1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this verify connection1 bad request response a status code equal to that given
func (o *VerifyConnection1BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *VerifyConnection1BadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/azureblobstore/test-connection][%d] verifyConnection1BadRequest ", 400)
}

func (o *VerifyConnection1BadRequest) String() string {
	return fmt.Sprintf("[POST /v1/azureblobstore/test-connection][%d] verifyConnection1BadRequest ", 400)
}

func (o *VerifyConnection1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVerifyConnection1Unauthorized creates a VerifyConnection1Unauthorized with default headers values
func NewVerifyConnection1Unauthorized() *VerifyConnection1Unauthorized {
	return &VerifyConnection1Unauthorized{}
}

/*
	VerifyConnection1Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type VerifyConnection1Unauthorized struct {
}

// IsSuccess returns true when this verify connection1 unauthorized response has a 2xx status code
func (o *VerifyConnection1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this verify connection1 unauthorized response has a 3xx status code
func (o *VerifyConnection1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this verify connection1 unauthorized response has a 4xx status code
func (o *VerifyConnection1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this verify connection1 unauthorized response has a 5xx status code
func (o *VerifyConnection1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this verify connection1 unauthorized response a status code equal to that given
func (o *VerifyConnection1Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *VerifyConnection1Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/azureblobstore/test-connection][%d] verifyConnection1Unauthorized ", 401)
}

func (o *VerifyConnection1Unauthorized) String() string {
	return fmt.Sprintf("[POST /v1/azureblobstore/test-connection][%d] verifyConnection1Unauthorized ", 401)
}

func (o *VerifyConnection1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVerifyConnection1Forbidden creates a VerifyConnection1Forbidden with default headers values
func NewVerifyConnection1Forbidden() *VerifyConnection1Forbidden {
	return &VerifyConnection1Forbidden{}
}

/*
	VerifyConnection1Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type VerifyConnection1Forbidden struct {
}

// IsSuccess returns true when this verify connection1 forbidden response has a 2xx status code
func (o *VerifyConnection1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this verify connection1 forbidden response has a 3xx status code
func (o *VerifyConnection1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this verify connection1 forbidden response has a 4xx status code
func (o *VerifyConnection1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this verify connection1 forbidden response has a 5xx status code
func (o *VerifyConnection1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this verify connection1 forbidden response a status code equal to that given
func (o *VerifyConnection1Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *VerifyConnection1Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/azureblobstore/test-connection][%d] verifyConnection1Forbidden ", 403)
}

func (o *VerifyConnection1Forbidden) String() string {
	return fmt.Sprintf("[POST /v1/azureblobstore/test-connection][%d] verifyConnection1Forbidden ", 403)
}

func (o *VerifyConnection1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
