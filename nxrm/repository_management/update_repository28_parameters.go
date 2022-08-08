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

	"Eldarrin/nexus-go-api/models"
)

// NewUpdateRepository28Params creates a new UpdateRepository28Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRepository28Params() *UpdateRepository28Params {
	return &UpdateRepository28Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRepository28ParamsWithTimeout creates a new UpdateRepository28Params object
// with the ability to set a timeout on a request.
func NewUpdateRepository28ParamsWithTimeout(timeout time.Duration) *UpdateRepository28Params {
	return &UpdateRepository28Params{
		timeout: timeout,
	}
}

// NewUpdateRepository28ParamsWithContext creates a new UpdateRepository28Params object
// with the ability to set a context for a request.
func NewUpdateRepository28ParamsWithContext(ctx context.Context) *UpdateRepository28Params {
	return &UpdateRepository28Params{
		Context: ctx,
	}
}

// NewUpdateRepository28ParamsWithHTTPClient creates a new UpdateRepository28Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRepository28ParamsWithHTTPClient(client *http.Client) *UpdateRepository28Params {
	return &UpdateRepository28Params{
		HTTPClient: client,
	}
}

/* UpdateRepository28Params contains all the parameters to send to the API endpoint
   for the update repository 28 operation.

   Typically these are written to a http.Request.
*/
type UpdateRepository28Params struct {

	// Body.
	Body *models.PypiProxyRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update repository 28 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRepository28Params) WithDefaults() *UpdateRepository28Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update repository 28 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRepository28Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update repository 28 params
func (o *UpdateRepository28Params) WithTimeout(timeout time.Duration) *UpdateRepository28Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update repository 28 params
func (o *UpdateRepository28Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update repository 28 params
func (o *UpdateRepository28Params) WithContext(ctx context.Context) *UpdateRepository28Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update repository 28 params
func (o *UpdateRepository28Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update repository 28 params
func (o *UpdateRepository28Params) WithHTTPClient(client *http.Client) *UpdateRepository28Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update repository 28 params
func (o *UpdateRepository28Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update repository 28 params
func (o *UpdateRepository28Params) WithBody(body *models.PypiProxyRepositoryAPIRequest) *UpdateRepository28Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update repository 28 params
func (o *UpdateRepository28Params) SetBody(body *models.PypiProxyRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the update repository 28 params
func (o *UpdateRepository28Params) WithRepositoryName(repositoryName string) *UpdateRepository28Params {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the update repository 28 params
func (o *UpdateRepository28Params) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRepository28Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
