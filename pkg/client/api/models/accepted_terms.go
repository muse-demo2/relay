// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AcceptedTerms Accepted terms and conditions
// swagger:model AcceptedTerms
type AcceptedTerms struct {

	// Timestamp when terms and conditions were accepted by an owner of this account
	// Required: true
	// Format: date-time
	AcceptedTermsAt *strfmt.DateTime `json:"accepted_terms_at"`

	// The version of the terms and conditions that was accepted
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this accepted terms
func (m *AcceptedTerms) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptedTermsAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AcceptedTerms) validateAcceptedTermsAt(formats strfmt.Registry) error {

	if err := validate.Required("accepted_terms_at", "body", m.AcceptedTermsAt); err != nil {
		return err
	}

	if err := validate.FormatOf("accepted_terms_at", "body", "date-time", m.AcceptedTermsAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AcceptedTerms) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AcceptedTerms) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AcceptedTerms) UnmarshalBinary(b []byte) error {
	var res AcceptedTerms
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}