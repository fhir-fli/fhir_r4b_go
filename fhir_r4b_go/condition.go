// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"
)

// Condition
// A clinical condition, problem, diagnosis, or other event, situation, issue, or clinical concept that has risen to a level of concern.
type Condition struct {
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
	Category []*CodeableConcept `json:"category,omitempty"`
	Severity *CodeableConcept `json:"severity,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	BodySite []*CodeableConcept `json:"bodysite,omitempty"`
	Subject *Reference `json:"subject,omitempty"`
	Encounter *Reference `json:"encounter,omitempty"`
	OnsetDateTime *FhirDateTime `json:"onsetdatetime,omitempty"`
	OnsetAge *Age `json:"onsetage,omitempty"`
	OnsetPeriod *Period `json:"onsetperiod,omitempty"`
	OnsetRange *Range `json:"onsetrange,omitempty"`
	OnsetString *FhirString `json:"onsetstring,omitempty"`
	AbatementDateTime *FhirDateTime `json:"abatementdatetime,omitempty"`
	AbatementAge *Age `json:"abatementage,omitempty"`
	AbatementPeriod *Period `json:"abatementperiod,omitempty"`
	AbatementRange *Range `json:"abatementrange,omitempty"`
	AbatementString *FhirString `json:"abatementstring,omitempty"`
	RecordedDate *FhirDateTime `json:"recordeddate,omitempty"`
	Recorder *Reference `json:"recorder,omitempty"`
	Asserter *Reference `json:"asserter,omitempty"`
	Stage []*ConditionStage `json:"stage,omitempty"`
	Evidence []*ConditionEvidence `json:"evidence,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
}

// NewCondition creates a new Condition instance.
func NewCondition() *Condition {
	return &Condition{}
}

// UnmarshalJSON populates Condition from JSON data.
func (m *Condition) UnmarshalJSON(data []byte) error {
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
		Category []*CodeableConcept `json:"category,omitempty"`
		Severity *CodeableConcept `json:"severity,omitempty"`
		Code *CodeableConcept `json:"code,omitempty"`
		BodySite []*CodeableConcept `json:"bodysite,omitempty"`
		Subject *Reference `json:"subject,omitempty"`
		Encounter *Reference `json:"encounter,omitempty"`
		OnsetDateTime *FhirDateTime `json:"onsetdatetime,omitempty"`
		OnsetAge *Age `json:"onsetage,omitempty"`
		OnsetPeriod *Period `json:"onsetperiod,omitempty"`
		OnsetRange *Range `json:"onsetrange,omitempty"`
		OnsetString *FhirString `json:"onsetstring,omitempty"`
		AbatementDateTime *FhirDateTime `json:"abatementdatetime,omitempty"`
		AbatementAge *Age `json:"abatementage,omitempty"`
		AbatementPeriod *Period `json:"abatementperiod,omitempty"`
		AbatementRange *Range `json:"abatementrange,omitempty"`
		AbatementString *FhirString `json:"abatementstring,omitempty"`
		RecordedDate *FhirDateTime `json:"recordeddate,omitempty"`
		Recorder *Reference `json:"recorder,omitempty"`
		Asserter *Reference `json:"asserter,omitempty"`
		Stage []*ConditionStage `json:"stage,omitempty"`
		Evidence []*ConditionEvidence `json:"evidence,omitempty"`
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
	m.ClinicalStatus = temp.ClinicalStatus
	m.VerificationStatus = temp.VerificationStatus
	m.Category = temp.Category
	m.Severity = temp.Severity
	m.Code = temp.Code
	m.BodySite = temp.BodySite
	m.Subject = temp.Subject
	m.Encounter = temp.Encounter
	m.OnsetDateTime = temp.OnsetDateTime
	m.OnsetAge = temp.OnsetAge
	m.OnsetPeriod = temp.OnsetPeriod
	m.OnsetRange = temp.OnsetRange
	m.OnsetString = temp.OnsetString
	m.AbatementDateTime = temp.AbatementDateTime
	m.AbatementAge = temp.AbatementAge
	m.AbatementPeriod = temp.AbatementPeriod
	m.AbatementRange = temp.AbatementRange
	m.AbatementString = temp.AbatementString
	m.RecordedDate = temp.RecordedDate
	m.Recorder = temp.Recorder
	m.Asserter = temp.Asserter
	m.Stage = temp.Stage
	m.Evidence = temp.Evidence
	m.Note = temp.Note
	return nil
}

