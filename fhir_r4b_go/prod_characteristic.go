// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json")

// ProdCharacteristic
// The marketing status describes the date when a medicinal product is actually put on the market or the date as of which it is no longer available.
type ProdCharacteristic struct {
	BackboneType
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Height *Quantity `json:"height,omitempty"`
	Width *Quantity `json:"width,omitempty"`
	Depth *Quantity `json:"depth,omitempty"`
	Weight *Quantity `json:"weight,omitempty"`
	NominalVolume *Quantity `json:"nominalvolume,omitempty"`
	ExternalDiameter *Quantity `json:"externaldiameter,omitempty"`
	Shape *FhirString `json:"shape,omitempty"`
	Color []*FhirString `json:"color,omitempty"`
	Imprint []*FhirString `json:"imprint,omitempty"`
	Image []*Attachment `json:"image,omitempty"`
	Scoring *CodeableConcept `json:"scoring,omitempty"`
}

// NewProdCharacteristic creates a new ProdCharacteristic instance
func NewProdCharacteristic() *ProdCharacteristic {
	return &ProdCharacteristic{}
}

// FromJSON populates ProdCharacteristic from JSON data
func (m *ProdCharacteristic) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts ProdCharacteristic to JSON data
func (m *ProdCharacteristic) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of ProdCharacteristic
func (m *ProdCharacteristic) Clone() *ProdCharacteristic {
	if m == nil { return nil }
	return &ProdCharacteristic{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Height: m.Height.Clone(),
		Width: m.Width.Clone(),
		Depth: m.Depth.Clone(),
		Weight: m.Weight.Clone(),
		NominalVolume: m.NominalVolume.Clone(),
		ExternalDiameter: m.ExternalDiameter.Clone(),
		Shape: m.Shape.Clone(),
		Color: cloneSlices(m.Color),
		Imprint: cloneSlices(m.Imprint),
		Image: cloneSlices(m.Image),
		Scoring: m.Scoring.Clone(),
	}
}

// Equals checks for equality with another ProdCharacteristic instance
func (m *ProdCharacteristic) Equals(other *ProdCharacteristic) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Height.Equals(other.Height) { return false }
	if !m.Width.Equals(other.Width) { return false }
	if !m.Depth.Equals(other.Depth) { return false }
	if !m.Weight.Equals(other.Weight) { return false }
	if !m.NominalVolume.Equals(other.NominalVolume) { return false }
	if !m.ExternalDiameter.Equals(other.ExternalDiameter) { return false }
	if !m.Shape.Equals(other.Shape) { return false }
	if !compareSlices(m.Color, other.Color) { return false }
	if !compareSlices(m.Imprint, other.Imprint) { return false }
	if !compareSlices(m.Image, other.Image) { return false }
	if !m.Scoring.Equals(other.Scoring) { return false }
	return true
}
