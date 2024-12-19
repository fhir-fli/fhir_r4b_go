package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// FhirDecimal represents the FHIR 'decimal' type.
type FhirDecimal struct {
	Value   *float64 `json:"-"`       // The actual value
	Element *Element `json:",inline"` // Metadata (FHIR element)
}

// NewFhirDecimal creates a new validated FhirDecimal.
func NewFhirDecimal(input interface{}, element *Element) (*FhirDecimal, error) {
	val, err := parseDecimal(input)
	if err != nil {
		return nil, err
	}
	return &FhirDecimal{Value: val, Element: element}, nil
}

// NewFhirDecimalFromMap creates a FhirDecimal instance from a map.
func NewFhirDecimalFromMap(data map[string]interface{}) (*FhirDecimal, error) {
	value, ok := data["value"].(float64)
	if !ok {
		return nil, errors.New("invalid or missing value for FhirDecimal")
	}

	decimal := &FhirDecimal{Value: &value}

	if elementData, ok := data["_value"].(map[string]interface{}); ok {
		decimal.Element = &Element{}
		if err := mapToStruct(elementData, decimal.Element); err != nil {
			return nil, fmt.Errorf("failed to parse _value for FhirDecimal: %v", err)
		}
	}

	return decimal, nil
}

// UnmarshalJSON deserializes JSON into FhirDecimal.
func (fd *FhirDecimal) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if rawValue, exists := raw["value"]; exists {
		var value string
		if err := json.Unmarshal(rawValue, &value); err != nil {
			return err
		}
		parsed, err := parseDecimal(value)
		if err != nil {
			return err
		}
		fd.Value = parsed
	}

	if rawElement, exists := raw["_value"]; exists {
		fd.Element = &Element{}
		if err := json.Unmarshal(rawElement, fd.Element); err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON serializes FhirDecimal to JSON.
func (fd *FhirDecimal) MarshalJSON() ([]byte, error) {
	raw := map[string]interface{}{}
	if fd.Value != nil {
		raw["value"] = strconv.FormatFloat(*fd.Value, 'f', -1, 64)
	}
	if fd.Element != nil {
		raw["_value"] = fd.Element
	}
	return json.Marshal(raw)
}

// Clone creates a deep copy of FhirDecimal.
func (fd *FhirDecimal) Clone() *FhirDecimal {
	if fd == nil {
		return nil
	}
	return &FhirDecimal{
		Value:   floatPtrIfNotNil(fd.Value),
		Element: fd.Element.Clone(),
	}
}

// Equals checks for equality between two FhirDecimal instances.
func (fd *FhirDecimal) Equals(other *FhirDecimal) bool {
	if fd == nil && other == nil {
		return true
	}
	if fd == nil || other == nil {
		return false
	}
	return floatEquals(fd.Value, other.Value) && fd.Element.Equals(other.Element)
}

// parseDecimal parses input into a valid decimal value.
func parseDecimal(input interface{}) (*float64, error) {
	switch v := input.(type) {
	case float64:
		return &v, nil
	case string:
		parsed, err := strconv.ParseFloat(v, 64)
		if err != nil || !isValidDecimalString(v) {
			return nil, errors.New("invalid FHIR decimal format")
		}
		return &parsed, nil
	default:
		return nil, errors.New("unsupported type for FhirDecimal")
	}
}

// isValidDecimalString validates a decimal string format.
func isValidDecimalString(input string) bool {
	// FHIR decimal allows numeric values with optional leading '+' or '-' and a single '.'.
	return strings.Contains(input, ".")
}

// ToJSON serializes FhirDecimal to JSON.
func (fd *FhirDecimal) ToJSON() ([]byte, error) {
	data := map[string]interface{}{}
	if fd.Value != nil {
		data["value"] = strconv.FormatFloat(*fd.Value, 'f', -1, 64)
	}
	if fd.Element != nil {
		elementJSON, err := fd.Element.ToJSON()
		if err != nil {
			return nil, err
		}
		data["_value"] = elementJSON
	}
	return json.Marshal(data)
}

// FromJSON initializes a FhirDecimal from JSON input.
func (fd *FhirDecimal) FromJSON(data []byte) error {
	temp := struct {
		Value   string   `json:"value"`
		Element *Element `json:"_value"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	parsed, err := parseDecimal(temp.Value)
	if err != nil {
		return err
	}

	fd.Value = parsed
	fd.Element = temp.Element
	return nil
}

// Add performs addition with another FhirDecimal.
func (fd *FhirDecimal) Add(other *FhirDecimal) (*FhirDecimal, error) {
	if fd.Value == nil || other.Value == nil {
		return nil, errors.New("cannot perform addition: one or both values are nil")
	}
	result := *fd.Value + *other.Value
	return NewFhirDecimal(result, fd.Element.Clone())
}

// Subtract performs subtraction with another FhirDecimal.
func (fd *FhirDecimal) Subtract(other *FhirDecimal) (*FhirDecimal, error) {
	if fd.Value == nil || other.Value == nil {
		return nil, errors.New("cannot perform subtraction: one or both values are nil")
	}
	result := *fd.Value - *other.Value
	return NewFhirDecimal(result, fd.Element.Clone())
}

// Multiply performs multiplication with another FhirDecimal.
func (fd *FhirDecimal) Multiply(other *FhirDecimal) (*FhirDecimal, error) {
	if fd.Value == nil || other.Value == nil {
		return nil, errors.New("cannot perform multiplication: one or both values are nil")
	}
	result := *fd.Value * *other.Value
	return NewFhirDecimal(result, fd.Element.Clone())
}

// Divide performs division with another FhirDecimal.
func (fd *FhirDecimal) Divide(other *FhirDecimal) (*FhirDecimal, error) {
	if fd.Value == nil || other.Value == nil {
		return nil, errors.New("cannot perform division: one or both values are nil")
	}
	if *other.Value == 0 {
		return nil, errors.New("division by zero is not allowed")
	}
	result := *fd.Value / *other.Value
	return NewFhirDecimal(result, fd.Element.Clone())
}

// String provides a string representation of the FhirDecimal.
func (fd *FhirDecimal) String() string {
	if fd.Value == nil {
		return ""
	}
	return strconv.FormatFloat(*fd.Value, 'f', -1, 64)
}

func floatEquals(a, b *float64) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}
