// Code generated by FHIR Generator. DO NOT EDIT.
// Performer function of an agent in an imaging study series
package fhir_r4b_go

// ImagingStudySeriesPerformerFunction represents the FHIR ValueSet as an enumeration
type ImagingStudySeriesPerformerFunction int

const (
	// consultant: 
	ImagingStudySeriesPerformerFunction_CON ImagingStudySeriesPerformerFunction = iota
	// verifier: 
	ImagingStudySeriesPerformerFunction_VRF ImagingStudySeriesPerformerFunction = iota
	// performer: 
	ImagingStudySeriesPerformerFunction_PRF ImagingStudySeriesPerformerFunction = iota
	// secondary performer: 
	ImagingStudySeriesPerformerFunction_SPRF ImagingStudySeriesPerformerFunction = iota
	// referrer: 
	ImagingStudySeriesPerformerFunction_REF ImagingStudySeriesPerformerFunction = iota
)

// String converts the enum to its string representation
func (e ImagingStudySeriesPerformerFunction) String() string {
	switch e {
	case ImagingStudySeriesPerformerFunction_CON: return "consultant"
	case ImagingStudySeriesPerformerFunction_VRF: return "verifier"
	case ImagingStudySeriesPerformerFunction_PRF: return "performer"
	case ImagingStudySeriesPerformerFunction_SPRF: return "secondary performer"
	case ImagingStudySeriesPerformerFunction_REF: return "referrer"
	default: return "Unknown"
	}
}
