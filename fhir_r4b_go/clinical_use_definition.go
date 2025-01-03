// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"
)

// ClinicalUseDefinition
// A single issue - either an indication, contraindication, interaction or an undesirable effect for a medicinal product, medication, device or procedure.
type ClinicalUseDefinition struct {
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
	Type *ClinicalUseDefinitionType `json:"type,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Subject []*Reference `json:"subject,omitempty"`
	Status *CodeableConcept `json:"status,omitempty"`
	Contraindication *ClinicalUseDefinitionContraindication `json:"contraindication,omitempty"`
	Indication *ClinicalUseDefinitionIndication `json:"indication,omitempty"`
	Interaction *ClinicalUseDefinitionInteraction `json:"interaction,omitempty"`
	Population []*Reference `json:"population,omitempty"`
	UndesirableEffect *ClinicalUseDefinitionUndesirableEffect `json:"undesirableeffect,omitempty"`
	Warning *ClinicalUseDefinitionWarning `json:"warning,omitempty"`
}

// NewClinicalUseDefinition creates a new ClinicalUseDefinition instance.
func NewClinicalUseDefinition() *ClinicalUseDefinition {
	return &ClinicalUseDefinition{}
}

// UnmarshalJSON populates ClinicalUseDefinition from JSON data.
func (m *ClinicalUseDefinition) UnmarshalJSON(data []byte) error {
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
		Type *ClinicalUseDefinitionType `json:"type,omitempty"`
		Category []*CodeableConcept `json:"category,omitempty"`
		Subject []*Reference `json:"subject,omitempty"`
		Status *CodeableConcept `json:"status,omitempty"`
		Contraindication *ClinicalUseDefinitionContraindication `json:"contraindication,omitempty"`
		Indication *ClinicalUseDefinitionIndication `json:"indication,omitempty"`
		Interaction *ClinicalUseDefinitionInteraction `json:"interaction,omitempty"`
		Population []*Reference `json:"population,omitempty"`
		UndesirableEffect *ClinicalUseDefinitionUndesirableEffect `json:"undesirableeffect,omitempty"`
		Warning *ClinicalUseDefinitionWarning `json:"warning,omitempty"`
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
	m.Type = temp.Type
	m.Category = temp.Category
	m.Subject = temp.Subject
	m.Status = temp.Status
	m.Contraindication = temp.Contraindication
	m.Indication = temp.Indication
	m.Interaction = temp.Interaction
	m.Population = temp.Population
	m.UndesirableEffect = temp.UndesirableEffect
	m.Warning = temp.Warning
	return nil
}

// MarshalJSON converts ClinicalUseDefinition to JSON data.
func (m *ClinicalUseDefinition) MarshalJSON() ([]byte, error) {
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
		Type *ClinicalUseDefinitionType `json:"type,omitempty"`
		Category []*CodeableConcept `json:"category,omitempty"`
		Subject []*Reference `json:"subject,omitempty"`
		Status *CodeableConcept `json:"status,omitempty"`
		Contraindication *ClinicalUseDefinitionContraindication `json:"contraindication,omitempty"`
		Indication *ClinicalUseDefinitionIndication `json:"indication,omitempty"`
		Interaction *ClinicalUseDefinitionInteraction `json:"interaction,omitempty"`
		Population []*Reference `json:"population,omitempty"`
		UndesirableEffect *ClinicalUseDefinitionUndesirableEffect `json:"undesirableeffect,omitempty"`
		Warning *ClinicalUseDefinitionWarning `json:"warning,omitempty"`
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
	output.Type = m.Type
	output.Category = m.Category
	output.Subject = m.Subject
	output.Status = m.Status
	output.Contraindication = m.Contraindication
	output.Indication = m.Indication
	output.Interaction = m.Interaction
	output.Population = m.Population
	output.UndesirableEffect = m.UndesirableEffect
	output.Warning = m.Warning
	return json.Marshal(output)
}

// Clone creates a deep copy of ClinicalUseDefinition.
func (m *ClinicalUseDefinition) Clone() *ClinicalUseDefinition {
	if m == nil { return nil }
	return &ClinicalUseDefinition{
		Id: m.Id.Clone(),
		Meta: m.Meta.Clone(),
		ImplicitRules: m.ImplicitRules.Clone(),
		Language: m.Language.Clone(),
		Text: m.Text.Clone(),
		Contained: cloneSlices(m.Contained),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Identifier: cloneSlices(m.Identifier),
		Type: m.Type.Clone(),
		Category: cloneSlices(m.Category),
		Subject: cloneSlices(m.Subject),
		Status: m.Status.Clone(),
		Contraindication: m.Contraindication.Clone(),
		Indication: m.Indication.Clone(),
		Interaction: m.Interaction.Clone(),
		Population: cloneSlices(m.Population),
		UndesirableEffect: m.UndesirableEffect.Clone(),
		Warning: m.Warning.Clone(),
	}
}

// Equals checks equality between two ClinicalUseDefinition instances.
func (m *ClinicalUseDefinition) Equals(other *ClinicalUseDefinition) bool {
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
	if !m.Type.Equals(other.Type) { return false }
	if !compareSlices(m.Category, other.Category) { return false }
	if !compareSlices(m.Subject, other.Subject) { return false }
	if !m.Status.Equals(other.Status) { return false }
	if !m.Contraindication.Equals(other.Contraindication) { return false }
	if !m.Indication.Equals(other.Indication) { return false }
	if !m.Interaction.Equals(other.Interaction) { return false }
	if !compareSlices(m.Population, other.Population) { return false }
	if !m.UndesirableEffect.Equals(other.UndesirableEffect) { return false }
	if !m.Warning.Equals(other.Warning) { return false }
	return true
}

// ClinicalUseDefinitionContraindication
// Specifics for when this is a contraindication.
type ClinicalUseDefinitionContraindication struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	DiseaseSymptomProcedure *CodeableReference `json:"diseasesymptomprocedure,omitempty"`
	DiseaseStatus *CodeableReference `json:"diseasestatus,omitempty"`
	Comorbidity []*CodeableReference `json:"comorbidity,omitempty"`
	Indication []*Reference `json:"indication,omitempty"`
	OtherTherapy []*ClinicalUseDefinitionOtherTherapy `json:"othertherapy,omitempty"`
}

// NewClinicalUseDefinitionContraindication creates a new ClinicalUseDefinitionContraindication instance.
func NewClinicalUseDefinitionContraindication() *ClinicalUseDefinitionContraindication {
	return &ClinicalUseDefinitionContraindication{}
}

// UnmarshalJSON populates ClinicalUseDefinitionContraindication from JSON data.
func (m *ClinicalUseDefinitionContraindication) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		DiseaseSymptomProcedure *CodeableReference `json:"diseasesymptomprocedure,omitempty"`
		DiseaseStatus *CodeableReference `json:"diseasestatus,omitempty"`
		Comorbidity []*CodeableReference `json:"comorbidity,omitempty"`
		Indication []*Reference `json:"indication,omitempty"`
		OtherTherapy []*ClinicalUseDefinitionOtherTherapy `json:"othertherapy,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.DiseaseSymptomProcedure = temp.DiseaseSymptomProcedure
	m.DiseaseStatus = temp.DiseaseStatus
	m.Comorbidity = temp.Comorbidity
	m.Indication = temp.Indication
	m.OtherTherapy = temp.OtherTherapy
	return nil
}

