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

// S3BlobStoreAPIBucket s3 blob store Api bucket
//
// swagger:model S3BlobStoreApiBucket
type S3BlobStoreAPIBucket struct {

	// How many days until deleted blobs are finally removed from the S3 bucket (-1 to disable)
	// Example: 3
	// Required: true
	Expiration *int32 `json:"expiration"`

	// The name of the S3 bucket
	// Required: true
	Name *string `json:"name"`

	// The S3 blob store (i.e S3 object) key prefix
	Prefix string `json:"prefix,omitempty"`

	// The AWS region to create a new S3 bucket in or an existing S3 bucket's region
	// Example: DEFAULT
	// Required: true
	Region *string `json:"region"`
}

// Validate validates this s3 blob store Api bucket
func (m *S3BlobStoreAPIBucket) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BlobStoreAPIBucket) validateExpiration(formats strfmt.Registry) error {

	if err := validate.Required("expiration", "body", m.Expiration); err != nil {
		return err
	}

	return nil
}

func (m *S3BlobStoreAPIBucket) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *S3BlobStoreAPIBucket) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this s3 blob store Api bucket based on context it is used
func (m *S3BlobStoreAPIBucket) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *S3BlobStoreAPIBucket) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BlobStoreAPIBucket) UnmarshalBinary(b []byte) error {
	var res S3BlobStoreAPIBucket
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
