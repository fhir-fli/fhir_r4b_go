// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"
	"fmt"
)

// Procedure
// An action that is or was performed on or for a patient. This can be a physical intervention like an operation, or less invasive like long term services, counseling, or hypnotherapy.
type Procedure struct {
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
	InstantiatesCanonical []*FhirCanonical `json:"instantiatescanonical,omitempty"`
	InstantiatesUri []*FhirUri `json:"instantiatesuri,omitempty"`
	BasedOn []*Reference `json:"basedon,omitempty"`
	PartOf []*Reference `json:"partof,omitempty"`
	Status *EventStatus `json:"status,omitempty"`
	StatusReason *CodeableConcept `json:"statusreason,omitempty"`
	Category *CodeableConcept `json:"category,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	Subject *Reference `json:"subject,omitempty"`
	Encounter *Reference `json:"encounter,omitempty"`
	PerformedDateTime *FhirDateTime `json:"performeddatetime,omitempty"`
	PerformedPeriod *Period `json:"performedperiod,omitempty"`
	PerformedString *FhirString `json:"performedstring,omitempty"`
	PerformedAge *Age `json:"performedage,omitempty"`
	PerformedRange *Range `json:"performedrange,omitempty"`
	Recorder *Reference `json:"recorder,omitempty"`
	Asserter *Reference `json:"asserter,omitempty"`
	Performer []*ProcedurePerformer `json:"performer,omitempty"`
	Location *Reference `json:"location,omitempty"`
	ReasonCode []*CodeableConcept `json:"reasoncode,omitempty"`
	ReasonReference []*Reference `json:"reasonreference,omitempty"`
	BodySite []*CodeableConcept `json:"bodysite,omitempty"`
	Outcome *CodeableConcept `json:"outcome,omitempty"`
	Report []*Reference `json:"report,omitempty"`
	Complication []*CodeableConcept `json:"complication,omitempty"`
	ComplicationDetail []*Reference `json:"complicationdetail,omitempty"`
	FollowUp []*CodeableConcept `json:"followup,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
	FocalDevice []*ProcedureFocalDevice `json:"focaldevice,omitempty"`
	UsedReference []*Reference `json:"usedreference,omitempty"`
	UsedCode []*CodeableConcept `json:"usedcode,omitempty"`
}

// NewProcedure creates a new Procedure instance.
func NewProcedure() *Procedure {
	return &Procedure{}
}

