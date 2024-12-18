package fhir_r4b_go

import (
	"errors"
	"fmt"
)

// FhirUnsignedInt represents the FHIR 'integer' type
type FhirUnsignedInt struct {
	value   int64
	element *Element
}

// NewFhirUnsignedInt creates a new validated FhirUnsignedInt
func NewFhirUnsignedInt(value int64, element *Element) *FhirUnsignedInt {
	return &FhirUnsignedInt{value: value, element: element}
}

// Value retrieves the integer's value as float64 (to satisfy FhirNumber)
func (fi *FhirUnsignedInt) Value() float64 {
	return float64(fi.value)
}

// ToJSON serializes FhirUnsignedInt to JSON
func (fi *FhirUnsignedInt) ToJSON() (map[string]interface{}, error) {
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
func (fi *FhirUnsignedInt) Equals(other FhirNumber) bool {
	if other == nil {
		return false
	}
	return fi.Value() == other.Value()
}

// Arithmetic operations
func (fi *FhirUnsignedInt) Add(other FhirNumber) FhirNumber {
	return NewFhirUnsignedInt(int64(fi.value)+int64(other.Value()), fi.element)
}

func (fi *FhirUnsignedInt) Subtract(other FhirNumber) FhirNumber {
	return NewFhirUnsignedInt(int64(fi.value)-int64(other.Value()), fi.element)
}

func (fi *FhirUnsignedInt) Multiply(other FhirNumber) FhirNumber {
	return NewFhirUnsignedInt(int64(fi.value)*int64(other.Value()), fi.element)
}

func (fi *FhirUnsignedInt) Divide(other FhirNumber) (FhirNumber, error) {
	if other.Value() == 0 {
		return nil, errors.New("division by zero is not allowed")
	}
	return NewFhirUnsignedInt(int64(fi.value)/int64(other.Value()), fi.element), nil
}

// Clone creates a deep copy of FhirUnsignedInt
func (fi *FhirUnsignedInt) Clone() FhirNumber {
	clonedElement := fi.element.Clone() // Clone directly returns *Element
	return NewFhirUnsignedInt(fi.value, clonedElement)
}

// String representation of FhirUnsignedInt
func (fi *FhirUnsignedInt) String() string {
	return fmt.Sprintf("%d", fi.value)
}
