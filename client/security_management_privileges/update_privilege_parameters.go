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

// NewUpdatePrivilegeParams creates a new UpdatePrivilegeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdatePrivilegeParams() *UpdatePrivilegeParams {
	return &UpdatePrivilegeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePrivilegeParamsWithTimeout creates a new UpdatePrivilegeParams object
// with the ability to set a timeout on a request.
func NewUpdatePrivilegeParamsWithTimeout(timeout time.Duration) *UpdatePrivilegeParams {
	return &UpdatePrivilegeParams{
		timeout: timeout,
	}
}

// NewUpdatePrivilegeParamsWithContext creates a new UpdatePrivilegeParams object
// with the ability to set a context for a request.
func NewUpdatePrivilegeParamsWithContext(ctx context.Context) *UpdatePrivilegeParams {
	return &UpdatePrivilegeParams{
		Context: ctx,
	}
}

// NewUpdatePrivilegeParamsWithHTTPClient creates a new UpdatePrivilegeParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdatePrivilegeParamsWithHTTPClient(client *http.Client) *UpdatePrivilegeParams {
	return &UpdatePrivilegeParams{
		HTTPClient: client,
	}
}

/*
UpdatePrivilegeParams contains all the parameters to send to the API endpoint

	for the update privilege operation.

	Typically these are written to a http.Request.
*/
type UpdatePrivilegeParams struct {

	/* Body.

	   The privilege to update.
	*/
	Body *models.APIPrivilegeApplicationRequest

	/* PrivilegeName.

	   The name of the privilege to update.
	*/
	PrivilegeName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update privilege params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePrivilegeParams) WithDefaults() *UpdatePrivilegeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update privilege params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePrivilegeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update privilege params
func (o *UpdatePrivilegeParams) WithTimeout(timeout time.Duration) *UpdatePrivilegeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update privilege params
func (o *UpdatePrivilegeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update privilege params
func (o *UpdatePrivilegeParams) WithContext(ctx context.Context) *UpdatePrivilegeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update privilege params
func (o *UpdatePrivilegeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update privilege params
func (o *UpdatePrivilegeParams) WithHTTPClient(client *http.Client) *UpdatePrivilegeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update privilege params
func (o *UpdatePrivilegeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update privilege params
func (o *UpdatePrivilegeParams) WithBody(body *models.APIPrivilegeApplicationRequest) *UpdatePrivilegeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update privilege params
func (o *UpdatePrivilegeParams) SetBody(body *models.APIPrivilegeApplicationRequest) {
	o.Body = body
}

// WithPrivilegeName adds the privilegeName to the update privilege params
func (o *UpdatePrivilegeParams) WithPrivilegeName(privilegeName string) *UpdatePrivilegeParams {
	o.SetPrivilegeName(privilegeName)
	return o
}

// SetPrivilegeName adds the privilegeName to the update privilege params
func (o *UpdatePrivilegeParams) SetPrivilegeName(privilegeName string) {
	o.PrivilegeName = privilegeName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePrivilegeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param privilegeName
	if err := r.SetPathParam("privilegeName", o.PrivilegeName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
