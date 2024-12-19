package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// FhirPositiveInt represents the FHIR 'integer' type.
type FhirPositiveInt struct {
	Value   *int64   `json:"-"`       // The actual value
	Element *Element `json:",inline"` // Metadata (FHIR element)
}

// NewFhirPositiveInt creates a new validated FhirPositiveInt.
func NewFhirPositiveInt(value int64, element *Element) *FhirPositiveInt {
	return &FhirPositiveInt{Value: &value, Element: element}
}

// NewFhirPositiveIntFromMap creates a FhirPositiveInt instance from a map.
func NewFhirPositiveIntFromMap(data map[string]interface{}) (*FhirPositiveInt, error) {
	value, ok := data["value"].(*int64)
	if !ok {
		return nil, errors.New("invalid or missing value for FhirPositiveInt")
	}

	integer := &FhirPositiveInt{Value: value}

	if elementData, ok := data["_value"].(map[string]interface{}); ok {
		integer.Element = &Element{}
		if err := mapToStruct(elementData, integer.Element); err != nil {
			return nil, fmt.Errorf("failed to parse _value for FhirPositiveInt: %v", err)
		}
	}

	return integer, nil
}

func (fi *FhirPositiveInt) UnmarshalJSON(data []byte) error {
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
		return fmt.Errorf("invalid FhirPositiveInt value: %w", err)
	}
	fi.Value = &parsed

	// Assign the Element if present
	fi.Element = temp.Element

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
