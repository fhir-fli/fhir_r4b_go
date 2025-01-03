// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"
)

// GraphDefinition
// A formal computable definition of a graph of resources - that is, a coherent set of resources that form a graph by following references. The Graph Definition resource defines a set and makes rules about the set.
type GraphDefinition struct {
	extends CanonicalResource
	Id *FhirString `json:"id,omitempty"`
	Meta *FhirMeta `json:"meta,omitempty"`
	ImplicitRules *FhirUri `json:"implicitrules,omitempty"`
	Language *CommonLanguages `json:"language,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Contained []*Resource `json:"contained,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Url *FhirUri `json:"url,omitempty"`
	Version *FhirString `json:"version,omitempty"`
	Name *FhirString `json:"name,omitempty"`
	Status *PublicationStatus `json:"status,omitempty"`
	Experimental *FhirBoolean `json:"experimental,omitempty"`
	Date *FhirDateTime `json:"date,omitempty"`
	Publisher *FhirString `json:"publisher,omitempty"`
	Contact []*ContactDetail `json:"contact,omitempty"`
	Description *FhirMarkdown `json:"description,omitempty"`
	UseContext []*UsageContext `json:"usecontext,omitempty"`
	Jurisdiction []*CodeableConcept `json:"jurisdiction,omitempty"`
	Purpose *FhirMarkdown `json:"purpose,omitempty"`
	Start *FhirCode `json:"start,omitempty"`
	Profile *FhirCanonical `json:"profile,omitempty"`
	Link []*GraphDefinitionLink `json:"link,omitempty"`
}

// NewGraphDefinition creates a new GraphDefinition instance.
func NewGraphDefinition() *GraphDefinition {
	return &GraphDefinition{}
}

// UnmarshalJSON populates GraphDefinition from JSON data.
func (m *GraphDefinition) UnmarshalJSON(data []byte) error {
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
		Version *FhirString `json:"version,omitempty"`
		Name *FhirString `json:"name,omitempty"`
		Status *PublicationStatus `json:"status,omitempty"`
		Experimental *FhirBoolean `json:"experimental,omitempty"`
		Date *FhirDateTime `json:"date,omitempty"`
		Publisher *FhirString `json:"publisher,omitempty"`
		Contact []*ContactDetail `json:"contact,omitempty"`
		Description *FhirMarkdown `json:"description,omitempty"`
		UseContext []*UsageContext `json:"usecontext,omitempty"`
		Jurisdiction []*CodeableConcept `json:"jurisdiction,omitempty"`
		Purpose *FhirMarkdown `json:"purpose,omitempty"`
		Start *FhirCode `json:"start,omitempty"`
		Profile *FhirCanonical `json:"profile,omitempty"`
		Link []*GraphDefinitionLink `json:"link,omitempty"`
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
	m.Version = temp.Version
	m.Name = temp.Name
	m.Status = temp.Status
	m.Experimental = temp.Experimental
	m.Date = temp.Date
	m.Publisher = temp.Publisher
	m.Contact = temp.Contact
	m.Description = temp.Description
	m.UseContext = temp.UseContext
	m.Jurisdiction = temp.Jurisdiction
	m.Purpose = temp.Purpose
	m.Start = temp.Start
	m.Profile = temp.Profile
	m.Link = temp.Link
	return nil
}

