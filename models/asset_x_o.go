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

// AssetXO asset x o
//
// swagger:model AssetXO
type AssetXO struct {

	// checksum
	Checksum map[string]interface{} `json:"checksum,omitempty"`

	// content type
	ContentType string `json:"contentType,omitempty"`

	// download Url
	DownloadURL string `json:"downloadUrl,omitempty"`

	// file size
	FileSize int64 `json:"fileSize,omitempty"`

	// format
	Format string `json:"format,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// last downloaded
	// Format: date-time
	LastDownloaded strfmt.DateTime `json:"lastDownloaded,omitempty"`

	// last modified
	// Format: date-time
	LastModified strfmt.DateTime `json:"lastModified,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// repository
	Repository string `json:"repository,omitempty"`

	// uploader
	Uploader string `json:"uploader,omitempty"`

	// uploader Ip
	UploaderIP string `json:"uploaderIp,omitempty"`
}

// Validate validates this asset x o
func (m *AssetXO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastDownloaded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModified(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetXO) validateLastDownloaded(formats strfmt.Registry) error {
	if swag.IsZero(m.LastDownloaded) { // not required
		return nil
	}

	if err := validate.FormatOf("lastDownloaded", "body", "date-time", m.LastDownloaded.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AssetXO) validateLastModified(formats strfmt.Registry) error {
	if swag.IsZero(m.LastModified) { // not required
		return nil
	}

	if err := validate.FormatOf("lastModified", "body", "date-time", m.LastModified.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this asset x o based on context it is used
func (m *AssetXO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AssetXO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetXO) UnmarshalBinary(b []byte) error {
	var res AssetXO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
