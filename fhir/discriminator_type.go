// Code generated by FHIR Generator. DO NOT EDIT.
// How an element value is interpreted when discrimination is evaluated.
package fhir_r4b_go

// DiscriminatorType represents the FHIR ValueSet as an enumeration
type DiscriminatorType int

const (
	// Value: The slices have different values in the nominated element.
	DiscriminatorType_Value DiscriminatorType = iota
	// Exists: The slices are differentiated by the presence or absence of the nominated element.
	DiscriminatorType_Exists DiscriminatorType = iota
	// Pattern: The slices have different values in the nominated element, as determined by testing them against the applicable ElementDefinition.pattern[x].
	DiscriminatorType_Pattern DiscriminatorType = iota
	// Type: The slices are differentiated by type of the nominated element.
	DiscriminatorType_Type_ DiscriminatorType = iota
	// Profile: The slices are differentiated by conformance of the nominated element to a specified profile. Note that if the path specifies .resolve() then the profile is the target profile on the reference. In this case, validation by the possible profiles is required to differentiate the slices.
	DiscriminatorType_Profile DiscriminatorType = iota
)

// String converts the enum to its string representation
func (e DiscriminatorType) String() string {
	switch e {
	case DiscriminatorType_Value: return "Value"
	case DiscriminatorType_Exists: return "Exists"
	case DiscriminatorType_Pattern: return "Pattern"
	case DiscriminatorType_Type_: return "Type"
	case DiscriminatorType_Profile: return "Profile"
	default: return "Unknown"
	}
}
