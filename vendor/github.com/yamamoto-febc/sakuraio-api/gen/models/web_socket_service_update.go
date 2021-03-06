package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// WebSocketServiceUpdate WebSocketService
// swagger:model WebSocketServiceUpdate
type WebSocketServiceUpdate struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// Project ID
	// Required: true
	Project *int64 `json:"project"`

	// "websocket"
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this web socket service update
func (m *WebSocketServiceUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebSocketServiceUpdate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *WebSocketServiceUpdate) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("project", "body", m.Project); err != nil {
		return err
	}

	return nil
}

var webSocketServiceUpdateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["websocket"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		webSocketServiceUpdateTypeTypePropEnum = append(webSocketServiceUpdateTypeTypePropEnum, v)
	}
}

const (
	webSocketServiceUpdateTypeWebsocket string = "websocket"
)

// prop value enum
func (m *WebSocketServiceUpdate) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, webSocketServiceUpdateTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WebSocketServiceUpdate) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}
