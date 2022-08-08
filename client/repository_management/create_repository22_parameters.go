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

	"github.com/Eldarrin/nexus-go-api/models"
)

// NewCreateRepository22Params creates a new CreateRepository22Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRepository22Params() *CreateRepository22Params {
	return &CreateRepository22Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepository22ParamsWithTimeout creates a new CreateRepository22Params object
// with the ability to set a timeout on a request.
func NewCreateRepository22ParamsWithTimeout(timeout time.Duration) *CreateRepository22Params {
	return &CreateRepository22Params{
		timeout: timeout,
	}
}

// NewCreateRepository22ParamsWithContext creates a new CreateRepository22Params object
// with the ability to set a context for a request.
func NewCreateRepository22ParamsWithContext(ctx context.Context) *CreateRepository22Params {
	return &CreateRepository22Params{
		Context: ctx,
	}
}

// NewCreateRepository22ParamsWithHTTPClient creates a new CreateRepository22Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRepository22ParamsWithHTTPClient(client *http.Client) *CreateRepository22Params {
	return &CreateRepository22Params{
		HTTPClient: client,
	}
}

/* CreateRepository22Params contains all the parameters to send to the API endpoint
   for the create repository 22 operation.

   Typically these are written to a http.Request.
*/
type CreateRepository22Params struct {

	// Body.
	Body *models.YumProxyRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 22 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository22Params) WithDefaults() *CreateRepository22Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 22 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository22Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 22 params
func (o *CreateRepository22Params) WithTimeout(timeout time.Duration) *CreateRepository22Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 22 params
func (o *CreateRepository22Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 22 params
func (o *CreateRepository22Params) WithContext(ctx context.Context) *CreateRepository22Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 22 params
func (o *CreateRepository22Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 22 params
func (o *CreateRepository22Params) WithHTTPClient(client *http.Client) *CreateRepository22Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 22 params
func (o *CreateRepository22Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 22 params
func (o *CreateRepository22Params) WithBody(body *models.YumProxyRepositoryAPIRequest) *CreateRepository22Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 22 params
func (o *CreateRepository22Params) SetBody(body *models.YumProxyRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepository22Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
