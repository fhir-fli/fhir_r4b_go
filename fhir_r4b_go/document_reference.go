// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json")

// DocumentReference
// A reference to a document of any kind for any purpose. Provides metadata about the document so that the document can be discovered and managed. The scope of a document is any seralized object with a mime-type, so includes formal patient centric documents (CDA), cliical notes, scanned paper, and non-patient specific documents like policy text.
type DocumentReference struct {
	DomainResource
	Id *FhirString `json:"id,omitempty"`
	Meta *FhirMeta `json:"meta,omitempty"`
	ImplicitRules *FhirUri `json:"implicitrules,omitempty"`
	Language *CommonLanguages `json:"language,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Contained []*Resource `json:"contained,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	MasterIdentifier *Identifier `json:"masteridentifier,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Status *DocumentReferenceStatus `json:"status,omitempty"`
	DocStatus *CompositionStatus `json:"docstatus,omitempty"`
	Type *CodeableConcept `json:"type,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Subject *Reference `json:"subject,omitempty"`
	Date *FhirInstant `json:"date,omitempty"`
	Author []*Reference `json:"author,omitempty"`
	Authenticator *Reference `json:"authenticator,omitempty"`
	Custodian *Reference `json:"custodian,omitempty"`
	RelatesTo []*DocumentReferenceRelatesTo `json:"relatesto,omitempty"`
	Description *FhirString `json:"description,omitempty"`
	SecurityLabel []*CodeableConcept `json:"securitylabel,omitempty"`
	Content []*DocumentReferenceContent `json:"content,omitempty"`
	Context *DocumentReferenceContext `json:"context,omitempty"`
}

// NewDocumentReference creates a new DocumentReference instance
func NewDocumentReference() *DocumentReference {
	return &DocumentReference{}
}

// FromJSON populates DocumentReference from JSON data
func (m *DocumentReference) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts DocumentReference to JSON data
func (m *DocumentReference) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of DocumentReference
func (m *DocumentReference) Clone() *DocumentReference {
	if m == nil { return nil }
	return &DocumentReference{
		Id: m.Id.Clone(),
		Meta: m.Meta.Clone(),
		ImplicitRules: m.ImplicitRules.Clone(),
		Language: m.Language.Clone(),
		Text: m.Text.Clone(),
		Contained: cloneSlices(m.Contained),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		MasterIdentifier: m.MasterIdentifier.Clone(),
		Identifier: cloneSlices(m.Identifier),
		Status: m.Status.Clone(),
		DocStatus: m.DocStatus.Clone(),
		Type: m.Type.Clone(),
		Category: cloneSlices(m.Category),
		Subject: m.Subject.Clone(),
		Date: m.Date.Clone(),
		Author: cloneSlices(m.Author),
		Authenticator: m.Authenticator.Clone(),
		Custodian: m.Custodian.Clone(),
		RelatesTo: cloneSlices(m.RelatesTo),
		Description: m.Description.Clone(),
		SecurityLabel: cloneSlices(m.SecurityLabel),
		Content: cloneSlices(m.Content),
		Context: m.Context.Clone(),
	}
}

// Equals checks for equality with another DocumentReference instance
func (m *DocumentReference) Equals(other *DocumentReference) bool {
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
	if !m.MasterIdentifier.Equals(other.MasterIdentifier) { return false }
	if !compareSlices(m.Identifier, other.Identifier) { return false }
	if !m.Status.Equals(other.Status) { return false }
	if !m.DocStatus.Equals(other.DocStatus) { return false }
	if !m.Type.Equals(other.Type) { return false }
	if !compareSlices(m.Category, other.Category) { return false }
	if !m.Subject.Equals(other.Subject) { return false }
	if !m.Date.Equals(other.Date) { return false }
	if !compareSlices(m.Author, other.Author) { return false }
	if !m.Authenticator.Equals(other.Authenticator) { return false }
	if !m.Custodian.Equals(other.Custodian) { return false }
	if !compareSlices(m.RelatesTo, other.RelatesTo) { return false }
	if !m.Description.Equals(other.Description) { return false }
	if !compareSlices(m.SecurityLabel, other.SecurityLabel) { return false }
	if !compareSlices(m.Content, other.Content) { return false }
	if !m.Context.Equals(other.Context) { return false }
	return true
}

// DocumentReferenceRelatesTo
// Relationships that this document has with other document references that already exist.
type DocumentReferenceRelatesTo struct {
	BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Code *DocumentRelationshipType `json:"code,omitempty"`
	Target *Reference `json:"target,omitempty"`
}

// NewDocumentReferenceRelatesTo creates a new DocumentReferenceRelatesTo instance
func NewDocumentReferenceRelatesTo() *DocumentReferenceRelatesTo {
	return &DocumentReferenceRelatesTo{}
}

// FromJSON populates DocumentReferenceRelatesTo from JSON data
func (m *DocumentReferenceRelatesTo) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts DocumentReferenceRelatesTo to JSON data
func (m *DocumentReferenceRelatesTo) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of DocumentReferenceRelatesTo
func (m *DocumentReferenceRelatesTo) Clone() *DocumentReferenceRelatesTo {
	if m == nil { return nil }
	return &DocumentReferenceRelatesTo{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Code: m.Code.Clone(),
		Target: m.Target.Clone(),
	}
}

// Equals checks for equality with another DocumentReferenceRelatesTo instance
func (m *DocumentReferenceRelatesTo) Equals(other *DocumentReferenceRelatesTo) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Code.Equals(other.Code) { return false }
	if !m.Target.Equals(other.Target) { return false }
	return true
}

// DocumentReferenceContent
// The document and format referenced. There may be multiple content element repetitions, each with a different format.
type DocumentReferenceContent struct {
	BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Attachment *Attachment `json:"attachment,omitempty"`
	Format *Coding `json:"format,omitempty"`
}

// NewDocumentReferenceContent creates a new DocumentReferenceContent instance
func NewDocumentReferenceContent() *DocumentReferenceContent {
	return &DocumentReferenceContent{}
}

// FromJSON populates DocumentReferenceContent from JSON data
func (m *DocumentReferenceContent) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts DocumentReferenceContent to JSON data
func (m *DocumentReferenceContent) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of DocumentReferenceContent
func (m *DocumentReferenceContent) Clone() *DocumentReferenceContent {
	if m == nil { return nil }
	return &DocumentReferenceContent{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Attachment: m.Attachment.Clone(),
		Format: m.Format.Clone(),
	}
}

// Equals checks for equality with another DocumentReferenceContent instance
func (m *DocumentReferenceContent) Equals(other *DocumentReferenceContent) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Attachment.Equals(other.Attachment) { return false }
	if !m.Format.Equals(other.Format) { return false }
	return true
}

// DocumentReferenceContext
// The clinical context in which the document was prepared.
type DocumentReferenceContext struct {
	BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Encounter []*Reference `json:"encounter,omitempty"`
	Event []*CodeableConcept `json:"event,omitempty"`
	Period *Period `json:"period,omitempty"`
	FacilityType *CodeableConcept `json:"facilitytype,omitempty"`
	PracticeSetting *CodeableConcept `json:"practicesetting,omitempty"`
	SourcePatientInfo *Reference `json:"sourcepatientinfo,omitempty"`
	Related []*Reference `json:"related,omitempty"`
}

// NewDocumentReferenceContext creates a new DocumentReferenceContext instance
func NewDocumentReferenceContext() *DocumentReferenceContext {
	return &DocumentReferenceContext{}
}

// FromJSON populates DocumentReferenceContext from JSON data
func (m *DocumentReferenceContext) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts DocumentReferenceContext to JSON data
func (m *DocumentReferenceContext) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of DocumentReferenceContext
func (m *DocumentReferenceContext) Clone() *DocumentReferenceContext {
	if m == nil { return nil }
	return &DocumentReferenceContext{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Encounter: cloneSlices(m.Encounter),
		Event: cloneSlices(m.Event),
		Period: m.Period.Clone(),
		FacilityType: m.FacilityType.Clone(),
		PracticeSetting: m.PracticeSetting.Clone(),
		SourcePatientInfo: m.SourcePatientInfo.Clone(),
		Related: cloneSlices(m.Related),
	}
}

// Equals checks for equality with another DocumentReferenceContext instance
func (m *DocumentReferenceContext) Equals(other *DocumentReferenceContext) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !compareSlices(m.Encounter, other.Encounter) { return false }
	if !compareSlices(m.Event, other.Event) { return false }
	if !m.Period.Equals(other.Period) { return false }
	if !m.FacilityType.Equals(other.FacilityType) { return false }
	if !m.PracticeSetting.Equals(other.PracticeSetting) { return false }
	if !m.SourcePatientInfo.Equals(other.SourcePatientInfo) { return false }
	if !compareSlices(m.Related, other.Related) { return false }
	return true
}
