// Code generated by FHIR Generator. DO NOT EDIT.
// Overall nature of the adverse event, e.g. real or potential.
package fhir_r4b_go

// AdverseEventActuality represents the FHIR ValueSet as an enumeration
type AdverseEventActuality int

const (
	// Adverse Event: The adverse event actually happened regardless of whether anyone was affected or harmed.
	AdverseEventActuality_Actual AdverseEventActuality = iota
	// Potential Adverse Event: A potential adverse event.
	AdverseEventActuality_Potential AdverseEventActuality = iota
)

// String converts the enum to its string representation
func (e AdverseEventActuality) String() string {
	switch e {
	case AdverseEventActuality_Actual: return "Adverse Event"
	case AdverseEventActuality_Potential: return "Potential Adverse Event"
	default: return "Unknown"
	}
}
