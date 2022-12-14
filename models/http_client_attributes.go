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

// HTTPClientAttributes Http client attributes
//
// swagger:model HttpClientAttributes
type HTTPClientAttributes struct {

	// authentication
	Authentication *HTTPClientConnectionAuthenticationAttributes `json:"authentication,omitempty"`

	// Whether to auto-block outbound connections if remote peer is detected as unreachable/unresponsive
	// Example: true
	// Required: true
	AutoBlock *bool `json:"autoBlock"`

	// Whether to block outbound connections on the repository
	// Example: false
	// Required: true
	Blocked *bool `json:"blocked"`

	// connection
	Connection *HTTPClientConnectionAttributes `json:"connection,omitempty"`
}

// Validate validates this Http client attributes
func (m *HTTPClientAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoBlock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlocked(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTTPClientAttributes) validateAuthentication(formats strfmt.Registry) error {
	if swag.IsZero(m.Authentication) { // not required
		return nil
	}

	if m.Authentication != nil {
		if err := m.Authentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authentication")
			}
			return err
		}
	}

	return nil
}

func (m *HTTPClientAttributes) validateAutoBlock(formats strfmt.Registry) error {

	if err := validate.Required("autoBlock", "body", m.AutoBlock); err != nil {
		return err
	}

	return nil
}

func (m *HTTPClientAttributes) validateBlocked(formats strfmt.Registry) error {

	if err := validate.Required("blocked", "body", m.Blocked); err != nil {
		return err
	}

	return nil
}

func (m *HTTPClientAttributes) validateConnection(formats strfmt.Registry) error {
	if swag.IsZero(m.Connection) { // not required
		return nil
	}

	if m.Connection != nil {
		if err := m.Connection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Http client attributes based on the context it is used
func (m *HTTPClientAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthentication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTTPClientAttributes) contextValidateAuthentication(ctx context.Context, formats strfmt.Registry) error {

	if m.Authentication != nil {
		if err := m.Authentication.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authentication")
			}
			return err
		}
	}

	return nil
}

func (m *HTTPClientAttributes) contextValidateConnection(ctx context.Context, formats strfmt.Registry) error {

	if m.Connection != nil {
		if err := m.Connection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HTTPClientAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTTPClientAttributes) UnmarshalBinary(b []byte) error {
	var res HTTPClientAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