// UnmarshalJSON populates Procedure from JSON data.
func (m *Procedure) UnmarshalJSON(data []byte) error {
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
		InstantiatesCanonical []interface{} `json:"instantiatescanonical,omitempty"`
		InstantiatesUri []interface{} `json:"instantiatesuri,omitempty"`
		BasedOn []*Reference `json:"basedon,omitempty"`
		PartOf []*Reference `json:"partof,omitempty"`
		Status *EventStatus `json:"status,omitempty"`
		StatusReason *CodeableConcept `json:"statusreason,omitempty"`
		Category *CodeableConcept `json:"category,omitempty"`
		Code *CodeableConcept `json:"code,omitempty"`
		Subject *Reference `json:"subject,omitempty"`
		Encounter *Reference `json:"encounter,omitempty"`
		PerformedDateTime *FhirDateTime `json:"performeddatetime,omitempty"`
		PerformedPeriod *Period `json:"performedperiod,omitempty"`
		PerformedString *FhirString `json:"performedstring,omitempty"`
		PerformedAge *Age `json:"performedage,omitempty"`
		PerformedRange *Range `json:"performedrange,omitempty"`
		Recorder *Reference `json:"recorder,omitempty"`
		Asserter *Reference `json:"asserter,omitempty"`
		Performer []*ProcedurePerformer `json:"performer,omitempty"`
		Location *Reference `json:"location,omitempty"`
		ReasonCode []*CodeableConcept `json:"reasoncode,omitempty"`
		ReasonReference []*Reference `json:"reasonreference,omitempty"`
		BodySite []*CodeableConcept `json:"bodysite,omitempty"`
		Outcome *CodeableConcept `json:"outcome,omitempty"`
		Report []*Reference `json:"report,omitempty"`
		Complication []*CodeableConcept `json:"complication,omitempty"`
		ComplicationDetail []*Reference `json:"complicationdetail,omitempty"`
		FollowUp []*CodeableConcept `json:"followup,omitempty"`
		Note []*Annotation `json:"note,omitempty"`
		FocalDevice []*ProcedureFocalDevice `json:"focaldevice,omitempty"`
		UsedReference []*Reference `json:"usedreference,omitempty"`
		UsedCode []*CodeableConcept `json:"usedcode,omitempty"`
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
	if len(temp.InstantiatesCanonical) > 0 {
		m.InstantiatesCanonical = make([]*FhirCanonical, len(temp.InstantiatesCanonical))
		for i := range temp.InstantiatesCanonical {
			itemMap, ok := temp.InstantiatesCanonical[i].(map[string]interface{})
			if !ok { return fmt.Errorf("invalid value for InstantiatesCanonical[%d]: expected map", i) }
			primitive, err := NewFhirCanonicalFromMap(itemMap)
			if err != nil { return fmt.Errorf("failed to parse InstantiatesCanonical[%d]: %v", i, err) }
			m.InstantiatesCanonical[i] = primitive
		}
	}
	if len(temp.InstantiatesUri) > 0 {
		m.InstantiatesUri = make([]*FhirUri, len(temp.InstantiatesUri))
		for i := range temp.InstantiatesUri {
			itemMap, ok := temp.InstantiatesUri[i].(map[string]interface{})
			if !ok { return fmt.Errorf("invalid value for InstantiatesUri[%d]: expected map", i) }
			primitive, err := NewFhirUriFromMap(itemMap)
			if err != nil { return fmt.Errorf("failed to parse InstantiatesUri[%d]: %v", i, err) }
			m.InstantiatesUri[i] = primitive
		}
	}
	m.BasedOn = temp.BasedOn
	m.PartOf = temp.PartOf
	m.Status = temp.Status
	m.StatusReason = temp.StatusReason
	m.Category = temp.Category
	m.Code = temp.Code
	m.Subject = temp.Subject
	m.Encounter = temp.Encounter
	m.PerformedDateTime = temp.PerformedDateTime
	m.PerformedPeriod = temp.PerformedPeriod
	m.PerformedString = temp.PerformedString
	m.PerformedAge = temp.PerformedAge
	m.PerformedRange = temp.PerformedRange
	m.Recorder = temp.Recorder
	m.Asserter = temp.Asserter
	m.Performer = temp.Performer
	m.Location = temp.Location
	m.ReasonCode = temp.ReasonCode
	m.ReasonReference = temp.ReasonReference
	m.BodySite = temp.BodySite
	m.Outcome = temp.Outcome
	m.Report = temp.Report
	m.Complication = temp.Complication
	m.ComplicationDetail = temp.ComplicationDetail
	m.FollowUp = temp.FollowUp
	m.Note = temp.Note
	m.FocalDevice = temp.FocalDevice
	m.UsedReference = temp.UsedReference
	m.UsedCode = temp.UsedCode
	return nil
}

