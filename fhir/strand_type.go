// Code generated by FHIR Generator. DO NOT EDIT.
// Type for strand.
package fhir_r4b_go

// StrandType represents the FHIR ValueSet as an enumeration
type StrandType int

const (
	// Watson strand of referenceSeq: Watson strand of reference sequence.
	StrandType_Watson StrandType = iota
	// Crick strand of referenceSeq: Crick strand of reference sequence.
	StrandType_Crick StrandType = iota
)

// String converts the enum to its string representation
func (e StrandType) String() string {
	switch e {
	case StrandType_Watson: return "Watson strand of referenceSeq"
	case StrandType_Crick: return "Crick strand of referenceSeq"
	default: return "Unknown"
	}
}
