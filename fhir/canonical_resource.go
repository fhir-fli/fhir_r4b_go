package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"fmt"
)

// CanonicalResource represents the base definition for all FHIR canonical resources.
type CanonicalResource struct {
	DomainResource
	URL          *FhirUri           `json:"url,omitempty"`
	Version      *FhirString        `json:"version,omitempty"`
	Status       *PublicationStatus `json:"status,omitempty"`
	Experimental *FhirBoolean       `json:"experimental,omitempty"`
	Date         *FhirDateTime      `json:"date,omitempty"`
	Publisher    *FhirString        `json:"publisher,omitempty"`
	Contact      []ContactDetail    `json:"contact,omitempty"`
	Description  *FhirMarkdown      `json:"description,omitempty"`
	UseContext   []UsageContext     `json:"useContext,omitempty"`
	Jurisdiction []CodeableConcept  `json:"jurisdiction,omitempty"`
}

// NewCanonicalResource creates a new CanonicalResource instance.
func NewCanonicalResource(
	resourceType ResourceType,
	id *FhirString,
	meta *FhirMeta,
	implicitRules *FhirUri,
	language *CommonLanguages,
	text *Narrative,
	contained []Resource,
	extension_ []FhirExtension,
	modifierExtension []FhirExtension,
	url *FhirUri,
	version *FhirString,
	status *PublicationStatus,
	experimental *FhirBoolean,
	date *FhirDateTime,
	publisher *FhirString,
	contact []ContactDetail,
	description *FhirMarkdown,
	useContext []UsageContext,
	jurisdiction []CodeableConcept,
) *CanonicalResource {
	return &CanonicalResource{
		DomainResource: *NewDomainResource(
			resourceType,
			id,
			meta,
			implicitRules,
			language,
			text,
			contained,
			extension_,
			modifierExtension,
		),
		URL:          url,
		Version:      version,
		Status:       status,
		Experimental: experimental,
		Date:         date,
		Publisher:    publisher,
		Contact:      contact,
		Description:  description,
		UseContext:   useContext,
		Jurisdiction: jurisdiction,
	}
}

// ToJSON converts the CanonicalResource to JSON data.
func (c *CanonicalResource) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}

// FromJSON populates a CanonicalResource from JSON data.
func (c *CanonicalResource) FromJSON(data []byte) error {
	return json.Unmarshal(data, c)
}

// CopyWith creates a copy of CanonicalResource with optional updates.
func (c *CanonicalResource) CopyWith(
	id *FhirString,
	meta *FhirMeta,
	implicitRules *FhirUri,
	language *CommonLanguages,
	text *Narrative,
	contained []Resource,
	extension_ []FhirExtension,
	modifierExtension []FhirExtension,
	url *FhirUri,
	version *FhirString,
	status *PublicationStatus,
	experimental *FhirBoolean,
	date *FhirDateTime,
	publisher *FhirString,
	contact []ContactDetail,
	description *FhirMarkdown,
	useContext []UsageContext,
	jurisdiction []CodeableConcept,
) *CanonicalResource {
	return &CanonicalResource{
		DomainResource: *c.DomainResource.CopyWith(
			id,
			meta,
			implicitRules,
			language,
			text,
			contained,
			extension_,
			modifierExtension,
		),
		URL:          ifNotNil(url, c.URL),
		Version:      ifNotNil(version, c.Version),
		Status:       ifNotNil(status, c.Status),
		Experimental: ifNotNil(experimental, c.Experimental),
		Date:         ifNotNil(date, c.Date),
		Publisher:    ifNotNil(publisher, c.Publisher),
		Contact:      copyContactDetails(ifNotNilSlice(contact, c.Contact)),
		Description:  ifNotNil(description, c.Description),
		UseContext:   copyUsageContexts(ifNotNilSlice(useContext, c.UseContext)),
		Jurisdiction: copyCodeableConcepts(ifNotNilSlice(jurisdiction, c.Jurisdiction)),
	}
}

