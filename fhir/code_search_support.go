// Code generated by FHIR Generator. DO NOT EDIT.
// The degree to which the server supports the code search parameter on ValueSet, if it is supported.
package fhir_r4b_go

// CodeSearchSupport represents the FHIR ValueSet as an enumeration
type CodeSearchSupport int

const (
	// Explicit Codes: The search for code on ValueSet only includes codes explicitly detailed on includes or expansions.
	CodeSearchSupport_Explicit CodeSearchSupport = iota
	// Implicit Codes: The search for code on ValueSet only includes all codes based on the expansion of the value set.
	CodeSearchSupport_All CodeSearchSupport = iota
)

// String converts the enum to its string representation
func (e CodeSearchSupport) String() string {
	switch e {
	case CodeSearchSupport_Explicit: return "Explicit Codes"
	case CodeSearchSupport_All: return "Implicit Codes"
	default: return "Unknown"
	}
}
