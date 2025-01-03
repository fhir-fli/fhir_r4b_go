// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"
)

// EvidenceVariable
// The EvidenceVariable resource describes an element that knowledge (Evidence) is about.
type EvidenceVariable struct {
	extends DomainResource
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
	Date *FhirDateTime `json:"date,omitempty"`
	Description *FhirMarkdown `json:"description,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
	UseContext []*UsageContext `json:"usecontext,omitempty"`
	Publisher *FhirString `json:"publisher,omitempty"`
	Contact []*ContactDetail `json:"contact,omitempty"`
	Author []*ContactDetail `json:"author,omitempty"`
	Editor []*ContactDetail `json:"editor,omitempty"`
	Reviewer []*ContactDetail `json:"reviewer,omitempty"`
	Endorser []*ContactDetail `json:"endorser,omitempty"`
	RelatedArtifact []*RelatedArtifact `json:"relatedartifact,omitempty"`
	Actual *FhirBoolean `json:"actual,omitempty"`
	CharacteristicCombination *CharacteristicCombination `json:"characteristiccombination,omitempty"`
	Characteristic []*EvidenceVariableCharacteristic `json:"characteristic,omitempty"`
	Handling *EvidenceVariableHandling `json:"handling,omitempty"`
	Category []*EvidenceVariableCategory `json:"category,omitempty"`
}

// NewEvidenceVariable creates a new EvidenceVariable instance.
func NewEvidenceVariable() *EvidenceVariable {
	return &EvidenceVariable{}
}

// UnmarshalJSON populates EvidenceVariable from JSON data.
func (m *EvidenceVariable) UnmarshalJSON(data []byte) error {
	temp := struct {
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
		Date *FhirDateTime `json:"date,omitempty"`
		Description *FhirMarkdown `json:"description,omitempty"`
		Note []*Annotation `json:"note,omitempty"`
		UseContext []*UsageContext `json:"usecontext,omitempty"`
		Publisher *FhirString `json:"publisher,omitempty"`
		Contact []*ContactDetail `json:"contact,omitempty"`
		Author []*ContactDetail `json:"author,omitempty"`
		Editor []*ContactDetail `json:"editor,omitempty"`
		Reviewer []*ContactDetail `json:"reviewer,omitempty"`
		Endorser []*ContactDetail `json:"endorser,omitempty"`
		RelatedArtifact []*RelatedArtifact `json:"relatedartifact,omitempty"`
		Actual *FhirBoolean `json:"actual,omitempty"`
		CharacteristicCombination *CharacteristicCombination `json:"characteristiccombination,omitempty"`
		Characteristic []*EvidenceVariableCharacteristic `json:"characteristic,omitempty"`
		Handling *EvidenceVariableHandling `json:"handling,omitempty"`
		Category []*EvidenceVariableCategory `json:"category,omitempty"`
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
	m.Url = temp.Url
	m.Identifier = temp.Identifier
	m.Version = temp.Version
	m.Name = temp.Name
	m.Title = temp.Title
	m.ShortTitle = temp.ShortTitle
	m.Subtitle = temp.Subtitle
	m.Status = temp.Status
	m.Date = temp.Date
	m.Description = temp.Description
	m.Note = temp.Note
	m.UseContext = temp.UseContext
	m.Publisher = temp.Publisher
	m.Contact = temp.Contact
	m.Author = temp.Author
	m.Editor = temp.Editor
	m.Reviewer = temp.Reviewer
	m.Endorser = temp.Endorser
	m.RelatedArtifact = temp.RelatedArtifact
	m.Actual = temp.Actual
	m.CharacteristicCombination = temp.CharacteristicCombination
	m.Characteristic = temp.Characteristic
	m.Handling = temp.Handling
	m.Category = temp.Category
	return nil
}

// MarshalJSON converts EvidenceVariable to JSON data.
func (m *EvidenceVariable) MarshalJSON() ([]byte, error) {
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
		Url interface{} `json:"url,omitempty"`
		UrlElement map[string]interface{} `json:"_url,omitempty"`
		Identifier []*Identifier `json:"identifier,omitempty"`
		Version interface{} `json:"version,omitempty"`
		VersionElement map[string]interface{} `json:"_version,omitempty"`
		Name interface{} `json:"name,omitempty"`
		NameElement map[string]interface{} `json:"_name,omitempty"`
		Title interface{} `json:"title,omitempty"`
		TitleElement map[string]interface{} `json:"_title,omitempty"`
		ShortTitle interface{} `json:"shorttitle,omitempty"`
		ShortTitleElement map[string]interface{} `json:"_shorttitle,omitempty"`
		Subtitle interface{} `json:"subtitle,omitempty"`
		SubtitleElement map[string]interface{} `json:"_subtitle,omitempty"`
		Status *PublicationStatus `json:"status,omitempty"`
		Date interface{} `json:"date,omitempty"`
		DateElement map[string]interface{} `json:"_date,omitempty"`
		Description interface{} `json:"description,omitempty"`
		DescriptionElement map[string]interface{} `json:"_description,omitempty"`
		Note []*Annotation `json:"note,omitempty"`
		UseContext []*UsageContext `json:"usecontext,omitempty"`
		Publisher interface{} `json:"publisher,omitempty"`
		PublisherElement map[string]interface{} `json:"_publisher,omitempty"`
		Contact []*ContactDetail `json:"contact,omitempty"`
		Author []*ContactDetail `json:"author,omitempty"`
		Editor []*ContactDetail `json:"editor,omitempty"`
		Reviewer []*ContactDetail `json:"reviewer,omitempty"`
		Endorser []*ContactDetail `json:"endorser,omitempty"`
		RelatedArtifact []*RelatedArtifact `json:"relatedartifact,omitempty"`
		Actual interface{} `json:"actual,omitempty"`
		ActualElement map[string]interface{} `json:"_actual,omitempty"`
		CharacteristicCombination *CharacteristicCombination `json:"characteristiccombination,omitempty"`
		Characteristic []*EvidenceVariableCharacteristic `json:"characteristic,omitempty"`
		Handling *EvidenceVariableHandling `json:"handling,omitempty"`
		Category []*EvidenceVariableCategory `json:"category,omitempty"`
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
	if m.Url != nil && m.Url.Value != nil {
		output.Url = m.Url.Value
		if m.Url.Element != nil {
			output.UrlElement = toMapOrNil(m.Url.Element.MarshalJSON())
		}
	}
	output.Identifier = m.Identifier
	if m.Version != nil && m.Version.Value != nil {
		output.Version = m.Version.Value
		if m.Version.Element != nil {
			output.VersionElement = toMapOrNil(m.Version.Element.MarshalJSON())
		}
	}
	if m.Name != nil && m.Name.Value != nil {
		output.Name = m.Name.Value
		if m.Name.Element != nil {
			output.NameElement = toMapOrNil(m.Name.Element.MarshalJSON())
		}
	}
	if m.Title != nil && m.Title.Value != nil {
		output.Title = m.Title.Value
		if m.Title.Element != nil {
			output.TitleElement = toMapOrNil(m.Title.Element.MarshalJSON())
		}
	}
	if m.ShortTitle != nil && m.ShortTitle.Value != nil {
		output.ShortTitle = m.ShortTitle.Value
		if m.ShortTitle.Element != nil {
			output.ShortTitleElement = toMapOrNil(m.ShortTitle.Element.MarshalJSON())
		}
	}
	if m.Subtitle != nil && m.Subtitle.Value != nil {
		output.Subtitle = m.Subtitle.Value
		if m.Subtitle.Element != nil {
			output.SubtitleElement = toMapOrNil(m.Subtitle.Element.MarshalJSON())
		}
	}
	output.Status = m.Status
	if m.Date != nil && m.Date.Value != nil {
		output.Date = m.Date.Value
		if m.Date.Element != nil {
			output.DateElement = toMapOrNil(m.Date.Element.MarshalJSON())
		}
	}
	if m.Description != nil && m.Description.Value != nil {
		output.Description = m.Description.Value
		if m.Description.Element != nil {
			output.DescriptionElement = toMapOrNil(m.Description.Element.MarshalJSON())
		}
	}
	output.Note = m.Note
	output.UseContext = m.UseContext
	if m.Publisher != nil && m.Publisher.Value != nil {
		output.Publisher = m.Publisher.Value
		if m.Publisher.Element != nil {
			output.PublisherElement = toMapOrNil(m.Publisher.Element.MarshalJSON())
		}
	}
	output.Contact = m.Contact
	output.Author = m.Author
	output.Editor = m.Editor
	output.Reviewer = m.Reviewer
	output.Endorser = m.Endorser
	output.RelatedArtifact = m.RelatedArtifact
	if m.Actual != nil && m.Actual.Value != nil {
		output.Actual = m.Actual.Value
		if m.Actual.Element != nil {
			output.ActualElement = toMapOrNil(m.Actual.Element.MarshalJSON())
		}
	}
	output.CharacteristicCombination = m.CharacteristicCombination
	output.Characteristic = m.Characteristic
	output.Handling = m.Handling
	output.Category = m.Category
	return json.Marshal(output)
}

// Clone creates a deep copy of EvidenceVariable.
func (m *EvidenceVariable) Clone() *EvidenceVariable {
	if m == nil { return nil }
	return &EvidenceVariable{
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
		Date: m.Date.Clone(),
		Description: m.Description.Clone(),
		Note: cloneSlices(m.Note),
		UseContext: cloneSlices(m.UseContext),
		Publisher: m.Publisher.Clone(),
		Contact: cloneSlices(m.Contact),
		Author: cloneSlices(m.Author),
		Editor: cloneSlices(m.Editor),
		Reviewer: cloneSlices(m.Reviewer),
		Endorser: cloneSlices(m.Endorser),
		RelatedArtifact: cloneSlices(m.RelatedArtifact),
		Actual: m.Actual.Clone(),
		CharacteristicCombination: m.CharacteristicCombination.Clone(),
		Characteristic: cloneSlices(m.Characteristic),
		Handling: m.Handling.Clone(),
		Category: cloneSlices(m.Category),
	}
}

// Equals checks equality between two EvidenceVariable instances.
func (m *EvidenceVariable) Equals(other *EvidenceVariable) bool {
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
	if !m.Date.Equals(other.Date) { return false }
	if !m.Description.Equals(other.Description) { return false }
	if !compareSlices(m.Note, other.Note) { return false }
	if !compareSlices(m.UseContext, other.UseContext) { return false }
	if !m.Publisher.Equals(other.Publisher) { return false }
	if !compareSlices(m.Contact, other.Contact) { return false }
	if !compareSlices(m.Author, other.Author) { return false }
	if !compareSlices(m.Editor, other.Editor) { return false }
	if !compareSlices(m.Reviewer, other.Reviewer) { return false }
	if !compareSlices(m.Endorser, other.Endorser) { return false }
	if !compareSlices(m.RelatedArtifact, other.RelatedArtifact) { return false }
	if !m.Actual.Equals(other.Actual) { return false }
	if !m.CharacteristicCombination.Equals(other.CharacteristicCombination) { return false }
	if !compareSlices(m.Characteristic, other.Characteristic) { return false }
	if !m.Handling.Equals(other.Handling) { return false }
	if !compareSlices(m.Category, other.Category) { return false }
	return true
}

// EvidenceVariableCharacteristic
// A characteristic that defines the members of the evidence element. Multiple characteristics are applied with "and" semantics.
type EvidenceVariableCharacteristic struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Description *FhirString `json:"description,omitempty"`
	DefinitionReference *Reference `json:"definitionreference,omitempty"`
	DefinitionCanonical *FhirCanonical `json:"definitioncanonical,omitempty"`
	DefinitionCodeableConcept *CodeableConcept `json:"definitioncodeableconcept,omitempty"`
	DefinitionExpression *FhirExpression `json:"definitionexpression,omitempty"`
	Method *CodeableConcept `json:"method,omitempty"`
	Device *Reference `json:"device,omitempty"`
	Exclude *FhirBoolean `json:"exclude,omitempty"`
	TimeFromStart *EvidenceVariableTimeFromStart `json:"timefromstart,omitempty"`
	GroupMeasure *GroupMeasure `json:"groupmeasure,omitempty"`
}

// NewEvidenceVariableCharacteristic creates a new EvidenceVariableCharacteristic instance.
func NewEvidenceVariableCharacteristic() *EvidenceVariableCharacteristic {
	return &EvidenceVariableCharacteristic{}
}

// UnmarshalJSON populates EvidenceVariableCharacteristic from JSON data.
func (m *EvidenceVariableCharacteristic) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Description *FhirString `json:"description,omitempty"`
		DefinitionReference *Reference `json:"definitionreference,omitempty"`
		DefinitionCanonical *FhirCanonical `json:"definitioncanonical,omitempty"`
		DefinitionCodeableConcept *CodeableConcept `json:"definitioncodeableconcept,omitempty"`
		DefinitionExpression *FhirExpression `json:"definitionexpression,omitempty"`
		Method *CodeableConcept `json:"method,omitempty"`
		Device *Reference `json:"device,omitempty"`
		Exclude *FhirBoolean `json:"exclude,omitempty"`
		TimeFromStart *EvidenceVariableTimeFromStart `json:"timefromstart,omitempty"`
		GroupMeasure *GroupMeasure `json:"groupmeasure,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Description = temp.Description
	m.DefinitionReference = temp.DefinitionReference
	m.DefinitionCanonical = temp.DefinitionCanonical
	m.DefinitionCodeableConcept = temp.DefinitionCodeableConcept
	m.DefinitionExpression = temp.DefinitionExpression
	m.Method = temp.Method
	m.Device = temp.Device
	m.Exclude = temp.Exclude
	m.TimeFromStart = temp.TimeFromStart
	m.GroupMeasure = temp.GroupMeasure
	return nil
}

