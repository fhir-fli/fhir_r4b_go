package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"fmt"
)

// FhirUnsignedInt represents the FHIR 'integer' type.
type FhirUnsignedInt struct {
	Value   *int64   `json:"-"`       // The actual value
	Element *Element `json:",inline"` // Metadata (FHIR element)
}

// NewFhirUnsignedInt creates a new validated FhirUnsignedInt.
func NewFhirUnsignedInt(value int64, element *Element) *FhirUnsignedInt {
	return &FhirUnsignedInt{Value: &value, Element: element}
}

func NewFhirUnsignedIntFromMap(data map[string]interface{}) (*FhirUnsignedInt, error) {
	value, ok := data["value"].(*int64)
	if !ok {
		return nil, errors.New("invalid or missing value for FhirUnsignedInt")
	}

	integer := &FhirUnsignedInt{Value: value}

	if elementData, ok := data["_value"].(map[string]interface{}); ok {
		integer.Element = &Element{}
		if err := mapToStruct(elementData, integer.Element); err != nil {
			return nil, fmt.Errorf("failed to parse _value for FhirUnsignedInt: %v", err)
		}
	}

	return integer, nil
}

// UnmarshalJSON deserializes JSON into FhirUnsignedInt.
func (fi *FhirUnsignedInt) UnmarshalJSON(data []byte) error {
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

// MarshalJSON serializes FhirUnsignedInt to JSON.
func (fi *FhirUnsignedInt) MarshalJSON() ([]byte, error) {
	raw := map[string]interface{}{}
	if fi.Value != nil {
		raw["value"] = *fi.Value
	}
	if fi.Element != nil {
		raw["_value"] = fi.Element
	}
	return json.Marshal(raw)
}

// Clone creates a deep copy of FhirUnsignedInt.
func (fi *FhirUnsignedInt) Clone() *FhirUnsignedInt {
	if fi == nil {
		return nil
	}
	return &FhirUnsignedInt{
		Value:   int64PtrIfNotNil(fi.Value),
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

// String returns the string representation of FhirUnsignedInt.
func (fi *FhirUnsignedInt) String() string {
	return fmt.Sprintf("%d", fi.Value)
}
