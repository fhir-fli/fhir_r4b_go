// Code generated by FHIR Generator. DO NOT EDIT.
// A code that indicates how the server supports conditional delete.
package fhir_r4b_go

// ConditionalDeleteStatus represents the FHIR ValueSet as an enumeration
type ConditionalDeleteStatus int

const (
	// Not Supported: No support for conditional deletes.
	ConditionalDeleteStatus_Not_supported ConditionalDeleteStatus = iota
	// Single Deletes Supported: Conditional deletes are supported, but only single resources at a time.
	ConditionalDeleteStatus_Single ConditionalDeleteStatus = iota
	// Multiple Deletes Supported: Conditional deletes are supported, and multiple resources can be deleted in a single interaction.
	ConditionalDeleteStatus_Multiple ConditionalDeleteStatus = iota
)

// String converts the enum to its string representation
func (e ConditionalDeleteStatus) String() string {
	switch e {
	case ConditionalDeleteStatus_Not_supported: return "Not Supported"
	case ConditionalDeleteStatus_Single: return "Single Deletes Supported"
	case ConditionalDeleteStatus_Multiple: return "Multiple Deletes Supported"
	default: return "Unknown"
	}
}