// MarshalJSON converts Procedure to JSON data.
func (m *Procedure) MarshalJSON() ([]byte, error) {
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
		InstantiatesCanonical []interface{} `json:"instantiatescanonical,omitempty"`
		InstantiatesCanonicalElement []map[string]interface{} `json:"_instantiatescanonical,omitempty"`
		InstantiatesUri []interface{} `json:"instantiatesuri,omitempty"`
		InstantiatesUriElement []map[string]interface{} `json:"_instantiatesuri,omitempty"`
		BasedOn []*Reference `json:"basedon,omitempty"`
		PartOf []*Reference `json:"partof,omitempty"`
		Status *EventStatus `json:"status,omitempty"`
		StatusReason *CodeableConcept `json:"statusreason,omitempty"`
		Category *CodeableConcept `json:"category,omitempty"`
		Code *CodeableConcept `json:"code,omitempty"`
		Subject *Reference `json:"subject,omitempty"`
		Encounter *Reference `json:"encounter,omitempty"`
		PerformedDateTime interface{} `json:"performeddatetime,omitempty"`
		PerformedDateTimeElement map[string]interface{} `json:"_performeddatetime,omitempty"`
		PerformedPeriod *Period `json:"performedperiod,omitempty"`
		PerformedString interface{} `json:"performedstring,omitempty"`
		PerformedStringElement map[string]interface{} `json:"_performedstring,omitempty"`
		PerformedAge *Age `json:"performedage,omitempty"`
		PerformedRange *Range `json:"performedrange,omitempty"`
		Recorder *Reference `json:"recorder,omitempty"`
		Asserter *Reference `json:"asserter,omitempty"`
		Performer []*ProcedurePerformer `json:"performer,omitempty"`
		Location *Reference `json:"location,omitempty"`
		ReasonCode []*CodeableConcept `json:"reasoncode,omitempty"`
		ReasonReference []*Reference `json:"reasonreference,omitempty"`
		BodySite []*CodeableConcept `json:"bodysite,omitempty"`
		Outcome *CodeableConcept `json:"outcome,omitempty"`
		Report []*Reference `json:"report,omitempty"`
		Complication []*CodeableConcept `json:"complication,omitempty"`
		ComplicationDetail []*Reference `json:"complicationdetail,omitempty"`
		FollowUp []*CodeableConcept `json:"followup,omitempty"`
		Note []*Annotation `json:"note,omitempty"`
		FocalDevice []*ProcedureFocalDevice `json:"focaldevice,omitempty"`
		UsedReference []*Reference `json:"usedreference,omitempty"`
		UsedCode []*CodeableConcept `json:"usedcode,omitempty"`
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
	if len(m.InstantiatesCanonical) > 0 {
		output.InstantiatesCanonical = make([]interface{}, len(m.InstantiatesCanonical))
		output.InstantiatesCanonicalElement = make([]map[string]interface{}, len(m.InstantiatesCanonical))
		for i, item := range m.InstantiatesCanonical {
			if item != nil && item.Value != nil {
				output.InstantiatesCanonical[i] = item.Value
			}
			if item != nil && item.Element != nil {
				output.InstantiatesCanonicalElement[i] = toMapOrNil(item.Element.MarshalJSON())
			}
		}
	}
	if len(m.InstantiatesUri) > 0 {
		output.InstantiatesUri = make([]interface{}, len(m.InstantiatesUri))
		output.InstantiatesUriElement = make([]map[string]interface{}, len(m.InstantiatesUri))
		for i, item := range m.InstantiatesUri {
			if item != nil && item.Value != nil {
				output.InstantiatesUri[i] = item.Value
			}
			if item != nil && item.Element != nil {
				output.InstantiatesUriElement[i] = toMapOrNil(item.Element.MarshalJSON())
			}
		}
	}
	output.BasedOn = m.BasedOn
	output.PartOf = m.PartOf
	output.Status = m.Status
	output.StatusReason = m.StatusReason
	output.Category = m.Category
	output.Code = m.Code
	output.Subject = m.Subject
	output.Encounter = m.Encounter
	if m.PerformedDateTime != nil && m.PerformedDateTime.Value != nil {
		output.PerformedDateTime = m.PerformedDateTime.Value
		if m.PerformedDateTime.Element != nil {
			output.PerformedDateTimeElement = toMapOrNil(m.PerformedDateTime.Element.MarshalJSON())
		}
	}
	output.PerformedPeriod = m.PerformedPeriod
	if m.PerformedString != nil && m.PerformedString.Value != nil {
		output.PerformedString = m.PerformedString.Value
		if m.PerformedString.Element != nil {
			output.PerformedStringElement = toMapOrNil(m.PerformedString.Element.MarshalJSON())
		}
	}
	output.PerformedAge = m.PerformedAge
	output.PerformedRange = m.PerformedRange
	output.Recorder = m.Recorder
	output.Asserter = m.Asserter
	output.Performer = m.Performer
	output.Location = m.Location
	output.ReasonCode = m.ReasonCode
	output.ReasonReference = m.ReasonReference
	output.BodySite = m.BodySite
	output.Outcome = m.Outcome
	output.Report = m.Report
	output.Complication = m.Complication
	output.ComplicationDetail = m.ComplicationDetail
	output.FollowUp = m.FollowUp
	output.Note = m.Note
	output.FocalDevice = m.FocalDevice
	output.UsedReference = m.UsedReference
	output.UsedCode = m.UsedCode
	return json.Marshal(output)
}

