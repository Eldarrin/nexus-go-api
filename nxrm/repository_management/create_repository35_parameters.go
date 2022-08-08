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

// NewCreateRepository35Params creates a new CreateRepository35Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRepository35Params() *CreateRepository35Params {
	return &CreateRepository35Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepository35ParamsWithTimeout creates a new CreateRepository35Params object
// with the ability to set a timeout on a request.
func NewCreateRepository35ParamsWithTimeout(timeout time.Duration) *CreateRepository35Params {
	return &CreateRepository35Params{
		timeout: timeout,
	}
}

// NewCreateRepository35ParamsWithContext creates a new CreateRepository35Params object
// with the ability to set a context for a request.
func NewCreateRepository35ParamsWithContext(ctx context.Context) *CreateRepository35Params {
	return &CreateRepository35Params{
		Context: ctx,
	}
}

// NewCreateRepository35ParamsWithHTTPClient creates a new CreateRepository35Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRepository35ParamsWithHTTPClient(client *http.Client) *CreateRepository35Params {
	return &CreateRepository35Params{
		HTTPClient: client,
	}
}

/* CreateRepository35Params contains all the parameters to send to the API endpoint
   for the create repository 35 operation.

   Typically these are written to a http.Request.
*/
type CreateRepository35Params struct {

	// Body.
	Body *models.GolangGroupRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 35 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository35Params) WithDefaults() *CreateRepository35Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 35 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository35Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 35 params
func (o *CreateRepository35Params) WithTimeout(timeout time.Duration) *CreateRepository35Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 35 params
func (o *CreateRepository35Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 35 params
func (o *CreateRepository35Params) WithContext(ctx context.Context) *CreateRepository35Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 35 params
func (o *CreateRepository35Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 35 params
func (o *CreateRepository35Params) WithHTTPClient(client *http.Client) *CreateRepository35Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 35 params
func (o *CreateRepository35Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 35 params
func (o *CreateRepository35Params) WithBody(body *models.GolangGroupRepositoryAPIRequest) *CreateRepository35Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 35 params
func (o *CreateRepository35Params) SetBody(body *models.GolangGroupRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepository35Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
