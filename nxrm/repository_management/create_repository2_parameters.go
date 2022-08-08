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

	"Eldarrin/nexus-go-api/models"
)

// NewCreateRepository2Params creates a new CreateRepository2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRepository2Params() *CreateRepository2Params {
	return &CreateRepository2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepository2ParamsWithTimeout creates a new CreateRepository2Params object
// with the ability to set a timeout on a request.
func NewCreateRepository2ParamsWithTimeout(timeout time.Duration) *CreateRepository2Params {
	return &CreateRepository2Params{
		timeout: timeout,
	}
}

// NewCreateRepository2ParamsWithContext creates a new CreateRepository2Params object
// with the ability to set a context for a request.
func NewCreateRepository2ParamsWithContext(ctx context.Context) *CreateRepository2Params {
	return &CreateRepository2Params{
		Context: ctx,
	}
}

// NewCreateRepository2ParamsWithHTTPClient creates a new CreateRepository2Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRepository2ParamsWithHTTPClient(client *http.Client) *CreateRepository2Params {
	return &CreateRepository2Params{
		HTTPClient: client,
	}
}

/* CreateRepository2Params contains all the parameters to send to the API endpoint
   for the create repository 2 operation.

   Typically these are written to a http.Request.
*/
type CreateRepository2Params struct {

	// Body.
	Body *models.MavenProxyRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository2Params) WithDefaults() *CreateRepository2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 2 params
func (o *CreateRepository2Params) WithTimeout(timeout time.Duration) *CreateRepository2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 2 params
func (o *CreateRepository2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 2 params
func (o *CreateRepository2Params) WithContext(ctx context.Context) *CreateRepository2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 2 params
func (o *CreateRepository2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 2 params
func (o *CreateRepository2Params) WithHTTPClient(client *http.Client) *CreateRepository2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 2 params
func (o *CreateRepository2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 2 params
func (o *CreateRepository2Params) WithBody(body *models.MavenProxyRepositoryAPIRequest) *CreateRepository2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 2 params
func (o *CreateRepository2Params) SetBody(body *models.MavenProxyRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepository2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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