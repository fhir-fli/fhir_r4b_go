// Code generated by FHIR Generator. DO NOT EDIT.
// Preferred value set for AllergyIntolerance Verification Status.
package fhir_r4b_go

// AllergyIntoleranceVerificationStatusCodes represents the FHIR ValueSet as an enumeration
type AllergyIntoleranceVerificationStatusCodes int

const (
	// Unconfirmed: A low level of certainty about the propensity for a reaction to the identified substance.
	AllergyIntoleranceVerificationStatusCodes_Unconfirmed AllergyIntoleranceVerificationStatusCodes = iota
	// Confirmed: A high level of certainty about the propensity for a reaction to the identified substance, which may include clinical evidence by testing or rechallenge.
	AllergyIntoleranceVerificationStatusCodes_Confirmed AllergyIntoleranceVerificationStatusCodes = iota
	// Refuted: A propensity for a reaction to the identified substance has been disputed or disproven with a sufficient level of clinical certainty to justify invalidating the assertion. This might or might not include testing or rechallenge.
	AllergyIntoleranceVerificationStatusCodes_Refuted AllergyIntoleranceVerificationStatusCodes = iota
	// Entered in Error: The statement was entered in error and is not valid.
	AllergyIntoleranceVerificationStatusCodes_Entered_in_error AllergyIntoleranceVerificationStatusCodes = iota
)

// String converts the enum to its string representation
func (e AllergyIntoleranceVerificationStatusCodes) String() string {
	switch e {
	case AllergyIntoleranceVerificationStatusCodes_Unconfirmed: return "Unconfirmed"
	case AllergyIntoleranceVerificationStatusCodes_Confirmed: return "Confirmed"
	case AllergyIntoleranceVerificationStatusCodes_Refuted: return "Refuted"
	case AllergyIntoleranceVerificationStatusCodes_Entered_in_error: return "Entered in Error"
	default: return "Unknown"
	}
}
