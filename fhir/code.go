package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"regexp"
)

// FhirCode represents the FHIR primitive type `code`.
type FhirCode struct {
	Value   *string  `json:"value,omitempty"`  // The actual code value
	Element *Element `json:"_value,omitempty"` // Additional metadata (FHIR element)
}

// NewFhirCode creates a new FhirCode with validation.
func NewFhirCode(input string, element *Element) (*FhirCode, error) {
	if err := validateFhirCode(input); err != nil {
		return nil, err
	}
	return &FhirCode{
		Value:   &input,
		Element: element,
	}, nil
}

// FromJSON initializes a FhirCode from JSON input.
func (fc *FhirCode) FromJSON(data []byte) error {
	return json.Unmarshal(data, fc)
}

// ToJSON serializes the FhirCode to JSON.
func (fc *FhirCode) ToJSON() ([]byte, error) {
	return json.Marshal(fc)
}

// Clone creates a deep copy of the FhirCode.
func (fc *FhirCode) Clone() *FhirCode {
	if fc == nil {
		return nil
	}
	return &FhirCode{
		Value:   fc.Value,
		Element: fc.Element.Clone(),
	}
}

// Equals checks for equality between two FhirCode instances.
func (fc *FhirCode) Equals(other *FhirCode) bool {
	if fc == nil && other == nil {
		return true
	}
	if fc == nil || other == nil {
		return false
	}
	return fc.Value == other.Value && fc.Element.Equals(other.Element)
}

// Helper: validateFhirCode checks if the input string matches the FHIR `code` format.
func validateFhirCode(input string) error {
	regex := regexp.MustCompile(`^[^\s]+(\s[^\s]+)*$`)
	if regex.MatchString(input) {
		return nil
	}
	return errors.New("invalid FHIR Code: does not match required format")
}
