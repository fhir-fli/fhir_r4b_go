// Code generated by FHIR Generator. DO NOT EDIT.
// This value set includes sample Claim Care Team Role codes.
package fhir_r4b_go

// ClaimCareTeamRoleCodes represents the FHIR ValueSet as an enumeration
type ClaimCareTeamRoleCodes int

const (
	// Primary provider: The primary care provider.
	ClaimCareTeamRoleCodes_Primary ClaimCareTeamRoleCodes = iota
	// Assisting Provider: Assisting care provider.
	ClaimCareTeamRoleCodes_Assist ClaimCareTeamRoleCodes = iota
	// Supervising Provider: Supervising care provider.
	ClaimCareTeamRoleCodes_Supervisor ClaimCareTeamRoleCodes = iota
	// Other: Other role on the care team.
	ClaimCareTeamRoleCodes_Other ClaimCareTeamRoleCodes = iota
)

// String converts the enum to its string representation
func (e ClaimCareTeamRoleCodes) String() string {
	switch e {
	case ClaimCareTeamRoleCodes_Primary: return "Primary provider"
	case ClaimCareTeamRoleCodes_Assist: return "Assisting Provider"
	case ClaimCareTeamRoleCodes_Supervisor: return "Supervising Provider"
	case ClaimCareTeamRoleCodes_Other: return "Other"
	default: return "Unknown"
	}
}
