// Code generated by FHIR Generator. DO NOT EDIT.
// A code that indicates how the server supports conditional read.
package fhir_r4b_go

// ConditionalReadStatus represents the FHIR ValueSet as an enumeration
type ConditionalReadStatus int

const (
	// Not Supported: No support for conditional reads.
	ConditionalReadStatus_Not_supported ConditionalReadStatus = iota
	// If-Modified-Since: Conditional reads are supported, but only with the If-Modified-Since HTTP Header.
	ConditionalReadStatus_Modified_since ConditionalReadStatus = iota
	// If-None-Match: Conditional reads are supported, but only with the If-None-Match HTTP Header.
	ConditionalReadStatus_Not_match ConditionalReadStatus = iota
	// Full Support: Conditional reads are supported, with both If-Modified-Since and If-None-Match HTTP Headers.
	ConditionalReadStatus_Full_support ConditionalReadStatus = iota
)

// String converts the enum to its string representation
func (e ConditionalReadStatus) String() string {
	switch e {
	case ConditionalReadStatus_Not_supported: return "Not Supported"
	case ConditionalReadStatus_Modified_since: return "If-Modified-Since"
	case ConditionalReadStatus_Not_match: return "If-None-Match"
	case ConditionalReadStatus_Full_support: return "Full Support"
	default: return "Unknown"
	}
}
