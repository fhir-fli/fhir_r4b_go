package fhir_r4b_go

import (
	"encoding/json"
	"errors"
)

// BackboneType represents a FHIR type that can carry modifier extensions.
type BackboneType struct {
	DataType
	ModifierExtension []*FhirExtension
}

// NewBackboneType creates a new BackboneType instance.
func NewBackboneType(
	id *string,
	extensions []*FhirExtension,
	modifierExtension []*FhirExtension,
	disallowExtension bool,
) *BackboneType {
	return &BackboneType{
		DataType:          *NewDataType(id, extensions, disallowExtension),
		ModifierExtension: modifierExtension,
	}
}

// FromJSON creates a BackboneType instance from JSON.
// This method throws an error as it's meant to be overridden.
func BackboneTypeFromJSON(input []byte) (*BackboneType, error) {
	return nil, errors.New("BackboneType.fromJson is not implemented")
}

// ToJSON converts the BackboneType to its JSON representation.
func (bt *BackboneType) ToJSON() ([]byte, error) {
	data := map[string]interface{}{}

	// Add ID
	if bt.ID != nil {
		data["id"] = *bt.ID
	}

	// Add Extension
	if len(bt.Extension) > 0 {
		exts := []interface{}{}
		for _, ext := range bt.Extension {
			extJSON, err := ext.ToJSON()
			if err != nil {
				return nil, err // Handle the error
			}

			var extData interface{}
			if err := json.Unmarshal(extJSON, &extData); err != nil {
				return nil, err // Handle unmarshalling error
			}

			exts = append(exts, extData)
		}
		data["extension"] = exts
	}

	// Add Modifier Extension
	if len(bt.ModifierExtension) > 0 {
		modExts := []interface{}{}
		for _, ext := range bt.ModifierExtension {
			modExtJSON, err := ext.ToJSON()
			if err != nil {
				return nil, err // Handle the error
			}

			var modExtData interface{}
			if err := json.Unmarshal(modExtJSON, &modExtData); err != nil {
				return nil, err // Handle unmarshalling error
			}

			modExts = append(modExts, modExtData)
		}
		data["modifierExtension"] = modExts
	}

	return json.Marshal(data)
}

// HasModifierExtension checks if there are any modifier extensions.
func (bt *BackboneType) HasModifierExtension() bool {
	return len(bt.ModifierExtension) > 0
}

// GetModifierExtensionFirstRep returns the first modifier extension.
func (bt *BackboneType) GetModifierExtensionFirstRep() *FhirExtension {
	if len(bt.ModifierExtension) == 0 {
		return &FhirExtension{
			Url: FhirString{Value: "fhirfli.dev"},
		}
	}
	return bt.ModifierExtension[0]
}

// GetModifierExtensionByURL retrieves all modifier extensions by a specific URL.
func (bt *BackboneType) GetModifierExtensionByURL(url string) []*FhirExtension {
	var result []*FhirExtension
	for _, ext := range bt.ModifierExtension {
		if ext.Url.Value == url {
			result = append(result, ext)
		}
	}
	return result
}

// AddModifierExtension adds a modifier extension to the BackboneType.
func (bt *BackboneType) AddModifierExtension(ext *FhirExtension) {
	bt.ModifierExtension = append(bt.ModifierExtension, ext)
}

// RemoveModifierExtensionByURL removes all modifier extensions matching a specific URL.
func (bt *BackboneType) RemoveModifierExtensionByURL(url string) {
	var filtered []*FhirExtension
	for _, ext := range bt.ModifierExtension {
		if ext.Url.Value != url {
			filtered = append(filtered, ext)
		}
	}
	bt.ModifierExtension = filtered
}

// EqualsDeep performs a deep equality check with another BackboneType.
func (bt *BackboneType) EqualsDeep(other *BackboneType) bool {
	if other == nil {
		return false
	}

	if !bt.DataType.EqualsDeep(&other.DataType) {
		return false
	}

	return CompareExtension(bt.ModifierExtension, other.ModifierExtension)
}

// IsEmpty checks whether the BackboneType is empty.
func (bt *BackboneType) IsEmpty() bool {
	return bt.DataType.IsEmpty() && len(bt.ModifierExtension) == 0
}

// CopyWith creates a copy of the BackboneType with optional modifications.
func (bt *BackboneType) CopyWith(
	id *string,
	extensions []*FhirExtension,
	modifierExtension []*FhirExtension,
) *BackboneType {
	newID := id
	if newID == nil {
		newID = bt.ID
	}

	newExtension := extensions
	if newExtension == nil {
		newExtension = bt.Extension
	}

	newModifierExtension := modifierExtension
	if newModifierExtension == nil {
		newModifierExtension = bt.ModifierExtension
	}

	return NewBackboneType(newID, newExtension, newModifierExtension, bt.DisallowExtension)
}
