// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json")

// PlanDefinition
// This resource allows for the definition of various types of plans as a sharable, consumable, and executable artifact. The resource is general enough to support the description of a broad range of clinical and non-clinical artifacts such as clinical decision support rules, order sets, protocols, and drug quality specifications.
type PlanDefinition struct {
	CanonicalResource
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
	Subtitle *FhirString `json:"subtitle,omitempty"`
	Type *CodeableConcept `json:"type,omitempty"`
	Status *PublicationStatus `json:"status,omitempty"`
	Experimental *FhirBoolean `json:"experimental,omitempty"`
	SubjectCodeableConcept *CodeableConcept `json:"subjectcodeableconcept,omitempty"`
	SubjectReference *Reference `json:"subjectreference,omitempty"`
	SubjectCanonical *SubjectType `json:"subjectcanonical,omitempty"`
	Date *FhirDateTime `json:"date,omitempty"`
	Publisher *FhirString `json:"publisher,omitempty"`
	Contact []*ContactDetail `json:"contact,omitempty"`
	Description *FhirMarkdown `json:"description,omitempty"`
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
	Goal []*PlanDefinitionGoal `json:"goal,omitempty"`
	Action []*PlanDefinitionAction `json:"action,omitempty"`
}

// NewPlanDefinition creates a new PlanDefinition instance
func NewPlanDefinition() *PlanDefinition {
	return &PlanDefinition{}
}

// FromJSON populates PlanDefinition from JSON data
func (m *PlanDefinition) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts PlanDefinition to JSON data
func (m *PlanDefinition) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of PlanDefinition
func (m *PlanDefinition) Clone() *PlanDefinition {
	if m == nil { return nil }
	return &PlanDefinition{
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
		Subtitle: m.Subtitle.Clone(),
		Type: m.Type.Clone(),
		Status: m.Status.Clone(),
		Experimental: m.Experimental.Clone(),
		SubjectCodeableConcept: m.SubjectCodeableConcept.Clone(),
		SubjectReference: m.SubjectReference.Clone(),
		SubjectCanonical: m.SubjectCanonical.Clone(),
		Date: m.Date.Clone(),
		Publisher: m.Publisher.Clone(),
		Contact: cloneSlices(m.Contact),
		Description: m.Description.Clone(),
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
		Goal: cloneSlices(m.Goal),
		Action: cloneSlices(m.Action),
	}
}

// Equals checks for equality with another PlanDefinition instance
func (m *PlanDefinition) Equals(other *PlanDefinition) bool {
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
	if !m.Subtitle.Equals(other.Subtitle) { return false }
	if !m.Type.Equals(other.Type) { return false }
	if !m.Status.Equals(other.Status) { return false }
	if !m.Experimental.Equals(other.Experimental) { return false }
	if !m.SubjectCodeableConcept.Equals(other.SubjectCodeableConcept) { return false }
	if !m.SubjectReference.Equals(other.SubjectReference) { return false }
	if !m.SubjectCanonical.Equals(other.SubjectCanonical) { return false }
	if !m.Date.Equals(other.Date) { return false }
	if !m.Publisher.Equals(other.Publisher) { return false }
	if !compareSlices(m.Contact, other.Contact) { return false }
	if !m.Description.Equals(other.Description) { return false }
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
	if !compareSlices(m.Goal, other.Goal) { return false }
	if !compareSlices(m.Action, other.Action) { return false }
	return true
}

// PlanDefinitionGoal
// A goal describes an expected outcome that activities within the plan are intended to achieve. For example, weight loss, restoring an activity of daily living, obtaining herd immunity via immunization, meeting a process improvement objective, meeting the acceptance criteria for a test as specified by a quality specification, etc.
type PlanDefinitionGoal struct {
	BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Category *CodeableConcept `json:"category,omitempty"`
	Description *CodeableConcept `json:"description,omitempty"`
	Priority *CodeableConcept `json:"priority,omitempty"`
	Start *CodeableConcept `json:"start,omitempty"`
	Addresses []*CodeableConcept `json:"addresses,omitempty"`
	Documentation []*RelatedArtifact `json:"documentation,omitempty"`
	Target []*PlanDefinitionTarget `json:"target,omitempty"`
}

// NewPlanDefinitionGoal creates a new PlanDefinitionGoal instance
func NewPlanDefinitionGoal() *PlanDefinitionGoal {
	return &PlanDefinitionGoal{}
}

