package fhir_r4b_go

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type FhirTime struct {
	Value       *string  `json:"-"` // Holds the string representation
	Hour        *int     `json:"hour,omitempty"`
	Minute      *int     `json:"minute,omitempty"`
	Second      *int     `json:"second,omitempty"`
	Millisecond *int     `json:"millisecond,omitempty"`
	Element     *Element `json:"_value,omitempty"`
}

// NewFhirTimeFromComponents creates a new FhirTime.
func NewFhirTimeFromComponents(hour, minute, second, millisecond *int, element *Element) *FhirTime {
	ft := &FhirTime{
		Hour:        hour,
		Minute:      minute,
		Second:      second,
		Millisecond: millisecond,
		Element:     element,
	}
	value := ft.ToString()
	ft.Value = &value
	return ft
}

// ToString updates the `Value` field with the string representation.
func (f *FhirTime) ToString() string {
	if f.Value != nil {
		return *f.Value
	}

	if f.Hour == nil {
		return ""
	}

	result := fmt.Sprintf("%02d", *f.Hour)
	if f.Minute != nil {
		result += fmt.Sprintf(":%02d", *f.Minute)
		if f.Second != nil {
			secondStr := fmt.Sprintf(":%02d", *f.Second)
			if f.Millisecond != nil {
				secondStr += fmt.Sprintf(".%03d", *f.Millisecond)
			}
			result += secondStr
		}
	}

	f.Value = &result // Cache the result
	return result
}

// NewFhirTimeFromString parses a FHIR-compliant time string.
func NewFhirTimeFromString(input string, element *Element) (*FhirTime, error) {
	if !validateFhirTime(input) {
		return nil, fmt.Errorf("invalid FHIR time format: %s", input)
	}

	parts := strings.Split(input, ":")
	hour := strToInt(parts[0])
	minute, second, millisecond := (*int)(nil), (*int)(nil), (*int)(nil)

	if len(parts) > 1 {
		minute = strToInt(parts[1])
		if len(parts) > 2 {
			secondParts := strings.Split(parts[2], ".")
			second = strToInt(secondParts[0])
			if len(secondParts) > 1 {
				millisecond = strToInt(secondParts[1])
			}
		}
	}

	return &FhirTime{
		Hour:        hour,
		Minute:      minute,
		Second:      second,
		Millisecond: millisecond,
		Element:     element,
	}, nil
}

// NewFhirTimeFromMap creates a new FhirTime from a map representation.
func NewFhirTimeFromMap(data map[string]interface{}) (*FhirTime, error) {
	// Validate and extract the "value" field
	value, ok := data["value"].(string)
	if !ok || value == "" {
		return nil, fmt.Errorf("missing or invalid 'value' in map")
	}

	// Parse the time string into its components
	parsedTime, err := NewFhirTimeFromString(value, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to parse FHIR time: %w", err)
	}

	// Handle optional "_value" field for metadata
	if elementData, ok := data["_value"].(map[string]interface{}); ok {
		parsedTime.Element = &Element{}
		if err := mapToStruct(elementData, parsedTime.Element); err != nil {
			return nil, fmt.Errorf("failed to parse '_value' element: %w", err)
		}
	}

	return parsedTime, nil
}

// MarshalJSON serializes FhirTime to JSON.
func (f *FhirTime) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"value": f.ToString(),
	}
	if f.Element != nil {
		data["_value"] = f.Element
	}
	return json.Marshal(data)
}

func (f *FhirTime) UnmarshalJSON(data []byte) error {
	temp := struct {
		Value   string   `json:"value"`
		Element *Element `json:"_value"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	parsed, err := NewFhirTimeFromString(temp.Value, temp.Element)
	if err != nil {
		return err
	}

	*f = *parsed
	return nil
}

// Clone creates a deep copy of FhirTime.
func (f *FhirTime) Clone() *FhirTime {
	if f == nil {
		return nil
	}
	return &FhirTime{
		Hour:        intPtrIfNotNil(f.Hour),
		Minute:      intPtrIfNotNil(f.Minute),
		Second:      intPtrIfNotNil(f.Second),
		Millisecond: intPtrIfNotNil(f.Millisecond),
		Element:     f.Element.Clone(),
	}
}

// Equals checks equality between two FhirTime instances.
func (f *FhirTime) Equals(other *FhirTime) bool {
	if f == nil && other == nil {
		return true
	}
	if f == nil || other == nil {
		return false
	}
	return f.ToString() == other.ToString() && f.Element.Equals(other.Element)
}

// validateFhirTime validates the FHIR time format using regex.
func validateFhirTime(input string) bool {
	timeRegex := `^([01][0-9]|2[0-3])(:([0-5][0-9])(:([0-5][0-9]|60)(\.[0-9]+)?)?)?$`
	re := regexp.MustCompile(timeRegex)
	return re.MatchString(input)
}

func strToInt(s string) *int {
	if val, err := strconv.Atoi(s); err == nil {
		return &val
	}
	return nil
}
