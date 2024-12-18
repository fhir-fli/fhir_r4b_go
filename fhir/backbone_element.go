package fhir_r4b_go

import (
	"encoding/json"
)

// BackboneElement represents elements inside a resource but not those in a data type.
type BackboneElement struct {
	DataType
	ModifierExtension []*FhirExtension `json:"modifierExtension,omitempty"`
}

// NewBackboneElement creates a new BackboneElement instance.
func NewBackboneElement(
	id *string,
	extensions []*FhirExtension,
	modifierExtension []*FhirExtension,
	disallowExtension bool,
) *BackboneElement {
	return &BackboneElement{
		DataType:          *NewDataType(id, extensions, disallowExtension),
		ModifierExtension: modifierExtension,
	}
}

// BackboneElementFromJSON creates a BackboneElement instance from JSON.
func BackboneElementFromJSON(input []byte) (*BackboneElement, error) {
	var be BackboneElement
	if err := json.Unmarshal(input, &be); err != nil {
		return nil, err
	}
	return &be, nil
}

// ToJSON converts the BackboneElement to JSON.
func (be *BackboneElement) ToJSON() ([]byte, error) {
	return json.Marshal(be)
}

// HasModifierExtension checks if the BackboneElement has any modifier extensions.
func (be *BackboneElement) HasModifierExtension() bool {
	return len(be.ModifierExtension) > 0
}

// GetModifierExtensionFirstRep returns the first modifier extension or a default value.
func (be *BackboneElement) GetModifierExtensionFirstRep() *FhirExtension {
	if len(be.ModifierExtension) == 0 {
		return &FhirExtension{
			Url: FhirString{Value: "fhirfli.dev"},
		}
	}
	return be.ModifierExtension[0]
}

// GetModifierExtensionByURL retrieves all modifier extensions matching a specific URL.
func (be *BackboneElement) GetModifierExtensionByURL(url string) []*FhirExtension {
	var result []*FhirExtension
	for _, ext := range be.ModifierExtension {
		if ext.Url.Value == url {
			result = append(result, ext)
		}
	}
	return result
}

// AddModifierExtension adds a modifier extension to the BackboneElement.
func (be *BackboneElement) AddModifierExtension(ext *FhirExtension) {
	be.ModifierExtension = append(be.ModifierExtension, ext)
}

// RemoveModifierExtensionByURL removes all modifier extensions with a specific URL.
func (be *BackboneElement) RemoveModifierExtensionByURL(url string) {
	var filtered []*FhirExtension
	for _, ext := range be.ModifierExtension {
		if ext.Url.Value != url {
			filtered = append(filtered, ext)
		}
	}
	be.ModifierExtension = filtered
}

// EqualsDeep checks deep equality with another BackboneElement.
func (be *BackboneElement) EqualsDeep(other *BackboneElement) bool {
	if other == nil {
		return false
	}

	// Compare the base DataType fields
	if !be.DataType.EqualsDeep(&other.DataType) {
		return false
	}

	// Compare ModifierExtension using CompareExtension
	return CompareExtension(be.ModifierExtension, other.ModifierExtension)
}

// IsEmpty checks if the BackboneElement is empty.
func (be *BackboneElement) IsEmpty() bool {
	return be.DataType.IsEmpty() && len(be.ModifierExtension) == 0
}

// CopyWith creates a copy of BackboneElement with optional modifications.
func (be *BackboneElement) CopyWith(
	id *string,
	extensions []*FhirExtension,
	modifierExtension []*FhirExtension,
) *BackboneElement {
	// Use existing values if new ones aren't provided
	if id == nil {
		id = be.ID
	}
	if extensions == nil {
		extensions = be.Extension
	}
	if modifierExtension == nil {
		modifierExtension = be.ModifierExtension
	}

	return NewBackboneElement(id, extensions, modifierExtension, be.DisallowExtension)
}

// BackboneElementFromYAML converts YAML input to a BackboneElement instance.
func BackboneElementFromYAML(yamlData interface{}) (*BackboneElement, error) {
	jsonData, err := convertYAMLToJSON(yamlData)
	if err != nil {
		return nil, err
	}
	return BackboneElementFromJSON(jsonData)
}
