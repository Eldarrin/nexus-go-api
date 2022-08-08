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

// NewCreateRepository39Params creates a new CreateRepository39Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRepository39Params() *CreateRepository39Params {
	return &CreateRepository39Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepository39ParamsWithTimeout creates a new CreateRepository39Params object
// with the ability to set a timeout on a request.
func NewCreateRepository39ParamsWithTimeout(timeout time.Duration) *CreateRepository39Params {
	return &CreateRepository39Params{
		timeout: timeout,
	}
}

// NewCreateRepository39ParamsWithContext creates a new CreateRepository39Params object
// with the ability to set a context for a request.
func NewCreateRepository39ParamsWithContext(ctx context.Context) *CreateRepository39Params {
	return &CreateRepository39Params{
		Context: ctx,
	}
}

// NewCreateRepository39ParamsWithHTTPClient creates a new CreateRepository39Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRepository39ParamsWithHTTPClient(client *http.Client) *CreateRepository39Params {
	return &CreateRepository39Params{
		HTTPClient: client,
	}
}

/* CreateRepository39Params contains all the parameters to send to the API endpoint
   for the create repository 39 operation.

   Typically these are written to a http.Request.
*/
type CreateRepository39Params struct {

	// Body.
	Body *models.BowerHostedRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 39 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository39Params) WithDefaults() *CreateRepository39Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 39 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository39Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 39 params
func (o *CreateRepository39Params) WithTimeout(timeout time.Duration) *CreateRepository39Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 39 params
func (o *CreateRepository39Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 39 params
func (o *CreateRepository39Params) WithContext(ctx context.Context) *CreateRepository39Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 39 params
func (o *CreateRepository39Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 39 params
func (o *CreateRepository39Params) WithHTTPClient(client *http.Client) *CreateRepository39Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 39 params
func (o *CreateRepository39Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 39 params
func (o *CreateRepository39Params) WithBody(body *models.BowerHostedRepositoryAPIRequest) *CreateRepository39Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 39 params
func (o *CreateRepository39Params) SetBody(body *models.BowerHostedRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepository39Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
