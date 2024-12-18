package fhir_r4b_go

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// FhirTime represents a FHIR-compliant time value.
type FhirTime struct {
	Value   string
	Element *Element
}

// NewFhirTime validates and creates a new FhirTime instance.
func NewFhirTime(input string, element *Element) (*FhirTime, error) {
	if !validateFhirTime(input) {
		return nil, fmt.Errorf("invalid FHIR time format: %s", input)
	}
	return &FhirTime{Value: input, Element: element}, nil
}

// validateFhirTime validates the FHIR time format using regex.
func validateFhirTime(input string) bool {
	timeRegex := `^([01][0-9]|2[0-3])(:([0-5][0-9])(:([0-5][0-9]|60)(\.[0-9]+)?)?)?$`
	re := regexp.MustCompile(timeRegex)
	return re.MatchString(input)
}

// NewFhirTimeFromUnits creates FhirTime from individual components.
func NewFhirTimeFromUnits(hour, minute, second, millisecond *int) *FhirTime {
	timeParts := []string{padZero(hour, 2)}

	if minute != nil {
		timeParts = append(timeParts, padZero(minute, 2))
		if second != nil {
			secondStr := padZero(second, 2)
			if millisecond != nil {
				secondStr += fmt.Sprintf(".%03d", *millisecond)
			}
			timeParts = append(timeParts, secondStr)
		}
	}

	return &FhirTime{Value: strings.Join(timeParts, ":")}
}

// padZero pads an integer to a string with leading zeros.
func padZero(value *int, length int) string {
	if value == nil {
		return ""
	}
	return fmt.Sprintf("%0*d", length, *value)
}

// Hour returns the hour part of the FhirTime.
func (f *FhirTime) Hour() *int {
	return parseTimePart(f.Value, 0)
}

// Minute returns the minute part of the FhirTime.
func (f *FhirTime) Minute() *int {
	return parseTimePart(f.Value, 1)
}

// Second returns the second part of the FhirTime.
func (f *FhirTime) Second() *int {
	if parts := strings.Split(f.Value, ":"); len(parts) > 2 {
		secParts := strings.Split(parts[2], ".")
		return strToInt(secParts[0])
	}
	return nil
}

// Millisecond returns the millisecond part of the FhirTime.
func (f *FhirTime) Millisecond() *int {
	if parts := strings.Split(f.Value, "."); len(parts) > 1 {
		ms := strings.TrimRight(parts[1], "0")
		if ms == "" {
			return nil
		}
		return strToInt(ms)
	}
	return nil
}

// parseTimePart extracts a part of the time string based on index.
func parseTimePart(timeStr string, index int) *int {
	parts := strings.Split(timeStr, ":")
	if len(parts) > index {
		return strToInt(parts[index])
	}
	return nil
}

// strToInt converts a string to an integer pointer.
func strToInt(s string) *int {
	if val, err := strconv.Atoi(s); err == nil {
		return &val
	}
	return nil
}

// Add adds hours, minutes, seconds, and milliseconds to FhirTime.
func (f *FhirTime) Add(hours, minutes, seconds, milliseconds int) *FhirTime {
	t := f.toTime()
	t = t.Add(time.Duration(hours)*time.Hour +
		time.Duration(minutes)*time.Minute +
		time.Duration(seconds)*time.Second +
		time.Duration(milliseconds)*time.Millisecond)

	// Store results in variables to take their addresses
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()
	millisecond := t.Nanosecond() / 1e6

	return NewFhirTimeFromUnits(&hour, &minute, &second, &millisecond)
}

// Subtract subtracts hours, minutes, seconds, and milliseconds from FhirTime.
func (f *FhirTime) Subtract(hours, minutes, seconds, milliseconds int) *FhirTime {
	return f.Add(-hours, -minutes, -seconds, -milliseconds)
}

// toTime converts FhirTime to a time.Time object.
func (f *FhirTime) toTime() time.Time {
	hour, minute, second := 0, 0, 0
	millisecond := 0

	if f.Hour() != nil {
		hour = *f.Hour()
	}
	if f.Minute() != nil {
		minute = *f.Minute()
	}
	if f.Second() != nil {
		second = *f.Second()
	}
	if f.Millisecond() != nil {
		millisecond = *f.Millisecond()
	}

	return time.Date(0, 1, 1, hour, minute, second, millisecond*1e6, time.UTC)
}

// MarshalJSON serializes FhirTime to JSON.
func (f *FhirTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{"value": f.Value})
}

// UnmarshalJSON deserializes JSON into FhirTime.
func (f *FhirTime) UnmarshalJSON(data []byte) error {
	var aux map[string]string
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if value, ok := aux["value"]; ok {
		if !validateFhirTime(value) {
			return fmt.Errorf("invalid FHIR time format: %s", value)
		}
		f.Value = value
		return nil
	}
	return fmt.Errorf("invalid JSON for FhirTime")
}

// CompareTo compares two FhirTime objects.
func (f *FhirTime) CompareTo(other *FhirTime) int {
	if f.Value > other.Value {
		return 1
	} else if f.Value < other.Value {
		return -1
	}
	return 0
}

// String returns the FhirTime as a string.
func (f *FhirTime) String() string {
	return f.Value
}