// MarshalJSON converts EvidenceVariableCharacteristic to JSON data.
func (m *EvidenceVariableCharacteristic) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Description interface{} `json:"description,omitempty"`
		DescriptionElement map[string]interface{} `json:"_description,omitempty"`
		DefinitionReference *Reference `json:"definitionreference,omitempty"`
		DefinitionCanonical interface{} `json:"definitioncanonical,omitempty"`
		DefinitionCanonicalElement map[string]interface{} `json:"_definitioncanonical,omitempty"`
		DefinitionCodeableConcept *CodeableConcept `json:"definitioncodeableconcept,omitempty"`
		DefinitionExpression *FhirExpression `json:"definitionexpression,omitempty"`
		Method *CodeableConcept `json:"method,omitempty"`
		Device *Reference `json:"device,omitempty"`
		Exclude interface{} `json:"exclude,omitempty"`
		ExcludeElement map[string]interface{} `json:"_exclude,omitempty"`
		TimeFromStart *EvidenceVariableTimeFromStart `json:"timefromstart,omitempty"`
		GroupMeasure *GroupMeasure `json:"groupmeasure,omitempty"`
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
	output.DefinitionReference = m.DefinitionReference
	if m.DefinitionCanonical != nil && m.DefinitionCanonical.Value != nil {
		output.DefinitionCanonical = m.DefinitionCanonical.Value
		if m.DefinitionCanonical.Element != nil {
			output.DefinitionCanonicalElement = toMapOrNil(m.DefinitionCanonical.Element.MarshalJSON())
		}
	}
	output.DefinitionCodeableConcept = m.DefinitionCodeableConcept
	output.DefinitionExpression = m.DefinitionExpression
	output.Method = m.Method
	output.Device = m.Device
	if m.Exclude != nil && m.Exclude.Value != nil {
		output.Exclude = m.Exclude.Value
		if m.Exclude.Element != nil {
			output.ExcludeElement = toMapOrNil(m.Exclude.Element.MarshalJSON())
		}
	}
	output.TimeFromStart = m.TimeFromStart
	output.GroupMeasure = m.GroupMeasure
	return json.Marshal(output)
}

