// Code generated by FHIR Generator. DO NOT EDIT.
// Detailed information about the type of the image - its kind, purpose, or the kind of equipment used to generate it.
package fhir_r4b_go

// MediaModality represents the FHIR ValueSet as an enumeration
type MediaModality int

const (
	// Diagram: A diagram. Often used in diagnostic reports
	MediaModality_Diagram MediaModality = iota
	// Fax: A digital record of a fax document
	MediaModality_Fax MediaModality = iota
	// Scanned Document: A digital scan of a document. This is reserved for when there is not enough metadata to create a document reference
	MediaModality_Scan MediaModality = iota
	// Retina Scan: A retinal image used for identification purposes
	MediaModality_Retina MediaModality = iota
	// Fingerprint: A finger print scan used for identification purposes
	MediaModality_Fingerprint MediaModality = iota
	// Iris Scan: An iris scan used for identification purposes
	MediaModality_Iris MediaModality = iota
	// Palm Scan: A palm scan used for identification purposes
	MediaModality_Palm MediaModality = iota
	// Face Scan: A face scan used for identification purposes
	MediaModality_Face MediaModality = iota
)

// String converts the enum to its string representation
func (e MediaModality) String() string {
	switch e {
	case MediaModality_Diagram: return "Diagram"
	case MediaModality_Fax: return "Fax"
	case MediaModality_Scan: return "Scanned Document"
	case MediaModality_Retina: return "Retina Scan"
	case MediaModality_Fingerprint: return "Fingerprint"
	case MediaModality_Iris: return "Iris Scan"
	case MediaModality_Palm: return "Palm Scan"
	case MediaModality_Face: return "Face Scan"
	default: return "Unknown"
	}
}