// MarshalJSON converts GraphDefinition to JSON data.
func (m *GraphDefinition) MarshalJSON() ([]byte, error) {
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
		Version interface{} `json:"version,omitempty"`
		VersionElement map[string]interface{} `json:"_version,omitempty"`
		Name interface{} `json:"name,omitempty"`
		NameElement map[string]interface{} `json:"_name,omitempty"`
		Status *PublicationStatus `json:"status,omitempty"`
		Experimental interface{} `json:"experimental,omitempty"`
		ExperimentalElement map[string]interface{} `json:"_experimental,omitempty"`
		Date interface{} `json:"date,omitempty"`
		DateElement map[string]interface{} `json:"_date,omitempty"`
		Publisher interface{} `json:"publisher,omitempty"`
		PublisherElement map[string]interface{} `json:"_publisher,omitempty"`
		Contact []*ContactDetail `json:"contact,omitempty"`
		Description interface{} `json:"description,omitempty"`
		DescriptionElement map[string]interface{} `json:"_description,omitempty"`
		UseContext []*UsageContext `json:"usecontext,omitempty"`
		Jurisdiction []*CodeableConcept `json:"jurisdiction,omitempty"`
		Purpose interface{} `json:"purpose,omitempty"`
		PurposeElement map[string]interface{} `json:"_purpose,omitempty"`
		Start interface{} `json:"start,omitempty"`
		StartElement map[string]interface{} `json:"_start,omitempty"`
		Profile interface{} `json:"profile,omitempty"`
		ProfileElement map[string]interface{} `json:"_profile,omitempty"`
		Link []*GraphDefinitionLink `json:"link,omitempty"`
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
	output.Status = m.Status
	if m.Experimental != nil && m.Experimental.Value != nil {
		output.Experimental = m.Experimental.Value
		if m.Experimental.Element != nil {
			output.ExperimentalElement = toMapOrNil(m.Experimental.Element.MarshalJSON())
		}
	}
	if m.Date != nil && m.Date.Value != nil {
		output.Date = m.Date.Value
		if m.Date.Element != nil {
			output.DateElement = toMapOrNil(m.Date.Element.MarshalJSON())
		}
	}
	if m.Publisher != nil && m.Publisher.Value != nil {
		output.Publisher = m.Publisher.Value
		if m.Publisher.Element != nil {
			output.PublisherElement = toMapOrNil(m.Publisher.Element.MarshalJSON())
		}
	}
	output.Contact = m.Contact
	if m.Description != nil && m.Description.Value != nil {
		output.Description = m.Description.Value
		if m.Description.Element != nil {
			output.DescriptionElement = toMapOrNil(m.Description.Element.MarshalJSON())
		}
	}
	output.UseContext = m.UseContext
	output.Jurisdiction = m.Jurisdiction
	if m.Purpose != nil && m.Purpose.Value != nil {
		output.Purpose = m.Purpose.Value
		if m.Purpose.Element != nil {
			output.PurposeElement = toMapOrNil(m.Purpose.Element.MarshalJSON())
		}
	}
	if m.Start != nil && m.Start.Value != nil {
		output.Start = m.Start.Value
		if m.Start.Element != nil {
			output.StartElement = toMapOrNil(m.Start.Element.MarshalJSON())
		}
	}
	if m.Profile != nil && m.Profile.Value != nil {
		output.Profile = m.Profile.Value
		if m.Profile.Element != nil {
			output.ProfileElement = toMapOrNil(m.Profile.Element.MarshalJSON())
		}
	}
	output.Link = m.Link
	return json.Marshal(output)
}

// Clone creates a deep copy of GraphDefinition.
func (m *GraphDefinition) Clone() *GraphDefinition {
	if m == nil { return nil }
	return &GraphDefinition{
		Id: m.Id.Clone(),
		Meta: m.Meta.Clone(),
		ImplicitRules: m.ImplicitRules.Clone(),
		Language: m.Language.Clone(),
		Text: m.Text.Clone(),
		Contained: cloneSlices(m.Contained),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Url: m.Url.Clone(),
		Version: m.Version.Clone(),
		Name: m.Name.Clone(),
		Status: m.Status.Clone(),
		Experimental: m.Experimental.Clone(),
		Date: m.Date.Clone(),
		Publisher: m.Publisher.Clone(),
		Contact: cloneSlices(m.Contact),
		Description: m.Description.Clone(),
		UseContext: cloneSlices(m.UseContext),
		Jurisdiction: cloneSlices(m.Jurisdiction),
		Purpose: m.Purpose.Clone(),
		Start: m.Start.Clone(),
		Profile: m.Profile.Clone(),
		Link: cloneSlices(m.Link),
	}
}

