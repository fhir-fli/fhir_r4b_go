// Code generated by FHIR Generator. DO NOT EDIT.
// This value set contract specific codes for decision modes.
package fhir_r4b_go

// ContractResourceDecisionModeCodes represents the FHIR ValueSet as an enumeration
type ContractResourceDecisionModeCodes int

const (
	// Policy: To be completed
	ContractResourceDecisionModeCodes_Policy ContractResourceDecisionModeCodes = iota
)

// String converts the enum to its string representation
func (e ContractResourceDecisionModeCodes) String() string {
	switch e {
	case ContractResourceDecisionModeCodes_Policy: return "Policy"
	default: return "Unknown"
	}
}
