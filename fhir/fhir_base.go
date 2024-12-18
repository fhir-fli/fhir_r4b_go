package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// Equalable is an interface for types that can compare themselves to others.
type Equalable interface {
	Equals(other Equalable) bool
}

// Cloneable is an interface for types that can clone themselves.
type Cloneable[T any] interface {
	Clone() *T
}

// FhirBase is the base type for all FHIR elements in Go.
type FhirBase struct {
	UserData           map[string]interface{} `json:"userData,omitempty"`
	FormatCommentsPre  []string               `json:"formatCommentsPre,omitempty"`
	FormatCommentsPost []string               `json:"formatCommentsPost,omitempty"`
	Annotations        []interface{}          `json:"annotations,omitempty"`
}

// NewFhirBase creates a new FhirBase instance.
func NewFhirBase() *FhirBase {
	return &FhirBase{
		UserData: make(map[string]interface{}),
	}
}

// IsEmpty checks if the object is empty.
func (fb *FhirBase) IsEmpty() bool {
	return len(fb.UserData) == 0 &&
		len(fb.FormatCommentsPre) == 0 &&
		len(fb.FormatCommentsPost) == 0 &&
		len(fb.Annotations) == 0
}

// HasUserData checks if user data exists for a given key.
func (fb *FhirBase) HasUserData(key string) bool {
	_, exists := fb.UserData[key]
	return exists
}

// GetUserData retrieves user data for a given key.
func (fb *FhirBase) GetUserData(key string) (interface{}, bool) {
	val, exists := fb.UserData[key]
	return val, exists
}

// SetUserData sets user data for a given key.
func (fb *FhirBase) SetUserData(key string, value interface{}) {
	fb.UserData[key] = value
}

// ClearUserData removes user data for a given key.
func (fb *FhirBase) ClearUserData(key string) {
	delete(fb.UserData, key)
}

// AddAnnotation adds an annotation.
func (fb *FhirBase) AddAnnotation(annotation interface{}) {
	fb.Annotations = append(fb.Annotations, annotation)
}

// RemoveAnnotations removes annotations of a specific type (by string key).
func (fb *FhirBase) RemoveAnnotations(targetType string) {
	var updatedAnnotations []interface{}
	for _, annotation := range fb.Annotations {
		if fmt.Sprintf("%T", annotation) != targetType {
			updatedAnnotations = append(updatedAnnotations, annotation)
		}
	}
	fb.Annotations = updatedAnnotations
}

// HasFormatComment checks if format comments exist.
func (fb *FhirBase) HasFormatComment() bool {
	return len(fb.FormatCommentsPre) > 0 || len(fb.FormatCommentsPost) > 0
}

// Clone creates a deep copy of FhirBase.
func (fb *FhirBase) Clone() *FhirBase {
	if fb == nil {
		return nil
	}

	clone := &FhirBase{
		UserData:           make(map[string]interface{}),
		FormatCommentsPre:  append([]string{}, fb.FormatCommentsPre...),
		FormatCommentsPost: append([]string{}, fb.FormatCommentsPost...),
		Annotations:        append([]interface{}{}, fb.Annotations...),
	}
	for k, v := range fb.UserData {
		clone.UserData[k] = v
	}
	return clone
}

// Equals compares equality between two FhirBase instances.
func (fb *FhirBase) Equals(other Equalable) bool {
	otherFhirBase, ok := other.(*FhirBase)
	if !ok {
		return false
	}

	if len(fb.UserData) != len(otherFhirBase.UserData) {
		return false
	}

	for k, v := range fb.UserData {
		if otherFhirBase.UserData[k] != v {
			return false
		}
	}

	return compareSlices(fb.FormatCommentsPre, otherFhirBase.FormatCommentsPre) &&
		compareSlices(fb.FormatCommentsPost, otherFhirBase.FormatCommentsPost) &&
		compareSlices(fb.Annotations, otherFhirBase.Annotations)
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

func cloneSlices[T Cloneable[T]](slice []T) []T {
	if slice == nil {
		return nil
	}
	clone := make([]T, len(slice))
	for i, item := range slice {
		if item != nil {
			clone[i] = item.Clone()
		} else {
			clone[i] = nil
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
