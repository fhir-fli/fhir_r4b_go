// Code generated by FHIR Generator. DO NOT EDIT.
// NLM codes Internet or Print.
package fhir_r4b_go

// CitedMedium represents the FHIR ValueSet as an enumeration
type CitedMedium int

const (
	// Internet: Online publication in a periodic release. Used to match NLM JournalIssue CitedMedium code for online version.
	CitedMedium_Internet CitedMedium = iota
	// Print: Print publication in a periodic release. Used to match NLM JournalIssue CitedMedium code for print version.
	CitedMedium_Print CitedMedium = iota
	// Offline Digital Storage: Publication in a physical device for electronic data storage, organized in issues for periodic release.
	CitedMedium_Offline_digital_storage CitedMedium = iota
	// Internet without issue: Online publication without any periodic release. Used for article specific publication date which could be the same as or different from journal issue publication date.
	CitedMedium_Internet_without_issue CitedMedium = iota
	// Print without issue: Print publication without any periodic release.
	CitedMedium_Print_without_issue CitedMedium = iota
	// Offline Digital Storage without issue: Publication in a physical device for electronic data storage, without any periodic release.
	CitedMedium_Offline_digital_storage_without_issue CitedMedium = iota
)

// String converts the enum to its string representation
func (e CitedMedium) String() string {
	switch e {
	case CitedMedium_Internet: return "Internet"
	case CitedMedium_Print: return "Print"
	case CitedMedium_Offline_digital_storage: return "Offline Digital Storage"
	case CitedMedium_Internet_without_issue: return "Internet without issue"
	case CitedMedium_Print_without_issue: return "Print without issue"
	case CitedMedium_Offline_digital_storage_without_issue: return "Offline Digital Storage without issue"
	default: return "Unknown"
	}
}
