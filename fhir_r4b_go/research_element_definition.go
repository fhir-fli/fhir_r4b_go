// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json")

// ResearchElementDefinition
// The ResearchElementDefinition resource describes a "PICO" element that knowledge (evidence, assertion, recommendation) is about.
type ResearchElementDefinition struct {
	DomainResource
	Id *FhirString `json:"id,omitempty"`
	Meta *FhirMeta `json:"meta,omitempty"`
	ImplicitRules *FhirUri `json:"implicitrules,omitempty"`
	Language *CommonLanguages `json:"language,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Contained []*Resource `json:"contained,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Url *FhirUri `json:"url,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Version *FhirString `json:"version,omitempty"`
	Name *FhirString `json:"name,omitempty"`
	Title *FhirString `json:"title,omitempty"`
	ShortTitle *FhirString `json:"shorttitle,omitempty"`
	Subtitle *FhirString `json:"subtitle,omitempty"`
	Status *PublicationStatus `json:"status,omitempty"`
	Experimental *FhirBoolean `json:"experimental,omitempty"`
	SubjectCodeableConcept *CodeableConcept `json:"subjectcodeableconcept,omitempty"`
	SubjectReference *Reference `json:"subjectreference,omitempty"`
	Date *FhirDateTime `json:"date,omitempty"`
	Publisher *FhirString `json:"publisher,omitempty"`
	Contact []*ContactDetail `json:"contact,omitempty"`
	Description *FhirMarkdown `json:"description,omitempty"`
	Comment []*FhirString `json:"comment,omitempty"`
	UseContext []*UsageContext `json:"usecontext,omitempty"`
	Jurisdiction []*CodeableConcept `json:"jurisdiction,omitempty"`
	Purpose *FhirMarkdown `json:"purpose,omitempty"`
	Usage *FhirString `json:"usage,omitempty"`
	Copyright *FhirMarkdown `json:"copyright,omitempty"`
	ApprovalDate *FhirDate `json:"approvaldate,omitempty"`
	LastReviewDate *FhirDate `json:"lastreviewdate,omitempty"`
	EffectivePeriod *Period `json:"effectiveperiod,omitempty"`
	Topic []*CodeableConcept `json:"topic,omitempty"`
	Author []*ContactDetail `json:"author,omitempty"`
	Editor []*ContactDetail `json:"editor,omitempty"`
	Reviewer []*ContactDetail `json:"reviewer,omitempty"`
	Endorser []*ContactDetail `json:"endorser,omitempty"`
	RelatedArtifact []*RelatedArtifact `json:"relatedartifact,omitempty"`
	Library_ []*FhirCanonical `json:"library,omitempty"`
	Type *ResearchElementType `json:"type,omitempty"`
	VariableType *VariableType `json:"variabletype,omitempty"`
	Characteristic []*ResearchElementDefinitionCharacteristic `json:"characteristic,omitempty"`
}

// NewResearchElementDefinition creates a new ResearchElementDefinition instance
func NewResearchElementDefinition() *ResearchElementDefinition {
	return &ResearchElementDefinition{}
}

// FromJSON populates ResearchElementDefinition from JSON data
func (m *ResearchElementDefinition) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts ResearchElementDefinition to JSON data
func (m *ResearchElementDefinition) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of ResearchElementDefinition
func (m *ResearchElementDefinition) Clone() *ResearchElementDefinition {
	if m == nil { return nil }
	return &ResearchElementDefinition{
		Id: m.Id.Clone(),
		Meta: m.Meta.Clone(),
		ImplicitRules: m.ImplicitRules.Clone(),
		Language: m.Language.Clone(),
		Text: m.Text.Clone(),
		Contained: cloneSlices(m.Contained),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Url: m.Url.Clone(),
		Identifier: cloneSlices(m.Identifier),
		Version: m.Version.Clone(),
		Name: m.Name.Clone(),
		Title: m.Title.Clone(),
		ShortTitle: m.ShortTitle.Clone(),
		Subtitle: m.Subtitle.Clone(),
		Status: m.Status.Clone(),
		Experimental: m.Experimental.Clone(),
		SubjectCodeableConcept: m.SubjectCodeableConcept.Clone(),
		SubjectReference: m.SubjectReference.Clone(),
		Date: m.Date.Clone(),
		Publisher: m.Publisher.Clone(),
		Contact: cloneSlices(m.Contact),
		Description: m.Description.Clone(),
		Comment: cloneSlices(m.Comment),
		UseContext: cloneSlices(m.UseContext),
		Jurisdiction: cloneSlices(m.Jurisdiction),
		Purpose: m.Purpose.Clone(),
		Usage: m.Usage.Clone(),
		Copyright: m.Copyright.Clone(),
		ApprovalDate: m.ApprovalDate.Clone(),
		LastReviewDate: m.LastReviewDate.Clone(),
		EffectivePeriod: m.EffectivePeriod.Clone(),
		Topic: cloneSlices(m.Topic),
		Author: cloneSlices(m.Author),
		Editor: cloneSlices(m.Editor),
		Reviewer: cloneSlices(m.Reviewer),
		Endorser: cloneSlices(m.Endorser),
		RelatedArtifact: cloneSlices(m.RelatedArtifact),
		Library_: cloneSlices(m.Library_),
		Type: m.Type.Clone(),
		VariableType: m.VariableType.Clone(),
		Characteristic: cloneSlices(m.Characteristic),
	}
}

