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

	"github.com/Eldarrin/nexus-go-api/models"
)

// NewCreateRoutingRuleParams creates a new CreateRoutingRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRoutingRuleParams() *CreateRoutingRuleParams {
	return &CreateRoutingRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRoutingRuleParamsWithTimeout creates a new CreateRoutingRuleParams object
// with the ability to set a timeout on a request.
func NewCreateRoutingRuleParamsWithTimeout(timeout time.Duration) *CreateRoutingRuleParams {
	return &CreateRoutingRuleParams{
		timeout: timeout,
	}
}

// NewCreateRoutingRuleParamsWithContext creates a new CreateRoutingRuleParams object
// with the ability to set a context for a request.
func NewCreateRoutingRuleParamsWithContext(ctx context.Context) *CreateRoutingRuleParams {
	return &CreateRoutingRuleParams{
		Context: ctx,
	}
}

// NewCreateRoutingRuleParamsWithHTTPClient creates a new CreateRoutingRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRoutingRuleParamsWithHTTPClient(client *http.Client) *CreateRoutingRuleParams {
	return &CreateRoutingRuleParams{
		HTTPClient: client,
	}
}

/* CreateRoutingRuleParams contains all the parameters to send to the API endpoint
   for the create routing rule operation.

   Typically these are written to a http.Request.
*/
type CreateRoutingRuleParams struct {

	/* Body.

	   A routing rule configuration
	*/
	Body *models.RoutingRuleXO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create routing rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRoutingRuleParams) WithDefaults() *CreateRoutingRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create routing rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRoutingRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create routing rule params
func (o *CreateRoutingRuleParams) WithTimeout(timeout time.Duration) *CreateRoutingRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create routing rule params
func (o *CreateRoutingRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create routing rule params
func (o *CreateRoutingRuleParams) WithContext(ctx context.Context) *CreateRoutingRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create routing rule params
func (o *CreateRoutingRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create routing rule params
func (o *CreateRoutingRuleParams) WithHTTPClient(client *http.Client) *CreateRoutingRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create routing rule params
func (o *CreateRoutingRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create routing rule params
func (o *CreateRoutingRuleParams) WithBody(body *models.RoutingRuleXO) *CreateRoutingRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create routing rule params
func (o *CreateRoutingRuleParams) SetBody(body *models.RoutingRuleXO) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRoutingRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
