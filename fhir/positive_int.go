package fhir_r4b_go

import (
	"errors"
	"fmt"
)

// FhirPositiveInt represents the FHIR 'integer' type
type FhirPositiveInt struct {
	value   int64
	element *Element
}

// NewFhirPositiveInt creates a new validated FhirPositiveInt
func NewFhirPositiveInt(value int64, element *Element) *FhirPositiveInt {
	return &FhirPositiveInt{value: value, element: element}
}

// Value retrieves the integer's value as float64 (to satisfy FhirNumber)
func (fi *FhirPositiveInt) Value() float64 {
	return float64(fi.value)
}

// ToJSON serializes FhirPositiveInt to JSON
func (fi *FhirPositiveInt) ToJSON() (map[string]interface{}, error) {
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
func (fi *FhirPositiveInt) Equals(other FhirNumber) bool {
	if other == nil {
		return false
	}
	return fi.Value() == other.Value()
}

// Arithmetic operations
func (fi *FhirPositiveInt) Add(other FhirNumber) FhirNumber {
	return NewFhirPositiveInt(int64(fi.value)+int64(other.Value()), fi.element)
}

func (fi *FhirPositiveInt) Subtract(other FhirNumber) FhirNumber {
	return NewFhirPositiveInt(int64(fi.value)-int64(other.Value()), fi.element)
}

func (fi *FhirPositiveInt) Multiply(other FhirNumber) FhirNumber {
	return NewFhirPositiveInt(int64(fi.value)*int64(other.Value()), fi.element)
}

func (fi *FhirPositiveInt) Divide(other FhirNumber) (FhirNumber, error) {
	if other.Value() == 0 {
		return nil, errors.New("division by zero is not allowed")
	}
	return NewFhirPositiveInt(int64(fi.value)/int64(other.Value()), fi.element), nil
}

// Clone creates a deep copy of FhirPositiveInt
func (fi *FhirPositiveInt) Clone() FhirNumber {
	return NewFhirPositiveInt(fi.value, fi.element.Clone())
}

// String representation of FhirPositiveInt
func (fi *FhirPositiveInt) String() string {
	return fmt.Sprintf("%d", fi.value)
}
