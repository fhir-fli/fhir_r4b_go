package fhir_r4b_go

import (
	"encoding/json"
	"errors"
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
		FhirDateTimeBase: *NewFhirDateTimeBase(&year, isUTC, &month, &day, &hour, &minute, &second, millisecond, microsecond, &offset, nil),
	}
}

func NewFhirInstantFromString(input string) (*FhirInstant, error) {
	base, err := FhirDateTimeBaseFromString(input)
	if err != nil {
		return nil, err
	}

	if base.Year == nil || base.Month == nil || base.Day == nil {
		return nil, errors.New("FHIR instant must include at least year, month, and day")
	}

	base.Value = &input // Cache the parsed string
	return &FhirInstant{FhirDateTimeBase: *base}, nil
}

// MarshalJSON serializes FhirInstant.
func (f *FhirInstant) MarshalJSON() ([]byte, error) {
	return f.FhirDateTimeBase.MarshalJSON()
}

func (f *FhirInstant) UnmarshalJSON(data []byte) error {
	temp := struct {
		Value   string   `json:"value"`
		Element *Element `json:"_value"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	parsed, err := NewFhirInstantFromString(temp.Value)
	if err != nil {
		return err
	}

	if temp.Element != nil {
		parsed.FhirDateTimeBase.Element = temp.Element
	}

	*f = *parsed
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
