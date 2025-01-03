// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"
)

// AllergyIntolerance
// Risk of harmful or undesirable, physiological response which is unique to an individual and associated with exposure to a substance.
type AllergyIntolerance struct {
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
	ClinicalStatus *CodeableConcept `json:"clinicalstatus,omitempty"`
	VerificationStatus *CodeableConcept `json:"verificationstatus,omitempty"`
	Type *AllergyIntoleranceType `json:"type,omitempty"`
	Category []*AllergyIntoleranceCategory `json:"category,omitempty"`
	Criticality *AllergyIntoleranceCriticality `json:"criticality,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	Patient *Reference `json:"patient,omitempty"`
	Encounter *Reference `json:"encounter,omitempty"`
	OnsetDateTime *FhirDateTime `json:"onsetdatetime,omitempty"`
	OnsetAge *Age `json:"onsetage,omitempty"`
	OnsetPeriod *Period `json:"onsetperiod,omitempty"`
	OnsetRange *Range `json:"onsetrange,omitempty"`
	OnsetString *FhirString `json:"onsetstring,omitempty"`
	RecordedDate *FhirDateTime `json:"recordeddate,omitempty"`
	Recorder *Reference `json:"recorder,omitempty"`
	Asserter *Reference `json:"asserter,omitempty"`
	LastOccurrence *FhirDateTime `json:"lastoccurrence,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
	Reaction []*AllergyIntoleranceReaction `json:"reaction,omitempty"`
}

// NewAllergyIntolerance creates a new AllergyIntolerance instance.
func NewAllergyIntolerance() *AllergyIntolerance {
	return &AllergyIntolerance{}
}

// UnmarshalJSON populates AllergyIntolerance from JSON data.
func (m *AllergyIntolerance) UnmarshalJSON(data []byte) error {
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
		ClinicalStatus *CodeableConcept `json:"clinicalstatus,omitempty"`
		VerificationStatus *CodeableConcept `json:"verificationstatus,omitempty"`
		Type *AllergyIntoleranceType `json:"type,omitempty"`
		Category []*AllergyIntoleranceCategory `json:"category,omitempty"`
		Criticality *AllergyIntoleranceCriticality `json:"criticality,omitempty"`
		Code *CodeableConcept `json:"code,omitempty"`
		Patient *Reference `json:"patient,omitempty"`
		Encounter *Reference `json:"encounter,omitempty"`
		OnsetDateTime *FhirDateTime `json:"onsetdatetime,omitempty"`
		OnsetAge *Age `json:"onsetage,omitempty"`
		OnsetPeriod *Period `json:"onsetperiod,omitempty"`
		OnsetRange *Range `json:"onsetrange,omitempty"`
		OnsetString *FhirString `json:"onsetstring,omitempty"`
		RecordedDate *FhirDateTime `json:"recordeddate,omitempty"`
		Recorder *Reference `json:"recorder,omitempty"`
		Asserter *Reference `json:"asserter,omitempty"`
		LastOccurrence *FhirDateTime `json:"lastoccurrence,omitempty"`
		Note []*Annotation `json:"note,omitempty"`
		Reaction []*AllergyIntoleranceReaction `json:"reaction,omitempty"`
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
	m.ClinicalStatus = temp.ClinicalStatus
	m.VerificationStatus = temp.VerificationStatus
	m.Type = temp.Type
	m.Category = temp.Category
	m.Criticality = temp.Criticality
	m.Code = temp.Code
	m.Patient = temp.Patient
	m.Encounter = temp.Encounter
	m.OnsetDateTime = temp.OnsetDateTime
	m.OnsetAge = temp.OnsetAge
	m.OnsetPeriod = temp.OnsetPeriod
	m.OnsetRange = temp.OnsetRange
	m.OnsetString = temp.OnsetString
	m.RecordedDate = temp.RecordedDate
	m.Recorder = temp.Recorder
	m.Asserter = temp.Asserter
	m.LastOccurrence = temp.LastOccurrence
	m.Note = temp.Note
	m.Reaction = temp.Reaction
	return nil
}

