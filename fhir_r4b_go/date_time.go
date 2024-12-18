package fhir_r4b_go

import (
	"encoding/json"
	"errors"
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
		FhirDateTimeBase: *NewFhirDateTimeBase(&year, isUTC, month, day, hour, minute, second, millisecond, nil, offset),
	}, nil
}

// NewFhirDateTimeFromString parses a FHIR-compliant dateTime string.
func NewFhirDateTimeFromString(input string) (*FhirDateTime, error) {
	base, err := FhirDateTimeBaseFromString(input)
	if err != nil {
		return nil, err
	}

	if base.Year == nil || base.Month == nil || base.Day == nil {
		return nil, errors.New("FHIR dateTime must include at least year, month, and day")
	}

	return &FhirDateTime{FhirDateTimeBase: *base}, nil
}

// MarshalJSON serializes FhirDateTime.
func (f *FhirDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{"value": f.ToString()})
}

// UnmarshalJSON deserializes JSON into FhirDateTime.
func (f *FhirDateTime) UnmarshalJSON(data []byte) error {
	var input map[string]string
	if err := json.Unmarshal(data, &input); err != nil {
		return err
	}
	value, ok := input["value"]
	if !ok {
		return nil
	}
	parsed, err := NewFhirDateTimeFromString(value)
	if err != nil {
		return err
	}
	*f = *parsed
	return nil
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
