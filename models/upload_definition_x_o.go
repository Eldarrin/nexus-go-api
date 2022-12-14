// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UploadDefinitionXO upload definition x o
//
// swagger:model UploadDefinitionXO
type UploadDefinitionXO struct {

	// asset fields
	AssetFields []*UploadFieldDefinitionXO `json:"assetFields"`

	// component fields
	ComponentFields []*UploadFieldDefinitionXO `json:"componentFields"`

	// format
	Format string `json:"format,omitempty"`

	// multiple upload
	MultipleUpload bool `json:"multipleUpload,omitempty"`
}

// Validate validates this upload definition x o
func (m *UploadDefinitionXO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssetFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponentFields(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UploadDefinitionXO) validateAssetFields(formats strfmt.Registry) error {
	if swag.IsZero(m.AssetFields) { // not required
		return nil
	}

	for i := 0; i < len(m.AssetFields); i++ {
		if swag.IsZero(m.AssetFields[i]) { // not required
			continue
		}

		if m.AssetFields[i] != nil {
			if err := m.AssetFields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assetFields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assetFields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UploadDefinitionXO) validateComponentFields(formats strfmt.Registry) error {
	if swag.IsZero(m.ComponentFields) { // not required
		return nil
	}

	for i := 0; i < len(m.ComponentFields); i++ {
		if swag.IsZero(m.ComponentFields[i]) { // not required
			continue
		}

		if m.ComponentFields[i] != nil {
			if err := m.ComponentFields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("componentFields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("componentFields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this upload definition x o based on the context it is used
func (m *UploadDefinitionXO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssetFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComponentFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UploadDefinitionXO) contextValidateAssetFields(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AssetFields); i++ {

		if m.AssetFields[i] != nil {
			if err := m.AssetFields[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assetFields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assetFields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UploadDefinitionXO) contextValidateComponentFields(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ComponentFields); i++ {

		if m.ComponentFields[i] != nil {
			if err := m.ComponentFields[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("componentFields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("componentFields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UploadDefinitionXO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UploadDefinitionXO) UnmarshalBinary(b []byte) error {
	var res UploadDefinitionXO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
