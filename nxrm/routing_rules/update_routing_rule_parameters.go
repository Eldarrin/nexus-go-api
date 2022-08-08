// Code generated by go-swagger; DO NOT EDIT.

package routing_rules

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

// NewUpdateRoutingRuleParams creates a new UpdateRoutingRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRoutingRuleParams() *UpdateRoutingRuleParams {
	return &UpdateRoutingRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRoutingRuleParamsWithTimeout creates a new UpdateRoutingRuleParams object
// with the ability to set a timeout on a request.
func NewUpdateRoutingRuleParamsWithTimeout(timeout time.Duration) *UpdateRoutingRuleParams {
	return &UpdateRoutingRuleParams{
		timeout: timeout,
	}
}

// NewUpdateRoutingRuleParamsWithContext creates a new UpdateRoutingRuleParams object
// with the ability to set a context for a request.
func NewUpdateRoutingRuleParamsWithContext(ctx context.Context) *UpdateRoutingRuleParams {
	return &UpdateRoutingRuleParams{
		Context: ctx,
	}
}

// NewUpdateRoutingRuleParamsWithHTTPClient creates a new UpdateRoutingRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRoutingRuleParamsWithHTTPClient(client *http.Client) *UpdateRoutingRuleParams {
	return &UpdateRoutingRuleParams{
		HTTPClient: client,
	}
}

/* UpdateRoutingRuleParams contains all the parameters to send to the API endpoint
   for the update routing rule operation.

   Typically these are written to a http.Request.
*/
type UpdateRoutingRuleParams struct {

	/* Body.

	   A routing rule configuration
	*/
	Body *models.RoutingRuleXO

	/* Name.

	   The name of the routing rule to update
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update routing rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRoutingRuleParams) WithDefaults() *UpdateRoutingRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update routing rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRoutingRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update routing rule params
func (o *UpdateRoutingRuleParams) WithTimeout(timeout time.Duration) *UpdateRoutingRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update routing rule params
func (o *UpdateRoutingRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update routing rule params
func (o *UpdateRoutingRuleParams) WithContext(ctx context.Context) *UpdateRoutingRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update routing rule params
func (o *UpdateRoutingRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update routing rule params
func (o *UpdateRoutingRuleParams) WithHTTPClient(client *http.Client) *UpdateRoutingRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update routing rule params
func (o *UpdateRoutingRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update routing rule params
func (o *UpdateRoutingRuleParams) WithBody(body *models.RoutingRuleXO) *UpdateRoutingRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update routing rule params
func (o *UpdateRoutingRuleParams) SetBody(body *models.RoutingRuleXO) {
	o.Body = body
}

// WithName adds the name to the update routing rule params
func (o *UpdateRoutingRuleParams) WithName(name string) *UpdateRoutingRuleParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update routing rule params
func (o *UpdateRoutingRuleParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRoutingRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
