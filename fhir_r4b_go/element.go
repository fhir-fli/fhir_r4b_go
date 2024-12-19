package fhir_r4b_go

import (
	"encoding/json"
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

// Clone creates a deep copy of the Element.
func (e *Element) Clone() *Element {
	if e == nil {
		return nil
	}

	// Deep copy extensions
	extensionsCopy := make([]*FhirExtension, len(e.Extension))
	for i, ext := range e.Extension {
		if ext != nil {
			extensionsCopy[i] = ext.Clone()
		}
	}

	// Deep copy ID
	var idCopy *string
	if e.ID != nil {
		idValue := *e.ID
		idCopy = &idValue
	}

	return &Element{
		ID:                idCopy,
		Extension:         extensionsCopy,
		DisallowExtension: e.DisallowExtension,
	}
}

// Equals performs a deep equality check with another Element.
func (e *Element) Equals(other *Element) bool {
	if e == nil && other == nil {
		return true
	}
	if e == nil || other == nil {
		return false
	}

	// Compare IDs
	if (e.ID == nil) != (other.ID == nil) || (e.ID != nil && *e.ID != *other.ID) {
		return false
	}

	// Compare extensions
	if len(e.Extension) != len(other.Extension) {
		return false
	}
	if !compareSlices(e.Extension, other.Extension) {
		return false
	}
	// Compare DisallowExtension
	return e.DisallowExtension == other.DisallowExtension
}

// IsEmpty checks whether the Element has no meaningful content.
func (e *Element) IsEmpty() bool {
	return e.ID == nil && len(e.Extension) == 0
}

// AddExtension appends an extension to the Element.
func (e *Element) AddExtension(ext *FhirExtension) {
	e.Extension = append(e.Extension, ext)
}

// GetExtensionByURL retrieves all extensions matching a specific URL.
func (e *Element) GetExtensionByURL(url string) []*FhirExtension {
	var matchingExtensions []*FhirExtension
	for _, ext := range e.Extension {
		if ext != nil && ext.Url.Value == &url {
			matchingExtensions = append(matchingExtensions, ext)
		}
	}
	return matchingExtensions
}

// RemoveExtensionByURL removes extensions that match a specific URL.
func (e *Element) RemoveExtensionByURL(url string) {
	var filteredExtensions []*FhirExtension
	for _, ext := range e.Extension {
		if ext != nil && ext.Url.Value != &url {
			filteredExtensions = append(filteredExtensions, ext)
		}
	}
	e.Extension = filteredExtensions
}

// String returns a string representation of the Element.
func (e *Element) String() string {
	data, err := e.ToJSON()
	if err != nil {
		return "<invalid Element>"
	}
	return string(data)
}
