// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json")

// ChargeItem
// The resource ChargeItem describes the provision of healthcare provider products for a certain patient, therefore referring not only to the product, but containing in addition details of the provision, like date, time, amounts and participating organizations and persons. Main Usage of the ChargeItem is to enable the billing process and internal cost allocation.
type ChargeItem struct {
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
	DefinitionUri []*FhirUri `json:"definitionuri,omitempty"`
	DefinitionCanonical []*FhirCanonical `json:"definitioncanonical,omitempty"`
	Status *ChargeItemStatus `json:"status,omitempty"`
	PartOf []*Reference `json:"partof,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	Subject *Reference `json:"subject,omitempty"`
	Context *Reference `json:"context,omitempty"`
	OccurrenceDateTime *FhirDateTime `json:"occurrencedatetime,omitempty"`
	OccurrencePeriod *Period `json:"occurrenceperiod,omitempty"`
	OccurrenceTiming *Timing `json:"occurrencetiming,omitempty"`
	Performer []*ChargeItemPerformer `json:"performer,omitempty"`
	PerformingOrganization *Reference `json:"performingorganization,omitempty"`
	RequestingOrganization *Reference `json:"requestingorganization,omitempty"`
	CostCenter *Reference `json:"costcenter,omitempty"`
	Quantity *Quantity `json:"quantity,omitempty"`
	Bodysite []*CodeableConcept `json:"bodysite,omitempty"`
	FactorOverride *FhirDecimal `json:"factoroverride,omitempty"`
	PriceOverride *Money `json:"priceoverride,omitempty"`
	OverrideReason *FhirString `json:"overridereason,omitempty"`
	Enterer *Reference `json:"enterer,omitempty"`
	EnteredDate *FhirDateTime `json:"entereddate,omitempty"`
	Reason []*CodeableConcept `json:"reason,omitempty"`
	Service []*Reference `json:"service,omitempty"`
	ProductReference *Reference `json:"productreference,omitempty"`
	ProductCodeableConcept *CodeableConcept `json:"productcodeableconcept,omitempty"`
	Account []*Reference `json:"account,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
	SupportingInformation []*Reference `json:"supportinginformation,omitempty"`
}

// NewChargeItem creates a new ChargeItem instance
func NewChargeItem() *ChargeItem {
	return &ChargeItem{}
}

// FromJSON populates ChargeItem from JSON data
func (m *ChargeItem) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts ChargeItem to JSON data
func (m *ChargeItem) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of ChargeItem
func (m *ChargeItem) Clone() *ChargeItem {
	if m == nil { return nil }
	return &ChargeItem{
		Id: m.Id.Clone(),
		Meta: m.Meta.Clone(),
		ImplicitRules: m.ImplicitRules.Clone(),
		Language: m.Language.Clone(),
		Text: m.Text.Clone(),
		Contained: cloneSlices(m.Contained),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Identifier: cloneSlices(m.Identifier),
		DefinitionUri: cloneSlices(m.DefinitionUri),
		DefinitionCanonical: cloneSlices(m.DefinitionCanonical),
		Status: m.Status.Clone(),
		PartOf: cloneSlices(m.PartOf),
		Code: m.Code.Clone(),
		Subject: m.Subject.Clone(),
		Context: m.Context.Clone(),
		OccurrenceDateTime: m.OccurrenceDateTime.Clone(),
		OccurrencePeriod: m.OccurrencePeriod.Clone(),
		OccurrenceTiming: m.OccurrenceTiming.Clone(),
		Performer: cloneSlices(m.Performer),
		PerformingOrganization: m.PerformingOrganization.Clone(),
		RequestingOrganization: m.RequestingOrganization.Clone(),
		CostCenter: m.CostCenter.Clone(),
		Quantity: m.Quantity.Clone(),
		Bodysite: cloneSlices(m.Bodysite),
		FactorOverride: m.FactorOverride.Clone(),
		PriceOverride: m.PriceOverride.Clone(),
		OverrideReason: m.OverrideReason.Clone(),
		Enterer: m.Enterer.Clone(),
		EnteredDate: m.EnteredDate.Clone(),
		Reason: cloneSlices(m.Reason),
		Service: cloneSlices(m.Service),
		ProductReference: m.ProductReference.Clone(),
		ProductCodeableConcept: m.ProductCodeableConcept.Clone(),
		Account: cloneSlices(m.Account),
		Note: cloneSlices(m.Note),
		SupportingInformation: cloneSlices(m.SupportingInformation),
	}
}