// MarshalJSON converts ClinicalUseDefinitionContraindication to JSON data.
func (m *ClinicalUseDefinitionContraindication) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		DiseaseSymptomProcedure *CodeableReference `json:"diseasesymptomprocedure,omitempty"`
		DiseaseStatus *CodeableReference `json:"diseasestatus,omitempty"`
		Comorbidity []*CodeableReference `json:"comorbidity,omitempty"`
		Indication []*Reference `json:"indication,omitempty"`
		OtherTherapy []*ClinicalUseDefinitionOtherTherapy `json:"othertherapy,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.DiseaseSymptomProcedure = m.DiseaseSymptomProcedure
	output.DiseaseStatus = m.DiseaseStatus
	output.Comorbidity = m.Comorbidity
	output.Indication = m.Indication
	output.OtherTherapy = m.OtherTherapy
	return json.Marshal(output)
}

// Clone creates a deep copy of ClinicalUseDefinitionContraindication.
func (m *ClinicalUseDefinitionContraindication) Clone() *ClinicalUseDefinitionContraindication {
	if m == nil { return nil }
	return &ClinicalUseDefinitionContraindication{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		DiseaseSymptomProcedure: m.DiseaseSymptomProcedure.Clone(),
		DiseaseStatus: m.DiseaseStatus.Clone(),
		Comorbidity: cloneSlices(m.Comorbidity),
		Indication: cloneSlices(m.Indication),
		OtherTherapy: cloneSlices(m.OtherTherapy),
	}
}

// Equals checks equality between two ClinicalUseDefinitionContraindication instances.
func (m *ClinicalUseDefinitionContraindication) Equals(other *ClinicalUseDefinitionContraindication) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.DiseaseSymptomProcedure.Equals(other.DiseaseSymptomProcedure) { return false }
	if !m.DiseaseStatus.Equals(other.DiseaseStatus) { return false }
	if !compareSlices(m.Comorbidity, other.Comorbidity) { return false }
	if !compareSlices(m.Indication, other.Indication) { return false }
	if !compareSlices(m.OtherTherapy, other.OtherTherapy) { return false }
	return true
}

// ClinicalUseDefinitionOtherTherapy
// Information about the use of the medicinal product in relation to other therapies described as part of the contraindication.
type ClinicalUseDefinitionOtherTherapy struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	RelationshipType *CodeableConcept `json:"relationshiptype,omitempty"`
	Therapy *CodeableReference `json:"therapy,omitempty"`
}

// NewClinicalUseDefinitionOtherTherapy creates a new ClinicalUseDefinitionOtherTherapy instance.
func NewClinicalUseDefinitionOtherTherapy() *ClinicalUseDefinitionOtherTherapy {
	return &ClinicalUseDefinitionOtherTherapy{}
}

// UnmarshalJSON populates ClinicalUseDefinitionOtherTherapy from JSON data.
func (m *ClinicalUseDefinitionOtherTherapy) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		RelationshipType *CodeableConcept `json:"relationshiptype,omitempty"`
		Therapy *CodeableReference `json:"therapy,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.RelationshipType = temp.RelationshipType
	m.Therapy = temp.Therapy
	return nil
}

