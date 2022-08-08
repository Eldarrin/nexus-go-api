// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

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

// NewCreatePrivilege2Params creates a new CreatePrivilege2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePrivilege2Params() *CreatePrivilege2Params {
	return &CreatePrivilege2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePrivilege2ParamsWithTimeout creates a new CreatePrivilege2Params object
// with the ability to set a timeout on a request.
func NewCreatePrivilege2ParamsWithTimeout(timeout time.Duration) *CreatePrivilege2Params {
	return &CreatePrivilege2Params{
		timeout: timeout,
	}
}

// NewCreatePrivilege2ParamsWithContext creates a new CreatePrivilege2Params object
// with the ability to set a context for a request.
func NewCreatePrivilege2ParamsWithContext(ctx context.Context) *CreatePrivilege2Params {
	return &CreatePrivilege2Params{
		Context: ctx,
	}
}

// NewCreatePrivilege2ParamsWithHTTPClient creates a new CreatePrivilege2Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePrivilege2ParamsWithHTTPClient(client *http.Client) *CreatePrivilege2Params {
	return &CreatePrivilege2Params{
		HTTPClient: client,
	}
}

/* CreatePrivilege2Params contains all the parameters to send to the API endpoint
   for the create privilege 2 operation.

   Typically these are written to a http.Request.
*/
type CreatePrivilege2Params struct {

	/* Body.

	   The privilege to create.
	*/
	Body *models.APIPrivilegeRepositoryContentSelectorRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create privilege 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePrivilege2Params) WithDefaults() *CreatePrivilege2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create privilege 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePrivilege2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create privilege 2 params
func (o *CreatePrivilege2Params) WithTimeout(timeout time.Duration) *CreatePrivilege2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create privilege 2 params
func (o *CreatePrivilege2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create privilege 2 params
func (o *CreatePrivilege2Params) WithContext(ctx context.Context) *CreatePrivilege2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create privilege 2 params
func (o *CreatePrivilege2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create privilege 2 params
func (o *CreatePrivilege2Params) WithHTTPClient(client *http.Client) *CreatePrivilege2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create privilege 2 params
func (o *CreatePrivilege2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create privilege 2 params
func (o *CreatePrivilege2Params) WithBody(body *models.APIPrivilegeRepositoryContentSelectorRequest) *CreatePrivilege2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create privilege 2 params
func (o *CreatePrivilege2Params) SetBody(body *models.APIPrivilegeRepositoryContentSelectorRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePrivilege2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
