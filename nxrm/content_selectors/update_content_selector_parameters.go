// Code generated by go-swagger; DO NOT EDIT.

package content_selectors

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

// NewUpdateContentSelectorParams creates a new UpdateContentSelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateContentSelectorParams() *UpdateContentSelectorParams {
	return &UpdateContentSelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateContentSelectorParamsWithTimeout creates a new UpdateContentSelectorParams object
// with the ability to set a timeout on a request.
func NewUpdateContentSelectorParamsWithTimeout(timeout time.Duration) *UpdateContentSelectorParams {
	return &UpdateContentSelectorParams{
		timeout: timeout,
	}
}

// NewUpdateContentSelectorParamsWithContext creates a new UpdateContentSelectorParams object
// with the ability to set a context for a request.
func NewUpdateContentSelectorParamsWithContext(ctx context.Context) *UpdateContentSelectorParams {
	return &UpdateContentSelectorParams{
		Context: ctx,
	}
}

// NewUpdateContentSelectorParamsWithHTTPClient creates a new UpdateContentSelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateContentSelectorParamsWithHTTPClient(client *http.Client) *UpdateContentSelectorParams {
	return &UpdateContentSelectorParams{
		HTTPClient: client,
	}
}

/* UpdateContentSelectorParams contains all the parameters to send to the API endpoint
   for the update content selector operation.

   Typically these are written to a http.Request.
*/
type UpdateContentSelectorParams struct {

	// Body.
	Body *models.ContentSelectorAPIUpdateRequest

	/* Name.

	   The content selector name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update content selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateContentSelectorParams) WithDefaults() *UpdateContentSelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update content selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateContentSelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update content selector params
func (o *UpdateContentSelectorParams) WithTimeout(timeout time.Duration) *UpdateContentSelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update content selector params
func (o *UpdateContentSelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update content selector params
func (o *UpdateContentSelectorParams) WithContext(ctx context.Context) *UpdateContentSelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update content selector params
func (o *UpdateContentSelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update content selector params
func (o *UpdateContentSelectorParams) WithHTTPClient(client *http.Client) *UpdateContentSelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update content selector params
func (o *UpdateContentSelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update content selector params
func (o *UpdateContentSelectorParams) WithBody(body *models.ContentSelectorAPIUpdateRequest) *UpdateContentSelectorParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update content selector params
func (o *UpdateContentSelectorParams) SetBody(body *models.ContentSelectorAPIUpdateRequest) {
	o.Body = body
}

// WithName adds the name to the update content selector params
func (o *UpdateContentSelectorParams) WithName(name string) *UpdateContentSelectorParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update content selector params
func (o *UpdateContentSelectorParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateContentSelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
