// Code generated by FHIR Generator. DO NOT EDIT.
// This value set includes SNOMED CT codes for materials that specimen containers are made of
package fhir_r4b_go

// ContainerMaterials represents the FHIR ValueSet as an enumeration
type ContainerMaterials int

const (
	// glass: 
	ContainerMaterials_Value32039001 ContainerMaterials = iota
	// plastic: 
	ContainerMaterials_Value61088005 ContainerMaterials = iota
	// metal: 
	ContainerMaterials_Value425620007 ContainerMaterials = iota
)

// String converts the enum to its string representation
func (e ContainerMaterials) String() string {
	switch e {
	case ContainerMaterials_Value32039001: return "glass"
	case ContainerMaterials_Value61088005: return "plastic"
	case ContainerMaterials_Value425620007: return "metal"
	default: return "Unknown"
	}
}