// MarshalJSON converts Condition to JSON data.
func (m *Condition) MarshalJSON() ([]byte, error) {
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
		Category []*CodeableConcept `json:"category,omitempty"`
		Severity *CodeableConcept `json:"severity,omitempty"`
		Code *CodeableConcept `json:"code,omitempty"`
		BodySite []*CodeableConcept `json:"bodysite,omitempty"`
		Subject *Reference `json:"subject,omitempty"`
		Encounter *Reference `json:"encounter,omitempty"`
		OnsetDateTime interface{} `json:"onsetdatetime,omitempty"`
		OnsetDateTimeElement map[string]interface{} `json:"_onsetdatetime,omitempty"`
		OnsetAge *Age `json:"onsetage,omitempty"`
		OnsetPeriod *Period `json:"onsetperiod,omitempty"`
		OnsetRange *Range `json:"onsetrange,omitempty"`
		OnsetString interface{} `json:"onsetstring,omitempty"`
		OnsetStringElement map[string]interface{} `json:"_onsetstring,omitempty"`
		AbatementDateTime interface{} `json:"abatementdatetime,omitempty"`
		AbatementDateTimeElement map[string]interface{} `json:"_abatementdatetime,omitempty"`
		AbatementAge *Age `json:"abatementage,omitempty"`
		AbatementPeriod *Period `json:"abatementperiod,omitempty"`
		AbatementRange *Range `json:"abatementrange,omitempty"`
		AbatementString interface{} `json:"abatementstring,omitempty"`
		AbatementStringElement map[string]interface{} `json:"_abatementstring,omitempty"`
		RecordedDate interface{} `json:"recordeddate,omitempty"`
		RecordedDateElement map[string]interface{} `json:"_recordeddate,omitempty"`
		Recorder *Reference `json:"recorder,omitempty"`
		Asserter *Reference `json:"asserter,omitempty"`
		Stage []*ConditionStage `json:"stage,omitempty"`
		Evidence []*ConditionEvidence `json:"evidence,omitempty"`
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
	output.ClinicalStatus = m.ClinicalStatus
	output.VerificationStatus = m.VerificationStatus
	output.Category = m.Category
	output.Severity = m.Severity
	output.Code = m.Code
	output.BodySite = m.BodySite
	output.Subject = m.Subject
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
	if m.AbatementDateTime != nil && m.AbatementDateTime.Value != nil {
		output.AbatementDateTime = m.AbatementDateTime.Value
		if m.AbatementDateTime.Element != nil {
			output.AbatementDateTimeElement = toMapOrNil(m.AbatementDateTime.Element.MarshalJSON())
		}
	}
	output.AbatementAge = m.AbatementAge
	output.AbatementPeriod = m.AbatementPeriod
	output.AbatementRange = m.AbatementRange
	if m.AbatementString != nil && m.AbatementString.Value != nil {
		output.AbatementString = m.AbatementString.Value
		if m.AbatementString.Element != nil {
			output.AbatementStringElement = toMapOrNil(m.AbatementString.Element.MarshalJSON())
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
	output.Stage = m.Stage
	output.Evidence = m.Evidence
	output.Note = m.Note
	return json.Marshal(output)
}

// Clone creates a deep copy of Condition.
func (m *Condition) Clone() *Condition {
	if m == nil { return nil }
	return &Condition{
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
		Category: cloneSlices(m.Category),
		Severity: m.Severity.Clone(),
		Code: m.Code.Clone(),
		BodySite: cloneSlices(m.BodySite),
		Subject: m.Subject.Clone(),
		Encounter: m.Encounter.Clone(),
		OnsetDateTime: m.OnsetDateTime.Clone(),
		OnsetAge: m.OnsetAge.Clone(),
		OnsetPeriod: m.OnsetPeriod.Clone(),
		OnsetRange: m.OnsetRange.Clone(),
		OnsetString: m.OnsetString.Clone(),
		AbatementDateTime: m.AbatementDateTime.Clone(),
		AbatementAge: m.AbatementAge.Clone(),
		AbatementPeriod: m.AbatementPeriod.Clone(),
		AbatementRange: m.AbatementRange.Clone(),
		AbatementString: m.AbatementString.Clone(),
		RecordedDate: m.RecordedDate.Clone(),
		Recorder: m.Recorder.Clone(),
		Asserter: m.Asserter.Clone(),
		Stage: cloneSlices(m.Stage),
		Evidence: cloneSlices(m.Evidence),
		Note: cloneSlices(m.Note),
	}
}

// Equals checks equality between two Condition instances.
func (m *Condition) Equals(other *Condition) bool {
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
	if !compareSlices(m.Category, other.Category) { return false }
	if !m.Severity.Equals(other.Severity) { return false }
	if !m.Code.Equals(other.Code) { return false }
	if !compareSlices(m.BodySite, other.BodySite) { return false }
	if !m.Subject.Equals(other.Subject) { return false }
	if !m.Encounter.Equals(other.Encounter) { return false }
	if !m.OnsetDateTime.Equals(other.OnsetDateTime) { return false }
	if !m.OnsetAge.Equals(other.OnsetAge) { return false }
	if !m.OnsetPeriod.Equals(other.OnsetPeriod) { return false }
	if !m.OnsetRange.Equals(other.OnsetRange) { return false }
	if !m.OnsetString.Equals(other.OnsetString) { return false }
	if !m.AbatementDateTime.Equals(other.AbatementDateTime) { return false }
	if !m.AbatementAge.Equals(other.AbatementAge) { return false }
	if !m.AbatementPeriod.Equals(other.AbatementPeriod) { return false }
	if !m.AbatementRange.Equals(other.AbatementRange) { return false }
	if !m.AbatementString.Equals(other.AbatementString) { return false }
	if !m.RecordedDate.Equals(other.RecordedDate) { return false }
	if !m.Recorder.Equals(other.Recorder) { return false }
	if !m.Asserter.Equals(other.Asserter) { return false }
	if !compareSlices(m.Stage, other.Stage) { return false }
	if !compareSlices(m.Evidence, other.Evidence) { return false }
	if !compareSlices(m.Note, other.Note) { return false }
	return true
}

// ConditionStage
// Clinical stage or grade of a condition. May include formal severity assessments.
type ConditionStage struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Summary *CodeableConcept `json:"summary,omitempty"`
	Assessment []*Reference `json:"assessment,omitempty"`
	Type *CodeableConcept `json:"type,omitempty"`
}

// NewConditionStage creates a new ConditionStage instance.
func NewConditionStage() *ConditionStage {
	return &ConditionStage{}
}

// UnmarshalJSON populates ConditionStage from JSON data.
func (m *ConditionStage) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Summary *CodeableConcept `json:"summary,omitempty"`
		Assessment []*Reference `json:"assessment,omitempty"`
		Type *CodeableConcept `json:"type,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Summary = temp.Summary
	m.Assessment = temp.Assessment
	m.Type = temp.Type
	return nil
}

// MarshalJSON converts ConditionStage to JSON data.
func (m *ConditionStage) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Summary *CodeableConcept `json:"summary,omitempty"`
		Assessment []*Reference `json:"assessment,omitempty"`
		Type *CodeableConcept `json:"type,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.Summary = m.Summary
	output.Assessment = m.Assessment
	output.Type = m.Type
	return json.Marshal(output)
}

