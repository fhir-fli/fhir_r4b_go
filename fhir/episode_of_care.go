// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"

)

// EpisodeOfCare
// An association between a patient and an organization / healthcare provider(s) during which time encounters may occur. The managing organization assumes a level of responsibility for the patient during this time.
type EpisodeOfCare struct {
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
	// The EpisodeOfCare may be known by different identifiers for different contexts of use, such as when an external agency is tracking the Episode for funding purposes.
	Identifier []Identifier `json:"identifier,omitempty"`
	// status
	// planned | waitlist | active | onhold | finished | cancelled.
	Status EpisodeOfCareStatus `json:"status,omitempty"`
	// statusHistory
	// The history of statuses that the EpisodeOfCare has been through (without requiring processing the history of the resource).
	StatusHistory []EpisodeOfCareStatusHistory `json:"statusHistory,omitempty"`
	// type
	// A classification of the type of episode of care; e.g. specialist referral, disease management, type of funded care.
	Type_ []CodeableConcept `json:"type,omitempty"`
	// diagnosis
	// The list of diagnosis relevant to this episode of care.
	Diagnosis []EpisodeOfCareDiagnosis `json:"diagnosis,omitempty"`
	// patient
	// The patient who is the focus of this episode of care.
	Patient Reference `json:"patient,omitempty"`
	// managingOrganization
	// The organization that has assumed the specific responsibilities for the specified duration.
	ManagingOrganization Reference `json:"managingOrganization,omitempty"`
	// period
	// The interval during which the managing organization assumes the defined responsibility.
	Period Period `json:"period,omitempty"`
	// referralRequest
	// Referral Request(s) that are fulfilled by this EpisodeOfCare, incoming referrals.
	ReferralRequest []Reference `json:"referralRequest,omitempty"`
	// careManager
	// The practitioner that is the care manager/care coordinator for this patient.
	CareManager Reference `json:"careManager,omitempty"`
	// team
	// The list of practitioners that may be facilitating this episode of care for specific purposes.
	Team []Reference `json:"team,omitempty"`
	// account
	// The set of accounts that may be used for billing for this EpisodeOfCare.
	Account []Reference `json:"account,omitempty"`
}

// NewEpisodeOfCare creates a new EpisodeOfCare instance
func NewEpisodeOfCare(
	id FhirString,
	meta FhirMeta,
	implicitRules FhirUri,
	language CommonLanguages,
	text Narrative,
	contained []Resource,
	extension_ []FhirExtension,
	modifierExtension []FhirExtension,
	identifier []Identifier,
	status EpisodeOfCareStatus,
	statusHistory []EpisodeOfCareStatusHistory,
	type_ []CodeableConcept,
	diagnosis []EpisodeOfCareDiagnosis,
	patient Reference,
	managingOrganization Reference,
	period Period,
	referralRequest []Reference,
	careManager Reference,
	team []Reference,
	account []Reference,
) *EpisodeOfCare {
	return &EpisodeOfCare{
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
		StatusHistory: statusHistory,
		Type_: type_,
		Diagnosis: diagnosis,
		Patient: patient,
		ManagingOrganization: managingOrganization,
		Period: period,
		ReferralRequest: referralRequest,
		CareManager: careManager,
		Team: team,
		Account: account,
	}
}
// FromJSON populates EpisodeOfCare from JSON data
func (m *EpisodeOfCare) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts EpisodeOfCare to JSON data
func (m *EpisodeOfCare) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// CopyWith creates a modified copy of EpisodeOfCare
func (m *EpisodeOfCare) CopyWith(
	id *FhirString,
	meta *FhirMeta,
	implicitRules *FhirUri,
	language *CommonLanguages,
	text *Narrative,
	contained *[]Resource,
	extension_ *[]FhirExtension,
	modifierExtension *[]FhirExtension,
	identifier *[]Identifier,
	status *EpisodeOfCareStatus,
	statusHistory *[]EpisodeOfCareStatusHistory,
	type_ *[]CodeableConcept,
	diagnosis *[]EpisodeOfCareDiagnosis,
	patient *Reference,
	managingOrganization *Reference,
	period *Period,
	referralRequest *[]Reference,
	careManager *Reference,
	team *[]Reference,
	account *[]Reference,
) *EpisodeOfCare {
	return &EpisodeOfCare{
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
		Status: func() EpisodeOfCareStatus {
			if status != nil { return *status }
			return m.Status
		}(),
		StatusHistory: func() []EpisodeOfCareStatusHistory {
			if statusHistory != nil { return *statusHistory }
			return m.StatusHistory
		}(),
		Type_: func() []CodeableConcept {
			if type_ != nil { return *type_ }
			return m.Type_
		}(),
		Diagnosis: func() []EpisodeOfCareDiagnosis {
			if diagnosis != nil { return *diagnosis }
			return m.Diagnosis
		}(),
		Patient: func() Reference {
			if patient != nil { return *patient }
			return m.Patient
		}(),
		ManagingOrganization: func() Reference {
			if managingOrganization != nil { return *managingOrganization }
			return m.ManagingOrganization
		}(),
		Period: func() Period {
			if period != nil { return *period }
			return m.Period
		}(),
		ReferralRequest: func() []Reference {
			if referralRequest != nil { return *referralRequest }
			return m.ReferralRequest
		}(),
		CareManager: func() Reference {
			if careManager != nil { return *careManager }
			return m.CareManager
		}(),
		Team: func() []Reference {
			if team != nil { return *team }
			return m.Team
		}(),
		Account: func() []Reference {
			if account != nil { return *account }
			return m.Account
		}(),
	}
}
// EpisodeOfCareStatusHistory
// The history of statuses that the EpisodeOfCare has been through (without requiring processing the history of the resource).
type EpisodeOfCareStatusHistory struct {
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
	// status
	// planned | waitlist | active | onhold | finished | cancelled.
	Status EpisodeOfCareStatus `json:"status,omitempty"`
	// period
	// The period during this EpisodeOfCare that the specific status applied.
	Period Period `json:"period,omitempty"`
}

