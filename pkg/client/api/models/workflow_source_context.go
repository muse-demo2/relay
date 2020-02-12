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

// WorkflowSourceContext Metadata about the source control origin of a yaml workflow definition
// swagger:model WorkflowSourceContext
type WorkflowSourceContext struct {

	// Git branch on which file lived, if relevant
	Branch string `json:"branch,omitempty"`

	// Git commit hash
	// Required: true
	Hash *string `json:"hash"`

	// integration
	// Required: true
	Integration *IntegrationRevision `json:"integration"`

	// Relative path from the repository root to the workflow file
	// Required: true
	Path *string `json:"path"`

	// A source repository slug
	// Required: true
	Repository *string `json:"repository"`
}

// Validate validates this workflow source context
func (m *WorkflowSourceContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntegration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowSourceContext) validateHash(formats strfmt.Registry) error {

	if err := validate.Required("hash", "body", m.Hash); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowSourceContext) validateIntegration(formats strfmt.Registry) error {

	if err := validate.Required("integration", "body", m.Integration); err != nil {
		return err
	}

	if m.Integration != nil {
		if err := m.Integration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integration")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowSourceContext) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowSourceContext) validateRepository(formats strfmt.Registry) error {

	if err := validate.Required("repository", "body", m.Repository); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowSourceContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowSourceContext) UnmarshalBinary(b []byte) error {
	var res WorkflowSourceContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}