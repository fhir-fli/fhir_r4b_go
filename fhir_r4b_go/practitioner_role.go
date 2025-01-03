// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"
)

// PractitionerRole
// A specific set of Roles/Locations/specialties/services that a practitioner may perform at an organization for a period of time.
type PractitionerRole struct {
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
	Period *Period `json:"period,omitempty"`
	Practitioner *Reference `json:"practitioner,omitempty"`
	Organization *Reference `json:"organization,omitempty"`
	Code []*CodeableConcept `json:"code,omitempty"`
	Specialty []*CodeableConcept `json:"specialty,omitempty"`
	Location []*Reference `json:"location,omitempty"`
	HealthcareService []*Reference `json:"healthcareservice,omitempty"`
	Telecom []*ContactPoint `json:"telecom,omitempty"`
	AvailableTime []*PractitionerRoleAvailableTime `json:"availabletime,omitempty"`
	NotAvailable []*PractitionerRoleNotAvailable `json:"notavailable,omitempty"`
	AvailabilityExceptions *FhirString `json:"availabilityexceptions,omitempty"`
	Endpoint []*Reference `json:"endpoint,omitempty"`
}

// NewPractitionerRole creates a new PractitionerRole instance.
func NewPractitionerRole() *PractitionerRole {
	return &PractitionerRole{}
}

// UnmarshalJSON populates PractitionerRole from JSON data.
func (m *PractitionerRole) UnmarshalJSON(data []byte) error {
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
		Period *Period `json:"period,omitempty"`
		Practitioner *Reference `json:"practitioner,omitempty"`
		Organization *Reference `json:"organization,omitempty"`
		Code []*CodeableConcept `json:"code,omitempty"`
		Specialty []*CodeableConcept `json:"specialty,omitempty"`
		Location []*Reference `json:"location,omitempty"`
		HealthcareService []*Reference `json:"healthcareservice,omitempty"`
		Telecom []*ContactPoint `json:"telecom,omitempty"`
		AvailableTime []*PractitionerRoleAvailableTime `json:"availabletime,omitempty"`
		NotAvailable []*PractitionerRoleNotAvailable `json:"notavailable,omitempty"`
		AvailabilityExceptions *FhirString `json:"availabilityexceptions,omitempty"`
		Endpoint []*Reference `json:"endpoint,omitempty"`
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
	m.Period = temp.Period
	m.Practitioner = temp.Practitioner
	m.Organization = temp.Organization
	m.Code = temp.Code
	m.Specialty = temp.Specialty
	m.Location = temp.Location
	m.HealthcareService = temp.HealthcareService
	m.Telecom = temp.Telecom
	m.AvailableTime = temp.AvailableTime
	m.NotAvailable = temp.NotAvailable
	m.AvailabilityExceptions = temp.AvailabilityExceptions
	m.Endpoint = temp.Endpoint
	return nil
}

// MarshalJSON converts PractitionerRole to JSON data.
func (m *PractitionerRole) MarshalJSON() ([]byte, error) {
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
		Period *Period `json:"period,omitempty"`
		Practitioner *Reference `json:"practitioner,omitempty"`
		Organization *Reference `json:"organization,omitempty"`
		Code []*CodeableConcept `json:"code,omitempty"`
		Specialty []*CodeableConcept `json:"specialty,omitempty"`
		Location []*Reference `json:"location,omitempty"`
		HealthcareService []*Reference `json:"healthcareservice,omitempty"`
		Telecom []*ContactPoint `json:"telecom,omitempty"`
		AvailableTime []*PractitionerRoleAvailableTime `json:"availabletime,omitempty"`
		NotAvailable []*PractitionerRoleNotAvailable `json:"notavailable,omitempty"`
		AvailabilityExceptions interface{} `json:"availabilityexceptions,omitempty"`
		AvailabilityExceptionsElement map[string]interface{} `json:"_availabilityexceptions,omitempty"`
		Endpoint []*Reference `json:"endpoint,omitempty"`
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
	output.Period = m.Period
	output.Practitioner = m.Practitioner
	output.Organization = m.Organization
	output.Code = m.Code
	output.Specialty = m.Specialty
	output.Location = m.Location
	output.HealthcareService = m.HealthcareService
	output.Telecom = m.Telecom
	output.AvailableTime = m.AvailableTime
	output.NotAvailable = m.NotAvailable
	if m.AvailabilityExceptions != nil && m.AvailabilityExceptions.Value != nil {
		output.AvailabilityExceptions = m.AvailabilityExceptions.Value
		if m.AvailabilityExceptions.Element != nil {
			output.AvailabilityExceptionsElement = toMapOrNil(m.AvailabilityExceptions.Element.MarshalJSON())
		}
	}
	output.Endpoint = m.Endpoint
	return json.Marshal(output)
}

