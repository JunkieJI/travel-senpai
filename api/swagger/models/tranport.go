// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Tranport Tranport
// swagger:model tranport

type Tranport struct {

	// arrival date
	ArrivalDate string `json:"arrival_date,omitempty"`

	// arrival time
	ArrivalTime string `json:"arrival_time,omitempty"`

	// cost
	Cost int64 `json:"cost,omitempty"`

	// departure date
	DepartureDate string `json:"departure_date,omitempty"`

	// departure time
	DepartureTime string `json:"departure_time,omitempty"`

	// destination id
	DestinationID int64 `json:"destination_id,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// layover
	Layover string `json:"layover,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

/* polymorph tranport arrival_date false */

/* polymorph tranport arrival_time false */

/* polymorph tranport cost false */

/* polymorph tranport departure_date false */

/* polymorph tranport departure_time false */

/* polymorph tranport destination_id false */

/* polymorph tranport id false */

/* polymorph tranport layover false */

/* polymorph tranport type false */

// Validate validates this tranport
func (m *Tranport) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Tranport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tranport) UnmarshalBinary(b []byte) error {
	var res Tranport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