// Clone creates a deep copy of ConditionStage.
func (m *ConditionStage) Clone() *ConditionStage {
	if m == nil { return nil }
	return &ConditionStage{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Summary: m.Summary.Clone(),
		Assessment: cloneSlices(m.Assessment),
		Type: m.Type.Clone(),
	}
}

// Equals checks equality between two ConditionStage instances.
func (m *ConditionStage) Equals(other *ConditionStage) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Summary.Equals(other.Summary) { return false }
	if !compareSlices(m.Assessment, other.Assessment) { return false }
	if !m.Type.Equals(other.Type) { return false }
	return true
}

// ConditionEvidence
// Supporting evidence / manifestations that are the basis of the Condition's verification status, such as evidence that confirmed or refuted the condition.
type ConditionEvidence struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Code []*CodeableConcept `json:"code,omitempty"`
	Detail []*Reference `json:"detail,omitempty"`
}

// NewConditionEvidence creates a new ConditionEvidence instance.
func NewConditionEvidence() *ConditionEvidence {
	return &ConditionEvidence{}
}

// UnmarshalJSON populates ConditionEvidence from JSON data.
func (m *ConditionEvidence) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Code []*CodeableConcept `json:"code,omitempty"`
		Detail []*Reference `json:"detail,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Code = temp.Code
	m.Detail = temp.Detail
	return nil
}

// MarshalJSON converts ConditionEvidence to JSON data.
func (m *ConditionEvidence) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Code []*CodeableConcept `json:"code,omitempty"`
		Detail []*Reference `json:"detail,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.Code = m.Code
	output.Detail = m.Detail
	return json.Marshal(output)
}

// Clone creates a deep copy of ConditionEvidence.
func (m *ConditionEvidence) Clone() *ConditionEvidence {
	if m == nil { return nil }
	return &ConditionEvidence{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Code: cloneSlices(m.Code),
		Detail: cloneSlices(m.Detail),
	}
}

// Equals checks equality between two ConditionEvidence instances.
func (m *ConditionEvidence) Equals(other *ConditionEvidence) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !compareSlices(m.Code, other.Code) { return false }
	if !compareSlices(m.Detail, other.Detail) { return false }
	return true
}

