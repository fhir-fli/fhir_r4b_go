// Code generated by FHIR Generator. DO NOT EDIT.
// This value set includes a smattering of Prescription Product codes.
package fhir_r4b_go

// ExampleVisionPrescriptionProductCodes represents the FHIR ValueSet as an enumeration
type ExampleVisionPrescriptionProductCodes int

const (
	// Lens: A lens to be fitted to a frame to comprise a pair of glasses.
	ExampleVisionPrescriptionProductCodes_Lens ExampleVisionPrescriptionProductCodes = iota
	// Contact Lens: A lens to be fitted for wearing directly on an eye.
	ExampleVisionPrescriptionProductCodes_Contact ExampleVisionPrescriptionProductCodes = iota
)

// String converts the enum to its string representation
func (e ExampleVisionPrescriptionProductCodes) String() string {
	switch e {
	case ExampleVisionPrescriptionProductCodes_Lens: return "Lens"
	case ExampleVisionPrescriptionProductCodes_Contact: return "Contact Lens"
	default: return "Unknown"
	}
}