// Clone creates a deep copy of Procedure.
func (m *Procedure) Clone() *Procedure {
	if m == nil { return nil }
	return &Procedure{
		Id: m.Id.Clone(),
		Meta: m.Meta.Clone(),
		ImplicitRules: m.ImplicitRules.Clone(),
		Language: m.Language.Clone(),
		Text: m.Text.Clone(),
		Contained: cloneSlices(m.Contained),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Identifier: cloneSlices(m.Identifier),
		InstantiatesCanonical: cloneSlices(m.InstantiatesCanonical),
		InstantiatesUri: cloneSlices(m.InstantiatesUri),
		BasedOn: cloneSlices(m.BasedOn),
		PartOf: cloneSlices(m.PartOf),
		Status: m.Status.Clone(),
		StatusReason: m.StatusReason.Clone(),
		Category: m.Category.Clone(),
		Code: m.Code.Clone(),
		Subject: m.Subject.Clone(),
		Encounter: m.Encounter.Clone(),
		PerformedDateTime: m.PerformedDateTime.Clone(),
		PerformedPeriod: m.PerformedPeriod.Clone(),
		PerformedString: m.PerformedString.Clone(),
		PerformedAge: m.PerformedAge.Clone(),
		PerformedRange: m.PerformedRange.Clone(),
		Recorder: m.Recorder.Clone(),
		Asserter: m.Asserter.Clone(),
		Performer: cloneSlices(m.Performer),
		Location: m.Location.Clone(),
		ReasonCode: cloneSlices(m.ReasonCode),
		ReasonReference: cloneSlices(m.ReasonReference),
		BodySite: cloneSlices(m.BodySite),
		Outcome: m.Outcome.Clone(),
		Report: cloneSlices(m.Report),
		Complication: cloneSlices(m.Complication),
		ComplicationDetail: cloneSlices(m.ComplicationDetail),
		FollowUp: cloneSlices(m.FollowUp),
		Note: cloneSlices(m.Note),
		FocalDevice: cloneSlices(m.FocalDevice),
		UsedReference: cloneSlices(m.UsedReference),
		UsedCode: cloneSlices(m.UsedCode),
	}
}

// Equals checks equality between two Procedure instances.
func (m *Procedure) Equals(other *Procedure) bool {
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
	if !compareSlices(m.InstantiatesCanonical, other.InstantiatesCanonical) { return false }
	if !compareSlices(m.InstantiatesUri, other.InstantiatesUri) { return false }
	if !compareSlices(m.BasedOn, other.BasedOn) { return false }
	if !compareSlices(m.PartOf, other.PartOf) { return false }
	if !m.Status.Equals(other.Status) { return false }
	if !m.StatusReason.Equals(other.StatusReason) { return false }
	if !m.Category.Equals(other.Category) { return false }
	if !m.Code.Equals(other.Code) { return false }
	if !m.Subject.Equals(other.Subject) { return false }
	if !m.Encounter.Equals(other.Encounter) { return false }
	if !m.PerformedDateTime.Equals(other.PerformedDateTime) { return false }
	if !m.PerformedPeriod.Equals(other.PerformedPeriod) { return false }
	if !m.PerformedString.Equals(other.PerformedString) { return false }
	if !m.PerformedAge.Equals(other.PerformedAge) { return false }
	if !m.PerformedRange.Equals(other.PerformedRange) { return false }
	if !m.Recorder.Equals(other.Recorder) { return false }
	if !m.Asserter.Equals(other.Asserter) { return false }
	if !compareSlices(m.Performer, other.Performer) { return false }
	if !m.Location.Equals(other.Location) { return false }
	if !compareSlices(m.ReasonCode, other.ReasonCode) { return false }
	if !compareSlices(m.ReasonReference, other.ReasonReference) { return false }
	if !compareSlices(m.BodySite, other.BodySite) { return false }
	if !m.Outcome.Equals(other.Outcome) { return false }
	if !compareSlices(m.Report, other.Report) { return false }
	if !compareSlices(m.Complication, other.Complication) { return false }
	if !compareSlices(m.ComplicationDetail, other.ComplicationDetail) { return false }
	if !compareSlices(m.FollowUp, other.FollowUp) { return false }
	if !compareSlices(m.Note, other.Note) { return false }
	if !compareSlices(m.FocalDevice, other.FocalDevice) { return false }
	if !compareSlices(m.UsedReference, other.UsedReference) { return false }
	if !compareSlices(m.UsedCode, other.UsedCode) { return false }
	return true
}

