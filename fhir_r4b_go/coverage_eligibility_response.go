// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"
)

// CoverageEligibilityResponse
// This resource provides eligibility and plan details from the processing of an CoverageEligibilityRequest resource.
type CoverageEligibilityResponse struct {
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
	Status *FinancialResourceStatusCodes `json:"status,omitempty"`
	Purpose []*EligibilityResponsePurpose `json:"purpose,omitempty"`
	Patient *Reference `json:"patient,omitempty"`
	ServicedDate *FhirDate `json:"serviceddate,omitempty"`
	ServicedPeriod *Period `json:"servicedperiod,omitempty"`
	Created *FhirDateTime `json:"created,omitempty"`
	Requestor *Reference `json:"requestor,omitempty"`
	Request *Reference `json:"request,omitempty"`
	Outcome *RemittanceOutcome `json:"outcome,omitempty"`
	Disposition *FhirString `json:"disposition,omitempty"`
	Insurer *Reference `json:"insurer,omitempty"`
	Insurance []*CoverageEligibilityResponseInsurance `json:"insurance,omitempty"`
	PreAuthRef *FhirString `json:"preauthref,omitempty"`
	Form *CodeableConcept `json:"form,omitempty"`
	Error []*CoverageEligibilityResponseError `json:"error,omitempty"`
}

// NewCoverageEligibilityResponse creates a new CoverageEligibilityResponse instance.
func NewCoverageEligibilityResponse() *CoverageEligibilityResponse {
	return &CoverageEligibilityResponse{}
}

// UnmarshalJSON populates CoverageEligibilityResponse from JSON data.
func (m *CoverageEligibilityResponse) UnmarshalJSON(data []byte) error {
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
		Status *FinancialResourceStatusCodes `json:"status,omitempty"`
		Purpose []*EligibilityResponsePurpose `json:"purpose,omitempty"`
		Patient *Reference `json:"patient,omitempty"`
		ServicedDate *FhirDate `json:"serviceddate,omitempty"`
		ServicedPeriod *Period `json:"servicedperiod,omitempty"`
		Created *FhirDateTime `json:"created,omitempty"`
		Requestor *Reference `json:"requestor,omitempty"`
		Request *Reference `json:"request,omitempty"`
		Outcome *RemittanceOutcome `json:"outcome,omitempty"`
		Disposition *FhirString `json:"disposition,omitempty"`
		Insurer *Reference `json:"insurer,omitempty"`
		Insurance []*CoverageEligibilityResponseInsurance `json:"insurance,omitempty"`
		PreAuthRef *FhirString `json:"preauthref,omitempty"`
		Form *CodeableConcept `json:"form,omitempty"`
		Error []*CoverageEligibilityResponseError `json:"error,omitempty"`
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
	m.Status = temp.Status
	m.Purpose = temp.Purpose
	m.Patient = temp.Patient
	m.ServicedDate = temp.ServicedDate
	m.ServicedPeriod = temp.ServicedPeriod
	m.Created = temp.Created
	m.Requestor = temp.Requestor
	m.Request = temp.Request
	m.Outcome = temp.Outcome
	m.Disposition = temp.Disposition
	m.Insurer = temp.Insurer
	m.Insurance = temp.Insurance
	m.PreAuthRef = temp.PreAuthRef
	m.Form = temp.Form
	m.Error = temp.Error
	return nil
}

