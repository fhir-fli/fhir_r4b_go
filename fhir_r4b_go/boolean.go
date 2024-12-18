package fhir_r4b_go

import (
	"encoding/json"
)

// FhirBoolean represents the FHIR primitive type `boolean`.
type FhirBoolean struct {
	Value   *bool   `json:"-"`           // The boolean value
	Element *Element `json:",inline"`    // Metadata (FHIR element)
}

// NewFhirBoolean creates a new FhirBoolean instance.
func NewFhirBoolean(value bool, element *Element) *FhirBoolean {
	return &FhirBoolean{
		Value:   &value,
		Element: element,
	}
}

// UnmarshalJSON initializes a FhirBoolean from JSON input.
func (fb *FhirBoolean) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Extract value if present
	if rawValue, exists := raw["value"]; exists {
		var value bool
		if err := json.Unmarshal(rawValue, &value); err != nil {
			return err
		}
		fb.Value = &value
	}

	// Extract element metadata if present
	if rawElement, exists := raw["_value"]; exists {
		fb.Element = &Element{}
		if err := json.Unmarshal(rawElement, fb.Element); err != nil {
			return err
		}
	}

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
