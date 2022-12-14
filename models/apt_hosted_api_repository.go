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

// AptHostedAPIRepository apt hosted Api repository
//
// swagger:model AptHostedApiRepository
type AptHostedAPIRepository struct {

	// apt
	// Required: true
	Apt *AptHostedRepositoriesAttributes `json:"apt"`

	// apt signing
	// Required: true
	AptSigning *AptSigningRepositoriesAttributes `json:"aptSigning"`

	// cleanup
	Cleanup *CleanupPolicyAttributes `json:"cleanup,omitempty"`

	// component
	Component *ComponentAttributes `json:"component,omitempty"`

	// A unique identifier for this repository
	// Example: internal
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name,omitempty"`

	// Whether this repository accepts incoming requests
	// Example: true
	// Required: true
	Online *bool `json:"online"`

	// storage
	// Required: true
	Storage *HostedStorageAttributes `json:"storage"`
}

// Validate validates this apt hosted Api repository
func (m *AptHostedAPIRepository) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAptSigning(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCleanup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AptHostedAPIRepository) validateApt(formats strfmt.Registry) error {

	if err := validate.Required("apt", "body", m.Apt); err != nil {
		return err
	}

	if m.Apt != nil {
		if err := m.Apt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apt")
			}
			return err
		}
	}

	return nil
}

func (m *AptHostedAPIRepository) validateAptSigning(formats strfmt.Registry) error {

	if err := validate.Required("aptSigning", "body", m.AptSigning); err != nil {
		return err
	}

	if m.AptSigning != nil {
		if err := m.AptSigning.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aptSigning")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aptSigning")
			}
			return err
		}
	}

	return nil
}

func (m *AptHostedAPIRepository) validateCleanup(formats strfmt.Registry) error {
	if swag.IsZero(m.Cleanup) { // not required
		return nil
	}

	if m.Cleanup != nil {
		if err := m.Cleanup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cleanup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cleanup")
			}
			return err
		}
	}

	return nil
}

func (m *AptHostedAPIRepository) validateComponent(formats strfmt.Registry) error {
	if swag.IsZero(m.Component) { // not required
		return nil
	}

	if m.Component != nil {
		if err := m.Component.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

func (m *AptHostedAPIRepository) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", m.Name, `^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AptHostedAPIRepository) validateOnline(formats strfmt.Registry) error {

	if err := validate.Required("online", "body", m.Online); err != nil {
		return err
	}

	return nil
}

func (m *AptHostedAPIRepository) validateStorage(formats strfmt.Registry) error {

	if err := validate.Required("storage", "body", m.Storage); err != nil {
		return err
	}

	if m.Storage != nil {
		if err := m.Storage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this apt hosted Api repository based on the context it is used
func (m *AptHostedAPIRepository) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAptSigning(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCleanup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComponent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AptHostedAPIRepository) contextValidateApt(ctx context.Context, formats strfmt.Registry) error {

	if m.Apt != nil {
		if err := m.Apt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apt")
			}
			return err
		}
	}

	return nil
}

func (m *AptHostedAPIRepository) contextValidateAptSigning(ctx context.Context, formats strfmt.Registry) error {

	if m.AptSigning != nil {
		if err := m.AptSigning.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aptSigning")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aptSigning")
			}
			return err
		}
	}

	return nil
}

func (m *AptHostedAPIRepository) contextValidateCleanup(ctx context.Context, formats strfmt.Registry) error {

	if m.Cleanup != nil {
		if err := m.Cleanup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cleanup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cleanup")
			}
			return err
		}
	}

	return nil
}

func (m *AptHostedAPIRepository) contextValidateComponent(ctx context.Context, formats strfmt.Registry) error {

	if m.Component != nil {
		if err := m.Component.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

func (m *AptHostedAPIRepository) contextValidateStorage(ctx context.Context, formats strfmt.Registry) error {

	if m.Storage != nil {
		if err := m.Storage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AptHostedAPIRepository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AptHostedAPIRepository) UnmarshalBinary(b []byte) error {
	var res AptHostedAPIRepository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
