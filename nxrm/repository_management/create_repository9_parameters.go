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

// NewCreateRepository9Params creates a new CreateRepository9Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRepository9Params() *CreateRepository9Params {
	return &CreateRepository9Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepository9ParamsWithTimeout creates a new CreateRepository9Params object
// with the ability to set a timeout on a request.
func NewCreateRepository9ParamsWithTimeout(timeout time.Duration) *CreateRepository9Params {
	return &CreateRepository9Params{
		timeout: timeout,
	}
}

// NewCreateRepository9ParamsWithContext creates a new CreateRepository9Params object
// with the ability to set a context for a request.
func NewCreateRepository9ParamsWithContext(ctx context.Context) *CreateRepository9Params {
	return &CreateRepository9Params{
		Context: ctx,
	}
}

// NewCreateRepository9ParamsWithHTTPClient creates a new CreateRepository9Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRepository9ParamsWithHTTPClient(client *http.Client) *CreateRepository9Params {
	return &CreateRepository9Params{
		HTTPClient: client,
	}
}

/* CreateRepository9Params contains all the parameters to send to the API endpoint
   for the create repository 9 operation.

   Typically these are written to a http.Request.
*/
type CreateRepository9Params struct {

	// Body.
	Body *models.NpmHostedRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 9 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository9Params) WithDefaults() *CreateRepository9Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 9 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository9Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 9 params
func (o *CreateRepository9Params) WithTimeout(timeout time.Duration) *CreateRepository9Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 9 params
func (o *CreateRepository9Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 9 params
func (o *CreateRepository9Params) WithContext(ctx context.Context) *CreateRepository9Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 9 params
func (o *CreateRepository9Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 9 params
func (o *CreateRepository9Params) WithHTTPClient(client *http.Client) *CreateRepository9Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 9 params
func (o *CreateRepository9Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 9 params
func (o *CreateRepository9Params) WithBody(body *models.NpmHostedRepositoryAPIRequest) *CreateRepository9Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 9 params
func (o *CreateRepository9Params) SetBody(body *models.NpmHostedRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepository9Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
