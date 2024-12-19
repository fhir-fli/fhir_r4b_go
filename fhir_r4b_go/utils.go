package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"reflect"
	"time"
)

// Equalable is an interface for types that can compare themselves to others.
type Equalable interface {
	Equals(other Equalable) bool
}

// Helper functions for creating pointers.
func intPtr(v int) *int { return &v }
func intPtrIfNotNil(v *int) *int {
	if v == nil {
		return nil
	}
	return intPtr(*v)
}

func int64Ptr(v int64) *int64 { return &v }
func int64PtrIfNotNil(v *int64) *int64 {
	if v == nil {
		return nil
	}
	return int64Ptr(*v)
}

func strPtrIfNotNil(v *string) *string {
	if v == nil {
		return nil
	}
	return strPtr(*v)
}

func strPtr(v string) *string { return &v }
func timePtr(t time.Time) *time.Time {
	return &t
}

// Helper: compareSlices compares slices of Equalable or primitive types.
func compareSlices(slice1, slice2 interface{}) bool {
	v1 := reflect.ValueOf(slice1)
	v2 := reflect.ValueOf(slice2)

	if v1.Kind() != reflect.Slice || v2.Kind() != reflect.Slice {
		return false
	}

	if v1.Len() != v2.Len() {
		return false
	}

	for i := 0; i < v1.Len(); i++ {
		item1 := v1.Index(i).Interface()
		item2 := v2.Index(i).Interface()

		// Check if the items implement Equalable
		equalable1, ok1 := item1.(Equalable)
		equalable2, ok2 := item2.(Equalable)

		if ok1 && ok2 {
			// Use the Equals method for comparison
			if !equalable1.Equals(equalable2) {
				return false
			}
		} else {
			// Fall back to reflect.DeepEqual for non-Equalable items
			if !reflect.DeepEqual(item1, item2) {
				return false
			}
		}
	}

	return true
}

func cloneSlices[T any, PT interface {
	*T
	Clone() PT
}](slice []PT) []PT {
	if slice == nil {
		return nil
	}
	clone := make([]PT, len(slice))
	for i, item := range slice {
		if item != nil { // Check if the item is non-nil
			clone[i] = item.Clone() // Clone the non-nil item
		} else {
			clone[i] = nil // Preserve nil for nil items
		}
	}
	return clone
}

// Helper for safely returning the non-nil value.
func ifNotNil[T any](newValue, oldValue *T) *T {
	if newValue != nil {
		return newValue
	}
	return oldValue
}

// Helper: convertYAMLToJSON converts YAML data to JSON format.
func convertYAMLToJSON(yamlData interface{}) ([]byte, error) {
	switch v := yamlData.(type) {
	case string:
		return json.Marshal(v)
	case map[string]interface{}:
		return json.Marshal(v)
	case []interface{}:
		return json.Marshal(v)
	default:
		return nil, errors.New("invalid YAML input: unsupported type")
	}
}

// Helper functions for pointer management and comparison.
func floatPtrIfNotNil(f *float64) *float64 {
	if f == nil {
		return nil
	}
	return f
}

// Helper function to map a map[string]interface{} to a struct.
func mapToStruct(data map[string]interface{}, target interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, target)
}

// toMapOrNil safely converts JSON bytes to a map[string]interface{}.
// Returns nil if an error occurs during the conversion.
func toMapOrNil(data []byte, err error) map[string]interface{} {
	if err != nil {
		return nil
	}
	var result map[string]interface{}
	if jsonErr := json.Unmarshal(data, &result); jsonErr != nil {
		return nil
	}
	return result
}