// MarshalJSON converts CoverageEligibilityResponse to JSON data.
func (m *CoverageEligibilityResponse) MarshalJSON() ([]byte, error) {
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
		Status *FinancialResourceStatusCodes `json:"status,omitempty"`
		Purpose []*EligibilityResponsePurpose `json:"purpose,omitempty"`
		Patient *Reference `json:"patient,omitempty"`
		ServicedDate interface{} `json:"serviceddate,omitempty"`
		ServicedDateElement map[string]interface{} `json:"_serviceddate,omitempty"`
		ServicedPeriod *Period `json:"servicedperiod,omitempty"`
		Created interface{} `json:"created,omitempty"`
		CreatedElement map[string]interface{} `json:"_created,omitempty"`
		Requestor *Reference `json:"requestor,omitempty"`
		Request *Reference `json:"request,omitempty"`
		Outcome *RemittanceOutcome `json:"outcome,omitempty"`
		Disposition interface{} `json:"disposition,omitempty"`
		DispositionElement map[string]interface{} `json:"_disposition,omitempty"`
		Insurer *Reference `json:"insurer,omitempty"`
		Insurance []*CoverageEligibilityResponseInsurance `json:"insurance,omitempty"`
		PreAuthRef interface{} `json:"preauthref,omitempty"`
		PreAuthRefElement map[string]interface{} `json:"_preauthref,omitempty"`
		Form *CodeableConcept `json:"form,omitempty"`
		Error []*CoverageEligibilityResponseError `json:"error,omitempty"`
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
	output.Status = m.Status
	output.Purpose = m.Purpose
	output.Patient = m.Patient
	if m.ServicedDate != nil && m.ServicedDate.Value != nil {
		output.ServicedDate = m.ServicedDate.Value
		if m.ServicedDate.Element != nil {
			output.ServicedDateElement = toMapOrNil(m.ServicedDate.Element.MarshalJSON())
		}
	}
	output.ServicedPeriod = m.ServicedPeriod
	if m.Created != nil && m.Created.Value != nil {
		output.Created = m.Created.Value
		if m.Created.Element != nil {
			output.CreatedElement = toMapOrNil(m.Created.Element.MarshalJSON())
		}
	}
	output.Requestor = m.Requestor
	output.Request = m.Request
	output.Outcome = m.Outcome
	if m.Disposition != nil && m.Disposition.Value != nil {
		output.Disposition = m.Disposition.Value
		if m.Disposition.Element != nil {
			output.DispositionElement = toMapOrNil(m.Disposition.Element.MarshalJSON())
		}
	}
	output.Insurer = m.Insurer
	output.Insurance = m.Insurance
	if m.PreAuthRef != nil && m.PreAuthRef.Value != nil {
		output.PreAuthRef = m.PreAuthRef.Value
		if m.PreAuthRef.Element != nil {
			output.PreAuthRefElement = toMapOrNil(m.PreAuthRef.Element.MarshalJSON())
		}
	}
	output.Form = m.Form
	output.Error = m.Error
	return json.Marshal(output)
}

// Clone creates a deep copy of CoverageEligibilityResponse.
func (m *CoverageEligibilityResponse) Clone() *CoverageEligibilityResponse {
	if m == nil { return nil }
	return &CoverageEligibilityResponse{
		Id: m.Id.Clone(),
		Meta: m.Meta.Clone(),
		ImplicitRules: m.ImplicitRules.Clone(),
		Language: m.Language.Clone(),
		Text: m.Text.Clone(),
		Contained: cloneSlices(m.Contained),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Identifier: cloneSlices(m.Identifier),
		Status: m.Status.Clone(),
		Purpose: cloneSlices(m.Purpose),
		Patient: m.Patient.Clone(),
		ServicedDate: m.ServicedDate.Clone(),
		ServicedPeriod: m.ServicedPeriod.Clone(),
		Created: m.Created.Clone(),
		Requestor: m.Requestor.Clone(),
		Request: m.Request.Clone(),
		Outcome: m.Outcome.Clone(),
		Disposition: m.Disposition.Clone(),
		Insurer: m.Insurer.Clone(),
		Insurance: cloneSlices(m.Insurance),
		PreAuthRef: m.PreAuthRef.Clone(),
		Form: m.Form.Clone(),
		Error: cloneSlices(m.Error),
	}
}

