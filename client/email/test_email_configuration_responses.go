// Code generated by go-swagger; DO NOT EDIT.

package email

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Eldarrin/nexus-go-api/models"
)

// TestEmailConfigurationReader is a Reader for the TestEmailConfiguration structure.
type TestEmailConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestEmailConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTestEmailConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewTestEmailConfigurationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTestEmailConfigurationOK creates a TestEmailConfigurationOK with default headers values
func NewTestEmailConfigurationOK() *TestEmailConfigurationOK {
	return &TestEmailConfigurationOK{}
}

/* TestEmailConfigurationOK describes a response with status code 200, with default header values.

Validation was complete, look at the body to determine success
*/
type TestEmailConfigurationOK struct {
	Payload *models.APIEmailValidation
}

// IsSuccess returns true when this test email configuration o k response has a 2xx status code
func (o *TestEmailConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this test email configuration o k response has a 3xx status code
func (o *TestEmailConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this test email configuration o k response has a 4xx status code
func (o *TestEmailConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this test email configuration o k response has a 5xx status code
func (o *TestEmailConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this test email configuration o k response a status code equal to that given
func (o *TestEmailConfigurationOK) IsCode(code int) bool {
	return code == 200
}

func (o *TestEmailConfigurationOK) Error() string {
	return fmt.Sprintf("[POST /v1/email/verify][%d] testEmailConfigurationOK  %+v", 200, o.Payload)
}

func (o *TestEmailConfigurationOK) String() string {
	return fmt.Sprintf("[POST /v1/email/verify][%d] testEmailConfigurationOK  %+v", 200, o.Payload)
}

func (o *TestEmailConfigurationOK) GetPayload() *models.APIEmailValidation {
	return o.Payload
}

func (o *TestEmailConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIEmailValidation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestEmailConfigurationForbidden creates a TestEmailConfigurationForbidden with default headers values
func NewTestEmailConfigurationForbidden() *TestEmailConfigurationForbidden {
	return &TestEmailConfigurationForbidden{}
}

/* TestEmailConfigurationForbidden describes a response with status code 403, with default header values.

Insufficient permissions to verify the email configuration
*/
type TestEmailConfigurationForbidden struct {
}

// IsSuccess returns true when this test email configuration forbidden response has a 2xx status code
func (o *TestEmailConfigurationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this test email configuration forbidden response has a 3xx status code
func (o *TestEmailConfigurationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this test email configuration forbidden response has a 4xx status code
func (o *TestEmailConfigurationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this test email configuration forbidden response has a 5xx status code
func (o *TestEmailConfigurationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this test email configuration forbidden response a status code equal to that given
func (o *TestEmailConfigurationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *TestEmailConfigurationForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/email/verify][%d] testEmailConfigurationForbidden ", 403)
}

func (o *TestEmailConfigurationForbidden) String() string {
	return fmt.Sprintf("[POST /v1/email/verify][%d] testEmailConfigurationForbidden ", 403)
}

func (o *TestEmailConfigurationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
