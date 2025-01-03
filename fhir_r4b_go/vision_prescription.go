// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"
)

// VisionPrescription
// An authorization for the provision of glasses and/or contact lenses to a patient.
type VisionPrescription struct {
	extends DomainResource
	Id *FhirString `json:"id,omitempty"`
	Meta *FhirMeta `json:"meta,omitempty"`
	ImplicitRules *FhirUri `json:"implicitrules,omitempty"`
	Language *CommonLanguages `json:"language,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Contained []*Resource `json:"contained,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Status *FinancialResourceStatusCodes `json:"status,omitempty"`
	Created *FhirDateTime `json:"created,omitempty"`
	Patient *Reference `json:"patient,omitempty"`
	Encounter *Reference `json:"encounter,omitempty"`
	DateWritten *FhirDateTime `json:"datewritten,omitempty"`
	Prescriber *Reference `json:"prescriber,omitempty"`
	LensSpecification []*VisionPrescriptionLensSpecification `json:"lensspecification,omitempty"`
}

// NewVisionPrescription creates a new VisionPrescription instance.
func NewVisionPrescription() *VisionPrescription {
	return &VisionPrescription{}
}

// UnmarshalJSON populates VisionPrescription from JSON data.
func (m *VisionPrescription) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Meta *FhirMeta `json:"meta,omitempty"`
		ImplicitRules *FhirUri `json:"implicitrules,omitempty"`
		Language *CommonLanguages `json:"language,omitempty"`
		Text *Narrative `json:"text,omitempty"`
		Contained []*Resource `json:"contained,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Identifier []*Identifier `json:"identifier,omitempty"`
		Status *FinancialResourceStatusCodes `json:"status,omitempty"`
		Created *FhirDateTime `json:"created,omitempty"`
		Patient *Reference `json:"patient,omitempty"`
		Encounter *Reference `json:"encounter,omitempty"`
		DateWritten *FhirDateTime `json:"datewritten,omitempty"`
		Prescriber *Reference `json:"prescriber,omitempty"`
		LensSpecification []*VisionPrescriptionLensSpecification `json:"lensspecification,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Meta = temp.Meta
	m.ImplicitRules = temp.ImplicitRules
	m.Language = temp.Language
	m.Text = temp.Text
	m.Contained = temp.Contained
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Identifier = temp.Identifier
	m.Status = temp.Status
	m.Created = temp.Created
	m.Patient = temp.Patient
	m.Encounter = temp.Encounter
	m.DateWritten = temp.DateWritten
	m.Prescriber = temp.Prescriber
	m.LensSpecification = temp.LensSpecification
	return nil
}

// MarshalJSON converts VisionPrescription to JSON data.
func (m *VisionPrescription) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Meta *FhirMeta `json:"meta,omitempty"`
		ImplicitRules interface{} `json:"implicitrules,omitempty"`
		ImplicitRulesElement map[string]interface{} `json:"_implicitrules,omitempty"`
		Language *CommonLanguages `json:"language,omitempty"`
		Text *Narrative `json:"text,omitempty"`
		Contained []*Resource `json:"contained,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Identifier []*Identifier `json:"identifier,omitempty"`
		Status *FinancialResourceStatusCodes `json:"status,omitempty"`
		Created interface{} `json:"created,omitempty"`
		CreatedElement map[string]interface{} `json:"_created,omitempty"`
		Patient *Reference `json:"patient,omitempty"`
		Encounter *Reference `json:"encounter,omitempty"`
		DateWritten interface{} `json:"datewritten,omitempty"`
		DateWrittenElement map[string]interface{} `json:"_datewritten,omitempty"`
		Prescriber *Reference `json:"prescriber,omitempty"`
		LensSpecification []*VisionPrescriptionLensSpecification `json:"lensspecification,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Meta = m.Meta
	if m.ImplicitRules != nil && m.ImplicitRules.Value != nil {
		output.ImplicitRules = m.ImplicitRules.Value
		if m.ImplicitRules.Element != nil {
			output.ImplicitRulesElement = toMapOrNil(m.ImplicitRules.Element.MarshalJSON())
		}
	}
	output.Language = m.Language
	output.Text = m.Text
	output.Contained = m.Contained
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.Identifier = m.Identifier
	output.Status = m.Status
	if m.Created != nil && m.Created.Value != nil {
		output.Created = m.Created.Value
		if m.Created.Element != nil {
			output.CreatedElement = toMapOrNil(m.Created.Element.MarshalJSON())
		}
	}
	output.Patient = m.Patient
	output.Encounter = m.Encounter
	if m.DateWritten != nil && m.DateWritten.Value != nil {
		output.DateWritten = m.DateWritten.Value
		if m.DateWritten.Element != nil {
			output.DateWrittenElement = toMapOrNil(m.DateWritten.Element.MarshalJSON())
		}
	}
	output.Prescriber = m.Prescriber
	output.LensSpecification = m.LensSpecification
	return json.Marshal(output)
}

