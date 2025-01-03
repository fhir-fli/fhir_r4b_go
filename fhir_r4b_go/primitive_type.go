package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"reflect"
)

// PrimitiveType represents the base for all FHIR primitive types.
type PrimitiveType[T any] struct {
	Value   *T       `json:"value,omitempty"`  // Primitive value
	Element *Element `json:"_value,omitempty"` // Optional metadata element
}

// NewPrimitiveType creates a new PrimitiveType with validation.
func NewPrimitiveType[T any](value *T, element *Element) (*PrimitiveType[T], error) {
	if value == nil && element == nil {
		return nil, errors.New("value or element is required")
	}
	return &PrimitiveType[T]{Value: value, Element: element}, nil
}

// MarshalJSON serializes the PrimitiveType to JSON.
func (pt *PrimitiveType[T]) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})
	if pt.Value != nil {
		jsonData["value"] = pt.Value
	}
	if pt.Element != nil {
		elementJSON, err := pt.Element.MarshalJSON()
		if err != nil {
			return nil, err
		}
		jsonData["_value"] = json.RawMessage(elementJSON)
	}
	return json.Marshal(jsonData)
}

// Equals checks equality with another PrimitiveType.
func (pt *PrimitiveType[T]) Equals(other *PrimitiveType[T]) bool {
	if pt == nil && other == nil {
		return true
	}
	if pt == nil || other == nil {
		return false
	}
	return equalValues(pt.Value, other.Value) && equalElements(pt.Element, other.Element)
}

// Clone creates a deep copy of the PrimitiveType.
func (pt *PrimitiveType[T]) Clone() *PrimitiveType[T] {
	if pt == nil {
		return nil
	}
	var clonedElement *Element
	if pt.Element != nil {
		clonedElement = pt.Element.Clone()
	}
	return &PrimitiveType[T]{
		Value:   cloneValue(pt.Value),
		Element: clonedElement,
	}
}

// Helper: Deep equality for values.
func equalValues[T any](a, b *T) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return reflect.DeepEqual(a, b)
}

// Helper: Deep copy for values.
func cloneValue[T any](value *T) *T {
	if value == nil {
		return nil
	}
	cloned := *value
	return &cloned
}

// Helper: Compare two Elements.
func equalElements(a, b *Element) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Equals(b)
}
