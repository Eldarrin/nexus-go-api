// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HTTPClientConnectionAttributes Http client connection attributes
//
// swagger:model HttpClientConnectionAttributes
type HTTPClientConnectionAttributes struct {

	// Whether to enable redirects to the same location (may be required by some servers)
	// Example: false
	EnableCircularRedirects bool `json:"enableCircularRedirects,omitempty"`

	// Whether to allow cookies to be stored and used
	// Example: false
	EnableCookies bool `json:"enableCookies,omitempty"`

	// Total retries if the initial connection attempt suffers a timeout
	// Example: 0
	// Maximum: 10
	// Minimum: 0
	Retries *int32 `json:"retries,omitempty"`

	// Seconds to wait for activity before stopping and retrying the connection
	// Example: 60
	// Maximum: 3600
	// Minimum: 1
	Timeout int32 `json:"timeout,omitempty"`

	// Use certificates stored in the Nexus Repository Manager truststore to connect to external systems
	// Example: false
	UseTrustStore bool `json:"useTrustStore,omitempty"`

	// Custom fragment to append to User-Agent header in HTTP requests
	UserAgentSuffix string `json:"userAgentSuffix,omitempty"`
}

// Validate validates this Http client connection attributes
func (m *HTTPClientConnectionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRetries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeout(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTTPClientConnectionAttributes) validateRetries(formats strfmt.Registry) error {
	if swag.IsZero(m.Retries) { // not required
		return nil
	}

	if err := validate.MinimumInt("retries", "body", int64(*m.Retries), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("retries", "body", int64(*m.Retries), 10, false); err != nil {
		return err
	}

	return nil
}

func (m *HTTPClientConnectionAttributes) validateTimeout(formats strfmt.Registry) error {
	if swag.IsZero(m.Timeout) { // not required
		return nil
	}

	if err := validate.MinimumInt("timeout", "body", int64(m.Timeout), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("timeout", "body", int64(m.Timeout), 3600, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Http client connection attributes based on context it is used
func (m *HTTPClientConnectionAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HTTPClientConnectionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTTPClientConnectionAttributes) UnmarshalBinary(b []byte) error {
	var res HTTPClientConnectionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