// MarshalJSON converts ClinicalUseDefinitionOtherTherapy to JSON data.
func (m *ClinicalUseDefinitionOtherTherapy) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		RelationshipType *CodeableConcept `json:"relationshiptype,omitempty"`
		Therapy *CodeableReference `json:"therapy,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.RelationshipType = m.RelationshipType
	output.Therapy = m.Therapy
	return json.Marshal(output)
}

// Clone creates a deep copy of ClinicalUseDefinitionOtherTherapy.
func (m *ClinicalUseDefinitionOtherTherapy) Clone() *ClinicalUseDefinitionOtherTherapy {
	if m == nil { return nil }
	return &ClinicalUseDefinitionOtherTherapy{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		RelationshipType: m.RelationshipType.Clone(),
		Therapy: m.Therapy.Clone(),
	}
}

// Equals checks equality between two ClinicalUseDefinitionOtherTherapy instances.
func (m *ClinicalUseDefinitionOtherTherapy) Equals(other *ClinicalUseDefinitionOtherTherapy) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.RelationshipType.Equals(other.RelationshipType) { return false }
	if !m.Therapy.Equals(other.Therapy) { return false }
	return true
}

// ClinicalUseDefinitionIndication
// Specifics for when this is an indication.
type ClinicalUseDefinitionIndication struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	DiseaseSymptomProcedure *CodeableReference `json:"diseasesymptomprocedure,omitempty"`
	DiseaseStatus *CodeableReference `json:"diseasestatus,omitempty"`
	Comorbidity []*CodeableReference `json:"comorbidity,omitempty"`
	IntendedEffect *CodeableReference `json:"intendedeffect,omitempty"`
	DurationRange *Range `json:"durationrange,omitempty"`
	DurationString *FhirString `json:"durationstring,omitempty"`
	UndesirableEffect []*Reference `json:"undesirableeffect,omitempty"`
	OtherTherapy []*ClinicalUseDefinitionOtherTherapy `json:"othertherapy,omitempty"`
}

// NewClinicalUseDefinitionIndication creates a new ClinicalUseDefinitionIndication instance.
func NewClinicalUseDefinitionIndication() *ClinicalUseDefinitionIndication {
	return &ClinicalUseDefinitionIndication{}
}

// UnmarshalJSON populates ClinicalUseDefinitionIndication from JSON data.
func (m *ClinicalUseDefinitionIndication) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		DiseaseSymptomProcedure *CodeableReference `json:"diseasesymptomprocedure,omitempty"`
		DiseaseStatus *CodeableReference `json:"diseasestatus,omitempty"`
		Comorbidity []*CodeableReference `json:"comorbidity,omitempty"`
		IntendedEffect *CodeableReference `json:"intendedeffect,omitempty"`
		DurationRange *Range `json:"durationrange,omitempty"`
		DurationString *FhirString `json:"durationstring,omitempty"`
		UndesirableEffect []*Reference `json:"undesirableeffect,omitempty"`
		OtherTherapy []*ClinicalUseDefinitionOtherTherapy `json:"othertherapy,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.DiseaseSymptomProcedure = temp.DiseaseSymptomProcedure
	m.DiseaseStatus = temp.DiseaseStatus
	m.Comorbidity = temp.Comorbidity
	m.IntendedEffect = temp.IntendedEffect
	m.DurationRange = temp.DurationRange
	m.DurationString = temp.DurationString
	m.UndesirableEffect = temp.UndesirableEffect
	m.OtherTherapy = temp.OtherTherapy
	return nil
}