// Equals checks equality between two CoverageEligibilityResponse instances.
func (m *CoverageEligibilityResponse) Equals(other *CoverageEligibilityResponse) bool {
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
	if !m.Status.Equals(other.Status) { return false }
	if !compareSlices(m.Purpose, other.Purpose) { return false }
	if !m.Patient.Equals(other.Patient) { return false }
	if !m.ServicedDate.Equals(other.ServicedDate) { return false }
	if !m.ServicedPeriod.Equals(other.ServicedPeriod) { return false }
	if !m.Created.Equals(other.Created) { return false }
	if !m.Requestor.Equals(other.Requestor) { return false }
	if !m.Request.Equals(other.Request) { return false }
	if !m.Outcome.Equals(other.Outcome) { return false }
	if !m.Disposition.Equals(other.Disposition) { return false }
	if !m.Insurer.Equals(other.Insurer) { return false }
	if !compareSlices(m.Insurance, other.Insurance) { return false }
	if !m.PreAuthRef.Equals(other.PreAuthRef) { return false }
	if !m.Form.Equals(other.Form) { return false }
	if !compareSlices(m.Error, other.Error) { return false }
	return true
}

// CoverageEligibilityResponseInsurance
// Financial instruments for reimbursement for the health care products and services.
type CoverageEligibilityResponseInsurance struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Coverage *Reference `json:"coverage,omitempty"`
	Inforce *FhirBoolean `json:"inforce,omitempty"`
	BenefitPeriod *Period `json:"benefitperiod,omitempty"`
	Item []*CoverageEligibilityResponseItem `json:"item,omitempty"`
}

// NewCoverageEligibilityResponseInsurance creates a new CoverageEligibilityResponseInsurance instance.
func NewCoverageEligibilityResponseInsurance() *CoverageEligibilityResponseInsurance {
	return &CoverageEligibilityResponseInsurance{}
}

// UnmarshalJSON populates CoverageEligibilityResponseInsurance from JSON data.
func (m *CoverageEligibilityResponseInsurance) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Coverage *Reference `json:"coverage,omitempty"`
		Inforce *FhirBoolean `json:"inforce,omitempty"`
		BenefitPeriod *Period `json:"benefitperiod,omitempty"`
		Item []*CoverageEligibilityResponseItem `json:"item,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Coverage = temp.Coverage
	m.Inforce = temp.Inforce
	m.BenefitPeriod = temp.BenefitPeriod
	m.Item = temp.Item
	return nil
}

// MarshalJSON converts CoverageEligibilityResponseInsurance to JSON data.
func (m *CoverageEligibilityResponseInsurance) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Coverage *Reference `json:"coverage,omitempty"`
		Inforce interface{} `json:"inforce,omitempty"`
		InforceElement map[string]interface{} `json:"_inforce,omitempty"`
		BenefitPeriod *Period `json:"benefitperiod,omitempty"`
		Item []*CoverageEligibilityResponseItem `json:"item,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.Coverage = m.Coverage
	if m.Inforce != nil && m.Inforce.Value != nil {
		output.Inforce = m.Inforce.Value
		if m.Inforce.Element != nil {
			output.InforceElement = toMapOrNil(m.Inforce.Element.MarshalJSON())
		}
	}
	output.BenefitPeriod = m.BenefitPeriod
	output.Item = m.Item
	return json.Marshal(output)
}

// Clone creates a deep copy of CoverageEligibilityResponseInsurance.
func (m *CoverageEligibilityResponseInsurance) Clone() *CoverageEligibilityResponseInsurance {
	if m == nil { return nil }
	return &CoverageEligibilityResponseInsurance{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Coverage: m.Coverage.Clone(),
		Inforce: m.Inforce.Clone(),
		BenefitPeriod: m.BenefitPeriod.Clone(),
		Item: cloneSlices(m.Item),
	}
}

// Equals checks equality between two CoverageEligibilityResponseInsurance instances.
func (m *CoverageEligibilityResponseInsurance) Equals(other *CoverageEligibilityResponseInsurance) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Coverage.Equals(other.Coverage) { return false }
	if !m.Inforce.Equals(other.Inforce) { return false }
	if !m.BenefitPeriod.Equals(other.BenefitPeriod) { return false }
	if !compareSlices(m.Item, other.Item) { return false }
	return true
}

