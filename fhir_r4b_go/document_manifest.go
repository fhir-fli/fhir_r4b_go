// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"
)

// DocumentManifest
// A collection of documents compiled for a purpose together with metadata that applies to the collection.
type DocumentManifest struct {
	extends DomainResource
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
	Type *CodeableConcept `json:"type,omitempty"`
	Subject *Reference `json:"subject,omitempty"`
	Created *FhirDateTime `json:"created,omitempty"`
	Author []*Reference `json:"author,omitempty"`
	Recipient []*Reference `json:"recipient,omitempty"`
	Source *FhirUri `json:"source,omitempty"`
	Description *FhirString `json:"description,omitempty"`
	Content []*Reference `json:"content,omitempty"`
	Related []*DocumentManifestRelated `json:"related,omitempty"`
}

// NewDocumentManifest creates a new DocumentManifest instance.
func NewDocumentManifest() *DocumentManifest {
	return &DocumentManifest{}
}

// UnmarshalJSON populates DocumentManifest from JSON data.
func (m *DocumentManifest) UnmarshalJSON(data []byte) error {
	temp := struct {
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
		Type *CodeableConcept `json:"type,omitempty"`
		Subject *Reference `json:"subject,omitempty"`
		Created *FhirDateTime `json:"created,omitempty"`
		Author []*Reference `json:"author,omitempty"`
		Recipient []*Reference `json:"recipient,omitempty"`
		Source *FhirUri `json:"source,omitempty"`
		Description *FhirString `json:"description,omitempty"`
		Content []*Reference `json:"content,omitempty"`
		Related []*DocumentManifestRelated `json:"related,omitempty"`
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
	m.MasterIdentifier = temp.MasterIdentifier
	m.Identifier = temp.Identifier
	m.Status = temp.Status
	m.Type = temp.Type
	m.Subject = temp.Subject
	m.Created = temp.Created
	m.Author = temp.Author
	m.Recipient = temp.Recipient
	m.Source = temp.Source
	m.Description = temp.Description
	m.Content = temp.Content
	m.Related = temp.Related
	return nil
}

// MarshalJSON converts DocumentManifest to JSON data.
func (m *DocumentManifest) MarshalJSON() ([]byte, error) {
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
		MasterIdentifier *Identifier `json:"masteridentifier,omitempty"`
		Identifier []*Identifier `json:"identifier,omitempty"`
		Status *DocumentReferenceStatus `json:"status,omitempty"`
		Type *CodeableConcept `json:"type,omitempty"`
		Subject *Reference `json:"subject,omitempty"`
		Created interface{} `json:"created,omitempty"`
		CreatedElement map[string]interface{} `json:"_created,omitempty"`
		Author []*Reference `json:"author,omitempty"`
		Recipient []*Reference `json:"recipient,omitempty"`
		Source interface{} `json:"source,omitempty"`
		SourceElement map[string]interface{} `json:"_source,omitempty"`
		Description interface{} `json:"description,omitempty"`
		DescriptionElement map[string]interface{} `json:"_description,omitempty"`
		Content []*Reference `json:"content,omitempty"`
		Related []*DocumentManifestRelated `json:"related,omitempty"`
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
	output.MasterIdentifier = m.MasterIdentifier
	output.Identifier = m.Identifier
	output.Status = m.Status
	output.Type = m.Type
	output.Subject = m.Subject
	if m.Created != nil && m.Created.Value != nil {
		output.Created = m.Created.Value
		if m.Created.Element != nil {
			output.CreatedElement = toMapOrNil(m.Created.Element.MarshalJSON())
		}
	}
	output.Author = m.Author
	output.Recipient = m.Recipient
	if m.Source != nil && m.Source.Value != nil {
		output.Source = m.Source.Value
		if m.Source.Element != nil {
			output.SourceElement = toMapOrNil(m.Source.Element.MarshalJSON())
		}
	}
	if m.Description != nil && m.Description.Value != nil {
		output.Description = m.Description.Value
		if m.Description.Element != nil {
			output.DescriptionElement = toMapOrNil(m.Description.Element.MarshalJSON())
		}
	}
	output.Content = m.Content
	output.Related = m.Related
	return json.Marshal(output)
}

// Clone creates a deep copy of DocumentManifest.
func (m *DocumentManifest) Clone() *DocumentManifest {
	if m == nil { return nil }
	return &DocumentManifest{
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
		Type: m.Type.Clone(),
		Subject: m.Subject.Clone(),
		Created: m.Created.Clone(),
		Author: cloneSlices(m.Author),
		Recipient: cloneSlices(m.Recipient),
		Source: m.Source.Clone(),
		Description: m.Description.Clone(),
		Content: cloneSlices(m.Content),
		Related: cloneSlices(m.Related),
	}
}

// Equals checks equality between two DocumentManifest instances.
func (m *DocumentManifest) Equals(other *DocumentManifest) bool {
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
	if !m.Type.Equals(other.Type) { return false }
	if !m.Subject.Equals(other.Subject) { return false }
	if !m.Created.Equals(other.Created) { return false }
	if !compareSlices(m.Author, other.Author) { return false }
	if !compareSlices(m.Recipient, other.Recipient) { return false }
	if !m.Source.Equals(other.Source) { return false }
	if !m.Description.Equals(other.Description) { return false }
	if !compareSlices(m.Content, other.Content) { return false }
	if !compareSlices(m.Related, other.Related) { return false }
	return true
}

// DocumentManifestRelated
// Related identifiers or resources associated with the DocumentManifest.
type DocumentManifestRelated struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Identifier *Identifier `json:"identifier,omitempty"`
	Ref *Reference `json:"ref,omitempty"`
}

// NewDocumentManifestRelated creates a new DocumentManifestRelated instance.
func NewDocumentManifestRelated() *DocumentManifestRelated {
	return &DocumentManifestRelated{}
}

// UnmarshalJSON populates DocumentManifestRelated from JSON data.
func (m *DocumentManifestRelated) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Identifier *Identifier `json:"identifier,omitempty"`
		Ref *Reference `json:"ref,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Identifier = temp.Identifier
	m.Ref = temp.Ref
	return nil
}

// MarshalJSON converts DocumentManifestRelated to JSON data.
func (m *DocumentManifestRelated) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Identifier *Identifier `json:"identifier,omitempty"`
		Ref *Reference `json:"ref,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.Identifier = m.Identifier
	output.Ref = m.Ref
	return json.Marshal(output)
}

// Clone creates a deep copy of DocumentManifestRelated.
func (m *DocumentManifestRelated) Clone() *DocumentManifestRelated {
	if m == nil { return nil }
	return &DocumentManifestRelated{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Identifier: m.Identifier.Clone(),
		Ref: m.Ref.Clone(),
	}
}

// Equals checks equality between two DocumentManifestRelated instances.
func (m *DocumentManifestRelated) Equals(other *DocumentManifestRelated) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Identifier.Equals(other.Identifier) { return false }
	if !m.Ref.Equals(other.Ref) { return false }
	return true
}

