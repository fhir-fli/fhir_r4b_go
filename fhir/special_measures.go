// Code generated by FHIR Generator. DO NOT EDIT.
// Extra measures defined for a Medicinal Product, such as a requirement to conduct post-authorisation studies.
package fhir_r4b_go

// SpecialMeasures represents the FHIR ValueSet as an enumeration
type SpecialMeasures int

const (
	// Requirement to conduct post-authorisation studies: Requirement to conduct post-authorisation studies
	SpecialMeasures_Post_authorisationStudies SpecialMeasures = iota
)

// String converts the enum to its string representation
func (e SpecialMeasures) String() string {
	switch e {
	case SpecialMeasures_Post_authorisationStudies: return "Requirement to conduct post-authorisation studies"
	default: return "Unknown"
	}
}
