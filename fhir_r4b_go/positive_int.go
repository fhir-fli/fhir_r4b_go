package fhir_r4b_go

import (
	"encoding/json"
	"fmt"
)

// FhirPositiveInt represents the FHIR 'integer' type.
type FhirPositiveInt struct {
	Value   *int64   `json:"-"`          // The actual value
	Element *Element `json:",inline"`    // Metadata (FHIR element)
}

// NewFhirPositiveInt creates a new validated FhirPositiveInt.
func NewFhirPositiveInt(value int64, element *Element) *FhirPositiveInt {
	return &FhirPositiveInt{Value: &value, Element: element}
}

// UnmarshalJSON deserializes JSON into FhirPositiveInt.
func (fi *FhirPositiveInt) UnmarshalJSON(data []byte) error {
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

// MarshalJSON serializes FhirPositiveInt to JSON.
func (fi *FhirPositiveInt) MarshalJSON() ([]byte, error) {
	raw := map[string]interface{}{}
	if fi.Value != nil {
		raw["value"] = *fi.Value
	}
	if fi.Element != nil {
		raw["_value"] = fi.Element
	}
	return json.Marshal(raw)
}

// Clone creates a deep copy of FhirPositiveInt.
func (fi *FhirPositiveInt) Clone() *FhirPositiveInt {
	if fi == nil {
		return nil
	}
	return &FhirPositiveInt{
		Value:   int64PtrIfNotNil(fi.Value),
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

// String returns the string representation of FhirPositiveInt.
func (fi *FhirPositiveInt) String() string {
	return fmt.Sprintf("%d", fi.Value)
}
