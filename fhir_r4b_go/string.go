package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"strings"
)

// FhirString represents a validated string used in FHIR resources.
type FhirString struct {
	Value   string   `json:"value,omitempty"`
	Element *Element `json:"_value,omitempty"`
}

// NewFhirString creates a new FhirString with validation.
func NewFhirString(input string, element *Element) (*FhirString, error) {
	if input == "" && element == nil {
		return nil, errors.New("a value or element is required")
	}
	return &FhirString{Value: input, Element: element}, nil
}

// MarshalJSON serializes FhirString into JSON.
func (f *FhirString) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{}
	if f.Value != "" {
		data["value"] = f.Value
	}
	if f.Element != nil {
		data["_value"] = f.Element
	}
	return json.Marshal(data)
}

// UnmarshalJSON deserializes JSON into FhirString.
func (f *FhirString) UnmarshalJSON(data []byte) error {
	temp := struct {
		Value   string   `json:"value"`
		Element *Element `json:"_value"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	f.Value = temp.Value
	f.Element = temp.Element
	return nil
}

// String returns the string value.
func (f *FhirString) String() string {
	return f.Value
}

// Clone creates a deep copy of FhirString.
func (f *FhirString) Clone() *FhirString {
	if f == nil {
		return nil
	}
	return &FhirString{
		Value:   f.Value,
		Element: f.Element.Clone(),
	}
}

// Equals checks equality between two FhirString instances.
func (f *FhirString) Equals(other *FhirString) bool {
	if f == nil && other == nil {
		return true
	}
	if f == nil || other == nil {
		return false
	}
	return f.Value == other.Value && f.Element.Equals(other.Element)
}

// Utility methods for FhirString.
func (f *FhirString) Length() int {
	return len(f.Value)
}

func (f *FhirString) IsEmpty() bool {
	return strings.TrimSpace(f.Value) == ""
}

func (f *FhirString) Trim() string {
	return strings.TrimSpace(f.Value)
}

func (f *FhirString) Concat(other string) string {
	return f.Value + other
}

