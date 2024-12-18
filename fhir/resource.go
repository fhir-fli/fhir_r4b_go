package fhir_r4b_go

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Resource represents the base definition for all FHIR resources.
type Resource struct {
	FhirBase
	ResourceType  ResourceType     `json:"resourceType,omitempty"`
	ID            *FhirString      `json:"id,omitempty"`
	Meta          *FhirMeta        `json:"meta,omitempty"`
	ImplicitRules *FhirUri         `json:"implicitRules,omitempty"`
	Language      *CommonLanguages `json:"language,omitempty"`
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

// Equals compares the current Resource with another Equalable.
func (r *Resource) Equals(other Equalable) bool {
	otherResource, ok := other.(*Resource)
	if !ok {
		return false
	}
	return r.ResourceType == otherResource.ResourceType &&
		r.ID.Equals(otherResource.ID) &&
		r.Meta.Equals(otherResource.Meta) &&
		r.ImplicitRules.Equals(otherResource.ImplicitRules) &&
		r.Language.Equals(otherResource.Language)
}

// Clone creates a deep copy of the Resource.
func (r *Resource) Clone() *Resource {
	if r == nil {
		return nil
	}
	return &Resource{
		FhirBase:      *r.FhirBase.Clone(),
		ResourceType:  r.ResourceType,
		ID:            r.ID.Clone(),
		Meta:          r.Meta.Clone(),
		ImplicitRules: r.ImplicitRules.Clone(),
		Language:      r.Language.Clone(),
	}
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

// GenerateNewID generates a new unique ID for the Resource.
func GenerateNewID() string {
	return uuid.New().String()
}

// UpdateMeta updates the metadata for a FHIR resource.
func UpdateMeta(currentMeta, oldMeta *FhirMeta, versionIDAsTime bool) *FhirMeta {
	newMeta := &FhirMeta{}
	if oldMeta != nil {
		*newMeta = *oldMeta
	}
	now := time.Now()
	newMeta.LastUpdated = FhirInstant{
		FhirDateTimeBase: NewFhirDateTimeBase(
			intPtr(now.Year()), now.Location() == time.UTC,
			intPtr(int(now.Month())), intPtr(now.Day()),
			intPtr(now.Hour()), intPtr(now.Minute()),
			intPtr(now.Second()), intPtr(now.Nanosecond()/1e6),
			nil, nil,
		),
	}
	if versionIDAsTime {
		newMeta.VersionId = FhirId{Value: strPtr(now.Format(time.RFC3339))}
	} else {
		newMeta.VersionId = FhirId{Value: strPtr("1")}
	}
	return newMeta
}
