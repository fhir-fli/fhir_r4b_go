package fhir_r4b_go

// FhirNumber interface for shared behavior across FHIR numeric types
type FhirNumber interface {
	Value() float64          // Retrieves the value as float64
	ToJSON() (map[string]interface{}, error) // Converts to JSON
	Equals(other FhirNumber) bool // Equality check
	Add(other FhirNumber) FhirNumber
	Subtract(other FhirNumber) FhirNumber
	Multiply(other FhirNumber) FhirNumber
	Divide(other FhirNumber) (FhirNumber, error)
	String() string
	Clone() FhirNumber
}
