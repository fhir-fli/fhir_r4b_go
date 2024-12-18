package fhir_r4b_go

import (
	"errors"
	"fmt"
)

// FhirDecimal represents the FHIR 'decimal' type
type FhirDecimal struct {
	value   float64
	element *Element
}

// NewFhirDecimal creates a new validated FhirDecimal
func NewFhirDecimal(value float64, element *Element) *FhirDecimal {
	return &FhirDecimal{value: value, element: element}
}

// Value retrieves the decimal's value
func (fd *FhirDecimal) Value() float64 {
	return fd.value
}

// ToJSON serializes FhirDecimal to JSON
func (fd *FhirDecimal) ToJSON() (map[string]interface{}, error) {
	result := map[string]interface{}{"value": fd.value}
	if fd.element != nil {
		elementJSON, err := fd.element.ToJSON()
		if err != nil {
			return nil, err
		}
		result["_value"] = elementJSON
	}
	return result, nil
}

// Equals checks equality between two FhirNumbers
func (fd *FhirDecimal) Equals(other FhirNumber) bool {
	if other == nil {
		return false
	}
	return fd.value == other.Value()
}

// Arithmetic operations
func (fd *FhirDecimal) Add(other FhirNumber) FhirNumber {
	return NewFhirDecimal(fd.value+other.Value(), fd.element)
}

func (fd *FhirDecimal) Subtract(other FhirNumber) FhirNumber {
	return NewFhirDecimal(fd.value-other.Value(), fd.element)
}

func (fd *FhirDecimal) Multiply(other FhirNumber) FhirNumber {
	return NewFhirDecimal(fd.value*other.Value(), fd.element)
}

func (fd *FhirDecimal) Divide(other FhirNumber) (FhirNumber, error) {
	if other.Value() == 0 {
		return nil, errors.New("division by zero is not allowed")
	}
	return NewFhirDecimal(fd.value/other.Value(), fd.element), nil
}

// Clone creates a deep copy of FhirDecimal
func (fd *FhirDecimal) Clone() FhirNumber {
	return NewFhirDecimal(fd.value, fd.element.Clone())
}

// String representation of FhirDecimal
func (fd *FhirDecimal) String() string {
	return fmt.Sprintf("%f", fd.value)
}
