// Code generated by FHIR Generator. DO NOT EDIT.
// The meaning of the hierarchy of concepts in a code system.
package fhir_r4b_go

// CodeSystemHierarchyMeaning represents the FHIR ValueSet as an enumeration
type CodeSystemHierarchyMeaning int

const (
	// Grouped By: No particular relationship between the concepts can be assumed, except what can be determined by inspection of the definitions of the elements (possible reasons to use this: importing from a source where this is not defined, or where various parts of the hierarchy have different meanings).
	CodeSystemHierarchyMeaning_Grouped_by CodeSystemHierarchyMeaning = iota
	// Is-A: A hierarchy where the child concepts have an IS-A relationship with the parents - that is, all the properties of the parent are also true for its child concepts. Not that is-a is a property of the concepts, so additional subsumption relationships may be defined using properties or the [subsumes](extension-codesystem-subsumes.html) extension.
	CodeSystemHierarchyMeaning_Is_a CodeSystemHierarchyMeaning = iota
	// Part Of: Child elements list the individual parts of a composite whole (e.g. body site).
	CodeSystemHierarchyMeaning_Part_of CodeSystemHierarchyMeaning = iota
	// Classified With: Child concepts in the hierarchy may have only one parent, and there is a presumption that the code system is a "closed world" meaning all things must be in the hierarchy. This results in concepts such as "not otherwise classified.".
	CodeSystemHierarchyMeaning_Classified_with CodeSystemHierarchyMeaning = iota
)

// String converts the enum to its string representation
func (e CodeSystemHierarchyMeaning) String() string {
	switch e {
	case CodeSystemHierarchyMeaning_Grouped_by: return "Grouped By"
	case CodeSystemHierarchyMeaning_Is_a: return "Is-A"
	case CodeSystemHierarchyMeaning_Part_of: return "Part Of"
	case CodeSystemHierarchyMeaning_Classified_with: return "Classified With"
	default: return "Unknown"
	}
}