// Equals checks for equality with another ResearchElementDefinition instance
func (m *ResearchElementDefinition) Equals(other *ResearchElementDefinition) bool {
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
	if !m.Url.Equals(other.Url) { return false }
	if !compareSlices(m.Identifier, other.Identifier) { return false }
	if !m.Version.Equals(other.Version) { return false }
	if !m.Name.Equals(other.Name) { return false }
	if !m.Title.Equals(other.Title) { return false }
	if !m.ShortTitle.Equals(other.ShortTitle) { return false }
	if !m.Subtitle.Equals(other.Subtitle) { return false }
	if !m.Status.Equals(other.Status) { return false }
	if !m.Experimental.Equals(other.Experimental) { return false }
	if !m.SubjectCodeableConcept.Equals(other.SubjectCodeableConcept) { return false }
	if !m.SubjectReference.Equals(other.SubjectReference) { return false }
	if !m.Date.Equals(other.Date) { return false }
	if !m.Publisher.Equals(other.Publisher) { return false }
	if !compareSlices(m.Contact, other.Contact) { return false }
	if !m.Description.Equals(other.Description) { return false }
	if !compareSlices(m.Comment, other.Comment) { return false }
	if !compareSlices(m.UseContext, other.UseContext) { return false }
	if !compareSlices(m.Jurisdiction, other.Jurisdiction) { return false }
	if !m.Purpose.Equals(other.Purpose) { return false }
	if !m.Usage.Equals(other.Usage) { return false }
	if !m.Copyright.Equals(other.Copyright) { return false }
	if !m.ApprovalDate.Equals(other.ApprovalDate) { return false }
	if !m.LastReviewDate.Equals(other.LastReviewDate) { return false }
	if !m.EffectivePeriod.Equals(other.EffectivePeriod) { return false }
	if !compareSlices(m.Topic, other.Topic) { return false }
	if !compareSlices(m.Author, other.Author) { return false }
	if !compareSlices(m.Editor, other.Editor) { return false }
	if !compareSlices(m.Reviewer, other.Reviewer) { return false }
	if !compareSlices(m.Endorser, other.Endorser) { return false }
	if !compareSlices(m.RelatedArtifact, other.RelatedArtifact) { return false }
	if !compareSlices(m.Library_, other.Library_) { return false }
	if !m.Type.Equals(other.Type) { return false }
	if !m.VariableType.Equals(other.VariableType) { return false }
	if !compareSlices(m.Characteristic, other.Characteristic) { return false }
	return true
}

// ResearchElementDefinitionCharacteristic
// A characteristic that defines the members of the research element. Multiple characteristics are applied with "and" semantics.
type ResearchElementDefinitionCharacteristic struct {
	BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	DefinitionCodeableConcept *CodeableConcept `json:"definitioncodeableconcept,omitempty"`
	DefinitionCanonical *FhirCanonical `json:"definitioncanonical,omitempty"`
	DefinitionExpression *FhirExpression `json:"definitionexpression,omitempty"`
	DefinitionDataRequirement *DataRequirement `json:"definitiondatarequirement,omitempty"`
	UsageContext []*UsageContext `json:"usagecontext,omitempty"`
	Exclude *FhirBoolean `json:"exclude,omitempty"`
	UnitOfMeasure *CodeableConcept `json:"unitofmeasure,omitempty"`
	StudyEffectiveDescription *FhirString `json:"studyeffectivedescription,omitempty"`
	StudyEffectiveDateTime *FhirDateTime `json:"studyeffectivedatetime,omitempty"`
	StudyEffectivePeriod *Period `json:"studyeffectiveperiod,omitempty"`
	StudyEffectiveDuration *FhirDuration `json:"studyeffectiveduration,omitempty"`
	StudyEffectiveTiming *Timing `json:"studyeffectivetiming,omitempty"`
	StudyEffectiveTimeFromStart *FhirDuration `json:"studyeffectivetimefromstart,omitempty"`
	StudyEffectiveGroupMeasure *GroupMeasure `json:"studyeffectivegroupmeasure,omitempty"`
	ParticipantEffectiveDescription *FhirString `json:"participanteffectivedescription,omitempty"`
	ParticipantEffectiveDateTime *FhirDateTime `json:"participanteffectivedatetime,omitempty"`
	ParticipantEffectivePeriod *Period `json:"participanteffectiveperiod,omitempty"`
	ParticipantEffectiveDuration *FhirDuration `json:"participanteffectiveduration,omitempty"`
	ParticipantEffectiveTiming *Timing `json:"participanteffectivetiming,omitempty"`
	ParticipantEffectiveTimeFromStart *FhirDuration `json:"participanteffectivetimefromstart,omitempty"`
	ParticipantEffectiveGroupMeasure *GroupMeasure `json:"participanteffectivegroupmeasure,omitempty"`
}

