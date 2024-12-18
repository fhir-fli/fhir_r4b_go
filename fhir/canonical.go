package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"net/url"
)

// FhirCanonical represents the FHIR primitive type `canonical`.
type FhirCanonical struct {
	Value   *url.URL `json:"value,omitempty"`  // The URL value
	Element *Element `json:"_value,omitempty"` // Additional metadata (FHIR element)
}

// NewFhirCanonical creates a new FhirCanonical with validation.
func NewFhirCanonical(input string, element *Element) (*FhirCanonical, error) {
	parsed, err := validateCanonical(input)
	if err != nil {
		return nil, err
	}
	return &FhirCanonical{
		Value:   parsed,
		Element: element,
	}, nil
}

// FromJSON initializes a FhirCanonical from JSON input.
func (fc *FhirCanonical) FromJSON(data []byte) error {
	return json.Unmarshal(data, fc)
}

// ToJSON converts the FhirCanonical to JSON.
func (fc *FhirCanonical) ToJSON() ([]byte, error) {
	return json.Marshal(fc)
}

// Clone creates a deep copy of the FhirCanonical.
func (fc *FhirCanonical) Clone() *FhirCanonical {
	if fc == nil {
		return nil
	}
	return &FhirCanonical{
		Value:   fc.Value,
		Element: fc.Element.Clone(),
	}
}

// Equals checks for equality between two FhirCanonical instances.
func (fc *FhirCanonical) Equals(other *FhirCanonical) bool {
	if fc == nil && other == nil {
		return true
	}
	if fc == nil || other == nil {
		return false
	}
	return fc.Value == other.Value && fc.Element.Equals(other.Element)
}

// Helper: Validates the input as a URL.
func validateCanonical(input string) (*url.URL, error) {
	parsed, err := url.Parse(input)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return nil, errors.New("invalid canonical URL")
	}
	return parsed, nil
}
