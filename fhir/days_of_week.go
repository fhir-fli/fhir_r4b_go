// Code generated by FHIR Generator. DO NOT EDIT.
// The days of the week.
package fhir_r4b_go

// DaysOfWeek represents the FHIR ValueSet as an enumeration
type DaysOfWeek int

const (
	// Monday: Monday.
	DaysOfWeek_Mon DaysOfWeek = iota
	// Tuesday: Tuesday.
	DaysOfWeek_Tue DaysOfWeek = iota
	// Wednesday: Wednesday.
	DaysOfWeek_Wed DaysOfWeek = iota
	// Thursday: Thursday.
	DaysOfWeek_Thu DaysOfWeek = iota
	// Friday: Friday.
	DaysOfWeek_Fri DaysOfWeek = iota
	// Saturday: Saturday.
	DaysOfWeek_Sat DaysOfWeek = iota
	// Sunday: Sunday.
	DaysOfWeek_Sun DaysOfWeek = iota
)

// String converts the enum to its string representation
func (e DaysOfWeek) String() string {
	switch e {
	case DaysOfWeek_Mon: return "Monday"
	case DaysOfWeek_Tue: return "Tuesday"
	case DaysOfWeek_Wed: return "Wednesday"
	case DaysOfWeek_Thu: return "Thursday"
	case DaysOfWeek_Fri: return "Friday"
	case DaysOfWeek_Sat: return "Saturday"
	case DaysOfWeek_Sun: return "Sunday"
	default: return "Unknown"
	}
}
