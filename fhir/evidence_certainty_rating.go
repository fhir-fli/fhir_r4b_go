// Code generated by FHIR Generator. DO NOT EDIT.
// The assessment of quality, confidence, or certainty.
package fhir_r4b_go

// EvidenceCertaintyRating represents the FHIR ValueSet as an enumeration
type EvidenceCertaintyRating int

const (
	// High quality: High quality evidence.
	EvidenceCertaintyRating_High EvidenceCertaintyRating = iota
	// Moderate quality: Moderate quality evidence.
	EvidenceCertaintyRating_Moderate EvidenceCertaintyRating = iota
	// Low quality: Low quality evidence.
	EvidenceCertaintyRating_Low EvidenceCertaintyRating = iota
	// Very low quality: Very low quality evidence.
	EvidenceCertaintyRating_Very_low EvidenceCertaintyRating = iota
	// no serious concern: no serious concern.
	EvidenceCertaintyRating_No_concern EvidenceCertaintyRating = iota
	// serious concern: serious concern.
	EvidenceCertaintyRating_Serious_concern EvidenceCertaintyRating = iota
	// very serious concern: very serious concern.
	EvidenceCertaintyRating_Very_serious_concern EvidenceCertaintyRating = iota
	// extremely serious concern: extremely serious concern.
	EvidenceCertaintyRating_Extremely_serious_concern EvidenceCertaintyRating = iota
	// present: possible reason for increasing quality rating was checked and found to be present.
	EvidenceCertaintyRating_Present EvidenceCertaintyRating = iota
	// absent: possible reason for increasing quality rating was checked and found to be absent.
	EvidenceCertaintyRating_Absent EvidenceCertaintyRating = iota
	// no change to rating: no change to quality rating.
	EvidenceCertaintyRating_No_change EvidenceCertaintyRating = iota
	// reduce rating: -1: reduce quality rating by 1.
	EvidenceCertaintyRating_Downcode1 EvidenceCertaintyRating = iota
	// reduce rating: -2: reduce quality rating by 2.
	EvidenceCertaintyRating_Downcode2 EvidenceCertaintyRating = iota
	// reduce rating: -3: reduce quality rating by 3.
	EvidenceCertaintyRating_Downcode3 EvidenceCertaintyRating = iota
	// increase rating: +1: increase quality rating by 1.
	EvidenceCertaintyRating_Upcode1 EvidenceCertaintyRating = iota
	// increase rating: +2: increase quality rating by 2.
	EvidenceCertaintyRating_Upcode2 EvidenceCertaintyRating = iota
)

// String converts the enum to its string representation
func (e EvidenceCertaintyRating) String() string {
	switch e {
	case EvidenceCertaintyRating_High: return "High quality"
	case EvidenceCertaintyRating_Moderate: return "Moderate quality"
	case EvidenceCertaintyRating_Low: return "Low quality"
	case EvidenceCertaintyRating_Very_low: return "Very low quality"
	case EvidenceCertaintyRating_No_concern: return "no serious concern"
	case EvidenceCertaintyRating_Serious_concern: return "serious concern"
	case EvidenceCertaintyRating_Very_serious_concern: return "very serious concern"
	case EvidenceCertaintyRating_Extremely_serious_concern: return "extremely serious concern"
	case EvidenceCertaintyRating_Present: return "present"
	case EvidenceCertaintyRating_Absent: return "absent"
	case EvidenceCertaintyRating_No_change: return "no change to rating"
	case EvidenceCertaintyRating_Downcode1: return "reduce rating: -1"
	case EvidenceCertaintyRating_Downcode2: return "reduce rating: -2"
	case EvidenceCertaintyRating_Downcode3: return "reduce rating: -3"
	case EvidenceCertaintyRating_Upcode1: return "increase rating: +1"
	case EvidenceCertaintyRating_Upcode2: return "increase rating: +2"
	default: return "Unknown"
	}
}
