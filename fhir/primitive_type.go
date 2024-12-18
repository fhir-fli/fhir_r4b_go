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

// ToJSON serializes the PrimitiveType to JSON.
func (pt *PrimitiveType[T]) ToJSON() (map[string]interface{}, error) {
	jsonData := make(map[string]interface{})
	if pt.Value != nil {
		jsonData["value"] = pt.Value
	}
	if pt.Element != nil {
		elementJSON, err := pt.Element.ToJSON()
		if err != nil {
			return nil, err
		}
		jsonData["_value"] = elementJSON
	}
	return jsonData, nil
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
	return a.EqualsDeep(b)
}

// Helper: Parse an element from a list safely.
func parseElement(elements []interface{}, index int) *Element {
	if elements == nil || index < 0 || index >= len(elements) || elements[index] == nil {
		// Return nil if the list is nil, index is out of bounds, or the element is nil
		return nil
	}

	// Assert the element is a map[string]interface{}
	if elMap, ok := elements[index].(map[string]interface{}); ok {
		// Convert the map to JSON bytes
		jsonBytes, err := json.Marshal(elMap)
		if err != nil {
			// Return nil if marshalling fails
			return nil
		}

		// Parse JSON bytes into an Element
		element, err := FromJSON(jsonBytes)
		if err != nil {
			// Return nil if parsing fails
			return nil
		}

		return element
	}

	// Return nil if type assertion fails
	return nil
}
