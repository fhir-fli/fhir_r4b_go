package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"
)

// FhirCanonical represents the FHIR primitive type `canonical` (a URL).
type FhirCanonical struct {
	Value   *url.URL `json:"value,omitempty"`   // The URL value
	Element *Element `json:"_value,omitempty"` // Additional metadata (FHIR element)
}

// NewFhirCanonical creates a new FhirCanonical with validation.
func NewFhirCanonical(input string, element *Element) (*FhirCanonical, error) {
	parsed, err := validateCanonical(input)
	if err != nil {
		return nil, err
	}
	return &FhirCanonical{
		Value:   parsed,
		Element: element,
	}, nil
}

// FromJSON initializes a FhirCanonical from JSON input.
func (fc *FhirCanonical) FromJSON(data []byte) error {
	return json.Unmarshal(data, fc)
}

// ToJSON converts the FhirCanonical to JSON.
func (fc *FhirCanonical) ToJSON() ([]byte, error) {
	return json.Marshal(fc)
}

// TryParseCanonical attempts to parse an input into a FhirCanonical.
func TryParseCanonical(input interface{}) (*FhirCanonical, error) {
	switch v := input.(type) {
	case string:
		return NewFhirCanonical(v, nil)
	case *url.URL:
		return &FhirCanonical{Value: v}, nil
	default:
		return nil, errors.New("invalid input for FhirCanonical")
	}
}

// Helper: Validates the input as a URL.
func validateCanonical(input string) (*url.URL, error) {
	parsed, err := url.Parse(input)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return nil, errors.New("invalid canonical URL")
	}
	return parsed, nil
}

// CopyWith creates a modified copy of FhirCanonical.
func (fc *FhirCanonical) CopyWith(newValue *url.URL, element *Element) *FhirCanonical {
	return &FhirCanonical{
		Value:   ifNotNil(newValue, fc.Value),
		Element: ifNotNil(element, fc.Element),
	}
}

// Clone creates a deep copy of the FhirCanonical.
func (fc *FhirCanonical) Clone() *FhirCanonical {
	if fc == nil {
		return nil
	}
	var newElement *Element
	if fc.Element != nil {
		newElement = fc.Element.Clone()
	}
	return &FhirCanonical{
		Value:   cloneURL(fc.Value),
		Element: newElement,
	}
}

// Helper: Deep clone a *url.URL.
func cloneURL(u *url.URL) *url.URL {
	if u == nil {
		return nil
	}
	cloned := *u
	return &cloned
}

// Equals checks for equality between two FhirCanonical instances.
func (fc *FhirCanonical) Equals(other *FhirCanonical) bool {
	if fc == nil && other == nil {
		return true
	}
	if fc == nil || other == nil {
		return false
	}
	return equalURLs(fc.Value, other.Value) && equalElements(fc.Element, other.Element)
}

// String returns the canonical URL as a string.
func (fc *FhirCanonical) String() string {
	if fc.Value != nil {
		return fc.Value.String()
	}
	return ""
}

// PathSegments returns the path segments of the canonical URL.
func (fc *FhirCanonical) PathSegments() []string {
	if fc.Value != nil {
		return strings.Split(fc.Value.Path, "/")
	}
	return nil
}

// Query returns the query string of the canonical URL.
func (fc *FhirCanonical) Query() string {
	if fc.Value != nil {
		return fc.Value.RawQuery
	}
	return ""
}

// Host returns the host portion of the canonical URL.
func (fc *FhirCanonical) Host() string {
	if fc.Value != nil {
		return fc.Value.Host
	}
	return ""
}

// Helper: Check if two URLs are equal.
func equalURLs(a, b *url.URL) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.String() == b.String()
}

// FhirCanonicalList represents a list of FhirCanonical values.
type FhirCanonicalList []*FhirCanonical

// FromJSONList converts JSON input into a list of FhirCanonical values.
func FromJSONListFhirCanonical(values []interface{}, elements []interface{}) (FhirCanonicalList, error) {
	if elements != nil && len(values) != len(elements) {
		return nil, errors.New("values and elements must have the same length")
	}

	list := make(FhirCanonicalList, len(values))
	for i, v := range values {
		element := parseElement(elements, i)
		parsed, err := TryParseCanonical(v)
		if err != nil {
			return nil, err
		}
		list[i] = parsed.CopyWith(nil, element)
	}
	return list, nil
}

// ToJSONList converts a list of FhirCanonical into JSON-compatible maps.
func (list FhirCanonicalList) ToJSONList() map[string]interface{} {
	values := make([]interface{}, len(list))
	elements := make([]interface{}, len(list))

	for i, fc := range list {
		if fc != nil {
			values[i] = fc.Value.String()
			elements[i] = fc.Element
		} else {
			values[i] = nil
			elements[i] = nil
		}
	}

	return map[string]interface{}{
		"value":  values,
		"_value": elements,
	}
}
