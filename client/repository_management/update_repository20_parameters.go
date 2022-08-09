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

// NewUpdateYumGroupRepositoryParams creates a new UpdateYumGroupRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateYumGroupRepositoryParams() *UpdateYumGroupRepositoryParams {
	return &UpdateYumGroupRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateYumGroupRepositoryParamsWithTimeout creates a new UpdateYumGroupRepositoryParams object
// with the ability to set a timeout on a request.
func NewUpdateYumGroupRepositoryParamsWithTimeout(timeout time.Duration) *UpdateYumGroupRepositoryParams {
	return &UpdateYumGroupRepositoryParams{
		timeout: timeout,
	}
}

// NewUpdateYumGroupRepositoryParamsWithContext creates a new UpdateYumGroupRepositoryParams object
// with the ability to set a context for a request.
func NewUpdateYumGroupRepositoryParamsWithContext(ctx context.Context) *UpdateYumGroupRepositoryParams {
	return &UpdateYumGroupRepositoryParams{
		Context: ctx,
	}
}

// NewUpdateYumGroupRepositoryParamsWithHTTPClient creates a new UpdateYumGroupRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateYumGroupRepositoryParamsWithHTTPClient(client *http.Client) *UpdateYumGroupRepositoryParams {
	return &UpdateYumGroupRepositoryParams{
		HTTPClient: client,
	}
}

/*
UpdateYumGroupRepositoryParams contains all the parameters to send to the API endpoint

	for the update repository 20 operation.

	Typically these are written to a http.Request.
*/
type UpdateYumGroupRepositoryParams struct {

	// Body.
	Body *models.YumGroupRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update repository 20 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateYumGroupRepositoryParams) WithDefaults() *UpdateYumGroupRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update repository 20 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateYumGroupRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update repository 20 params
func (o *UpdateYumGroupRepositoryParams) WithTimeout(timeout time.Duration) *UpdateYumGroupRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update repository 20 params
func (o *UpdateYumGroupRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update repository 20 params
func (o *UpdateYumGroupRepositoryParams) WithContext(ctx context.Context) *UpdateYumGroupRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update repository 20 params
func (o *UpdateYumGroupRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update repository 20 params
func (o *UpdateYumGroupRepositoryParams) WithHTTPClient(client *http.Client) *UpdateYumGroupRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update repository 20 params
func (o *UpdateYumGroupRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update repository 20 params
func (o *UpdateYumGroupRepositoryParams) WithBody(body *models.YumGroupRepositoryAPIRequest) *UpdateYumGroupRepositoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update repository 20 params
func (o *UpdateYumGroupRepositoryParams) SetBody(body *models.YumGroupRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the update repository 20 params
func (o *UpdateYumGroupRepositoryParams) WithRepositoryName(repositoryName string) *UpdateYumGroupRepositoryParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the update repository 20 params
func (o *UpdateYumGroupRepositoryParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateYumGroupRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
