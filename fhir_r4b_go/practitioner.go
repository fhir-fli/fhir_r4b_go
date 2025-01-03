// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"
)

// Practitioner
// A person who is directly or indirectly involved in the provisioning of healthcare.
type Practitioner struct {
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
	Active *FhirBoolean `json:"active,omitempty"`
	Name []*HumanName `json:"name,omitempty"`
	Telecom []*ContactPoint `json:"telecom,omitempty"`
	Address []*Address `json:"address,omitempty"`
	Gender *AdministrativeGender `json:"gender,omitempty"`
	BirthDate *FhirDate `json:"birthdate,omitempty"`
	Photo []*Attachment `json:"photo,omitempty"`
	Qualification []*PractitionerQualification `json:"qualification,omitempty"`
	Communication []*CodeableConcept `json:"communication,omitempty"`
}

// NewPractitioner creates a new Practitioner instance.
func NewPractitioner() *Practitioner {
	return &Practitioner{}
}

// UnmarshalJSON populates Practitioner from JSON data.
func (m *Practitioner) UnmarshalJSON(data []byte) error {
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
		Active *FhirBoolean `json:"active,omitempty"`
		Name []*HumanName `json:"name,omitempty"`
		Telecom []*ContactPoint `json:"telecom,omitempty"`
		Address []*Address `json:"address,omitempty"`
		Gender *AdministrativeGender `json:"gender,omitempty"`
		BirthDate *FhirDate `json:"birthdate,omitempty"`
		Photo []*Attachment `json:"photo,omitempty"`
		Qualification []*PractitionerQualification `json:"qualification,omitempty"`
		Communication []*CodeableConcept `json:"communication,omitempty"`
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
	m.Active = temp.Active
	m.Name = temp.Name
	m.Telecom = temp.Telecom
	m.Address = temp.Address
	m.Gender = temp.Gender
	m.BirthDate = temp.BirthDate
	m.Photo = temp.Photo
	m.Qualification = temp.Qualification
	m.Communication = temp.Communication
	return nil
}

// MarshalJSON converts Practitioner to JSON data.
func (m *Practitioner) MarshalJSON() ([]byte, error) {
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
		Active interface{} `json:"active,omitempty"`
		ActiveElement map[string]interface{} `json:"_active,omitempty"`
		Name []*HumanName `json:"name,omitempty"`
		Telecom []*ContactPoint `json:"telecom,omitempty"`
		Address []*Address `json:"address,omitempty"`
		Gender *AdministrativeGender `json:"gender,omitempty"`
		BirthDate interface{} `json:"birthdate,omitempty"`
		BirthDateElement map[string]interface{} `json:"_birthdate,omitempty"`
		Photo []*Attachment `json:"photo,omitempty"`
		Qualification []*PractitionerQualification `json:"qualification,omitempty"`
		Communication []*CodeableConcept `json:"communication,omitempty"`
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
	if m.Active != nil && m.Active.Value != nil {
		output.Active = m.Active.Value
		if m.Active.Element != nil {
			output.ActiveElement = toMapOrNil(m.Active.Element.MarshalJSON())
		}
	}
	output.Name = m.Name
	output.Telecom = m.Telecom
	output.Address = m.Address
	output.Gender = m.Gender
	if m.BirthDate != nil && m.BirthDate.Value != nil {
		output.BirthDate = m.BirthDate.Value
		if m.BirthDate.Element != nil {
			output.BirthDateElement = toMapOrNil(m.BirthDate.Element.MarshalJSON())
		}
	}
	output.Photo = m.Photo
	output.Qualification = m.Qualification
	output.Communication = m.Communication
	return json.Marshal(output)
}

// Clone creates a deep copy of Practitioner.
func (m *Practitioner) Clone() *Practitioner {
	if m == nil { return nil }
	return &Practitioner{
		Id: m.Id.Clone(),
		Meta: m.Meta.Clone(),
		ImplicitRules: m.ImplicitRules.Clone(),
		Language: m.Language.Clone(),
		Text: m.Text.Clone(),
		Contained: cloneSlices(m.Contained),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Identifier: cloneSlices(m.Identifier),
		Active: m.Active.Clone(),
		Name: cloneSlices(m.Name),
		Telecom: cloneSlices(m.Telecom),
		Address: cloneSlices(m.Address),
		Gender: m.Gender.Clone(),
		BirthDate: m.BirthDate.Clone(),
		Photo: cloneSlices(m.Photo),
		Qualification: cloneSlices(m.Qualification),
		Communication: cloneSlices(m.Communication),
	}
}

// Equals checks equality between two Practitioner instances.
func (m *Practitioner) Equals(other *Practitioner) bool {
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
	if !m.Active.Equals(other.Active) { return false }
	if !compareSlices(m.Name, other.Name) { return false }
	if !compareSlices(m.Telecom, other.Telecom) { return false }
	if !compareSlices(m.Address, other.Address) { return false }
	if !m.Gender.Equals(other.Gender) { return false }
	if !m.BirthDate.Equals(other.BirthDate) { return false }
	if !compareSlices(m.Photo, other.Photo) { return false }
	if !compareSlices(m.Qualification, other.Qualification) { return false }
	if !compareSlices(m.Communication, other.Communication) { return false }
	return true
}

// PractitionerQualification
// The official certifications, training, and licenses that authorize or otherwise pertain to the provision of care by the practitioner.  For example, a medical license issued by a medical board authorizing the practitioner to practice medicine within a certian locality.
type PractitionerQualification struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	Period *Period `json:"period,omitempty"`
	Issuer *Reference `json:"issuer,omitempty"`
}

// NewPractitionerQualification creates a new PractitionerQualification instance.
func NewPractitionerQualification() *PractitionerQualification {
	return &PractitionerQualification{}
}

// UnmarshalJSON populates PractitionerQualification from JSON data.
func (m *PractitionerQualification) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Identifier []*Identifier `json:"identifier,omitempty"`
		Code *CodeableConcept `json:"code,omitempty"`
		Period *Period `json:"period,omitempty"`
		Issuer *Reference `json:"issuer,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Identifier = temp.Identifier
	m.Code = temp.Code
	m.Period = temp.Period
	m.Issuer = temp.Issuer
	return nil
}

// MarshalJSON converts PractitionerQualification to JSON data.
func (m *PractitionerQualification) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Identifier []*Identifier `json:"identifier,omitempty"`
		Code *CodeableConcept `json:"code,omitempty"`
		Period *Period `json:"period,omitempty"`
		Issuer *Reference `json:"issuer,omitempty"`
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
	output.Code = m.Code
	output.Period = m.Period
	output.Issuer = m.Issuer
	return json.Marshal(output)
}

// Clone creates a deep copy of PractitionerQualification.
func (m *PractitionerQualification) Clone() *PractitionerQualification {
	if m == nil { return nil }
	return &PractitionerQualification{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Identifier: cloneSlices(m.Identifier),
		Code: m.Code.Clone(),
		Period: m.Period.Clone(),
		Issuer: m.Issuer.Clone(),
	}
}

// Equals checks equality between two PractitionerQualification instances.
func (m *PractitionerQualification) Equals(other *PractitionerQualification) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !compareSlices(m.Identifier, other.Identifier) { return false }
	if !m.Code.Equals(other.Code) { return false }
	if !m.Period.Equals(other.Period) { return false }
	if !m.Issuer.Equals(other.Issuer) { return false }
	return true
}

