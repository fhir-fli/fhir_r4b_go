// Code generated by FHIR Generator. DO NOT EDIT.
// The value set to instantiate this attribute should be drawn from a terminologically robust code system that consists of or contains concepts to support describing the disease that the dose is being administered against. This value set is provided as a suggestive example and includes the SNOMED CT concepts from the 64572001 (Disease) hierarchy.
package fhir_r4b_go

// ImmunizationTargetDiseaseCodes represents the FHIR ValueSet as an enumeration
type ImmunizationTargetDiseaseCodes int

const (
	// : 
	ImmunizationTargetDiseaseCodes_Value1857005 ImmunizationTargetDiseaseCodes = iota
	// : 
	ImmunizationTargetDiseaseCodes_Value397430003 ImmunizationTargetDiseaseCodes = iota
	// : 
	ImmunizationTargetDiseaseCodes_Value14189004 ImmunizationTargetDiseaseCodes = iota
	// : 
	ImmunizationTargetDiseaseCodes_Value36989005 ImmunizationTargetDiseaseCodes = iota
	// : 
	ImmunizationTargetDiseaseCodes_Value36653000 ImmunizationTargetDiseaseCodes = iota
	// : 
	ImmunizationTargetDiseaseCodes_Value76902006 ImmunizationTargetDiseaseCodes = iota
	// : 
	ImmunizationTargetDiseaseCodes_Value709410003 ImmunizationTargetDiseaseCodes = iota
	// : 
	ImmunizationTargetDiseaseCodes_Value27836007 ImmunizationTargetDiseaseCodes = iota
	// : 
	ImmunizationTargetDiseaseCodes_Value398102009 ImmunizationTargetDiseaseCodes = iota
)

// String converts the enum to its string representation
func (e ImmunizationTargetDiseaseCodes) String() string {
	switch e {
	case ImmunizationTargetDiseaseCodes_Value1857005: return ""
	case ImmunizationTargetDiseaseCodes_Value397430003: return ""
	case ImmunizationTargetDiseaseCodes_Value14189004: return ""
	case ImmunizationTargetDiseaseCodes_Value36989005: return ""
	case ImmunizationTargetDiseaseCodes_Value36653000: return ""
	case ImmunizationTargetDiseaseCodes_Value76902006: return ""
	case ImmunizationTargetDiseaseCodes_Value709410003: return ""
	case ImmunizationTargetDiseaseCodes_Value27836007: return ""
	case ImmunizationTargetDiseaseCodes_Value398102009: return ""
	default: return "Unknown"
	}
}
