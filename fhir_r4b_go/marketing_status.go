// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json")

// MarketingStatus
// The marketing status describes the date when a medicinal product is actually put on the market or the date as of which it is no longer available.
type MarketingStatus struct {
	BackboneType
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Country *CodeableConcept `json:"country,omitempty"`
	Jurisdiction *CodeableConcept `json:"jurisdiction,omitempty"`
	Status *CodeableConcept `json:"status,omitempty"`
	DateRange *Period `json:"daterange,omitempty"`
	RestoreDate *FhirDateTime `json:"restoredate,omitempty"`
}

// NewMarketingStatus creates a new MarketingStatus instance
func NewMarketingStatus() *MarketingStatus {
	return &MarketingStatus{}
}

// FromJSON populates MarketingStatus from JSON data
func (m *MarketingStatus) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts MarketingStatus to JSON data
func (m *MarketingStatus) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of MarketingStatus
func (m *MarketingStatus) Clone() *MarketingStatus {
	if m == nil { return nil }
	return &MarketingStatus{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Country: m.Country.Clone(),
		Jurisdiction: m.Jurisdiction.Clone(),
		Status: m.Status.Clone(),
		DateRange: m.DateRange.Clone(),
		RestoreDate: m.RestoreDate.Clone(),
	}
}

// Equals checks for equality with another MarketingStatus instance
func (m *MarketingStatus) Equals(other *MarketingStatus) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Country.Equals(other.Country) { return false }
	if !m.Jurisdiction.Equals(other.Jurisdiction) { return false }
	if !m.Status.Equals(other.Status) { return false }
	if !m.DateRange.Equals(other.DateRange) { return false }
	if !m.RestoreDate.Equals(other.RestoreDate) { return false }
	return true
}
