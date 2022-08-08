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

// NewUpdateFileBlobStoreParams creates a new UpdateFileBlobStoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateFileBlobStoreParams() *UpdateFileBlobStoreParams {
	return &UpdateFileBlobStoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateFileBlobStoreParamsWithTimeout creates a new UpdateFileBlobStoreParams object
// with the ability to set a timeout on a request.
func NewUpdateFileBlobStoreParamsWithTimeout(timeout time.Duration) *UpdateFileBlobStoreParams {
	return &UpdateFileBlobStoreParams{
		timeout: timeout,
	}
}

// NewUpdateFileBlobStoreParamsWithContext creates a new UpdateFileBlobStoreParams object
// with the ability to set a context for a request.
func NewUpdateFileBlobStoreParamsWithContext(ctx context.Context) *UpdateFileBlobStoreParams {
	return &UpdateFileBlobStoreParams{
		Context: ctx,
	}
}

// NewUpdateFileBlobStoreParamsWithHTTPClient creates a new UpdateFileBlobStoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateFileBlobStoreParamsWithHTTPClient(client *http.Client) *UpdateFileBlobStoreParams {
	return &UpdateFileBlobStoreParams{
		HTTPClient: client,
	}
}

/* UpdateFileBlobStoreParams contains all the parameters to send to the API endpoint
   for the update file blob store operation.

   Typically these are written to a http.Request.
*/
type UpdateFileBlobStoreParams struct {

	// Body.
	Body *models.FileBlobStoreAPIUpdateRequest

	/* Name.

	   The name of the file blob store to update
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update file blob store params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateFileBlobStoreParams) WithDefaults() *UpdateFileBlobStoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update file blob store params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateFileBlobStoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update file blob store params
func (o *UpdateFileBlobStoreParams) WithTimeout(timeout time.Duration) *UpdateFileBlobStoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update file blob store params
func (o *UpdateFileBlobStoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update file blob store params
func (o *UpdateFileBlobStoreParams) WithContext(ctx context.Context) *UpdateFileBlobStoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update file blob store params
func (o *UpdateFileBlobStoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update file blob store params
func (o *UpdateFileBlobStoreParams) WithHTTPClient(client *http.Client) *UpdateFileBlobStoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update file blob store params
func (o *UpdateFileBlobStoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update file blob store params
func (o *UpdateFileBlobStoreParams) WithBody(body *models.FileBlobStoreAPIUpdateRequest) *UpdateFileBlobStoreParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update file blob store params
func (o *UpdateFileBlobStoreParams) SetBody(body *models.FileBlobStoreAPIUpdateRequest) {
	o.Body = body
}

// WithName adds the name to the update file blob store params
func (o *UpdateFileBlobStoreParams) WithName(name string) *UpdateFileBlobStoreParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update file blob store params
func (o *UpdateFileBlobStoreParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateFileBlobStoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
