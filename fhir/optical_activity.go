// Code generated by FHIR Generator. DO NOT EDIT.
// The optical rotation type of a substance.
package fhir_r4b_go

// OpticalActivity represents the FHIR ValueSet as an enumeration
type OpticalActivity int

const (
	// dextrorotary: 
	OpticalActivity__ OpticalActivity = iota
	// levorotary: 
	OpticalActivity__ OpticalActivity = iota
)

// String converts the enum to its string representation
func (e OpticalActivity) String() string {
	switch e {
	case OpticalActivity__: return "dextrorotary"
	case OpticalActivity__: return "levorotary"
	default: return "Unknown"
	}
}
