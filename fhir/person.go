// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"

)

// Person
// Demographics and administrative information about a person independent of a specific health-related context.
type Person struct {
	DomainResource
	// id
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id FhirString `json:"id,omitempty"`
	// meta
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta FhirMeta `json:"meta,omitempty"`
	// implicitRules
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules FhirUri `json:"implicitRules,omitempty"`
	// language
	// The base language in which the resource is written.
	Language CommonLanguages `json:"language,omitempty"`
	// text
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text Narrative `json:"text,omitempty"`
	// contained
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []Resource `json:"contained,omitempty"`
	// extension
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension_ []FhirExtension `json:"extension,omitempty"`
	// modifierExtension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
// 
// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []FhirExtension `json:"modifierExtension,omitempty"`
	// identifier
	// Identifier for a person within a particular scope.
	Identifier []Identifier `json:"identifier,omitempty"`
	// name
	// A name associated with the person.
	Name []HumanName `json:"name,omitempty"`
	// telecom
	// A contact detail for the person, e.g. a telephone number or an email address.
	Telecom []ContactPoint `json:"telecom,omitempty"`
	// gender
	// Administrative Gender.
	Gender AdministrativeGender `json:"gender,omitempty"`
	// birthDate
	// The birth date for the person.
	BirthDate FhirDate `json:"birthDate,omitempty"`
	// address
	// One or more addresses for the person.
	Address []Address `json:"address,omitempty"`
	// photo
	// An image that can be displayed as a thumbnail of the person to enhance the identification of the individual.
	Photo Attachment `json:"photo,omitempty"`
	// managingOrganization
	// The organization that is the custodian of the person record.
	ManagingOrganization Reference `json:"managingOrganization,omitempty"`
	// active
	// Whether this person's record is in active use.
	Active FhirBoolean `json:"active,omitempty"`
	// link
	// Link to a resource that concerns the same actual person.
	Link []PersonLink `json:"link,omitempty"`
}

