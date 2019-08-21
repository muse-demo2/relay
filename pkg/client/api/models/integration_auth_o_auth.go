// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IntegrationAuthOAuth OAuth authentication
// swagger:model IntegrationAuthOAuth
type IntegrationAuthOAuth struct {

	// The temporary code to exchange
	// Required: true
	Code *string `json:"code"`

	// The OAuth state as specified in the initial request
	// Required: true
	State *string `json:"state"`

	// type
	// Required: true
	// Enum: [oauth]
	Type *string `json:"type"`
}

// Validate validates this integration auth o auth
func (m *IntegrationAuthOAuth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationAuthOAuth) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *IntegrationAuthOAuth) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

var integrationAuthOAuthTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["oauth"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		integrationAuthOAuthTypeTypePropEnum = append(integrationAuthOAuthTypeTypePropEnum, v)
	}
}

const (

	// IntegrationAuthOAuthTypeOauth captures enum value "oauth"
	IntegrationAuthOAuthTypeOauth string = "oauth"
)

// prop value enum
func (m *IntegrationAuthOAuth) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, integrationAuthOAuthTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IntegrationAuthOAuth) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationAuthOAuth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationAuthOAuth) UnmarshalBinary(b []byte) error {
	var res IntegrationAuthOAuth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