// MarshalJSON converts AllergyIntolerance to JSON data.
func (m *AllergyIntolerance) MarshalJSON() ([]byte, error) {
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
		ClinicalStatus *CodeableConcept `json:"clinicalstatus,omitempty"`
		VerificationStatus *CodeableConcept `json:"verificationstatus,omitempty"`
		Type *AllergyIntoleranceType `json:"type,omitempty"`
		Category []*AllergyIntoleranceCategory `json:"category,omitempty"`
		Criticality *AllergyIntoleranceCriticality `json:"criticality,omitempty"`
		Code *CodeableConcept `json:"code,omitempty"`
		Patient *Reference `json:"patient,omitempty"`
		Encounter *Reference `json:"encounter,omitempty"`
		OnsetDateTime interface{} `json:"onsetdatetime,omitempty"`
		OnsetDateTimeElement map[string]interface{} `json:"_onsetdatetime,omitempty"`
		OnsetAge *Age `json:"onsetage,omitempty"`
		OnsetPeriod *Period `json:"onsetperiod,omitempty"`
		OnsetRange *Range `json:"onsetrange,omitempty"`
		OnsetString interface{} `json:"onsetstring,omitempty"`
		OnsetStringElement map[string]interface{} `json:"_onsetstring,omitempty"`
		RecordedDate interface{} `json:"recordeddate,omitempty"`
		RecordedDateElement map[string]interface{} `json:"_recordeddate,omitempty"`
		Recorder *Reference `json:"recorder,omitempty"`
		Asserter *Reference `json:"asserter,omitempty"`
		LastOccurrence interface{} `json:"lastoccurrence,omitempty"`
		LastOccurrenceElement map[string]interface{} `json:"_lastoccurrence,omitempty"`
		Note []*Annotation `json:"note,omitempty"`
		Reaction []*AllergyIntoleranceReaction `json:"reaction,omitempty"`
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
	output.ClinicalStatus = m.ClinicalStatus
	output.VerificationStatus = m.VerificationStatus
	output.Type = m.Type
	output.Category = m.Category
	output.Criticality = m.Criticality
	output.Code = m.Code
	output.Patient = m.Patient
	output.Encounter = m.Encounter
	if m.OnsetDateTime != nil && m.OnsetDateTime.Value != nil {
		output.OnsetDateTime = m.OnsetDateTime.Value
		if m.OnsetDateTime.Element != nil {
			output.OnsetDateTimeElement = toMapOrNil(m.OnsetDateTime.Element.MarshalJSON())
		}
	}
	output.OnsetAge = m.OnsetAge
	output.OnsetPeriod = m.OnsetPeriod
	output.OnsetRange = m.OnsetRange
	if m.OnsetString != nil && m.OnsetString.Value != nil {
		output.OnsetString = m.OnsetString.Value
		if m.OnsetString.Element != nil {
			output.OnsetStringElement = toMapOrNil(m.OnsetString.Element.MarshalJSON())
		}
	}
	if m.RecordedDate != nil && m.RecordedDate.Value != nil {
		output.RecordedDate = m.RecordedDate.Value
		if m.RecordedDate.Element != nil {
			output.RecordedDateElement = toMapOrNil(m.RecordedDate.Element.MarshalJSON())
		}
	}
	output.Recorder = m.Recorder
	output.Asserter = m.Asserter
	if m.LastOccurrence != nil && m.LastOccurrence.Value != nil {
		output.LastOccurrence = m.LastOccurrence.Value
		if m.LastOccurrence.Element != nil {
			output.LastOccurrenceElement = toMapOrNil(m.LastOccurrence.Element.MarshalJSON())
		}
	}
	output.Note = m.Note
	output.Reaction = m.Reaction
	return json.Marshal(output)
}

// Clone creates a deep copy of AllergyIntolerance.
func (m *AllergyIntolerance) Clone() *AllergyIntolerance {
	if m == nil { return nil }
	return &AllergyIntolerance{
		Id: m.Id.Clone(),
		Meta: m.Meta.Clone(),
		ImplicitRules: m.ImplicitRules.Clone(),
		Language: m.Language.Clone(),
		Text: m.Text.Clone(),
		Contained: cloneSlices(m.Contained),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Identifier: cloneSlices(m.Identifier),
		ClinicalStatus: m.ClinicalStatus.Clone(),
		VerificationStatus: m.VerificationStatus.Clone(),
		Type: m.Type.Clone(),
		Category: cloneSlices(m.Category),
		Criticality: m.Criticality.Clone(),
		Code: m.Code.Clone(),
		Patient: m.Patient.Clone(),
		Encounter: m.Encounter.Clone(),
		OnsetDateTime: m.OnsetDateTime.Clone(),
		OnsetAge: m.OnsetAge.Clone(),
		OnsetPeriod: m.OnsetPeriod.Clone(),
		OnsetRange: m.OnsetRange.Clone(),
		OnsetString: m.OnsetString.Clone(),
		RecordedDate: m.RecordedDate.Clone(),
		Recorder: m.Recorder.Clone(),
		Asserter: m.Asserter.Clone(),
		LastOccurrence: m.LastOccurrence.Clone(),
		Note: cloneSlices(m.Note),
		Reaction: cloneSlices(m.Reaction),
	}
}

