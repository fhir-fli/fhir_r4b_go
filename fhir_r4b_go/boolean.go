package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"fmt"
)

// FhirBoolean represents the FHIR primitive type `boolean`.
type FhirBoolean struct {
	Value   *bool    `json:"-"`       // The boolean value
	Element *Element `json:",inline"` // Metadata (FHIR element)
}

// NewFhirBoolean creates a new FhirBoolean instance.
func NewFhirBoolean(value bool, element *Element) *FhirBoolean {
	return &FhirBoolean{
		Value:   &value,
		Element: element,
	}
}

// NewFhirBooleanFromMap creates a FhirBoolean instance from a map.
func NewFhirBooleanFromMap(data map[string]interface{}) (*FhirBoolean, error) {
	value, ok := data["value"].(bool)
	if !ok {
		return nil, errors.New("invalid or missing value for FhirBoolean")
	}

	boolean := &FhirBoolean{Value: &value}

	if elementData, ok := data["_value"].(map[string]interface{}); ok {
		boolean.Element = &Element{}
		if err := mapToStruct(elementData, boolean.Element); err != nil {
			return nil, fmt.Errorf("failed to parse _value for FhirBoolean: %v", err)
		}
	}

	return boolean, nil
}

// UnmarshalJSON initializes a FhirBoolean from JSON input.
func (fb *FhirBoolean) UnmarshalJSON(data []byte) error {
	// Handle case where the input is a simple boolean
	var boolValue bool
	if err := json.Unmarshal(data, &boolValue); err == nil {
		fb.Value = &boolValue
		fb.Element = nil
		return nil
	}

	// Handle case where the input is a JSON object
	var raw struct {
		Value   *bool    `json:"value"`
		Element *Element `json:"_value"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	fb.Value = raw.Value
	fb.Element = raw.Element
	return nil
}

// MarshalJSON converts the FhirBoolean to JSON.
func (fb *FhirBoolean) MarshalJSON() ([]byte, error) {
	raw := make(map[string]interface{})

	if fb.Value != nil {
		raw["value"] = fb.Value
	}

	if fb.Element != nil {
		elementJSON, err := json.Marshal(fb.Element)
		if err != nil {
			return nil, err
		}
		var elementMap map[string]interface{}
		if err := json.Unmarshal(elementJSON, &elementMap); err != nil {
			return nil, err
		}
		raw["_value"] = elementMap
	}

	return json.Marshal(raw)
}

// Clone creates a deep copy of the FhirBoolean.
func (fb *FhirBoolean) Clone() *FhirBoolean {
	if fb == nil {
		return nil
	}
	return &FhirBoolean{
		Value:   fb.Value,
		Element: fb.Element.Clone(),
	}
}

// Equals checks for equality between two FhirBoolean instances.
func (fb *FhirBoolean) Equals(other *FhirBoolean) bool {
	if fb == nil && other == nil {
		return true
	}
	if fb == nil || other == nil {
		return false
	}
	return fb.Value == other.Value && fb.Element.Equals(other.Element)
}

// String returns the string representation of the FhirBoolean.
func (fb *FhirBoolean) String() string {
	if fb.Value == nil {
		return "null"
	}
	if *fb.Value {
		return "true"
	}
	return "false"
}
