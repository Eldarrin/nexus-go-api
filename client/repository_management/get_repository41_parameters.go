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
)

// NewGetBowerProxyRepositoryParams creates a new GetBowerProxyRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBowerProxyRepositoryParams() *GetBowerProxyRepositoryParams {
	return &GetBowerProxyRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBowerProxyRepositoryParamsWithTimeout creates a new GetBowerProxyRepositoryParams object
// with the ability to set a timeout on a request.
func NewGetBowerProxyRepositoryParamsWithTimeout(timeout time.Duration) *GetBowerProxyRepositoryParams {
	return &GetBowerProxyRepositoryParams{
		timeout: timeout,
	}
}

// NewGetBowerProxyRepositoryParamsWithContext creates a new GetBowerProxyRepositoryParams object
// with the ability to set a context for a request.
func NewGetBowerProxyRepositoryParamsWithContext(ctx context.Context) *GetBowerProxyRepositoryParams {
	return &GetBowerProxyRepositoryParams{
		Context: ctx,
	}
}

// NewGetBowerProxyRepositoryParamsWithHTTPClient creates a new GetBowerProxyRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBowerProxyRepositoryParamsWithHTTPClient(client *http.Client) *GetBowerProxyRepositoryParams {
	return &GetBowerProxyRepositoryParams{
		HTTPClient: client,
	}
}

/*
GetBowerProxyRepositoryParams contains all the parameters to send to the API endpoint

	for the get repository 41 operation.

	Typically these are written to a http.Request.
*/
type GetBowerProxyRepositoryParams struct {

	// RepositoryName.
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repository 41 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBowerProxyRepositoryParams) WithDefaults() *GetBowerProxyRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repository 41 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBowerProxyRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get repository 41 params
func (o *GetBowerProxyRepositoryParams) WithTimeout(timeout time.Duration) *GetBowerProxyRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repository 41 params
func (o *GetBowerProxyRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repository 41 params
func (o *GetBowerProxyRepositoryParams) WithContext(ctx context.Context) *GetBowerProxyRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repository 41 params
func (o *GetBowerProxyRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repository 41 params
func (o *GetBowerProxyRepositoryParams) WithHTTPClient(client *http.Client) *GetBowerProxyRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repository 41 params
func (o *GetBowerProxyRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the get repository 41 params
func (o *GetBowerProxyRepositoryParams) WithRepositoryName(repositoryName string) *GetBowerProxyRepositoryParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get repository 41 params
func (o *GetBowerProxyRepositoryParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetBowerProxyRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repositoryName
	if err := r.SetPathParam("repositoryName", o.RepositoryName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
