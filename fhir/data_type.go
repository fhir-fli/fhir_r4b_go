package fhir_r4b_go

import (
	"encoding/json"
	"errors"
)

// DataType represents the base class for all reusable types in FHIR.
type DataType struct {
	Element
}

// NewDataType creates a new DataType instance.
func NewDataType(
	id *string,
	extensions []*FhirExtension,
	disallowExtension bool,
) *DataType {
	return &DataType{
		Element: *NewElement(id, extensions, disallowExtension),
	}
}

// FromJSON creates a DataType instance from JSON.
// This method throws an error as it's meant to be overridden.
func DataTypeFromJSON(input []byte) (*DataType, error) {
	return nil, errors.New("DataType.fromJson is not implemented")
}

// FromYAML creates a DataType instance from YAML input.
func DataTypeFromYAML(yamlData interface{}) (*DataType, error) {
	jsonData, err := convertYAMLToJSON(yamlData)
	if err != nil {
		return nil, err
	}
	return DataTypeFromJSON(jsonData)
}

// ToJSON converts the DataType to its JSON representation.
func (dt *DataType) ToJSON() ([]byte, error) {
	data := map[string]interface{}{}

	// Add ID if present
	if dt.ID != nil {
		data["id"] = *dt.ID
	}

	// Add Extension if present
	if len(dt.Extension) > 0 {
		exts := []interface{}{}
		for _, ext := range dt.Extension {
			// Call ToJSON and handle the error
			extJSON, err := ext.ToJSON()
			if err != nil {
				return nil, err // Return the error if ToJSON fails
			}

			// Unmarshal the JSON bytes into an interface
			var extData interface{}
			if err := json.Unmarshal(extJSON, &extData); err != nil {
				return nil, err // Handle unmarshalling error
			}

			exts = append(exts, extData)
		}
		data["extension"] = exts
	}

	// Marshal the map to JSON
	return json.Marshal(data)
}

// CopyWith creates a copy of the DataType with optional modifications.
func (dt *DataType) CopyWith(
	id *string,
	extensions []*FhirExtension,
) *DataType {
	newID := id
	if newID == nil {
		newID = dt.ID
	}

	newExtension := extensions
	if newExtension == nil {
		newExtension = dt.Extension
	}

	return NewDataType(newID, newExtension, dt.DisallowExtension)
}

// Helper: convertYAMLToJSON converts YAML data to JSON.
func convertYAMLToJSON(yamlData interface{}) ([]byte, error) {
	switch data := yamlData.(type) {
	case string:
		return json.Marshal(data)
	case map[string]interface{}:
		return json.Marshal(data)
	default:
		return nil, errors.New("invalid YAML input for DataType")
	}
}

// EqualsDeep performs a deep equality check between two DataType instances.
func (dt *DataType) EqualsDeep(other *DataType) bool {
	if other == nil {
		return false
	}

	// Compare the embedded Element fields
	if !dt.Element.EqualsDeep(&other.Element) {
		return false
	}

	return true
}