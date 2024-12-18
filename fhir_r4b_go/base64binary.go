package fhir_r4b_go

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

// FhirBase64Binary represents the FHIR primitive type `FhirBase64Binary`.
type FhirBase64Binary struct {
	Value   *string  `json:"value,omitempty"`  // The Base64 encoded string
	Element *Element `json:"_value,omitempty"` // Additional metadata (FHIR element)
}

// NewFhirBase64Binary creates a new FhirBase64Binary instance with validation.
func NewFhirBase64Binary(input string, element *Element) (*FhirBase64Binary, error) {
	cleanedInput := strings.ReplaceAll(input, " ", "")
	if !isValidBase64(cleanedInput) {
		return nil, errors.New("invalid Base64 encoded string")
	}
	return &FhirBase64Binary{
		Value:   &cleanedInput,
		Element: element,
	}, nil
}

// isValidBase64 checks whether a string is a valid Base64 encoded string.
func isValidBase64(input string) bool {
	if len(input)%4 != 0 {
		return false
	}
	_, err := base64.StdEncoding.DecodeString(input)
	return err == nil
}

// FromJSON initializes a FhirBase64Binary from JSON input.
func (b *FhirBase64Binary) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, b); err != nil {
		return errors.New("failed to parse FhirBase64Binary from JSON: " + err.Error())
	}
	return nil
}

// ToJSON converts the FhirBase64Binary to JSON.
func (b *FhirBase64Binary) ToJSON() ([]byte, error) {
	return json.Marshal(b)
}

// Decode decodes the FhirBase64Binary value into bytes.
func (b *FhirBase64Binary) Decode() ([]byte, error) {
	if b.Value == nil {
		return nil, errors.New("no Base64 value to decode")
	}
	data, err := base64.StdEncoding.DecodeString(*b.Value)
	if err != nil {
		return nil, errors.New("failed to decode Base64 string: " + err.Error())
	}
	return data, nil
}

// DetectFileType attempts to detect the file type from the decoded data.
func (b *FhirBase64Binary) DetectFileType() FileType {
	data, err := b.Decode()
	if err != nil || len(data) < 4 {
		return UnknownFileType
	}

	switch {
	case len(data) >= 3 && data[0] == 0xFF && data[1] == 0xD8 && data[2] == 0xFF:
		return JPEG
	case len(data) >= 8 && data[0] == 0x89 && string(data[1:4]) == "PNG":
		return PNG
	case len(data) >= 4 && string(data[:4]) == "%PDF":
		return PDF
	case len(data) >= 2 && data[0] == 0x50 && data[1] == 0x4B:
		return ZIP
	case len(data) >= 3 && data[0] == 0x47 && string(data[1:3]) == "IF":
		return GIF
	case len(data) >= 2 && data[0] == 0x42 && data[1] == 0x4D:
		return BMP
	default:
		return UnknownFileType
	}
}

// Clone creates a deep copy of FhirBase64Binary.
func (b *FhirBase64Binary) Clone() *FhirBase64Binary {
	if b == nil {
		return nil
	}
	clone := &FhirBase64Binary{}
	if b.Value != nil {
		val := *b.Value
		clone.Value = &val
	}
	if b.Element != nil {
		clone.Element = b.Element.Clone()
	}
	return clone
}

// String returns the string representation of the FhirBase64Binary.
func (b *FhirBase64Binary) String() string {
	if b.Value != nil {
		return *b.Value
	}
	return ""
}

// FileType defines possible file types detected from FhirBase64Binary.
type FileType string

const (
	JPEG            FileType = "jpeg"
	PNG             FileType = "png"
	PDF             FileType = "pdf"
	ZIP             FileType = "zip"
	GIF             FileType = "gif"
	BMP             FileType = "bmp"
	UnknownFileType FileType = "unknown"
)

// Equals checks for equality between FhirBase64Binary instances.
func (b *FhirBase64Binary) Equals(other *FhirBase64Binary) bool {
	if b == nil && other == nil {
		return true
	}
	if b == nil || other == nil {
		return false
	}
	if (b.Value == nil && other.Value != nil) || (b.Value != nil && other.Value == nil) {
		return false
	}
	if b.Value != nil && *b.Value != *other.Value {
		return false
	}
	return b.Element.Equals(other.Element)
}
