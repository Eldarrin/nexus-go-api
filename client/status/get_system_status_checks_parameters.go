// Code generated by go-swagger; DO NOT EDIT.

package status

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
)

// NewGetSystemStatusChecksParams creates a new GetSystemStatusChecksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSystemStatusChecksParams() *GetSystemStatusChecksParams {
	return &GetSystemStatusChecksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemStatusChecksParamsWithTimeout creates a new GetSystemStatusChecksParams object
// with the ability to set a timeout on a request.
func NewGetSystemStatusChecksParamsWithTimeout(timeout time.Duration) *GetSystemStatusChecksParams {
	return &GetSystemStatusChecksParams{
		timeout: timeout,
	}
}

// NewGetSystemStatusChecksParamsWithContext creates a new GetSystemStatusChecksParams object
// with the ability to set a context for a request.
func NewGetSystemStatusChecksParamsWithContext(ctx context.Context) *GetSystemStatusChecksParams {
	return &GetSystemStatusChecksParams{
		Context: ctx,
	}
}

// NewGetSystemStatusChecksParamsWithHTTPClient creates a new GetSystemStatusChecksParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSystemStatusChecksParamsWithHTTPClient(client *http.Client) *GetSystemStatusChecksParams {
	return &GetSystemStatusChecksParams{
		HTTPClient: client,
	}
}

/*
GetSystemStatusChecksParams contains all the parameters to send to the API endpoint

	for the get system status checks operation.

	Typically these are written to a http.Request.
*/
type GetSystemStatusChecksParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get system status checks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemStatusChecksParams) WithDefaults() *GetSystemStatusChecksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get system status checks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemStatusChecksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get system status checks params
func (o *GetSystemStatusChecksParams) WithTimeout(timeout time.Duration) *GetSystemStatusChecksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system status checks params
func (o *GetSystemStatusChecksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system status checks params
func (o *GetSystemStatusChecksParams) WithContext(ctx context.Context) *GetSystemStatusChecksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system status checks params
func (o *GetSystemStatusChecksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system status checks params
func (o *GetSystemStatusChecksParams) WithHTTPClient(client *http.Client) *GetSystemStatusChecksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system status checks params
func (o *GetSystemStatusChecksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemStatusChecksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
