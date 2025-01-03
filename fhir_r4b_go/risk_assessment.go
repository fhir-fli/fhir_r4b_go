// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"
)

// RiskAssessment
// An assessment of the likely outcome(s) for a patient or other subject as well as the likelihood of each outcome.
type RiskAssessment struct {
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
	BasedOn *Reference `json:"basedon,omitempty"`
	Parent *Reference `json:"parent,omitempty"`
	Status *ObservationStatus `json:"status,omitempty"`
	Method *CodeableConcept `json:"method,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	Subject *Reference `json:"subject,omitempty"`
	Encounter *Reference `json:"encounter,omitempty"`
	OccurrenceDateTime *FhirDateTime `json:"occurrencedatetime,omitempty"`
	OccurrencePeriod *Period `json:"occurrenceperiod,omitempty"`
	Condition *Reference `json:"condition,omitempty"`
	Performer *Reference `json:"performer,omitempty"`
	ReasonCode []*CodeableConcept `json:"reasoncode,omitempty"`
	ReasonReference []*Reference `json:"reasonreference,omitempty"`
	Basis []*Reference `json:"basis,omitempty"`
	Prediction []*RiskAssessmentPrediction `json:"prediction,omitempty"`
	Mitigation *FhirString `json:"mitigation,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
}

// NewRiskAssessment creates a new RiskAssessment instance.
func NewRiskAssessment() *RiskAssessment {
	return &RiskAssessment{}
}

// UnmarshalJSON populates RiskAssessment from JSON data.
func (m *RiskAssessment) UnmarshalJSON(data []byte) error {
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
		BasedOn *Reference `json:"basedon,omitempty"`
		Parent *Reference `json:"parent,omitempty"`
		Status *ObservationStatus `json:"status,omitempty"`
		Method *CodeableConcept `json:"method,omitempty"`
		Code *CodeableConcept `json:"code,omitempty"`
		Subject *Reference `json:"subject,omitempty"`
		Encounter *Reference `json:"encounter,omitempty"`
		OccurrenceDateTime *FhirDateTime `json:"occurrencedatetime,omitempty"`
		OccurrencePeriod *Period `json:"occurrenceperiod,omitempty"`
		Condition *Reference `json:"condition,omitempty"`
		Performer *Reference `json:"performer,omitempty"`
		ReasonCode []*CodeableConcept `json:"reasoncode,omitempty"`
		ReasonReference []*Reference `json:"reasonreference,omitempty"`
		Basis []*Reference `json:"basis,omitempty"`
		Prediction []*RiskAssessmentPrediction `json:"prediction,omitempty"`
		Mitigation *FhirString `json:"mitigation,omitempty"`
		Note []*Annotation `json:"note,omitempty"`
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
	m.BasedOn = temp.BasedOn
	m.Parent = temp.Parent
	m.Status = temp.Status
	m.Method = temp.Method
	m.Code = temp.Code
	m.Subject = temp.Subject
	m.Encounter = temp.Encounter
	m.OccurrenceDateTime = temp.OccurrenceDateTime
	m.OccurrencePeriod = temp.OccurrencePeriod
	m.Condition = temp.Condition
	m.Performer = temp.Performer
	m.ReasonCode = temp.ReasonCode
	m.ReasonReference = temp.ReasonReference
	m.Basis = temp.Basis
	m.Prediction = temp.Prediction
	m.Mitigation = temp.Mitigation
	m.Note = temp.Note
	return nil
}

// MarshalJSON converts RiskAssessment to JSON data.
func (m *RiskAssessment) MarshalJSON() ([]byte, error) {
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
		BasedOn *Reference `json:"basedon,omitempty"`
		Parent *Reference `json:"parent,omitempty"`
		Status *ObservationStatus `json:"status,omitempty"`
		Method *CodeableConcept `json:"method,omitempty"`
		Code *CodeableConcept `json:"code,omitempty"`
		Subject *Reference `json:"subject,omitempty"`
		Encounter *Reference `json:"encounter,omitempty"`
		OccurrenceDateTime interface{} `json:"occurrencedatetime,omitempty"`
		OccurrenceDateTimeElement map[string]interface{} `json:"_occurrencedatetime,omitempty"`
		OccurrencePeriod *Period `json:"occurrenceperiod,omitempty"`
		Condition *Reference `json:"condition,omitempty"`
		Performer *Reference `json:"performer,omitempty"`
		ReasonCode []*CodeableConcept `json:"reasoncode,omitempty"`
		ReasonReference []*Reference `json:"reasonreference,omitempty"`
		Basis []*Reference `json:"basis,omitempty"`
		Prediction []*RiskAssessmentPrediction `json:"prediction,omitempty"`
		Mitigation interface{} `json:"mitigation,omitempty"`
		MitigationElement map[string]interface{} `json:"_mitigation,omitempty"`
		Note []*Annotation `json:"note,omitempty"`
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
	output.BasedOn = m.BasedOn
	output.Parent = m.Parent
	output.Status = m.Status
	output.Method = m.Method
	output.Code = m.Code
	output.Subject = m.Subject
	output.Encounter = m.Encounter
	if m.OccurrenceDateTime != nil && m.OccurrenceDateTime.Value != nil {
		output.OccurrenceDateTime = m.OccurrenceDateTime.Value
		if m.OccurrenceDateTime.Element != nil {
			output.OccurrenceDateTimeElement = toMapOrNil(m.OccurrenceDateTime.Element.MarshalJSON())
		}
	}
	output.OccurrencePeriod = m.OccurrencePeriod
	output.Condition = m.Condition
	output.Performer = m.Performer
	output.ReasonCode = m.ReasonCode
	output.ReasonReference = m.ReasonReference
	output.Basis = m.Basis
	output.Prediction = m.Prediction
	if m.Mitigation != nil && m.Mitigation.Value != nil {
		output.Mitigation = m.Mitigation.Value
		if m.Mitigation.Element != nil {
			output.MitigationElement = toMapOrNil(m.Mitigation.Element.MarshalJSON())
		}
	}
	output.Note = m.Note
	return json.Marshal(output)
}

