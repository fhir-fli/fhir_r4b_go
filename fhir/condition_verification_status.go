// Code generated by FHIR Generator. DO NOT EDIT.
// The verification status to support or decline the clinical status of the condition or diagnosis.
package fhir_r4b_go

// ConditionVerificationStatus represents the FHIR ValueSet as an enumeration
type ConditionVerificationStatus int

const (
	// Unconfirmed: There is not sufficient diagnostic and/or clinical evidence to treat this as a confirmed condition.
	ConditionVerificationStatus_Unconfirmed ConditionVerificationStatus = iota
	// Provisional: This is a tentative diagnosis - still a candidate that is under consideration.
	ConditionVerificationStatus_Provisional ConditionVerificationStatus = iota
	// Differential: One of a set of potential (and typically mutually exclusive) diagnoses asserted to further guide the diagnostic process and preliminary treatment.
	ConditionVerificationStatus_Differential ConditionVerificationStatus = iota
	// Confirmed: There is sufficient diagnostic and/or clinical evidence to treat this as a confirmed condition.
	ConditionVerificationStatus_Confirmed ConditionVerificationStatus = iota
	// Refuted: This condition has been ruled out by diagnostic and clinical evidence.
	ConditionVerificationStatus_Refuted ConditionVerificationStatus = iota
	// Entered in Error: The statement was entered in error and is not valid.
	ConditionVerificationStatus_Entered_in_error ConditionVerificationStatus = iota
)

// String converts the enum to its string representation
func (e ConditionVerificationStatus) String() string {
	switch e {
	case ConditionVerificationStatus_Unconfirmed: return "Unconfirmed"
	case ConditionVerificationStatus_Provisional: return "Provisional"
	case ConditionVerificationStatus_Differential: return "Differential"
	case ConditionVerificationStatus_Confirmed: return "Confirmed"
	case ConditionVerificationStatus_Refuted: return "Refuted"
	case ConditionVerificationStatus_Entered_in_error: return "Entered in Error"
	default: return "Unknown"
	}
}