// Equals checks equality between two GraphDefinition instances.
func (m *GraphDefinition) Equals(other *GraphDefinition) bool {
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
	if !m.Version.Equals(other.Version) { return false }
	if !m.Name.Equals(other.Name) { return false }
	if !m.Status.Equals(other.Status) { return false }
	if !m.Experimental.Equals(other.Experimental) { return false }
	if !m.Date.Equals(other.Date) { return false }
	if !m.Publisher.Equals(other.Publisher) { return false }
	if !compareSlices(m.Contact, other.Contact) { return false }
	if !m.Description.Equals(other.Description) { return false }
	if !compareSlices(m.UseContext, other.UseContext) { return false }
	if !compareSlices(m.Jurisdiction, other.Jurisdiction) { return false }
	if !m.Purpose.Equals(other.Purpose) { return false }
	if !m.Start.Equals(other.Start) { return false }
	if !m.Profile.Equals(other.Profile) { return false }
	if !compareSlices(m.Link, other.Link) { return false }
	return true
}

// GraphDefinitionLink
// Links this graph makes rules about.
type GraphDefinitionLink struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Path *FhirString `json:"path,omitempty"`
	SliceName *FhirString `json:"slicename,omitempty"`
	Min *FhirInteger `json:"min,omitempty"`
	Max *FhirString `json:"max,omitempty"`
	Description *FhirString `json:"description,omitempty"`
	Target []*GraphDefinitionTarget `json:"target,omitempty"`
}

// NewGraphDefinitionLink creates a new GraphDefinitionLink instance.
func NewGraphDefinitionLink() *GraphDefinitionLink {
	return &GraphDefinitionLink{}
}

// UnmarshalJSON populates GraphDefinitionLink from JSON data.
func (m *GraphDefinitionLink) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Path *FhirString `json:"path,omitempty"`
		SliceName *FhirString `json:"slicename,omitempty"`
		Min *FhirInteger `json:"min,omitempty"`
		Max *FhirString `json:"max,omitempty"`
		Description *FhirString `json:"description,omitempty"`
		Target []*GraphDefinitionTarget `json:"target,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Path = temp.Path
	m.SliceName = temp.SliceName
	m.Min = temp.Min
	m.Max = temp.Max
	m.Description = temp.Description
	m.Target = temp.Target
	return nil
}

// MarshalJSON converts GraphDefinitionLink to JSON data.
func (m *GraphDefinitionLink) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Path interface{} `json:"path,omitempty"`
		PathElement map[string]interface{} `json:"_path,omitempty"`
		SliceName interface{} `json:"slicename,omitempty"`
		SliceNameElement map[string]interface{} `json:"_slicename,omitempty"`
		Min interface{} `json:"min,omitempty"`
		MinElement map[string]interface{} `json:"_min,omitempty"`
		Max interface{} `json:"max,omitempty"`
		MaxElement map[string]interface{} `json:"_max,omitempty"`
		Description interface{} `json:"description,omitempty"`
		DescriptionElement map[string]interface{} `json:"_description,omitempty"`
		Target []*GraphDefinitionTarget `json:"target,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	if m.Path != nil && m.Path.Value != nil {
		output.Path = m.Path.Value
		if m.Path.Element != nil {
			output.PathElement = toMapOrNil(m.Path.Element.MarshalJSON())
		}
	}
	if m.SliceName != nil && m.SliceName.Value != nil {
		output.SliceName = m.SliceName.Value
		if m.SliceName.Element != nil {
			output.SliceNameElement = toMapOrNil(m.SliceName.Element.MarshalJSON())
		}
	}
	if m.Min != nil && m.Min.Value != nil {
		output.Min = m.Min.Value
		if m.Min.Element != nil {
			output.MinElement = toMapOrNil(m.Min.Element.MarshalJSON())
		}
	}
	if m.Max != nil && m.Max.Value != nil {
		output.Max = m.Max.Value
		if m.Max.Element != nil {
			output.MaxElement = toMapOrNil(m.Max.Element.MarshalJSON())
		}
	}
	if m.Description != nil && m.Description.Value != nil {
		output.Description = m.Description.Value
		if m.Description.Element != nil {
			output.DescriptionElement = toMapOrNil(m.Description.Element.MarshalJSON())
		}
	}
	output.Target = m.Target
	return json.Marshal(output)
}

