// Code generated by go-swagger; DO NOT EDIT.

package blob_store

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

	"Eldarrin/nexus-go-api/models"
)

// NewCreateFileBlobStoreParams creates a new CreateFileBlobStoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateFileBlobStoreParams() *CreateFileBlobStoreParams {
	return &CreateFileBlobStoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateFileBlobStoreParamsWithTimeout creates a new CreateFileBlobStoreParams object
// with the ability to set a timeout on a request.
func NewCreateFileBlobStoreParamsWithTimeout(timeout time.Duration) *CreateFileBlobStoreParams {
	return &CreateFileBlobStoreParams{
		timeout: timeout,
	}
}

// NewCreateFileBlobStoreParamsWithContext creates a new CreateFileBlobStoreParams object
// with the ability to set a context for a request.
func NewCreateFileBlobStoreParamsWithContext(ctx context.Context) *CreateFileBlobStoreParams {
	return &CreateFileBlobStoreParams{
		Context: ctx,
	}
}

// NewCreateFileBlobStoreParamsWithHTTPClient creates a new CreateFileBlobStoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateFileBlobStoreParamsWithHTTPClient(client *http.Client) *CreateFileBlobStoreParams {
	return &CreateFileBlobStoreParams{
		HTTPClient: client,
	}
}

/* CreateFileBlobStoreParams contains all the parameters to send to the API endpoint
   for the create file blob store operation.

   Typically these are written to a http.Request.
*/
type CreateFileBlobStoreParams struct {

	// Body.
	Body *models.FileBlobStoreAPICreateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create file blob store params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateFileBlobStoreParams) WithDefaults() *CreateFileBlobStoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create file blob store params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateFileBlobStoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create file blob store params
func (o *CreateFileBlobStoreParams) WithTimeout(timeout time.Duration) *CreateFileBlobStoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create file blob store params
func (o *CreateFileBlobStoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create file blob store params
func (o *CreateFileBlobStoreParams) WithContext(ctx context.Context) *CreateFileBlobStoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create file blob store params
func (o *CreateFileBlobStoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create file blob store params
func (o *CreateFileBlobStoreParams) WithHTTPClient(client *http.Client) *CreateFileBlobStoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create file blob store params
func (o *CreateFileBlobStoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create file blob store params
func (o *CreateFileBlobStoreParams) WithBody(body *models.FileBlobStoreAPICreateRequest) *CreateFileBlobStoreParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create file blob store params
func (o *CreateFileBlobStoreParams) SetBody(body *models.FileBlobStoreAPICreateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateFileBlobStoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
