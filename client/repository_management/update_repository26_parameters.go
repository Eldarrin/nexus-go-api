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

// NewUpdateRepository26Params creates a new UpdateRepository26Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRepository26Params() *UpdateRepository26Params {
	return &UpdateRepository26Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRepository26ParamsWithTimeout creates a new UpdateRepository26Params object
// with the ability to set a timeout on a request.
func NewUpdateRepository26ParamsWithTimeout(timeout time.Duration) *UpdateRepository26Params {
	return &UpdateRepository26Params{
		timeout: timeout,
	}
}

// NewUpdateRepository26ParamsWithContext creates a new UpdateRepository26Params object
// with the ability to set a context for a request.
func NewUpdateRepository26ParamsWithContext(ctx context.Context) *UpdateRepository26Params {
	return &UpdateRepository26Params{
		Context: ctx,
	}
}

// NewUpdateRepository26ParamsWithHTTPClient creates a new UpdateRepository26Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRepository26ParamsWithHTTPClient(client *http.Client) *UpdateRepository26Params {
	return &UpdateRepository26Params{
		HTTPClient: client,
	}
}

/* UpdateRepository26Params contains all the parameters to send to the API endpoint
   for the update repository 26 operation.

   Typically these are written to a http.Request.
*/
type UpdateRepository26Params struct {

	// Body.
	Body *models.PypiGroupRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update repository 26 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRepository26Params) WithDefaults() *UpdateRepository26Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update repository 26 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRepository26Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update repository 26 params
func (o *UpdateRepository26Params) WithTimeout(timeout time.Duration) *UpdateRepository26Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update repository 26 params
func (o *UpdateRepository26Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update repository 26 params
func (o *UpdateRepository26Params) WithContext(ctx context.Context) *UpdateRepository26Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update repository 26 params
func (o *UpdateRepository26Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update repository 26 params
func (o *UpdateRepository26Params) WithHTTPClient(client *http.Client) *UpdateRepository26Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update repository 26 params
func (o *UpdateRepository26Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update repository 26 params
func (o *UpdateRepository26Params) WithBody(body *models.PypiGroupRepositoryAPIRequest) *UpdateRepository26Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update repository 26 params
func (o *UpdateRepository26Params) SetBody(body *models.PypiGroupRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the update repository 26 params
func (o *UpdateRepository26Params) WithRepositoryName(repositoryName string) *UpdateRepository26Params {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the update repository 26 params
func (o *UpdateRepository26Params) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRepository26Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