// NewResearchElementDefinitionCharacteristic creates a new ResearchElementDefinitionCharacteristic instance
func NewResearchElementDefinitionCharacteristic() *ResearchElementDefinitionCharacteristic {
	return &ResearchElementDefinitionCharacteristic{}
}

// FromJSON populates ResearchElementDefinitionCharacteristic from JSON data
func (m *ResearchElementDefinitionCharacteristic) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts ResearchElementDefinitionCharacteristic to JSON data
func (m *ResearchElementDefinitionCharacteristic) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of ResearchElementDefinitionCharacteristic
func (m *ResearchElementDefinitionCharacteristic) Clone() *ResearchElementDefinitionCharacteristic {
	if m == nil { return nil }
	return &ResearchElementDefinitionCharacteristic{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		DefinitionCodeableConcept: m.DefinitionCodeableConcept.Clone(),
		DefinitionCanonical: m.DefinitionCanonical.Clone(),
		DefinitionExpression: m.DefinitionExpression.Clone(),
		DefinitionDataRequirement: m.DefinitionDataRequirement.Clone(),
		UsageContext: cloneSlices(m.UsageContext),
		Exclude: m.Exclude.Clone(),
		UnitOfMeasure: m.UnitOfMeasure.Clone(),
		StudyEffectiveDescription: m.StudyEffectiveDescription.Clone(),
		StudyEffectiveDateTime: m.StudyEffectiveDateTime.Clone(),
		StudyEffectivePeriod: m.StudyEffectivePeriod.Clone(),
		StudyEffectiveDuration: m.StudyEffectiveDuration.Clone(),
		StudyEffectiveTiming: m.StudyEffectiveTiming.Clone(),
		StudyEffectiveTimeFromStart: m.StudyEffectiveTimeFromStart.Clone(),
		StudyEffectiveGroupMeasure: m.StudyEffectiveGroupMeasure.Clone(),
		ParticipantEffectiveDescription: m.ParticipantEffectiveDescription.Clone(),
		ParticipantEffectiveDateTime: m.ParticipantEffectiveDateTime.Clone(),
		ParticipantEffectivePeriod: m.ParticipantEffectivePeriod.Clone(),
		ParticipantEffectiveDuration: m.ParticipantEffectiveDuration.Clone(),
		ParticipantEffectiveTiming: m.ParticipantEffectiveTiming.Clone(),
		ParticipantEffectiveTimeFromStart: m.ParticipantEffectiveTimeFromStart.Clone(),
		ParticipantEffectiveGroupMeasure: m.ParticipantEffectiveGroupMeasure.Clone(),
	}
}

// Equals checks for equality with another ResearchElementDefinitionCharacteristic instance
func (m *ResearchElementDefinitionCharacteristic) Equals(other *ResearchElementDefinitionCharacteristic) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.DefinitionCodeableConcept.Equals(other.DefinitionCodeableConcept) { return false }
	if !m.DefinitionCanonical.Equals(other.DefinitionCanonical) { return false }
	if !m.DefinitionExpression.Equals(other.DefinitionExpression) { return false }
	if !m.DefinitionDataRequirement.Equals(other.DefinitionDataRequirement) { return false }
	if !compareSlices(m.UsageContext, other.UsageContext) { return false }
	if !m.Exclude.Equals(other.Exclude) { return false }
	if !m.UnitOfMeasure.Equals(other.UnitOfMeasure) { return false }
	if !m.StudyEffectiveDescription.Equals(other.StudyEffectiveDescription) { return false }
	if !m.StudyEffectiveDateTime.Equals(other.StudyEffectiveDateTime) { return false }
	if !m.StudyEffectivePeriod.Equals(other.StudyEffectivePeriod) { return false }
	if !m.StudyEffectiveDuration.Equals(other.StudyEffectiveDuration) { return false }
	if !m.StudyEffectiveTiming.Equals(other.StudyEffectiveTiming) { return false }
	if !m.StudyEffectiveTimeFromStart.Equals(other.StudyEffectiveTimeFromStart) { return false }
	if !m.StudyEffectiveGroupMeasure.Equals(other.StudyEffectiveGroupMeasure) { return false }
	if !m.ParticipantEffectiveDescription.Equals(other.ParticipantEffectiveDescription) { return false }
	if !m.ParticipantEffectiveDateTime.Equals(other.ParticipantEffectiveDateTime) { return false }
	if !m.ParticipantEffectivePeriod.Equals(other.ParticipantEffectivePeriod) { return false }
	if !m.ParticipantEffectiveDuration.Equals(other.ParticipantEffectiveDuration) { return false }
	if !m.ParticipantEffectiveTiming.Equals(other.ParticipantEffectiveTiming) { return false }
	if !m.ParticipantEffectiveTimeFromStart.Equals(other.ParticipantEffectiveTimeFromStart) { return false }
	if !m.ParticipantEffectiveGroupMeasure.Equals(other.ParticipantEffectiveGroupMeasure) { return false }
	return true
}
