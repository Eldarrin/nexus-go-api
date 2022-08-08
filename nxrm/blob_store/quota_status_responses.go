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

// QuotaStatusReader is a Reader for the QuotaStatus structure.
type QuotaStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuotaStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQuotaStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQuotaStatusOK creates a QuotaStatusOK with default headers values
func NewQuotaStatusOK() *QuotaStatusOK {
	return &QuotaStatusOK{}
}

/* QuotaStatusOK describes a response with status code 200, with default header values.

successful operation
*/
type QuotaStatusOK struct {
	Payload *models.BlobStoreQuotaResultXO
}

// IsSuccess returns true when this quota status o k response has a 2xx status code
func (o *QuotaStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this quota status o k response has a 3xx status code
func (o *QuotaStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this quota status o k response has a 4xx status code
func (o *QuotaStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this quota status o k response has a 5xx status code
func (o *QuotaStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this quota status o k response a status code equal to that given
func (o *QuotaStatusOK) IsCode(code int) bool {
	return code == 200
}

func (o *QuotaStatusOK) Error() string {
	return fmt.Sprintf("[GET /v1/blobstores/{name}/quota-status][%d] quotaStatusOK  %+v", 200, o.Payload)
}

func (o *QuotaStatusOK) String() string {
	return fmt.Sprintf("[GET /v1/blobstores/{name}/quota-status][%d] quotaStatusOK  %+v", 200, o.Payload)
}

func (o *QuotaStatusOK) GetPayload() *models.BlobStoreQuotaResultXO {
	return o.Payload
}

func (o *QuotaStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BlobStoreQuotaResultXO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
