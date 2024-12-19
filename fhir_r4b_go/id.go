package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"regexp"
)

// FhirId represents the FHIR primitive type 'id'
type FhirId struct {
	Value   *string  `json:"value,omitempty"`  // ID value
	Element *Element `json:"_value,omitempty"` // Optional metadata element
}

// NewFhirId creates a new FhirId with validation
func NewFhirId(input string, element *Element) (*FhirId, error) {
	if err := validateFhirId(input); err != nil {
		return nil, err
	}
	return &FhirId{Value: &input, Element: element}, nil
}

// validateFhirId ensures input matches FHIR ID format
func validateFhirId(input string) error {
	regex := regexp.MustCompile(`^[A-Za-z0-9\-\.][A-Za-z0-9\-._]{0,63}$`)
	if !regex.MatchString(input) {
		return errors.New("invalid FhirId format")
	}
	return nil
}

// NewFhirIdFromMap creates a new FhirId instance from a map representation.
func NewFhirIdFromMap(data map[string]interface{}) (*FhirId, error) {
	// Extract and validate the "value" field
	rawValue, ok := data["value"].(string)
	if !ok || rawValue == "" {
		return nil, errors.New("missing or invalid 'value' in map for FhirId")
	}
	if err := validateFhirId(rawValue); err != nil {
		return nil, err
	}

	// Create a new FhirId instance with the value
	fhirId := &FhirId{Value: &rawValue}

	// Extract and parse the "_value" field for metadata if it exists
	if rawElement, ok := data["_value"].(map[string]interface{}); ok {
		element := &Element{}
		if err := mapToStruct(rawElement, element); err != nil {
			return nil, errors.New("failed to parse '_value' metadata for FhirId")
		}
		fhirId.Element = element
	}

	return fhirId, nil
}

func (fi *FhirId) UnmarshalJSON(data []byte) error {
	temp := struct {
		Value   *string  `json:"value"`
		Element *Element `json:"_value"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	if temp.Value != nil && validateFhirId(*temp.Value) != nil {
		return errors.New("invalid FhirId in JSON")
	}
	fi.Value = temp.Value
	fi.Element = temp.Element
	return nil
}

func (fi *FhirId) MarshalJSON() ([]byte, error) {
	result := map[string]interface{}{}
	if fi.Value != nil {
		result["value"] = *fi.Value
	}
	if fi.Element != nil {
		elementJSON, err := fi.Element.MarshalJSON()
		if err != nil {
			return nil, err
		}
		result["_value"] = elementJSON
	}
	return json.Marshal(result)
}

// Equals compares two FhirId instances
func (fi *FhirId) Equals(other *FhirId) bool {
	if fi == nil || other == nil {
		return false
	}
	if fi.Value == nil && other.Value == nil {
		return fi.Element.Equals(other.Element)
	}
	if fi.Value != nil && other.Value != nil {
		return *fi.Value == *other.Value && fi.Element.Equals(other.Element)
	}
	return false
}

func (fi *FhirId) Clone() *FhirId {
	if fi == nil {
		return nil
	}
	var clonedValue *string
	if fi.Value != nil {
		val := *fi.Value
		clonedValue = &val
	}
	return &FhirId{
		Value:   clonedValue,
		Element: fi.Element.Clone(),
	}
}