// Clone creates a deep copy of GraphDefinitionLink.
func (m *GraphDefinitionLink) Clone() *GraphDefinitionLink {
	if m == nil { return nil }
	return &GraphDefinitionLink{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Path: m.Path.Clone(),
		SliceName: m.SliceName.Clone(),
		Min: m.Min.Clone(),
		Max: m.Max.Clone(),
		Description: m.Description.Clone(),
		Target: cloneSlices(m.Target),
	}
}

// Equals checks equality between two GraphDefinitionLink instances.
func (m *GraphDefinitionLink) Equals(other *GraphDefinitionLink) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Path.Equals(other.Path) { return false }
	if !m.SliceName.Equals(other.SliceName) { return false }
	if !m.Min.Equals(other.Min) { return false }
	if !m.Max.Equals(other.Max) { return false }
	if !m.Description.Equals(other.Description) { return false }
	if !compareSlices(m.Target, other.Target) { return false }
	return true
}

// GraphDefinitionTarget
// Potential target for the link.
type GraphDefinitionTarget struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Type *FhirCode `json:"type,omitempty"`
	Params *FhirString `json:"params,omitempty"`
	Profile *FhirCanonical `json:"profile,omitempty"`
	Compartment []*GraphDefinitionCompartment `json:"compartment,omitempty"`
	Link []*GraphDefinitionLink `json:"link,omitempty"`
}

// NewGraphDefinitionTarget creates a new GraphDefinitionTarget instance.
func NewGraphDefinitionTarget() *GraphDefinitionTarget {
	return &GraphDefinitionTarget{}
}

// UnmarshalJSON populates GraphDefinitionTarget from JSON data.
func (m *GraphDefinitionTarget) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Type *FhirCode `json:"type,omitempty"`
		Params *FhirString `json:"params,omitempty"`
		Profile *FhirCanonical `json:"profile,omitempty"`
		Compartment []*GraphDefinitionCompartment `json:"compartment,omitempty"`
		Link []*GraphDefinitionLink `json:"link,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Type = temp.Type
	m.Params = temp.Params
	m.Profile = temp.Profile
	m.Compartment = temp.Compartment
	m.Link = temp.Link
	return nil
}

// MarshalJSON converts GraphDefinitionTarget to JSON data.
func (m *GraphDefinitionTarget) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Type interface{} `json:"type,omitempty"`
		TypeElement map[string]interface{} `json:"_type,omitempty"`
		Params interface{} `json:"params,omitempty"`
		ParamsElement map[string]interface{} `json:"_params,omitempty"`
		Profile interface{} `json:"profile,omitempty"`
		ProfileElement map[string]interface{} `json:"_profile,omitempty"`
		Compartment []*GraphDefinitionCompartment `json:"compartment,omitempty"`
		Link []*GraphDefinitionLink `json:"link,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	if m.Type != nil && m.Type.Value != nil {
		output.Type = m.Type.Value
		if m.Type.Element != nil {
			output.TypeElement = toMapOrNil(m.Type.Element.MarshalJSON())
		}
	}
	if m.Params != nil && m.Params.Value != nil {
		output.Params = m.Params.Value
		if m.Params.Element != nil {
			output.ParamsElement = toMapOrNil(m.Params.Element.MarshalJSON())
		}
	}
	if m.Profile != nil && m.Profile.Value != nil {
		output.Profile = m.Profile.Value
		if m.Profile.Element != nil {
			output.ProfileElement = toMapOrNil(m.Profile.Element.MarshalJSON())
		}
	}
	output.Compartment = m.Compartment
	output.Link = m.Link
	return json.Marshal(output)
}