// FromJSON populates PlanDefinitionGoal from JSON data
func (m *PlanDefinitionGoal) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts PlanDefinitionGoal to JSON data
func (m *PlanDefinitionGoal) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of PlanDefinitionGoal
func (m *PlanDefinitionGoal) Clone() *PlanDefinitionGoal {
	if m == nil { return nil }
	return &PlanDefinitionGoal{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Category: m.Category.Clone(),
		Description: m.Description.Clone(),
		Priority: m.Priority.Clone(),
		Start: m.Start.Clone(),
		Addresses: cloneSlices(m.Addresses),
		Documentation: cloneSlices(m.Documentation),
		Target: cloneSlices(m.Target),
	}
}

// Equals checks for equality with another PlanDefinitionGoal instance
func (m *PlanDefinitionGoal) Equals(other *PlanDefinitionGoal) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Category.Equals(other.Category) { return false }
	if !m.Description.Equals(other.Description) { return false }
	if !m.Priority.Equals(other.Priority) { return false }
	if !m.Start.Equals(other.Start) { return false }
	if !compareSlices(m.Addresses, other.Addresses) { return false }
	if !compareSlices(m.Documentation, other.Documentation) { return false }
	if !compareSlices(m.Target, other.Target) { return false }
	return true
}

// PlanDefinitionTarget
// Indicates what should be done and within what timeframe.
type PlanDefinitionTarget struct {
	BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Measure *CodeableConcept `json:"measure,omitempty"`
	DetailQuantity *Quantity `json:"detailquantity,omitempty"`
	DetailRange *Range `json:"detailrange,omitempty"`
	DetailCodeableConcept *CodeableConcept `json:"detailcodeableconcept,omitempty"`
	Due *FhirDuration `json:"due,omitempty"`
}

// NewPlanDefinitionTarget creates a new PlanDefinitionTarget instance
func NewPlanDefinitionTarget() *PlanDefinitionTarget {
	return &PlanDefinitionTarget{}
}

// FromJSON populates PlanDefinitionTarget from JSON data
func (m *PlanDefinitionTarget) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts PlanDefinitionTarget to JSON data
func (m *PlanDefinitionTarget) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of PlanDefinitionTarget
func (m *PlanDefinitionTarget) Clone() *PlanDefinitionTarget {
	if m == nil { return nil }
	return &PlanDefinitionTarget{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Measure: m.Measure.Clone(),
		DetailQuantity: m.DetailQuantity.Clone(),
		DetailRange: m.DetailRange.Clone(),
		DetailCodeableConcept: m.DetailCodeableConcept.Clone(),
		Due: m.Due.Clone(),
	}
}

// Equals checks for equality with another PlanDefinitionTarget instance
func (m *PlanDefinitionTarget) Equals(other *PlanDefinitionTarget) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Measure.Equals(other.Measure) { return false }
	if !m.DetailQuantity.Equals(other.DetailQuantity) { return false }
	if !m.DetailRange.Equals(other.DetailRange) { return false }
	if !m.DetailCodeableConcept.Equals(other.DetailCodeableConcept) { return false }
	if !m.Due.Equals(other.Due) { return false }
	return true
}

