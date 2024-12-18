package fhir_r4b_go

import (
	"encoding/json"
	"fmt"
)

// FhirPositiveInt represents the FHIR 'integer' type.
type FhirPositiveInt struct {
	Value   int64    `json:"value,omitempty"`
	Element *Element `json:"_value,omitempty"`
}

// NewFhirPositiveInt creates a new validated FhirPositiveInt.
func NewFhirPositiveInt(value int64, element *Element) *FhirPositiveInt {
	return &FhirPositiveInt{Value: value, Element: element}
}

// Clone creates a deep copy of FhirPositiveInt.
func (fi *FhirPositiveInt) Clone() *FhirPositiveInt {
	if fi == nil {
		return nil
	}
	return &FhirPositiveInt{
		Value:   fi.Value,
		Element: fi.Element.Clone(),
	}
}

// Equals checks equality between two FhirPositiveInt instances.
func (fi *FhirPositiveInt) Equals(other *FhirPositiveInt) bool {
	if fi == nil && other == nil {
		return true
	}
	if fi == nil || other == nil {
		return false
	}
	return fi.Value == other.Value && fi.Element.Equals(other.Element)
}

// MarshalJSON serializes FhirPositiveInt to JSON.
func (fi *FhirPositiveInt) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"value": fi.Value,
	}
	if fi.Element != nil {
		data["_value"] = fi.Element
	}
	return json.Marshal(data)
}

// UnmarshalJSON deserializes JSON into FhirPositiveInt.
func (fi *FhirPositiveInt) UnmarshalJSON(data []byte) error {
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

// String returns the string representation of FhirPositiveInt.
func (fi *FhirPositiveInt) String() string {
	return fmt.Sprintf("%d", fi.Value)
}
