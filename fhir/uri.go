package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"
)

// FhirUri represents a validated URI in FHIR resources.
type FhirUri struct {
	Value   *url.URL                `json:"value,omitempty"`
	Element *Element `json:"_value,omitempty"`
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
		Value   string                  `json:"value"`
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

// String returns the URI as a string.
func (f *FhirUri) String() string {
	if f.Value != nil {
		return f.Value.String()
	}
	return ""
}

// Host returns the host portion of the URI.
func (f *FhirUri) Host() string {
	if f.Value != nil {
		return f.Value.Host
	}
	return ""
}

// PathSegments splits the URL path into segments.
func (f *FhirUri) PathSegments() []string {
	if f.Value != nil {
		parsedURL, err := url.Parse(f.Value.String())
		if err == nil {
			return splitPathSegments(parsedURL.Path)
		}
	}
	return nil
}

// Helper function to split and clean path segments
func splitPathSegments(path string) []string {
	segments := []string{}
	for _, segment := range strings.Split(path, "/") {
		if segment != "" {
			segments = append(segments, segment)
		}
	}
	return segments
}

// Clone creates a deep copy of FhirUri.
func (f *FhirUri) Clone() *FhirUri {
	var elementCopy *Element
	if f.Element != nil {
		elementCopy = f.Element.Copy()
	}
	return &FhirUri{
		Value:   f.Value,
		Element: elementCopy,
	}
}

// Equal checks equality between two FhirUri instances.
func (f *FhirUri) Equal(other *FhirUri) bool {
	return f.String() == other.String() && compareElements(f.Element, other.Element)
}
