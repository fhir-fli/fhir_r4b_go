// Code generated by FHIR Generator. DO NOT EDIT.
// HTTP verbs (in the HTTP command line). See [HTTP rfc](https://tools.ietf.org/html/rfc7231) for details.
package fhir_r4b_go

// HTTPVerb represents the FHIR ValueSet as an enumeration
type HTTPVerb int

const (
	// GET: HTTP GET Command.
	HTTPVerb_GET HTTPVerb = iota
	// HEAD: HTTP HEAD Command.
	HTTPVerb_HEAD HTTPVerb = iota
	// POST: HTTP POST Command.
	HTTPVerb_POST HTTPVerb = iota
	// PUT: HTTP PUT Command.
	HTTPVerb_PUT HTTPVerb = iota
	// DELETE: HTTP DELETE Command.
	HTTPVerb_DELETE HTTPVerb = iota
	// PATCH: HTTP PATCH Command.
	HTTPVerb_PATCH HTTPVerb = iota
)

// String converts the enum to its string representation
func (e HTTPVerb) String() string {
	switch e {
	case HTTPVerb_GET: return "GET"
	case HTTPVerb_HEAD: return "HEAD"
	case HTTPVerb_POST: return "POST"
	case HTTPVerb_PUT: return "PUT"
	case HTTPVerb_DELETE: return "DELETE"
	case HTTPVerb_PATCH: return "PATCH"
	default: return "Unknown"
	}
}
