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

// NewCreateRepository29Params creates a new CreateRepository29Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRepository29Params() *CreateRepository29Params {
	return &CreateRepository29Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepository29ParamsWithTimeout creates a new CreateRepository29Params object
// with the ability to set a timeout on a request.
func NewCreateRepository29ParamsWithTimeout(timeout time.Duration) *CreateRepository29Params {
	return &CreateRepository29Params{
		timeout: timeout,
	}
}

// NewCreateRepository29ParamsWithContext creates a new CreateRepository29Params object
// with the ability to set a context for a request.
func NewCreateRepository29ParamsWithContext(ctx context.Context) *CreateRepository29Params {
	return &CreateRepository29Params{
		Context: ctx,
	}
}

// NewCreateRepository29ParamsWithHTTPClient creates a new CreateRepository29Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRepository29ParamsWithHTTPClient(client *http.Client) *CreateRepository29Params {
	return &CreateRepository29Params{
		HTTPClient: client,
	}
}

/* CreateRepository29Params contains all the parameters to send to the API endpoint
   for the create repository 29 operation.

   Typically these are written to a http.Request.
*/
type CreateRepository29Params struct {

	// Body.
	Body *models.CondaProxyRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 29 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository29Params) WithDefaults() *CreateRepository29Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 29 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository29Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 29 params
func (o *CreateRepository29Params) WithTimeout(timeout time.Duration) *CreateRepository29Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 29 params
func (o *CreateRepository29Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 29 params
func (o *CreateRepository29Params) WithContext(ctx context.Context) *CreateRepository29Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 29 params
func (o *CreateRepository29Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 29 params
func (o *CreateRepository29Params) WithHTTPClient(client *http.Client) *CreateRepository29Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 29 params
func (o *CreateRepository29Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 29 params
func (o *CreateRepository29Params) WithBody(body *models.CondaProxyRepositoryAPIRequest) *CreateRepository29Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 29 params
func (o *CreateRepository29Params) SetBody(body *models.CondaProxyRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepository29Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
