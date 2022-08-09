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

// NewUpdateHelmProxyRepositoryParams creates a new UpdateHelmProxyRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateHelmProxyRepositoryParams() *UpdateHelmProxyRepositoryParams {
	return &UpdateHelmProxyRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateHelmProxyRepositoryParamsWithTimeout creates a new UpdateHelmProxyRepositoryParams object
// with the ability to set a timeout on a request.
func NewUpdateHelmProxyRepositoryParamsWithTimeout(timeout time.Duration) *UpdateHelmProxyRepositoryParams {
	return &UpdateHelmProxyRepositoryParams{
		timeout: timeout,
	}
}

// NewUpdateHelmProxyRepositoryParamsWithContext creates a new UpdateHelmProxyRepositoryParams object
// with the ability to set a context for a request.
func NewUpdateHelmProxyRepositoryParamsWithContext(ctx context.Context) *UpdateHelmProxyRepositoryParams {
	return &UpdateHelmProxyRepositoryParams{
		Context: ctx,
	}
}

// NewUpdateHelmProxyRepositoryParamsWithHTTPClient creates a new UpdateHelmProxyRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateHelmProxyRepositoryParamsWithHTTPClient(client *http.Client) *UpdateHelmProxyRepositoryParams {
	return &UpdateHelmProxyRepositoryParams{
		HTTPClient: client,
	}
}

/*
UpdateHelmProxyRepositoryParams contains all the parameters to send to the API endpoint

	for the update repository 24 operation.

	Typically these are written to a http.Request.
*/
type UpdateHelmProxyRepositoryParams struct {

	// Body.
	Body *models.HelmProxyRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update repository 24 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateHelmProxyRepositoryParams) WithDefaults() *UpdateHelmProxyRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update repository 24 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateHelmProxyRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update repository 24 params
func (o *UpdateHelmProxyRepositoryParams) WithTimeout(timeout time.Duration) *UpdateHelmProxyRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update repository 24 params
func (o *UpdateHelmProxyRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update repository 24 params
func (o *UpdateHelmProxyRepositoryParams) WithContext(ctx context.Context) *UpdateHelmProxyRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update repository 24 params
func (o *UpdateHelmProxyRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update repository 24 params
func (o *UpdateHelmProxyRepositoryParams) WithHTTPClient(client *http.Client) *UpdateHelmProxyRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update repository 24 params
func (o *UpdateHelmProxyRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update repository 24 params
func (o *UpdateHelmProxyRepositoryParams) WithBody(body *models.HelmProxyRepositoryAPIRequest) *UpdateHelmProxyRepositoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update repository 24 params
func (o *UpdateHelmProxyRepositoryParams) SetBody(body *models.HelmProxyRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the update repository 24 params
func (o *UpdateHelmProxyRepositoryParams) WithRepositoryName(repositoryName string) *UpdateHelmProxyRepositoryParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the update repository 24 params
func (o *UpdateHelmProxyRepositoryParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateHelmProxyRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param repositoryName
	if err := r.SetPathParam("repositoryName", o.RepositoryName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
