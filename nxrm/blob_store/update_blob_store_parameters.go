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

// NewUpdateBlobStoreParams creates a new UpdateBlobStoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateBlobStoreParams() *UpdateBlobStoreParams {
	return &UpdateBlobStoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBlobStoreParamsWithTimeout creates a new UpdateBlobStoreParams object
// with the ability to set a timeout on a request.
func NewUpdateBlobStoreParamsWithTimeout(timeout time.Duration) *UpdateBlobStoreParams {
	return &UpdateBlobStoreParams{
		timeout: timeout,
	}
}

// NewUpdateBlobStoreParamsWithContext creates a new UpdateBlobStoreParams object
// with the ability to set a context for a request.
func NewUpdateBlobStoreParamsWithContext(ctx context.Context) *UpdateBlobStoreParams {
	return &UpdateBlobStoreParams{
		Context: ctx,
	}
}

// NewUpdateBlobStoreParamsWithHTTPClient creates a new UpdateBlobStoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateBlobStoreParamsWithHTTPClient(client *http.Client) *UpdateBlobStoreParams {
	return &UpdateBlobStoreParams{
		HTTPClient: client,
	}
}

/* UpdateBlobStoreParams contains all the parameters to send to the API endpoint
   for the update blob store operation.

   Typically these are written to a http.Request.
*/
type UpdateBlobStoreParams struct {

	// Body.
	Body *models.S3BlobStoreAPIModel

	/* Name.

	   Name of the blob store to update
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update blob store params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBlobStoreParams) WithDefaults() *UpdateBlobStoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update blob store params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBlobStoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update blob store params
func (o *UpdateBlobStoreParams) WithTimeout(timeout time.Duration) *UpdateBlobStoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update blob store params
func (o *UpdateBlobStoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update blob store params
func (o *UpdateBlobStoreParams) WithContext(ctx context.Context) *UpdateBlobStoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update blob store params
func (o *UpdateBlobStoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update blob store params
func (o *UpdateBlobStoreParams) WithHTTPClient(client *http.Client) *UpdateBlobStoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update blob store params
func (o *UpdateBlobStoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update blob store params
func (o *UpdateBlobStoreParams) WithBody(body *models.S3BlobStoreAPIModel) *UpdateBlobStoreParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update blob store params
func (o *UpdateBlobStoreParams) SetBody(body *models.S3BlobStoreAPIModel) {
	o.Body = body
}

// WithName adds the name to the update blob store params
func (o *UpdateBlobStoreParams) WithName(name string) *UpdateBlobStoreParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update blob store params
func (o *UpdateBlobStoreParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBlobStoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
