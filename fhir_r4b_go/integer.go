package fhir_r4b_go

import (
	"encoding/json"
	"fmt"
)

// FhirInteger represents the FHIR 'integer' type.
type FhirInteger struct {
	Value   int64    `json:"value,omitempty"`
	Element *Element `json:"_value,omitempty"`
}

// NewFhirInteger creates a new validated FhirInteger.
func NewFhirInteger(value int64, element *Element) *FhirInteger {
	return &FhirInteger{Value: value, Element: element}
}

// Clone creates a deep copy of FhirInteger.
func (fi *FhirInteger) Clone() *FhirInteger {
	if fi == nil {
		return nil
	}
	return &FhirInteger{
		Value:   fi.Value,
		Element: fi.Element.Clone(),
	}
}

// Equals checks equality between two FhirInteger instances.
func (fi *FhirInteger) Equals(other *FhirInteger) bool {
	if fi == nil && other == nil {
		return true
	}
	if fi == nil || other == nil {
		return false
	}
	return fi.Value == other.Value && fi.Element.Equals(other.Element)
}

// MarshalJSON serializes FhirInteger to JSON.
func (fi *FhirInteger) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"value": fi.Value,
	}
	if fi.Element != nil {
		data["_value"] = fi.Element
	}
	return json.Marshal(data)
}

// UnmarshalJSON deserializes JSON into FhirInteger.
func (fi *FhirInteger) UnmarshalJSON(data []byte) error {
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

// String returns the string representation of FhirInteger.
func (fi *FhirInteger) String() string {
	return fmt.Sprintf("%d", fi.Value)
}
