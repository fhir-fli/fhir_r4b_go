// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json")

// EnrollmentResponse
// This resource provides enrollment and plan details from the processing of an EnrollmentRequest resource.
type EnrollmentResponse struct {
	DomainResource
	Id *FhirString `json:"id,omitempty"`
	Meta *FhirMeta `json:"meta,omitempty"`
	ImplicitRules *FhirUri `json:"implicitrules,omitempty"`
	Language *CommonLanguages `json:"language,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Contained []*Resource `json:"contained,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Status *FinancialResourceStatusCodes `json:"status,omitempty"`
	Request *Reference `json:"request,omitempty"`
	Outcome *RemittanceOutcome `json:"outcome,omitempty"`
	Disposition *FhirString `json:"disposition,omitempty"`
	Created *FhirDateTime `json:"created,omitempty"`
	Organization *Reference `json:"organization,omitempty"`
	RequestProvider *Reference `json:"requestprovider,omitempty"`
}

// NewEnrollmentResponse creates a new EnrollmentResponse instance
func NewEnrollmentResponse() *EnrollmentResponse {
	return &EnrollmentResponse{}
}

// FromJSON populates EnrollmentResponse from JSON data
func (m *EnrollmentResponse) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts EnrollmentResponse to JSON data
func (m *EnrollmentResponse) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of EnrollmentResponse
func (m *EnrollmentResponse) Clone() *EnrollmentResponse {
	if m == nil { return nil }
	return &EnrollmentResponse{
		Id: m.Id.Clone(),
		Meta: m.Meta.Clone(),
		ImplicitRules: m.ImplicitRules.Clone(),
		Language: m.Language.Clone(),
		Text: m.Text.Clone(),
		Contained: cloneSlices(m.Contained),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Identifier: cloneSlices(m.Identifier),
		Status: m.Status.Clone(),
		Request: m.Request.Clone(),
		Outcome: m.Outcome.Clone(),
		Disposition: m.Disposition.Clone(),
		Created: m.Created.Clone(),
		Organization: m.Organization.Clone(),
		RequestProvider: m.RequestProvider.Clone(),
	}
}

// Equals checks for equality with another EnrollmentResponse instance
func (m *EnrollmentResponse) Equals(other *EnrollmentResponse) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !m.Meta.Equals(other.Meta) { return false }
	if !m.ImplicitRules.Equals(other.ImplicitRules) { return false }
	if !m.Language.Equals(other.Language) { return false }
	if !m.Text.Equals(other.Text) { return false }
	if !compareSlices(m.Contained, other.Contained) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !compareSlices(m.Identifier, other.Identifier) { return false }
	if !m.Status.Equals(other.Status) { return false }
	if !m.Request.Equals(other.Request) { return false }
	if !m.Outcome.Equals(other.Outcome) { return false }
	if !m.Disposition.Equals(other.Disposition) { return false }
	if !m.Created.Equals(other.Created) { return false }
	if !m.Organization.Equals(other.Organization) { return false }
	if !m.RequestProvider.Equals(other.RequestProvider) { return false }
	return true
}
