// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"

)

// ProductShelfLife
// The shelf-life and storage information for a medicinal product item or container can be described using this class.
type ProductShelfLife struct {
	BackboneType
	// id
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id FhirString `json:"id,omitempty"`
	// extension
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension_ []FhirExtension `json:"extension,omitempty"`
	// modifierExtension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
// 
// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []FhirExtension `json:"modifierExtension,omitempty"`
	// identifier
	// Unique identifier for the packaged Medicinal Product.
	Identifier Identifier `json:"identifier,omitempty"`
	// type
	// This describes the shelf life, taking into account various scenarios such as shelf life of the packaged Medicinal Product itself, shelf life after transformation where necessary and shelf life after the first opening of a bottle, etc. The shelf life type shall be specified using an appropriate controlled vocabulary The controlled term and the controlled term identifier shall be specified.
	Type_ CodeableConcept `json:"type,omitempty"`
	// period
	// The shelf life time period can be specified using a numerical value for the period of time and its unit of time measurement The unit of measurement shall be specified in accordance with ISO 11240 and the resulting terminology The symbol and the symbol identifier shall be used.
	Period Quantity `json:"period,omitempty"`
	// specialPrecautionsForStorage
	// Special precautions for storage, if any, can be specified using an appropriate controlled vocabulary The controlled term and the controlled term identifier shall be specified.
	SpecialPrecautionsForStorage []CodeableConcept `json:"specialPrecautionsForStorage,omitempty"`
}

// NewProductShelfLife creates a new ProductShelfLife instance
func NewProductShelfLife(
	id FhirString,
	extension_ []FhirExtension,
	modifierExtension []FhirExtension,
	identifier Identifier,
	type_ CodeableConcept,
	period Quantity,
	specialPrecautionsForStorage []CodeableConcept,
) *ProductShelfLife {
	return &ProductShelfLife{
		Id: id,
		Extension_: extension_,
		ModifierExtension: modifierExtension,
		Identifier: identifier,
		Type_: type_,
		Period: period,
		SpecialPrecautionsForStorage: specialPrecautionsForStorage,
	}
}
// FromJSON populates ProductShelfLife from JSON data
func (m *ProductShelfLife) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts ProductShelfLife to JSON data
func (m *ProductShelfLife) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// CopyWith creates a modified copy of ProductShelfLife
func (m *ProductShelfLife) CopyWith(
	id *FhirString,
	extension_ *[]FhirExtension,
	modifierExtension *[]FhirExtension,
	identifier *Identifier,
	type_ *CodeableConcept,
	period *Quantity,
	specialPrecautionsForStorage *[]CodeableConcept,
) *ProductShelfLife {
	return &ProductShelfLife{
		Id: func() FhirString {
			if id != nil { return *id }
			return m.Id
		}(),
		Extension_: func() []FhirExtension {
			if extension_ != nil { return *extension_ }
			return m.Extension_
		}(),
		ModifierExtension: func() []FhirExtension {
			if modifierExtension != nil { return *modifierExtension }
			return m.ModifierExtension
		}(),
		Identifier: func() Identifier {
			if identifier != nil { return *identifier }
			return m.Identifier
		}(),
		Type_: func() CodeableConcept {
			if type_ != nil { return *type_ }
			return m.Type_
		}(),
		Period: func() Quantity {
			if period != nil { return *period }
			return m.Period
		}(),
		SpecialPrecautionsForStorage: func() []CodeableConcept {
			if specialPrecautionsForStorage != nil { return *specialPrecautionsForStorage }
			return m.SpecialPrecautionsForStorage
		}(),
	}
}