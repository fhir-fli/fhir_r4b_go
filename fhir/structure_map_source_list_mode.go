// Code generated by FHIR Generator. DO NOT EDIT.
// If field is a list, how to manage the source.
package fhir_r4b_go

// StructureMapSourceListMode represents the FHIR ValueSet as an enumeration
type StructureMapSourceListMode int

const (
	// First: Only process this rule for the first in the list.
	StructureMapSourceListMode_First StructureMapSourceListMode = iota
	// All but the first: Process this rule for all but the first.
	StructureMapSourceListMode_Not_first StructureMapSourceListMode = iota
	// Last: Only process this rule for the last in the list.
	StructureMapSourceListMode_Last StructureMapSourceListMode = iota
	// All but the last: Process this rule for all but the last.
	StructureMapSourceListMode_Not_last StructureMapSourceListMode = iota
	// Enforce only one: Only process this rule is there is only item.
	StructureMapSourceListMode_Only_one StructureMapSourceListMode = iota
)

// String converts the enum to its string representation
func (e StructureMapSourceListMode) String() string {
	switch e {
	case StructureMapSourceListMode_First: return "First"
	case StructureMapSourceListMode_Not_first: return "All but the first"
	case StructureMapSourceListMode_Last: return "Last"
	case StructureMapSourceListMode_Not_last: return "All but the last"
	case StructureMapSourceListMode_Only_one: return "Enforce only one"
	default: return "Unknown"
	}
}
