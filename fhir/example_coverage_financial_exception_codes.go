// Code generated by FHIR Generator. DO NOT EDIT.
// This value set includes Example Coverage Financial Exception Codes.
package fhir_r4b_go

// ExampleCoverageFinancialExceptionCodes represents the FHIR ValueSet as an enumeration
type ExampleCoverageFinancialExceptionCodes int

const (
	// Retired: Retired persons have all copays and deductibles reduced.
	ExampleCoverageFinancialExceptionCodes_Retired ExampleCoverageFinancialExceptionCodes = iota
	// Foster child: Children in the foster care have all copays and deductibles waived.
	ExampleCoverageFinancialExceptionCodes_Foster ExampleCoverageFinancialExceptionCodes = iota
)

// String converts the enum to its string representation
func (e ExampleCoverageFinancialExceptionCodes) String() string {
	switch e {
	case ExampleCoverageFinancialExceptionCodes_Retired: return "Retired"
	case ExampleCoverageFinancialExceptionCodes_Foster: return "Foster child"
	default: return "Unknown"
	}
}