// FromYAML populates a CanonicalResource from YAML input.
func CanonicalResourceFromYAML(yamlData interface{}) (*CanonicalResource, error) {
	jsonData, err := convertYAMLToJSON(yamlData)
	if err != nil {
		return nil, err
	}
	var res CanonicalResource
	if err := json.Unmarshal(jsonData, &res); err != nil {
		return nil, errors.New("failed to parse CanonicalResource from YAML")
	}
	return &res, nil
}

// Path returns the local reference for this CanonicalResource in the form "ResourceType/Id".
func (c *CanonicalResource) Path() string {
	if c.ID != nil {
		return fmt.Sprintf("%s/%s", c.ResourceType, *c.ID)
	}
	return fmt.Sprintf("%s/", c.ResourceType)
}

// IsEmpty checks if the CanonicalResource is empty.
func (c *CanonicalResource) IsEmpty() bool {
	return c.DomainResource.IsEmpty() &&
		c.URL == nil &&
		c.Version == nil &&
		c.Status == nil &&
		c.Experimental == nil &&
		c.Date == nil &&
		c.Publisher == nil &&
		len(c.Contact) == 0 &&
		c.Description == nil &&
		len(c.UseContext) == 0 &&
		len(c.Jurisdiction) == 0
}

// EqualsDeep performs a deep equality check with another CanonicalResource.
func (c *CanonicalResource) EqualsDeep(other *CanonicalResource) bool {
	if other == nil {
		return false
	}

	if !c.DomainResource.EqualsDeep(&other.DomainResource) {
		return false
	}

	return comparePointer(c.URL, other.URL) &&
		comparePointer(c.Version, other.Version) &&
		comparePointer(c.Status, other.Status) &&
		comparePointer(c.Experimental, other.Experimental) &&
		comparePointer(c.Date, other.Date) &&
		comparePointer(c.Publisher, other.Publisher) &&
		compareContactDetails(c.Contact, other.Contact) &&
		comparePointer(c.Description, other.Description) &&
		compareUsageContexts(c.UseContext, other.UseContext) &&
		compareCodeableConcepts(c.Jurisdiction, other.Jurisdiction)
}

// Helpers:

// copyContactDetails creates a deep copy of a slice of ContactDetail.
func copyContactDetails(details []ContactDetail) []ContactDetail {
	if details == nil {
		return nil
	}
	copy := make([]ContactDetail, len(details))
	for i, detail := range details {
		copy[i] = *detail.Clone()
	}
	return copy
}

// compareContactDetails checks deep equality for slices of ContactDetail.
func compareContactDetails(slice1, slice2 []ContactDetail) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if !slice1[i].EqualsDeep(&slice2[i]) {
			return false
		}
	}
	return true
}

// copyUsageContexts creates a deep copy of a slice of UsageContext.
func copyUsageContexts(contexts []UsageContext) []UsageContext {
	if contexts == nil {
		return nil
	}
	copy := make([]UsageContext, len(contexts))
	for i, context := range contexts {
		copy[i] = *context.Clone()
	}
	return copy
}

// compareUsageContexts checks deep equality for slices of UsageContext.
func compareUsageContexts(slice1, slice2 []UsageContext) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if !slice1[i].EqualsDeep(&slice2[i]) {
			return false
		}
	}
	return true
}

// copyCodeableConcepts creates a deep copy of a slice of CodeableConcept.
func copyCodeableConcepts(concepts []CodeableConcept) []CodeableConcept {
	if concepts == nil {
		return nil
	}
	copy := make([]CodeableConcept, len(concepts))
	for i, concept := range concepts {
		copy[i] = *concept.Clone()
	}
	return copy
}

// compareCodeableConcepts checks deep equality for slices of CodeableConcept.
func compareCodeableConcepts(slice1, slice2 []CodeableConcept) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if !slice1[i].EqualsDeep(&slice2[i]) {
			return false
		}
	}
	return true
}
