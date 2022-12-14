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

// NewUpdateMavenHostedRepositoryParams creates a new UpdateMavenHostedRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateMavenHostedRepositoryParams() *UpdateMavenHostedRepositoryParams {
	return &UpdateMavenHostedRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMavenHostedRepositoryParamsWithTimeout creates a new UpdateMavenHostedRepositoryParams object
// with the ability to set a timeout on a request.
func NewUpdateMavenHostedRepositoryParamsWithTimeout(timeout time.Duration) *UpdateMavenHostedRepositoryParams {
	return &UpdateMavenHostedRepositoryParams{
		timeout: timeout,
	}
}

// NewUpdateMavenHostedRepositoryParamsWithContext creates a new UpdateMavenHostedRepositoryParams object
// with the ability to set a context for a request.
func NewUpdateMavenHostedRepositoryParamsWithContext(ctx context.Context) *UpdateMavenHostedRepositoryParams {
	return &UpdateMavenHostedRepositoryParams{
		Context: ctx,
	}
}

// NewUpdateMavenHostedRepositoryParamsWithHTTPClient creates a new UpdateMavenHostedRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateMavenHostedRepositoryParamsWithHTTPClient(client *http.Client) *UpdateMavenHostedRepositoryParams {
	return &UpdateMavenHostedRepositoryParams{
		HTTPClient: client,
	}
}

/*
UpdateMavenHostedRepositoryParams contains all the parameters to send to the API endpoint

	for the update repository 1 operation.

	Typically these are written to a http.Request.
*/
type UpdateMavenHostedRepositoryParams struct {

	// Body.
	Body *models.MavenHostedRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update repository 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMavenHostedRepositoryParams) WithDefaults() *UpdateMavenHostedRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update repository 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMavenHostedRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update repository 1 params
func (o *UpdateMavenHostedRepositoryParams) WithTimeout(timeout time.Duration) *UpdateMavenHostedRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update repository 1 params
func (o *UpdateMavenHostedRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update repository 1 params
func (o *UpdateMavenHostedRepositoryParams) WithContext(ctx context.Context) *UpdateMavenHostedRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update repository 1 params
func (o *UpdateMavenHostedRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update repository 1 params
func (o *UpdateMavenHostedRepositoryParams) WithHTTPClient(client *http.Client) *UpdateMavenHostedRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update repository 1 params
func (o *UpdateMavenHostedRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update repository 1 params
func (o *UpdateMavenHostedRepositoryParams) WithBody(body *models.MavenHostedRepositoryAPIRequest) *UpdateMavenHostedRepositoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update repository 1 params
func (o *UpdateMavenHostedRepositoryParams) SetBody(body *models.MavenHostedRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the update repository 1 params
func (o *UpdateMavenHostedRepositoryParams) WithRepositoryName(repositoryName string) *UpdateMavenHostedRepositoryParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the update repository 1 params
func (o *UpdateMavenHostedRepositoryParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMavenHostedRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
