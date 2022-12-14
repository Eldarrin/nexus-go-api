// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SupportZipXO support zip x o
//
// swagger:model SupportZipXO
type SupportZipXO struct {

	// file
	File string `json:"file,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// size
	Size string `json:"size,omitempty"`

	// truncated
	Truncated bool `json:"truncated,omitempty"`
}

// Validate validates this support zip x o
func (m *SupportZipXO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this support zip x o based on context it is used
func (m *SupportZipXO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SupportZipXO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupportZipXO) UnmarshalBinary(b []byte) error {
	var res SupportZipXO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
