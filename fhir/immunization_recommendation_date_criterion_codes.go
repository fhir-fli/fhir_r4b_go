// Code generated by FHIR Generator. DO NOT EDIT.
// The value set to instantiate this attribute should be drawn from a terminologically robust code system that consists of or contains concepts to support the definition of dates relevant to recommendations for future doses of vaccines. This value set is provided as a suggestive example.
package fhir_r4b_go

// ImmunizationRecommendationDateCriterionCodes represents the FHIR ValueSet as an enumeration
type ImmunizationRecommendationDateCriterionCodes int

const (
	// : 
	ImmunizationRecommendationDateCriterionCodes_Value30981_5 ImmunizationRecommendationDateCriterionCodes = iota
	// : 
	ImmunizationRecommendationDateCriterionCodes_Value30980_7 ImmunizationRecommendationDateCriterionCodes = iota
	// : 
	ImmunizationRecommendationDateCriterionCodes_Value59777_3 ImmunizationRecommendationDateCriterionCodes = iota
	// : 
	ImmunizationRecommendationDateCriterionCodes_Value59778_1 ImmunizationRecommendationDateCriterionCodes = iota
)

// String converts the enum to its string representation
func (e ImmunizationRecommendationDateCriterionCodes) String() string {
	switch e {
	case ImmunizationRecommendationDateCriterionCodes_Value30981_5: return ""
	case ImmunizationRecommendationDateCriterionCodes_Value30980_7: return ""
	case ImmunizationRecommendationDateCriterionCodes_Value59777_3: return ""
	case ImmunizationRecommendationDateCriterionCodes_Value59778_1: return ""
	default: return "Unknown"
	}
}