// Clone creates a deep copy of VisionPrescription.
func (m *VisionPrescription) Clone() *VisionPrescription {
	if m == nil { return nil }
	return &VisionPrescription{
		Id: m.Id.Clone(),
		Meta: m.Meta.Clone(),
		ImplicitRules: m.ImplicitRules.Clone(),
		Language: m.Language.Clone(),
		Text: m.Text.Clone(),
		Contained: cloneSlices(m.Contained),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Identifier: cloneSlices(m.Identifier),
		Status: m.Status.Clone(),
		Created: m.Created.Clone(),
		Patient: m.Patient.Clone(),
		Encounter: m.Encounter.Clone(),
		DateWritten: m.DateWritten.Clone(),
		Prescriber: m.Prescriber.Clone(),
		LensSpecification: cloneSlices(m.LensSpecification),
	}
}

// Equals checks equality between two VisionPrescription instances.
func (m *VisionPrescription) Equals(other *VisionPrescription) bool {
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
	if !compareSlices(m.Identifier, other.Identifier) { return false }
	if !m.Status.Equals(other.Status) { return false }
	if !m.Created.Equals(other.Created) { return false }
	if !m.Patient.Equals(other.Patient) { return false }
	if !m.Encounter.Equals(other.Encounter) { return false }
	if !m.DateWritten.Equals(other.DateWritten) { return false }
	if !m.Prescriber.Equals(other.Prescriber) { return false }
	if !compareSlices(m.LensSpecification, other.LensSpecification) { return false }
	return true
}

// VisionPrescriptionLensSpecification
// Contain the details of  the individual lens specifications and serves as the authorization for the fullfillment by certified professionals.
type VisionPrescriptionLensSpecification struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Product *CodeableConcept `json:"product,omitempty"`
	Eye *VisionEyes `json:"eye,omitempty"`
	Sphere *FhirDecimal `json:"sphere,omitempty"`
	Cylinder *FhirDecimal `json:"cylinder,omitempty"`
	Axis *FhirInteger `json:"axis,omitempty"`
	Prism []*VisionPrescriptionPrism `json:"prism,omitempty"`
	Add *FhirDecimal `json:"add,omitempty"`
	Power *FhirDecimal `json:"power,omitempty"`
	BackCurve *FhirDecimal `json:"backcurve,omitempty"`
	Diameter *FhirDecimal `json:"diameter,omitempty"`
	Duration *Quantity `json:"duration,omitempty"`
	Color *FhirString `json:"color,omitempty"`
	Brand *FhirString `json:"brand,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
}

// NewVisionPrescriptionLensSpecification creates a new VisionPrescriptionLensSpecification instance.
func NewVisionPrescriptionLensSpecification() *VisionPrescriptionLensSpecification {
	return &VisionPrescriptionLensSpecification{}
}

// UnmarshalJSON populates VisionPrescriptionLensSpecification from JSON data.
func (m *VisionPrescriptionLensSpecification) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Product *CodeableConcept `json:"product,omitempty"`
		Eye *VisionEyes `json:"eye,omitempty"`
		Sphere *FhirDecimal `json:"sphere,omitempty"`
		Cylinder *FhirDecimal `json:"cylinder,omitempty"`
		Axis *FhirInteger `json:"axis,omitempty"`
		Prism []*VisionPrescriptionPrism `json:"prism,omitempty"`
		Add *FhirDecimal `json:"add,omitempty"`
		Power *FhirDecimal `json:"power,omitempty"`
		BackCurve *FhirDecimal `json:"backcurve,omitempty"`
		Diameter *FhirDecimal `json:"diameter,omitempty"`
		Duration *Quantity `json:"duration,omitempty"`
		Color *FhirString `json:"color,omitempty"`
		Brand *FhirString `json:"brand,omitempty"`
		Note []*Annotation `json:"note,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Product = temp.Product
	m.Eye = temp.Eye
	m.Sphere = temp.Sphere
	m.Cylinder = temp.Cylinder
	m.Axis = temp.Axis
	m.Prism = temp.Prism
	m.Add = temp.Add
	m.Power = temp.Power
	m.BackCurve = temp.BackCurve
	m.Diameter = temp.Diameter
	m.Duration = temp.Duration
	m.Color = temp.Color
	m.Brand = temp.Brand
	m.Note = temp.Note
	return nil
}

