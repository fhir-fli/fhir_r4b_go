package fhir_r4b_go

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

// FhirBase64Binary represents the FHIR primitive type `FhirBase64Binary`.
type FhirBase64Binary struct {
	Value   *string     `json:"value,omitempty"`   // The Base64 encoded string
	Element *Element    `json:"_value,omitempty"`  // Additional metadata (FHIR element)
}

// NewFhirBase64Binary creates a new FhirBase64Binary instance with validation.
func NewFhirBase64Binary(input string, element *Element) (*FhirBase64Binary, error) {
	if isValidBase64(input) {
		return &FhirBase64Binary{
			Value:   &input,
			Element: element,
		}, nil
	}

	// Attempt to clean whitespace and validate again
	cleanedInput := strings.ReplaceAll(input, " ", "")
	if isValidBase64(cleanedInput) {
		return &FhirBase64Binary{
			Value:   &cleanedInput,
			Element: element,
		}, nil
	}
	return nil, errors.New("invalid Base64 encoded string")
}

// IsValidBase64 checks whether a string is a valid Base64 encoded string.
func isValidBase64(input string) bool {
	_, err := base64.StdEncoding.DecodeString(input)
	return err == nil && len(input)%4 == 0
}

// FromJSON initializes a FhirBase64Binary from JSON input.
func (b *FhirBase64Binary) FromJSON(data []byte) error {
	return json.Unmarshal(data, b)
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
	return base64.StdEncoding.DecodeString(*b.Value)
}

// DetectFileType attempts to detect the file type from the decoded data.
func (b *FhirBase64Binary) DetectFileType() FileType {
	data, err := b.Decode()
	if err != nil || len(data) < 4 {
		return UnknownFileType
	}

	switch {
	case data[0] == 0xFF && data[1] == 0xD8 && data[2] == 0xFF:
		return JPEG
	case data[0] == 0x89 && string(data[1:4]) == "PNG":
		return PNG
	case string(data[:4]) == "%PDF":
		return PDF
	case data[0] == 0x50 && data[1] == 0x4B:
		return ZIP
	case data[0] == 0x47 && string(data[1:3]) == "IF":
		return GIF
	case data[0] == 0x42 && data[1] == 0x4D:
		return BMP
	default:
		return UnknownFileType
	}
}

// CopyWith creates a modified copy of FhirBase64Binary.
func (b *FhirBase64Binary) CopyWith(newValue *string, element *Element) *FhirBase64Binary {
	value := ifNotNil(newValue, b.Value)
	return &FhirBase64Binary{
		Value:   value,
		Element: ifNotNil(element, b.Element),
	}
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
