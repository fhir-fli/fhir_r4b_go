package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

// FhirCanonical represents the FHIR primitive type `canonical`.
type FhirCanonical struct {
	Value   *url.URL `json:"-"`       // The URL value
	Element *Element `json:",inline"` // Metadata (FHIR element)
}

// NewFhirCanonical creates a new FhirCanonical with validation.
func NewFhirCanonical(input string, element *Element) (*FhirCanonical, error) {
	parsed, err := url.Parse(input)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return nil, errors.New("invalid canonical URL")
	}
	return &FhirCanonical{
		Value:   parsed,
		Element: element,
	}, nil
}

// NewFhirCanonicalFromMap creates a FhirCanonical instance from a map.
func NewFhirCanonicalFromMap(data map[string]interface{}) (*FhirCanonical, error) {
	valueStr, ok := data["value"].(string)
	if !ok {
		return nil, errors.New("missing or invalid value for FhirCanonical")
	}

	parsedValue, err := url.Parse(valueStr)
	if err != nil {
		return nil, fmt.Errorf("invalid URL value for FhirCanonical: %v", err)
	}

	canonical := &FhirCanonical{Value: parsedValue}

	if elementData, ok := data["_value"].(map[string]interface{}); ok {
		canonical.Element = &Element{}
		if err := mapToStruct(elementData, canonical.Element); err != nil {
			return nil, fmt.Errorf("failed to parse _value for FhirCanonical: %v", err)
		}
	}

	return canonical, nil
}

func (fc *FhirCanonical) UnmarshalJSON(data []byte) error {
	temp := struct {
		Value   string   `json:"value"`
		Element *Element `json:"_value"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Validate and parse the canonical URL
	parsed, err := url.Parse(temp.Value)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return errors.New("invalid FhirCanonical: missing or invalid URL")
	}
	fc.Value = parsed

	// Assign the Element if present
	fc.Element = temp.Element

	return nil
}

// MarshalJSON converts the FhirCanonical to JSON.
func (fc *FhirCanonical) MarshalJSON() ([]byte, error) {
	raw := make(map[string]interface{})

	if fc.Value != nil {
		raw["value"] = fc.Value.String()
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

// Clone creates a deep copy of the FhirCanonical.
func (fc *FhirCanonical) Clone() *FhirCanonical {
	if fc == nil {
		return nil
	}
	clone := &FhirCanonical{}
	if fc.Value != nil {
		val := *fc.Value
		clone.Value = &val
	}
	if fc.Element != nil {
		clone.Element = fc.Element.Clone()
	}
	return clone
}

// Equals checks for equality between two FhirCanonical instances.
func (fc *FhirCanonical) Equals(other *FhirCanonical) bool {
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
