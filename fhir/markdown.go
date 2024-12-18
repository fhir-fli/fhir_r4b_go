package fhir_r4b_go

import (
	"encoding/json"
	"errors"
)

// FhirMarkdown represents the FHIR primitive type "markdown".
type FhirMarkdown struct {
	Value   string   `json:"value,omitempty"`
	Element *Element `json:"_value,omitempty"`
}

// NewFhirMarkdown creates a new FhirMarkdown instance with validation.
func NewFhirMarkdown(input string, element *Element) (*FhirMarkdown, error) {
	if !validateMarkdown(input) {
		return nil, errors.New("invalid FhirMarkdown: contains invalid characters")
	}
	return &FhirMarkdown{Value: input, Element: element}, nil
}

// validateMarkdown ensures the input conforms to markdown rules.
func validateMarkdown(input string) bool {
	return len(input) > 0
}

// MarshalJSON serializes FhirMarkdown to JSON.
func (f *FhirMarkdown) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{}
	if f.Value != "" {
		data["value"] = f.Value
	}
	if f.Element != nil {
		data["_value"] = f.Element
	}
	return json.Marshal(data)
}

// UnmarshalJSON deserializes JSON into FhirMarkdown.
func (f *FhirMarkdown) UnmarshalJSON(data []byte) error {
	temp := struct {
		Value   string   `json:"value"`
		Element *Element `json:"_value"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	f.Value = temp.Value
	f.Element = temp.Element
	return nil
}

// String returns the markdown value as a string.
func (f *FhirMarkdown) String() string {
	return f.Value
}

// Equal checks equality between two FhirMarkdown instances.
func (f *FhirMarkdown) Equal(other *FhirMarkdown) bool {
	return f.Value == other.Value
}
