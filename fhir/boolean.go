package fhir_r4b_go

import (
	"encoding/json"
)

// FhirBoolean represents the FHIR primitive type `boolean`.
type FhirBoolean struct {
	Value   *bool    `json:"value,omitempty"`  // The boolean value
	Element *Element `json:"_value,omitempty"` // Additional metadata (FHIR element)
}

// NewFhirBoolean creates a new FhirBoolean instance.
func NewFhirBoolean(value bool, element *Element) *FhirBoolean {
	return &FhirBoolean{
		Value:   &value,
		Element: element,
	}
}

// FromJSON initializes a FhirBoolean from JSON input.
func (fb *FhirBoolean) FromJSON(data []byte) error {
	return json.Unmarshal(data, fb)
}

// ToJSON converts the FhirBoolean to JSON.
func (fb *FhirBoolean) ToJSON() ([]byte, error) {
	return json.Marshal(fb)
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
