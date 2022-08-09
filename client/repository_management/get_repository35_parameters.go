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

// NewGetCocoapodsProxyRepositoryParams creates a new GetCocoapodsProxyRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCocoapodsProxyRepositoryParams() *GetCocoapodsProxyRepositoryParams {
	return &GetCocoapodsProxyRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCocoapodsProxyRepositoryParamsWithTimeout creates a new GetCocoapodsProxyRepositoryParams object
// with the ability to set a timeout on a request.
func NewGetCocoapodsProxyRepositoryParamsWithTimeout(timeout time.Duration) *GetCocoapodsProxyRepositoryParams {
	return &GetCocoapodsProxyRepositoryParams{
		timeout: timeout,
	}
}

// NewGetCocoapodsProxyRepositoryParamsWithContext creates a new GetCocoapodsProxyRepositoryParams object
// with the ability to set a context for a request.
func NewGetCocoapodsProxyRepositoryParamsWithContext(ctx context.Context) *GetCocoapodsProxyRepositoryParams {
	return &GetCocoapodsProxyRepositoryParams{
		Context: ctx,
	}
}

// NewGetCocoapodsProxyRepositoryParamsWithHTTPClient creates a new GetCocoapodsProxyRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCocoapodsProxyRepositoryParamsWithHTTPClient(client *http.Client) *GetCocoapodsProxyRepositoryParams {
	return &GetCocoapodsProxyRepositoryParams{
		HTTPClient: client,
	}
}

/*
GetCocoapodsProxyRepositoryParams contains all the parameters to send to the API endpoint

	for the get repository 35 operation.

	Typically these are written to a http.Request.
*/
type GetCocoapodsProxyRepositoryParams struct {

	// RepositoryName.
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repository 35 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCocoapodsProxyRepositoryParams) WithDefaults() *GetCocoapodsProxyRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repository 35 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCocoapodsProxyRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get repository 35 params
func (o *GetCocoapodsProxyRepositoryParams) WithTimeout(timeout time.Duration) *GetCocoapodsProxyRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repository 35 params
func (o *GetCocoapodsProxyRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repository 35 params
func (o *GetCocoapodsProxyRepositoryParams) WithContext(ctx context.Context) *GetCocoapodsProxyRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repository 35 params
func (o *GetCocoapodsProxyRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repository 35 params
func (o *GetCocoapodsProxyRepositoryParams) WithHTTPClient(client *http.Client) *GetCocoapodsProxyRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repository 35 params
func (o *GetCocoapodsProxyRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the get repository 35 params
func (o *GetCocoapodsProxyRepositoryParams) WithRepositoryName(repositoryName string) *GetCocoapodsProxyRepositoryParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get repository 35 params
func (o *GetCocoapodsProxyRepositoryParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetCocoapodsProxyRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
