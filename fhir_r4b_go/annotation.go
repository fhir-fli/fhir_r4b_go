// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json")

// Annotation
// A  text note which also  contains information about who made the statement and when.
type Annotation struct {
	DataType
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	AuthorReference *Reference `json:"authorreference,omitempty"`
	AuthorString *FhirString `json:"authorstring,omitempty"`
	Time *FhirDateTime `json:"time,omitempty"`
	Text *FhirMarkdown `json:"text,omitempty"`
}

// NewAnnotation creates a new Annotation instance
func NewAnnotation() *Annotation {
	return &Annotation{}
}

// FromJSON populates Annotation from JSON data
func (m *Annotation) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts Annotation to JSON data
func (m *Annotation) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of Annotation
func (m *Annotation) Clone() *Annotation {
	if m == nil { return nil }
	return &Annotation{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		AuthorReference: m.AuthorReference.Clone(),
		AuthorString: m.AuthorString.Clone(),
		Time: m.Time.Clone(),
		Text: m.Text.Clone(),
	}
}

// Equals checks for equality with another Annotation instance
func (m *Annotation) Equals(other *Annotation) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !m.AuthorReference.Equals(other.AuthorReference) { return false }
	if !m.AuthorString.Equals(other.AuthorString) { return false }
	if !m.Time.Equals(other.Time) { return false }
	if !m.Text.Equals(other.Text) { return false }
	return true
}
