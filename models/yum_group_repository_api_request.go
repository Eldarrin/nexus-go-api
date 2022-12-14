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

// YumGroupRepositoryAPIRequest yum group repository Api request
//
// swagger:model YumGroupRepositoryApiRequest
type YumGroupRepositoryAPIRequest struct {

	// group
	// Required: true
	Group *GroupAttributes `json:"group"`

	// A unique identifier for this repository
	// Example: internal
	// Required: true
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name *string `json:"name"`

	// Whether this repository accepts incoming requests
	// Example: true
	// Required: true
	Online *bool `json:"online"`

	// storage
	// Required: true
	Storage *StorageAttributes `json:"storage"`

	// yum signing
	YumSigning *YumSigningRepositoriesAttributes `json:"yumSigning,omitempty"`
}

// Validate validates this yum group repository Api request
func (m *YumGroupRepositoryAPIRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroup(formats); err != nil {
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

	if err := m.validateYumSigning(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *YumGroupRepositoryAPIRequest) validateGroup(formats strfmt.Registry) error {

	if err := validate.Required("group", "body", m.Group); err != nil {
		return err
	}

	if m.Group != nil {
		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *YumGroupRepositoryAPIRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", *m.Name, `^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$`); err != nil {
		return err
	}

	return nil
}

func (m *YumGroupRepositoryAPIRequest) validateOnline(formats strfmt.Registry) error {

	if err := validate.Required("online", "body", m.Online); err != nil {
		return err
	}

	return nil
}

func (m *YumGroupRepositoryAPIRequest) validateStorage(formats strfmt.Registry) error {

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

func (m *YumGroupRepositoryAPIRequest) validateYumSigning(formats strfmt.Registry) error {
	if swag.IsZero(m.YumSigning) { // not required
		return nil
	}

	if m.YumSigning != nil {
		if err := m.YumSigning.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("yumSigning")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("yumSigning")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this yum group repository Api request based on the context it is used
func (m *YumGroupRepositoryAPIRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateYumSigning(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *YumGroupRepositoryAPIRequest) contextValidateGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.Group != nil {
		if err := m.Group.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *YumGroupRepositoryAPIRequest) contextValidateStorage(ctx context.Context, formats strfmt.Registry) error {

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

func (m *YumGroupRepositoryAPIRequest) contextValidateYumSigning(ctx context.Context, formats strfmt.Registry) error {

	if m.YumSigning != nil {
		if err := m.YumSigning.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("yumSigning")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("yumSigning")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *YumGroupRepositoryAPIRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *YumGroupRepositoryAPIRequest) UnmarshalBinary(b []byte) error {
	var res YumGroupRepositoryAPIRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