// NewEpisodeOfCareStatusHistory creates a new EpisodeOfCareStatusHistory instance
func NewEpisodeOfCareStatusHistory(
	id FhirString,
	extension_ []FhirExtension,
	modifierExtension []FhirExtension,
	status EpisodeOfCareStatus,
	period Period,
) *EpisodeOfCareStatusHistory {
	return &EpisodeOfCareStatusHistory{
		Id: id,
		Extension_: extension_,
		ModifierExtension: modifierExtension,
		Status: status,
		Period: period,
	}
}
// FromJSON populates EpisodeOfCareStatusHistory from JSON data
func (m *EpisodeOfCareStatusHistory) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts EpisodeOfCareStatusHistory to JSON data
func (m *EpisodeOfCareStatusHistory) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// CopyWith creates a modified copy of EpisodeOfCareStatusHistory
func (m *EpisodeOfCareStatusHistory) CopyWith(
	id *FhirString,
	extension_ *[]FhirExtension,
	modifierExtension *[]FhirExtension,
	status *EpisodeOfCareStatus,
	period *Period,
) *EpisodeOfCareStatusHistory {
	return &EpisodeOfCareStatusHistory{
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
		Status: func() EpisodeOfCareStatus {
			if status != nil { return *status }
			return m.Status
		}(),
		Period: func() Period {
			if period != nil { return *period }
			return m.Period
		}(),
	}
}
// EpisodeOfCareDiagnosis
// The list of diagnosis relevant to this episode of care.
type EpisodeOfCareDiagnosis struct {
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
	// condition
	// A list of conditions/problems/diagnoses that this episode of care is intended to be providing care for.
	Condition Reference `json:"condition,omitempty"`
	// role
	// Role that this diagnosis has within the episode of care (e.g. admission, billing, discharge …).
	Role CodeableConcept `json:"role,omitempty"`
	// rank
	// Ranking of the diagnosis (for each role type).
	Rank FhirPositiveInt `json:"rank,omitempty"`
}

// NewEpisodeOfCareDiagnosis creates a new EpisodeOfCareDiagnosis instance
func NewEpisodeOfCareDiagnosis(
	id FhirString,
	extension_ []FhirExtension,
	modifierExtension []FhirExtension,
	condition Reference,
	role CodeableConcept,
	rank FhirPositiveInt,
) *EpisodeOfCareDiagnosis {
	return &EpisodeOfCareDiagnosis{
		Id: id,
		Extension_: extension_,
		ModifierExtension: modifierExtension,
		Condition: condition,
		Role: role,
		Rank: rank,
	}
}
// FromJSON populates EpisodeOfCareDiagnosis from JSON data
func (m *EpisodeOfCareDiagnosis) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts EpisodeOfCareDiagnosis to JSON data
func (m *EpisodeOfCareDiagnosis) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// CopyWith creates a modified copy of EpisodeOfCareDiagnosis
func (m *EpisodeOfCareDiagnosis) CopyWith(
	id *FhirString,
	extension_ *[]FhirExtension,
	modifierExtension *[]FhirExtension,
	condition *Reference,
	role *CodeableConcept,
	rank *FhirPositiveInt,
) *EpisodeOfCareDiagnosis {
	return &EpisodeOfCareDiagnosis{
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
		Condition: func() Reference {
			if condition != nil { return *condition }
			return m.Condition
		}(),
		Role: func() CodeableConcept {
			if role != nil { return *role }
			return m.Role
		}(),
		Rank: func() FhirPositiveInt {
			if rank != nil { return *rank }
			return m.Rank
		}(),
	}
}