// PlanDefinitionAction
// An action or group of actions to be taken as part of the plan. For example, in clinical care, an action would be to prescribe a particular indicated medication, or perform a particular test as appropriate. In pharmaceutical quality, an action would be the test that needs to be performed on a drug product as defined in the quality specification.
type PlanDefinitionAction struct {
	BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Prefix *FhirString `json:"prefix,omitempty"`
	Title *FhirString `json:"title,omitempty"`
	Description *FhirString `json:"description,omitempty"`
	TextEquivalent *FhirString `json:"textequivalent,omitempty"`
	Priority *RequestPriority `json:"priority,omitempty"`
	Code []*CodeableConcept `json:"code,omitempty"`
	Reason []*CodeableConcept `json:"reason,omitempty"`
	Documentation []*RelatedArtifact `json:"documentation,omitempty"`
	GoalId []*FhirId `json:"goalid,omitempty"`
	SubjectCodeableConcept *CodeableConcept `json:"subjectcodeableconcept,omitempty"`
	SubjectReference *Reference `json:"subjectreference,omitempty"`
	SubjectCanonical *SubjectType `json:"subjectcanonical,omitempty"`
	Trigger []*TriggerDefinition `json:"trigger,omitempty"`
	Condition []*PlanDefinitionCondition `json:"condition,omitempty"`
	Input []*DataRequirement `json:"input,omitempty"`
	Output []*DataRequirement `json:"output,omitempty"`
	RelatedAction []*PlanDefinitionRelatedAction `json:"relatedaction,omitempty"`
	TimingDateTime *FhirDateTime `json:"timingdatetime,omitempty"`
	TimingAge *Age `json:"timingage,omitempty"`
	TimingPeriod *Period `json:"timingperiod,omitempty"`
	TimingDuration *FhirDuration `json:"timingduration,omitempty"`
	TimingRange *Range `json:"timingrange,omitempty"`
	TimingTiming *Timing `json:"timingtiming,omitempty"`
	Participant []*PlanDefinitionParticipant `json:"participant,omitempty"`
	Type *CodeableConcept `json:"type,omitempty"`
	GroupingBehavior *ActionGroupingBehavior `json:"groupingbehavior,omitempty"`
	SelectionBehavior *ActionSelectionBehavior `json:"selectionbehavior,omitempty"`
	RequiredBehavior *ActionRequiredBehavior `json:"requiredbehavior,omitempty"`
	PrecheckBehavior *ActionPrecheckBehavior `json:"precheckbehavior,omitempty"`
	CardinalityBehavior *ActionCardinalityBehavior `json:"cardinalitybehavior,omitempty"`
	DefinitionCanonical *FhirCanonical `json:"definitioncanonical,omitempty"`
	DefinitionUri *FhirUri `json:"definitionuri,omitempty"`
	Transform *FhirCanonical `json:"transform,omitempty"`
	DynamicValue []*PlanDefinitionDynamicValue `json:"dynamicvalue,omitempty"`
	Action []*PlanDefinitionAction `json:"action,omitempty"`
}

// NewPlanDefinitionAction creates a new PlanDefinitionAction instance
func NewPlanDefinitionAction() *PlanDefinitionAction {
	return &PlanDefinitionAction{}
}

// FromJSON populates PlanDefinitionAction from JSON data
func (m *PlanDefinitionAction) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts PlanDefinitionAction to JSON data
func (m *PlanDefinitionAction) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of PlanDefinitionAction
func (m *PlanDefinitionAction) Clone() *PlanDefinitionAction {
	if m == nil { return nil }
	return &PlanDefinitionAction{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Prefix: m.Prefix.Clone(),
		Title: m.Title.Clone(),
		Description: m.Description.Clone(),
		TextEquivalent: m.TextEquivalent.Clone(),
		Priority: m.Priority.Clone(),
		Code: cloneSlices(m.Code),
		Reason: cloneSlices(m.Reason),
		Documentation: cloneSlices(m.Documentation),
		GoalId: cloneSlices(m.GoalId),
		SubjectCodeableConcept: m.SubjectCodeableConcept.Clone(),
		SubjectReference: m.SubjectReference.Clone(),
		SubjectCanonical: m.SubjectCanonical.Clone(),
		Trigger: cloneSlices(m.Trigger),
		Condition: cloneSlices(m.Condition),
		Input: cloneSlices(m.Input),
		Output: cloneSlices(m.Output),
		RelatedAction: cloneSlices(m.RelatedAction),
		TimingDateTime: m.TimingDateTime.Clone(),
		TimingAge: m.TimingAge.Clone(),
		TimingPeriod: m.TimingPeriod.Clone(),
		TimingDuration: m.TimingDuration.Clone(),
		TimingRange: m.TimingRange.Clone(),
		TimingTiming: m.TimingTiming.Clone(),
		Participant: cloneSlices(m.Participant),
		Type: m.Type.Clone(),
		GroupingBehavior: m.GroupingBehavior.Clone(),
		SelectionBehavior: m.SelectionBehavior.Clone(),
		RequiredBehavior: m.RequiredBehavior.Clone(),
		PrecheckBehavior: m.PrecheckBehavior.Clone(),
		CardinalityBehavior: m.CardinalityBehavior.Clone(),
		DefinitionCanonical: m.DefinitionCanonical.Clone(),
		DefinitionUri: m.DefinitionUri.Clone(),
		Transform: m.Transform.Clone(),
		DynamicValue: cloneSlices(m.DynamicValue),
		Action: cloneSlices(m.Action),
	}
}

