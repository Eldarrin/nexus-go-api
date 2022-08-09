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

// NewCreateRubyGemsProxyRepositoryParams creates a new CreateRubyGemsProxyRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRubyGemsProxyRepositoryParams() *CreateRubyGemsProxyRepositoryParams {
	return &CreateRubyGemsProxyRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRubyGemsProxyRepositoryParamsWithTimeout creates a new CreateRubyGemsProxyRepositoryParams object
// with the ability to set a timeout on a request.
func NewCreateRubyGemsProxyRepositoryParamsWithTimeout(timeout time.Duration) *CreateRubyGemsProxyRepositoryParams {
	return &CreateRubyGemsProxyRepositoryParams{
		timeout: timeout,
	}
}

// NewCreateRubyGemsProxyRepositoryParamsWithContext creates a new CreateRubyGemsProxyRepositoryParams object
// with the ability to set a context for a request.
func NewCreateRubyGemsProxyRepositoryParamsWithContext(ctx context.Context) *CreateRubyGemsProxyRepositoryParams {
	return &CreateRubyGemsProxyRepositoryParams{
		Context: ctx,
	}
}

// NewCreateRubyGemsProxyRepositoryParamsWithHTTPClient creates a new CreateRubyGemsProxyRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRubyGemsProxyRepositoryParamsWithHTTPClient(client *http.Client) *CreateRubyGemsProxyRepositoryParams {
	return &CreateRubyGemsProxyRepositoryParams{
		HTTPClient: client,
	}
}

/*
CreateRubyGemsProxyRepositoryParams contains all the parameters to send to the API endpoint

	for the create repository 16 operation.

	Typically these are written to a http.Request.
*/
type CreateRubyGemsProxyRepositoryParams struct {

	// Body.
	Body *models.RubyGemsProxyRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 16 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRubyGemsProxyRepositoryParams) WithDefaults() *CreateRubyGemsProxyRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 16 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRubyGemsProxyRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 16 params
func (o *CreateRubyGemsProxyRepositoryParams) WithTimeout(timeout time.Duration) *CreateRubyGemsProxyRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 16 params
func (o *CreateRubyGemsProxyRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 16 params
func (o *CreateRubyGemsProxyRepositoryParams) WithContext(ctx context.Context) *CreateRubyGemsProxyRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 16 params
func (o *CreateRubyGemsProxyRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 16 params
func (o *CreateRubyGemsProxyRepositoryParams) WithHTTPClient(client *http.Client) *CreateRubyGemsProxyRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 16 params
func (o *CreateRubyGemsProxyRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 16 params
func (o *CreateRubyGemsProxyRepositoryParams) WithBody(body *models.RubyGemsProxyRepositoryAPIRequest) *CreateRubyGemsProxyRepositoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 16 params
func (o *CreateRubyGemsProxyRepositoryParams) SetBody(body *models.RubyGemsProxyRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRubyGemsProxyRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
