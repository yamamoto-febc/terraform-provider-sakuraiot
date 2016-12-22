package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// Module module
// swagger:model Module
type Module struct {

	// id
	ID string `json:"id,omitempty"`

	// is online
	IsOnline bool `json:"is_online,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Project ID
	Project int64 `json:"project,omitempty"`
}

// Validate validates this module
func (m *Module) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
