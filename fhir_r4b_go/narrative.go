// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"
)

// Narrative
// A human-readable summary of the resource conveying the essential clinical and business information for the resource.
type Narrative struct {
	extends DataType
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	Status *NarrativeStatus `json:"status,omitempty"`
	Div *FhirXhtml `json:"div,omitempty"`
}

// NewNarrative creates a new Narrative instance.
func NewNarrative() *Narrative {
	return &Narrative{}
}

// UnmarshalJSON populates Narrative from JSON data.
func (m *Narrative) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		Status *NarrativeStatus `json:"status,omitempty"`
		Div *FhirXhtml `json:"div,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.Status = temp.Status
	m.Div = temp.Div
	return nil
}

// MarshalJSON converts Narrative to JSON data.
func (m *Narrative) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		Status *NarrativeStatus `json:"status,omitempty"`
		Div interface{} `json:"div,omitempty"`
		DivElement map[string]interface{} `json:"_div,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.Status = m.Status
	if m.Div != nil && m.Div.Value != nil {
		output.Div = m.Div.Value
		if m.Div.Element != nil {
			output.DivElement = toMapOrNil(m.Div.Element.MarshalJSON())
		}
	}
	return json.Marshal(output)
}

// Clone creates a deep copy of Narrative.
func (m *Narrative) Clone() *Narrative {
	if m == nil { return nil }
	return &Narrative{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		Status: m.Status.Clone(),
		Div: m.Div.Clone(),
	}
}

// Equals checks equality between two Narrative instances.
func (m *Narrative) Equals(other *Narrative) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !m.Status.Equals(other.Status) { return false }
	if !m.Div.Equals(other.Div) { return false }
	return true
}

