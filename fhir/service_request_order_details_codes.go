// Code generated by FHIR Generator. DO NOT EDIT.
// An example value set of Codified order entry details concepts. These concepts only make sense in the context of what is being ordered. This example is for a patient ventilation order
package fhir_r4b_go

// ServiceRequestOrderDetailsCodes represents the FHIR ValueSet as an enumeration
type ServiceRequestOrderDetailsCodes int

const (
	// Continuous positive airway pressure ventilation treatment (regime/therapy): 
	ServiceRequestOrderDetailsCodes_Value47545007 ServiceRequestOrderDetailsCodes = iota
	// Pressure controlled ventilation (procedure): 
	ServiceRequestOrderDetailsCodes_Value286812008 ServiceRequestOrderDetailsCodes = iota
	// Patient triggered inspiratory assistance (procedure): 
	ServiceRequestOrderDetailsCodes_Value243144002 ServiceRequestOrderDetailsCodes = iota
	// Assisted controlled mandatory ventilation (procedure): 
	ServiceRequestOrderDetailsCodes_Value243150007 ServiceRequestOrderDetailsCodes = iota
	// Synchronized intermittent mandatory ventilation (procedure): 
	ServiceRequestOrderDetailsCodes_Value59427005 ServiceRequestOrderDetailsCodes = iota
)

// String converts the enum to its string representation
func (e ServiceRequestOrderDetailsCodes) String() string {
	switch e {
	case ServiceRequestOrderDetailsCodes_Value47545007: return "Continuous positive airway pressure ventilation treatment (regime/therapy)"
	case ServiceRequestOrderDetailsCodes_Value286812008: return "Pressure controlled ventilation (procedure)"
	case ServiceRequestOrderDetailsCodes_Value243144002: return "Patient triggered inspiratory assistance (procedure)"
	case ServiceRequestOrderDetailsCodes_Value243150007: return "Assisted controlled mandatory ventilation (procedure)"
	case ServiceRequestOrderDetailsCodes_Value59427005: return "Synchronized intermittent mandatory ventilation (procedure)"
	default: return "Unknown"
	}
}
