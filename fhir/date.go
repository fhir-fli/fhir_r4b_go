package fhir_r4b_go

import (
	"encoding/json"
)

// FhirDate represents FHIR-compliant date (year-month-day).
type FhirDate struct {
	FhirDateTimeBase
}

// NewFhirDateFromComponents creates a new FhirDate.
func NewFhirDateFromComponents(year int, month, day *int, isUTC bool) *FhirDate {
	return &FhirDate{
		FhirDateTimeBase: *NewFhirDateTimeBase(&year, isUTC, month, day, nil, nil, nil, nil, nil, nil),
	}
}

// NewFhirDateFromString parses a date string into FhirDate.
func NewFhirDateFromString(input string) (*FhirDate, error) {
	base, err := FhirDateTimeBaseFromString(input)
	if err != nil {
		return nil, err
	}
	return &FhirDate{FhirDateTimeBase: *base}, nil
}

// MarshalJSON serializes FhirDate.
func (f *FhirDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{"value": f.ToString()})
}

// Clone creates a deep copy of FhirDate.
func (f *FhirDate) Clone() *FhirDate {
	if f == nil {
		return nil
	}
	return &FhirDate{FhirDateTimeBase: *f.FhirDateTimeBase.Clone()}
}

// Equals checks equality between two FhirDate instances.
func (f *FhirDate) Equals(other *FhirDate) bool {
	if f == nil && other == nil {
		return true
	}
	if f == nil || other == nil {
		return false
	}
	return f.FhirDateTimeBase.CompareTo(&other.FhirDateTimeBase) == 0
}
