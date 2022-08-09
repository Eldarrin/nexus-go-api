// Code generated by go-swagger; DO NOT EDIT.

package content_selectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteContentSelectorReader is a Reader for the DeleteContentSelector structure.
type DeleteContentSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteContentSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteContentSelectorNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteContentSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteContentSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteContentSelectorNoContent creates a DeleteContentSelectorNoContent with default headers values
func NewDeleteContentSelectorNoContent() *DeleteContentSelectorNoContent {
	return &DeleteContentSelectorNoContent{}
}

/*
	DeleteContentSelectorNoContent describes a response with status code 204, with default header values.

Content selector deleted successfully
*/
type DeleteContentSelectorNoContent struct {
}

// IsSuccess returns true when this delete content selector no content response has a 2xx status code
func (o *DeleteContentSelectorNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete content selector no content response has a 3xx status code
func (o *DeleteContentSelectorNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete content selector no content response has a 4xx status code
func (o *DeleteContentSelectorNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete content selector no content response has a 5xx status code
func (o *DeleteContentSelectorNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete content selector no content response a status code equal to that given
func (o *DeleteContentSelectorNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteContentSelectorNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/security/content-selectors/{name}][%d] deleteContentSelectorNoContent ", 204)
}

func (o *DeleteContentSelectorNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/security/content-selectors/{name}][%d] deleteContentSelectorNoContent ", 204)
}

func (o *DeleteContentSelectorNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteContentSelectorBadRequest creates a DeleteContentSelectorBadRequest with default headers values
func NewDeleteContentSelectorBadRequest() *DeleteContentSelectorBadRequest {
	return &DeleteContentSelectorBadRequest{}
}

/*
	DeleteContentSelectorBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type DeleteContentSelectorBadRequest struct {
}

// IsSuccess returns true when this delete content selector bad request response has a 2xx status code
func (o *DeleteContentSelectorBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete content selector bad request response has a 3xx status code
func (o *DeleteContentSelectorBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete content selector bad request response has a 4xx status code
func (o *DeleteContentSelectorBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete content selector bad request response has a 5xx status code
func (o *DeleteContentSelectorBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete content selector bad request response a status code equal to that given
func (o *DeleteContentSelectorBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteContentSelectorBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/security/content-selectors/{name}][%d] deleteContentSelectorBadRequest ", 400)
}

func (o *DeleteContentSelectorBadRequest) String() string {
	return fmt.Sprintf("[DELETE /v1/security/content-selectors/{name}][%d] deleteContentSelectorBadRequest ", 400)
}

func (o *DeleteContentSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteContentSelectorForbidden creates a DeleteContentSelectorForbidden with default headers values
func NewDeleteContentSelectorForbidden() *DeleteContentSelectorForbidden {
	return &DeleteContentSelectorForbidden{}
}

/*
	DeleteContentSelectorForbidden describes a response with status code 403, with default header values.

Insufficient permissions to delete the content selector
*/
type DeleteContentSelectorForbidden struct {
}

// IsSuccess returns true when this delete content selector forbidden response has a 2xx status code
func (o *DeleteContentSelectorForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete content selector forbidden response has a 3xx status code
func (o *DeleteContentSelectorForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete content selector forbidden response has a 4xx status code
func (o *DeleteContentSelectorForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete content selector forbidden response has a 5xx status code
func (o *DeleteContentSelectorForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete content selector forbidden response a status code equal to that given
func (o *DeleteContentSelectorForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteContentSelectorForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/security/content-selectors/{name}][%d] deleteContentSelectorForbidden ", 403)
}

func (o *DeleteContentSelectorForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/security/content-selectors/{name}][%d] deleteContentSelectorForbidden ", 403)
}

func (o *DeleteContentSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