// MarshalJSON converts ClinicalUseDefinitionIndication to JSON data.
func (m *ClinicalUseDefinitionIndication) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		DiseaseSymptomProcedure *CodeableReference `json:"diseasesymptomprocedure,omitempty"`
		DiseaseStatus *CodeableReference `json:"diseasestatus,omitempty"`
		Comorbidity []*CodeableReference `json:"comorbidity,omitempty"`
		IntendedEffect *CodeableReference `json:"intendedeffect,omitempty"`
		DurationRange *Range `json:"durationrange,omitempty"`
		DurationString interface{} `json:"durationstring,omitempty"`
		DurationStringElement map[string]interface{} `json:"_durationstring,omitempty"`
		UndesirableEffect []*Reference `json:"undesirableeffect,omitempty"`
		OtherTherapy []*ClinicalUseDefinitionOtherTherapy `json:"othertherapy,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.DiseaseSymptomProcedure = m.DiseaseSymptomProcedure
	output.DiseaseStatus = m.DiseaseStatus
	output.Comorbidity = m.Comorbidity
	output.IntendedEffect = m.IntendedEffect
	output.DurationRange = m.DurationRange
	if m.DurationString != nil && m.DurationString.Value != nil {
		output.DurationString = m.DurationString.Value
		if m.DurationString.Element != nil {
			output.DurationStringElement = toMapOrNil(m.DurationString.Element.MarshalJSON())
		}
	}
	output.UndesirableEffect = m.UndesirableEffect
	output.OtherTherapy = m.OtherTherapy
	return json.Marshal(output)
}

// Clone creates a deep copy of ClinicalUseDefinitionIndication.
func (m *ClinicalUseDefinitionIndication) Clone() *ClinicalUseDefinitionIndication {
	if m == nil { return nil }
	return &ClinicalUseDefinitionIndication{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		DiseaseSymptomProcedure: m.DiseaseSymptomProcedure.Clone(),
		DiseaseStatus: m.DiseaseStatus.Clone(),
		Comorbidity: cloneSlices(m.Comorbidity),
		IntendedEffect: m.IntendedEffect.Clone(),
		DurationRange: m.DurationRange.Clone(),
		DurationString: m.DurationString.Clone(),
		UndesirableEffect: cloneSlices(m.UndesirableEffect),
		OtherTherapy: cloneSlices(m.OtherTherapy),
	}
}

// Equals checks equality between two ClinicalUseDefinitionIndication instances.
func (m *ClinicalUseDefinitionIndication) Equals(other *ClinicalUseDefinitionIndication) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.DiseaseSymptomProcedure.Equals(other.DiseaseSymptomProcedure) { return false }
	if !m.DiseaseStatus.Equals(other.DiseaseStatus) { return false }
	if !compareSlices(m.Comorbidity, other.Comorbidity) { return false }
	if !m.IntendedEffect.Equals(other.IntendedEffect) { return false }
	if !m.DurationRange.Equals(other.DurationRange) { return false }
	if !m.DurationString.Equals(other.DurationString) { return false }
	if !compareSlices(m.UndesirableEffect, other.UndesirableEffect) { return false }
	if !compareSlices(m.OtherTherapy, other.OtherTherapy) { return false }
	return true
}

// ClinicalUseDefinitionInteraction
// Specifics for when this is an interaction.
type ClinicalUseDefinitionInteraction struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Interactant []*ClinicalUseDefinitionInteractant `json:"interactant,omitempty"`
	Type *CodeableConcept `json:"type,omitempty"`
	Effect *CodeableReference `json:"effect,omitempty"`
	Incidence *CodeableConcept `json:"incidence,omitempty"`
	Management []*CodeableConcept `json:"management,omitempty"`
}

// NewClinicalUseDefinitionInteraction creates a new ClinicalUseDefinitionInteraction instance.
func NewClinicalUseDefinitionInteraction() *ClinicalUseDefinitionInteraction {
	return &ClinicalUseDefinitionInteraction{}
}

// UnmarshalJSON populates ClinicalUseDefinitionInteraction from JSON data.
func (m *ClinicalUseDefinitionInteraction) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Interactant []*ClinicalUseDefinitionInteractant `json:"interactant,omitempty"`
		Type *CodeableConcept `json:"type,omitempty"`
		Effect *CodeableReference `json:"effect,omitempty"`
		Incidence *CodeableConcept `json:"incidence,omitempty"`
		Management []*CodeableConcept `json:"management,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Interactant = temp.Interactant
	m.Type = temp.Type
	m.Effect = temp.Effect
	m.Incidence = temp.Incidence
	m.Management = temp.Management
	return nil
}

