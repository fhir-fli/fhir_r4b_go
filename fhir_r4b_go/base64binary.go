package fhir_r4b_go

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

// FhirBase64Binary represents the FHIR primitive type `base64Binary`.
type FhirBase64Binary struct {
	Value   *string  `json:"-"`       // The Base64 encoded string
	Element *Element `json:",inline"` // Metadata (FHIR element)
}

// NewFhirBase64Binary creates a new FhirBase64Binary with validation.
func NewFhirBase64Binary(input string, element *Element) (*FhirBase64Binary, error) {
	cleanedInput := strings.ReplaceAll(input, " ", "")
	if !isValidBase64(cleanedInput) {
		return nil, errors.New("invalid Base64 encoded string")
	}
	return &FhirBase64Binary{
		Value:   &cleanedInput,
		Element: element,
	}, nil
}

// UnmarshalJSON initializes a FhirBase64Binary from JSON input.
func (b *FhirBase64Binary) UnmarshalJSON(data []byte) error {
	// Handle case where the input is a simple string
	var strValue string
	if err := json.Unmarshal(data, &strValue); err == nil {
		if !isValidBase64(strValue) {
			return errors.New("invalid Base64 encoded string")
		}
		b.Value = &strValue
		b.Element = nil
		return nil
	}

	// Handle case where the input is a JSON object
	var raw struct {
		Value   *string  `json:"value"`
		Element *Element `json:"_value"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if raw.Value != nil && !isValidBase64(*raw.Value) {
		return errors.New("invalid Base64 encoded string")
	}

	b.Value = raw.Value
	b.Element = raw.Element
	return nil
}

// MarshalJSON converts the FhirBase64Binary to JSON.
func (b *FhirBase64Binary) MarshalJSON() ([]byte, error) {
	raw := make(map[string]interface{})

	if b.Value != nil {
		raw["value"] = *b.Value
	}
	if b.Element != nil {
		elementJSON, err := json.Marshal(b.Element)
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

// isValidBase64 checks if a string is a valid Base64 encoded string.
func isValidBase64(input string) bool {
	if len(input)%4 != 0 {
		return false
	}
	_, err := base64.StdEncoding.DecodeString(input)
	return err == nil
}

// Clone creates a deep copy of FhirBase64Binary.
func (b *FhirBase64Binary) Clone() *FhirBase64Binary {
	if b == nil {
		return nil
	}
	clone := &FhirBase64Binary{}
	if b.Value != nil {
		val := *b.Value
		clone.Value = &val
	}
	if b.Element != nil {
		clone.Element = b.Element.Clone()
	}
	return clone
}

// Equals checks for equality between two FhirBase64Binary instances.
func (b *FhirBase64Binary) Equals(other *FhirBase64Binary) bool {
	if b == nil && other == nil {
		return true
	}
	if b == nil || other == nil {
		return false
	}
	if (b.Value == nil && other.Value != nil) || (b.Value != nil && other.Value == nil) {
		return false
	}
	if b.Value != nil && *b.Value != *other.Value {
		return false
	}
	return b.Element.Equals(other.Element)
}
