package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

// FhirUri represents the FHIR primitive type `canonical`.
type FhirUri struct {
	Value   *url.URL `json:"-"`       // The URL value
	Element *Element `json:",inline"` // Metadata (FHIR element)
}

// NewFhirUri creates a new FhirUri with validation.
func NewFhirUri(input string, element *Element) (*FhirUri, error) {
	parsed, err := url.Parse(input)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return nil, errors.New("invalid canonical URL")
	}
	return &FhirUri{
		Value:   parsed,
		Element: element,
	}, nil
}

// NewFhirUtiFromMap creates a FhirUti instance from a map.
func NewFhirUriFromMap(data map[string]interface{}) (*FhirUri, error) {
	valueStr, ok := data["value"].(string)
	if !ok {
		return nil, errors.New("missing or invalid value for FhirCanonical")
	}

	parsedValue, err := url.Parse(valueStr)
	if err != nil {
		return nil, fmt.Errorf("invalid URL value for FhirCanonical: %v", err)
	}

	uri := &FhirUri{Value: parsedValue}

	if elementData, ok := data["_value"].(map[string]interface{}); ok {
		uri.Element = &Element{}
		if err := mapToStruct(elementData, uri.Element); err != nil {
			return nil, fmt.Errorf("failed to parse _value for FhirCanonical: %v", err)
		}
	}

	return uri, nil
}

// UnmarshalJSON initializes a FhirUri from JSON input.
func (fc *FhirUri) UnmarshalJSON(data []byte) error {
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
		parsed, err := url.Parse(value)
		if err != nil || parsed.Scheme == "" || parsed.Host == "" {
			return errors.New("invalid canonical URL")
		}
		fc.Value = parsed
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

// MarshalJSON converts the FhirUri to JSON.
func (fc *FhirUri) MarshalJSON() ([]byte, error) {
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

// Clone creates a deep copy of the FhirUri.
func (fc *FhirUri) Clone() *FhirUri {
	if fc == nil {
		return nil
	}
	clone := &FhirUri{}
	if fc.Value != nil {
		val := *fc.Value
		clone.Value = &val
	}
	if fc.Element != nil {
		clone.Element = fc.Element.Clone()
	}
	return clone
}

// Equals checks for equality between two FhirUri instances.
func (fc *FhirUri) Equals(other *FhirUri) bool {
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
