// Code generated by FHIR Generator. DO NOT EDIT.
// This value set includes Status codes.
package fhir_r4b_go

// FinancialResourceStatusCodes represents the FHIR ValueSet as an enumeration
type FinancialResourceStatusCodes int

const (
	// Active: The instance is currently in-force.
	FinancialResourceStatusCodes_Active FinancialResourceStatusCodes = iota
	// Cancelled: The instance is withdrawn, rescinded or reversed.
	FinancialResourceStatusCodes_Cancelled FinancialResourceStatusCodes = iota
	// Draft: A new instance the contents of which is not complete.
	FinancialResourceStatusCodes_Draft FinancialResourceStatusCodes = iota
	// Entered in Error: The instance was entered in error.
	FinancialResourceStatusCodes_Entered_in_error FinancialResourceStatusCodes = iota
)

// String converts the enum to its string representation
func (e FinancialResourceStatusCodes) String() string {
	switch e {
	case FinancialResourceStatusCodes_Active: return "Active"
	case FinancialResourceStatusCodes_Cancelled: return "Cancelled"
	case FinancialResourceStatusCodes_Draft: return "Draft"
	case FinancialResourceStatusCodes_Entered_in_error: return "Entered in Error"
	default: return "Unknown"
	}
}
