package fhir_r4b_go

import (
	"encoding/json"
)

// DomainResource represents the base definition for FHIR domain resources.
type DomainResource struct {
	Resource
	Text              *Narrative      `json:"text,omitempty"`
	Contained         []*Resource      `json:"contained,omitempty"`
	Extension_        []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierExtension,omitempty"`
}

// NewDomainResource creates a new DomainResource instance.
func NewDomainResource(
	resourceType ResourceType,
	id *FhirString,
	meta *FhirMeta,
	implicitRules *FhirUri,
	language *CommonLanguages,
	text *Narrative,
	contained []*Resource,
	extension_ []*FhirExtension,
	modifierExtension []*FhirExtension,
) *DomainResource {
	return &DomainResource{
		Resource: Resource{
			ResourceType:  resourceType,
			ID:            id,
			Meta:          meta,
			ImplicitRules: implicitRules,
			Language:      language,
		},
		Text:              text,
		Contained:         contained,
		Extension_:        extension_,
		ModifierExtension: modifierExtension,
	}
}

// Equals compares DomainResource with another Equalable instance.
func (d *DomainResource) Equals(other Equalable) bool {
	otherDomainResource, ok := other.(*DomainResource)
	if !ok {
		return false
	}
	return d.Resource.Equals(&otherDomainResource.Resource) &&
		d.Text.Equals(otherDomainResource.Text) &&
		compareSlices(d.Contained, otherDomainResource.Contained) &&
		compareSlices(d.Extension_, otherDomainResource.Extension_) &&
		compareSlices(d.ModifierExtension, otherDomainResource.ModifierExtension)
}

// Clone creates a deep copy of the DomainResource.
func (d *DomainResource) Clone() *DomainResource {
	if d == nil {
		return nil
	}
	return &DomainResource{
		Resource:          *d.Resource.Clone(),
		Text:              d.Text.Clone(),
		Contained:         cloneSlices(d.Contained),
		Extension_:        cloneSlices(d.Extension_),
		ModifierExtension: cloneSlices(d.ModifierExtension),
	}
}

// ToJSON converts the DomainResource to JSON.
func (d *DomainResource) ToJSON() ([]byte, error) {
	return json.Marshal(d)
}