// NewPerson creates a new Person instance
func NewPerson(
	id FhirString,
	meta FhirMeta,
	implicitRules FhirUri,
	language CommonLanguages,
	text Narrative,
	contained []Resource,
	extension_ []FhirExtension,
	modifierExtension []FhirExtension,
	identifier []Identifier,
	name []HumanName,
	telecom []ContactPoint,
	gender AdministrativeGender,
	birthDate FhirDate,
	address []Address,
	photo Attachment,
	managingOrganization Reference,
	active FhirBoolean,
	link []PersonLink,
) *Person {
	return &Person{
		Id: id,
		Meta: meta,
		ImplicitRules: implicitRules,
		Language: language,
		Text: text,
		Contained: contained,
		Extension_: extension_,
		ModifierExtension: modifierExtension,
		Identifier: identifier,
		Name: name,
		Telecom: telecom,
		Gender: gender,
		BirthDate: birthDate,
		Address: address,
		Photo: photo,
		ManagingOrganization: managingOrganization,
		Active: active,
		Link: link,
	}
}
// FromJSON populates Person from JSON data
func (m *Person) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts Person to JSON data
func (m *Person) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// CopyWith creates a modified copy of Person
func (m *Person) CopyWith(
	id *FhirString,
	meta *FhirMeta,
	implicitRules *FhirUri,
	language *CommonLanguages,
	text *Narrative,
	contained *[]Resource,
	extension_ *[]FhirExtension,
	modifierExtension *[]FhirExtension,
	identifier *[]Identifier,
	name *[]HumanName,
	telecom *[]ContactPoint,
	gender *AdministrativeGender,
	birthDate *FhirDate,
	address *[]Address,
	photo *Attachment,
	managingOrganization *Reference,
	active *FhirBoolean,
	link *[]PersonLink,
) *Person {
	return &Person{
		Id: func() FhirString {
			if id != nil { return *id }
			return m.Id
		}(),
		Meta: func() FhirMeta {
			if meta != nil { return *meta }
			return m.Meta
		}(),
		ImplicitRules: func() FhirUri {
			if implicitRules != nil { return *implicitRules }
			return m.ImplicitRules
		}(),
		Language: func() CommonLanguages {
			if language != nil { return *language }
			return m.Language
		}(),
		Text: func() Narrative {
			if text != nil { return *text }
			return m.Text
		}(),
		Contained: func() []Resource {
			if contained != nil { return *contained }
			return m.Contained
		}(),
		Extension_: func() []FhirExtension {
			if extension_ != nil { return *extension_ }
			return m.Extension_
		}(),
		ModifierExtension: func() []FhirExtension {
			if modifierExtension != nil { return *modifierExtension }
			return m.ModifierExtension
		}(),
		Identifier: func() []Identifier {
			if identifier != nil { return *identifier }
			return m.Identifier
		}(),
		Name: func() []HumanName {
			if name != nil { return *name }
			return m.Name
		}(),
		Telecom: func() []ContactPoint {
			if telecom != nil { return *telecom }
			return m.Telecom
		}(),
		Gender: func() AdministrativeGender {
			if gender != nil { return *gender }
			return m.Gender
		}(),
		BirthDate: func() FhirDate {
			if birthDate != nil { return *birthDate }
			return m.BirthDate
		}(),
		Address: func() []Address {
			if address != nil { return *address }
			return m.Address
		}(),
		Photo: func() Attachment {
			if photo != nil { return *photo }
			return m.Photo
		}(),
		ManagingOrganization: func() Reference {
			if managingOrganization != nil { return *managingOrganization }
			return m.ManagingOrganization
		}(),
		Active: func() FhirBoolean {
			if active != nil { return *active }
			return m.Active
		}(),
		Link: func() []PersonLink {
			if link != nil { return *link }
			return m.Link
		}(),
	}
}
// PersonLink
// Link to a resource that concerns the same actual person.
type PersonLink struct {
	BackboneElement
	// id
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id FhirString `json:"id,omitempty"`
	// extension
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension_ []FhirExtension `json:"extension,omitempty"`
	// modifierExtension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
// 
// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []FhirExtension `json:"modifierExtension,omitempty"`
	// target
	// The resource to which this actual person is associated.
	Target Reference `json:"target,omitempty"`
	// assurance
	// Level of assurance that this link is associated with the target resource.
	Assurance IdentityAssuranceLevel `json:"assurance,omitempty"`
}

// NewPersonLink creates a new PersonLink instance
func NewPersonLink(
	id FhirString,
	extension_ []FhirExtension,
	modifierExtension []FhirExtension,
	target Reference,
	assurance IdentityAssuranceLevel,
) *PersonLink {
	return &PersonLink{
		Id: id,
		Extension_: extension_,
		ModifierExtension: modifierExtension,
		Target: target,
		Assurance: assurance,
	}
}
// FromJSON populates PersonLink from JSON data
func (m *PersonLink) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts PersonLink to JSON data
func (m *PersonLink) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// CopyWith creates a modified copy of PersonLink
func (m *PersonLink) CopyWith(
	id *FhirString,
	extension_ *[]FhirExtension,
	modifierExtension *[]FhirExtension,
	target *Reference,
	assurance *IdentityAssuranceLevel,
) *PersonLink {
	return &PersonLink{
		Id: func() FhirString {
			if id != nil { return *id }
			return m.Id
		}(),
		Extension_: func() []FhirExtension {
			if extension_ != nil { return *extension_ }
			return m.Extension_
		}(),
		ModifierExtension: func() []FhirExtension {
			if modifierExtension != nil { return *modifierExtension }
			return m.ModifierExtension
		}(),
		Target: func() Reference {
			if target != nil { return *target }
			return m.Target
		}(),
		Assurance: func() IdentityAssuranceLevel {
			if assurance != nil { return *assurance }
			return m.Assurance
		}(),
	}
}