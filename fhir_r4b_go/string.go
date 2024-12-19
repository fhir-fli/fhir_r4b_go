package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// FhirString represents a validated string used in FHIR resources.
type FhirString struct {
	Value   *string  `json:"value,omitempty"`
	Element *Element `json:"_value,omitempty"`
}

// NewFhirString creates a new FhirString with validation.
func NewFhirString(input string, element *Element) (*FhirString, error) {
	if input == "" && element == nil {
		return nil, errors.New("a value or element is required")
	}
	return &FhirString{
		Value:   &input,
		Element: element,
	}, nil
}

// NewFhirStringFromMap creates a FhirString instance from a map.
func NewFhirStringFromMap(data map[string]interface{}) (*FhirString, error) {
	value, ok := data["value"].(string)
	if !ok {
		return nil, errors.New("invalid or missing value for FhirString")
	}

	str := &FhirString{Value: &value}

	if elementData, ok := data["_value"].(map[string]interface{}); ok {
		str.Element = &Element{}
		if err := mapToStruct(elementData, str.Element); err != nil {
			return nil, fmt.Errorf("failed to parse _value for FhirString: %v", err)
		}
	}

	return str, nil
}

// MarshalJSON serializes FhirString into JSON.
func (f *FhirString) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{}
	if f.Value != nil && *f.Value != "" {
		data["value"] = f.Value
	}
	if f.Element != nil {
		data["_value"] = f.Element
	}
	return json.Marshal(data)
}

func (f *FhirString) UnmarshalJSON(data []byte) error {
	var temp interface{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Handle case where the input is a simple string
	if str, ok := temp.(string); ok {
		f.Value = &str
		f.Element = nil
		return nil
	}

	// Handle case where the input is a JSON object
	tempStruct := struct {
		Value   *string  `json:"value"`
		Element *Element `json:"_value"`
	}{}
	if err := json.Unmarshal(data, &tempStruct); err != nil {
		return err
	}

	f.Value = tempStruct.Value
	f.Element = tempStruct.Element
	return nil
}

// String returns the string value or an empty string if nil.
func (f *FhirString) String() string {
	if f.Value == nil {
		return ""
	}
	return *f.Value
}

// Clone creates a deep copy of FhirString.
func (f *FhirString) Clone() *FhirString {
	if f == nil {
		return nil
	}
	var clonedValue *string
	if f.Value != nil {
		val := *f.Value
		clonedValue = &val
	}
	return &FhirString{
		Value:   clonedValue,
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
	if (f.Value == nil && other.Value != nil) || (f.Value != nil && other.Value == nil) {
		return false
	}
	return (f.Value == nil || *f.Value == *other.Value) && f.Element.Equals(other.Element)
}

// Utility methods for FhirString.
func (f *FhirString) Length() int {
	if f.Value == nil {
		return 0
	}
	return len(*f.Value)
}

func (f *FhirString) IsEmpty() bool {
	if f.Value == nil {
		return true
	}
	return strings.TrimSpace(*f.Value) == ""
}

func (f *FhirString) Trim() string {
	if f.Value == nil {
		return ""
	}
	return strings.TrimSpace(*f.Value)
}

func (f *FhirString) Concat(other string) string {
	if f.Value == nil {
		return other
	}
	return *f.Value + other
}
