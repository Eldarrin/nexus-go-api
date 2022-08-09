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

// NewUpdateGitLfsHostedRepositoryParams creates a new UpdateGitLfsHostedRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateGitLfsHostedRepositoryParams() *UpdateGitLfsHostedRepositoryParams {
	return &UpdateGitLfsHostedRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateGitLfsHostedRepositoryParamsWithTimeout creates a new UpdateGitLfsHostedRepositoryParams object
// with the ability to set a timeout on a request.
func NewUpdateGitLfsHostedRepositoryParamsWithTimeout(timeout time.Duration) *UpdateGitLfsHostedRepositoryParams {
	return &UpdateGitLfsHostedRepositoryParams{
		timeout: timeout,
	}
}

// NewUpdateGitLfsHostedRepositoryParamsWithContext creates a new UpdateGitLfsHostedRepositoryParams object
// with the ability to set a context for a request.
func NewUpdateGitLfsHostedRepositoryParamsWithContext(ctx context.Context) *UpdateGitLfsHostedRepositoryParams {
	return &UpdateGitLfsHostedRepositoryParams{
		Context: ctx,
	}
}

// NewUpdateGitLfsHostedRepositoryParamsWithHTTPClient creates a new UpdateGitLfsHostedRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateGitLfsHostedRepositoryParamsWithHTTPClient(client *http.Client) *UpdateGitLfsHostedRepositoryParams {
	return &UpdateGitLfsHostedRepositoryParams{
		HTTPClient: client,
	}
}

/*
UpdateGitLfsHostedRepositoryParams contains all the parameters to send to the API endpoint

	for the update repository 25 operation.

	Typically these are written to a http.Request.
*/
type UpdateGitLfsHostedRepositoryParams struct {

	// Body.
	Body *models.GitLfsHostedRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update repository 25 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGitLfsHostedRepositoryParams) WithDefaults() *UpdateGitLfsHostedRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update repository 25 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGitLfsHostedRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update repository 25 params
func (o *UpdateGitLfsHostedRepositoryParams) WithTimeout(timeout time.Duration) *UpdateGitLfsHostedRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update repository 25 params
func (o *UpdateGitLfsHostedRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update repository 25 params
func (o *UpdateGitLfsHostedRepositoryParams) WithContext(ctx context.Context) *UpdateGitLfsHostedRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update repository 25 params
func (o *UpdateGitLfsHostedRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update repository 25 params
func (o *UpdateGitLfsHostedRepositoryParams) WithHTTPClient(client *http.Client) *UpdateGitLfsHostedRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update repository 25 params
func (o *UpdateGitLfsHostedRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update repository 25 params
func (o *UpdateGitLfsHostedRepositoryParams) WithBody(body *models.GitLfsHostedRepositoryAPIRequest) *UpdateGitLfsHostedRepositoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update repository 25 params
func (o *UpdateGitLfsHostedRepositoryParams) SetBody(body *models.GitLfsHostedRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the update repository 25 params
func (o *UpdateGitLfsHostedRepositoryParams) WithRepositoryName(repositoryName string) *UpdateGitLfsHostedRepositoryParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the update repository 25 params
func (o *UpdateGitLfsHostedRepositoryParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateGitLfsHostedRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
