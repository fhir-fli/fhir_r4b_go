package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

// FhirDecimal represents the FHIR 'decimal' type.
type FhirDecimal struct {
	Value   *float64 `json:"value,omitempty"`
	Element *Element `json:"_value,omitempty"`
}

// NewFhirDecimal creates a new validated FhirDecimal.
func NewFhirDecimal(input interface{}, element *Element) (*FhirDecimal, error) {
	val, err := parseDecimal(input)
	if err != nil {
		return nil, err
	}
	return &FhirDecimal{Value: val, Element: element}, nil
}

// parseDecimal parses input into a valid decimal value.
func parseDecimal(input interface{}) (*float64, error) {
	switch v := input.(type) {
	case float64:
		return &v, nil
	case string:
		parsed, err := strconv.ParseFloat(v, 64)
		if err != nil || !isValidDecimalString(v) {
			return nil, errors.New("invalid FHIR decimal format")
		}
		return &parsed, nil
	default:
		return nil, errors.New("unsupported type for FhirDecimal")
	}
}

// isValidDecimalString validates a decimal string format.
func isValidDecimalString(input string) bool {
	// FHIR decimal allows numeric values with optional leading '+' or '-' and a single '.'.
	return strings.Contains(input, ".")
}

// ToJSON serializes FhirDecimal to JSON.
func (fd *FhirDecimal) ToJSON() ([]byte, error) {
	data := map[string]interface{}{}
	if fd.Value != nil {
		data["value"] = strconv.FormatFloat(*fd.Value, 'f', -1, 64)
	}
	if fd.Element != nil {
		elementJSON, err := fd.Element.ToJSON()
		if err != nil {
			return nil, err
		}
		data["_value"] = elementJSON
	}
	return json.Marshal(data)
}

// FromJSON initializes a FhirDecimal from JSON input.
func (fd *FhirDecimal) FromJSON(data []byte) error {
	temp := struct {
		Value   string   `json:"value"`
		Element *Element `json:"_value"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	parsed, err := parseDecimal(temp.Value)
	if err != nil {
		return err
	}

	fd.Value = parsed
	fd.Element = temp.Element
	return nil
}

// Add performs addition with another FhirDecimal.
func (fd *FhirDecimal) Add(other *FhirDecimal) (*FhirDecimal, error) {
	if fd.Value == nil || other.Value == nil {
		return nil, errors.New("cannot perform addition: one or both values are nil")
	}
	result := *fd.Value + *other.Value
	return NewFhirDecimal(result, fd.Element.Clone())
}

// Subtract performs subtraction with another FhirDecimal.
func (fd *FhirDecimal) Subtract(other *FhirDecimal) (*FhirDecimal, error) {
	if fd.Value == nil || other.Value == nil {
		return nil, errors.New("cannot perform subtraction: one or both values are nil")
	}
	result := *fd.Value - *other.Value
	return NewFhirDecimal(result, fd.Element.Clone())
}

// Multiply performs multiplication with another FhirDecimal.
func (fd *FhirDecimal) Multiply(other *FhirDecimal) (*FhirDecimal, error) {
	if fd.Value == nil || other.Value == nil {
		return nil, errors.New("cannot perform multiplication: one or both values are nil")
	}
	result := *fd.Value * *other.Value
	return NewFhirDecimal(result, fd.Element.Clone())
}

// Divide performs division with another FhirDecimal.
func (fd *FhirDecimal) Divide(other *FhirDecimal) (*FhirDecimal, error) {
	if fd.Value == nil || other.Value == nil {
		return nil, errors.New("cannot perform division: one or both values are nil")
	}
	if *other.Value == 0 {
		return nil, errors.New("division by zero is not allowed")
	}
	result := *fd.Value / *other.Value
	return NewFhirDecimal(result, fd.Element.Clone())
}

// Clone creates a deep copy of FhirDecimal.
func (fd *FhirDecimal) Clone() *FhirDecimal {
	if fd == nil {
		return nil
	}
	return &FhirDecimal{
		Value:   floatPtrIfNotNil(fd.Value),
		Element: fd.Element.Clone(),
	}
}

// Equal checks for equality between two FhirDecimal instances.
func (fd *FhirDecimal) Equals(other *FhirDecimal) bool {
	if fd == nil && other == nil {
		return true
	}
	if fd == nil || other == nil {
		return false
	}
	return floatEquals(fd.Value, other.Value) && fd.Element.Equals(other.Element)
}

// String provides a string representation of the FhirDecimal.
func (fd *FhirDecimal) String() string {
	if fd.Value == nil {
		return ""
	}
	return strconv.FormatFloat(*fd.Value, 'f', -1, 64)
}

func floatEquals(a, b *float64) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}