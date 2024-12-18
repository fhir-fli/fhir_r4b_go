package fhir_r4b_go

import (
	"encoding/json"
	"errors"

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
		Value   string   `json:"value"`
		Element *Element `json:"_value"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	if temp.Value != "" {
		parsedUUID, err := uuid.Parse(temp.Value)
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
