// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json")

// Ingredient
// An ingredient of a manufactured item or pharmaceutical product.
type Ingredient struct {
	DomainResource
	Id *FhirString `json:"id,omitempty"`
	Meta *FhirMeta `json:"meta,omitempty"`
	ImplicitRules *FhirUri `json:"implicitrules,omitempty"`
	Language *CommonLanguages `json:"language,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Contained []*Resource `json:"contained,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Identifier *Identifier `json:"identifier,omitempty"`
	Status *PublicationStatus `json:"status,omitempty"`
	For_ []*Reference `json:"for,omitempty"`
	Role *CodeableConcept `json:"role,omitempty"`
	Function_ []*CodeableConcept `json:"function,omitempty"`
	AllergenicIndicator *FhirBoolean `json:"allergenicindicator,omitempty"`
	Manufacturer []*IngredientManufacturer `json:"manufacturer,omitempty"`
	Substance *IngredientSubstance `json:"substance,omitempty"`
}

// NewIngredient creates a new Ingredient instance
func NewIngredient() *Ingredient {
	return &Ingredient{}
}

// FromJSON populates Ingredient from JSON data
func (m *Ingredient) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts Ingredient to JSON data
func (m *Ingredient) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of Ingredient
func (m *Ingredient) Clone() *Ingredient {
	if m == nil { return nil }
	return &Ingredient{
		Id: m.Id.Clone(),
		Meta: m.Meta.Clone(),
		ImplicitRules: m.ImplicitRules.Clone(),
		Language: m.Language.Clone(),
		Text: m.Text.Clone(),
		Contained: cloneSlices(m.Contained),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Identifier: m.Identifier.Clone(),
		Status: m.Status.Clone(),
		For_: cloneSlices(m.For_),
		Role: m.Role.Clone(),
		Function_: cloneSlices(m.Function_),
		AllergenicIndicator: m.AllergenicIndicator.Clone(),
		Manufacturer: cloneSlices(m.Manufacturer),
		Substance: m.Substance.Clone(),
	}
}

// Equals checks for equality with another Ingredient instance
func (m *Ingredient) Equals(other *Ingredient) bool {
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
	if !m.Identifier.Equals(other.Identifier) { return false }
	if !m.Status.Equals(other.Status) { return false }
	if !compareSlices(m.For_, other.For_) { return false }
	if !m.Role.Equals(other.Role) { return false }
	if !compareSlices(m.Function_, other.Function_) { return false }
	if !m.AllergenicIndicator.Equals(other.AllergenicIndicator) { return false }
	if !compareSlices(m.Manufacturer, other.Manufacturer) { return false }
	if !m.Substance.Equals(other.Substance) { return false }
	return true
}

// IngredientManufacturer
// The organization(s) that manufacture this ingredient. Can be used to indicate:         1) Organizations we are aware of that manufacture this ingredient         2) Specific Manufacturer(s) currently being used         3) Set of organisations allowed to manufacture this ingredient for this product         Users must be clear on the application of context relevant to their use case.
type IngredientManufacturer struct {
	BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Role *IngredientManufacturerRole `json:"role,omitempty"`
	Manufacturer *Reference `json:"manufacturer,omitempty"`
}

// NewIngredientManufacturer creates a new IngredientManufacturer instance
func NewIngredientManufacturer() *IngredientManufacturer {
	return &IngredientManufacturer{}
}

// FromJSON populates IngredientManufacturer from JSON data
func (m *IngredientManufacturer) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts IngredientManufacturer to JSON data
func (m *IngredientManufacturer) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of IngredientManufacturer
func (m *IngredientManufacturer) Clone() *IngredientManufacturer {
	if m == nil { return nil }
	return &IngredientManufacturer{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Role: m.Role.Clone(),
		Manufacturer: m.Manufacturer.Clone(),
	}
}

// Equals checks for equality with another IngredientManufacturer instance
func (m *IngredientManufacturer) Equals(other *IngredientManufacturer) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Role.Equals(other.Role) { return false }
	if !m.Manufacturer.Equals(other.Manufacturer) { return false }
	return true
}

// IngredientSubstance
// The substance that comprises this ingredient.
type IngredientSubstance struct {
	BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Code *CodeableReference `json:"code,omitempty"`
	Strength []*IngredientStrength `json:"strength,omitempty"`
}

// NewIngredientSubstance creates a new IngredientSubstance instance
func NewIngredientSubstance() *IngredientSubstance {
	return &IngredientSubstance{}
}

// FromJSON populates IngredientSubstance from JSON data
func (m *IngredientSubstance) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts IngredientSubstance to JSON data
func (m *IngredientSubstance) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of IngredientSubstance
func (m *IngredientSubstance) Clone() *IngredientSubstance {
	if m == nil { return nil }
	return &IngredientSubstance{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Code: m.Code.Clone(),
		Strength: cloneSlices(m.Strength),
	}
}

// Equals checks for equality with another IngredientSubstance instance
func (m *IngredientSubstance) Equals(other *IngredientSubstance) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Code.Equals(other.Code) { return false }
	if !compareSlices(m.Strength, other.Strength) { return false }
	return true
}

