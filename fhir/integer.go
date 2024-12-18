package fhir_r4b_go

import (
	"errors"
	"fmt"
)

// FhirInteger represents the FHIR 'integer' type
type FhirInteger struct {
	value   int64
	element *Element
}

// NewFhirInteger creates a new validated FhirInteger
func NewFhirInteger(value int64, element *Element) *FhirInteger {
	return &FhirInteger{value: value, element: element}
}

// Value retrieves the integer's value as float64 (to satisfy FhirNumber)
func (fi *FhirInteger) Value() float64 {
	return float64(fi.value)
}

// ToJSON serializes FhirInteger to JSON
func (fi *FhirInteger) ToJSON() (map[string]interface{}, error) {
	result := map[string]interface{}{"value": fi.value}
	if fi.element != nil {
		elementJSON, err := fi.element.ToJSON()
		if err != nil {
			return nil, err
		}
		result["_value"] = elementJSON
	}
	return result, nil
}

// Equals checks equality between two FhirNumbers
func (fi *FhirInteger) Equals(other FhirNumber) bool {
	if other == nil {
		return false
	}
	return fi.Value() == other.Value()
}

// Arithmetic operations
func (fi *FhirInteger) Add(other FhirNumber) FhirNumber {
	return NewFhirInteger(int64(fi.value)+int64(other.Value()), fi.element)
}

func (fi *FhirInteger) Subtract(other FhirNumber) FhirNumber {
	return NewFhirInteger(int64(fi.value)-int64(other.Value()), fi.element)
}

func (fi *FhirInteger) Multiply(other FhirNumber) FhirNumber {
	return NewFhirInteger(int64(fi.value)*int64(other.Value()), fi.element)
}

func (fi *FhirInteger) Divide(other FhirNumber) (FhirNumber, error) {
	if other.Value() == 0 {
		return nil, errors.New("division by zero is not allowed")
	}
	return NewFhirInteger(int64(fi.value)/int64(other.Value()), fi.element), nil
}

// Clone creates a deep copy of FhirInteger
func (fi *FhirInteger) Clone() FhirNumber {
	return NewFhirInteger(fi.value, fi.element.Clone())
}

// String representation of FhirInteger
func (fi *FhirInteger) String() string {
	return fmt.Sprintf("%d", fi.value)
}
