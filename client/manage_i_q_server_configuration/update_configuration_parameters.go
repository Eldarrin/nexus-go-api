// Code generated by go-swagger; DO NOT EDIT.

package manage_i_q_server_configuration

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

// NewUpdateConfigurationParams creates a new UpdateConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateConfigurationParams() *UpdateConfigurationParams {
	return &UpdateConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateConfigurationParamsWithTimeout creates a new UpdateConfigurationParams object
// with the ability to set a timeout on a request.
func NewUpdateConfigurationParamsWithTimeout(timeout time.Duration) *UpdateConfigurationParams {
	return &UpdateConfigurationParams{
		timeout: timeout,
	}
}

// NewUpdateConfigurationParamsWithContext creates a new UpdateConfigurationParams object
// with the ability to set a context for a request.
func NewUpdateConfigurationParamsWithContext(ctx context.Context) *UpdateConfigurationParams {
	return &UpdateConfigurationParams{
		Context: ctx,
	}
}

// NewUpdateConfigurationParamsWithHTTPClient creates a new UpdateConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateConfigurationParamsWithHTTPClient(client *http.Client) *UpdateConfigurationParams {
	return &UpdateConfigurationParams{
		HTTPClient: client,
	}
}

/* UpdateConfigurationParams contains all the parameters to send to the API endpoint
   for the update configuration operation.

   Typically these are written to a http.Request.
*/
type UpdateConfigurationParams struct {

	// Body.
	Body *models.IqConnectionXo

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateConfigurationParams) WithDefaults() *UpdateConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update configuration params
func (o *UpdateConfigurationParams) WithTimeout(timeout time.Duration) *UpdateConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update configuration params
func (o *UpdateConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update configuration params
func (o *UpdateConfigurationParams) WithContext(ctx context.Context) *UpdateConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update configuration params
func (o *UpdateConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update configuration params
func (o *UpdateConfigurationParams) WithHTTPClient(client *http.Client) *UpdateConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update configuration params
func (o *UpdateConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update configuration params
func (o *UpdateConfigurationParams) WithBody(body *models.IqConnectionXo) *UpdateConfigurationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update configuration params
func (o *UpdateConfigurationParams) SetBody(body *models.IqConnectionXo) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
