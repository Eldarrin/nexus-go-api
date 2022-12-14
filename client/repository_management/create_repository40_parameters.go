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

// NewCreateBowerProxyRepositoryParams creates a new CreateBowerProxyRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateBowerProxyRepositoryParams() *CreateBowerProxyRepositoryParams {
	return &CreateBowerProxyRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBowerProxyRepositoryParamsWithTimeout creates a new CreateBowerProxyRepositoryParams object
// with the ability to set a timeout on a request.
func NewCreateBowerProxyRepositoryParamsWithTimeout(timeout time.Duration) *CreateBowerProxyRepositoryParams {
	return &CreateBowerProxyRepositoryParams{
		timeout: timeout,
	}
}

// NewCreateBowerProxyRepositoryParamsWithContext creates a new CreateBowerProxyRepositoryParams object
// with the ability to set a context for a request.
func NewCreateBowerProxyRepositoryParamsWithContext(ctx context.Context) *CreateBowerProxyRepositoryParams {
	return &CreateBowerProxyRepositoryParams{
		Context: ctx,
	}
}

// NewCreateBowerProxyRepositoryParamsWithHTTPClient creates a new CreateBowerProxyRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateBowerProxyRepositoryParamsWithHTTPClient(client *http.Client) *CreateBowerProxyRepositoryParams {
	return &CreateBowerProxyRepositoryParams{
		HTTPClient: client,
	}
}

/*
CreateBowerProxyRepositoryParams contains all the parameters to send to the API endpoint

	for the create repository 40 operation.

	Typically these are written to a http.Request.
*/
type CreateBowerProxyRepositoryParams struct {

	// Body.
	Body *models.BowerProxyRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 40 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBowerProxyRepositoryParams) WithDefaults() *CreateBowerProxyRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 40 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBowerProxyRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 40 params
func (o *CreateBowerProxyRepositoryParams) WithTimeout(timeout time.Duration) *CreateBowerProxyRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 40 params
func (o *CreateBowerProxyRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 40 params
func (o *CreateBowerProxyRepositoryParams) WithContext(ctx context.Context) *CreateBowerProxyRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 40 params
func (o *CreateBowerProxyRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 40 params
func (o *CreateBowerProxyRepositoryParams) WithHTTPClient(client *http.Client) *CreateBowerProxyRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 40 params
func (o *CreateBowerProxyRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 40 params
func (o *CreateBowerProxyRepositoryParams) WithBody(body *models.BowerProxyRepositoryAPIRequest) *CreateBowerProxyRepositoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 40 params
func (o *CreateBowerProxyRepositoryParams) SetBody(body *models.BowerProxyRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBowerProxyRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
