// Code generated by FHIR Generator. DO NOT EDIT.
// This value set includes a smattering of Service Place codes.
package fhir_r4b_go

// ExampleServicePlaceCodes represents the FHIR ValueSet as an enumeration
type ExampleServicePlaceCodes int

const (
	// Pharmacy: A facility or location where drugs and other medically related items and services are sold, dispensed, or otherwise provided directly to patients.
	ExampleServicePlaceCodes_Value01 ExampleServicePlaceCodes = iota
	// School: A facility whose primary purpose is education.
	ExampleServicePlaceCodes_Value03 ExampleServicePlaceCodes = iota
	// Homeless Shelter: A facility or location whose primary purpose is to provide temporary housing to homeless individuals (e.g., emergency shelters, individual or family shelters).
	ExampleServicePlaceCodes_Value04 ExampleServicePlaceCodes = iota
	// Indian Health Service Free-standing Facility: A facility or location, owned and operated by the Indian Health Service, which provides diagnostic, therapeutic (surgical and nonsurgical), and rehabilitation services to American Indians and Alaska Natives who do not require hospitalization.
	ExampleServicePlaceCodes_Value05 ExampleServicePlaceCodes = iota
	// Indian Health Service Provider-based Facility: A facility or location, owned and operated by the Indian Health Service, which provides diagnostic, therapeutic (surgical and nonsurgical), and rehabilitation services rendered by, or under the supervision of, physicians to American Indians and Alaska Natives admitted as inpatients or outpatients.
	ExampleServicePlaceCodes_Value06 ExampleServicePlaceCodes = iota
	// Tribal 638 Free-Standing Facility: A facility or location owned and operated by a federally recognized American Indian or Alaska Native tribe or tribal organization under a 638 agreement, which provides diagnostic, therapeutic (surgical and nonsurgical), and rehabilitation services to tribal members who do not require hospitalization.
	ExampleServicePlaceCodes_Value07 ExampleServicePlaceCodes = iota
	// Tribal 638 Provider-Based Facility: A facility or location owned and operated by a federally recognized American Indian or Alaska Native tribe or tribal organization under a 638 agreement, which provides diagnostic, therapeutic (surgical and nonsurgical), and rehabilitation services to tribal members admitted as inpatients or outpatients.
	ExampleServicePlaceCodes_Value08 ExampleServicePlaceCodes = iota
	// Prison/Correctional Facility: A prison, jail, reformatory, work farm, detention center, or any other similar facility maintained by either Federal, State or local authorities for the purpose of confinement or rehabilitation of adult or juvenile criminal offenders.
	ExampleServicePlaceCodes_Value09 ExampleServicePlaceCodes = iota
	// Office: Location, other than a hospital, skilled nursing facility (SNF), military treatment facility, community health center, State or local public health clinic, or intermediate care facility (ICF), where the health professional routinely provides health examinations, diagnosis, and treatment of illness or injury on an ambulatory basis.
	ExampleServicePlaceCodes_Value11 ExampleServicePlaceCodes = iota
	// Home: Location, other than a hospital or other facility, where the patient receives care in a private residence.
	ExampleServicePlaceCodes_Value12 ExampleServicePlaceCodes = iota
	// Assisted Living Fa: Congregate residential facility with self-contained living units providing assessment of each resident's needs and on-site support 24 hours a day, 7 days a week, with the capacity to deliver or arrange for services including some health care and other services.
	ExampleServicePlaceCodes_Value13 ExampleServicePlaceCodes = iota
	// Group Home: A residence, with shared living areas, where clients receive supervision and other services such as social and/or behavioral services, custodial service, and minimal services (e.g., medication administration).
	ExampleServicePlaceCodes_Value14 ExampleServicePlaceCodes = iota
	// Mobile Unit: A facility/unit that moves from place-to-place equipped to provide preventive, screening, diagnostic, and/or treatment services.
	ExampleServicePlaceCodes_Value15 ExampleServicePlaceCodes = iota
	// Off Campus-Outpatient Hospital: portion of an off-campus hospital provider-based department which provides diagnostic, therapeutic (both surgical and nonsurgical), and rehabilitation services to sick or injured persons who do not require hospitalization or institutionalization.
	ExampleServicePlaceCodes_Value19 ExampleServicePlaceCodes = iota
	// Urgent Care Facility: Location, distinct from a hospital emergency room, an office, or a clinic, whose purpose is to diagnose and treat illness or injury for unscheduled, ambulatory patients seeking immediate medical attention.
	ExampleServicePlaceCodes_Value20 ExampleServicePlaceCodes = iota
	// Inpatient Hospital: A facility, other than psychiatric, which primarily provides diagnostic, therapeutic (both surgical and nonsurgical), and rehabilitation services by, or under, the supervision of physicians to patients admitted for a variety of medical conditions.
	ExampleServicePlaceCodes_Value21 ExampleServicePlaceCodes = iota
	// Ambulance—Land: A land vehicle specifically designed, equipped and staffed for lifesaving and transporting the sick or injured.
	ExampleServicePlaceCodes_Value41 ExampleServicePlaceCodes = iota
)

// String converts the enum to its string representation
func (e ExampleServicePlaceCodes) String() string {
	switch e {
	case ExampleServicePlaceCodes_Value01: return "Pharmacy"
	case ExampleServicePlaceCodes_Value03: return "School"
	case ExampleServicePlaceCodes_Value04: return "Homeless Shelter"
	case ExampleServicePlaceCodes_Value05: return "Indian Health Service Free-standing Facility"
	case ExampleServicePlaceCodes_Value06: return "Indian Health Service Provider-based Facility"
	case ExampleServicePlaceCodes_Value07: return "Tribal 638 Free-Standing Facility"
	case ExampleServicePlaceCodes_Value08: return "Tribal 638 Provider-Based Facility"
	case ExampleServicePlaceCodes_Value09: return "Prison/Correctional Facility"
	case ExampleServicePlaceCodes_Value11: return "Office"
	case ExampleServicePlaceCodes_Value12: return "Home"
	case ExampleServicePlaceCodes_Value13: return "Assisted Living Fa"
	case ExampleServicePlaceCodes_Value14: return "Group Home"
	case ExampleServicePlaceCodes_Value15: return "Mobile Unit"
	case ExampleServicePlaceCodes_Value19: return "Off Campus-Outpatient Hospital"
	case ExampleServicePlaceCodes_Value20: return "Urgent Care Facility"
	case ExampleServicePlaceCodes_Value21: return "Inpatient Hospital"
	case ExampleServicePlaceCodes_Value41: return "Ambulance—Land"
	default: return "Unknown"
	}
}
