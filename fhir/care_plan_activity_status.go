// Code generated by FHIR Generator. DO NOT EDIT.
// Codes that reflect the current state of a care plan activity within its overall life cycle.
package fhir_r4b_go

// CarePlanActivityStatus represents the FHIR ValueSet as an enumeration
type CarePlanActivityStatus int

const (
	// Not Started: Care plan activity is planned but no action has yet been taken.
	CarePlanActivityStatus_Not_started CarePlanActivityStatus = iota
	// Scheduled: Appointment or other booking has occurred but activity has not yet begun.
	CarePlanActivityStatus_Scheduled CarePlanActivityStatus = iota
	// In Progress: Care plan activity has been started but is not yet complete.
	CarePlanActivityStatus_In_progress CarePlanActivityStatus = iota
	// On Hold: Care plan activity was started but has temporarily ceased with an expectation of resumption at a future time.
	CarePlanActivityStatus_On_hold CarePlanActivityStatus = iota
	// Completed: Care plan activity has been completed (more or less) as planned.
	CarePlanActivityStatus_Completed CarePlanActivityStatus = iota
	// Cancelled: The planned care plan activity has been withdrawn.
	CarePlanActivityStatus_Cancelled CarePlanActivityStatus = iota
	// Stopped: The planned care plan activity has been ended prior to completion after the activity was started.
	CarePlanActivityStatus_Stopped CarePlanActivityStatus = iota
	// Unknown: The current state of the care plan activity is not known. Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply, but the authoring/source system does not know which one.
	CarePlanActivityStatus_Unknown CarePlanActivityStatus = iota
	// Entered in Error: Care plan activity was entered in error and voided.
	CarePlanActivityStatus_Entered_in_error CarePlanActivityStatus = iota
)

// String converts the enum to its string representation
func (e CarePlanActivityStatus) String() string {
	switch e {
	case CarePlanActivityStatus_Not_started: return "Not Started"
	case CarePlanActivityStatus_Scheduled: return "Scheduled"
	case CarePlanActivityStatus_In_progress: return "In Progress"
	case CarePlanActivityStatus_On_hold: return "On Hold"
	case CarePlanActivityStatus_Completed: return "Completed"
	case CarePlanActivityStatus_Cancelled: return "Cancelled"
	case CarePlanActivityStatus_Stopped: return "Stopped"
	case CarePlanActivityStatus_Unknown: return "Unknown"
	case CarePlanActivityStatus_Entered_in_error: return "Entered in Error"
	default: return "Unknown"
	}
}