// Clone creates a deep copy of EvidenceVariableCharacteristic.
func (m *EvidenceVariableCharacteristic) Clone() *EvidenceVariableCharacteristic {
	if m == nil { return nil }
	return &EvidenceVariableCharacteristic{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Description: m.Description.Clone(),
		DefinitionReference: m.DefinitionReference.Clone(),
		DefinitionCanonical: m.DefinitionCanonical.Clone(),
		DefinitionCodeableConcept: m.DefinitionCodeableConcept.Clone(),
		DefinitionExpression: m.DefinitionExpression.Clone(),
		Method: m.Method.Clone(),
		Device: m.Device.Clone(),
		Exclude: m.Exclude.Clone(),
		TimeFromStart: m.TimeFromStart.Clone(),
		GroupMeasure: m.GroupMeasure.Clone(),
	}
}

// Equals checks equality between two EvidenceVariableCharacteristic instances.
func (m *EvidenceVariableCharacteristic) Equals(other *EvidenceVariableCharacteristic) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Description.Equals(other.Description) { return false }
	if !m.DefinitionReference.Equals(other.DefinitionReference) { return false }
	if !m.DefinitionCanonical.Equals(other.DefinitionCanonical) { return false }
	if !m.DefinitionCodeableConcept.Equals(other.DefinitionCodeableConcept) { return false }
	if !m.DefinitionExpression.Equals(other.DefinitionExpression) { return false }
	if !m.Method.Equals(other.Method) { return false }
	if !m.Device.Equals(other.Device) { return false }
	if !m.Exclude.Equals(other.Exclude) { return false }
	if !m.TimeFromStart.Equals(other.TimeFromStart) { return false }
	if !m.GroupMeasure.Equals(other.GroupMeasure) { return false }
	return true
}

