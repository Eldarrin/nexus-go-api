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

// NewCreateNpmProxyRepositoryParams creates a new CreateNpmProxyRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNpmProxyRepositoryParams() *CreateNpmProxyRepositoryParams {
	return &CreateNpmProxyRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNpmProxyRepositoryParamsWithTimeout creates a new CreateNpmProxyRepositoryParams object
// with the ability to set a timeout on a request.
func NewCreateNpmProxyRepositoryParamsWithTimeout(timeout time.Duration) *CreateNpmProxyRepositoryParams {
	return &CreateNpmProxyRepositoryParams{
		timeout: timeout,
	}
}

// NewCreateNpmProxyRepositoryParamsWithContext creates a new CreateNpmProxyRepositoryParams object
// with the ability to set a context for a request.
func NewCreateNpmProxyRepositoryParamsWithContext(ctx context.Context) *CreateNpmProxyRepositoryParams {
	return &CreateNpmProxyRepositoryParams{
		Context: ctx,
	}
}

// NewCreateNpmProxyRepositoryParamsWithHTTPClient creates a new CreateNpmProxyRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNpmProxyRepositoryParamsWithHTTPClient(client *http.Client) *CreateNpmProxyRepositoryParams {
	return &CreateNpmProxyRepositoryParams{
		HTTPClient: client,
	}
}

/*
CreateNpmProxyRepositoryParams contains all the parameters to send to the API endpoint

	for the create repository 10 operation.

	Typically these are written to a http.Request.
*/
type CreateNpmProxyRepositoryParams struct {

	// Body.
	Body *models.NpmProxyRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 10 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNpmProxyRepositoryParams) WithDefaults() *CreateNpmProxyRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 10 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNpmProxyRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 10 params
func (o *CreateNpmProxyRepositoryParams) WithTimeout(timeout time.Duration) *CreateNpmProxyRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 10 params
func (o *CreateNpmProxyRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 10 params
func (o *CreateNpmProxyRepositoryParams) WithContext(ctx context.Context) *CreateNpmProxyRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 10 params
func (o *CreateNpmProxyRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 10 params
func (o *CreateNpmProxyRepositoryParams) WithHTTPClient(client *http.Client) *CreateNpmProxyRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 10 params
func (o *CreateNpmProxyRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 10 params
func (o *CreateNpmProxyRepositoryParams) WithBody(body *models.NpmProxyRepositoryAPIRequest) *CreateNpmProxyRepositoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 10 params
func (o *CreateNpmProxyRepositoryParams) SetBody(body *models.NpmProxyRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNpmProxyRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