// Clone creates a deep copy of RiskAssessment.
func (m *RiskAssessment) Clone() *RiskAssessment {
	if m == nil { return nil }
	return &RiskAssessment{
		Id: m.Id.Clone(),
		Meta: m.Meta.Clone(),
		ImplicitRules: m.ImplicitRules.Clone(),
		Language: m.Language.Clone(),
		Text: m.Text.Clone(),
		Contained: cloneSlices(m.Contained),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Identifier: cloneSlices(m.Identifier),
		BasedOn: m.BasedOn.Clone(),
		Parent: m.Parent.Clone(),
		Status: m.Status.Clone(),
		Method: m.Method.Clone(),
		Code: m.Code.Clone(),
		Subject: m.Subject.Clone(),
		Encounter: m.Encounter.Clone(),
		OccurrenceDateTime: m.OccurrenceDateTime.Clone(),
		OccurrencePeriod: m.OccurrencePeriod.Clone(),
		Condition: m.Condition.Clone(),
		Performer: m.Performer.Clone(),
		ReasonCode: cloneSlices(m.ReasonCode),
		ReasonReference: cloneSlices(m.ReasonReference),
		Basis: cloneSlices(m.Basis),
		Prediction: cloneSlices(m.Prediction),
		Mitigation: m.Mitigation.Clone(),
		Note: cloneSlices(m.Note),
	}
}

// Equals checks equality between two RiskAssessment instances.
func (m *RiskAssessment) Equals(other *RiskAssessment) bool {
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
	if !m.BasedOn.Equals(other.BasedOn) { return false }
	if !m.Parent.Equals(other.Parent) { return false }
	if !m.Status.Equals(other.Status) { return false }
	if !m.Method.Equals(other.Method) { return false }
	if !m.Code.Equals(other.Code) { return false }
	if !m.Subject.Equals(other.Subject) { return false }
	if !m.Encounter.Equals(other.Encounter) { return false }
	if !m.OccurrenceDateTime.Equals(other.OccurrenceDateTime) { return false }
	if !m.OccurrencePeriod.Equals(other.OccurrencePeriod) { return false }
	if !m.Condition.Equals(other.Condition) { return false }
	if !m.Performer.Equals(other.Performer) { return false }
	if !compareSlices(m.ReasonCode, other.ReasonCode) { return false }
	if !compareSlices(m.ReasonReference, other.ReasonReference) { return false }
	if !compareSlices(m.Basis, other.Basis) { return false }
	if !compareSlices(m.Prediction, other.Prediction) { return false }
	if !m.Mitigation.Equals(other.Mitigation) { return false }
	if !compareSlices(m.Note, other.Note) { return false }
	return true
}

// RiskAssessmentPrediction
// Describes the expected outcome for the subject.
type RiskAssessmentPrediction struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Outcome *CodeableConcept `json:"outcome,omitempty"`
	ProbabilityDecimal *FhirDecimal `json:"probabilitydecimal,omitempty"`
	ProbabilityRange *Range `json:"probabilityrange,omitempty"`
	QualitativeRisk *CodeableConcept `json:"qualitativerisk,omitempty"`
	RelativeRisk *FhirDecimal `json:"relativerisk,omitempty"`
	WhenPeriod *Period `json:"whenperiod,omitempty"`
	WhenRange *Range `json:"whenrange,omitempty"`
	Rationale *FhirString `json:"rationale,omitempty"`
}

// NewRiskAssessmentPrediction creates a new RiskAssessmentPrediction instance.
func NewRiskAssessmentPrediction() *RiskAssessmentPrediction {
	return &RiskAssessmentPrediction{}
}

