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

// NewUpdateRubyGemsGroupRepositoryParams creates a new UpdateRubyGemsGroupRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRubyGemsGroupRepositoryParams() *UpdateRubyGemsGroupRepositoryParams {
	return &UpdateRubyGemsGroupRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRubyGemsGroupRepositoryParamsWithTimeout creates a new UpdateRubyGemsGroupRepositoryParams object
// with the ability to set a timeout on a request.
func NewUpdateRubyGemsGroupRepositoryParamsWithTimeout(timeout time.Duration) *UpdateRubyGemsGroupRepositoryParams {
	return &UpdateRubyGemsGroupRepositoryParams{
		timeout: timeout,
	}
}

// NewUpdateRubyGemsGroupRepositoryParamsWithContext creates a new UpdateRubyGemsGroupRepositoryParams object
// with the ability to set a context for a request.
func NewUpdateRubyGemsGroupRepositoryParamsWithContext(ctx context.Context) *UpdateRubyGemsGroupRepositoryParams {
	return &UpdateRubyGemsGroupRepositoryParams{
		Context: ctx,
	}
}

// NewUpdateRubyGemsGroupRepositoryParamsWithHTTPClient creates a new UpdateRubyGemsGroupRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRubyGemsGroupRepositoryParamsWithHTTPClient(client *http.Client) *UpdateRubyGemsGroupRepositoryParams {
	return &UpdateRubyGemsGroupRepositoryParams{
		HTTPClient: client,
	}
}

/*
UpdateRubyGemsGroupRepositoryParams contains all the parameters to send to the API endpoint

	for the update repository 14 operation.

	Typically these are written to a http.Request.
*/
type UpdateRubyGemsGroupRepositoryParams struct {

	// Body.
	Body *models.RubyGemsGroupRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update repository 14 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRubyGemsGroupRepositoryParams) WithDefaults() *UpdateRubyGemsGroupRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update repository 14 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRubyGemsGroupRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update repository 14 params
func (o *UpdateRubyGemsGroupRepositoryParams) WithTimeout(timeout time.Duration) *UpdateRubyGemsGroupRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update repository 14 params
func (o *UpdateRubyGemsGroupRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update repository 14 params
func (o *UpdateRubyGemsGroupRepositoryParams) WithContext(ctx context.Context) *UpdateRubyGemsGroupRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update repository 14 params
func (o *UpdateRubyGemsGroupRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update repository 14 params
func (o *UpdateRubyGemsGroupRepositoryParams) WithHTTPClient(client *http.Client) *UpdateRubyGemsGroupRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update repository 14 params
func (o *UpdateRubyGemsGroupRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update repository 14 params
func (o *UpdateRubyGemsGroupRepositoryParams) WithBody(body *models.RubyGemsGroupRepositoryAPIRequest) *UpdateRubyGemsGroupRepositoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update repository 14 params
func (o *UpdateRubyGemsGroupRepositoryParams) SetBody(body *models.RubyGemsGroupRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the update repository 14 params
func (o *UpdateRubyGemsGroupRepositoryParams) WithRepositoryName(repositoryName string) *UpdateRubyGemsGroupRepositoryParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the update repository 14 params
func (o *UpdateRubyGemsGroupRepositoryParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRubyGemsGroupRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
