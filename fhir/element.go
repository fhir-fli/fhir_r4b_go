package fhir_r4b_go

import (
	"encoding/json"
	"reflect"
)

// Element represents the base definition for all FHIR elements.
type Element struct {
	FhirBase
	ID                *string          `json:"id,omitempty"`
	Extension         []*FhirExtension `json:"extension,omitempty"`
	DisallowExtension bool             `json:"disallowExtension,omitempty"`
}

// NewElement creates a new Element instance.
func NewElement(id *string, extensions []*FhirExtension, disallowExtension bool) *Element {
	return &Element{
		ID:                id,
		Extension:         extensions,
		DisallowExtension: disallowExtension,
	}
}

// FromJSON populates an Element from JSON data.
func (e *Element) FromJSON(data []byte) error {
	return json.Unmarshal(data, e)
}

// ToJSON converts an Element to its JSON representation.
func (e *Element) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}

// HasID checks if the element has an ID.
func (e *Element) HasID() bool {
	return e.ID != nil && *e.ID != ""
}

// HasExtension checks if the element has extensions.
func (e *Element) HasExtension() bool {
	return len(e.Extension) > 0
}

// GetExtensionByURL retrieves extensions by their URL.
func (e *Element) GetExtensionByURL(url string) []*FhirExtension {
	var result []*FhirExtension
	for _, ext := range e.Extension {
		if ext != nil && ext.Url.Value == url {
			result = append(result, ext)
		}
	}
	return result
}

// AddExtension adds an extension to the element.
func (e *Element) AddExtension(ext *FhirExtension) {
	e.Extension = append(e.Extension, ext)
}

// RemoveExtensionByURL removes extensions matching the given URL.
func (e *Element) RemoveExtensionByURL(url string) {
	var filtered []*FhirExtension
	for _, ext := range e.Extension {
		if ext != nil && ext.Url.Value != url {
			filtered = append(filtered, ext)
		}
	}
	e.Extension = filtered
}

// Copy creates a deep copy of the Element.
func (e *Element) Copy() *Element {
	return e.Clone()
}

// EqualsDeep performs a deep equality check with another Element.
func (e *Element) EqualsDeep(other *Element) bool {
	if other == nil {
		return false
	}
	if !reflect.DeepEqual(e.ID, other.ID) {
		return false
	}
	return compareFhirExtensions(e.Extension, other.Extension)
}

// Clone creates a deep copy of the Element.
func (e *Element) Clone() *Element {
	if e == nil {
		return nil
	}

	// Deep copy the Extensions slice
	extensionsCopy := make([]*FhirExtension, len(e.Extension))
	for i, ext := range e.Extension {
		if ext != nil {
			extensionsCopy[i] = ext.Clone()
		}
	}

	// Deep copy the ID pointer
	var idCopy *string
	if e.ID != nil {
		idVal := *e.ID
		idCopy = &idVal
	}

	return &Element{
		ID:                idCopy,
		Extension:         extensionsCopy,
		DisallowExtension: e.DisallowExtension,
	}
}

// compareFhirExtensions compares two slices of FhirExtension for deep equality.
func compareFhirExtensions(ext1, ext2 []*FhirExtension) bool {
	if len(ext1) != len(ext2) {
		return false
	}
	for i := range ext1 {
		if !compareOptionalExtensions(ext1[i], ext2[i]) {
			return false
		}
	}
	return true
}

// compareOptionalExtensions compares two FhirExtension pointers safely.
func compareOptionalExtensions(ext1, ext2 *FhirExtension) bool {
	if ext1 == nil || ext2 == nil {
		return ext1 == ext2 // Both must be nil to be equal
	}
	return ext1.EqualsDeep(ext2)
}
