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

// NewGetMavenHostedRepositoryParams creates a new GetMavenHostedRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMavenHostedRepositoryParams() *GetMavenHostedRepositoryParams {
	return &GetMavenHostedRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMavenHostedRepositoryParamsWithTimeout creates a new GetMavenHostedRepositoryParams object
// with the ability to set a timeout on a request.
func NewGetMavenHostedRepositoryParamsWithTimeout(timeout time.Duration) *GetMavenHostedRepositoryParams {
	return &GetMavenHostedRepositoryParams{
		timeout: timeout,
	}
}

// NewGetMavenHostedRepositoryParamsWithContext creates a new GetMavenHostedRepositoryParams object
// with the ability to set a context for a request.
func NewGetMavenHostedRepositoryParamsWithContext(ctx context.Context) *GetMavenHostedRepositoryParams {
	return &GetMavenHostedRepositoryParams{
		Context: ctx,
	}
}

// NewGetMavenHostedRepositoryParamsWithHTTPClient creates a new GetMavenHostedRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMavenHostedRepositoryParamsWithHTTPClient(client *http.Client) *GetMavenHostedRepositoryParams {
	return &GetMavenHostedRepositoryParams{
		HTTPClient: client,
	}
}

/*
GetMavenHostedRepositoryParams contains all the parameters to send to the API endpoint

	for the get repository 2 operation.

	Typically these are written to a http.Request.
*/
type GetMavenHostedRepositoryParams struct {

	// RepositoryName.
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repository 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMavenHostedRepositoryParams) WithDefaults() *GetMavenHostedRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repository 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMavenHostedRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get repository 2 params
func (o *GetMavenHostedRepositoryParams) WithTimeout(timeout time.Duration) *GetMavenHostedRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repository 2 params
func (o *GetMavenHostedRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repository 2 params
func (o *GetMavenHostedRepositoryParams) WithContext(ctx context.Context) *GetMavenHostedRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repository 2 params
func (o *GetMavenHostedRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repository 2 params
func (o *GetMavenHostedRepositoryParams) WithHTTPClient(client *http.Client) *GetMavenHostedRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repository 2 params
func (o *GetMavenHostedRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the get repository 2 params
func (o *GetMavenHostedRepositoryParams) WithRepositoryName(repositoryName string) *GetMavenHostedRepositoryParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get repository 2 params
func (o *GetMavenHostedRepositoryParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetMavenHostedRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
