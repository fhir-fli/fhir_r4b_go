// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"

)

// Account
// A financial tool for tracking value accrued for a particular purpose.  In the healthcare field, used to track charges for a patient, cost centers, etc.
type Account struct {
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
	// Unique identifier used to reference the account.  Might or might not be intended for human use (e.g. credit card number).
	Identifier []Identifier `json:"identifier,omitempty"`
	// status
	// Indicates whether the account is presently used/usable or not.
	Status AccountStatus `json:"status,omitempty"`
	// type
	// Categorizes the account for reporting and searching purposes.
	Type_ CodeableConcept `json:"type,omitempty"`
	// name
	// Name used for the account when displaying it to humans in reports, etc.
	Name FhirString `json:"name,omitempty"`
	// subject
	// Identifies the entity which incurs the expenses. While the immediate recipients of services or goods might be entities related to the subject, the expenses were ultimately incurred by the subject of the Account.
	Subject []Reference `json:"subject,omitempty"`
	// servicePeriod
	// The date range of services associated with this account.
	ServicePeriod Period `json:"servicePeriod,omitempty"`
	// coverage
	// The party(s) that are responsible for covering the payment of this account, and what order should they be applied to the account.
	Coverage []AccountCoverage `json:"coverage,omitempty"`
	// owner
	// Indicates the service area, hospital, department, etc. with responsibility for managing the Account.
	Owner Reference `json:"owner,omitempty"`
	// description
	// Provides additional information about what the account tracks and how it is used.
	Description FhirString `json:"description,omitempty"`
	// guarantor
	// The parties responsible for balancing the account if other payment options fall short.
	Guarantor []AccountGuarantor `json:"guarantor,omitempty"`
	// partOf
	// Reference to a parent Account.
	PartOf Reference `json:"partOf,omitempty"`
}

// NewAccount creates a new Account instance
func NewAccount(
	id FhirString,
	meta FhirMeta,
	implicitRules FhirUri,
	language CommonLanguages,
	text Narrative,
	contained []Resource,
	extension_ []FhirExtension,
	modifierExtension []FhirExtension,
	identifier []Identifier,
	status AccountStatus,
	type_ CodeableConcept,
	name FhirString,
	subject []Reference,
	servicePeriod Period,
	coverage []AccountCoverage,
	owner Reference,
	description FhirString,
	guarantor []AccountGuarantor,
	partOf Reference,
) *Account {
	return &Account{
		Id: id,
		Meta: meta,
		ImplicitRules: implicitRules,
		Language: language,
		Text: text,
		Contained: contained,
		Extension_: extension_,
		ModifierExtension: modifierExtension,
		Identifier: identifier,
		Status: status,
		Type_: type_,
		Name: name,
		Subject: subject,
		ServicePeriod: servicePeriod,
		Coverage: coverage,
		Owner: owner,
		Description: description,
		Guarantor: guarantor,
		PartOf: partOf,
	}
}
// FromJSON populates Account from JSON data
func (m *Account) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts Account to JSON data
func (m *Account) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// CopyWith creates a modified copy of Account
func (m *Account) CopyWith(
	id *FhirString,
	meta *FhirMeta,
	implicitRules *FhirUri,
	language *CommonLanguages,
	text *Narrative,
	contained *[]Resource,
	extension_ *[]FhirExtension,
	modifierExtension *[]FhirExtension,
	identifier *[]Identifier,
	status *AccountStatus,
	type_ *CodeableConcept,
	name *FhirString,
	subject *[]Reference,
	servicePeriod *Period,
	coverage *[]AccountCoverage,
	owner *Reference,
	description *FhirString,
	guarantor *[]AccountGuarantor,
	partOf *Reference,
) *Account {
	return &Account{
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
		Status: func() AccountStatus {
			if status != nil { return *status }
			return m.Status
		}(),
		Type_: func() CodeableConcept {
			if type_ != nil { return *type_ }
			return m.Type_
		}(),
		Name: func() FhirString {
			if name != nil { return *name }
			return m.Name
		}(),
		Subject: func() []Reference {
			if subject != nil { return *subject }
			return m.Subject
		}(),
		ServicePeriod: func() Period {
			if servicePeriod != nil { return *servicePeriod }
			return m.ServicePeriod
		}(),
		Coverage: func() []AccountCoverage {
			if coverage != nil { return *coverage }
			return m.Coverage
		}(),
		Owner: func() Reference {
			if owner != nil { return *owner }
			return m.Owner
		}(),
		Description: func() FhirString {
			if description != nil { return *description }
			return m.Description
		}(),
		Guarantor: func() []AccountGuarantor {
			if guarantor != nil { return *guarantor }
			return m.Guarantor
		}(),
		PartOf: func() Reference {
			if partOf != nil { return *partOf }
			return m.PartOf
		}(),
	}
}
// AccountCoverage
// The party(s) that are responsible for covering the payment of this account, and what order should they be applied to the account.
type AccountCoverage struct {
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
	// coverage
	// The party(s) that contribute to payment (or part of) of the charges applied to this account (including self-pay).
// 
// A coverage may only be responsible for specific types of charges, and the sequence of the coverages in the account could be important when processing billing.
	Coverage Reference `json:"coverage,omitempty"`
	// priority
	// The priority of the coverage in the context of this account.
	Priority FhirPositiveInt `json:"priority,omitempty"`
}

