package fhir_r4b_go

import (
	"encoding/json"
	"errors"
)

// FhirBoolean represents the FHIR primitive type `boolean`.
type FhirBoolean struct {
	Value   *bool    `json:"value,omitempty"`  // The boolean value
	Element *Element `json:"_value,omitempty"` // Additional metadata (FHIR element)
}

// NewFhirBoolean creates a new FhirBoolean instance with validation.
func NewFhirBoolean(value bool, element *Element) *FhirBoolean {
	return &FhirBoolean{
		Value:   &value,
		Element: element,
	}
}

// FromJSON initializes a FhirBoolean from JSON input.
func (fb *FhirBoolean) FromJSON(data []byte) error {
	return json.Unmarshal(data, fb)
}

// ToJSON converts the FhirBoolean to JSON.
func (fb *FhirBoolean) ToJSON() ([]byte, error) {
	return json.Marshal(fb)
}

// TryParse attempts to create a FhirBoolean from dynamic input.
func TryParseBoolean(input interface{}) (*FhirBoolean, error) {
	switch v := input.(type) {
	case bool:
		return NewFhirBoolean(v, nil), nil
	case string:
		if v == "true" {
			return NewFhirBoolean(true, nil), nil
		} else if v == "false" {
			return NewFhirBoolean(false, nil), nil
		}
	}
	return nil, errors.New("invalid input for FhirBoolean")
}

// CopyWith creates a modified copy of FhirBoolean.
func (fb *FhirBoolean) CopyWith(newValue *bool, element *Element) *FhirBoolean {
	return &FhirBoolean{
		Value:   ifNotNil(newValue, fb.Value),
		Element: ifNotNil(element, fb.Element),
	}
}

// Clone creates a deep copy of the FhirBoolean instance.
func (fb *FhirBoolean) Clone() *FhirBoolean {
	if fb == nil {
		return nil
	}

	var newElement *Element
	if fb.Element != nil {
		// Manually deep copy the Element
		var copiedExtensions []*FhirExtension
		if fb.Element.Extension != nil {
			copiedExtensions = make([]*FhirExtension, len(fb.Element.Extension))
			for i, ext := range fb.Element.Extension {
				if ext != nil {
					// Directly call Clone() on ext as it is already *FhirExtension
					copiedExtensions[i] = ext.Clone()
				}
			}
		}

		newElement = &Element{
			ID:        fb.Element.ID,
			Extension: copiedExtensions,
		}
	}

	return &FhirBoolean{
		Value:   fb.Value,
		Element: newElement,
	}
}

// Equals checks for equality between FhirBoolean instances.
func (fb *FhirBoolean) Equals(other *FhirBoolean) bool {
	if fb == nil && other == nil {
		return true
	}
	if fb == nil || other == nil {
		return false
	}
	return equalBool(fb.Value, other.Value) && equalElements(fb.Element, other.Element)
}

// Helper: Check if two *bool values are equal.
func equalBool(a, b *bool) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// String returns the string representation of the FhirBoolean.
func (fb *FhirBoolean) String() string {
	if fb.Value == nil {
		return "null"
	}
	if *fb.Value {
		return "true"
	}
	return "false"
}

// FhirBooleanList represents a list of FhirBoolean values.
type FhirBooleanList []*FhirBoolean

// FromJSONList converts a list of JSON values into a list of FhirBoolean.
func FromJSONListFhirBoolean(values []interface{}, elements []interface{}) (FhirBooleanList, error) {
	if elements != nil && len(values) != len(elements) {
		return nil, errors.New("values and elements must have the same length")
	}

	list := make(FhirBooleanList, len(values))
	for i, v := range values {
		element := parseElement(elements, i)
		parsed, err := TryParseBoolean(v)
		if err != nil {
			return nil, err
		}
		list[i] = parsed.CopyWith(nil, element)
	}
	return list, nil
}

// ToJSONList converts a list of FhirBoolean into JSON-compatible maps.
func (list FhirBooleanList) ToJSONList() map[string]interface{} {
	values := make([]interface{}, len(list))
	elements := make([]interface{}, len(list))

	for i, fb := range list {
		if fb != nil {
			values[i] = fb.Value
			elements[i] = fb.Element
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
