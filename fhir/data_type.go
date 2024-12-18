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
// Returns an error because this method should be overridden.
func DataTypeFromJSON(input []byte) (*DataType, error) {
	return nil, errors.New("DataType.FromJSON is not implemented; override this method for specific types")
}

// FromYAML creates a DataType instance from YAML input by converting it to JSON first.
func DataTypeFromYAML(yamlData interface{}) (*DataType, error) {
	jsonData, err := convertYAMLToJSON(yamlData)
	if err != nil {
		return nil, err
	}
	return DataTypeFromJSON(jsonData)
}

// ToJSON converts the DataType to its JSON representation.
func (dt *DataType) ToJSON() ([]byte, error) {
	// Use Go's JSON marshalling unless customization is required
	return json.Marshal(dt)
}

// CopyWith creates a copy of the DataType with optional modifications.
func (dt *DataType) CopyWith(
	id *string,
	extensions []*FhirExtension,
) *DataType {
	// Use existing values if new ones aren't provided
	if id == nil {
		id = dt.ID
	}
	if extensions == nil {
		extensions = dt.Extension
	}

	return NewDataType(id, extensions, dt.DisallowExtension)
}

// EqualsDeep performs a deep equality check between two DataType instances.
func (dt *DataType) EqualsDeep(other *DataType) bool {
	if other == nil {
		return false
	}
	// Use Element's EqualsDeep to compare common fields
	return dt.Element.EqualsDeep(&other.Element)
}

