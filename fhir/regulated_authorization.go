// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"

)

// RegulatedAuthorization
// Regulatory approval, clearance or licencing related to a regulated product, treatment, facility or activity that is cited in a guidance, regulation, rule or legislative act. An example is Market Authorization relating to a Medicinal Product.
type RegulatedAuthorization struct {
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
	// Business identifier for the authorization, typically assigned by the authorizing body.
	Identifier []Identifier `json:"identifier,omitempty"`
	// subject
	// The product type, treatment, facility or activity that is being authorized.
	Subject []Reference `json:"subject,omitempty"`
	// type
	// Overall type of this authorization, for example drug marketing approval, orphan drug designation.
	Type_ CodeableConcept `json:"type,omitempty"`
	// description
	// General textual supporting information.
	Description FhirMarkdown `json:"description,omitempty"`
	// region
	// The territory (e.g., country, jurisdiction etc.) in which the authorization has been granted.
	Region []CodeableConcept `json:"region,omitempty"`
	// status
	// The status that is authorised e.g. approved. Intermediate states and actions can be tracked with cases and applications.
	Status CodeableConcept `json:"status,omitempty"`
	// statusDate
	// The date at which the current status was assigned.
	StatusDate FhirDateTime `json:"statusDate,omitempty"`
	// validityPeriod
	// The time period in which the regulatory approval, clearance or licencing is in effect. As an example, a Marketing Authorization includes the date of authorization and/or an expiration date.
	ValidityPeriod Period `json:"validityPeriod,omitempty"`
	// indication
	// Condition for which the use of the regulated product applies.
	Indication CodeableReference `json:"indication,omitempty"`
	// intendedUse
	// The intended use of the product, e.g. prevention, treatment, diagnosis.
	IntendedUse CodeableConcept `json:"intendedUse,omitempty"`
	// basis
	// The legal or regulatory framework against which this authorization is granted, or other reasons for it.
	Basis []CodeableConcept `json:"basis,omitempty"`
	// holder
	// The organization that has been granted this authorization, by some authoritative body (the 'regulator').
	Holder Reference `json:"holder,omitempty"`
	// regulator
	// The regulatory authority or authorizing body granting the authorization. For example, European Medicines Agency (EMA), Food and Drug Administration (FDA), Health Canada (HC), etc.
	Regulator Reference `json:"regulator,omitempty"`
	// case
	// The case or regulatory procedure for granting or amending a regulated authorization. An authorization is granted in response to submissions/applications by those seeking authorization. A case is the administrative process that deals with the application(s) that relate to this and assesses them. Note: This area is subject to ongoing review and the workgroup is seeking implementer feedback on its use (see link at bottom of page).
	Case_ RegulatedAuthorizationCase `json:"case,omitempty"`
}

