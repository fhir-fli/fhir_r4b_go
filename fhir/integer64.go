package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"math/big"
)

// FhirInteger64 represents a 64-bit integer in FHIR.
type FhirInteger64 struct {
	Value   *big.Int `json:"value,omitempty"`
	Element *Element `json:"_value,omitempty"`
}

// NewFhirInteger64 creates a new FhirInteger64 with validation.
func NewFhirInteger64(input string, element *Element) (*FhirInteger64, error) {
	val, success := new(big.Int).SetString(input, 10)
	if !success {
		return nil, errors.New("invalid FhirInteger64: " + input)
	}
	return &FhirInteger64{Value: val, Element: element}, nil
}

// Clone creates a deep copy of FhirInteger64.
func (f *FhirInteger64) Clone() *FhirInteger64 {
	if f == nil {
		return nil
	}
	var valueCopy *big.Int
	if f.Value != nil {
		valueCopy = new(big.Int).Set(f.Value)
	}
	return &FhirInteger64{
		Value:   valueCopy,
		Element: f.Element.Clone(),
	}
}

// Equals checks equality with another FhirInteger64.
func (f *FhirInteger64) Equals(other *FhirInteger64) bool {
	if f == nil && other == nil {
		return true
	}
	if f == nil || other == nil {
		return false
	}
	return f.Value.Cmp(other.Value) == 0 && f.Element.Equals(other.Element)
}

// MarshalJSON serializes FhirInteger64 to JSON.
func (f *FhirInteger64) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{}
	if f.Value != nil {
		data["value"] = f.Value.String()
	}
	if f.Element != nil {
		data["_value"] = f.Element
	}
	return json.Marshal(data)
}

// UnmarshalJSON deserializes JSON into FhirInteger64.
func (f *FhirInteger64) UnmarshalJSON(data []byte) error {
	temp := struct {
		Value   string   `json:"value"`
		Element *Element `json:"_value"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	if temp.Value != "" {
		val, success := new(big.Int).SetString(temp.Value, 10)
		if !success {
			return errors.New("invalid FhirInteger64 value")
		}
		f.Value = val
	}
	f.Element = temp.Element
	return nil
}

// String returns the string representation of FhirInteger64.
func (f *FhirInteger64) String() string {
	if f.Value != nil {
		return f.Value.String()
	}
	return ""
}
