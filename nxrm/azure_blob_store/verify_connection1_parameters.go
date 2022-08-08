// Code generated by go-swagger; DO NOT EDIT.

package azure_blob_store

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

// NewVerifyConnection1Params creates a new VerifyConnection1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVerifyConnection1Params() *VerifyConnection1Params {
	return &VerifyConnection1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewVerifyConnection1ParamsWithTimeout creates a new VerifyConnection1Params object
// with the ability to set a timeout on a request.
func NewVerifyConnection1ParamsWithTimeout(timeout time.Duration) *VerifyConnection1Params {
	return &VerifyConnection1Params{
		timeout: timeout,
	}
}

// NewVerifyConnection1ParamsWithContext creates a new VerifyConnection1Params object
// with the ability to set a context for a request.
func NewVerifyConnection1ParamsWithContext(ctx context.Context) *VerifyConnection1Params {
	return &VerifyConnection1Params{
		Context: ctx,
	}
}

// NewVerifyConnection1ParamsWithHTTPClient creates a new VerifyConnection1Params object
// with the ability to set a custom HTTPClient for a request.
func NewVerifyConnection1ParamsWithHTTPClient(client *http.Client) *VerifyConnection1Params {
	return &VerifyConnection1Params{
		HTTPClient: client,
	}
}

/* VerifyConnection1Params contains all the parameters to send to the API endpoint
   for the verify connection 1 operation.

   Typically these are written to a http.Request.
*/
type VerifyConnection1Params struct {

	// Body.
	Body *models.AzureConnectionXO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the verify connection 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VerifyConnection1Params) WithDefaults() *VerifyConnection1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the verify connection 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VerifyConnection1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the verify connection 1 params
func (o *VerifyConnection1Params) WithTimeout(timeout time.Duration) *VerifyConnection1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the verify connection 1 params
func (o *VerifyConnection1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the verify connection 1 params
func (o *VerifyConnection1Params) WithContext(ctx context.Context) *VerifyConnection1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the verify connection 1 params
func (o *VerifyConnection1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the verify connection 1 params
func (o *VerifyConnection1Params) WithHTTPClient(client *http.Client) *VerifyConnection1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the verify connection 1 params
func (o *VerifyConnection1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the verify connection 1 params
func (o *VerifyConnection1Params) WithBody(body *models.AzureConnectionXO) *VerifyConnection1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the verify connection 1 params
func (o *VerifyConnection1Params) SetBody(body *models.AzureConnectionXO) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *VerifyConnection1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
