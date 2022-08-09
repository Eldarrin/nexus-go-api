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

// NewCreateGitLfsHostedRepositoryParams creates a new CreateGitLfsHostedRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateGitLfsHostedRepositoryParams() *CreateGitLfsHostedRepositoryParams {
	return &CreateGitLfsHostedRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateGitLfsHostedRepositoryParamsWithTimeout creates a new CreateGitLfsHostedRepositoryParams object
// with the ability to set a timeout on a request.
func NewCreateGitLfsHostedRepositoryParamsWithTimeout(timeout time.Duration) *CreateGitLfsHostedRepositoryParams {
	return &CreateGitLfsHostedRepositoryParams{
		timeout: timeout,
	}
}

// NewCreateGitLfsHostedRepositoryParamsWithContext creates a new CreateGitLfsHostedRepositoryParams object
// with the ability to set a context for a request.
func NewCreateGitLfsHostedRepositoryParamsWithContext(ctx context.Context) *CreateGitLfsHostedRepositoryParams {
	return &CreateGitLfsHostedRepositoryParams{
		Context: ctx,
	}
}

// NewCreateGitLfsHostedRepositoryParamsWithHTTPClient creates a new CreateGitLfsHostedRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateGitLfsHostedRepositoryParamsWithHTTPClient(client *http.Client) *CreateGitLfsHostedRepositoryParams {
	return &CreateGitLfsHostedRepositoryParams{
		HTTPClient: client,
	}
}

/*
CreateGitLfsHostedRepositoryParams contains all the parameters to send to the API endpoint

	for the create repository 25 operation.

	Typically these are written to a http.Request.
*/
type CreateGitLfsHostedRepositoryParams struct {

	// Body.
	Body *models.GitLfsHostedRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 25 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGitLfsHostedRepositoryParams) WithDefaults() *CreateGitLfsHostedRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 25 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGitLfsHostedRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 25 params
func (o *CreateGitLfsHostedRepositoryParams) WithTimeout(timeout time.Duration) *CreateGitLfsHostedRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 25 params
func (o *CreateGitLfsHostedRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 25 params
func (o *CreateGitLfsHostedRepositoryParams) WithContext(ctx context.Context) *CreateGitLfsHostedRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 25 params
func (o *CreateGitLfsHostedRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 25 params
func (o *CreateGitLfsHostedRepositoryParams) WithHTTPClient(client *http.Client) *CreateGitLfsHostedRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 25 params
func (o *CreateGitLfsHostedRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 25 params
func (o *CreateGitLfsHostedRepositoryParams) WithBody(body *models.GitLfsHostedRepositoryAPIRequest) *CreateGitLfsHostedRepositoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 25 params
func (o *CreateGitLfsHostedRepositoryParams) SetBody(body *models.GitLfsHostedRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateGitLfsHostedRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