// ProcedurePerformer
// Limited to "real" people rather than equipment.
type ProcedurePerformer struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Function_ *CodeableConcept `json:"function,omitempty"`
	Actor *Reference `json:"actor,omitempty"`
	OnBehalfOf *Reference `json:"onbehalfof,omitempty"`
}

// NewProcedurePerformer creates a new ProcedurePerformer instance.
func NewProcedurePerformer() *ProcedurePerformer {
	return &ProcedurePerformer{}
}

// UnmarshalJSON populates ProcedurePerformer from JSON data.
func (m *ProcedurePerformer) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Function_ *CodeableConcept `json:"function,omitempty"`
		Actor *Reference `json:"actor,omitempty"`
		OnBehalfOf *Reference `json:"onbehalfof,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Function_ = temp.Function_
	m.Actor = temp.Actor
	m.OnBehalfOf = temp.OnBehalfOf
	return nil
}

// MarshalJSON converts ProcedurePerformer to JSON data.
func (m *ProcedurePerformer) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Function_ *CodeableConcept `json:"function,omitempty"`
		Actor *Reference `json:"actor,omitempty"`
		OnBehalfOf *Reference `json:"onbehalfof,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.Function_ = m.Function_
	output.Actor = m.Actor
	output.OnBehalfOf = m.OnBehalfOf
	return json.Marshal(output)
}

// Clone creates a deep copy of ProcedurePerformer.
func (m *ProcedurePerformer) Clone() *ProcedurePerformer {
	if m == nil { return nil }
	return &ProcedurePerformer{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Function_: m.Function_.Clone(),
		Actor: m.Actor.Clone(),
		OnBehalfOf: m.OnBehalfOf.Clone(),
	}
}

// Equals checks equality between two ProcedurePerformer instances.
func (m *ProcedurePerformer) Equals(other *ProcedurePerformer) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Function_.Equals(other.Function_) { return false }
	if !m.Actor.Equals(other.Actor) { return false }
	if !m.OnBehalfOf.Equals(other.OnBehalfOf) { return false }
	return true
}

// ProcedureFocalDevice
// A device that is implanted, removed or otherwise manipulated (calibration, battery replacement, fitting a prosthesis, attaching a wound-vac, etc.) as a focal portion of the Procedure.
type ProcedureFocalDevice struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Action *CodeableConcept `json:"action,omitempty"`
	Manipulated *Reference `json:"manipulated,omitempty"`
}

// NewProcedureFocalDevice creates a new ProcedureFocalDevice instance.
func NewProcedureFocalDevice() *ProcedureFocalDevice {
	return &ProcedureFocalDevice{}
}

// UnmarshalJSON populates ProcedureFocalDevice from JSON data.
func (m *ProcedureFocalDevice) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Action *CodeableConcept `json:"action,omitempty"`
		Manipulated *Reference `json:"manipulated,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Action = temp.Action
	m.Manipulated = temp.Manipulated
	return nil
}

// MarshalJSON converts ProcedureFocalDevice to JSON data.
func (m *ProcedureFocalDevice) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Action *CodeableConcept `json:"action,omitempty"`
		Manipulated *Reference `json:"manipulated,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.Action = m.Action
	output.Manipulated = m.Manipulated
	return json.Marshal(output)
}

// Clone creates a deep copy of ProcedureFocalDevice.
func (m *ProcedureFocalDevice) Clone() *ProcedureFocalDevice {
	if m == nil { return nil }
	return &ProcedureFocalDevice{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Action: m.Action.Clone(),
		Manipulated: m.Manipulated.Clone(),
	}
}

// Equals checks equality between two ProcedureFocalDevice instances.
func (m *ProcedureFocalDevice) Equals(other *ProcedureFocalDevice) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Action.Equals(other.Action) { return false }
	if !m.Manipulated.Equals(other.Manipulated) { return false }
	return true
}

