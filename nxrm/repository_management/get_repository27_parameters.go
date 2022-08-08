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

// NewGetRepository27Params creates a new GetRepository27Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRepository27Params() *GetRepository27Params {
	return &GetRepository27Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepository27ParamsWithTimeout creates a new GetRepository27Params object
// with the ability to set a timeout on a request.
func NewGetRepository27ParamsWithTimeout(timeout time.Duration) *GetRepository27Params {
	return &GetRepository27Params{
		timeout: timeout,
	}
}

// NewGetRepository27ParamsWithContext creates a new GetRepository27Params object
// with the ability to set a context for a request.
func NewGetRepository27ParamsWithContext(ctx context.Context) *GetRepository27Params {
	return &GetRepository27Params{
		Context: ctx,
	}
}

// NewGetRepository27ParamsWithHTTPClient creates a new GetRepository27Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetRepository27ParamsWithHTTPClient(client *http.Client) *GetRepository27Params {
	return &GetRepository27Params{
		HTTPClient: client,
	}
}

/* GetRepository27Params contains all the parameters to send to the API endpoint
   for the get repository 27 operation.

   Typically these are written to a http.Request.
*/
type GetRepository27Params struct {

	// RepositoryName.
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repository 27 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepository27Params) WithDefaults() *GetRepository27Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repository 27 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepository27Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get repository 27 params
func (o *GetRepository27Params) WithTimeout(timeout time.Duration) *GetRepository27Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repository 27 params
func (o *GetRepository27Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repository 27 params
func (o *GetRepository27Params) WithContext(ctx context.Context) *GetRepository27Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repository 27 params
func (o *GetRepository27Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repository 27 params
func (o *GetRepository27Params) WithHTTPClient(client *http.Client) *GetRepository27Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repository 27 params
func (o *GetRepository27Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the get repository 27 params
func (o *GetRepository27Params) WithRepositoryName(repositoryName string) *GetRepository27Params {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get repository 27 params
func (o *GetRepository27Params) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepository27Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
