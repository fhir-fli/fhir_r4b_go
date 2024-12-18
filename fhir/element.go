package fhir_r4b_go

import (
	"encoding/json"
	"reflect"
)

// Element represents the base definition for all FHIR elements.
type Element struct {
	FhirBase
	ID                *string
	Extension         []*FhirExtension
	DisallowExtension bool
}

// NewElement creates a new Element instance.
func NewElement(id *string, extensions []*FhirExtension, disallowExtension bool) *Element {
	return &Element{
		ID:                id,
		Extension:         extensions,
		DisallowExtension: disallowExtension,
	}
}

// FromJSON creates an Element from JSON input.
func FromJSON(input []byte) (*Element, error) {
	var data map[string]interface{}
	if err := json.Unmarshal(input, &data); err != nil {
		return nil, err
	}

	// Extract ID
	var id *string
	if val, ok := data["id"].(string); ok {
		id = &val
	}

	// Extract Extension
	var extensions []*FhirExtension
	if extData, ok := data["extension"].([]interface{}); ok {
		for _, e := range extData {
			if extMap, ok := e.(map[string]interface{}); ok {
				// Marshal the map back into JSON
				extBytes, err := json.Marshal(extMap)
				if err != nil {
					return nil, err
				}

				// Unmarshal into a FhirExtension
				var extension FhirExtension
				if err := extension.FromJSON(extBytes); err != nil {
					return nil, err
				}

				extensions = append(extensions, &extension)
			}
		}
	}

	return NewElement(id, extensions, false), nil
}

// ToJSON converts Element to its JSON representation.
func (e *Element) ToJSON() ([]byte, error) {
	data := map[string]interface{}{}

	if e.ID != nil {
		data["id"] = *e.ID
	}

	if len(e.Extension) > 0 {
		exts := []interface{}{}
		for _, ext := range e.Extension {
			// Call ext.ToJSON and handle the error
			extJSON, err := ext.ToJSON()
			if err != nil {
				return nil, err // Return error if ToJSON fails
			}

			// Unmarshal extJSON to an interface{} for compatibility
			var extData interface{}
			if err := json.Unmarshal(extJSON, &extData); err != nil {
				return nil, err // Handle unmarshalling error
			}

			exts = append(exts, extData)
		}
		data["extension"] = exts
	}

	// Marshal the entire Element to JSON
	return json.Marshal(data)
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
		if ext.Url.Value == url {
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
		if ext.Url.Value != url {
			filtered = append(filtered, ext)
		}
	}
	e.Extension = filtered
}

// Copy creates a deep copy of the Element.
func (e *Element) Copy() *Element {
	extensionsCopy := []*FhirExtension{}
	for _, ext := range e.Extension {
		extensionsCopy = append(extensionsCopy, ext.Clone())
	}

	return &Element{
		ID:                e.ID,
		Extension:         extensionsCopy,
		DisallowExtension: e.DisallowExtension,
	}
}

// EqualsDeep performs a deep equality check.
func (e *Element) EqualsDeep(other *Element) bool {
	if other == nil {
		return false
	}

	if !reflect.DeepEqual(e.ID, other.ID) {
		return false
	}

	return CompareExtension(e.Extension, other.Extension)
}

// Helper: CompareExtension checks deep equality of extensions.
func CompareExtension(ext1, ext2 []*FhirExtension) bool {
	if len(ext1) != len(ext2) {
		return false
	}

	for i := range ext1 {
		if !ext1[i].EqualsDeep(ext2[i]) { // Now works correctly with updated EqualsDeep
			return false
		}
	}
	return true
}

// Clone creates a deep copy of the Element instance.
func (e *Element) Clone() *Element {
	if e == nil {
		return nil
	}

	// Deep copy the Extensions slice
	var extensionsCopy []*FhirExtension
	if e.Extension != nil {
		extensionsCopy = make([]*FhirExtension, len(e.Extension))
		for i, ext := range e.Extension {
			if ext != nil {
				extensionsCopy[i] = ext.Clone() // Use FhirExtension's Copy method
			}
		}
	}

	// Deep copy the ID pointer
	var idCopy *string
	if e.ID != nil {
		idVal := *e.ID
		idCopy = &idVal
	}

	// Return the cloned Element instance
	return &Element{
		ID:                idCopy,
		Extension:         extensionsCopy,
		DisallowExtension: e.DisallowExtension,
	}
}
