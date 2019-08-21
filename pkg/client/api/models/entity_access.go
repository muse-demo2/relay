// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// EntityAccess Access control information for a given entity
// swagger:model EntityAccess
type EntityAccess struct {

	// The grants available relative to this entity
	PermissionGrants []*PermissionGrant `json:"permission_grants"`
}

// Validate validates this entity access
func (m *EntityAccess) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissionGrants(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntityAccess) validatePermissionGrants(formats strfmt.Registry) error {

	if swag.IsZero(m.PermissionGrants) { // not required
		return nil
	}

	for i := 0; i < len(m.PermissionGrants); i++ {
		if swag.IsZero(m.PermissionGrants[i]) { // not required
			continue
		}

		if m.PermissionGrants[i] != nil {
			if err := m.PermissionGrants[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permission_grants" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EntityAccess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntityAccess) UnmarshalBinary(b []byte) error {
	var res EntityAccess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
