// Code generated by FHIR Generator. DO NOT EDIT.
// Codes indicating the degree of authority/intentionality associated with a care plan.
package fhir_r4b_go

// CarePlanIntent represents the FHIR ValueSet as an enumeration
type CarePlanIntent int

const (
	// Proposal: The request is a suggestion made by someone/something that does not have an intention to ensure it occurs and without providing an authorization to act.
	CarePlanIntent_Proposal CarePlanIntent = iota
	// Plan: The request represents an intention to ensure something occurs without providing an authorization for others to act.
	CarePlanIntent_Plan CarePlanIntent = iota
	// Directive: The request represents a legally binding instruction authored by a Patient or RelatedPerson.
	CarePlanIntent_Directive CarePlanIntent = iota
	// Order: The request represents a request/demand and authorization for action by a Practitioner.
	CarePlanIntent_Order CarePlanIntent = iota
	// Original Order: The request represents an original authorization for action.
	CarePlanIntent_Original_order CarePlanIntent = iota
	// Reflex Order: The request represents an automatically generated supplemental authorization for action based on a parent authorization together with initial results of the action taken against that parent authorization.
	CarePlanIntent_Reflex_order CarePlanIntent = iota
	// Filler Order: The request represents the view of an authorization instantiated by a fulfilling system representing the details of the fulfiller's intention to act upon a submitted order.
	CarePlanIntent_Filler_order CarePlanIntent = iota
	// Instance Order: An order created in fulfillment of a broader order that represents the authorization for a single activity occurrence. E.g. The administration of a single dose of a drug.
	CarePlanIntent_Instance_order CarePlanIntent = iota
	// Option: The request represents a component or option for a RequestGroup that establishes timing, conditionality and/or other constraints among a set of requests. Refer to [[[RequestGroup]]] for additional information on how this status is used.
	CarePlanIntent_Option CarePlanIntent = iota
)

// String converts the enum to its string representation
func (e CarePlanIntent) String() string {
	switch e {
	case CarePlanIntent_Proposal: return "Proposal"
	case CarePlanIntent_Plan: return "Plan"
	case CarePlanIntent_Directive: return "Directive"
	case CarePlanIntent_Order: return "Order"
	case CarePlanIntent_Original_order: return "Original Order"
	case CarePlanIntent_Reflex_order: return "Reflex Order"
	case CarePlanIntent_Filler_order: return "Filler Order"
	case CarePlanIntent_Instance_order: return "Instance Order"
	case CarePlanIntent_Option: return "Option"
	default: return "Unknown"
	}
}
