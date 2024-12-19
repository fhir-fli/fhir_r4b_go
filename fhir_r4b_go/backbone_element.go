package fhir_r4b_go

import (
	"encoding/json"
)

// BackboneElement represents elements inside a resource but not those in a data type.
type BackboneElement struct {
	DataType
	ModifierExtension []*FhirExtension `json:"modifierExtension,omitempty"`
}

// NewBackboneElement creates a new BackboneElement instance.
func NewBackboneElement(
	id *string,
	extensions []*FhirExtension,
	modifierExtension []*FhirExtension,
	disallowExtension bool,
) *BackboneElement {
	return &BackboneElement{
		DataType:          *NewDataType(id, extensions, disallowExtension),
		ModifierExtension: modifierExtension,
	}
}

// Equals compares BackboneElement with another Equalable instance.
func (be *BackboneElement) Equals(other Equalable) bool {
	otherBE, ok := other.(*BackboneElement)
	if !ok {
		return false
	}

	if !be.DataType.Equals(&otherBE.DataType) {
		return false
	}

	if !compareSlices(be.ModifierExtension, otherBE.ModifierExtension) {
		return false
	}

	return true
}

// Clone creates a deep copy of the BackboneElement.
func (be *BackboneElement) Clone() *BackboneElement {
	if be == nil {
		return nil
	}
	return &BackboneElement{
		DataType:          *be.DataType.Clone(),
		ModifierExtension: cloneSlices(be.ModifierExtension),
	}
}

// MarshalJSON converts the BackboneElement to JSON.
func (be *BackboneElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(be)
}
