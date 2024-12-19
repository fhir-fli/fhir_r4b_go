package fhir_r4b_go

import (
	"errors"
	"fmt"
)

// FhirDateTime represents FHIR-compliant dateTime.
type FhirDateTime struct {
	FhirDateTimeBase
}

// NewFhirDateTimeFromComponents creates a new FhirDateTime from components.
func NewFhirDateTimeFromComponents(
	year int,
	month, day, hour, minute, second, millisecond *int,
	isUTC bool,
	offset *float64,
) (*FhirDateTime, error) {
	if year == 0 || month == nil || day == nil {
		return nil, errors.New("FHIR dateTime must include at least year, month, and day")
	}

	return &FhirDateTime{
		FhirDateTimeBase: *NewFhirDateTimeBase(&year, isUTC, month, day, hour, minute, second, millisecond, nil, offset, nil),
	}, nil
}

// MarshalJSON serializes FhirDateTime.
func (f *FhirDateTime) MarshalJSON() ([]byte, error) {
	return f.FhirDateTimeBase.MarshalJSON()
}

// UnmarshalJSON deserializes JSON into FhirDateTime.
func (f *FhirDateTime) UnmarshalJSON(data []byte) error {
	return f.FhirDateTimeBase.UnmarshalJSON(data)
}

func NewFhirDateTimeFromString(input string) (*FhirDateTime, error) {
	base, err := FhirDateTimeBaseFromString(input)
	if err != nil {
		return nil, err
	}

	if base.Year == nil || base.Month == nil || base.Day == nil {
		return nil, errors.New("FHIR dateTime must include at least year, month, and day")
	}

	base.Value = &input // Cache the parsed string
	return &FhirDateTime{FhirDateTimeBase: *base}, nil
}

// NewFhirDateTimeFromMap creates a new FhirDateTime from a map representation.
func NewFhirDateTimeFromMap(data map[string]interface{}) (*FhirDateTime, error) {
	// Validate that the map contains the necessary keys
	value, ok := data["value"].(string)
	if !ok || value == "" {
		return nil, fmt.Errorf("missing or invalid 'value' in map")
	}

	// Parse the date-time string into a FhirDateTimeBase
	base, err := FhirDateTimeBaseFromString(value)
	if err != nil {
		return nil, fmt.Errorf("failed to parse FHIR dateTime: %w", err)
	}

	// Check that required fields (year, month, day) are populated
	if base.Year == nil || base.Month == nil || base.Day == nil {
		return nil, fmt.Errorf("FHIR dateTime must include at least year, month, and day")
	}

	// Initialize the FhirDateTime instance
	result := &FhirDateTime{
		FhirDateTimeBase: *base,
	}

	// Parse the optional _value element if present
	if elementData, ok := data["_value"].(map[string]interface{}); ok {
		result.FhirDateTimeBase.Element = &Element{}
		if err := mapToStruct(elementData, result.FhirDateTimeBase.Element); err != nil {
			return nil, fmt.Errorf("failed to parse '_value' element: %w", err)
		}
	}

	return result, nil
}

// Clone creates a deep copy of FhirDateTime.
func (f *FhirDateTime) Clone() *FhirDateTime {
	if f == nil {
		return nil
	}
	return &FhirDateTime{FhirDateTimeBase: *f.FhirDateTimeBase.Clone()}
}

// Equals checks for equality between two FhirDateTime instances.
func (f *FhirDateTime) Equals(other *FhirDateTime) bool {
	if f == nil && other == nil {
		return true
	}
	if f == nil || other == nil {
		return false
	}

	return f.FhirDateTimeBase.CompareTo(&other.FhirDateTimeBase) == 0
}
