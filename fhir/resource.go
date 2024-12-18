package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
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
	baseURL := fmt.Sprintf("http://hl7.org/fhir/%s", r.ResourceType)
	parsedURL, _ := url.Parse(baseURL) // Assume valid URL, handle errors appropriately in practice

	return &Reference{
		Reference: FhirString{Value: r.Path()},
		Type_:     FhirUri{Value: parsedURL},
	}
}

// NewIDIfNoID returns the Resource with a new ID if there is no current ID.
func (r *Resource) NewIDIfNoID() *Resource {
	if r.ID == nil {
		return r.NewID()
	}
	return r
}

// NewID generates a unique ID for the Resource dynamically.
func (r *Resource) NewID() *Resource {
	newID := GenerateNewID()
	rCopy := *r
	rCopy.ID = &FhirString{Value: newID}
	return &rCopy
}

// UpdateVersion updates the meta field, increments the version, and modifies lastUpdated.
func (r *Resource) UpdateVersion(oldMeta *FhirMeta, versionIDAsTime bool) *Resource {
	newMeta := UpdateMeta(r.Meta, oldMeta, versionIDAsTime)
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

// GenerateNewID generates a new unique ID for the Resource.
func GenerateNewID() string {
	return uuid.New().String()
}

// UpdateMeta updates the metadata for a FHIR resource, including version and lastUpdated fields.
func UpdateMeta(currentMeta, oldMeta *FhirMeta, versionIDAsTime bool) *FhirMeta {
	// Create a new meta instance, copying values from the old meta if it exists
	newMeta := &FhirMeta{}
	if oldMeta != nil {
		*newMeta = *oldMeta // Deep copy of oldMeta
	}

	// Set the VersionId field
	if versionIDAsTime {
		newMeta.VersionId = FhirId{Value: strPtr(time.Now().Format(time.RFC3339))}
	} else {
		newMeta.VersionId = FhirId{Value: strPtr(incrementVersion(getStringPointerValue(oldMeta.VersionId.Value, "1")))}
	}

	// Set the LastUpdated field using the current timestamp
	now := time.Now()
	newMeta.LastUpdated = FhirInstant{
		FhirDateTimeBase: NewFhirDateTimeBase(
			intPtr(now.Year()), now.Location() == time.UTC,
			intPtr(int(now.Month())), intPtr(now.Day()),
			intPtr(now.Hour()), intPtr(now.Minute()),
			intPtr(now.Second()), intPtr(now.Nanosecond()/1e6),
			nil, nil, // No microseconds or additional time zone provided
		),
	}

	return newMeta
}

// Helper function to retrieve a string value from a pointer, returning a default if nil.
func getStringPointerValue(ptr *string, defaultValue string) string {
	if ptr != nil {
		return *ptr
	}
	return defaultValue
}

// Helper function to increment a numeric version string.
func incrementVersion(version string) string {
	var versionNum int
	fmt.Sscanf(version, "%d", &versionNum)
	return fmt.Sprintf("%d", versionNum+1)
}

// ResourceFromJSON acts as a factory to construct a Resource from JSON data.
func ResourceFromJSON(data []byte) (*Resource, error) {
	var res Resource
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, errors.New("failed to parse resource from JSON")
	}
	return &res, nil
}

// Helper for converting a string to a pointer
func strPtr(s string) *string {
	return &s
}