// NewAccountCoverage creates a new AccountCoverage instance
func NewAccountCoverage(
	id FhirString,
	extension_ []FhirExtension,
	modifierExtension []FhirExtension,
	coverage Reference,
	priority FhirPositiveInt,
) *AccountCoverage {
	return &AccountCoverage{
		Id: id,
		Extension_: extension_,
		ModifierExtension: modifierExtension,
		Coverage: coverage,
		Priority: priority,
	}
}
// FromJSON populates AccountCoverage from JSON data
func (m *AccountCoverage) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts AccountCoverage to JSON data
func (m *AccountCoverage) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// CopyWith creates a modified copy of AccountCoverage
func (m *AccountCoverage) CopyWith(
	id *FhirString,
	extension_ *[]FhirExtension,
	modifierExtension *[]FhirExtension,
	coverage *Reference,
	priority *FhirPositiveInt,
) *AccountCoverage {
	return &AccountCoverage{
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
		Coverage: func() Reference {
			if coverage != nil { return *coverage }
			return m.Coverage
		}(),
		Priority: func() FhirPositiveInt {
			if priority != nil { return *priority }
			return m.Priority
		}(),
	}
}
// AccountGuarantor
// The parties responsible for balancing the account if other payment options fall short.
type AccountGuarantor struct {
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
	// party
	// The entity who is responsible.
	Party Reference `json:"party,omitempty"`
	// onHold
	// A guarantor may be placed on credit hold or otherwise have their role temporarily suspended.
	OnHold FhirBoolean `json:"onHold,omitempty"`
	// period
	// The timeframe during which the guarantor accepts responsibility for the account.
	Period Period `json:"period,omitempty"`
}

// NewAccountGuarantor creates a new AccountGuarantor instance
func NewAccountGuarantor(
	id FhirString,
	extension_ []FhirExtension,
	modifierExtension []FhirExtension,
	party Reference,
	onHold FhirBoolean,
	period Period,
) *AccountGuarantor {
	return &AccountGuarantor{
		Id: id,
		Extension_: extension_,
		ModifierExtension: modifierExtension,
		Party: party,
		OnHold: onHold,
		Period: period,
	}
}
// FromJSON populates AccountGuarantor from JSON data
func (m *AccountGuarantor) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts AccountGuarantor to JSON data
func (m *AccountGuarantor) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// CopyWith creates a modified copy of AccountGuarantor
func (m *AccountGuarantor) CopyWith(
	id *FhirString,
	extension_ *[]FhirExtension,
	modifierExtension *[]FhirExtension,
	party *Reference,
	onHold *FhirBoolean,
	period *Period,
) *AccountGuarantor {
	return &AccountGuarantor{
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
		Party: func() Reference {
			if party != nil { return *party }
			return m.Party
		}(),
		OnHold: func() FhirBoolean {
			if onHold != nil { return *onHold }
			return m.OnHold
		}(),
		Period: func() Period {
			if period != nil { return *period }
			return m.Period
		}(),
	}
}