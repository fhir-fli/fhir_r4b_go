package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"regexp"
)

// FhirOid represents a validated OID value in the FHIR standard.
type FhirOid struct {
	Value   string   `json:"value,omitempty"`
	Element *Element `json:"_value,omitempty"`
}

// NewFhirOid creates a new FhirOid instance with validation.
func NewFhirOid(input string, element *Element) (*FhirOid, error) {
	if err := validateOid(input); err != nil {
		return nil, err
	}
	return &FhirOid{Value: input, Element: element}, nil
}

// validateOid ensures the input matches the OID pattern.
func validateOid(input string) error {
	pattern := regexp.MustCompile(`^urn:oid:[0-2](\.(0|[1-9][0-9]*))+$`)
	if !pattern.MatchString(input) {
		return errors.New("invalid FhirOid: does not match OID format")
	}
	return nil
}

// MarshalJSON serializes FhirOid into JSON.
func (f *FhirOid) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{}
	if f.Value != "" {
		data["value"] = f.Value
	}
	if f.Element != nil {
		data["_value"] = f.Element
	}
	return json.Marshal(data)
}

// UnmarshalJSON deserializes JSON into FhirOid.
func (f *FhirOid) UnmarshalJSON(data []byte) error {
	temp := struct {
		Value   string   `json:"value"`
		Element *Element `json:"_value"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	if err := validateOid(temp.Value); err != nil {
		return err
	}

	f.Value = temp.Value
	f.Element = temp.Element
	return nil
}

// String returns the string representation of the OID.
func (f *FhirOid) String() string {
	return f.Value
}

// Equal checks equality between two FhirOid instances.
func (f *FhirOid) Equals(other *FhirOid) bool {
	return f.Value == other.Value && f.Element.Equals(other.Element)
}

// Clone creates a deep copy of the FhirOid instance.
func (f *FhirOid) Clone() *FhirOid {
	var elementCopy *Element
	if f.Element != nil {
		elementCopy = f.Element.Clone()
	}
	return &FhirOid{
		Value:   f.Value,
		Element: elementCopy,
	}
}

// FromJSONList converts a list of JSON values into a slice of FhirOid.
func FromJSONListFhirOid(values []json.RawMessage) ([]*FhirOid, error) {
	var oids []*FhirOid
	for _, v := range values {
		var oid FhirOid
		if err := json.Unmarshal(v, &oid); err != nil {
			return nil, err
		}
		oids = append(oids, &oid)
	}
	return oids, nil
}

// ToJSONList serializes a slice of FhirOid into JSON.
func ToJSONList(oids []*FhirOid) ([]byte, error) {
	return json.Marshal(oids)
}
