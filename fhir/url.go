package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"net/url"
)

// FhirUrl represents a validated URL in the FHIR standard.
type FhirUrl struct {
	Value   *url.URL `json:"value,omitempty"`
	Element *Element `json:"_value,omitempty"`
}

// NewFhirUrl creates a new FhirUrl instance with validation.
func NewFhirUrl(input string, element *Element) (*FhirUrl, error) {
	parsedURL, err := url.Parse(input)
	if err != nil || parsedURL.Scheme == "" {
		return nil, errors.New("invalid FhirUrl: " + input)
	}
	return &FhirUrl{Value: parsedURL, Element: element}, nil
}

// MarshalJSON serializes FhirUrl to JSON.
func (f *FhirUrl) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{}
	if f.Value != nil {
		data["value"] = f.Value.String()
	}
	if f.Element != nil {
		data["_value"] = f.Element
	}
	return json.Marshal(data)
}

// UnmarshalJSON deserializes JSON into FhirUrl.
func (f *FhirUrl) UnmarshalJSON(data []byte) error {
	temp := struct {
		Value   string   `json:"value"`
		Element *Element `json:"_value"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	parsedURL, err := url.Parse(temp.Value)
	if err != nil {
		return errors.New("invalid FhirUrl value")
	}

	f.Value = parsedURL
	f.Element = temp.Element
	return nil
}

// String returns the URL as a string.
func (f *FhirUrl) String() string {
	if f.Value != nil {
		return f.Value.String()
	}
	return ""
}

// Clone creates a deep copy of FhirUrl.
func (f *FhirUrl) Clone() *FhirUrl {
	var elementCopy *Element
	if f.Element != nil {
		elementCopy = f.Element.Copy()
	}
	return &FhirUrl{
		Value:   f.Value,
		Element: elementCopy,
	}
}

// Equal checks equality between two FhirUrl instances.
func (f *FhirUrl) Equal(other *FhirUrl) bool {
	return f.String() == other.String() && compareElements(f.Element, other.Element)
}
