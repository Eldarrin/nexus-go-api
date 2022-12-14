// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GroupDeployAttributes group deploy attributes
//
// swagger:model GroupDeployAttributes
type GroupDeployAttributes struct {

	// Member repositories' names
	MemberNames []string `json:"memberNames"`

	// Pro-only: This field is for the Group Deployment feature available in NXRM Pro.
	WritableMember string `json:"writableMember,omitempty"`
}

// Validate validates this group deploy attributes
func (m *GroupDeployAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this group deploy attributes based on context it is used
func (m *GroupDeployAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GroupDeployAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupDeployAttributes) UnmarshalBinary(b []byte) error {
	var res GroupDeployAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
