// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"
)

// Contributor
// A contributor to the content of a knowledge asset, including authors, editors, reviewers, and endorsers.
type Contributor struct {
	extends DataType
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	Type *ContributorType `json:"type,omitempty"`
	Name *FhirString `json:"name,omitempty"`
	Contact []*ContactDetail `json:"contact,omitempty"`
}

// NewContributor creates a new Contributor instance.
func NewContributor() *Contributor {
	return &Contributor{}
}

// UnmarshalJSON populates Contributor from JSON data.
func (m *Contributor) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		Type *ContributorType `json:"type,omitempty"`
		Name *FhirString `json:"name,omitempty"`
		Contact []*ContactDetail `json:"contact,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.Type = temp.Type
	m.Name = temp.Name
	m.Contact = temp.Contact
	return nil
}

// MarshalJSON converts Contributor to JSON data.
func (m *Contributor) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		Type *ContributorType `json:"type,omitempty"`
		Name interface{} `json:"name,omitempty"`
		NameElement map[string]interface{} `json:"_name,omitempty"`
		Contact []*ContactDetail `json:"contact,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.Type = m.Type
	if m.Name != nil && m.Name.Value != nil {
		output.Name = m.Name.Value
		if m.Name.Element != nil {
			output.NameElement = toMapOrNil(m.Name.Element.MarshalJSON())
		}
	}
	output.Contact = m.Contact
	return json.Marshal(output)
}

// Clone creates a deep copy of Contributor.
func (m *Contributor) Clone() *Contributor {
	if m == nil { return nil }
	return &Contributor{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		Type: m.Type.Clone(),
		Name: m.Name.Clone(),
		Contact: cloneSlices(m.Contact),
	}
}

// Equals checks equality between two Contributor instances.
func (m *Contributor) Equals(other *Contributor) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !m.Type.Equals(other.Type) { return false }
	if !m.Name.Equals(other.Name) { return false }
	if !compareSlices(m.Contact, other.Contact) { return false }
	return true
}