// Equals checks equality between two AllergyIntolerance instances.
func (m *AllergyIntolerance) Equals(other *AllergyIntolerance) bool {
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
	if !m.ClinicalStatus.Equals(other.ClinicalStatus) { return false }
	if !m.VerificationStatus.Equals(other.VerificationStatus) { return false }
	if !m.Type.Equals(other.Type) { return false }
	if !compareSlices(m.Category, other.Category) { return false }
	if !m.Criticality.Equals(other.Criticality) { return false }
	if !m.Code.Equals(other.Code) { return false }
	if !m.Patient.Equals(other.Patient) { return false }
	if !m.Encounter.Equals(other.Encounter) { return false }
	if !m.OnsetDateTime.Equals(other.OnsetDateTime) { return false }
	if !m.OnsetAge.Equals(other.OnsetAge) { return false }
	if !m.OnsetPeriod.Equals(other.OnsetPeriod) { return false }
	if !m.OnsetRange.Equals(other.OnsetRange) { return false }
	if !m.OnsetString.Equals(other.OnsetString) { return false }
	if !m.RecordedDate.Equals(other.RecordedDate) { return false }
	if !m.Recorder.Equals(other.Recorder) { return false }
	if !m.Asserter.Equals(other.Asserter) { return false }
	if !m.LastOccurrence.Equals(other.LastOccurrence) { return false }
	if !compareSlices(m.Note, other.Note) { return false }
	if !compareSlices(m.Reaction, other.Reaction) { return false }
	return true
}

// AllergyIntoleranceReaction
// Details about each adverse reaction event linked to exposure to the identified substance.
type AllergyIntoleranceReaction struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Substance *CodeableConcept `json:"substance,omitempty"`
	Manifestation []*CodeableConcept `json:"manifestation,omitempty"`
	Description *FhirString `json:"description,omitempty"`
	Onset *FhirDateTime `json:"onset,omitempty"`
	Severity *AllergyIntoleranceSeverity `json:"severity,omitempty"`
	ExposureRoute *CodeableConcept `json:"exposureroute,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
}

// NewAllergyIntoleranceReaction creates a new AllergyIntoleranceReaction instance.
func NewAllergyIntoleranceReaction() *AllergyIntoleranceReaction {
	return &AllergyIntoleranceReaction{}
}

// UnmarshalJSON populates AllergyIntoleranceReaction from JSON data.
func (m *AllergyIntoleranceReaction) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Substance *CodeableConcept `json:"substance,omitempty"`
		Manifestation []*CodeableConcept `json:"manifestation,omitempty"`
		Description *FhirString `json:"description,omitempty"`
		Onset *FhirDateTime `json:"onset,omitempty"`
		Severity *AllergyIntoleranceSeverity `json:"severity,omitempty"`
		ExposureRoute *CodeableConcept `json:"exposureroute,omitempty"`
		Note []*Annotation `json:"note,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Substance = temp.Substance
	m.Manifestation = temp.Manifestation
	m.Description = temp.Description
	m.Onset = temp.Onset
	m.Severity = temp.Severity
	m.ExposureRoute = temp.ExposureRoute
	m.Note = temp.Note
	return nil
}

// MarshalJSON converts AllergyIntoleranceReaction to JSON data.
func (m *AllergyIntoleranceReaction) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Substance *CodeableConcept `json:"substance,omitempty"`
		Manifestation []*CodeableConcept `json:"manifestation,omitempty"`
		Description interface{} `json:"description,omitempty"`
		DescriptionElement map[string]interface{} `json:"_description,omitempty"`
		Onset interface{} `json:"onset,omitempty"`
		OnsetElement map[string]interface{} `json:"_onset,omitempty"`
		Severity *AllergyIntoleranceSeverity `json:"severity,omitempty"`
		ExposureRoute *CodeableConcept `json:"exposureroute,omitempty"`
		Note []*Annotation `json:"note,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.Substance = m.Substance
	output.Manifestation = m.Manifestation
	if m.Description != nil && m.Description.Value != nil {
		output.Description = m.Description.Value
		if m.Description.Element != nil {
			output.DescriptionElement = toMapOrNil(m.Description.Element.MarshalJSON())
		}
	}
	if m.Onset != nil && m.Onset.Value != nil {
		output.Onset = m.Onset.Value
		if m.Onset.Element != nil {
			output.OnsetElement = toMapOrNil(m.Onset.Element.MarshalJSON())
		}
	}
	output.Severity = m.Severity
	output.ExposureRoute = m.ExposureRoute
	output.Note = m.Note
	return json.Marshal(output)
}

// Clone creates a deep copy of AllergyIntoleranceReaction.
func (m *AllergyIntoleranceReaction) Clone() *AllergyIntoleranceReaction {
	if m == nil { return nil }
	return &AllergyIntoleranceReaction{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Substance: m.Substance.Clone(),
		Manifestation: cloneSlices(m.Manifestation),
		Description: m.Description.Clone(),
		Onset: m.Onset.Clone(),
		Severity: m.Severity.Clone(),
		ExposureRoute: m.ExposureRoute.Clone(),
		Note: cloneSlices(m.Note),
	}
}

// Equals checks equality between two AllergyIntoleranceReaction instances.
func (m *AllergyIntoleranceReaction) Equals(other *AllergyIntoleranceReaction) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Substance.Equals(other.Substance) { return false }
	if !compareSlices(m.Manifestation, other.Manifestation) { return false }
	if !m.Description.Equals(other.Description) { return false }
	if !m.Onset.Equals(other.Onset) { return false }
	if !m.Severity.Equals(other.Severity) { return false }
	if !m.ExposureRoute.Equals(other.ExposureRoute) { return false }
	if !compareSlices(m.Note, other.Note) { return false }
	return true
}