// MarshalJSON converts ClinicalUseDefinitionInteraction to JSON data.
func (m *ClinicalUseDefinitionInteraction) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Interactant []*ClinicalUseDefinitionInteractant `json:"interactant,omitempty"`
		Type *CodeableConcept `json:"type,omitempty"`
		Effect *CodeableReference `json:"effect,omitempty"`
		Incidence *CodeableConcept `json:"incidence,omitempty"`
		Management []*CodeableConcept `json:"management,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.Interactant = m.Interactant
	output.Type = m.Type
	output.Effect = m.Effect
	output.Incidence = m.Incidence
	output.Management = m.Management
	return json.Marshal(output)
}

// Clone creates a deep copy of ClinicalUseDefinitionInteraction.
func (m *ClinicalUseDefinitionInteraction) Clone() *ClinicalUseDefinitionInteraction {
	if m == nil { return nil }
	return &ClinicalUseDefinitionInteraction{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Interactant: cloneSlices(m.Interactant),
		Type: m.Type.Clone(),
		Effect: m.Effect.Clone(),
		Incidence: m.Incidence.Clone(),
		Management: cloneSlices(m.Management),
	}
}

// Equals checks equality between two ClinicalUseDefinitionInteraction instances.
func (m *ClinicalUseDefinitionInteraction) Equals(other *ClinicalUseDefinitionInteraction) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !compareSlices(m.Interactant, other.Interactant) { return false }
	if !m.Type.Equals(other.Type) { return false }
	if !m.Effect.Equals(other.Effect) { return false }
	if !m.Incidence.Equals(other.Incidence) { return false }
	if !compareSlices(m.Management, other.Management) { return false }
	return true
}

// ClinicalUseDefinitionInteractant
// The specific medication, food, substance or laboratory test that interacts.
type ClinicalUseDefinitionInteractant struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	ItemReference *Reference `json:"itemreference,omitempty"`
	ItemCodeableConcept *CodeableConcept `json:"itemcodeableconcept,omitempty"`
}

// NewClinicalUseDefinitionInteractant creates a new ClinicalUseDefinitionInteractant instance.
func NewClinicalUseDefinitionInteractant() *ClinicalUseDefinitionInteractant {
	return &ClinicalUseDefinitionInteractant{}
}

// UnmarshalJSON populates ClinicalUseDefinitionInteractant from JSON data.
func (m *ClinicalUseDefinitionInteractant) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		ItemReference *Reference `json:"itemreference,omitempty"`
		ItemCodeableConcept *CodeableConcept `json:"itemcodeableconcept,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.ItemReference = temp.ItemReference
	m.ItemCodeableConcept = temp.ItemCodeableConcept
	return nil
}

