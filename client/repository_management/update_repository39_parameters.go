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

// NewUpdateBowerHostedRepositoryParams creates a new UpdateBowerHostedRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateBowerHostedRepositoryParams() *UpdateBowerHostedRepositoryParams {
	return &UpdateBowerHostedRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBowerHostedRepositoryParamsWithTimeout creates a new UpdateBowerHostedRepositoryParams object
// with the ability to set a timeout on a request.
func NewUpdateBowerHostedRepositoryParamsWithTimeout(timeout time.Duration) *UpdateBowerHostedRepositoryParams {
	return &UpdateBowerHostedRepositoryParams{
		timeout: timeout,
	}
}

// NewUpdateBowerHostedRepositoryParamsWithContext creates a new UpdateBowerHostedRepositoryParams object
// with the ability to set a context for a request.
func NewUpdateBowerHostedRepositoryParamsWithContext(ctx context.Context) *UpdateBowerHostedRepositoryParams {
	return &UpdateBowerHostedRepositoryParams{
		Context: ctx,
	}
}

// NewUpdateBowerHostedRepositoryParamsWithHTTPClient creates a new UpdateBowerHostedRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateBowerHostedRepositoryParamsWithHTTPClient(client *http.Client) *UpdateBowerHostedRepositoryParams {
	return &UpdateBowerHostedRepositoryParams{
		HTTPClient: client,
	}
}

/*
UpdateBowerHostedRepositoryParams contains all the parameters to send to the API endpoint

	for the update repository 39 operation.

	Typically these are written to a http.Request.
*/
type UpdateBowerHostedRepositoryParams struct {

	// Body.
	Body *models.BowerHostedRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update repository 39 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBowerHostedRepositoryParams) WithDefaults() *UpdateBowerHostedRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update repository 39 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBowerHostedRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update repository 39 params
func (o *UpdateBowerHostedRepositoryParams) WithTimeout(timeout time.Duration) *UpdateBowerHostedRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update repository 39 params
func (o *UpdateBowerHostedRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update repository 39 params
func (o *UpdateBowerHostedRepositoryParams) WithContext(ctx context.Context) *UpdateBowerHostedRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update repository 39 params
func (o *UpdateBowerHostedRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update repository 39 params
func (o *UpdateBowerHostedRepositoryParams) WithHTTPClient(client *http.Client) *UpdateBowerHostedRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update repository 39 params
func (o *UpdateBowerHostedRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update repository 39 params
func (o *UpdateBowerHostedRepositoryParams) WithBody(body *models.BowerHostedRepositoryAPIRequest) *UpdateBowerHostedRepositoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update repository 39 params
func (o *UpdateBowerHostedRepositoryParams) SetBody(body *models.BowerHostedRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the update repository 39 params
func (o *UpdateBowerHostedRepositoryParams) WithRepositoryName(repositoryName string) *UpdateBowerHostedRepositoryParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the update repository 39 params
func (o *UpdateBowerHostedRepositoryParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBowerHostedRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