// EvidenceVariableTimeFromStart
// Indicates duration, period, or point of observation from the participant's study entry.
type EvidenceVariableTimeFromStart struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Description *FhirString `json:"description,omitempty"`
	Quantity *Quantity `json:"quantity,omitempty"`
	Range *Range `json:"range,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
}

// NewEvidenceVariableTimeFromStart creates a new EvidenceVariableTimeFromStart instance.
func NewEvidenceVariableTimeFromStart() *EvidenceVariableTimeFromStart {
	return &EvidenceVariableTimeFromStart{}
}

// UnmarshalJSON populates EvidenceVariableTimeFromStart from JSON data.
func (m *EvidenceVariableTimeFromStart) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Description *FhirString `json:"description,omitempty"`
		Quantity *Quantity `json:"quantity,omitempty"`
		Range *Range `json:"range,omitempty"`
		Note []*Annotation `json:"note,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Description = temp.Description
	m.Quantity = temp.Quantity
	m.Range = temp.Range
	m.Note = temp.Note
	return nil
}

// MarshalJSON converts EvidenceVariableTimeFromStart to JSON data.
func (m *EvidenceVariableTimeFromStart) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Description interface{} `json:"description,omitempty"`
		DescriptionElement map[string]interface{} `json:"_description,omitempty"`
		Quantity *Quantity `json:"quantity,omitempty"`
		Range *Range `json:"range,omitempty"`
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
	if m.Description != nil && m.Description.Value != nil {
		output.Description = m.Description.Value
		if m.Description.Element != nil {
			output.DescriptionElement = toMapOrNil(m.Description.Element.MarshalJSON())
		}
	}
	output.Quantity = m.Quantity
	output.Range = m.Range
	output.Note = m.Note
	return json.Marshal(output)
}

