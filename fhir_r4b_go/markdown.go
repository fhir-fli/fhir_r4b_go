// FhirMarkdown represents the FHIR primitive type "markdown".
package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
)

type FhirMarkdown struct {
	Value   *string  `json:"value,omitempty"`
	Element *Element `json:"_value,omitempty"`
}

// NewFhirMarkdown creates a new FhirMarkdown instance with validation.
func NewFhirMarkdown(input string, element *Element) (*FhirMarkdown, error) {
	if !validateMarkdown(input) {
		return nil, errors.New("invalid FhirMarkdown: contains invalid characters")
	}
	return &FhirMarkdown{
		Value:   &input,
		Element: element,
	}, nil
}

// NewFhirMarkdownFromMap creates a FhirMarkdown instance from a map.
func NewFhirMarkdownFromMap(data map[string]interface{}) (*FhirMarkdown, error) {
	value, ok := data["value"].(string)
	if !ok {
		return nil, errors.New("invalid or missing value for FhirMarkdown")
	}

	markdown := &FhirMarkdown{Value: &value}

	if elementData, ok := data["_value"].(map[string]interface{}); ok {
		markdown.Element = &Element{}
		if err := mapToStruct(elementData, markdown.Element); err != nil {
			return nil, fmt.Errorf("failed to parse _value for FhirMarkdown: %v", err)
		}
	}

	return markdown, nil
}

// validateMarkdown ensures the input conforms to markdown rules.
func validateMarkdown(input string) bool {
	regex := regexp.MustCompile(`^[^\x00-\x08\x0B\x0C\x0E-\x1F]*$`)
	return regex.MatchString(input)
}

// MarshalJSON serializes FhirMarkdown to JSON.
func (f *FhirMarkdown) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{}
	if *f.Value != "" {
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
	if !validateMarkdown(temp.Value) {
		return errors.New("invalid FhirMarkdown: contains invalid characters")
	}
	f.Value = &temp.Value
	f.Element = temp.Element
	return nil
}

// String returns the markdown value as a string.
func (f *FhirMarkdown) String() string {
	return *f.Value
}

// Clone creates a deep copy of FhirMarkdown.
func (f *FhirMarkdown) Clone() *FhirMarkdown {
	if f == nil {
		return nil
	}
	return &FhirMarkdown{
		Value:   f.Value,
		Element: f.Element.Clone(),
	}
}

// Equals checks equality between two FhirMarkdown instances.
func (f *FhirMarkdown) Equals(other *FhirMarkdown) bool {
	if f == nil && other == nil {
		return true
	}
	if f == nil || other == nil {
		return false
	}
	return f.Value == other.Value && f.Element.Equals(other.Element)
}
