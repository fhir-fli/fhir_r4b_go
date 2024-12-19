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

// NewFhirUriFromMap creates a FhirUri instance from a map.
func NewFhirUriFromMap(data map[string]interface{}) (*FhirUri, error) {
	valueStr, ok := data["value"].(string)
	if !ok {
		return nil, errors.New("missing or invalid value for FhirUri")
	}

	parsedValue, err := url.Parse(valueStr)
	if err != nil {
		return nil, fmt.Errorf("invalid URL value for FhirUri: %v", err)
	}

	canonical := &FhirUri{Value: parsedValue}

	if elementData, ok := data["_value"].(map[string]interface{}); ok {
		canonical.Element = &Element{}
		if err := mapToStruct(elementData, canonical.Element); err != nil {
			return nil, fmt.Errorf("failed to parse _value for FhirUri: %v", err)
		}
	}

	return canonical, nil
}

func (fc *FhirUri) UnmarshalJSON(data []byte) error {
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
		return errors.New("invalid FhirUri: missing or invalid URL")
	}
	fc.Value = parsed

	// Assign the Element if present
	fc.Element = temp.Element

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