// Clone creates a deep copy of EvidenceVariableTimeFromStart.
func (m *EvidenceVariableTimeFromStart) Clone() *EvidenceVariableTimeFromStart {
	if m == nil { return nil }
	return &EvidenceVariableTimeFromStart{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Description: m.Description.Clone(),
		Quantity: m.Quantity.Clone(),
		Range: m.Range.Clone(),
		Note: cloneSlices(m.Note),
	}
}

// Equals checks equality between two EvidenceVariableTimeFromStart instances.
func (m *EvidenceVariableTimeFromStart) Equals(other *EvidenceVariableTimeFromStart) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Description.Equals(other.Description) { return false }
	if !m.Quantity.Equals(other.Quantity) { return false }
	if !m.Range.Equals(other.Range) { return false }
	if !compareSlices(m.Note, other.Note) { return false }
	return true
}

// EvidenceVariableCategory
// A grouping (or set of values) described along with other groupings to specify the set of groupings allowed for the variable.
type EvidenceVariableCategory struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Name *FhirString `json:"name,omitempty"`
	ValueCodeableConcept *CodeableConcept `json:"valuecodeableconcept,omitempty"`
	ValueQuantity *Quantity `json:"valuequantity,omitempty"`
	ValueRange *Range `json:"valuerange,omitempty"`
}

