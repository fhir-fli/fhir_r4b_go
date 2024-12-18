package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"fmt"
)

// DomainResource represents the base definition for FHIR domain resources.
type DomainResource struct {
	Resource
	Text              *Narrative      `json:"text,omitempty"`
	Contained         []Resource      `json:"contained,omitempty"`
	Extension_        []FhirExtension `json:"extension,omitempty"`
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
		Text:              text,
		Contained:         contained,
		Extension_:        extension_,
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
		Text:              ifNotNil(text, d.Text),
		Contained:         copyResourceSlice(ifNotNilSlice(contained, d.Contained)),
		Extension_:        copyExtensionSlice(ifNotNilSlice(extension_, d.Extension_)),
		ModifierExtension: copyExtensionSlice(ifNotNilSlice(modifierExtension, d.ModifierExtension)),
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

// IsEmpty checks if the DomainResource is empty.
func (d *DomainResource) IsEmpty() bool {
	return d.Resource.IsEmpty() &&
		d.Text == nil &&
		len(d.Contained) == 0 &&
		len(d.Extension_) == 0 &&
		len(d.ModifierExtension) == 0
}

// EqualsDeep performs a deep equality check with another DomainResource.
func (d *DomainResource) EqualsDeep(other *DomainResource) bool {
	if other == nil {
		return false
	}

	// Compare the embedded Resource field
	if !compareResources(&d.Resource, &other.Resource) {
		return false
	}

	// Compare all other fields
	return comparePointer(d.Text, other.Text) &&
		compareResourceSlices(d.Contained, other.Contained) &&
		compareExtensionSlices(d.Extension_, other.Extension_) &&
		compareExtensionSlices(d.ModifierExtension, other.ModifierExtension)
}

// compareResources compares two Resource instances deeply.
func compareResources(r1, r2 *Resource) bool {
	if r1 == nil || r2 == nil {
		return r1 == r2
	}

	return r1.ResourceType == r2.ResourceType &&
		comparePointer(r1.ID, r2.ID) &&
		comparePointer(r1.Meta, r2.Meta) &&
		comparePointer(r1.ImplicitRules, r2.ImplicitRules) &&
		comparePointer(r1.Language, r2.Language)
}

// Helpers:

// copyResourceSlice creates a deep copy of a slice of Resource.
func copyResourceSlice(resources []Resource) []Resource {
	if resources == nil {
		return nil
	}
	copy := make([]Resource, len(resources))
	for i, res := range resources {
		copy[i] = *resourceDeepCopy(&res)
	}
	return copy
}

// resourceDeepCopy creates a deep copy of a Resource.
func resourceDeepCopy(r *Resource) *Resource {
	if r == nil {
		return nil
	}

	return &Resource{
		ResourceType:  r.ResourceType,
		ID:            deepCopyFhirString(r.ID),
		Meta:          metaDeepCopy(r.Meta),
		ImplicitRules: deepCopyFhirUri(r.ImplicitRules),
		Language:      r.Language, // Assuming it's immutable
	}
}

// copyExtensionSlice creates a deep copy of a slice of FhirExtension.
func copyExtensionSlice(extensions []FhirExtension) []FhirExtension {
	if extensions == nil {
		return nil
	}
	copy := make([]FhirExtension, len(extensions))
	for i, ext := range extensions {
		copy[i] = *ext.Clone()
	}
	return copy
}

// compareResourceSlices compares slices of Resources deeply.
func compareResourceSlices(slice1, slice2 []Resource) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if !compareResources(&slice1[i], &slice2[i]) {
			return false
		}
	}
	return true
}

// compareExtensionSlices compares slices of FhirExtensions deeply.
func compareExtensionSlices(slice1, slice2 []FhirExtension) bool {
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

// comparePointer compares two pointers deeply.
func comparePointer[T comparable](p1, p2 *T) bool {
	if p1 == nil || p2 == nil {
		return p1 == p2
	}
	return *p1 == *p2
}

// ifNotNilSlice returns the new slice if it's not nil; otherwise, returns the old slice.
func ifNotNilSlice[T any](newSlice, oldSlice []T) []T {
	if newSlice != nil {
		return newSlice
	}
	return oldSlice
}
