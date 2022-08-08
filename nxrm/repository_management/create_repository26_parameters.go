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

// NewCreateRepository26Params creates a new CreateRepository26Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRepository26Params() *CreateRepository26Params {
	return &CreateRepository26Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepository26ParamsWithTimeout creates a new CreateRepository26Params object
// with the ability to set a timeout on a request.
func NewCreateRepository26ParamsWithTimeout(timeout time.Duration) *CreateRepository26Params {
	return &CreateRepository26Params{
		timeout: timeout,
	}
}

// NewCreateRepository26ParamsWithContext creates a new CreateRepository26Params object
// with the ability to set a context for a request.
func NewCreateRepository26ParamsWithContext(ctx context.Context) *CreateRepository26Params {
	return &CreateRepository26Params{
		Context: ctx,
	}
}

// NewCreateRepository26ParamsWithHTTPClient creates a new CreateRepository26Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRepository26ParamsWithHTTPClient(client *http.Client) *CreateRepository26Params {
	return &CreateRepository26Params{
		HTTPClient: client,
	}
}

/* CreateRepository26Params contains all the parameters to send to the API endpoint
   for the create repository 26 operation.

   Typically these are written to a http.Request.
*/
type CreateRepository26Params struct {

	// Body.
	Body *models.PypiGroupRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 26 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository26Params) WithDefaults() *CreateRepository26Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 26 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository26Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 26 params
func (o *CreateRepository26Params) WithTimeout(timeout time.Duration) *CreateRepository26Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 26 params
func (o *CreateRepository26Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 26 params
func (o *CreateRepository26Params) WithContext(ctx context.Context) *CreateRepository26Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 26 params
func (o *CreateRepository26Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 26 params
func (o *CreateRepository26Params) WithHTTPClient(client *http.Client) *CreateRepository26Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 26 params
func (o *CreateRepository26Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 26 params
func (o *CreateRepository26Params) WithBody(body *models.PypiGroupRepositoryAPIRequest) *CreateRepository26Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 26 params
func (o *CreateRepository26Params) SetBody(body *models.PypiGroupRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepository26Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