// MarshalJSON converts VisionPrescriptionLensSpecification to JSON data.
func (m *VisionPrescriptionLensSpecification) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Product *CodeableConcept `json:"product,omitempty"`
		Eye *VisionEyes `json:"eye,omitempty"`
		Sphere interface{} `json:"sphere,omitempty"`
		SphereElement map[string]interface{} `json:"_sphere,omitempty"`
		Cylinder interface{} `json:"cylinder,omitempty"`
		CylinderElement map[string]interface{} `json:"_cylinder,omitempty"`
		Axis interface{} `json:"axis,omitempty"`
		AxisElement map[string]interface{} `json:"_axis,omitempty"`
		Prism []*VisionPrescriptionPrism `json:"prism,omitempty"`
		Add interface{} `json:"add,omitempty"`
		AddElement map[string]interface{} `json:"_add,omitempty"`
		Power interface{} `json:"power,omitempty"`
		PowerElement map[string]interface{} `json:"_power,omitempty"`
		BackCurve interface{} `json:"backcurve,omitempty"`
		BackCurveElement map[string]interface{} `json:"_backcurve,omitempty"`
		Diameter interface{} `json:"diameter,omitempty"`
		DiameterElement map[string]interface{} `json:"_diameter,omitempty"`
		Duration *Quantity `json:"duration,omitempty"`
		Color interface{} `json:"color,omitempty"`
		ColorElement map[string]interface{} `json:"_color,omitempty"`
		Brand interface{} `json:"brand,omitempty"`
		BrandElement map[string]interface{} `json:"_brand,omitempty"`
		Note []*Annotation `json:"note,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	output.Product = m.Product
	output.Eye = m.Eye
	if m.Sphere != nil && m.Sphere.Value != nil {
		output.Sphere = m.Sphere.Value
		if m.Sphere.Element != nil {
			output.SphereElement = toMapOrNil(m.Sphere.Element.MarshalJSON())
		}
	}
	if m.Cylinder != nil && m.Cylinder.Value != nil {
		output.Cylinder = m.Cylinder.Value
		if m.Cylinder.Element != nil {
			output.CylinderElement = toMapOrNil(m.Cylinder.Element.MarshalJSON())
		}
	}
	if m.Axis != nil && m.Axis.Value != nil {
		output.Axis = m.Axis.Value
		if m.Axis.Element != nil {
			output.AxisElement = toMapOrNil(m.Axis.Element.MarshalJSON())
		}
	}
	output.Prism = m.Prism
	if m.Add != nil && m.Add.Value != nil {
		output.Add = m.Add.Value
		if m.Add.Element != nil {
			output.AddElement = toMapOrNil(m.Add.Element.MarshalJSON())
		}
	}
	if m.Power != nil && m.Power.Value != nil {
		output.Power = m.Power.Value
		if m.Power.Element != nil {
			output.PowerElement = toMapOrNil(m.Power.Element.MarshalJSON())
		}
	}
	if m.BackCurve != nil && m.BackCurve.Value != nil {
		output.BackCurve = m.BackCurve.Value
		if m.BackCurve.Element != nil {
			output.BackCurveElement = toMapOrNil(m.BackCurve.Element.MarshalJSON())
		}
	}
	if m.Diameter != nil && m.Diameter.Value != nil {
		output.Diameter = m.Diameter.Value
		if m.Diameter.Element != nil {
			output.DiameterElement = toMapOrNil(m.Diameter.Element.MarshalJSON())
		}
	}
	output.Duration = m.Duration
	if m.Color != nil && m.Color.Value != nil {
		output.Color = m.Color.Value
		if m.Color.Element != nil {
			output.ColorElement = toMapOrNil(m.Color.Element.MarshalJSON())
		}
	}
	if m.Brand != nil && m.Brand.Value != nil {
		output.Brand = m.Brand.Value
		if m.Brand.Element != nil {
			output.BrandElement = toMapOrNil(m.Brand.Element.MarshalJSON())
		}
	}
	output.Note = m.Note
	return json.Marshal(output)
}

// Clone creates a deep copy of VisionPrescriptionLensSpecification.
func (m *VisionPrescriptionLensSpecification) Clone() *VisionPrescriptionLensSpecification {
	if m == nil { return nil }
	return &VisionPrescriptionLensSpecification{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Product: m.Product.Clone(),
		Eye: m.Eye.Clone(),
		Sphere: m.Sphere.Clone(),
		Cylinder: m.Cylinder.Clone(),
		Axis: m.Axis.Clone(),
		Prism: cloneSlices(m.Prism),
		Add: m.Add.Clone(),
		Power: m.Power.Clone(),
		BackCurve: m.BackCurve.Clone(),
		Diameter: m.Diameter.Clone(),
		Duration: m.Duration.Clone(),
		Color: m.Color.Clone(),
		Brand: m.Brand.Clone(),
		Note: cloneSlices(m.Note),
	}
}

// Equals checks equality between two VisionPrescriptionLensSpecification instances.
func (m *VisionPrescriptionLensSpecification) Equals(other *VisionPrescriptionLensSpecification) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Product.Equals(other.Product) { return false }
	if !m.Eye.Equals(other.Eye) { return false }
	if !m.Sphere.Equals(other.Sphere) { return false }
	if !m.Cylinder.Equals(other.Cylinder) { return false }
	if !m.Axis.Equals(other.Axis) { return false }
	if !compareSlices(m.Prism, other.Prism) { return false }
	if !m.Add.Equals(other.Add) { return false }
	if !m.Power.Equals(other.Power) { return false }
	if !m.BackCurve.Equals(other.BackCurve) { return false }
	if !m.Diameter.Equals(other.Diameter) { return false }
	if !m.Duration.Equals(other.Duration) { return false }
	if !m.Color.Equals(other.Color) { return false }
	if !m.Brand.Equals(other.Brand) { return false }
	if !compareSlices(m.Note, other.Note) { return false }
	return true
}

// VisionPrescriptionPrism
// Allows for adjustment on two axis.
type VisionPrescriptionPrism struct {
	extends BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Amount *FhirDecimal `json:"amount,omitempty"`
	Base *VisionBase `json:"base,omitempty"`
}

// NewVisionPrescriptionPrism creates a new VisionPrescriptionPrism instance.
func NewVisionPrescriptionPrism() *VisionPrescriptionPrism {
	return &VisionPrescriptionPrism{}
}

// UnmarshalJSON populates VisionPrescriptionPrism from JSON data.
func (m *VisionPrescriptionPrism) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Amount *FhirDecimal `json:"amount,omitempty"`
		Base *VisionBase `json:"base,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.ModifierExtension = temp.ModifierExtension
	m.Amount = temp.Amount
	m.Base = temp.Base
	return nil
}

// MarshalJSON converts VisionPrescriptionPrism to JSON data.
func (m *VisionPrescriptionPrism) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
		Amount interface{} `json:"amount,omitempty"`
		AmountElement map[string]interface{} `json:"_amount,omitempty"`
		Base *VisionBase `json:"base,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.ModifierExtension = m.ModifierExtension
	if m.Amount != nil && m.Amount.Value != nil {
		output.Amount = m.Amount.Value
		if m.Amount.Element != nil {
			output.AmountElement = toMapOrNil(m.Amount.Element.MarshalJSON())
		}
	}
	output.Base = m.Base
	return json.Marshal(output)
}

// Clone creates a deep copy of VisionPrescriptionPrism.
func (m *VisionPrescriptionPrism) Clone() *VisionPrescriptionPrism {
	if m == nil { return nil }
	return &VisionPrescriptionPrism{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Amount: m.Amount.Clone(),
		Base: m.Base.Clone(),
	}
}

// Equals checks equality between two VisionPrescriptionPrism instances.
func (m *VisionPrescriptionPrism) Equals(other *VisionPrescriptionPrism) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Amount.Equals(other.Amount) { return false }
	if !m.Base.Equals(other.Base) { return false }
	return true
}