// CoverageEligibilityResponseItem
// Benefits and optionally current balances, and authorization details by category or service.
type CoverageEligibilityResponseItem struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Category *CodeableConcept `json:"category,omitempty"`
	ProductOrService *CodeableConcept `json:"productorservice,omitempty"`
	Modifier []*CodeableConcept `json:"modifier,omitempty"`
	Provider *Reference `json:"provider,omitempty"`
	Excluded *FhirBoolean `json:"excluded,omitempty"`
	Name *FhirString `json:"name,omitempty"`
	Description *FhirString `json:"description,omitempty"`
	Network *CodeableConcept `json:"network,omitempty"`
	Unit *CodeableConcept `json:"unit,omitempty"`
	Term *CodeableConcept `json:"term,omitempty"`
	Benefit []*CoverageEligibilityResponseBenefit `json:"benefit,omitempty"`
	AuthorizationRequired *FhirBoolean `json:"authorizationrequired,omitempty"`
	AuthorizationSupporting []*CodeableConcept `json:"authorizationsupporting,omitempty"`
	AuthorizationUrl *FhirUri `json:"authorizationurl,omitempty"`
}

// NewCoverageEligibilityResponseItem creates a new CoverageEligibilityResponseItem instance.
func NewCoverageEligibilityResponseItem() *CoverageEligibilityResponseItem {
	return &CoverageEligibilityResponseItem{}
}

