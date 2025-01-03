// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"
)

// UsageContext
// Specifies clinical/business/etc. metadata that can be used to retrieve, index and/or categorize an artifact. This metadata can either be specific to the applicable population (e.g., age category, DRG) or the specific context of care (e.g., venue, care setting, provider of care).
type UsageContext struct {
	extends DataType
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	Code *Coding `json:"code,omitempty"`
	ValueCodeableConcept *CodeableConcept `json:"valuecodeableconcept,omitempty"`
	ValueQuantity *Quantity `json:"valuequantity,omitempty"`
	ValueRange *Range `json:"valuerange,omitempty"`
	ValueReference *Reference `json:"valuereference,omitempty"`
}

// NewUsageContext creates a new UsageContext instance.
func NewUsageContext() *UsageContext {
	return &UsageContext{}
}

// UnmarshalJSON populates UsageContext from JSON data.
func (m *UsageContext) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		Code *Coding `json:"code,omitempty"`
		ValueCodeableConcept *CodeableConcept `json:"valuecodeableconcept,omitempty"`
		ValueQuantity *Quantity `json:"valuequantity,omitempty"`
		ValueRange *Range `json:"valuerange,omitempty"`
		ValueReference *Reference `json:"valuereference,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.Code = temp.Code
	m.ValueCodeableConcept = temp.ValueCodeableConcept
	m.ValueQuantity = temp.ValueQuantity
	m.ValueRange = temp.ValueRange
	m.ValueReference = temp.ValueReference
	return nil
}

// MarshalJSON converts UsageContext to JSON data.
func (m *UsageContext) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		Code *Coding `json:"code,omitempty"`
		ValueCodeableConcept *CodeableConcept `json:"valuecodeableconcept,omitempty"`
		ValueQuantity *Quantity `json:"valuequantity,omitempty"`
		ValueRange *Range `json:"valuerange,omitempty"`
		ValueReference *Reference `json:"valuereference,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.Code = m.Code
	output.ValueCodeableConcept = m.ValueCodeableConcept
	output.ValueQuantity = m.ValueQuantity
	output.ValueRange = m.ValueRange
	output.ValueReference = m.ValueReference
	return json.Marshal(output)
}

// Clone creates a deep copy of UsageContext.
func (m *UsageContext) Clone() *UsageContext {
	if m == nil { return nil }
	return &UsageContext{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		Code: m.Code.Clone(),
		ValueCodeableConcept: m.ValueCodeableConcept.Clone(),
		ValueQuantity: m.ValueQuantity.Clone(),
		ValueRange: m.ValueRange.Clone(),
		ValueReference: m.ValueReference.Clone(),
	}
}

// Equals checks equality between two UsageContext instances.
func (m *UsageContext) Equals(other *UsageContext) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !m.Code.Equals(other.Code) { return false }
	if !m.ValueCodeableConcept.Equals(other.ValueCodeableConcept) { return false }
	if !m.ValueQuantity.Equals(other.ValueQuantity) { return false }
	if !m.ValueRange.Equals(other.ValueRange) { return false }
	if !m.ValueReference.Equals(other.ValueReference) { return false }
	return true
}

