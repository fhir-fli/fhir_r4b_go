// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json")

// Appointment
// A booking of a healthcare event among patient(s), practitioner(s), related person(s) and/or device(s) for a specific date/time. This may result in one or more Encounter(s).
type Appointment struct {
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
	Status *AppointmentStatus `json:"status,omitempty"`
	CancelationReason *CodeableConcept `json:"cancelationreason,omitempty"`
	ServiceCategory []*CodeableConcept `json:"servicecategory,omitempty"`
	ServiceType []*CodeableConcept `json:"servicetype,omitempty"`
	Specialty []*CodeableConcept `json:"specialty,omitempty"`
	AppointmentType *CodeableConcept `json:"appointmenttype,omitempty"`
	ReasonCode []*CodeableConcept `json:"reasoncode,omitempty"`
	ReasonReference []*Reference `json:"reasonreference,omitempty"`
	Priority *FhirUnsignedInt `json:"priority,omitempty"`
	Description *FhirString `json:"description,omitempty"`
	SupportingInformation []*Reference `json:"supportinginformation,omitempty"`
	Start *FhirInstant `json:"start,omitempty"`
	End *FhirInstant `json:"end,omitempty"`
	MinutesDuration *FhirPositiveInt `json:"minutesduration,omitempty"`
	Slot []*Reference `json:"slot,omitempty"`
	Created *FhirDateTime `json:"created,omitempty"`
	Comment *FhirString `json:"comment,omitempty"`
	PatientInstruction *FhirString `json:"patientinstruction,omitempty"`
	BasedOn []*Reference `json:"basedon,omitempty"`
	Participant []*AppointmentParticipant `json:"participant,omitempty"`
	RequestedPeriod []*Period `json:"requestedperiod,omitempty"`
}

// NewAppointment creates a new Appointment instance
func NewAppointment() *Appointment {
	return &Appointment{}
}

// FromJSON populates Appointment from JSON data
func (m *Appointment) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts Appointment to JSON data
func (m *Appointment) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of Appointment
func (m *Appointment) Clone() *Appointment {
	if m == nil { return nil }
	return &Appointment{
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
		CancelationReason: m.CancelationReason.Clone(),
		ServiceCategory: cloneSlices(m.ServiceCategory),
		ServiceType: cloneSlices(m.ServiceType),
		Specialty: cloneSlices(m.Specialty),
		AppointmentType: m.AppointmentType.Clone(),
		ReasonCode: cloneSlices(m.ReasonCode),
		ReasonReference: cloneSlices(m.ReasonReference),
		Priority: m.Priority.Clone(),
		Description: m.Description.Clone(),
		SupportingInformation: cloneSlices(m.SupportingInformation),
		Start: m.Start.Clone(),
		End: m.End.Clone(),
		MinutesDuration: m.MinutesDuration.Clone(),
		Slot: cloneSlices(m.Slot),
		Created: m.Created.Clone(),
		Comment: m.Comment.Clone(),
		PatientInstruction: m.PatientInstruction.Clone(),
		BasedOn: cloneSlices(m.BasedOn),
		Participant: cloneSlices(m.Participant),
		RequestedPeriod: cloneSlices(m.RequestedPeriod),
	}
}

// Equals checks for equality with another Appointment instance
func (m *Appointment) Equals(other *Appointment) bool {
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
	if !m.CancelationReason.Equals(other.CancelationReason) { return false }
	if !compareSlices(m.ServiceCategory, other.ServiceCategory) { return false }
	if !compareSlices(m.ServiceType, other.ServiceType) { return false }
	if !compareSlices(m.Specialty, other.Specialty) { return false }
	if !m.AppointmentType.Equals(other.AppointmentType) { return false }
	if !compareSlices(m.ReasonCode, other.ReasonCode) { return false }
	if !compareSlices(m.ReasonReference, other.ReasonReference) { return false }
	if !m.Priority.Equals(other.Priority) { return false }
	if !m.Description.Equals(other.Description) { return false }
	if !compareSlices(m.SupportingInformation, other.SupportingInformation) { return false }
	if !m.Start.Equals(other.Start) { return false }
	if !m.End.Equals(other.End) { return false }
	if !m.MinutesDuration.Equals(other.MinutesDuration) { return false }
	if !compareSlices(m.Slot, other.Slot) { return false }
	if !m.Created.Equals(other.Created) { return false }
	if !m.Comment.Equals(other.Comment) { return false }
	if !m.PatientInstruction.Equals(other.PatientInstruction) { return false }
	if !compareSlices(m.BasedOn, other.BasedOn) { return false }
	if !compareSlices(m.Participant, other.Participant) { return false }
	if !compareSlices(m.RequestedPeriod, other.RequestedPeriod) { return false }
	return true
}

// AppointmentParticipant
// List of participants involved in the appointment.
type AppointmentParticipant struct {
	BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Type []*CodeableConcept `json:"type,omitempty"`
	Actor *Reference `json:"actor,omitempty"`
	Required_ *ParticipantRequired `json:"required,omitempty"`
	Status *ParticipationStatus `json:"status,omitempty"`
	Period *Period `json:"period,omitempty"`
}

// NewAppointmentParticipant creates a new AppointmentParticipant instance
func NewAppointmentParticipant() *AppointmentParticipant {
	return &AppointmentParticipant{}
}

// FromJSON populates AppointmentParticipant from JSON data
func (m *AppointmentParticipant) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts AppointmentParticipant to JSON data
func (m *AppointmentParticipant) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of AppointmentParticipant
func (m *AppointmentParticipant) Clone() *AppointmentParticipant {
	if m == nil { return nil }
	return &AppointmentParticipant{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Type: cloneSlices(m.Type),
		Actor: m.Actor.Clone(),
		Required_: m.Required_.Clone(),
		Status: m.Status.Clone(),
		Period: m.Period.Clone(),
	}
}

// Equals checks for equality with another AppointmentParticipant instance
func (m *AppointmentParticipant) Equals(other *AppointmentParticipant) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !compareSlices(m.Type, other.Type) { return false }
	if !m.Actor.Equals(other.Actor) { return false }
	if !m.Required_.Equals(other.Required_) { return false }
	if !m.Status.Equals(other.Status) { return false }
	if !m.Period.Equals(other.Period) { return false }
	return true
}
