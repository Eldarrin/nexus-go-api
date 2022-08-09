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

// NewUpdateGolangGroupRepositoryParams creates a new UpdateGolangGroupRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateGolangGroupRepositoryParams() *UpdateGolangGroupRepositoryParams {
	return &UpdateGolangGroupRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateGolangGroupRepositoryParamsWithTimeout creates a new UpdateGolangGroupRepositoryParams object
// with the ability to set a timeout on a request.
func NewUpdateGolangGroupRepositoryParamsWithTimeout(timeout time.Duration) *UpdateGolangGroupRepositoryParams {
	return &UpdateGolangGroupRepositoryParams{
		timeout: timeout,
	}
}

// NewUpdateGolangGroupRepositoryParamsWithContext creates a new UpdateGolangGroupRepositoryParams object
// with the ability to set a context for a request.
func NewUpdateGolangGroupRepositoryParamsWithContext(ctx context.Context) *UpdateGolangGroupRepositoryParams {
	return &UpdateGolangGroupRepositoryParams{
		Context: ctx,
	}
}

// NewUpdateGolangGroupRepositoryParamsWithHTTPClient creates a new UpdateGolangGroupRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateGolangGroupRepositoryParamsWithHTTPClient(client *http.Client) *UpdateGolangGroupRepositoryParams {
	return &UpdateGolangGroupRepositoryParams{
		HTTPClient: client,
	}
}

/*
UpdateGolangGroupRepositoryParams contains all the parameters to send to the API endpoint

	for the update repository 35 operation.

	Typically these are written to a http.Request.
*/
type UpdateGolangGroupRepositoryParams struct {

	// Body.
	Body *models.GolangGroupRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update repository 35 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGolangGroupRepositoryParams) WithDefaults() *UpdateGolangGroupRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update repository 35 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGolangGroupRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update repository 35 params
func (o *UpdateGolangGroupRepositoryParams) WithTimeout(timeout time.Duration) *UpdateGolangGroupRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update repository 35 params
func (o *UpdateGolangGroupRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update repository 35 params
func (o *UpdateGolangGroupRepositoryParams) WithContext(ctx context.Context) *UpdateGolangGroupRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update repository 35 params
func (o *UpdateGolangGroupRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update repository 35 params
func (o *UpdateGolangGroupRepositoryParams) WithHTTPClient(client *http.Client) *UpdateGolangGroupRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update repository 35 params
func (o *UpdateGolangGroupRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update repository 35 params
func (o *UpdateGolangGroupRepositoryParams) WithBody(body *models.GolangGroupRepositoryAPIRequest) *UpdateGolangGroupRepositoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update repository 35 params
func (o *UpdateGolangGroupRepositoryParams) SetBody(body *models.GolangGroupRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the update repository 35 params
func (o *UpdateGolangGroupRepositoryParams) WithRepositoryName(repositoryName string) *UpdateGolangGroupRepositoryParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the update repository 35 params
func (o *UpdateGolangGroupRepositoryParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateGolangGroupRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
