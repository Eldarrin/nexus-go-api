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

// NewUpdateRawGroupRepositoryParams creates a new UpdateRawGroupRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRawGroupRepositoryParams() *UpdateRawGroupRepositoryParams {
	return &UpdateRawGroupRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRawGroupRepositoryParamsWithTimeout creates a new UpdateRawGroupRepositoryParams object
// with the ability to set a timeout on a request.
func NewUpdateRawGroupRepositoryParamsWithTimeout(timeout time.Duration) *UpdateRawGroupRepositoryParams {
	return &UpdateRawGroupRepositoryParams{
		timeout: timeout,
	}
}

// NewUpdateRawGroupRepositoryParamsWithContext creates a new UpdateRawGroupRepositoryParams object
// with the ability to set a context for a request.
func NewUpdateRawGroupRepositoryParamsWithContext(ctx context.Context) *UpdateRawGroupRepositoryParams {
	return &UpdateRawGroupRepositoryParams{
		Context: ctx,
	}
}

// NewUpdateRawGroupRepositoryParamsWithHTTPClient creates a new UpdateRawGroupRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRawGroupRepositoryParamsWithHTTPClient(client *http.Client) *UpdateRawGroupRepositoryParams {
	return &UpdateRawGroupRepositoryParams{
		HTTPClient: client,
	}
}

/*
UpdateRawGroupRepositoryParams contains all the parameters to send to the API endpoint

	for the update repository 5 operation.

	Typically these are written to a http.Request.
*/
type UpdateRawGroupRepositoryParams struct {

	// Body.
	Body *models.RawGroupRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update repository 5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRawGroupRepositoryParams) WithDefaults() *UpdateRawGroupRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update repository 5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRawGroupRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update repository 5 params
func (o *UpdateRawGroupRepositoryParams) WithTimeout(timeout time.Duration) *UpdateRawGroupRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update repository 5 params
func (o *UpdateRawGroupRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update repository 5 params
func (o *UpdateRawGroupRepositoryParams) WithContext(ctx context.Context) *UpdateRawGroupRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update repository 5 params
func (o *UpdateRawGroupRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update repository 5 params
func (o *UpdateRawGroupRepositoryParams) WithHTTPClient(client *http.Client) *UpdateRawGroupRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update repository 5 params
func (o *UpdateRawGroupRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update repository 5 params
func (o *UpdateRawGroupRepositoryParams) WithBody(body *models.RawGroupRepositoryAPIRequest) *UpdateRawGroupRepositoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update repository 5 params
func (o *UpdateRawGroupRepositoryParams) SetBody(body *models.RawGroupRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the update repository 5 params
func (o *UpdateRawGroupRepositoryParams) WithRepositoryName(repositoryName string) *UpdateRawGroupRepositoryParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the update repository 5 params
func (o *UpdateRawGroupRepositoryParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRawGroupRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
