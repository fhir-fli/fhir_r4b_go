package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Resource represents the base definition for all FHIR resources.
type Resource struct {
	FhirBase
	ResourceType   ResourceType `json:"resourceType,omitempty"`
	ID             *FhirString  `json:"id,omitempty"`
	Meta           *FhirMeta    `json:"meta,omitempty"`
	ImplicitRules  *FhirUri     `json:"implicitRules,omitempty"`
	Language       *CommonLanguages `json:"language,omitempty"`
}

// NewResource creates a new Resource instance.
func NewResource(
	resourceType ResourceType,
	id *FhirString,
	meta *FhirMeta,
	implicitRules *FhirUri,
	language *CommonLanguages,
) *Resource {
	return &Resource{
		ResourceType:  resourceType,
		ID:            id,
		Meta:          meta,
		ImplicitRules: implicitRules,
		Language:      language,
	}
}

// FromJSON populates a Resource from JSON data.
func (r *Resource) FromJSON(data []byte) error {
	return json.Unmarshal(data, r)
}

// ToJSON converts a Resource to JSON data.
func (r *Resource) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

// Path returns the local reference for this Resource in the form "ResourceType/Id".
func (r *Resource) Path() string {
	if r.ID != nil {
		return fmt.Sprintf("%s/%s", r.ResourceType, *r.ID)
	}
	return fmt.Sprintf("%s/", r.ResourceType)
}

// ThisReference returns a Reference referring to the current Resource.
func (r *Resource) ThisReference() *Reference {
	return &Reference{
		Reference: &FhirString{Value: r.Path()},
		Type:      &FhirUri{Value: string(r.ResourceType)},
	}
}

// NewIDIfNoID returns the Resource with a new ID if there is no current ID.
func (r *Resource) NewIDIfNoID() *Resource {
	if r.ID == nil {
		return r.NewID()
	}
	return r
}

// NewID returns the Resource with a new generated ID.
func (r *Resource) NewID() *Resource {
	newID := GenerateNewID() // Function to generate new ID
	rCopy := *r
	rCopy.ID = &FhirString{Value: newID}
	return &rCopy
}

// UpdateVersion updates the meta field, increments the version, and modifies lastUpdated.
func (r *Resource) UpdateVersion(oldMeta *FhirMeta, versionIDAsTime bool) *Resource {
	newMeta := UpdateMeta(r.Meta, oldMeta, versionIDAsTime) // Reuse a helper function
	rCopy := *r
	rCopy.Meta = newMeta
	return &rCopy
}

// CopyWith creates a copy of the Resource with optional updates.
func (r *Resource) CopyWith(
	id *FhirString,
	meta *FhirMeta,
	implicitRules *FhirUri,
	language *CommonLanguages,
) *Resource {
	return &Resource{
		ResourceType:  r.ResourceType,
		ID:            ifNotNil(id, r.ID),
		Meta:          ifNotNil(meta, r.Meta),
		ImplicitRules: ifNotNil(implicitRules, r.ImplicitRules),
		Language:      ifNotNil(language, r.Language),
	}
}

// GenerateNewID generates a new unique ID for the resource.
func GenerateNewID() string {
	// Implement UUID generation logic or any unique ID generator here
	return "new-id-placeholder"
}

// UpdateMeta updates the meta data for version and lastUpdated.
func UpdateMeta(currentMeta, oldMeta *FhirMeta, versionIDAsTime bool) *FhirMeta {
	// Logic to update meta version, increment, and add lastUpdated timestamp
	newMeta := &FhirMeta{}
	if oldMeta != nil {
		newMeta = *oldMeta
	}
	// Mock example, replace with logic for version and timestamps
	newMeta.VersionID = "new-version"
	newMeta.LastUpdated = "2024-06-17T00:00:00Z"
	return newMeta
}

// ResourceFromJSON acts as a factory to construct a Resource from JSON data.
func ResourceFromJSON(data []byte) (*Resource, error) {
	var res Resource
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, errors.New("failed to parse resource from JSON")
	}
	return &res, nil
}

func ifNotNilSlice[T any](newSlice, oldSlice []T) []T {
	if newSlice != nil {
		return newSlice
	}
	return oldSlice
}
