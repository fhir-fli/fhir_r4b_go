package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// FhirUuid represents a validated UUID in the FHIR standard.
type FhirUuid struct {
	Value   *uuid.UUID `json:"value,omitempty"`
	Element *Element   `json:"_value,omitempty"`
}

// NewFhirUuid creates a new FhirUuid instance with validation.
func NewFhirUuid(input string, element *Element) (*FhirUuid, error) {
	parsedUUID, err := uuid.Parse(input)
	if err != nil {
		return nil, errors.New("invalid FhirUuid: " + input)
	}
	return &FhirUuid{Value: &parsedUUID, Element: element}, nil
}

// NewFhirUuidFromMap creates a FhirUuid instance from a map.
func NewFhirUuidFromMap(data map[string]interface{}) (*FhirUuid, error) {
	valueStr, ok := data["value"].(string)
	if !ok {
		return nil, errors.New("invalid or missing value for FhirUuid")
	}

	uuidValue, err := uuid.Parse(valueStr)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID value for FhirUuid: %v", err)
	}

	uuidInstance := &FhirUuid{Value: &uuidValue}

	if elementData, ok := data["_value"].(map[string]interface{}); ok {
		uuidInstance.Element = &Element{}
		if err := mapToStruct(elementData, uuidInstance.Element); err != nil {
			return nil, fmt.Errorf("failed to parse _value for FhirUuid: %v", err)
		}
	}

	return uuidInstance, nil
}

// GenerateFhirUuidV4 generates a new version 4 UUID.
func GenerateFhirUuidV4(element *Element) *FhirUuid {
	newUUID := uuid.New()
	return &FhirUuid{Value: &newUUID, Element: element}
}

// MarshalJSON serializes FhirUuid to JSON.
func (f *FhirUuid) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{}
	if f.Value != nil {
		data["value"] = f.Value.String()
	}
	if f.Element != nil {
		data["_value"] = f.Element
	}
	return json.Marshal(data)
}

// UnmarshalJSON deserializes JSON into FhirUuid.
func (f *FhirUuid) UnmarshalJSON(data []byte) error {
	temp := struct {
		Value   *string  `json:"value"`
		Element *Element `json:"_value"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	if *temp.Value != "" {
		parsedUUID, err := uuid.Parse(*temp.Value)
		if err != nil {
			return errors.New("invalid FhirUuid value")
		}
		f.Value = &parsedUUID
	}

	f.Element = temp.Element
	return nil
}

// Clone creates a deep copy of FhirUuid.
func (f *FhirUuid) Clone() *FhirUuid {
	if f == nil {
		return nil
	}
	var elementCopy *Element
	if f.Element != nil {
		elementCopy = f.Element.Clone()
	}
	var uuidCopy *uuid.UUID
	if f.Value != nil {
		uuidValue := *f.Value
		uuidCopy = &uuidValue
	}
	return &FhirUuid{
		Value:   uuidCopy,
		Element: elementCopy,
	}
}

// Equals checks equality between two FhirUuid instances.
func (f *FhirUuid) Equals(other *FhirUuid) bool {
	if f == nil && other == nil {
		return true
	}
	if f == nil || other == nil {
		return false
	}
	if (f.Value == nil && other.Value != nil) || (f.Value != nil && other.Value == nil) {
		return false
	}
	if f.Value != nil && other.Value != nil && *f.Value != *other.Value {
		return false
	}
	return f.Element.Equals(other.Element)
}
