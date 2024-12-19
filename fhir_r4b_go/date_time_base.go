package fhir_r4b_go

import (
	"encoding/json"
	"fmt"
	"time"
)

// FhirDateTimeBase represents the base structure for FHIR date and time types.
type FhirDateTimeBase struct {
	Year        *int     `json:"year,omitempty"`
	Month       *int     `json:"month,omitempty"`
	Day         *int     `json:"day,omitempty"`
	Hour        *int     `json:"hour,omitempty"`
	Minute      *int     `json:"minute,omitempty"`
	Second      *int     `json:"second,omitempty"`
	Millisecond *int     `json:"millisecond,omitempty"`
	Microsecond *string  `json:"microsecond,omitempty"`
	TimeZone    *float64 `json:"timezone,omitempty"` // Offset in hours
	IsUTC       bool     `json:"isUTC"`
	Value       *string  `json:"-"`                // Cached string representation
	Element     *Element `json:"_value,omitempty"` // Metadata element
}

// NewFhirDateTimeBase creates a new FhirDateTimeBase instance.
func NewFhirDateTimeBase(
	year *int,
	isUTC bool,
	month, day, hour, minute, second, millisecond *int,
	microsecond *string,
	timeZone *float64,
	element *Element,
) *FhirDateTimeBase {
	return &FhirDateTimeBase{
		Year:        year,
		Month:       month,
		Day:         day,
		Hour:        hour,
		Minute:      minute,
		Second:      second,
		Millisecond: millisecond,
		Microsecond: microsecond,
		TimeZone:    timeZone,
		IsUTC:       isUTC,
		Element:     element,
	}
}

// Value converts FhirDateTimeBase into a time.Time object.
func (f *FhirDateTimeBase) DateTimeValue() *time.Time {
	if f.Year == nil {
		return nil
	}

	year := *f.Year
	month := 1
	day := 1
	hour := 0
	minute := 0
	second := 0
	millisecond := 0

	if f.Month != nil {
		month = *f.Month
	}
	if f.Day != nil {
		day = *f.Day
	}
	if f.Hour != nil {
		hour = *f.Hour
	}
	if f.Minute != nil {
		minute = *f.Minute
	}
	if f.Second != nil {
		second = *f.Second
	}
	if f.Millisecond != nil {
		millisecond = *f.Millisecond
	}

	loc := time.UTC
	if !f.IsUTC && f.TimeZone != nil {
		zoneOffset := int(*f.TimeZone * 3600) // Convert hours to seconds
		loc = time.FixedZone("Offset", zoneOffset)
	}

	return timePtr(time.Date(year, time.Month(month), day, hour, minute, second, millisecond*1e6, loc))
}

// ToString formats the FHIR date-time as a string.
func (f *FhirDateTimeBase) ToString() string {
	if f.Year == nil {
		return ""
	}

	result := fmt.Sprintf("%04d", *f.Year)

	if f.Month != nil {
		result += fmt.Sprintf("-%02d", *f.Month)
	}
	if f.Day != nil {
		result += fmt.Sprintf("-%02d", *f.Day)
	}

	if f.Hour != nil {
		result += fmt.Sprintf("T%02d", *f.Hour)
		if f.Minute != nil {
			result += fmt.Sprintf(":%02d", *f.Minute)
			if f.Second != nil {
				result += fmt.Sprintf(":%02d", *f.Second)
				if f.Millisecond != nil || f.Microsecond != nil {
					milli := 0
					if f.Millisecond != nil {
						milli = *f.Millisecond
					}
					micro := ""
					if f.Microsecond != nil {
						micro = *f.Microsecond
					}
					result += fmt.Sprintf(".%03d%s", milli, micro)
				}
			}
		}
	}

	if f.IsUTC {
		result += "Z"
	} else if f.TimeZone != nil {
		hours := int(*f.TimeZone)
		minutes := int((*f.TimeZone - float64(hours)) * 60)
		sign := "+"
		if *f.TimeZone < 0 {
			sign = "-"
			hours = -hours
		}
		result += fmt.Sprintf("%s%02d:%02d", sign, hours, minutes)
	}

	return result
}

// Clone creates a deep copy of FhirDateTimeBase.
func (f *FhirDateTimeBase) Clone() *FhirDateTimeBase {
	if f == nil {
		return nil
	}
	return &FhirDateTimeBase{
		Year:        intPtrIfNotNil(f.Year),
		Month:       intPtrIfNotNil(f.Month),
		Day:         intPtrIfNotNil(f.Day),
		Hour:        intPtrIfNotNil(f.Hour),
		Minute:      intPtrIfNotNil(f.Minute),
		Second:      intPtrIfNotNil(f.Second),
		Millisecond: intPtrIfNotNil(f.Millisecond),
		Microsecond: strPtrIfNotNil(f.Microsecond),
		TimeZone:    floatPtrIfNotNil(f.TimeZone),
		IsUTC:       f.IsUTC,
		Element:     f.Element.Clone(),
	}
}

// MarshalJSON serializes FhirDateTimeBase to JSON.
func (f *FhirDateTimeBase) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"value": f.ToString(),
	}
	if f.Element != nil {
		data["_value"] = f.Element
	}
	return json.Marshal(data)
}

// UnmarshalJSON deserializes JSON into FhirDateTimeBase.
func (f *FhirDateTimeBase) UnmarshalJSON(data []byte) error {
	temp := struct {
		Value   string   `json:"value"`
		Element *Element `json:"_value"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	parsed, err := FhirDateTimeBaseFromString(temp.Value)
	if err != nil {
		return err
	}

	parsed.Element = temp.Element
	*f = *parsed
	return nil
}

// FhirDateTimeBaseFromString parses a FHIR date-time string into FhirDateTimeBase.
func FhirDateTimeBaseFromString(input string) (*FhirDateTimeBase, error) {
	t, err := time.Parse(time.RFC3339Nano, input)
	if err != nil {
		return nil, fmt.Errorf("invalid FHIR dateTime format: %v", err)
	}

	_, zoneOffset := t.Zone()
	timeZone := float64(zoneOffset) / 3600

	return &FhirDateTimeBase{
		Year:        intPtr(t.Year()),
		IsUTC:       t.Location() == time.UTC,
		Month:       intPtr(int(t.Month())),
		Day:         intPtr(t.Day()),
		Hour:        intPtr(t.Hour()),
		Minute:      intPtr(t.Minute()),
		Second:      intPtr(t.Second()),
		Millisecond: intPtr(t.Nanosecond() / 1e6),
		Microsecond: nil,
		TimeZone:    &timeZone,
		Element:     nil,
	}, nil
}

// CompareTo compares two FhirDateTimeBase instances.
func (f *FhirDateTimeBase) CompareTo(other *FhirDateTimeBase) int {
	t1 := f.DateTimeValue()
	t2 := other.DateTimeValue()

	if t1 == nil || t2 == nil {
		if t1 == nil && t2 == nil {
			return 0
		}
		if t1 == nil {
			return -1
		}
		return 1
	}

	if t1.Before(*t2) {
		return -1
	} else if t1.After(*t2) {
		return 1
	}
	return 0
}
