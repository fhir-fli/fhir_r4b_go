package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"net/url"
)

// FhirUrl represents the FHIR primitive type `canonical`.
type FhirUrl struct {
	Value   *url.URL `json:"-"`          // The URL value
	Element *Element `json:",inline"`    // Metadata (FHIR element)
}

// NewFhirUrl creates a new FhirUrl with validation.
func NewFhirUrl(input string, element *Element) (*FhirUrl, error) {
	parsed, err := url.Parse(input)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return nil, errors.New("invalid canonical URL")
	}
	return &FhirUrl{
		Value:   parsed,
		Element: element,
	}, nil
}

// UnmarshalJSON initializes a FhirUrl from JSON input.
func (fc *FhirUrl) UnmarshalJSON(data []byte) error {
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

// MarshalJSON converts the FhirUrl to JSON.
func (fc *FhirUrl) MarshalJSON() ([]byte, error) {
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

// Clone creates a deep copy of the FhirUrl.
func (fc *FhirUrl) Clone() *FhirUrl {
	if fc == nil {
		return nil
	}
	clone := &FhirUrl{}
	if fc.Value != nil {
		val := *fc.Value
		clone.Value = &val
	}
	if fc.Element != nil {
		clone.Element = fc.Element.Clone()
	}
	return clone
}

// Equals checks for equality between two FhirUrl instances.
func (fc *FhirUrl) Equals(other *FhirUrl) bool {
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
