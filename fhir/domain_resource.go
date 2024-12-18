package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"fmt"
)

// DomainResource represents the base definition for FHIR domain resources.
type DomainResource struct {
	Resource
	Text             *Narrative       `json:"text,omitempty"`
	Contained        []Resource       `json:"contained,omitempty"`
	Extension_       []FhirExtension  `json:"extension,omitempty"`
	ModifierExtension []FhirExtension `json:"modifierExtension,omitempty"`
}

// NewDomainResource creates a new DomainResource instance.
func NewDomainResource(
	resourceType ResourceType,
	id *FhirString,
	meta *FhirMeta,
	implicitRules *FhirUri,
	language *CommonLanguages,
	text *Narrative,
	contained []Resource,
	extension_ []FhirExtension,
	modifierExtension []FhirExtension,
) *DomainResource {
	return &DomainResource{
		Resource: Resource{
			ResourceType:  resourceType,
			ID:            id,
			Meta:          meta,
			ImplicitRules: implicitRules,
			Language:      language,
		},
		Text:             text,
		Contained:        contained,
		Extension_:       extension_,
		ModifierExtension: modifierExtension,
	}
}

// ToJSON converts the DomainResource to JSON data.
func (d *DomainResource) ToJSON() ([]byte, error) {
	return json.Marshal(d)
}

// FromJSON populates a DomainResource from JSON data.
func (d *DomainResource) FromJSON(data []byte) error {
	return json.Unmarshal(data, d)
}

// CopyWith creates a copy of the DomainResource with optional updates.
func (d *DomainResource) CopyWith(
	id *FhirString,
	meta *FhirMeta,
	implicitRules *FhirUri,
	language *CommonLanguages,
	text *Narrative,
	contained []Resource,
	extension_ []FhirExtension,
	modifierExtension []FhirExtension,
) *DomainResource {
	return &DomainResource{
		Resource: Resource{
			ResourceType:  d.ResourceType,
			ID:            ifNotNil(id, d.ID),
			Meta:          ifNotNil(meta, d.Meta),
			ImplicitRules: ifNotNil(implicitRules, d.ImplicitRules),
			Language:      ifNotNil(language, d.Language),
		},
		Text:             ifNotNil(text, d.Text),
		Contained:        ifNotNilSlice(contained, d.Contained),
		Extension_:       ifNotNilSlice(extension_, d.Extension_),
		ModifierExtension: ifNotNilSlice(modifierExtension, d.ModifierExtension),
	}
}

// FromYAML populates a DomainResource from YAML input.
func DomainResourceFromYAML(yamlData interface{}) (*DomainResource, error) {
	jsonData, err := convertYAMLToJSON(yamlData)
	if err != nil {
		return nil, err
	}
	var res DomainResource
	if err := json.Unmarshal(jsonData, &res); err != nil {
		return nil, errors.New("failed to parse DomainResource from YAML")
	}
	return &res, nil
}

// Path returns the local reference for this DomainResource in the form "ResourceType/Id".
func (d *DomainResource) Path() string {
	if d.ID != nil {
		return fmt.Sprintf("%s/%s", d.ResourceType, *d.ID)
	}
	return fmt.Sprintf("%s/", d.ResourceType)
}
