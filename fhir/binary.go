// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"

)

// Binary
// A resource that represents the data of a single raw artifact as digital content accessible in its native format.  A Binary resource can contain any content, whether text, image, pdf, zip archive, etc.
type Binary struct {
	Resource
	// id
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id FhirString `json:"id,omitempty"`
	// meta
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta FhirMeta `json:"meta,omitempty"`
	// implicitRules
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules FhirUri `json:"implicitRules,omitempty"`
	// language
	// The base language in which the resource is written.
	Language CommonLanguages `json:"language,omitempty"`
	// contentType
	// MimeType of the binary content represented as a standard MimeType (BCP 13).
	ContentType FhirCode `json:"contentType,omitempty"`
	// securityContext
	// This element identifies another resource that can be used as a proxy of the security sensitivity to use when deciding and enforcing access control rules for the Binary resource. Given that the Binary resource contains very few elements that can be used to determine the sensitivity of the data and relationships to individuals, the referenced resource stands in as a proxy equivalent for this purpose. This referenced resource may be related to the Binary (e.g. Media, DocumentReference), or may be some non-related Resource purely as a security proxy. E.g. to identify that the binary resource relates to a patient, and access should only be granted to applications that have access to the patient.
	SecurityContext Reference `json:"securityContext,omitempty"`
	// data
	// The actual content, base64 encoded.
	Data FhirBase64Binary `json:"data,omitempty"`
}

// NewBinary creates a new Binary instance
func NewBinary(
	id FhirString,
	meta FhirMeta,
	implicitRules FhirUri,
	language CommonLanguages,
	contentType FhirCode,
	securityContext Reference,
	data FhirBase64Binary,
) *Binary {
	return &Binary{
		Id: id,
		Meta: meta,
		ImplicitRules: implicitRules,
		Language: language,
		ContentType: contentType,
		SecurityContext: securityContext,
		Data: data,
	}
}
// FromJSON populates Binary from JSON data
func (m *Binary) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts Binary to JSON data
func (m *Binary) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// CopyWith creates a modified copy of Binary
func (m *Binary) CopyWith(
	id *FhirString,
	meta *FhirMeta,
	implicitRules *FhirUri,
	language *CommonLanguages,
	contentType *FhirCode,
	securityContext *Reference,
	data *FhirBase64Binary,
) *Binary {
	return &Binary{
		Id: func() FhirString {
			if id != nil { return *id }
			return m.Id
		}(),
		Meta: func() FhirMeta {
			if meta != nil { return *meta }
			return m.Meta
		}(),
		ImplicitRules: func() FhirUri {
			if implicitRules != nil { return *implicitRules }
			return m.ImplicitRules
		}(),
		Language: func() CommonLanguages {
			if language != nil { return *language }
			return m.Language
		}(),
		ContentType: func() FhirCode {
			if contentType != nil { return *contentType }
			return m.ContentType
		}(),
		SecurityContext: func() Reference {
			if securityContext != nil { return *securityContext }
			return m.SecurityContext
		}(),
		Data: func() FhirBase64Binary {
			if data != nil { return *data }
			return m.Data
		}(),
	}
}