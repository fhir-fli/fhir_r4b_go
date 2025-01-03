// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"
)

// Appointment
// A booking of a healthcare event among patient(s), practitioner(s), related person(s) and/or device(s) for a specific date/time. This may result in one or more Encounter(s).
type Appointment struct {
	extends DomainResource
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

// NewAppointment creates a new Appointment instance.
func NewAppointment() *Appointment {
	return &Appointment{}
}

// UnmarshalJSON populates Appointment from JSON data.
func (m *Appointment) UnmarshalJSON(data []byte) error {
	temp := struct {
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
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Meta = temp.Meta
	m.ImplicitRules = temp.ImplicitRules
	m.Language = temp.Language
	m.Text = temp.Text
	m.Contained = temp.Contained
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Identifier = temp.Identifier
	m.Status = temp.Status
	m.CancelationReason = temp.CancelationReason
	m.ServiceCategory = temp.ServiceCategory
	m.ServiceType = temp.ServiceType
	m.Specialty = temp.Specialty
	m.AppointmentType = temp.AppointmentType
	m.ReasonCode = temp.ReasonCode
	m.ReasonReference = temp.ReasonReference
	m.Priority = temp.Priority
	m.Description = temp.Description
	m.SupportingInformation = temp.SupportingInformation
	m.Start = temp.Start
	m.End = temp.End
	m.MinutesDuration = temp.MinutesDuration
	m.Slot = temp.Slot
	m.Created = temp.Created
	m.Comment = temp.Comment
	m.PatientInstruction = temp.PatientInstruction
	m.BasedOn = temp.BasedOn
	m.Participant = temp.Participant
	m.RequestedPeriod = temp.RequestedPeriod
	return nil
}

// MarshalJSON converts Appointment to JSON data.
func (m *Appointment) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Meta *FhirMeta `json:"meta,omitempty"`
		ImplicitRules interface{} `json:"implicitrules,omitempty"`
		ImplicitRulesElement map[string]interface{} `json:"_implicitrules,omitempty"`
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
		Priority interface{} `json:"priority,omitempty"`
		PriorityElement map[string]interface{} `json:"_priority,omitempty"`
		Description interface{} `json:"description,omitempty"`
		DescriptionElement map[string]interface{} `json:"_description,omitempty"`
		SupportingInformation []*Reference `json:"supportinginformation,omitempty"`
		Start interface{} `json:"start,omitempty"`
		StartElement map[string]interface{} `json:"_start,omitempty"`
		End interface{} `json:"end,omitempty"`
		EndElement map[string]interface{} `json:"_end,omitempty"`
		MinutesDuration interface{} `json:"minutesduration,omitempty"`
		MinutesDurationElement map[string]interface{} `json:"_minutesduration,omitempty"`
		Slot []*Reference `json:"slot,omitempty"`
		Created interface{} `json:"created,omitempty"`
		CreatedElement map[string]interface{} `json:"_created,omitempty"`
		Comment interface{} `json:"comment,omitempty"`
		CommentElement map[string]interface{} `json:"_comment,omitempty"`
		PatientInstruction interface{} `json:"patientinstruction,omitempty"`
		PatientInstructionElement map[string]interface{} `json:"_patientinstruction,omitempty"`
		BasedOn []*Reference `json:"basedon,omitempty"`
		Participant []*AppointmentParticipant `json:"participant,omitempty"`
		RequestedPeriod []*Period `json:"requestedperiod,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Meta = m.Meta
	if m.ImplicitRules != nil && m.ImplicitRules.Value != nil {
		output.ImplicitRules = m.ImplicitRules.Value
		if m.ImplicitRules.Element != nil {
			output.ImplicitRulesElement = toMapOrNil(m.ImplicitRules.Element.MarshalJSON())
		}
	}
	output.Language = m.Language
	output.Text = m.Text
	output.Contained = m.Contained
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.Identifier = m.Identifier
	output.Status = m.Status
	output.CancelationReason = m.CancelationReason
	output.ServiceCategory = m.ServiceCategory
	output.ServiceType = m.ServiceType
	output.Specialty = m.Specialty
	output.AppointmentType = m.AppointmentType
	output.ReasonCode = m.ReasonCode
	output.ReasonReference = m.ReasonReference
	if m.Priority != nil && m.Priority.Value != nil {
		output.Priority = m.Priority.Value
		if m.Priority.Element != nil {
			output.PriorityElement = toMapOrNil(m.Priority.Element.MarshalJSON())
		}
	}
	if m.Description != nil && m.Description.Value != nil {
		output.Description = m.Description.Value
		if m.Description.Element != nil {
			output.DescriptionElement = toMapOrNil(m.Description.Element.MarshalJSON())
		}
	}
	output.SupportingInformation = m.SupportingInformation
	if m.Start != nil && m.Start.Value != nil {
		output.Start = m.Start.Value
		if m.Start.Element != nil {
			output.StartElement = toMapOrNil(m.Start.Element.MarshalJSON())
		}
	}
	if m.End != nil && m.End.Value != nil {
		output.End = m.End.Value
		if m.End.Element != nil {
			output.EndElement = toMapOrNil(m.End.Element.MarshalJSON())
		}
	}
	if m.MinutesDuration != nil && m.MinutesDuration.Value != nil {
		output.MinutesDuration = m.MinutesDuration.Value
		if m.MinutesDuration.Element != nil {
			output.MinutesDurationElement = toMapOrNil(m.MinutesDuration.Element.MarshalJSON())
		}
	}
	output.Slot = m.Slot
	if m.Created != nil && m.Created.Value != nil {
		output.Created = m.Created.Value
		if m.Created.Element != nil {
			output.CreatedElement = toMapOrNil(m.Created.Element.MarshalJSON())
		}
	}
	if m.Comment != nil && m.Comment.Value != nil {
		output.Comment = m.Comment.Value
		if m.Comment.Element != nil {
			output.CommentElement = toMapOrNil(m.Comment.Element.MarshalJSON())
		}
	}
	if m.PatientInstruction != nil && m.PatientInstruction.Value != nil {
		output.PatientInstruction = m.PatientInstruction.Value
		if m.PatientInstruction.Element != nil {
			output.PatientInstructionElement = toMapOrNil(m.PatientInstruction.Element.MarshalJSON())
		}
	}
	output.BasedOn = m.BasedOn
	output.Participant = m.Participant
	output.RequestedPeriod = m.RequestedPeriod
	return json.Marshal(output)
}

// Clone creates a deep copy of Appointment.
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

// Equals checks equality between two Appointment instances.
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
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Type []*CodeableConcept `json:"type,omitempty"`
	Actor *Reference `json:"actor,omitempty"`
	Required_ *ParticipantRequired `json:"required,omitempty"`
	Status *ParticipationStatus `json:"status,omitempty"`
	Period *Period `json:"period,omitempty"`
}

// NewAppointmentParticipant creates a new AppointmentParticipant instance.
func NewAppointmentParticipant() *AppointmentParticipant {
	return &AppointmentParticipant{}
}

// UnmarshalJSON populates AppointmentParticipant from JSON data.
func (m *AppointmentParticipant) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Type []*CodeableConcept `json:"type,omitempty"`
		Actor *Reference `json:"actor,omitempty"`
		Required_ *ParticipantRequired `json:"required,omitempty"`
		Status *ParticipationStatus `json:"status,omitempty"`
		Period *Period `json:"period,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Type = temp.Type
	m.Actor = temp.Actor
	m.Required_ = temp.Required_
	m.Status = temp.Status
	m.Period = temp.Period
	return nil
}

// MarshalJSON converts AppointmentParticipant to JSON data.
func (m *AppointmentParticipant) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Type []*CodeableConcept `json:"type,omitempty"`
		Actor *Reference `json:"actor,omitempty"`
		Required_ *ParticipantRequired `json:"required,omitempty"`
		Status *ParticipationStatus `json:"status,omitempty"`
		Period *Period `json:"period,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.Type = m.Type
	output.Actor = m.Actor
	output.Required_ = m.Required_
	output.Status = m.Status
	output.Period = m.Period
	return json.Marshal(output)
}

// Clone creates a deep copy of AppointmentParticipant.
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

// Equals checks equality between two AppointmentParticipant instances.
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

