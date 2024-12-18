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
func NewFhirInteger64(input interface{}, element *Element) (*FhirInteger64, error) {
	val, err := parseBigInt(input)
	if err != nil {
		return nil, err
	}
	return &FhirInteger64{Value: val, Element: element}, nil
}

// parseBigInt validates and converts input into *big.Int.
func parseBigInt(input interface{}) (*big.Int, error) {
	switch v := input.(type) {
	case int, int64:
		return big.NewInt(int64(v.(int))), nil
	case string:
		bi := new(big.Int)
		_, success := bi.SetString(v, 10)
		if !success {
			return nil, errors.New("invalid FhirInteger64 string format")
		}
		return bi, nil
	case *big.Int:
		return v, nil
	default:
		return nil, errors.New("unsupported type for FhirInteger64")
	}
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
		bi := new(big.Int)
		_, success := bi.SetString(temp.Value, 10)
		if !success {
			return errors.New("invalid FhirInteger64 value")
		}
		f.Value = bi
	}
	f.Element = temp.Element
	return nil
}

// String returns the string representation of the value.
func (f *FhirInteger64) String() string {
	if f.Value != nil {
		return f.Value.String()
	}
	return ""
}

// Arithmetic and Comparison Methods

// Add adds another FhirInteger64 or integer.
func (f *FhirInteger64) Add(other *FhirInteger64) *FhirInteger64 {
	result := new(big.Int).Add(f.Value, other.Value)
	return &FhirInteger64{Value: result}
}

// Sub subtracts another FhirInteger64 or integer.
func (f *FhirInteger64) Sub(other *FhirInteger64) *FhirInteger64 {
	result := new(big.Int).Sub(f.Value, other.Value)
	return &FhirInteger64{Value: result}
}

// Equal checks equality with another FhirInteger64.
func (f *FhirInteger64) Equal(other *FhirInteger64) bool {
	return f.Value.Cmp(other.Value) == 0
}
