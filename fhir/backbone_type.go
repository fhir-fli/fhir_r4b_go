package fhir_r4b_go

import (
	"encoding/json"
)

// BackboneType represents a FHIR type that can carry modifier extensions.
type BackboneType struct {
	DataType
	ModifierExtension []*FhirExtension `json:"modifierExtension,omitempty"`
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

// BackboneTypeFromJSON creates a BackboneType instance from JSON.
func BackboneTypeFromJSON(input []byte) (*BackboneType, error) {
	var bt BackboneType
	if err := json.Unmarshal(input, &bt); err != nil {
		return nil, err
	}
	return &bt, nil
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
		exts, err := convertExtensionsToJSON(bt.Extension)
		if err != nil {
			return nil, err
		}
		data["extension"] = exts
	}

	// Add Modifier Extension
	if len(bt.ModifierExtension) > 0 {
		modExts, err := convertExtensionsToJSON(bt.ModifierExtension)
		if err != nil {
			return nil, err
		}
		data["modifierExtension"] = modExts
	}

	return json.Marshal(data)
}

// Helper: convertExtensionsToJSON handles conversion of a slice of FhirExtension to JSON-compatible objects.
func convertExtensionsToJSON(extensions []*FhirExtension) ([]interface{}, error) {
	exts := []interface{}{}
	for _, ext := range extensions {
		extJSON, err := ext.ToJSON()
		if err != nil {
			return nil, err
		}

		var extData interface{}
		if err := json.Unmarshal(extJSON, &extData); err != nil {
			return nil, err
		}

		exts = append(exts, extData)
	}
	return exts, nil
}

// HasModifierExtension checks if there are any modifier extensions.
func (bt *BackboneType) HasModifierExtension() bool {
	return len(bt.ModifierExtension) > 0
}

// GetModifierExtensionFirstRep returns the first modifier extension or a default value.
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

	// Compare base DataType fields
	if !bt.DataType.EqualsDeep(&other.DataType) {
		return false
	}

	// Compare ModifierExtension using CompareExtension
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
	// Use existing values if new ones aren't provided
	if id == nil {
		id = bt.ID
	}
	if extensions == nil {
		extensions = bt.Extension
	}
	if modifierExtension == nil {
		modifierExtension = bt.ModifierExtension
	}

	return NewBackboneType(id, extensions, modifierExtension, bt.DisallowExtension)
}

// BackboneTypeFromYAML converts YAML input to a BackboneType instance.
func BackboneTypeFromYAML(yamlData interface{}) (*BackboneType, error) {
	jsonData, err := convertYAMLToJSON(yamlData)
	if err != nil {
		return nil, err
	}
	return BackboneTypeFromJSON(jsonData)
}
