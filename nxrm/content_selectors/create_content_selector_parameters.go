// Code generated by go-swagger; DO NOT EDIT.

package content_selectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"nexus-go-api/models"
)

// NewCreateContentSelectorParams creates a new CreateContentSelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateContentSelectorParams() *CreateContentSelectorParams {
	return &CreateContentSelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateContentSelectorParamsWithTimeout creates a new CreateContentSelectorParams object
// with the ability to set a timeout on a request.
func NewCreateContentSelectorParamsWithTimeout(timeout time.Duration) *CreateContentSelectorParams {
	return &CreateContentSelectorParams{
		timeout: timeout,
	}
}

// NewCreateContentSelectorParamsWithContext creates a new CreateContentSelectorParams object
// with the ability to set a context for a request.
func NewCreateContentSelectorParamsWithContext(ctx context.Context) *CreateContentSelectorParams {
	return &CreateContentSelectorParams{
		Context: ctx,
	}
}

// NewCreateContentSelectorParamsWithHTTPClient creates a new CreateContentSelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateContentSelectorParamsWithHTTPClient(client *http.Client) *CreateContentSelectorParams {
	return &CreateContentSelectorParams{
		HTTPClient: client,
	}
}

/* CreateContentSelectorParams contains all the parameters to send to the API endpoint
   for the create content selector operation.

   Typically these are written to a http.Request.
*/
type CreateContentSelectorParams struct {

	// Body.
	Body *models.ContentSelectorAPICreateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create content selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateContentSelectorParams) WithDefaults() *CreateContentSelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create content selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateContentSelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create content selector params
func (o *CreateContentSelectorParams) WithTimeout(timeout time.Duration) *CreateContentSelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create content selector params
func (o *CreateContentSelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create content selector params
func (o *CreateContentSelectorParams) WithContext(ctx context.Context) *CreateContentSelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create content selector params
func (o *CreateContentSelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create content selector params
func (o *CreateContentSelectorParams) WithHTTPClient(client *http.Client) *CreateContentSelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create content selector params
func (o *CreateContentSelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create content selector params
func (o *CreateContentSelectorParams) WithBody(body *models.ContentSelectorAPICreateRequest) *CreateContentSelectorParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create content selector params
func (o *CreateContentSelectorParams) SetBody(body *models.ContentSelectorAPICreateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateContentSelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
