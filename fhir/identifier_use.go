// Code generated by FHIR Generator. DO NOT EDIT.
// Identifies the purpose for this identifier, if known .
package fhir_r4b_go

// IdentifierUse represents the FHIR ValueSet as an enumeration
type IdentifierUse int

const (
	// Usual: The identifier recommended for display and use in real-world interactions.
	IdentifierUse_Usual IdentifierUse = iota
	// Official: The identifier considered to be most trusted for the identification of this item. Sometimes also known as "primary" and "main". The determination of "official" is subjective and implementation guides often provide additional guidelines for use.
	IdentifierUse_Official IdentifierUse = iota
	// Temp: A temporary identifier.
	IdentifierUse_Temp IdentifierUse = iota
	// Secondary: An identifier that was assigned in secondary use - it serves to identify the object in a relative context, but cannot be consistently assigned to the same object again in a different context.
	IdentifierUse_Secondary IdentifierUse = iota
	// Old: The identifier id no longer considered valid, but may be relevant for search purposes. E.g. Changes to identifier schemes, account merges, etc.
	IdentifierUse_Old IdentifierUse = iota
)

// String converts the enum to its string representation
func (e IdentifierUse) String() string {
	switch e {
	case IdentifierUse_Usual: return "Usual"
	case IdentifierUse_Official: return "Official"
	case IdentifierUse_Temp: return "Temp"
	case IdentifierUse_Secondary: return "Secondary"
	case IdentifierUse_Old: return "Old"
	default: return "Unknown"
	}
}
