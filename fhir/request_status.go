// Code generated by FHIR Generator. DO NOT EDIT.
// Codes identifying the lifecycle stage of a request.
package fhir_r4b_go

// RequestStatus represents the FHIR ValueSet as an enumeration
type RequestStatus int

const (
	// Draft: The request has been created but is not yet complete or ready for action.
	RequestStatus_Draft RequestStatus = iota
	// Active: The request is in force and ready to be acted upon.
	RequestStatus_Active RequestStatus = iota
	// On Hold: The request (and any implicit authorization to act) has been temporarily withdrawn but is expected to resume in the future.
	RequestStatus_On_hold RequestStatus = iota
	// Revoked: The request (and any implicit authorization to act) has been terminated prior to the known full completion of the intended actions. No further activity should occur.
	RequestStatus_Revoked RequestStatus = iota
	// Completed: The activity described by the request has been fully performed. No further activity will occur.
	RequestStatus_Completed RequestStatus = iota
	// Entered in Error: This request should never have existed and should be considered 'void'. (It is possible that real-world decisions were based on it. If real-world activity has occurred, the status should be "revoked" rather than "entered-in-error".).
	RequestStatus_Entered_in_error RequestStatus = iota
	// Unknown: The authoring/source system does not know which of the status values currently applies for this request. Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply, but the authoring/source system does not know which.
	RequestStatus_Unknown RequestStatus = iota
)

// String converts the enum to its string representation
func (e RequestStatus) String() string {
	switch e {
	case RequestStatus_Draft: return "Draft"
	case RequestStatus_Active: return "Active"
	case RequestStatus_On_hold: return "On Hold"
	case RequestStatus_Revoked: return "Revoked"
	case RequestStatus_Completed: return "Completed"
	case RequestStatus_Entered_in_error: return "Entered in Error"
	case RequestStatus_Unknown: return "Unknown"
	default: return "Unknown"
	}
}