// UnmarshalJSON populates CoverageEligibilityResponseItem from JSON data.
func (m *CoverageEligibilityResponseItem) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Category *CodeableConcept `json:"category,omitempty"`
		ProductOrService *CodeableConcept `json:"productorservice,omitempty"`
		Modifier []*CodeableConcept `json:"modifier,omitempty"`
		Provider *Reference `json:"provider,omitempty"`
		Excluded *FhirBoolean `json:"excluded,omitempty"`
		Name *FhirString `json:"name,omitempty"`
		Description *FhirString `json:"description,omitempty"`
		Network *CodeableConcept `json:"network,omitempty"`
		Unit *CodeableConcept `json:"unit,omitempty"`
		Term *CodeableConcept `json:"term,omitempty"`
		Benefit []*CoverageEligibilityResponseBenefit `json:"benefit,omitempty"`
		AuthorizationRequired *FhirBoolean `json:"authorizationrequired,omitempty"`
		AuthorizationSupporting []*CodeableConcept `json:"authorizationsupporting,omitempty"`
		AuthorizationUrl *FhirUri `json:"authorizationurl,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Category = temp.Category
	m.ProductOrService = temp.ProductOrService
	m.Modifier = temp.Modifier
	m.Provider = temp.Provider
	m.Excluded = temp.Excluded
	m.Name = temp.Name
	m.Description = temp.Description
	m.Network = temp.Network
	m.Unit = temp.Unit
	m.Term = temp.Term
	m.Benefit = temp.Benefit
	m.AuthorizationRequired = temp.AuthorizationRequired
	m.AuthorizationSupporting = temp.AuthorizationSupporting
	m.AuthorizationUrl = temp.AuthorizationUrl
	return nil
}

// MarshalJSON converts CoverageEligibilityResponseItem to JSON data.
func (m *CoverageEligibilityResponseItem) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Category *CodeableConcept `json:"category,omitempty"`
		ProductOrService *CodeableConcept `json:"productorservice,omitempty"`
		Modifier []*CodeableConcept `json:"modifier,omitempty"`
		Provider *Reference `json:"provider,omitempty"`
		Excluded interface{} `json:"excluded,omitempty"`
		ExcludedElement map[string]interface{} `json:"_excluded,omitempty"`
		Name interface{} `json:"name,omitempty"`
		NameElement map[string]interface{} `json:"_name,omitempty"`
		Description interface{} `json:"description,omitempty"`
		DescriptionElement map[string]interface{} `json:"_description,omitempty"`
		Network *CodeableConcept `json:"network,omitempty"`
		Unit *CodeableConcept `json:"unit,omitempty"`
		Term *CodeableConcept `json:"term,omitempty"`
		Benefit []*CoverageEligibilityResponseBenefit `json:"benefit,omitempty"`
		AuthorizationRequired interface{} `json:"authorizationrequired,omitempty"`
		AuthorizationRequiredElement map[string]interface{} `json:"_authorizationrequired,omitempty"`
		AuthorizationSupporting []*CodeableConcept `json:"authorizationsupporting,omitempty"`
		AuthorizationUrl interface{} `json:"authorizationurl,omitempty"`
		AuthorizationUrlElement map[string]interface{} `json:"_authorizationurl,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.Category = m.Category
	output.ProductOrService = m.ProductOrService
	output.Modifier = m.Modifier
	output.Provider = m.Provider
	if m.Excluded != nil && m.Excluded.Value != nil {
		output.Excluded = m.Excluded.Value
		if m.Excluded.Element != nil {
			output.ExcludedElement = toMapOrNil(m.Excluded.Element.MarshalJSON())
		}
	}
	if m.Name != nil && m.Name.Value != nil {
		output.Name = m.Name.Value
		if m.Name.Element != nil {
			output.NameElement = toMapOrNil(m.Name.Element.MarshalJSON())
		}
	}
	if m.Description != nil && m.Description.Value != nil {
		output.Description = m.Description.Value
		if m.Description.Element != nil {
			output.DescriptionElement = toMapOrNil(m.Description.Element.MarshalJSON())
		}
	}
	output.Network = m.Network
	output.Unit = m.Unit
	output.Term = m.Term
	output.Benefit = m.Benefit
	if m.AuthorizationRequired != nil && m.AuthorizationRequired.Value != nil {
		output.AuthorizationRequired = m.AuthorizationRequired.Value
		if m.AuthorizationRequired.Element != nil {
			output.AuthorizationRequiredElement = toMapOrNil(m.AuthorizationRequired.Element.MarshalJSON())
		}
	}
	output.AuthorizationSupporting = m.AuthorizationSupporting
	if m.AuthorizationUrl != nil && m.AuthorizationUrl.Value != nil {
		output.AuthorizationUrl = m.AuthorizationUrl.Value
		if m.AuthorizationUrl.Element != nil {
			output.AuthorizationUrlElement = toMapOrNil(m.AuthorizationUrl.Element.MarshalJSON())
		}
	}
	return json.Marshal(output)
}

// Clone creates a deep copy of CoverageEligibilityResponseItem.
func (m *CoverageEligibilityResponseItem) Clone() *CoverageEligibilityResponseItem {
	if m == nil { return nil }
	return &CoverageEligibilityResponseItem{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Category: m.Category.Clone(),
		ProductOrService: m.ProductOrService.Clone(),
		Modifier: cloneSlices(m.Modifier),
		Provider: m.Provider.Clone(),
		Excluded: m.Excluded.Clone(),
		Name: m.Name.Clone(),
		Description: m.Description.Clone(),
		Network: m.Network.Clone(),
		Unit: m.Unit.Clone(),
		Term: m.Term.Clone(),
		Benefit: cloneSlices(m.Benefit),
		AuthorizationRequired: m.AuthorizationRequired.Clone(),
		AuthorizationSupporting: cloneSlices(m.AuthorizationSupporting),
		AuthorizationUrl: m.AuthorizationUrl.Clone(),
	}
}

