// Code generated by FHIR Generator. DO NOT EDIT.
// The method used to elucidate the structure or characterization of the drug substance.
package fhir_r4b_go

// StructureTechnique represents the FHIR ValueSet as an enumeration
type StructureTechnique int

const (
	// X-ray: 
	StructureTechnique_X_Ray StructureTechnique = iota
	// HPLC: 
	StructureTechnique_HPLC StructureTechnique = iota
	// NMR: 
	StructureTechnique_NMR StructureTechnique = iota
	// Peptide mapping: 
	StructureTechnique_PeptideMapping StructureTechnique = iota
	// Ligand binding assay: 
	StructureTechnique_LigandBindingAssay StructureTechnique = iota
)

// String converts the enum to its string representation
func (e StructureTechnique) String() string {
	switch e {
	case StructureTechnique_X_Ray: return "X-ray"
	case StructureTechnique_HPLC: return "HPLC"
	case StructureTechnique_NMR: return "NMR"
	case StructureTechnique_PeptideMapping: return "Peptide mapping"
	case StructureTechnique_LigandBindingAssay: return "Ligand binding assay"
	default: return "Unknown"
	}
}
