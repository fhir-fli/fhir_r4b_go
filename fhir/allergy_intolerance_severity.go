// Code generated by FHIR Generator. DO NOT EDIT.
// Clinical assessment of the severity of a reaction event as a whole, potentially considering multiple different manifestations.
package fhir_r4b_go

// AllergyIntoleranceSeverity represents the FHIR ValueSet as an enumeration
type AllergyIntoleranceSeverity int

const (
	// Mild: Causes mild physiological effects.
	AllergyIntoleranceSeverity_Mild AllergyIntoleranceSeverity = iota
	// Moderate: Causes moderate physiological effects.
	AllergyIntoleranceSeverity_Moderate AllergyIntoleranceSeverity = iota
	// Severe: Causes severe physiological effects.
	AllergyIntoleranceSeverity_Severe AllergyIntoleranceSeverity = iota
)

// String converts the enum to its string representation
func (e AllergyIntoleranceSeverity) String() string {
	switch e {
	case AllergyIntoleranceSeverity_Mild: return "Mild"
	case AllergyIntoleranceSeverity_Moderate: return "Moderate"
	case AllergyIntoleranceSeverity_Severe: return "Severe"
	default: return "Unknown"
	}
}
