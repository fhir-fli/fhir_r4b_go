// Code generated by FHIR Generator. DO NOT EDIT.
// TODO.
package fhir_r4b_go

// AdverseEventCausalityMethod represents the FHIR ValueSet as an enumeration
type AdverseEventCausalityMethod int

const (
	// Probability Scale: 
	AdverseEventCausalityMethod_ProbabilityScale AdverseEventCausalityMethod = iota
	// Bayesian: 
	AdverseEventCausalityMethod_Bayesian AdverseEventCausalityMethod = iota
	// Checklist: 
	AdverseEventCausalityMethod_Checklist AdverseEventCausalityMethod = iota
)

// String converts the enum to its string representation
func (e AdverseEventCausalityMethod) String() string {
	switch e {
	case AdverseEventCausalityMethod_ProbabilityScale: return "Probability Scale"
	case AdverseEventCausalityMethod_Bayesian: return "Bayesian"
	case AdverseEventCausalityMethod_Checklist: return "Checklist"
	default: return "Unknown"
	}
}
