package fhir_r4b_go

import (
	"encoding/json"
	"time"
)

// FhirDateTime extends FhirDateTimeBase for FHIR-compliant dateTime.
type FhirDateTime struct {
	FhirDateTimeBase
}

// NewFhirDateTimeFromComponents creates a new FhirDateTime from components.
func NewFhirDateTimeFromComponents(year int, month, day, hour, minute, second, millisecond *int, isUTC bool, offset *float64) FhirDateTime {
	return FhirDateTime{
		FhirDateTimeBase: NewFhirDateTimeBase(&year, isUTC, month, day, hour, minute, second, millisecond, nil, offset),
	}
}

// NewFhirDateTimeFromString parses a FHIR-compliant dateTime string.
func NewFhirDateTimeFromString(input string) (*FhirDateTime, error) {
	base, err := FromString(input)
	if err != nil {
		return nil, err
	}
	return &FhirDateTime{FhirDateTimeBase: *base}, nil
}

// AddDuration adds a time.Duration to FhirDateTime.
func (f FhirDateTime) AddDuration(d time.Duration) FhirDateTime {
	t := f.Value().Add(d)
	return NewFhirDateTimeFromTime(t)
}

// MarshalJSON serializes FhirDateTime.
func (f FhirDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{"value": f.ToString()})
}

// NewFhirDateTimeFromTime creates a FhirDateTime from Go's time.Time.
// NewFhirDateTimeFromTime creates a FhirDateTime from Go's time.Time.
func NewFhirDateTimeFromTime(t time.Time) FhirDateTime {
	_, zoneOffset := t.Zone() // Capture the zone name (ignore it) and the offset in seconds
	offset := float64(zoneOffset) / 3600 // Convert offset from seconds to hours

	return FhirDateTime{
		FhirDateTimeBase: NewFhirDateTimeBase(
			intPtr(t.Year()), t.Location() == time.UTC,
			intPtr(int(t.Month())), intPtr(t.Day()), intPtr(t.Hour()),
			intPtr(t.Minute()), intPtr(t.Second()), intPtr(t.Nanosecond()/1e6),
			nil, &offset,
		),
	}
}
