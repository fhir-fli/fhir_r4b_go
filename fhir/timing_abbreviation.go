// Code generated by FHIR Generator. DO NOT EDIT.
// Code for a known / defined timing pattern.
package fhir_r4b_go

// TimingAbbreviation represents the FHIR ValueSet as an enumeration
type TimingAbbreviation int

const (
	// BID: 
	TimingAbbreviation_BID TimingAbbreviation = iota
	// TID: 
	TimingAbbreviation_TID TimingAbbreviation = iota
	// QID: 
	TimingAbbreviation_QID TimingAbbreviation = iota
	// AM: 
	TimingAbbreviation_AM TimingAbbreviation = iota
	// PM: 
	TimingAbbreviation_PM TimingAbbreviation = iota
	// QD: 
	TimingAbbreviation_QD TimingAbbreviation = iota
	// QOD: 
	TimingAbbreviation_QOD TimingAbbreviation = iota
	// every hour: 
	TimingAbbreviation_Q1H TimingAbbreviation = iota
	// every 2 hours: 
	TimingAbbreviation_Q2H TimingAbbreviation = iota
	// every 3 hours: 
	TimingAbbreviation_Q3H TimingAbbreviation = iota
	// Q4H: 
	TimingAbbreviation_Q4H TimingAbbreviation = iota
	// Q6H: 
	TimingAbbreviation_Q6H TimingAbbreviation = iota
	// every 8 hours: 
	TimingAbbreviation_Q8H TimingAbbreviation = iota
	// at bedtime: 
	TimingAbbreviation_BED TimingAbbreviation = iota
	// weekly: 
	TimingAbbreviation_WK TimingAbbreviation = iota
	// monthly: 
	TimingAbbreviation_MO TimingAbbreviation = iota
)

// String converts the enum to its string representation
func (e TimingAbbreviation) String() string {
	switch e {
	case TimingAbbreviation_BID: return "BID"
	case TimingAbbreviation_TID: return "TID"
	case TimingAbbreviation_QID: return "QID"
	case TimingAbbreviation_AM: return "AM"
	case TimingAbbreviation_PM: return "PM"
	case TimingAbbreviation_QD: return "QD"
	case TimingAbbreviation_QOD: return "QOD"
	case TimingAbbreviation_Q1H: return "every hour"
	case TimingAbbreviation_Q2H: return "every 2 hours"
	case TimingAbbreviation_Q3H: return "every 3 hours"
	case TimingAbbreviation_Q4H: return "Q4H"
	case TimingAbbreviation_Q6H: return "Q6H"
	case TimingAbbreviation_Q8H: return "every 8 hours"
	case TimingAbbreviation_BED: return "at bedtime"
	case TimingAbbreviation_WK: return "weekly"
	case TimingAbbreviation_MO: return "monthly"
	default: return "Unknown"
	}
}
