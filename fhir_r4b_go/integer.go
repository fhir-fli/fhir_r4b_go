package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// FhirInteger represents the FHIR 'integer' type.
type FhirInteger struct {
	Value   *int64   `json:"-"`       // The actual value
	Element *Element `json:",inline"` // Metadata (FHIR element)
}

// NewFhirInteger creates a new validated FhirInteger.
func NewFhirInteger(value int64, element *Element) *FhirInteger {
	return &FhirInteger{Value: &value, Element: element}
}

// NewFhirIntegerFromMap creates a FhirInteger instance from a map.
func NewFhirIntegerFromMap(data map[string]interface{}) (*FhirInteger, error) {
	value, ok := data["value"].(*int64)
	if !ok {
		return nil, errors.New("invalid or missing value for FhirInteger")
	}

	integer := &FhirInteger{Value: value}

	if elementData, ok := data["_value"].(map[string]interface{}); ok {
		integer.Element = &Element{}
		if err := mapToStruct(elementData, integer.Element); err != nil {
			return nil, fmt.Errorf("failed to parse _value for FhirInteger: %v", err)
		}
	}

	return integer, nil
}

func (fi *FhirInteger) UnmarshalJSON(data []byte) error {
	temp := struct {
		Value   string   `json:"value"`
		Element *Element `json:"_value"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Parse the value string into an int64
	parsed, err := strconv.ParseInt(temp.Value, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid FhirInteger value: %w", err)
	}
	fi.Value = &parsed

	// Assign the Element if present
	fi.Element = temp.Element

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
