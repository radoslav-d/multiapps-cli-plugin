// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse200 inline response 200
// swagger:model inline_response_200

type InlineResponse200 struct {

	// action ids
	ActionIds []string `json:"actionIds"`
}

/* polymorph inline_response_200 actionIds false */

// Validate validates this inline response 200
func (m *InlineResponse200) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse200) validateActionIds(formats strfmt.Registry) error {

	if swag.IsZero(m.ActionIds) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse200) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse200) UnmarshalBinary(b []byte) error {
	var res InlineResponse200
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
