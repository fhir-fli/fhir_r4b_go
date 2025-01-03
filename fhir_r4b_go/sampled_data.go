// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"
)

// SampledData
// A series of measurements taken by a device, with upper and lower limits. There may be more than one dimension in the data.
type SampledData struct {
	extends DataType
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	Origin *Quantity `json:"origin,omitempty"`
	Period *FhirDecimal `json:"period,omitempty"`
	Factor *FhirDecimal `json:"factor,omitempty"`
	LowerLimit *FhirDecimal `json:"lowerlimit,omitempty"`
	UpperLimit *FhirDecimal `json:"upperlimit,omitempty"`
	Dimensions *FhirPositiveInt `json:"dimensions,omitempty"`
	Data *FhirString `json:"data,omitempty"`
}

// NewSampledData creates a new SampledData instance.
func NewSampledData() *SampledData {
	return &SampledData{}
}

// UnmarshalJSON populates SampledData from JSON data.
func (m *SampledData) UnmarshalJSON(data []byte) error {
	temp := struct {
		Id *FhirString `json:"id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		Origin *Quantity `json:"origin,omitempty"`
		Period *FhirDecimal `json:"period,omitempty"`
		Factor *FhirDecimal `json:"factor,omitempty"`
		LowerLimit *FhirDecimal `json:"lowerlimit,omitempty"`
		UpperLimit *FhirDecimal `json:"upperlimit,omitempty"`
		Dimensions *FhirPositiveInt `json:"dimensions,omitempty"`
		Data *FhirString `json:"data,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Id = temp.Id
	m.Extension_ = temp.Extension_
	m.Origin = temp.Origin
	m.Period = temp.Period
	m.Factor = temp.Factor
	m.LowerLimit = temp.LowerLimit
	m.UpperLimit = temp.UpperLimit
	m.Dimensions = temp.Dimensions
	m.Data = temp.Data
	return nil
}

// MarshalJSON converts SampledData to JSON data.
func (m *SampledData) MarshalJSON() ([]byte, error) {
	output := struct {
		Id interface{} `json:"id,omitempty"`
		IdElement map[string]interface{} `json:"_id,omitempty"`
		Extension_ []*FhirExtension `json:"extension,omitempty"`
		Origin *Quantity `json:"origin,omitempty"`
		Period interface{} `json:"period,omitempty"`
		PeriodElement map[string]interface{} `json:"_period,omitempty"`
		Factor interface{} `json:"factor,omitempty"`
		FactorElement map[string]interface{} `json:"_factor,omitempty"`
		LowerLimit interface{} `json:"lowerlimit,omitempty"`
		LowerLimitElement map[string]interface{} `json:"_lowerlimit,omitempty"`
		UpperLimit interface{} `json:"upperlimit,omitempty"`
		UpperLimitElement map[string]interface{} `json:"_upperlimit,omitempty"`
		Dimensions interface{} `json:"dimensions,omitempty"`
		DimensionsElement map[string]interface{} `json:"_dimensions,omitempty"`
		Data interface{} `json:"data,omitempty"`
		DataElement map[string]interface{} `json:"_data,omitempty"`
	}{}
	if m.Id != nil && m.Id.Value != nil {
		output.Id = m.Id.Value
		if m.Id.Element != nil {
			output.IdElement = toMapOrNil(m.Id.Element.MarshalJSON())
		}
	}
	output.Extension_ = m.Extension_
	output.Origin = m.Origin
	if m.Period != nil && m.Period.Value != nil {
		output.Period = m.Period.Value
		if m.Period.Element != nil {
			output.PeriodElement = toMapOrNil(m.Period.Element.MarshalJSON())
		}
	}
	if m.Factor != nil && m.Factor.Value != nil {
		output.Factor = m.Factor.Value
		if m.Factor.Element != nil {
			output.FactorElement = toMapOrNil(m.Factor.Element.MarshalJSON())
		}
	}
	if m.LowerLimit != nil && m.LowerLimit.Value != nil {
		output.LowerLimit = m.LowerLimit.Value
		if m.LowerLimit.Element != nil {
			output.LowerLimitElement = toMapOrNil(m.LowerLimit.Element.MarshalJSON())
		}
	}
	if m.UpperLimit != nil && m.UpperLimit.Value != nil {
		output.UpperLimit = m.UpperLimit.Value
		if m.UpperLimit.Element != nil {
			output.UpperLimitElement = toMapOrNil(m.UpperLimit.Element.MarshalJSON())
		}
	}
	if m.Dimensions != nil && m.Dimensions.Value != nil {
		output.Dimensions = m.Dimensions.Value
		if m.Dimensions.Element != nil {
			output.DimensionsElement = toMapOrNil(m.Dimensions.Element.MarshalJSON())
		}
	}
	if m.Data != nil && m.Data.Value != nil {
		output.Data = m.Data.Value
		if m.Data.Element != nil {
			output.DataElement = toMapOrNil(m.Data.Element.MarshalJSON())
		}
	}
	return json.Marshal(output)
}

// Clone creates a deep copy of SampledData.
func (m *SampledData) Clone() *SampledData {
	if m == nil { return nil }
	return &SampledData{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		Origin: m.Origin.Clone(),
		Period: m.Period.Clone(),
		Factor: m.Factor.Clone(),
		LowerLimit: m.LowerLimit.Clone(),
		UpperLimit: m.UpperLimit.Clone(),
		Dimensions: m.Dimensions.Clone(),
		Data: m.Data.Clone(),
	}
}

// Equals checks equality between two SampledData instances.
func (m *SampledData) Equals(other *SampledData) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !m.Origin.Equals(other.Origin) { return false }
	if !m.Period.Equals(other.Period) { return false }
	if !m.Factor.Equals(other.Factor) { return false }
	if !m.LowerLimit.Equals(other.LowerLimit) { return false }
	if !m.UpperLimit.Equals(other.UpperLimit) { return false }
	if !m.Dimensions.Equals(other.Dimensions) { return false }
	if !m.Data.Equals(other.Data) { return false }
	return true
}

