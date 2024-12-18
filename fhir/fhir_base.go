package fhir_r4b_go

import (
	"encoding/json"
	"reflect"
)

// FhirBase is the base type for all FHIR elements in Go.
type FhirBase struct {
	UserData          map[string]interface{}
	FormatCommentsPre []string
	FormatCommentsPost []string
	Annotations       []interface{}
}

// FhirType defines behavior shared by all FHIR types.
type FhirType interface {
	ToJSON() (map[string]interface{}, error)
	Clone() FhirType
	Equals(other FhirType) bool
	String() string
}

// NewFhirBase creates a new FhirBase instance.
func NewFhirBase() *FhirBase {
	return &FhirBase{
		UserData: make(map[string]interface{}),
	}
}

// FhirType returns the type of the FHIR element.
func (fb *FhirBase) FhirType() string {
	return "FhirBase"
}

// IsPrimitive checks if the element is primitive.
func (fb *FhirBase) IsPrimitive() bool {
	return false
}

// HasPrimitiveValue checks if a primitive value exists.
func (fb *FhirBase) HasPrimitiveValue() bool {
	return fb.IsPrimitive()
}

// PrimitiveValue retrieves the primitive value.
func (fb *FhirBase) PrimitiveValue() *string {
	return nil
}

// IsEmpty checks if the object is empty.
func (fb *FhirBase) IsEmpty() bool {
	return len(fb.UserData) == 0 &&
		len(fb.FormatCommentsPre) == 0 &&
		len(fb.FormatCommentsPost) == 0
}

// HasUserData checks if user data exists for a given key.
func (fb *FhirBase) HasUserData(key string) bool {
	_, exists := fb.UserData[key]
	return exists
}

// GetUserData retrieves user data for a given key.
func (fb *FhirBase) GetUserData(key string) interface{} {
	return fb.UserData[key]
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

// RemoveAnnotations removes annotations of a specific type.
func (fb *FhirBase) RemoveAnnotations(targetType reflect.Type) {
	var updatedAnnotations []interface{}
	for _, annotation := range fb.Annotations {
		if reflect.TypeOf(annotation) != targetType {
			updatedAnnotations = append(updatedAnnotations, annotation)
		}
	}
	fb.Annotations = updatedAnnotations
}

// GetUserString retrieves user data as a string.
func (fb *FhirBase) GetUserString(key string) string {
	if val, exists := fb.UserData[key]; exists {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

// GetUserInt retrieves user data as an integer.
func (fb *FhirBase) GetUserInt(key string) int {
	if val, exists := fb.UserData[key]; exists {
		if i, ok := val.(int); ok {
			return i
		}
	}
	return 0
}

// HasFormatComment checks if format comments exist.
func (fb *FhirBase) HasFormatComment() bool {
	return len(fb.FormatCommentsPre) > 0 || len(fb.FormatCommentsPost) > 0
}

// EqualsDeep checks deep equality with another FhirBase.
func (fb *FhirBase) EqualsDeep(other *FhirBase) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(fb, other)
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
func (fb *FhirBase) ToJSON() (string, error) {
	jsonData, err := json.Marshal(fb)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// Clone creates a deep copy of FhirBase.
func (fb *FhirBase) Clone() *FhirBase {
	clone := *fb
	clone.UserData = make(map[string]interface{})
	for k, v := range fb.UserData {
		clone.UserData[k] = v
	}
	return &clone
}
