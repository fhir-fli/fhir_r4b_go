package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"regexp"
)

// FhirCode represents the FHIR primitive type `code`.
type FhirCode struct {
	Value   *string  `json:"-"`          // The actual code value
	Element *Element `json:",inline"`    // Metadata (FHIR element)
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

// UnmarshalJSON initializes a FhirCode from JSON input.
func (fc *FhirCode) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Extract value
	if rawValue, exists := raw["value"]; exists {
		var value string
		if err := json.Unmarshal(rawValue, &value); err != nil {
			return err
		}
		if err := validateFhirCode(value); err != nil {
			return err
		}
		fc.Value = &value
	}

	// Extract metadata
	if rawElement, exists := raw["_value"]; exists {
		fc.Element = &Element{}
		if err := json.Unmarshal(rawElement, fc.Element); err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON converts the FhirCode to JSON.
func (fc *FhirCode) MarshalJSON() ([]byte, error) {
	raw := make(map[string]interface{})

	if fc.Value != nil {
		raw["value"] = *fc.Value
	}
	if fc.Element != nil {
		elementJSON, err := json.Marshal(fc.Element)
		if err != nil {
			return nil, err
		}
		var elementMap map[string]interface{}
		if err := json.Unmarshal(elementJSON, &elementMap); err != nil {
			return nil, err
		}
		raw["_value"] = elementMap
	}

	return json.Marshal(raw)
}

// validateFhirCode checks if the input string matches the FHIR `code` format.
func validateFhirCode(input string) error {
	regex := regexp.MustCompile(`^[^\s]+(\s[^\s]+)*$`)
	if regex.MatchString(input) {
		return nil
	}
	return errors.New("invalid FHIR Code: does not match required format")
}

// Clone creates a deep copy of the FhirCode.
func (fc *FhirCode) Clone() *FhirCode {
	if fc == nil {
		return nil
	}
	clone := &FhirCode{}
	if fc.Value != nil {
		val := *fc.Value
		clone.Value = &val
	}
	if fc.Element != nil {
		clone.Element = fc.Element.Clone()
	}
	return clone
}

// Equals checks for equality between two FhirCode instances.
func (fc *FhirCode) Equals(other *FhirCode) bool {
	if fc == nil && other == nil {
		return true
	}
	if fc == nil || other == nil {
		return false
	}
	if (fc.Value == nil && other.Value != nil) || (fc.Value != nil && other.Value == nil) {
		return false
	}
	if fc.Value != nil && *fc.Value != *other.Value {
		return false
	}
	return fc.Element.Equals(other.Element)
}