// Clone creates a deep copy of GraphDefinitionTarget.
func (m *GraphDefinitionTarget) Clone() *GraphDefinitionTarget {
	if m == nil { return nil }
	return &GraphDefinitionTarget{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Type: m.Type.Clone(),
		Params: m.Params.Clone(),
		Profile: m.Profile.Clone(),
		Compartment: cloneSlices(m.Compartment),
		Link: cloneSlices(m.Link),
	}
}

// Equals checks equality between two GraphDefinitionTarget instances.
func (m *GraphDefinitionTarget) Equals(other *GraphDefinitionTarget) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Type.Equals(other.Type) { return false }
	if !m.Params.Equals(other.Params) { return false }
	if !m.Profile.Equals(other.Profile) { return false }
	if !compareSlices(m.Compartment, other.Compartment) { return false }
	if !compareSlices(m.Link, other.Link) { return false }
	return true
}

// GraphDefinitionCompartment
// Compartment Consistency Rules.
type GraphDefinitionCompartment struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Use *GraphCompartmentUse `json:"use,omitempty"`
	Code *CompartmentType `json:"code,omitempty"`
	Rule *GraphCompartmentRule `json:"rule,omitempty"`
	Expression *FhirString `json:"expression,omitempty"`
	Description *FhirString `json:"description,omitempty"`
}

// NewGraphDefinitionCompartment creates a new GraphDefinitionCompartment instance.
func NewGraphDefinitionCompartment() *GraphDefinitionCompartment {
	return &GraphDefinitionCompartment{}
}

// UnmarshalJSON populates GraphDefinitionCompartment from JSON data.
func (m *GraphDefinitionCompartment) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Use *GraphCompartmentUse `json:"use,omitempty"`
		Code *CompartmentType `json:"code,omitempty"`
		Rule *GraphCompartmentRule `json:"rule,omitempty"`
		Expression *FhirString `json:"expression,omitempty"`
		Description *FhirString `json:"description,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Use = temp.Use
	m.Code = temp.Code
	m.Rule = temp.Rule
	m.Expression = temp.Expression
	m.Description = temp.Description
	return nil
}

// MarshalJSON converts GraphDefinitionCompartment to JSON data.
func (m *GraphDefinitionCompartment) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Use *GraphCompartmentUse `json:"use,omitempty"`
		Code *CompartmentType `json:"code,omitempty"`
		Rule *GraphCompartmentRule `json:"rule,omitempty"`
		Expression interface{} `json:"expression,omitempty"`
		ExpressionElement map[string]interface{} `json:"_expression,omitempty"`
		Description interface{} `json:"description,omitempty"`
		DescriptionElement map[string]interface{} `json:"_description,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.Use = m.Use
	output.Code = m.Code
	output.Rule = m.Rule
	if m.Expression != nil && m.Expression.Value != nil {
		output.Expression = m.Expression.Value
		if m.Expression.Element != nil {
			output.ExpressionElement = toMapOrNil(m.Expression.Element.MarshalJSON())
		}
	}
	if m.Description != nil && m.Description.Value != nil {
		output.Description = m.Description.Value
		if m.Description.Element != nil {
			output.DescriptionElement = toMapOrNil(m.Description.Element.MarshalJSON())
		}
	}
	return json.Marshal(output)
}

// Clone creates a deep copy of GraphDefinitionCompartment.
func (m *GraphDefinitionCompartment) Clone() *GraphDefinitionCompartment {
	if m == nil { return nil }
	return &GraphDefinitionCompartment{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Use: m.Use.Clone(),
		Code: m.Code.Clone(),
		Rule: m.Rule.Clone(),
		Expression: m.Expression.Clone(),
		Description: m.Description.Clone(),
	}
}

// Equals checks equality between two GraphDefinitionCompartment instances.
func (m *GraphDefinitionCompartment) Equals(other *GraphDefinitionCompartment) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Use.Equals(other.Use) { return false }
	if !m.Code.Equals(other.Code) { return false }
	if !m.Rule.Equals(other.Rule) { return false }
	if !m.Expression.Equals(other.Expression) { return false }
	if !m.Description.Equals(other.Description) { return false }
	return true
}

