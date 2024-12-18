package fhir_r4b_go

import (
	"fmt"
	"time"
)

// FhirDateTimeBase is the base struct for FHIR date types.
type FhirDateTimeBase struct {
	Year        *int
	Month       *int
	Day         *int
	Hour        *int
	Minute      *int
	Second      *int
	Millisecond *int
	Microsecond *string
	TimeZone    *float64 // Represent timezone offset as float (e.g., +02:00 → 2.0)
	IsUTC       bool
}

// NewFhirDateTimeBase initializes a new FhirDateTimeBase.
func NewFhirDateTimeBase(year *int, isUTC bool, month, day, hour, minute, second, millisecond *int, microsecond *string, timeZone *float64) FhirDateTimeBase {
	return FhirDateTimeBase{
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
	}
}

// Value returns the FhirDateTimeBase as a time.Time object.
func (f FhirDateTimeBase) Value() *time.Time {
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

	// Construct time.Time with microsecond handling
	t := time.Date(year, time.Month(month), day, hour, minute, second, millisecond*1000, loc)
	return &t
}

// ToString returns the formatted FHIR date-time string.
func (f FhirDateTimeBase) ToString() string {
	if f.Year == nil {
		return ""
	}

	// Start building the date string
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

	// Append time zone
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

// CompareTo compares two FhirDateTimeBase instances.
func (f FhirDateTimeBase) CompareTo(other FhirDateTimeBase) int {
	t1 := f.Value()
	t2 := other.Value()

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

// ToJSON serializes the FhirDateTimeBase into JSON format.
func (f FhirDateTimeBase) ToJSON() map[string]interface{} {
	if f.Year == nil {
		return nil
	}
	return map[string]interface{}{
		"value": f.ToString(),
	}
}

// intPtr returns a pointer to an int.
func intPtr(v int) *int {
	return &v
}

func FhirDateTimeBaseFromString(input string) (*FhirDateTimeBase, error) {
	t, err := time.Parse(time.RFC3339Nano, input)
	if err != nil {
		return nil, fmt.Errorf("invalid FHIR dateTime format: %v", err)
	}

	// Capture the zone name (unused) and offset in seconds
	_, zoneOffset := t.Zone()
	timeZone := float64(zoneOffset) / 3600 // Convert offset to hours

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
		TimeZone:    &timeZone, // Use TimeZone instead of Offset
	}, nil
}
