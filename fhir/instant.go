package fhir_r4b_go

import (
	"encoding/json"
	"time"
)

// FhirInstant represents an instant in time with full precision.
type FhirInstant struct {
	FhirDateTimeBase
}

// NewFhirInstantFromComponents creates a new FhirInstant.
func NewFhirInstantFromComponents(year, month, day, hour, minute, second int, millisecond *int, microsecond *string, offset float64, isUTC bool) FhirInstant {
	return FhirInstant{
		FhirDateTimeBase: NewFhirDateTimeBase(&year, isUTC, &month, &day, &hour, &minute, &second, millisecond, microsecond, &offset),
	}
}

// NewFhirInstantFromTime creates a FhirInstant from Go's time.Time.
func NewFhirInstantFromTime(t time.Time) FhirInstant {
	_, offsetInSeconds := t.Zone() // Extract only the offset
	offset := float64(offsetInSeconds) / 3600
	return FhirInstant{
		FhirDateTimeBase: NewFhirDateTimeBase(
			intPtr(t.Year()), t.Location() == time.UTC,
			intPtr(int(t.Month())), intPtr(t.Day()), intPtr(t.Hour()),
			intPtr(t.Minute()), intPtr(t.Second()), intPtr(t.Nanosecond()/1e6),
			nil, &offset,
		),
	}
}

// MarshalJSON serializes FhirInstant.
func (f FhirInstant) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{"value": f.ToString()})
}
