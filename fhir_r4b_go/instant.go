package fhir_r4b_go

import (
	"encoding/json"
)

// FhirInstant represents an instant in time with full precision.
type FhirInstant struct {
	FhirDateTimeBase
}

// NewFhirInstantFromComponents creates a new FhirInstant.
func NewFhirInstantFromComponents(
	year, month, day, hour, minute, second int,
	millisecond *int,
	microsecond *string,
	offset float64,
	isUTC bool,
) *FhirInstant {
	return &FhirInstant{
		FhirDateTimeBase: *NewFhirDateTimeBase(&year, isUTC, &month, &day, &hour, &minute, &second, millisecond, microsecond, &offset),
	}
}

// MarshalJSON serializes FhirInstant.
func (f *FhirInstant) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{"value": f.ToString()})
}

// UnmarshalJSON deserializes JSON into FhirInstant.
func (f *FhirInstant) UnmarshalJSON(data []byte) error {
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
	*f = FhirInstant{FhirDateTimeBase: *parsed.FhirDateTimeBase.Clone()}
	return nil
}

// Clone creates a deep copy of FhirInstant.
func (f *FhirInstant) Clone() *FhirInstant {
	if f == nil {
		return nil
	}
	return &FhirInstant{FhirDateTimeBase: *f.FhirDateTimeBase.Clone()}
}

// Equals checks equality between two FhirInstant instances.
func (f *FhirInstant) Equals(other *FhirInstant) bool {
	if f == nil && other == nil {
		return true
	}
	if f == nil || other == nil {
		return false
	}
	return f.FhirDateTimeBase.CompareTo(&other.FhirDateTimeBase) == 0
}
