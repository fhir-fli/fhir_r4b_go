// Code generated by FHIR Generator. DO NOT EDIT.
// Lifecycle status of the questionnaire response.
package fhir_r4b_go

// QuestionnaireResponseStatus represents the FHIR ValueSet as an enumeration
type QuestionnaireResponseStatus int

const (
	// In Progress: This QuestionnaireResponse has been partially filled out with answers but changes or additions are still expected to be made to it.
	QuestionnaireResponseStatus_In_progress QuestionnaireResponseStatus = iota
	// Completed: This QuestionnaireResponse has been filled out with answers and the current content is regarded as definitive.
	QuestionnaireResponseStatus_Completed QuestionnaireResponseStatus = iota
	// Amended: This QuestionnaireResponse has been filled out with answers, then marked as complete, yet changes or additions have been made to it afterwards.
	QuestionnaireResponseStatus_Amended QuestionnaireResponseStatus = iota
	// Entered in Error: This QuestionnaireResponse was entered in error and voided.
	QuestionnaireResponseStatus_Entered_in_error QuestionnaireResponseStatus = iota
	// Stopped: This QuestionnaireResponse has been partially filled out with answers but has been abandoned. It is unknown whether changes or additions are expected to be made to it.
	QuestionnaireResponseStatus_Stopped QuestionnaireResponseStatus = iota
)

// String converts the enum to its string representation
func (e QuestionnaireResponseStatus) String() string {
	switch e {
	case QuestionnaireResponseStatus_In_progress: return "In Progress"
	case QuestionnaireResponseStatus_Completed: return "Completed"
	case QuestionnaireResponseStatus_Amended: return "Amended"
	case QuestionnaireResponseStatus_Entered_in_error: return "Entered in Error"
	case QuestionnaireResponseStatus_Stopped: return "Stopped"
	default: return "Unknown"
	}
}
