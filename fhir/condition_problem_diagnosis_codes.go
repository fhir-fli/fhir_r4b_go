// Code generated by FHIR Generator. DO NOT EDIT.
// Example value set for Condition/Problem/Diagnosis codes.
package fhir_r4b_go

// ConditionProblemDiagnosisCodes represents the FHIR ValueSet as an enumeration
type ConditionProblemDiagnosisCodes int

const (
	// No current problems or disability: 
	ConditionProblemDiagnosisCodes_Value160245001 ConditionProblemDiagnosisCodes = iota
)

// String converts the enum to its string representation
func (e ConditionProblemDiagnosisCodes) String() string {
	switch e {
	case ConditionProblemDiagnosisCodes_Value160245001: return "No current problems or disability"
	default: return "Unknown"
	}
}
