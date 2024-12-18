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
	newMeta.LastUpdated = &FhirInstant{
		FhirDateTimeBase: *NewFhirDateTimeBase(
			intPtr(now.Year()), now.Location() == time.UTC,
			intPtr(int(now.Month())), intPtr(now.Day()),
			intPtr(now.Hour()), intPtr(now.Minute()),
			intPtr(now.Second()), intPtr(now.Nanosecond()/1e6),
			nil, nil,
		),
	}
	if versionIDAsTime {
		newMeta.VersionId = &FhirId{Value: strPtr(now.Format(time.RFC3339))}
	} else {
		newMeta.VersionId = &FhirId{Value: strPtr("1")}
	}
	return newMeta
}

// newResourceFromType creates a new instance of a FHIR resource based on its type.
func newResourceFromType(resourceType string) interface{} {
	switch resourceType {
	case "Account":
		return &Account{}
	case "ActivityDefinition":
		return &ActivityDefinition{}
	case "AdministrableProductDefinition":
		return &AdministrableProductDefinition{}
	case "AdverseEvent":
		return &AdverseEvent{}
	case "AllergyIntolerance":
		return &AllergyIntolerance{}
	case "Appointment":
		return &Appointment{}
	case "AppointmentResponse":
		return &AppointmentResponse{}
	case "AuditEvent":
		return &AuditEvent{}
	case "Basic":
		return &Basic{}
	case "Binary":
		return &Binary{}
	case "BiologicallyDerivedProduct":
		return &BiologicallyDerivedProduct{}
	case "BodyStructure":
		return &BodyStructure{}
	case "Bundle":
		return &Bundle{}
	case "CapabilityStatement":
		return &CapabilityStatement{}
	case "CarePlan":
		return &CarePlan{}
	case "CareTeam":
		return &CareTeam{}
	case "CatalogEntry":
		return &CatalogEntry{}
	case "ChargeItem":
		return &ChargeItem{}
	case "ChargeItemDefinition":
		return &ChargeItemDefinition{}
	case "Citation":
		return &Citation{}
	case "Claim":
		return &Claim{}
	case "ClaimResponse":
		return &ClaimResponse{}
	case "ClinicalImpression":
		return &ClinicalImpression{}
	case "ClinicalUseDefinition":
		return &ClinicalUseDefinition{}
	case "CodeSystem":
		return &CodeSystem{}
	case "Communication":
		return &Communication{}
	case "CommunicationRequest":
		return &CommunicationRequest{}
	case "CompartmentDefinition":
		return &CompartmentDefinition{}
	case "Composition":
		return &Composition{}
	case "ConceptMap":
		return &ConceptMap{}
	case "Condition":
		return &Condition{}
	case "Consent":
		return &Consent{}
	case "Contract":
		return &Contract{}
	case "Coverage":
		return &Coverage{}
	case "CoverageEligibilityRequest":
		return &CoverageEligibilityRequest{}
	case "CoverageEligibilityResponse":
		return &CoverageEligibilityResponse{}
	case "DetectedIssue":
		return &DetectedIssue{}
	case "Device":
		return &Device{}
	case "DeviceDefinition":
		return &DeviceDefinition{}
	case "DeviceMetric":
		return &DeviceMetric{}
	case "DeviceRequest":
		return &DeviceRequest{}
	case "DeviceUseStatement":
		return &DeviceUseStatement{}
	case "DiagnosticReport":
		return &DiagnosticReport{}
	case "DocumentManifest":
		return &DocumentManifest{}
	case "DocumentReference":
		return &DocumentReference{}
	case "Encounter":
		return &Encounter{}
	case "EnrollmentRequest":
		return &EnrollmentRequest{}
	case "EnrollmentResponse":
		return &EnrollmentResponse{}
	case "EpisodeOfCare":
		return &EpisodeOfCare{}
	case "EventDefinition":
		return &EventDefinition{}
	case "Evidence":
		return &Evidence{}
	case "EvidenceReport":
		return &EvidenceReport{}
	case "EvidenceVariable":
		return &EvidenceVariable{}
	case "ExampleScenario":
		return &ExampleScenario{}
	case "ExplanationOfBenefit":
		return &ExplanationOfBenefit{}
	case "FamilyMemberHistory":
		return &FamilyMemberHistory{}
	case "Endpoint":
		return &FhirEndpoint{}
	case "Group":
		return &FhirGroup{}
	case "List":
		return &FhirList{}
	case "Flag":
		return &Flag{}
	case "Goal":
		return &Goal{}
	case "GraphDefinition":
		return &GraphDefinition{}
	case "GuidanceResponse":
		return &GuidanceResponse{}
	case "HealthcareService":
		return &HealthcareService{}
	case "ImagingStudy":
		return &ImagingStudy{}
	case "Immunization":
		return &Immunization{}
	case "ImmunizationEvaluation":
		return &ImmunizationEvaluation{}
	case "ImmunizationRecommendation":
		return &ImmunizationRecommendation{}
	case "ImplementationGuide":
		return &ImplementationGuide{}
	case "Ingredient":
		return &Ingredient{}
	case "InsurancePlan":
		return &InsurancePlan{}
	case "Invoice":
		return &Invoice{}
	case "Library":
		return &Library{}
	case "Linkage":
		return &Linkage{}
	case "Location":
		return &Location{}
	case "ManufacturedItemDefinition":
		return &ManufacturedItemDefinition{}
	case "Measure":
		return &Measure{}
	case "MeasureReport":
		return &MeasureReport{}
	case "Media":
		return &Media{}
	case "Medication":
		return &Medication{}
	case "MedicationAdministration":
		return &MedicationAdministration{}
	case "MedicationDispense":
		return &MedicationDispense{}
	case "MedicationKnowledge":
		return &MedicationKnowledge{}
	case "MedicationRequest":
		return &MedicationRequest{}
	case "MedicationStatement":
		return &MedicationStatement{}
	case "MedicinalProductDefinition":
		return &MedicinalProductDefinition{}
	case "MessageDefinition":
		return &MessageDefinition{}
	case "MessageHeader":
		return &MessageHeader{}
	case "MolecularSequence":
		return &MolecularSequence{}
	case "NamingSystem":
		return &NamingSystem{}
	case "NutritionOrder":
		return &NutritionOrder{}
	case "NutritionProduct":
		return &NutritionProduct{}
	case "Observation":
		return &Observation{}
	case "ObservationDefinition":
		return &ObservationDefinition{}
	case "OperationDefinition":
		return &OperationDefinition{}
	case "OperationOutcome":
		return &OperationOutcome{}
	case "Organization":
		return &Organization{}
	case "OrganizationAffiliation":
		return &OrganizationAffiliation{}
	case "PackagedProductDefinition":
		return &PackagedProductDefinition{}
	case "Parameters":
		return &Parameters{}
	case "Patient":
		return &Patient{}
	case "PaymentNotice":
		return &PaymentNotice{}
	case "PaymentReconciliation":
		return &PaymentReconciliation{}
	case "Person":
		return &Person{}
	case "PlanDefinition":
		return &PlanDefinition{}
	case "Practitioner":
		return &Practitioner{}
	case "PractitionerRole":
		return &PractitionerRole{}
	case "Procedure":
		return &Procedure{}
	case "Provenance":
		return &Provenance{}
	case "Questionnaire":
		return &Questionnaire{}
	case "QuestionnaireResponse":
		return &QuestionnaireResponse{}
	case "RegulatedAuthorization":
		return &RegulatedAuthorization{}
	case "RelatedPerson":
		return &RelatedPerson{}
	case "RequestGroup":
		return &RequestGroup{}
	case "ResearchDefinition":
		return &ResearchDefinition{}
	case "ResearchElementDefinition":
		return &ResearchElementDefinition{}
	case "ResearchStudy":
		return &ResearchStudy{}
	case "ResearchSubject":
		return &ResearchSubject{}
	case "RiskAssessment":
		return &RiskAssessment{}
	case "Schedule":
		return &Schedule{}
	case "SearchParameter":
		return &SearchParameter{}
	case "ServiceRequest":
		return &ServiceRequest{}
	case "Slot":
		return &Slot{}
	case "Specimen":
		return &Specimen{}
	case "SpecimenDefinition":
		return &SpecimenDefinition{}
	case "StructureDefinition":
		return &StructureDefinition{}
	case "StructureMap":
		return &StructureMap{}
	case "Subscription":
		return &Subscription{}
	case "SubscriptionStatus":
		return &SubscriptionStatus{}
	case "SubscriptionTopic":
		return &SubscriptionTopic{}
	case "Substance":
		return &Substance{}
	case "SubstanceDefinition":
		return &SubstanceDefinition{}
	case "SupplyDelivery":
		return &SupplyDelivery{}
	case "SupplyRequest":
		return &SupplyRequest{}
	case "Task":
		return &Task{}
	case "TerminologyCapabilities":
		return &TerminologyCapabilities{}
	case "TestReport":
		return &TestReport{}
	case "TestScript":
		return &TestScript{}
	case "ValueSet":
		return &ValueSet{}
	case "VerificationResult":
		return &VerificationResult{}
	case "VisionPrescription":
		return &VisionPrescription{}
	// Add more cases for other resource types
	default:
		return nil
	}
}
