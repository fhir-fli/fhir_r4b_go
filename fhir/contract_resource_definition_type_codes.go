// Code generated by FHIR Generator. DO NOT EDIT.
// This value set contract specific codes for status.
package fhir_r4b_go

// ContractResourceDefinitionTypeCodes represents the FHIR ValueSet as an enumeration
type ContractResourceDefinitionTypeCodes int

const (
	// Temporary Value: To be completed
	ContractResourceDefinitionTypeCodes_Temp ContractResourceDefinitionTypeCodes = iota
)

// String converts the enum to its string representation
func (e ContractResourceDefinitionTypeCodes) String() string {
	switch e {
	case ContractResourceDefinitionTypeCodes_Temp: return "Temporary Value"
	default: return "Unknown"
	}
}
