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

	"github.com/Eldarrin/nexus-go-api/models"
)

// NewCreateBlobStore1Params creates a new CreateBlobStore1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateBlobStore1Params() *CreateBlobStore1Params {
	return &CreateBlobStore1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBlobStore1ParamsWithTimeout creates a new CreateBlobStore1Params object
// with the ability to set a timeout on a request.
func NewCreateBlobStore1ParamsWithTimeout(timeout time.Duration) *CreateBlobStore1Params {
	return &CreateBlobStore1Params{
		timeout: timeout,
	}
}

// NewCreateBlobStore1ParamsWithContext creates a new CreateBlobStore1Params object
// with the ability to set a context for a request.
func NewCreateBlobStore1ParamsWithContext(ctx context.Context) *CreateBlobStore1Params {
	return &CreateBlobStore1Params{
		Context: ctx,
	}
}

// NewCreateBlobStore1ParamsWithHTTPClient creates a new CreateBlobStore1Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateBlobStore1ParamsWithHTTPClient(client *http.Client) *CreateBlobStore1Params {
	return &CreateBlobStore1Params{
		HTTPClient: client,
	}
}

/*
CreateBlobStore1Params contains all the parameters to send to the API endpoint

	for the create blob store 1 operation.

	Typically these are written to a http.Request.
*/
type CreateBlobStore1Params struct {

	// Body.
	Body *models.AzureBlobStoreAPIModel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create blob store 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBlobStore1Params) WithDefaults() *CreateBlobStore1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create blob store 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBlobStore1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create blob store 1 params
func (o *CreateBlobStore1Params) WithTimeout(timeout time.Duration) *CreateBlobStore1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create blob store 1 params
func (o *CreateBlobStore1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create blob store 1 params
func (o *CreateBlobStore1Params) WithContext(ctx context.Context) *CreateBlobStore1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create blob store 1 params
func (o *CreateBlobStore1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create blob store 1 params
func (o *CreateBlobStore1Params) WithHTTPClient(client *http.Client) *CreateBlobStore1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create blob store 1 params
func (o *CreateBlobStore1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create blob store 1 params
func (o *CreateBlobStore1Params) WithBody(body *models.AzureBlobStoreAPIModel) *CreateBlobStore1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create blob store 1 params
func (o *CreateBlobStore1Params) SetBody(body *models.AzureBlobStoreAPIModel) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBlobStore1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