// Clone creates a deep copy of PractitionerRole.
func (m *PractitionerRole) Clone() *PractitionerRole {
	if m == nil { return nil }
	return &PractitionerRole{
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
		Period: m.Period.Clone(),
		Practitioner: m.Practitioner.Clone(),
		Organization: m.Organization.Clone(),
		Code: cloneSlices(m.Code),
		Specialty: cloneSlices(m.Specialty),
		Location: cloneSlices(m.Location),
		HealthcareService: cloneSlices(m.HealthcareService),
		Telecom: cloneSlices(m.Telecom),
		AvailableTime: cloneSlices(m.AvailableTime),
		NotAvailable: cloneSlices(m.NotAvailable),
		AvailabilityExceptions: m.AvailabilityExceptions.Clone(),
		Endpoint: cloneSlices(m.Endpoint),
	}
}

// Equals checks equality between two PractitionerRole instances.
func (m *PractitionerRole) Equals(other *PractitionerRole) bool {
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
	if !m.Period.Equals(other.Period) { return false }
	if !m.Practitioner.Equals(other.Practitioner) { return false }
	if !m.Organization.Equals(other.Organization) { return false }
	if !compareSlices(m.Code, other.Code) { return false }
	if !compareSlices(m.Specialty, other.Specialty) { return false }
	if !compareSlices(m.Location, other.Location) { return false }
	if !compareSlices(m.HealthcareService, other.HealthcareService) { return false }
	if !compareSlices(m.Telecom, other.Telecom) { return false }
	if !compareSlices(m.AvailableTime, other.AvailableTime) { return false }
	if !compareSlices(m.NotAvailable, other.NotAvailable) { return false }
	if !m.AvailabilityExceptions.Equals(other.AvailabilityExceptions) { return false }
	if !compareSlices(m.Endpoint, other.Endpoint) { return false }
	return true
}

// PractitionerRoleAvailableTime
// A collection of times the practitioner is available or performing this role at the location and/or healthcareservice.
type PractitionerRoleAvailableTime struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	DaysOfWeek []*DaysOfWeek `json:"daysofweek,omitempty"`
	AllDay *FhirBoolean `json:"allday,omitempty"`
	AvailableStartTime *FhirTime `json:"availablestarttime,omitempty"`
	AvailableEndTime *FhirTime `json:"availableendtime,omitempty"`
}

// NewPractitionerRoleAvailableTime creates a new PractitionerRoleAvailableTime instance.
func NewPractitionerRoleAvailableTime() *PractitionerRoleAvailableTime {
	return &PractitionerRoleAvailableTime{}
}

