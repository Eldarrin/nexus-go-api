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

// NewGetMavenProxyRepositoryParams creates a new GetMavenProxyRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMavenProxyRepositoryParams() *GetMavenProxyRepositoryParams {
	return &GetMavenProxyRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMavenProxyRepositoryParamsWithTimeout creates a new GetMavenProxyRepositoryParams object
// with the ability to set a timeout on a request.
func NewGetMavenProxyRepositoryParamsWithTimeout(timeout time.Duration) *GetMavenProxyRepositoryParams {
	return &GetMavenProxyRepositoryParams{
		timeout: timeout,
	}
}

// NewGetMavenProxyRepositoryParamsWithContext creates a new GetMavenProxyRepositoryParams object
// with the ability to set a context for a request.
func NewGetMavenProxyRepositoryParamsWithContext(ctx context.Context) *GetMavenProxyRepositoryParams {
	return &GetMavenProxyRepositoryParams{
		Context: ctx,
	}
}

// NewGetMavenProxyRepositoryParamsWithHTTPClient creates a new GetMavenProxyRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMavenProxyRepositoryParamsWithHTTPClient(client *http.Client) *GetMavenProxyRepositoryParams {
	return &GetMavenProxyRepositoryParams{
		HTTPClient: client,
	}
}

/*
GetMavenProxyRepositoryParams contains all the parameters to send to the API endpoint

	for the get repository 3 operation.

	Typically these are written to a http.Request.
*/
type GetMavenProxyRepositoryParams struct {

	// RepositoryName.
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repository 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMavenProxyRepositoryParams) WithDefaults() *GetMavenProxyRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repository 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMavenProxyRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get repository 3 params
func (o *GetMavenProxyRepositoryParams) WithTimeout(timeout time.Duration) *GetMavenProxyRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repository 3 params
func (o *GetMavenProxyRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repository 3 params
func (o *GetMavenProxyRepositoryParams) WithContext(ctx context.Context) *GetMavenProxyRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repository 3 params
func (o *GetMavenProxyRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repository 3 params
func (o *GetMavenProxyRepositoryParams) WithHTTPClient(client *http.Client) *GetMavenProxyRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repository 3 params
func (o *GetMavenProxyRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the get repository 3 params
func (o *GetMavenProxyRepositoryParams) WithRepositoryName(repositoryName string) *GetMavenProxyRepositoryParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get repository 3 params
func (o *GetMavenProxyRepositoryParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetMavenProxyRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
