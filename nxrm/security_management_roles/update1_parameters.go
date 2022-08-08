// Code generated by go-swagger; DO NOT EDIT.

package security_management_roles

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

// NewUpdate1Params creates a new Update1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdate1Params() *Update1Params {
	return &Update1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdate1ParamsWithTimeout creates a new Update1Params object
// with the ability to set a timeout on a request.
func NewUpdate1ParamsWithTimeout(timeout time.Duration) *Update1Params {
	return &Update1Params{
		timeout: timeout,
	}
}

// NewUpdate1ParamsWithContext creates a new Update1Params object
// with the ability to set a context for a request.
func NewUpdate1ParamsWithContext(ctx context.Context) *Update1Params {
	return &Update1Params{
		Context: ctx,
	}
}

// NewUpdate1ParamsWithHTTPClient creates a new Update1Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdate1ParamsWithHTTPClient(client *http.Client) *Update1Params {
	return &Update1Params{
		HTTPClient: client,
	}
}

/* Update1Params contains all the parameters to send to the API endpoint
   for the update 1 operation.

   Typically these are written to a http.Request.
*/
type Update1Params struct {

	/* Body.

	   A role configuration
	*/
	Body *models.RoleXORequest

	/* ID.

	   The id of the role to update
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Update1Params) WithDefaults() *Update1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Update1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update 1 params
func (o *Update1Params) WithTimeout(timeout time.Duration) *Update1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update 1 params
func (o *Update1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update 1 params
func (o *Update1Params) WithContext(ctx context.Context) *Update1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update 1 params
func (o *Update1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update 1 params
func (o *Update1Params) WithHTTPClient(client *http.Client) *Update1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update 1 params
func (o *Update1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update 1 params
func (o *Update1Params) WithBody(body *models.RoleXORequest) *Update1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update 1 params
func (o *Update1Params) SetBody(body *models.RoleXORequest) {
	o.Body = body
}

// WithID adds the id to the update 1 params
func (o *Update1Params) WithID(id string) *Update1Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the update 1 params
func (o *Update1Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *Update1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
