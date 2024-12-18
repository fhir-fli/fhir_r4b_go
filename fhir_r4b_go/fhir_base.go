package fhir_r4b_go

import (
	"fmt"
)

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
