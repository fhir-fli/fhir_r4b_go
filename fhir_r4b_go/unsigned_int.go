package fhir_r4b_go

import (
	"encoding/json"
	"fmt"
)

// FhirUnsignedInt represents the FHIR 'integer' type.
type FhirUnsignedInt struct {
	Value   int64    `json:"value,omitempty"`
	Element *Element `json:"_value,omitempty"`
}

// NewFhirUnsignedInt creates a new validated FhirUnsignedInt.
func NewFhirUnsignedInt(value int64, element *Element) *FhirUnsignedInt {
	return &FhirUnsignedInt{Value: value, Element: element}
}

// Clone creates a deep copy of FhirUnsignedInt.
func (fi *FhirUnsignedInt) Clone() *FhirUnsignedInt {
	if fi == nil {
		return nil
	}
	return &FhirUnsignedInt{
		Value:   fi.Value,
		Element: fi.Element.Clone(),
	}
}

// Equals checks equality between two FhirUnsignedInt instances.
func (fi *FhirUnsignedInt) Equals(other *FhirUnsignedInt) bool {
	if fi == nil && other == nil {
		return true
	}
	if fi == nil || other == nil {
		return false
	}
	return fi.Value == other.Value && fi.Element.Equals(other.Element)
}

// MarshalJSON serializes FhirUnsignedInt to JSON.
func (fi *FhirUnsignedInt) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"value": fi.Value,
	}
	if fi.Element != nil {
		data["_value"] = fi.Element
	}
	return json.Marshal(data)
}

// UnmarshalJSON deserializes JSON into FhirUnsignedInt.
func (fi *FhirUnsignedInt) UnmarshalJSON(data []byte) error {
	temp := struct {
		Value   int64    `json:"value"`
		Element *Element `json:"_value"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	fi.Value = temp.Value
	fi.Element = temp.Element
	return nil
}

// String returns the string representation of FhirUnsignedInt.
func (fi *FhirUnsignedInt) String() string {
	return fmt.Sprintf("%d", fi.Value)
}