// NewRegulatedAuthorization creates a new RegulatedAuthorization instance
func NewRegulatedAuthorization(
	id FhirString,
	meta FhirMeta,
	implicitRules FhirUri,
	language CommonLanguages,
	text Narrative,
	contained []Resource,
	extension_ []FhirExtension,
	modifierExtension []FhirExtension,
	identifier []Identifier,
	subject []Reference,
	type_ CodeableConcept,
	description FhirMarkdown,
	region []CodeableConcept,
	status CodeableConcept,
	statusDate FhirDateTime,
	validityPeriod Period,
	indication CodeableReference,
	intendedUse CodeableConcept,
	basis []CodeableConcept,
	holder Reference,
	regulator Reference,
	case_ RegulatedAuthorizationCase,
) *RegulatedAuthorization {
	return &RegulatedAuthorization{
		Id: id,
		Meta: meta,
		ImplicitRules: implicitRules,
		Language: language,
		Text: text,
		Contained: contained,
		Extension_: extension_,
		ModifierExtension: modifierExtension,
		Identifier: identifier,
		Subject: subject,
		Type_: type_,
		Description: description,
		Region: region,
		Status: status,
		StatusDate: statusDate,
		ValidityPeriod: validityPeriod,
		Indication: indication,
		IntendedUse: intendedUse,
		Basis: basis,
		Holder: holder,
		Regulator: regulator,
		Case_: case_,
	}
}
// FromJSON populates RegulatedAuthorization from JSON data
func (m *RegulatedAuthorization) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts RegulatedAuthorization to JSON data
func (m *RegulatedAuthorization) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// CopyWith creates a modified copy of RegulatedAuthorization
func (m *RegulatedAuthorization) CopyWith(
	id *FhirString,
	meta *FhirMeta,
	implicitRules *FhirUri,
	language *CommonLanguages,
	text *Narrative,
	contained *[]Resource,
	extension_ *[]FhirExtension,
	modifierExtension *[]FhirExtension,
	identifier *[]Identifier,
	subject *[]Reference,
	type_ *CodeableConcept,
	description *FhirMarkdown,
	region *[]CodeableConcept,
	status *CodeableConcept,
	statusDate *FhirDateTime,
	validityPeriod *Period,
	indication *CodeableReference,
	intendedUse *CodeableConcept,
	basis *[]CodeableConcept,
	holder *Reference,
	regulator *Reference,
	case_ *RegulatedAuthorizationCase,
) *RegulatedAuthorization {
	return &RegulatedAuthorization{
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
		Subject: func() []Reference {
			if subject != nil { return *subject }
			return m.Subject
		}(),
		Type_: func() CodeableConcept {
			if type_ != nil { return *type_ }
			return m.Type_
		}(),
		Description: func() FhirMarkdown {
			if description != nil { return *description }
			return m.Description
		}(),
		Region: func() []CodeableConcept {
			if region != nil { return *region }
			return m.Region
		}(),
		Status: func() CodeableConcept {
			if status != nil { return *status }
			return m.Status
		}(),
		StatusDate: func() FhirDateTime {
			if statusDate != nil { return *statusDate }
			return m.StatusDate
		}(),
		ValidityPeriod: func() Period {
			if validityPeriod != nil { return *validityPeriod }
			return m.ValidityPeriod
		}(),
		Indication: func() CodeableReference {
			if indication != nil { return *indication }
			return m.Indication
		}(),
		IntendedUse: func() CodeableConcept {
			if intendedUse != nil { return *intendedUse }
			return m.IntendedUse
		}(),
		Basis: func() []CodeableConcept {
			if basis != nil { return *basis }
			return m.Basis
		}(),
		Holder: func() Reference {
			if holder != nil { return *holder }
			return m.Holder
		}(),
		Regulator: func() Reference {
			if regulator != nil { return *regulator }
			return m.Regulator
		}(),
		Case_: func() RegulatedAuthorizationCase {
			if case_ != nil { return *case_ }
			return m.Case_
		}(),
	}
}
// RegulatedAuthorizationCase
// The case or regulatory procedure for granting or amending a regulated authorization. An authorization is granted in response to submissions/applications by those seeking authorization. A case is the administrative process that deals with the application(s) that relate to this and assesses them. Note: This area is subject to ongoing review and the workgroup is seeking implementer feedback on its use (see link at bottom of page).
type RegulatedAuthorizationCase struct {
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
	// identifier
	// Identifier by which this case can be referenced.
	Identifier Identifier `json:"identifier,omitempty"`
	// type
	// The defining type of case.
	Type_ CodeableConcept `json:"type,omitempty"`
	// status
	// The status associated with the case.
	Status CodeableConcept `json:"status,omitempty"`
	// datePeriod
	// Relevant date for this case.
	DatePeriod Period `json:"datePeriod,omitempty"`
	// dateDateTime
	// Relevant date for this case.
	DateDateTime FhirDateTime `json:"dateDateTime,omitempty"`
	// application
	// A regulatory submission from an organization to a regulator, as part of an assessing case. Multiple applications may occur over time, with more or different information to support or modify the submission or the authorization. The applications can be considered as steps within the longer running case or procedure for this authorization process.
	Application []RegulatedAuthorizationCase `json:"application,omitempty"`
}

// NewRegulatedAuthorizationCase creates a new RegulatedAuthorizationCase instance
func NewRegulatedAuthorizationCase(
	id FhirString,
	extension_ []FhirExtension,
	modifierExtension []FhirExtension,
	identifier Identifier,
	type_ CodeableConcept,
	status CodeableConcept,
	datePeriod Period,
	dateDateTime FhirDateTime,
	application []RegulatedAuthorizationCase,
) *RegulatedAuthorizationCase {
	return &RegulatedAuthorizationCase{
		Id: id,
		Extension_: extension_,
		ModifierExtension: modifierExtension,
		Identifier: identifier,
		Type_: type_,
		Status: status,
		DatePeriod: datePeriod,
		DateDateTime: dateDateTime,
		Application: application,
	}
}
// FromJSON populates RegulatedAuthorizationCase from JSON data
func (m *RegulatedAuthorizationCase) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts RegulatedAuthorizationCase to JSON data
func (m *RegulatedAuthorizationCase) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// CopyWith creates a modified copy of RegulatedAuthorizationCase
func (m *RegulatedAuthorizationCase) CopyWith(
	id *FhirString,
	extension_ *[]FhirExtension,
	modifierExtension *[]FhirExtension,
	identifier *Identifier,
	type_ *CodeableConcept,
	status *CodeableConcept,
	datePeriod *Period,
	dateDateTime *FhirDateTime,
	application *[]RegulatedAuthorizationCase,
) *RegulatedAuthorizationCase {
	return &RegulatedAuthorizationCase{
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
		Identifier: func() Identifier {
			if identifier != nil { return *identifier }
			return m.Identifier
		}(),
		Type_: func() CodeableConcept {
			if type_ != nil { return *type_ }
			return m.Type_
		}(),
		Status: func() CodeableConcept {
			if status != nil { return *status }
			return m.Status
		}(),
		DatePeriod: func() Period {
			if datePeriod != nil { return *datePeriod }
			return m.DatePeriod
		}(),
		DateDateTime: func() FhirDateTime {
			if dateDateTime != nil { return *dateDateTime }
			return m.DateDateTime
		}(),
		Application: func() []RegulatedAuthorizationCase {
			if application != nil { return *application }
			return m.Application
		}(),
	}
}