// MarshalJSON converts ClinicalUseDefinitionInteractant to JSON data.
func (m *ClinicalUseDefinitionInteractant) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		ItemReference *Reference `json:"itemreference,omitempty"`
		ItemCodeableConcept *CodeableConcept `json:"itemcodeableconcept,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.ItemReference = m.ItemReference
	output.ItemCodeableConcept = m.ItemCodeableConcept
	return json.Marshal(output)
}

// Clone creates a deep copy of ClinicalUseDefinitionInteractant.
func (m *ClinicalUseDefinitionInteractant) Clone() *ClinicalUseDefinitionInteractant {
	if m == nil { return nil }
	return &ClinicalUseDefinitionInteractant{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		ItemReference: m.ItemReference.Clone(),
		ItemCodeableConcept: m.ItemCodeableConcept.Clone(),
	}
}

// Equals checks equality between two ClinicalUseDefinitionInteractant instances.
func (m *ClinicalUseDefinitionInteractant) Equals(other *ClinicalUseDefinitionInteractant) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.ItemReference.Equals(other.ItemReference) { return false }
	if !m.ItemCodeableConcept.Equals(other.ItemCodeableConcept) { return false }
	return true
}

// ClinicalUseDefinitionUndesirableEffect
// Describe the possible undesirable effects (negative outcomes) from the use of the medicinal product as treatment.
type ClinicalUseDefinitionUndesirableEffect struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	SymptomConditionEffect *CodeableReference `json:"symptomconditioneffect,omitempty"`
	Classification *CodeableConcept `json:"classification,omitempty"`
	FrequencyOfOccurrence *CodeableConcept `json:"frequencyofoccurrence,omitempty"`
}

// NewClinicalUseDefinitionUndesirableEffect creates a new ClinicalUseDefinitionUndesirableEffect instance.
func NewClinicalUseDefinitionUndesirableEffect() *ClinicalUseDefinitionUndesirableEffect {
	return &ClinicalUseDefinitionUndesirableEffect{}
}

// UnmarshalJSON populates ClinicalUseDefinitionUndesirableEffect from JSON data.
func (m *ClinicalUseDefinitionUndesirableEffect) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		SymptomConditionEffect *CodeableReference `json:"symptomconditioneffect,omitempty"`
		Classification *CodeableConcept `json:"classification,omitempty"`
		FrequencyOfOccurrence *CodeableConcept `json:"frequencyofoccurrence,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.SymptomConditionEffect = temp.SymptomConditionEffect
	m.Classification = temp.Classification
	m.FrequencyOfOccurrence = temp.FrequencyOfOccurrence
	return nil
}