// Equals checks equality between two CoverageEligibilityResponseItem instances.
func (m *CoverageEligibilityResponseItem) Equals(other *CoverageEligibilityResponseItem) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Category.Equals(other.Category) { return false }
	if !m.ProductOrService.Equals(other.ProductOrService) { return false }
	if !compareSlices(m.Modifier, other.Modifier) { return false }
	if !m.Provider.Equals(other.Provider) { return false }
	if !m.Excluded.Equals(other.Excluded) { return false }
	if !m.Name.Equals(other.Name) { return false }
	if !m.Description.Equals(other.Description) { return false }
	if !m.Network.Equals(other.Network) { return false }
	if !m.Unit.Equals(other.Unit) { return false }
	if !m.Term.Equals(other.Term) { return false }
	if !compareSlices(m.Benefit, other.Benefit) { return false }
	if !m.AuthorizationRequired.Equals(other.AuthorizationRequired) { return false }
	if !compareSlices(m.AuthorizationSupporting, other.AuthorizationSupporting) { return false }
	if !m.AuthorizationUrl.Equals(other.AuthorizationUrl) { return false }
	return true
}

// CoverageEligibilityResponseBenefit
// Benefits used to date.
type CoverageEligibilityResponseBenefit struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Type *CodeableConcept `json:"type,omitempty"`
	AllowedUnsignedInt *FhirUnsignedInt `json:"allowedunsignedint,omitempty"`
	AllowedString *FhirString `json:"allowedstring,omitempty"`
	AllowedMoney *Money `json:"allowedmoney,omitempty"`
	UsedUnsignedInt *FhirUnsignedInt `json:"usedunsignedint,omitempty"`
	UsedString *FhirString `json:"usedstring,omitempty"`
	UsedMoney *Money `json:"usedmoney,omitempty"`
}

// NewCoverageEligibilityResponseBenefit creates a new CoverageEligibilityResponseBenefit instance.
func NewCoverageEligibilityResponseBenefit() *CoverageEligibilityResponseBenefit {
	return &CoverageEligibilityResponseBenefit{}
}

// UnmarshalJSON populates CoverageEligibilityResponseBenefit from JSON data.
func (m *CoverageEligibilityResponseBenefit) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Type *CodeableConcept `json:"type,omitempty"`
		AllowedUnsignedInt *FhirUnsignedInt `json:"allowedunsignedint,omitempty"`
		AllowedString *FhirString `json:"allowedstring,omitempty"`
		AllowedMoney *Money `json:"allowedmoney,omitempty"`
		UsedUnsignedInt *FhirUnsignedInt `json:"usedunsignedint,omitempty"`
		UsedString *FhirString `json:"usedstring,omitempty"`
		UsedMoney *Money `json:"usedmoney,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Type = temp.Type
	m.AllowedUnsignedInt = temp.AllowedUnsignedInt
	m.AllowedString = temp.AllowedString
	m.AllowedMoney = temp.AllowedMoney
	m.UsedUnsignedInt = temp.UsedUnsignedInt
	m.UsedString = temp.UsedString
	m.UsedMoney = temp.UsedMoney
	return nil
}

// MarshalJSON converts CoverageEligibilityResponseBenefit to JSON data.
func (m *CoverageEligibilityResponseBenefit) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Type *CodeableConcept `json:"type,omitempty"`
		AllowedUnsignedInt interface{} `json:"allowedunsignedint,omitempty"`
		AllowedUnsignedIntElement map[string]interface{} `json:"_allowedunsignedint,omitempty"`
		AllowedString interface{} `json:"allowedstring,omitempty"`
		AllowedStringElement map[string]interface{} `json:"_allowedstring,omitempty"`
		AllowedMoney *Money `json:"allowedmoney,omitempty"`
		UsedUnsignedInt interface{} `json:"usedunsignedint,omitempty"`
		UsedUnsignedIntElement map[string]interface{} `json:"_usedunsignedint,omitempty"`
		UsedString interface{} `json:"usedstring,omitempty"`
		UsedStringElement map[string]interface{} `json:"_usedstring,omitempty"`
		UsedMoney *Money `json:"usedmoney,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.Type = m.Type
	if m.AllowedUnsignedInt != nil && m.AllowedUnsignedInt.Value != nil {
		output.AllowedUnsignedInt = m.AllowedUnsignedInt.Value
		if m.AllowedUnsignedInt.Element != nil {
			output.AllowedUnsignedIntElement = toMapOrNil(m.AllowedUnsignedInt.Element.MarshalJSON())
		}
	}
	if m.AllowedString != nil && m.AllowedString.Value != nil {
		output.AllowedString = m.AllowedString.Value
		if m.AllowedString.Element != nil {
			output.AllowedStringElement = toMapOrNil(m.AllowedString.Element.MarshalJSON())
		}
	}
	output.AllowedMoney = m.AllowedMoney
	if m.UsedUnsignedInt != nil && m.UsedUnsignedInt.Value != nil {
		output.UsedUnsignedInt = m.UsedUnsignedInt.Value
		if m.UsedUnsignedInt.Element != nil {
			output.UsedUnsignedIntElement = toMapOrNil(m.UsedUnsignedInt.Element.MarshalJSON())
		}
	}
	if m.UsedString != nil && m.UsedString.Value != nil {
		output.UsedString = m.UsedString.Value
		if m.UsedString.Element != nil {
			output.UsedStringElement = toMapOrNil(m.UsedString.Element.MarshalJSON())
		}
	}
	output.UsedMoney = m.UsedMoney
	return json.Marshal(output)
}

