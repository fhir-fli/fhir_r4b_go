// Code generated by FHIR Generator. DO NOT EDIT.
// The intended usage for supplemental data elements in the measure.
package fhir_r4b_go

// MeasureDataUsage represents the FHIR ValueSet as an enumeration
type MeasureDataUsage int

const (
	// Supplemental Data: The data is intended to be provided as additional information alongside the measure results.
	MeasureDataUsage_Supplemental_data MeasureDataUsage = iota
	// Risk Adjustment Factor: The data is intended to be used to calculate and apply a risk adjustment model for the measure.
	MeasureDataUsage_Risk_adjustment_factor MeasureDataUsage = iota
)

// String converts the enum to its string representation
func (e MeasureDataUsage) String() string {
	switch e {
	case MeasureDataUsage_Supplemental_data: return "Supplemental Data"
	case MeasureDataUsage_Risk_adjustment_factor: return "Risk Adjustment Factor"
	default: return "Unknown"
	}
}
