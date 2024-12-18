package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"regexp"
)

// FhirId represents the FHIR primitive type 'id'
type FhirId struct {
	Value   *string `json:"value,omitempty"`   // ID value
	Element *Element `json:"_value,omitempty"` // Optional metadata element
}

// NewFhirId creates a new FhirId with validation
func NewFhirId(input string, element *Element) (*FhirId, error) {
	if err := validateFhirId(input); err != nil {
		return nil, err
	}
	return &FhirId{Value: &input, Element: element}, nil
}

// validateFhirId ensures input matches FHIR ID format
func validateFhirId(input string) error {
	regex := regexp.MustCompile(`^[A-Za-z0-9\-\.][A-Za-z0-9\-._]{0,63}$`)
	if !regex.MatchString(input) {
		return errors.New("invalid FhirId format")
	}
	return nil
}

// FromJSON creates FhirId from JSON
func (fi *FhirId) FromJSON(data []byte) error {
	return json.Unmarshal(data, fi)
}

// ToJSON converts FhirId to JSON map
func (fi *FhirId) ToJSON() (map[string]interface{}, error) {
	result := make(map[string]interface{})
	if fi.Value != nil {
		result["value"] = *fi.Value
	}
	if fi.Element != nil {
		elementJSON, err := fi.Element.ToJSON()
		if err != nil {
			return nil, err
		}
		result["_value"] = elementJSON
	}
	return result, nil
}

// Equals compares two FhirId instances
func (fi *FhirId) Equals(other *FhirId) bool {
	if fi == nil || other == nil {
		return false
	}
	if fi.Value == nil && other.Value == nil {
		return fi.Element.Equals(other.Element)
	}
	if fi.Value != nil && other.Value != nil {
		return *fi.Value == *other.Value && fi.Element.Equals(other.Element)
	}
	return false
}

// Clone creates a deep copy of FhirId
func (fi *FhirId) Clone() *FhirId {
	if fi == nil {
		return nil
	}
	clonedElement := fi.Element.Clone()
	clonedValue := *fi.Value
	return &FhirId{Value: &clonedValue, Element: clonedElement}
}
