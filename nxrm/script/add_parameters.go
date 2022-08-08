// Code generated by go-swagger; DO NOT EDIT.

package script

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

	"nexus-go-api/models"
)

// NewAddParams creates a new AddParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddParams() *AddParams {
	return &AddParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddParamsWithTimeout creates a new AddParams object
// with the ability to set a timeout on a request.
func NewAddParamsWithTimeout(timeout time.Duration) *AddParams {
	return &AddParams{
		timeout: timeout,
	}
}

// NewAddParamsWithContext creates a new AddParams object
// with the ability to set a context for a request.
func NewAddParamsWithContext(ctx context.Context) *AddParams {
	return &AddParams{
		Context: ctx,
	}
}

// NewAddParamsWithHTTPClient creates a new AddParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddParamsWithHTTPClient(client *http.Client) *AddParams {
	return &AddParams{
		HTTPClient: client,
	}
}

/* AddParams contains all the parameters to send to the API endpoint
   for the add operation.

   Typically these are written to a http.Request.
*/
type AddParams struct {

	// Body.
	Body *models.ScriptXO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddParams) WithDefaults() *AddParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add params
func (o *AddParams) WithTimeout(timeout time.Duration) *AddParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add params
func (o *AddParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add params
func (o *AddParams) WithContext(ctx context.Context) *AddParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add params
func (o *AddParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add params
func (o *AddParams) WithHTTPClient(client *http.Client) *AddParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add params
func (o *AddParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add params
func (o *AddParams) WithBody(body *models.ScriptXO) *AddParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add params
func (o *AddParams) SetBody(body *models.ScriptXO) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
