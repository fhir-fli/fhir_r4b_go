package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
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

func (fb *FhirBase) EqualsDeep(other *FhirBase) bool {
	if other == nil {
		return false
	}

	if !reflect.DeepEqual(fb.UserData, other.UserData) {
		return false
	}

	// Handle slices with compareSlices
	if !compareSlices(fb.FormatCommentsPre, other.FormatCommentsPre) {
		return false
	}

	if !compareSlices(fb.FormatCommentsPost, other.FormatCommentsPost) {
		return false
	}

	if !compareSlices(fb.Annotations, other.Annotations) {
		return false
	}

	return true
}

// CompareDeep compares two FhirBase objects for equality.
func CompareDeep(f1, f2 *FhirBase, allowNull bool) bool {
	if allowNull {
		noLeft := f1 == nil || f1.IsEmpty()
		noRight := f2 == nil || f2.IsEmpty()
		if noLeft && noRight {
			return true
		}
	}
	if f1 == nil || f2 == nil {
		return false
	}
	return f1.EqualsDeep(f2)
}

// CompareDeepLists compares two lists of FhirBase objects deeply.
func CompareDeepLists(list1, list2 []*FhirBase, allowNull bool) bool {
	if allowNull {
		noLeft := len(list1) == 0
		noRight := len(list2) == 0
		if noLeft && noRight {
			return true
		}
	}
	if len(list1) != len(list2) {
		return false
	}
	for i := range list1 {
		if !CompareDeep(list1[i], list2[i], allowNull) {
			return false
		}
	}
	return true
}

// ToJSON converts FhirBase to a JSON representation.
func (fb *FhirBase) ToJSON() ([]byte, error) {
	return json.Marshal(fb)
}

// MarshalJSON provides custom JSON serialization.
func (fb *FhirBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		UserData           map[string]interface{} `json:"userData,omitempty"`
		FormatCommentsPre  []string               `json:"formatCommentsPre,omitempty"`
		FormatCommentsPost []string               `json:"formatCommentsPost,omitempty"`
		Annotations        []interface{}          `json:"annotations,omitempty"`
	}{
		UserData:           fb.UserData,
		FormatCommentsPre:  fb.FormatCommentsPre,
		FormatCommentsPost: fb.FormatCommentsPost,
		Annotations:        fb.Annotations,
	})
}

// UnmarshalJSON provides custom JSON deserialization.
func (fb *FhirBase) UnmarshalJSON(data []byte) error {
	temp := struct {
		UserData           map[string]interface{} `json:"userData,omitempty"`
		FormatCommentsPre  []string               `json:"formatCommentsPre,omitempty"`
		FormatCommentsPost []string               `json:"formatCommentsPost,omitempty"`
		Annotations        []interface{}          `json:"annotations,omitempty"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	fb.UserData = temp.UserData
	fb.FormatCommentsPre = temp.FormatCommentsPre
	fb.FormatCommentsPost = temp.FormatCommentsPost
	fb.Annotations = temp.Annotations
	return nil
}

// Clone creates a deep copy of FhirBase.
func (fb *FhirBase) Clone() *FhirBase {
	clone := *fb
	clone.UserData = make(map[string]interface{})
	for k, v := range fb.UserData {
		clone.UserData[k] = v
	}
	clone.FormatCommentsPre = append([]string{}, fb.FormatCommentsPre...)
	clone.FormatCommentsPost = append([]string{}, fb.FormatCommentsPost...)
	clone.Annotations = append([]interface{}{}, fb.Annotations...)
	return &clone
}

// Helper: compareMaps compares two maps.
func compareMaps(m1, m2 map[string]interface{}) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}

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
		if !reflect.DeepEqual(v1.Index(i).Interface(), v2.Index(i).Interface()) {
			return false
		}
	}

	return true
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
		// Convert simple YAML string
		return json.Marshal(v)
	case map[string]interface{}:
		// Convert complex YAML structure
		return json.Marshal(v)
	case []interface{}:
		// Handle YAML lists
		return json.Marshal(v)
	default:
		return nil, errors.New("invalid YAML input: unsupported type")
	}
}
