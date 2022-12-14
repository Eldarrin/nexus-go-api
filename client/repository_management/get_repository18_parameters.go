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

// NewGetDockerGroupRepositoryParams creates a new GetDockerGroupRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDockerGroupRepositoryParams() *GetDockerGroupRepositoryParams {
	return &GetDockerGroupRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDockerGroupRepositoryParamsWithTimeout creates a new GetDockerGroupRepositoryParams object
// with the ability to set a timeout on a request.
func NewGetDockerGroupRepositoryParamsWithTimeout(timeout time.Duration) *GetDockerGroupRepositoryParams {
	return &GetDockerGroupRepositoryParams{
		timeout: timeout,
	}
}

// NewGetDockerGroupRepositoryParamsWithContext creates a new GetDockerGroupRepositoryParams object
// with the ability to set a context for a request.
func NewGetDockerGroupRepositoryParamsWithContext(ctx context.Context) *GetDockerGroupRepositoryParams {
	return &GetDockerGroupRepositoryParams{
		Context: ctx,
	}
}

// NewGetDockerGroupRepositoryParamsWithHTTPClient creates a new GetDockerGroupRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDockerGroupRepositoryParamsWithHTTPClient(client *http.Client) *GetDockerGroupRepositoryParams {
	return &GetDockerGroupRepositoryParams{
		HTTPClient: client,
	}
}

/*
GetDockerGroupRepositoryParams contains all the parameters to send to the API endpoint

	for the get repository 18 operation.

	Typically these are written to a http.Request.
*/
type GetDockerGroupRepositoryParams struct {

	// RepositoryName.
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repository 18 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDockerGroupRepositoryParams) WithDefaults() *GetDockerGroupRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repository 18 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDockerGroupRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get repository 18 params
func (o *GetDockerGroupRepositoryParams) WithTimeout(timeout time.Duration) *GetDockerGroupRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repository 18 params
func (o *GetDockerGroupRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repository 18 params
func (o *GetDockerGroupRepositoryParams) WithContext(ctx context.Context) *GetDockerGroupRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repository 18 params
func (o *GetDockerGroupRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repository 18 params
func (o *GetDockerGroupRepositoryParams) WithHTTPClient(client *http.Client) *GetDockerGroupRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repository 18 params
func (o *GetDockerGroupRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the get repository 18 params
func (o *GetDockerGroupRepositoryParams) WithRepositoryName(repositoryName string) *GetDockerGroupRepositoryParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get repository 18 params
func (o *GetDockerGroupRepositoryParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetDockerGroupRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
