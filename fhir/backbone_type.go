package fhir_r4b_go

import (
	"encoding/json"
)

// BackboneType represents a FHIR type that can carry modifier extensions.
type BackboneType struct {
	DataType
	ModifierExtension []*FhirExtension `json:"modifierExtension,omitempty"`
}

// NewBackboneType creates a new BackboneType instance.
func NewBackboneType(
	id *string,
	extensions []*FhirExtension,
	modifierExtension []*FhirExtension,
	disallowExtension bool,
) *BackboneType {
	return &BackboneType{
		DataType:          *NewDataType(id, extensions, disallowExtension),
		ModifierExtension: modifierExtension,
	}
}

// Equals compares BackboneType with another Equalable instance.
func (bt *BackboneType) Equals(other Equalable) bool {
	otherBT, ok := other.(*BackboneType)
	if !ok {
		return false
	}

	if !bt.DataType.Equals(&otherBT.DataType) {
		return false
	}

	if !compareSlices(bt.ModifierExtension, otherBT.ModifierExtension) {
		return false
	}

	return true
}

// Clone creates a deep copy of the BackboneType.
func (bt *BackboneType) Clone() *BackboneType {
	if bt == nil {
		return nil
	}
	return &BackboneType{
		DataType:          *bt.DataType.Clone(),
		ModifierExtension: cloneSlices(bt.ModifierExtension),
	}
}

// ToJSON converts the BackboneType to its JSON representation.
func (bt *BackboneType) ToJSON() ([]byte, error) {
	return json.Marshal(bt)
}