// MarshalJSON converts ClinicalUseDefinitionUndesirableEffect to JSON data.
func (m *ClinicalUseDefinitionUndesirableEffect) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		SymptomConditionEffect *CodeableReference `json:"symptomconditioneffect,omitempty"`
		Classification *CodeableConcept `json:"classification,omitempty"`
		FrequencyOfOccurrence *CodeableConcept `json:"frequencyofoccurrence,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.SymptomConditionEffect = m.SymptomConditionEffect
	output.Classification = m.Classification
	output.FrequencyOfOccurrence = m.FrequencyOfOccurrence
	return json.Marshal(output)
}

// Clone creates a deep copy of ClinicalUseDefinitionUndesirableEffect.
func (m *ClinicalUseDefinitionUndesirableEffect) Clone() *ClinicalUseDefinitionUndesirableEffect {
	if m == nil { return nil }
	return &ClinicalUseDefinitionUndesirableEffect{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		SymptomConditionEffect: m.SymptomConditionEffect.Clone(),
		Classification: m.Classification.Clone(),
		FrequencyOfOccurrence: m.FrequencyOfOccurrence.Clone(),
	}
}

// Equals checks equality between two ClinicalUseDefinitionUndesirableEffect instances.
func (m *ClinicalUseDefinitionUndesirableEffect) Equals(other *ClinicalUseDefinitionUndesirableEffect) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.SymptomConditionEffect.Equals(other.SymptomConditionEffect) { return false }
	if !m.Classification.Equals(other.Classification) { return false }
	if !m.FrequencyOfOccurrence.Equals(other.FrequencyOfOccurrence) { return false }
	return true
}

// ClinicalUseDefinitionWarning
// A critical piece of information about environmental, health or physical risks or hazards that serve as caution to the user. For example 'Do not operate heavy machinery', 'May cause drowsiness', or 'Get medical advice/attention if you feel unwell'.
type ClinicalUseDefinitionWarning struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Description *FhirMarkdown `json:"description,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
}

// NewClinicalUseDefinitionWarning creates a new ClinicalUseDefinitionWarning instance.
func NewClinicalUseDefinitionWarning() *ClinicalUseDefinitionWarning {
	return &ClinicalUseDefinitionWarning{}
}

// UnmarshalJSON populates ClinicalUseDefinitionWarning from JSON data.
func (m *ClinicalUseDefinitionWarning) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Description *FhirMarkdown `json:"description,omitempty"`
		Code *CodeableConcept `json:"code,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Description = temp.Description
	m.Code = temp.Code
	return nil
}

// MarshalJSON converts ClinicalUseDefinitionWarning to JSON data.
func (m *ClinicalUseDefinitionWarning) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Description interface{} `json:"description,omitempty"`
		DescriptionElement map[string]interface{} `json:"_description,omitempty"`
		Code *CodeableConcept `json:"code,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	if m.Description != nil && m.Description.Value != nil {
		output.Description = m.Description.Value
		if m.Description.Element != nil {
			output.DescriptionElement = toMapOrNil(m.Description.Element.MarshalJSON())
		}
	}
	output.Code = m.Code
	return json.Marshal(output)
}

// Clone creates a deep copy of ClinicalUseDefinitionWarning.
func (m *ClinicalUseDefinitionWarning) Clone() *ClinicalUseDefinitionWarning {
	if m == nil { return nil }
	return &ClinicalUseDefinitionWarning{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Description: m.Description.Clone(),
		Code: m.Code.Clone(),
	}
}

// Equals checks equality between two ClinicalUseDefinitionWarning instances.
func (m *ClinicalUseDefinitionWarning) Equals(other *ClinicalUseDefinitionWarning) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Description.Equals(other.Description) { return false }
	if !m.Code.Equals(other.Code) { return false }
	return true
}

