package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"regexp"
)

// FhirCode represents the FHIR primitive type `code`.
type FhirCode struct {
	Value   *string  `json:"value,omitempty"`   // The actual code value
	Element *Element `json:"_value,omitempty"` // Additional metadata (FHIR element)
}

// NewFhirCode creates a new FhirCode with validation.
func NewFhirCode(input string, element *Element) (*FhirCode, error) {
	if err := validateFhirCode(input); err != nil {
		return nil, err
	}
	return &FhirCode{
		Value:   &input,
		Element: element,
	}, nil
}

// FromJSON initializes a FhirCode from JSON input.
func (fc *FhirCode) FromJSON(data []byte) error {
	return json.Unmarshal(data, fc)
}

// ToJSON serializes the FhirCode to JSON.
func (fc *FhirCode) ToJSON() ([]byte, error) {
	return json.Marshal(fc)
}

// TryParseFhirCode attempts to parse input into a FhirCode, returning nil if invalid.
func TryParseFhirCode(input interface{}) (*FhirCode, error) {
	if str, ok := input.(string); ok {
		return NewFhirCode(str, nil)
	}
	return nil, errors.New("invalid input for FhirCode")
}

// validateFhirCode checks if the input string matches the FHIR `code` format.
func validateFhirCode(input string) error {
	regex := regexp.MustCompile(`^[^\s]+(\s[^\s]+)*$`)
	if regex.MatchString(input) {
		return nil
	}
	return errors.New("invalid FHIR Code: does not match required format")
}

// CopyWith creates a modified copy of the FhirCode with updated values.
func (fc *FhirCode) CopyWith(newValue *string, element *Element) *FhirCode {
	return &FhirCode{
		Value:   ifNotNil(newValue, fc.Value),
		Element: ifNotNil(element, fc.Element),
	}
}

// Clone creates a deep copy of the FhirCode.
func (fc *FhirCode) Clone() *FhirCode {
	if fc == nil {
		return nil
	}
	var newElement *Element
	if fc.Element != nil {
		newElement = fc.Element.Clone()
	}
	return &FhirCode{
		Value:   cloneString(fc.Value),
		Element: newElement,
	}
}

// Equals checks for equality between two FhirCode instances.
func (fc *FhirCode) Equals(other *FhirCode) bool {
	if fc == nil && other == nil {
		return true
	}
	if fc == nil || other == nil {
		return false
	}
	return equalStrings(fc.Value, other.Value) && equalElements(fc.Element, other.Element)
}

// String returns the FhirCode as a string.
func (fc *FhirCode) String() string {
	if fc.Value != nil {
		return *fc.Value
	}
	return ""
}

// FhirCodeList represents a list of FhirCode instances.
type FhirCodeList []*FhirCode

// FromJSONList initializes a list of FhirCode instances from JSON values and elements.
func FromJSONListFhirCode(values []interface{}, elements []interface{}) (FhirCodeList, error) {
	if elements != nil && len(values) != len(elements) {
		return nil, errors.New("values and elements must have the same length")
	}

	list := make(FhirCodeList, len(values))
	for i, v := range values {
		element := parseElement(elements, i)
		codeStr, ok := v.(string)
		if !ok {
			return nil, errors.New("invalid value for FhirCode")
		}
		code, err := NewFhirCode(codeStr, element)
		if err != nil {
			return nil, err
		}
		list[i] = code
	}
	return list, nil
}

// ToJSONList converts a FhirCodeList into JSON-compatible maps.
func (list FhirCodeList) ToJSONList() map[string]interface{} {
	values := make([]interface{}, len(list))
	elements := make([]interface{}, len(list))

	for i, fc := range list {
		if fc != nil {
			values[i] = fc.Value
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

// Helper: Deep clone a *string.
func cloneString(s *string) *string {
	if s == nil {
		return nil
	}
	cloned := *s
	return &cloned
}

// Helper: Compare two *string values.
func equalStrings(a, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