// Equals checks for equality with another PlanDefinitionAction instance
func (m *PlanDefinitionAction) Equals(other *PlanDefinitionAction) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Prefix.Equals(other.Prefix) { return false }
	if !m.Title.Equals(other.Title) { return false }
	if !m.Description.Equals(other.Description) { return false }
	if !m.TextEquivalent.Equals(other.TextEquivalent) { return false }
	if !m.Priority.Equals(other.Priority) { return false }
	if !compareSlices(m.Code, other.Code) { return false }
	if !compareSlices(m.Reason, other.Reason) { return false }
	if !compareSlices(m.Documentation, other.Documentation) { return false }
	if !compareSlices(m.GoalId, other.GoalId) { return false }
	if !m.SubjectCodeableConcept.Equals(other.SubjectCodeableConcept) { return false }
	if !m.SubjectReference.Equals(other.SubjectReference) { return false }
	if !m.SubjectCanonical.Equals(other.SubjectCanonical) { return false }
	if !compareSlices(m.Trigger, other.Trigger) { return false }
	if !compareSlices(m.Condition, other.Condition) { return false }
	if !compareSlices(m.Input, other.Input) { return false }
	if !compareSlices(m.Output, other.Output) { return false }
	if !compareSlices(m.RelatedAction, other.RelatedAction) { return false }
	if !m.TimingDateTime.Equals(other.TimingDateTime) { return false }
	if !m.TimingAge.Equals(other.TimingAge) { return false }
	if !m.TimingPeriod.Equals(other.TimingPeriod) { return false }
	if !m.TimingDuration.Equals(other.TimingDuration) { return false }
	if !m.TimingRange.Equals(other.TimingRange) { return false }
	if !m.TimingTiming.Equals(other.TimingTiming) { return false }
	if !compareSlices(m.Participant, other.Participant) { return false }
	if !m.Type.Equals(other.Type) { return false }
	if !m.GroupingBehavior.Equals(other.GroupingBehavior) { return false }
	if !m.SelectionBehavior.Equals(other.SelectionBehavior) { return false }
	if !m.RequiredBehavior.Equals(other.RequiredBehavior) { return false }
	if !m.PrecheckBehavior.Equals(other.PrecheckBehavior) { return false }
	if !m.CardinalityBehavior.Equals(other.CardinalityBehavior) { return false }
	if !m.DefinitionCanonical.Equals(other.DefinitionCanonical) { return false }
	if !m.DefinitionUri.Equals(other.DefinitionUri) { return false }
	if !m.Transform.Equals(other.Transform) { return false }
	if !compareSlices(m.DynamicValue, other.DynamicValue) { return false }
	if !compareSlices(m.Action, other.Action) { return false }
	return true
}

// PlanDefinitionCondition
// An expression that describes applicability criteria or start/stop conditions for the action.
type PlanDefinitionCondition struct {
	BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Kind *ActionConditionKind `json:"kind,omitempty"`
	Expression *FhirExpression `json:"expression,omitempty"`
}

// NewPlanDefinitionCondition creates a new PlanDefinitionCondition instance
func NewPlanDefinitionCondition() *PlanDefinitionCondition {
	return &PlanDefinitionCondition{}
}

// FromJSON populates PlanDefinitionCondition from JSON data
func (m *PlanDefinitionCondition) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts PlanDefinitionCondition to JSON data
func (m *PlanDefinitionCondition) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of PlanDefinitionCondition
func (m *PlanDefinitionCondition) Clone() *PlanDefinitionCondition {
	if m == nil { return nil }
	return &PlanDefinitionCondition{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Kind: m.Kind.Clone(),
		Expression: m.Expression.Clone(),
	}
}

// Equals checks for equality with another PlanDefinitionCondition instance
func (m *PlanDefinitionCondition) Equals(other *PlanDefinitionCondition) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Kind.Equals(other.Kind) { return false }
	if !m.Expression.Equals(other.Expression) { return false }
	return true
}

