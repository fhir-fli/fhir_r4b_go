package fhir_r4b_go

import (
	"encoding/json"
)

// DataType represents the base class for all reusable types in FHIR.
type DataType struct {
	Element
}

// NewDataType creates a new DataType instance.
func NewDataType(
	id *string,
	extensions []*FhirExtension,
	disallowExtension bool,
) *DataType {
	return &DataType{
		Element: *NewElement(id, extensions, disallowExtension),
	}
}

// Equals compares DataType with another Equalable instance.
func (dt *DataType) Equals(other Equalable) bool {
	otherDataType, ok := other.(*DataType)
	if !ok {
		return false
	}
	return dt.Element.Equals(&otherDataType.Element)
}

// Clone creates a deep copy of the DataType.
func (dt *DataType) Clone() *DataType {
	if dt == nil {
		return nil
	}
	return &DataType{
		Element: *dt.Element.Clone(),
	}
}

// MarshalJSON converts the DataType to its JSON representation.
func (dt *DataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(dt)
}