// UnmarshalJSON populates PractitionerRoleAvailableTime from JSON data.
func (m *PractitionerRoleAvailableTime) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		DaysOfWeek []*DaysOfWeek `json:"daysofweek,omitempty"`
		AllDay *FhirBoolean `json:"allday,omitempty"`
		AvailableStartTime *FhirTime `json:"availablestarttime,omitempty"`
		AvailableEndTime *FhirTime `json:"availableendtime,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.DaysOfWeek = temp.DaysOfWeek
	m.AllDay = temp.AllDay
	m.AvailableStartTime = temp.AvailableStartTime
	m.AvailableEndTime = temp.AvailableEndTime
	return nil
}

// MarshalJSON converts PractitionerRoleAvailableTime to JSON data.
func (m *PractitionerRoleAvailableTime) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		DaysOfWeek []*DaysOfWeek `json:"daysofweek,omitempty"`
		AllDay interface{} `json:"allday,omitempty"`
		AllDayElement map[string]interface{} `json:"_allday,omitempty"`
		AvailableStartTime interface{} `json:"availablestarttime,omitempty"`
		AvailableStartTimeElement map[string]interface{} `json:"_availablestarttime,omitempty"`
		AvailableEndTime interface{} `json:"availableendtime,omitempty"`
		AvailableEndTimeElement map[string]interface{} `json:"_availableendtime,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.DaysOfWeek = m.DaysOfWeek
	if m.AllDay != nil && m.AllDay.Value != nil {
		output.AllDay = m.AllDay.Value
		if m.AllDay.Element != nil {
			output.AllDayElement = toMapOrNil(m.AllDay.Element.MarshalJSON())
		}
	}
	if m.AvailableStartTime != nil && m.AvailableStartTime.Value != nil {
		output.AvailableStartTime = m.AvailableStartTime.Value
		if m.AvailableStartTime.Element != nil {
			output.AvailableStartTimeElement = toMapOrNil(m.AvailableStartTime.Element.MarshalJSON())
		}
	}
	if m.AvailableEndTime != nil && m.AvailableEndTime.Value != nil {
		output.AvailableEndTime = m.AvailableEndTime.Value
		if m.AvailableEndTime.Element != nil {
			output.AvailableEndTimeElement = toMapOrNil(m.AvailableEndTime.Element.MarshalJSON())
		}
	}
	return json.Marshal(output)
}

// Clone creates a deep copy of PractitionerRoleAvailableTime.
func (m *PractitionerRoleAvailableTime) Clone() *PractitionerRoleAvailableTime {
	if m == nil { return nil }
	return &PractitionerRoleAvailableTime{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		DaysOfWeek: cloneSlices(m.DaysOfWeek),
		AllDay: m.AllDay.Clone(),
		AvailableStartTime: m.AvailableStartTime.Clone(),
		AvailableEndTime: m.AvailableEndTime.Clone(),
	}
}

// Equals checks equality between two PractitionerRoleAvailableTime instances.
func (m *PractitionerRoleAvailableTime) Equals(other *PractitionerRoleAvailableTime) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !compareSlices(m.DaysOfWeek, other.DaysOfWeek) { return false }
	if !m.AllDay.Equals(other.AllDay) { return false }
	if !m.AvailableStartTime.Equals(other.AvailableStartTime) { return false }
	if !m.AvailableEndTime.Equals(other.AvailableEndTime) { return false }
	return true
}

// PractitionerRoleNotAvailable
// The practitioner is not available or performing this role during this period of time due to the provided reason.
type PractitionerRoleNotAvailable struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Description *FhirString `json:"description,omitempty"`
	During *Period `json:"during,omitempty"`
}

// NewPractitionerRoleNotAvailable creates a new PractitionerRoleNotAvailable instance.
func NewPractitionerRoleNotAvailable() *PractitionerRoleNotAvailable {
	return &PractitionerRoleNotAvailable{}
}

// UnmarshalJSON populates PractitionerRoleNotAvailable from JSON data.
func (m *PractitionerRoleNotAvailable) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Description *FhirString `json:"description,omitempty"`
		During *Period `json:"during,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Description = temp.Description
	m.During = temp.During
	return nil
}

// MarshalJSON converts PractitionerRoleNotAvailable to JSON data.
func (m *PractitionerRoleNotAvailable) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Description interface{} `json:"description,omitempty"`
		DescriptionElement map[string]interface{} `json:"_description,omitempty"`
		During *Period `json:"during,omitempty"`
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
	output.During = m.During
	return json.Marshal(output)
}

// Clone creates a deep copy of PractitionerRoleNotAvailable.
func (m *PractitionerRoleNotAvailable) Clone() *PractitionerRoleNotAvailable {
	if m == nil { return nil }
	return &PractitionerRoleNotAvailable{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Description: m.Description.Clone(),
		During: m.During.Clone(),
	}
}

// Equals checks equality between two PractitionerRoleNotAvailable instances.
func (m *PractitionerRoleNotAvailable) Equals(other *PractitionerRoleNotAvailable) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Description.Equals(other.Description) { return false }
	if !m.During.Equals(other.During) { return false }
	return true
}