// PlanDefinitionRelatedAction
// A relationship to another action such as "before" or "30-60 minutes after start of".
type PlanDefinitionRelatedAction struct {
	BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	ActionId *FhirId `json:"actionid,omitempty"`
	Relationship *ActionRelationshipType `json:"relationship,omitempty"`
	OffsetDuration *FhirDuration `json:"offsetduration,omitempty"`
	OffsetRange *Range `json:"offsetrange,omitempty"`
}

// NewPlanDefinitionRelatedAction creates a new PlanDefinitionRelatedAction instance
func NewPlanDefinitionRelatedAction() *PlanDefinitionRelatedAction {
	return &PlanDefinitionRelatedAction{}
}

// FromJSON populates PlanDefinitionRelatedAction from JSON data
func (m *PlanDefinitionRelatedAction) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts PlanDefinitionRelatedAction to JSON data
func (m *PlanDefinitionRelatedAction) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of PlanDefinitionRelatedAction
func (m *PlanDefinitionRelatedAction) Clone() *PlanDefinitionRelatedAction {
	if m == nil { return nil }
	return &PlanDefinitionRelatedAction{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		ActionId: m.ActionId.Clone(),
		Relationship: m.Relationship.Clone(),
		OffsetDuration: m.OffsetDuration.Clone(),
		OffsetRange: m.OffsetRange.Clone(),
	}
}

// Equals checks for equality with another PlanDefinitionRelatedAction instance
func (m *PlanDefinitionRelatedAction) Equals(other *PlanDefinitionRelatedAction) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.ActionId.Equals(other.ActionId) { return false }
	if !m.Relationship.Equals(other.Relationship) { return false }
	if !m.OffsetDuration.Equals(other.OffsetDuration) { return false }
	if !m.OffsetRange.Equals(other.OffsetRange) { return false }
	return true
}

// PlanDefinitionParticipant
// Indicates who should participate in performing the action described.
type PlanDefinitionParticipant struct {
	BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Type *ActionParticipantType `json:"type,omitempty"`
	Role *CodeableConcept `json:"role,omitempty"`
}

// NewPlanDefinitionParticipant creates a new PlanDefinitionParticipant instance
func NewPlanDefinitionParticipant() *PlanDefinitionParticipant {
	return &PlanDefinitionParticipant{}
}

// FromJSON populates PlanDefinitionParticipant from JSON data
func (m *PlanDefinitionParticipant) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts PlanDefinitionParticipant to JSON data
func (m *PlanDefinitionParticipant) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of PlanDefinitionParticipant
func (m *PlanDefinitionParticipant) Clone() *PlanDefinitionParticipant {
	if m == nil { return nil }
	return &PlanDefinitionParticipant{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Type: m.Type.Clone(),
		Role: m.Role.Clone(),
	}
}

// Equals checks for equality with another PlanDefinitionParticipant instance
func (m *PlanDefinitionParticipant) Equals(other *PlanDefinitionParticipant) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Type.Equals(other.Type) { return false }
	if !m.Role.Equals(other.Role) { return false }
	return true
}

// PlanDefinitionDynamicValue
// Customizations that should be applied to the statically defined resource. For example, if the dosage of a medication must be computed based on the patient's weight, a customization would be used to specify an expression that calculated the weight, and the path on the resource that would contain the result.
type PlanDefinitionDynamicValue struct {
	BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Path *FhirString `json:"path,omitempty"`
	Expression *FhirExpression `json:"expression,omitempty"`
}

// NewPlanDefinitionDynamicValue creates a new PlanDefinitionDynamicValue instance
func NewPlanDefinitionDynamicValue() *PlanDefinitionDynamicValue {
	return &PlanDefinitionDynamicValue{}
}

// FromJSON populates PlanDefinitionDynamicValue from JSON data
func (m *PlanDefinitionDynamicValue) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts PlanDefinitionDynamicValue to JSON data
func (m *PlanDefinitionDynamicValue) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of PlanDefinitionDynamicValue
func (m *PlanDefinitionDynamicValue) Clone() *PlanDefinitionDynamicValue {
	if m == nil { return nil }
	return &PlanDefinitionDynamicValue{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Path: m.Path.Clone(),
		Expression: m.Expression.Clone(),
	}
}

// Equals checks for equality with another PlanDefinitionDynamicValue instance
func (m *PlanDefinitionDynamicValue) Equals(other *PlanDefinitionDynamicValue) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Path.Equals(other.Path) { return false }
	if !m.Expression.Equals(other.Expression) { return false }
	return true
}