// Equals checks for equality with another ChargeItem instance
func (m *ChargeItem) Equals(other *ChargeItem) bool {
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
	if !compareSlices(m.DefinitionUri, other.DefinitionUri) { return false }
	if !compareSlices(m.DefinitionCanonical, other.DefinitionCanonical) { return false }
	if !m.Status.Equals(other.Status) { return false }
	if !compareSlices(m.PartOf, other.PartOf) { return false }
	if !m.Code.Equals(other.Code) { return false }
	if !m.Subject.Equals(other.Subject) { return false }
	if !m.Context.Equals(other.Context) { return false }
	if !m.OccurrenceDateTime.Equals(other.OccurrenceDateTime) { return false }
	if !m.OccurrencePeriod.Equals(other.OccurrencePeriod) { return false }
	if !m.OccurrenceTiming.Equals(other.OccurrenceTiming) { return false }
	if !compareSlices(m.Performer, other.Performer) { return false }
	if !m.PerformingOrganization.Equals(other.PerformingOrganization) { return false }
	if !m.RequestingOrganization.Equals(other.RequestingOrganization) { return false }
	if !m.CostCenter.Equals(other.CostCenter) { return false }
	if !m.Quantity.Equals(other.Quantity) { return false }
	if !compareSlices(m.Bodysite, other.Bodysite) { return false }
	if !m.FactorOverride.Equals(other.FactorOverride) { return false }
	if !m.PriceOverride.Equals(other.PriceOverride) { return false }
	if !m.OverrideReason.Equals(other.OverrideReason) { return false }
	if !m.Enterer.Equals(other.Enterer) { return false }
	if !m.EnteredDate.Equals(other.EnteredDate) { return false }
	if !compareSlices(m.Reason, other.Reason) { return false }
	if !compareSlices(m.Service, other.Service) { return false }
	if !m.ProductReference.Equals(other.ProductReference) { return false }
	if !m.ProductCodeableConcept.Equals(other.ProductCodeableConcept) { return false }
	if !compareSlices(m.Account, other.Account) { return false }
	if !compareSlices(m.Note, other.Note) { return false }
	if !compareSlices(m.SupportingInformation, other.SupportingInformation) { return false }
	return true
}

// ChargeItemPerformer
// Indicates who or what performed or participated in the charged service.
type ChargeItemPerformer struct {
	BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Function_ *CodeableConcept `json:"function,omitempty"`
	Actor *Reference `json:"actor,omitempty"`
}

// NewChargeItemPerformer creates a new ChargeItemPerformer instance
func NewChargeItemPerformer() *ChargeItemPerformer {
	return &ChargeItemPerformer{}
}

// FromJSON populates ChargeItemPerformer from JSON data
func (m *ChargeItemPerformer) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts ChargeItemPerformer to JSON data
func (m *ChargeItemPerformer) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of ChargeItemPerformer
func (m *ChargeItemPerformer) Clone() *ChargeItemPerformer {
	if m == nil { return nil }
	return &ChargeItemPerformer{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Function_: m.Function_.Clone(),
		Actor: m.Actor.Clone(),
	}
}

// Equals checks for equality with another ChargeItemPerformer instance
func (m *ChargeItemPerformer) Equals(other *ChargeItemPerformer) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Function_.Equals(other.Function_) { return false }
	if !m.Actor.Equals(other.Actor) { return false }
	return true
}
