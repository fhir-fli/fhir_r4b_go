// Code generated by FHIR Generator. DO NOT EDIT.
// A coded concept listing the eye codes.
package fhir_r4b_go

// VisionEyes represents the FHIR ValueSet as an enumeration
type VisionEyes int

const (
	// Right Eye: Right Eye.
	VisionEyes_Right VisionEyes = iota
	// Left Eye: Left Eye.
	VisionEyes_Left VisionEyes = iota
)

// String converts the enum to its string representation
func (e VisionEyes) String() string {
	switch e {
	case VisionEyes_Right: return "Right Eye"
	case VisionEyes_Left: return "Left Eye"
	default: return "Unknown"
	}
}
