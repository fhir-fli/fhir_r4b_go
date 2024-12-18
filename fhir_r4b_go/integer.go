package fhir_r4b_go

import (
	"encoding/json"
	"fmt"
)

// FhirInteger represents the FHIR 'integer' type.
type FhirInteger struct {
	Value   *int64   `json:"-"`          // The actual value
	Element *Element `json:",inline"`    // Metadata (FHIR element)
}

// NewFhirInteger creates a new validated FhirInteger.
func NewFhirInteger(value int64, element *Element) *FhirInteger {
	return &FhirInteger{Value: &value, Element: element}
}

// UnmarshalJSON deserializes JSON into FhirInteger.
func (fi *FhirInteger) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if rawValue, exists := raw["value"]; exists {
		var value int64
		if err := json.Unmarshal(rawValue, &value); err != nil {
			return err
		}
		fi.Value = &value
	}

	if rawElement, exists := raw["_value"]; exists {
		fi.Element = &Element{}
		if err := json.Unmarshal(rawElement, fi.Element); err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON serializes FhirInteger to JSON.
func (fi *FhirInteger) MarshalJSON() ([]byte, error) {
	raw := map[string]interface{}{}
	if fi.Value != nil {
		raw["value"] = *fi.Value
	}
	if fi.Element != nil {
		raw["_value"] = fi.Element
	}
	return json.Marshal(raw)
}

// Clone creates a deep copy of FhirInteger.
func (fi *FhirInteger) Clone() *FhirInteger {
	if fi == nil {
		return nil
	}
	return &FhirInteger{
		Value:   int64PtrIfNotNil(fi.Value),
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

// String returns the string representation of FhirInteger.
func (fi *FhirInteger) String() string {
	return fmt.Sprintf("%d", fi.Value)
}
