// Code generated by FHIR Generator. DO NOT EDIT.
// Unified Code for Units of Measure (UCUM). This value set includes all UCUM codes
package fhir_r4b_go

// CommonUCUMCodesForDuration represents the FHIR ValueSet as an enumeration
type CommonUCUMCodesForDuration int

const (
	// milliseconds: 
	CommonUCUMCodesForDuration_Ms CommonUCUMCodesForDuration = iota
	// seconds: 
	CommonUCUMCodesForDuration_S CommonUCUMCodesForDuration = iota
	// minutes: 
	CommonUCUMCodesForDuration_Min CommonUCUMCodesForDuration = iota
	// hours: 
	CommonUCUMCodesForDuration_H CommonUCUMCodesForDuration = iota
	// days: 
	CommonUCUMCodesForDuration_D CommonUCUMCodesForDuration = iota
	// weeks: 
	CommonUCUMCodesForDuration_Wk CommonUCUMCodesForDuration = iota
	// months: 
	CommonUCUMCodesForDuration_Mo CommonUCUMCodesForDuration = iota
	// years: 
	CommonUCUMCodesForDuration_A CommonUCUMCodesForDuration = iota
)

// String converts the enum to its string representation
func (e CommonUCUMCodesForDuration) String() string {
	switch e {
	case CommonUCUMCodesForDuration_Ms: return "milliseconds"
	case CommonUCUMCodesForDuration_S: return "seconds"
	case CommonUCUMCodesForDuration_Min: return "minutes"
	case CommonUCUMCodesForDuration_H: return "hours"
	case CommonUCUMCodesForDuration_D: return "days"
	case CommonUCUMCodesForDuration_Wk: return "weeks"
	case CommonUCUMCodesForDuration_Mo: return "months"
	case CommonUCUMCodesForDuration_A: return "years"
	default: return "Unknown"
	}
}
