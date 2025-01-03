package fhir_r4b_go

import (
	"encoding/json"
	"errors"
)

// FhirDate represents FHIR-compliant date (year-month-day).
type FhirDate struct {
	FhirDateTimeBase
}

// NewFhirDateFromComponents creates a new FhirDate.
func NewFhirDateFromComponents(year int, month, day *int, isUTC bool) *FhirDate {
	return &FhirDate{
		FhirDateTimeBase: *NewFhirDateTimeBase(&year, isUTC, month, day, nil, nil, nil, nil, nil, nil, nil),
	}
}

// MarshalJSON serializes FhirDate.
func (f *FhirDate) MarshalJSON() ([]byte, error) {
	return f.FhirDateTimeBase.MarshalJSON()
}

func (f *FhirDate) UnmarshalJSON(data []byte) error {
	temp := struct {
		Value   string   `json:"value"`
		Element *Element `json:"_value"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	parsed, err := NewFhirDateFromString(temp.Value)
	if err != nil {
		return err
	}

	if temp.Element != nil {
		parsed.FhirDateTimeBase.Element = temp.Element
	}

	*f = *parsed
	return nil
}

func NewFhirDateFromString(input string) (*FhirDate, error) {
	base, err := FhirDateTimeBaseFromString(input)
	if err != nil {
		return nil, err
	}

	if base.Year == nil || base.Month == nil || base.Day == nil {
		return nil, errors.New("FHIR date must include at least year, month, and day")
	}

	base.Value = &input // Cache the parsed string
	return &FhirDate{FhirDateTimeBase: *base}, nil
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
