// Code generated by FHIR Generator. DO NOT EDIT.
// This value set includes example Procedure Type codes.
package fhir_r4b_go

// ExampleProcedureTypeCodes represents the FHIR ValueSet as an enumeration
type ExampleProcedureTypeCodes int

const (
	// Primary procedure: The first procedure in a series required to produce and overall patient outcome.
	ExampleProcedureTypeCodes_Primary ExampleProcedureTypeCodes = iota
	// Secondary procedure: The second procedure in a series required to produce and overall patient outcome.
	ExampleProcedureTypeCodes_Secondary ExampleProcedureTypeCodes = iota
)

// String converts the enum to its string representation
func (e ExampleProcedureTypeCodes) String() string {
	switch e {
	case ExampleProcedureTypeCodes_Primary: return "Primary procedure"
	case ExampleProcedureTypeCodes_Secondary: return "Secondary procedure"
	default: return "Unknown"
	}
}