// UnmarshalJSON populates RiskAssessmentPrediction from JSON data.
func (m *RiskAssessmentPrediction) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Outcome *CodeableConcept `json:"outcome,omitempty"`
		ProbabilityDecimal *FhirDecimal `json:"probabilitydecimal,omitempty"`
		ProbabilityRange *Range `json:"probabilityrange,omitempty"`
		QualitativeRisk *CodeableConcept `json:"qualitativerisk,omitempty"`
		RelativeRisk *FhirDecimal `json:"relativerisk,omitempty"`
		WhenPeriod *Period `json:"whenperiod,omitempty"`
		WhenRange *Range `json:"whenrange,omitempty"`
		Rationale *FhirString `json:"rationale,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Outcome = temp.Outcome
	m.ProbabilityDecimal = temp.ProbabilityDecimal
	m.ProbabilityRange = temp.ProbabilityRange
	m.QualitativeRisk = temp.QualitativeRisk
	m.RelativeRisk = temp.RelativeRisk
	m.WhenPeriod = temp.WhenPeriod
	m.WhenRange = temp.WhenRange
	m.Rationale = temp.Rationale
	return nil
}

// MarshalJSON converts RiskAssessmentPrediction to JSON data.
func (m *RiskAssessmentPrediction) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Outcome *CodeableConcept `json:"outcome,omitempty"`
		ProbabilityDecimal interface{} `json:"probabilitydecimal,omitempty"`
		ProbabilityDecimalElement map[string]interface{} `json:"_probabilitydecimal,omitempty"`
		ProbabilityRange *Range `json:"probabilityrange,omitempty"`
		QualitativeRisk *CodeableConcept `json:"qualitativerisk,omitempty"`
		RelativeRisk interface{} `json:"relativerisk,omitempty"`
		RelativeRiskElement map[string]interface{} `json:"_relativerisk,omitempty"`
		WhenPeriod *Period `json:"whenperiod,omitempty"`
		WhenRange *Range `json:"whenrange,omitempty"`
		Rationale interface{} `json:"rationale,omitempty"`
		RationaleElement map[string]interface{} `json:"_rationale,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.Outcome = m.Outcome
	if m.ProbabilityDecimal != nil && m.ProbabilityDecimal.Value != nil {
		output.ProbabilityDecimal = m.ProbabilityDecimal.Value
		if m.ProbabilityDecimal.Element != nil {
			output.ProbabilityDecimalElement = toMapOrNil(m.ProbabilityDecimal.Element.MarshalJSON())
		}
	}
	output.ProbabilityRange = m.ProbabilityRange
	output.QualitativeRisk = m.QualitativeRisk
	if m.RelativeRisk != nil && m.RelativeRisk.Value != nil {
		output.RelativeRisk = m.RelativeRisk.Value
		if m.RelativeRisk.Element != nil {
			output.RelativeRiskElement = toMapOrNil(m.RelativeRisk.Element.MarshalJSON())
		}
	}
	output.WhenPeriod = m.WhenPeriod
	output.WhenRange = m.WhenRange
	if m.Rationale != nil && m.Rationale.Value != nil {
		output.Rationale = m.Rationale.Value
		if m.Rationale.Element != nil {
			output.RationaleElement = toMapOrNil(m.Rationale.Element.MarshalJSON())
		}
	}
	return json.Marshal(output)
}

// Clone creates a deep copy of RiskAssessmentPrediction.
func (m *RiskAssessmentPrediction) Clone() *RiskAssessmentPrediction {
	if m == nil { return nil }
	return &RiskAssessmentPrediction{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Outcome: m.Outcome.Clone(),
		ProbabilityDecimal: m.ProbabilityDecimal.Clone(),
		ProbabilityRange: m.ProbabilityRange.Clone(),
		QualitativeRisk: m.QualitativeRisk.Clone(),
		RelativeRisk: m.RelativeRisk.Clone(),
		WhenPeriod: m.WhenPeriod.Clone(),
		WhenRange: m.WhenRange.Clone(),
		Rationale: m.Rationale.Clone(),
	}
}

// Equals checks equality between two RiskAssessmentPrediction instances.
func (m *RiskAssessmentPrediction) Equals(other *RiskAssessmentPrediction) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Outcome.Equals(other.Outcome) { return false }
	if !m.ProbabilityDecimal.Equals(other.ProbabilityDecimal) { return false }
	if !m.ProbabilityRange.Equals(other.ProbabilityRange) { return false }
	if !m.QualitativeRisk.Equals(other.QualitativeRisk) { return false }
	if !m.RelativeRisk.Equals(other.RelativeRisk) { return false }
	if !m.WhenPeriod.Equals(other.WhenPeriod) { return false }
	if !m.WhenRange.Equals(other.WhenRange) { return false }
	if !m.Rationale.Equals(other.Rationale) { return false }
	return true
}

