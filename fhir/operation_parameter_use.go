// Code generated by FHIR Generator. DO NOT EDIT.
// Whether an operation parameter is an input or an output parameter.
package fhir_r4b_go

// OperationParameterUse represents the FHIR ValueSet as an enumeration
type OperationParameterUse int

const (
	// In: This is an input parameter.
	OperationParameterUse_In OperationParameterUse = iota
	// Out: This is an output parameter.
	OperationParameterUse_Out OperationParameterUse = iota
)

// String converts the enum to its string representation
func (e OperationParameterUse) String() string {
	switch e {
	case OperationParameterUse_In: return "In"
	case OperationParameterUse_Out: return "Out"
	default: return "Unknown"
	}
}