// IngredientStrength
// The quantity of substance in the unit of presentation, or in the volume (or mass) of the single pharmaceutical product or manufactured item. The allowed repetitions do not represent different strengths, but are different representations - mathematically equivalent - of a single strength.
type IngredientStrength struct {
	BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	PresentationRatio *Ratio `json:"presentationratio,omitempty"`
	PresentationRatioRange *RatioRange `json:"presentationratiorange,omitempty"`
	TextPresentation *FhirString `json:"textpresentation,omitempty"`
	ConcentrationRatio *Ratio `json:"concentrationratio,omitempty"`
	ConcentrationRatioRange *RatioRange `json:"concentrationratiorange,omitempty"`
	TextConcentration *FhirString `json:"textconcentration,omitempty"`
	MeasurementPoint *FhirString `json:"measurementpoint,omitempty"`
	Country []*CodeableConcept `json:"country,omitempty"`
	ReferenceStrength []*IngredientReferenceStrength `json:"referencestrength,omitempty"`
}

// NewIngredientStrength creates a new IngredientStrength instance
func NewIngredientStrength() *IngredientStrength {
	return &IngredientStrength{}
}

// FromJSON populates IngredientStrength from JSON data
func (m *IngredientStrength) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts IngredientStrength to JSON data
func (m *IngredientStrength) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of IngredientStrength
func (m *IngredientStrength) Clone() *IngredientStrength {
	if m == nil { return nil }
	return &IngredientStrength{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		PresentationRatio: m.PresentationRatio.Clone(),
		PresentationRatioRange: m.PresentationRatioRange.Clone(),
		TextPresentation: m.TextPresentation.Clone(),
		ConcentrationRatio: m.ConcentrationRatio.Clone(),
		ConcentrationRatioRange: m.ConcentrationRatioRange.Clone(),
		TextConcentration: m.TextConcentration.Clone(),
		MeasurementPoint: m.MeasurementPoint.Clone(),
		Country: cloneSlices(m.Country),
		ReferenceStrength: cloneSlices(m.ReferenceStrength),
	}
}

// Equals checks for equality with another IngredientStrength instance
func (m *IngredientStrength) Equals(other *IngredientStrength) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.PresentationRatio.Equals(other.PresentationRatio) { return false }
	if !m.PresentationRatioRange.Equals(other.PresentationRatioRange) { return false }
	if !m.TextPresentation.Equals(other.TextPresentation) { return false }
	if !m.ConcentrationRatio.Equals(other.ConcentrationRatio) { return false }
	if !m.ConcentrationRatioRange.Equals(other.ConcentrationRatioRange) { return false }
	if !m.TextConcentration.Equals(other.TextConcentration) { return false }
	if !m.MeasurementPoint.Equals(other.MeasurementPoint) { return false }
	if !compareSlices(m.Country, other.Country) { return false }
	if !compareSlices(m.ReferenceStrength, other.ReferenceStrength) { return false }
	return true
}

// IngredientReferenceStrength
// Strength expressed in terms of a reference substance. For when the ingredient strength is additionally expressed as equivalent to the strength of some other closely related substance (e.g. salt vs. base). Reference strength represents the strength (quantitative composition) of the active moiety of the active substance. There are situations when the active substance and active moiety are different, therefore both a strength and a reference strength are needed.
type IngredientReferenceStrength struct {
	BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Substance *CodeableReference `json:"substance,omitempty"`
	StrengthRatio *Ratio `json:"strengthratio,omitempty"`
	StrengthRatioRange *RatioRange `json:"strengthratiorange,omitempty"`
	MeasurementPoint *FhirString `json:"measurementpoint,omitempty"`
	Country []*CodeableConcept `json:"country,omitempty"`
}

// NewIngredientReferenceStrength creates a new IngredientReferenceStrength instance
func NewIngredientReferenceStrength() *IngredientReferenceStrength {
	return &IngredientReferenceStrength{}
}

// FromJSON populates IngredientReferenceStrength from JSON data
func (m *IngredientReferenceStrength) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts IngredientReferenceStrength to JSON data
func (m *IngredientReferenceStrength) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of IngredientReferenceStrength
func (m *IngredientReferenceStrength) Clone() *IngredientReferenceStrength {
	if m == nil { return nil }
	return &IngredientReferenceStrength{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Substance: m.Substance.Clone(),
		StrengthRatio: m.StrengthRatio.Clone(),
		StrengthRatioRange: m.StrengthRatioRange.Clone(),
		MeasurementPoint: m.MeasurementPoint.Clone(),
		Country: cloneSlices(m.Country),
	}
}

// Equals checks for equality with another IngredientReferenceStrength instance
func (m *IngredientReferenceStrength) Equals(other *IngredientReferenceStrength) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Substance.Equals(other.Substance) { return false }
	if !m.StrengthRatio.Equals(other.StrengthRatio) { return false }
	if !m.StrengthRatioRange.Equals(other.StrengthRatioRange) { return false }
	if !m.MeasurementPoint.Equals(other.MeasurementPoint) { return false }
	if !compareSlices(m.Country, other.Country) { return false }
	return true
}
