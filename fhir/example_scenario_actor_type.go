// Code generated by FHIR Generator. DO NOT EDIT.
// The type of actor - system or human.
package fhir_r4b_go

// ExampleScenarioActorType represents the FHIR ValueSet as an enumeration
type ExampleScenarioActorType int

const (
	// Person: A person.
	ExampleScenarioActorType_Person ExampleScenarioActorType = iota
	// System: A system.
	ExampleScenarioActorType_Entity ExampleScenarioActorType = iota
)

// String converts the enum to its string representation
func (e ExampleScenarioActorType) String() string {
	switch e {
	case ExampleScenarioActorType_Person: return "Person"
	case ExampleScenarioActorType_Entity: return "System"
	default: return "Unknown"
	}
}
