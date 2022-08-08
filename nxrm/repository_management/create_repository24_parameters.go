// Code generated by go-swagger; DO NOT EDIT.

package repository_management

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

// NewCreateRepository24Params creates a new CreateRepository24Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRepository24Params() *CreateRepository24Params {
	return &CreateRepository24Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepository24ParamsWithTimeout creates a new CreateRepository24Params object
// with the ability to set a timeout on a request.
func NewCreateRepository24ParamsWithTimeout(timeout time.Duration) *CreateRepository24Params {
	return &CreateRepository24Params{
		timeout: timeout,
	}
}

// NewCreateRepository24ParamsWithContext creates a new CreateRepository24Params object
// with the ability to set a context for a request.
func NewCreateRepository24ParamsWithContext(ctx context.Context) *CreateRepository24Params {
	return &CreateRepository24Params{
		Context: ctx,
	}
}

// NewCreateRepository24ParamsWithHTTPClient creates a new CreateRepository24Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRepository24ParamsWithHTTPClient(client *http.Client) *CreateRepository24Params {
	return &CreateRepository24Params{
		HTTPClient: client,
	}
}

/* CreateRepository24Params contains all the parameters to send to the API endpoint
   for the create repository 24 operation.

   Typically these are written to a http.Request.
*/
type CreateRepository24Params struct {

	// Body.
	Body *models.HelmProxyRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 24 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository24Params) WithDefaults() *CreateRepository24Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 24 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository24Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 24 params
func (o *CreateRepository24Params) WithTimeout(timeout time.Duration) *CreateRepository24Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 24 params
func (o *CreateRepository24Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 24 params
func (o *CreateRepository24Params) WithContext(ctx context.Context) *CreateRepository24Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 24 params
func (o *CreateRepository24Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 24 params
func (o *CreateRepository24Params) WithHTTPClient(client *http.Client) *CreateRepository24Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 24 params
func (o *CreateRepository24Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 24 params
func (o *CreateRepository24Params) WithBody(body *models.HelmProxyRepositoryAPIRequest) *CreateRepository24Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 24 params
func (o *CreateRepository24Params) SetBody(body *models.HelmProxyRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepository24Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