// NewEvidenceVariableCategory creates a new EvidenceVariableCategory instance.
func NewEvidenceVariableCategory() *EvidenceVariableCategory {
	return &EvidenceVariableCategory{}
}

// UnmarshalJSON populates EvidenceVariableCategory from JSON data.
func (m *EvidenceVariableCategory) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Name *FhirString `json:"name,omitempty"`
		ValueCodeableConcept *CodeableConcept `json:"valuecodeableconcept,omitempty"`
		ValueQuantity *Quantity `json:"valuequantity,omitempty"`
		ValueRange *Range `json:"valuerange,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Name = temp.Name
	m.ValueCodeableConcept = temp.ValueCodeableConcept
	m.ValueQuantity = temp.ValueQuantity
	m.ValueRange = temp.ValueRange
	return nil
}

// MarshalJSON converts EvidenceVariableCategory to JSON data.
func (m *EvidenceVariableCategory) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Name interface{} `json:"name,omitempty"`
		NameElement map[string]interface{} `json:"_name,omitempty"`
		ValueCodeableConcept *CodeableConcept `json:"valuecodeableconcept,omitempty"`
		ValueQuantity *Quantity `json:"valuequantity,omitempty"`
		ValueRange *Range `json:"valuerange,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	if m.Name != nil && m.Name.Value != nil {
		output.Name = m.Name.Value
		if m.Name.Element != nil {
			output.NameElement = toMapOrNil(m.Name.Element.MarshalJSON())
		}
	}
	output.ValueCodeableConcept = m.ValueCodeableConcept
	output.ValueQuantity = m.ValueQuantity
	output.ValueRange = m.ValueRange
	return json.Marshal(output)
}

// Clone creates a deep copy of EvidenceVariableCategory.
func (m *EvidenceVariableCategory) Clone() *EvidenceVariableCategory {
	if m == nil { return nil }
	return &EvidenceVariableCategory{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Name: m.Name.Clone(),
		ValueCodeableConcept: m.ValueCodeableConcept.Clone(),
		ValueQuantity: m.ValueQuantity.Clone(),
		ValueRange: m.ValueRange.Clone(),
	}
}

// Equals checks equality between two EvidenceVariableCategory instances.
func (m *EvidenceVariableCategory) Equals(other *EvidenceVariableCategory) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Name.Equals(other.Name) { return false }
	if !m.ValueCodeableConcept.Equals(other.ValueCodeableConcept) { return false }
	if !m.ValueQuantity.Equals(other.ValueQuantity) { return false }
	if !m.ValueRange.Equals(other.ValueRange) { return false }
	return true
}