// Clone creates a deep copy of CoverageEligibilityResponseBenefit.
func (m *CoverageEligibilityResponseBenefit) Clone() *CoverageEligibilityResponseBenefit {
	if m == nil { return nil }
	return &CoverageEligibilityResponseBenefit{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Type: m.Type.Clone(),
		AllowedUnsignedInt: m.AllowedUnsignedInt.Clone(),
		AllowedString: m.AllowedString.Clone(),
		AllowedMoney: m.AllowedMoney.Clone(),
		UsedUnsignedInt: m.UsedUnsignedInt.Clone(),
		UsedString: m.UsedString.Clone(),
		UsedMoney: m.UsedMoney.Clone(),
	}
}

// Equals checks equality between two CoverageEligibilityResponseBenefit instances.
func (m *CoverageEligibilityResponseBenefit) Equals(other *CoverageEligibilityResponseBenefit) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Type.Equals(other.Type) { return false }
	if !m.AllowedUnsignedInt.Equals(other.AllowedUnsignedInt) { return false }
	if !m.AllowedString.Equals(other.AllowedString) { return false }
	if !m.AllowedMoney.Equals(other.AllowedMoney) { return false }
	if !m.UsedUnsignedInt.Equals(other.UsedUnsignedInt) { return false }
	if !m.UsedString.Equals(other.UsedString) { return false }
	if !m.UsedMoney.Equals(other.UsedMoney) { return false }
	return true
}

// CoverageEligibilityResponseError
// Errors encountered during the processing of the request.
type CoverageEligibilityResponseError struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
}

// NewCoverageEligibilityResponseError creates a new CoverageEligibilityResponseError instance.
func NewCoverageEligibilityResponseError() *CoverageEligibilityResponseError {
	return &CoverageEligibilityResponseError{}
}

// UnmarshalJSON populates CoverageEligibilityResponseError from JSON data.
func (m *CoverageEligibilityResponseError) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Code *CodeableConcept `json:"code,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Code = temp.Code
	return nil
}

// MarshalJSON converts CoverageEligibilityResponseError to JSON data.
func (m *CoverageEligibilityResponseError) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Code *CodeableConcept `json:"code,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.Code = m.Code
	return json.Marshal(output)
}

// Clone creates a deep copy of CoverageEligibilityResponseError.
func (m *CoverageEligibilityResponseError) Clone() *CoverageEligibilityResponseError {
	if m == nil { return nil }
	return &CoverageEligibilityResponseError{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Code: m.Code.Clone(),
	}
}

// Equals checks equality between two CoverageEligibilityResponseError instances.
func (m *CoverageEligibilityResponseError) Equals(other *CoverageEligibilityResponseError) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Code.Equals(other.Code) { return false }
	return true
}

