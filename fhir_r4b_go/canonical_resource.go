package fhir_r4b_go

import (
	"encoding/json"
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
	Contact      []*ContactDetail   `json:"contact,omitempty"`
	Description  *FhirMarkdown      `json:"description,omitempty"`
	UseContext   []*UsageContext    `json:"useContext,omitempty"`
	Jurisdiction []*CodeableConcept `json:"jurisdiction,omitempty"`
}

// NewCanonicalResource creates a new CanonicalResource instance.
func NewCanonicalResource(
	resourceType ResourceType,
	id *FhirString,
	meta *FhirMeta,
	implicitRules *FhirUri,
	language *CommonLanguages,
	text *Narrative,
	contained []*Resource,
	extension_ []*FhirExtension,
	modifierExtension []*FhirExtension,
	url *FhirUri,
	version *FhirString,
	status *PublicationStatus,
	experimental *FhirBoolean,
	date *FhirDateTime,
	publisher *FhirString,
	contact []*ContactDetail,
	description *FhirMarkdown,
	useContext []*UsageContext,
	jurisdiction []*CodeableConcept,
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

// Clone creates a deep copy of the CanonicalResource.
func (c *CanonicalResource) Clone() *CanonicalResource {
	if c == nil {
		return nil
	}
	return &CanonicalResource{
		DomainResource: *c.DomainResource.Clone(),
		URL:            c.URL.Clone(),
		Version:        c.Version.Clone(),
		Status:         c.Status.Clone(),
		Experimental:   c.Experimental.Clone(),
		Date:           c.Date.Clone(),
		Publisher:      c.Publisher.Clone(),
		Contact:        cloneSlices(c.Contact),
		Description:    c.Description.Clone(),
		UseContext:     cloneSlices(c.UseContext),
		Jurisdiction:   cloneSlices(c.Jurisdiction),
	}
}

// Equals checks for equality with another CanonicalResource.
func (c *CanonicalResource) Equals(other Equalable) bool {
	otherCanonicalResource, ok := other.(*CanonicalResource)
	if !ok {
		return false
	}
	return c.DomainResource.Equals(&otherCanonicalResource.DomainResource) &&
		c.URL.Equals(otherCanonicalResource.URL) &&
		c.Version.Equals(otherCanonicalResource.Version) &&
		c.Status.Equals(otherCanonicalResource.Status) &&
		c.Experimental.Equals(otherCanonicalResource.Experimental) &&
		c.Date.Equals(otherCanonicalResource.Date) &&
		c.Publisher.Equals(otherCanonicalResource.Publisher) &&
		compareSlices(c.Contact, otherCanonicalResource.Contact) &&
		c.Description.Equals(otherCanonicalResource.Description) &&
		compareSlices(c.UseContext, otherCanonicalResource.UseContext) &&
		compareSlices(c.Jurisdiction, otherCanonicalResource.Jurisdiction)
}

// MarshalJSON converts the CanonicalResource to JSON.
func (c *CanonicalResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(c)
}

// UnmarshalJSON populates the CanonicalResource from JSON data.
func (c *CanonicalResource) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, c)
}
