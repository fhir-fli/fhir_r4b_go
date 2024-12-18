package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"net/url"
)

// FhirUri represents a validated URI in FHIR resources.
type FhirUri struct {
	Value   *url.URL  `json:"value,omitempty"`
	Element *Element  `json:"_value,omitempty"`
}

// NewFhirUri creates a new FhirUri instance with validation.
func NewFhirUri(input string, element *Element) (*FhirUri, error) {
	parsedURL, err := url.Parse(input)
	if err != nil || parsedURL.Scheme == "" {
		return nil, errors.New("invalid URI: " + input)
	}
	return &FhirUri{Value: parsedURL, Element: element}, nil
}

// MarshalJSON serializes FhirUri into JSON.
func (f *FhirUri) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{}
	if f.Value != nil {
		data["value"] = f.Value.String()
	}
	if f.Element != nil {
		data["_value"] = f.Element
	}
	return json.Marshal(data)
}

// UnmarshalJSON deserializes JSON into FhirUri.
func (f *FhirUri) UnmarshalJSON(data []byte) error {
	temp := struct {
		Value   string   `json:"value"`
		Element *Element `json:"_value"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	parsedURL, err := url.Parse(temp.Value)
	if err != nil {
		return errors.New("invalid URI: " + temp.Value)
	}

	f.Value = parsedURL
	f.Element = temp.Element
	return nil
}

// Clone creates a deep copy of FhirUri.
func (f *FhirUri) Clone() *FhirUri {
	if f == nil {
		return nil
	}
	var elementCopy *Element
	if f.Element != nil {
		elementCopy = f.Element.Clone()
	}
	return &FhirUri{
		Value:   f.Value, // URLs are immutable; cloning is not required
		Element: elementCopy,
	}
}

// Equals checks equality between two FhirUri instances.
func (f *FhirUri) Equals(other *FhirUri) bool {
	if f == nil && other == nil {
		return true
	}
	if f == nil || other == nil {
		return false
	}
	return f.String() == other.String() && f.Element.Equals(other.Element)
}

// String returns the URI as a string.
func (f *FhirUri) String() string {
	if f.Value != nil {
		return f.Value.String()
	}
	return ""
}
