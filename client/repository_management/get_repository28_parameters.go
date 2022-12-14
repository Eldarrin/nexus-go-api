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

// NewGetPypiHostedRepositoryParams creates a new GetPypiHostedRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPypiHostedRepositoryParams() *GetPypiHostedRepositoryParams {
	return &GetPypiHostedRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPypiHostedRepositoryParamsWithTimeout creates a new GetPypiHostedRepositoryParams object
// with the ability to set a timeout on a request.
func NewGetPypiHostedRepositoryParamsWithTimeout(timeout time.Duration) *GetPypiHostedRepositoryParams {
	return &GetPypiHostedRepositoryParams{
		timeout: timeout,
	}
}

// NewGetPypiHostedRepositoryParamsWithContext creates a new GetPypiHostedRepositoryParams object
// with the ability to set a context for a request.
func NewGetPypiHostedRepositoryParamsWithContext(ctx context.Context) *GetPypiHostedRepositoryParams {
	return &GetPypiHostedRepositoryParams{
		Context: ctx,
	}
}

// NewGetPypiHostedRepositoryParamsWithHTTPClient creates a new GetPypiHostedRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPypiHostedRepositoryParamsWithHTTPClient(client *http.Client) *GetPypiHostedRepositoryParams {
	return &GetPypiHostedRepositoryParams{
		HTTPClient: client,
	}
}

/*
GetPypiHostedRepositoryParams contains all the parameters to send to the API endpoint

	for the get repository 28 operation.

	Typically these are written to a http.Request.
*/
type GetPypiHostedRepositoryParams struct {

	// RepositoryName.
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repository 28 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPypiHostedRepositoryParams) WithDefaults() *GetPypiHostedRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repository 28 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPypiHostedRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get repository 28 params
func (o *GetPypiHostedRepositoryParams) WithTimeout(timeout time.Duration) *GetPypiHostedRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repository 28 params
func (o *GetPypiHostedRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repository 28 params
func (o *GetPypiHostedRepositoryParams) WithContext(ctx context.Context) *GetPypiHostedRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repository 28 params
func (o *GetPypiHostedRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repository 28 params
func (o *GetPypiHostedRepositoryParams) WithHTTPClient(client *http.Client) *GetPypiHostedRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repository 28 params
func (o *GetPypiHostedRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the get repository 28 params
func (o *GetPypiHostedRepositoryParams) WithRepositoryName(repositoryName string) *GetPypiHostedRepositoryParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get repository 28 params
func (o *GetPypiHostedRepositoryParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetPypiHostedRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
