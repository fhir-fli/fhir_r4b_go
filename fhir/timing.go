// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json"

)

// Timing
// Specifies an event that may occur multiple times. Timing schedules are used to record when things are planned, expected or requested to occur. The most common usage is in dosage instructions for medications. They are also used when planning care of various kinds, and may be used for reporting the schedule to which past regular activities were carried out.
type Timing struct {
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
	// event
	// Identifies specific times when the event occurs.
	Event []FhirDateTime `json:"event,omitempty"`
	// repeat
	// A set of rules that describe when the event is scheduled.
	Repeat TimingRepeat `json:"repeat,omitempty"`
	// code
	// A code for the timing schedule (or just text in code.text). Some codes such as BID are ubiquitous, but many institutions define their own additional codes. If a code is provided, the code is understood to be a complete statement of whatever is specified in the structured timing data, and either the code or the data may be used to interpret the Timing, with the exception that .repeat.bounds still applies over the code (and is not contained in the code).
	Code CodeableConcept `json:"code,omitempty"`
}

// NewTiming creates a new Timing instance
func NewTiming(
	id FhirString,
	extension_ []FhirExtension,
	modifierExtension []FhirExtension,
	event []FhirDateTime,
	repeat TimingRepeat,
	code CodeableConcept,
) *Timing {
	return &Timing{
		Id: id,
		Extension_: extension_,
		ModifierExtension: modifierExtension,
		Event: event,
		Repeat: repeat,
		Code: code,
	}
}
// FromJSON populates Timing from JSON data
func (m *Timing) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts Timing to JSON data
func (m *Timing) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// CopyWith creates a modified copy of Timing
func (m *Timing) CopyWith(
	id *FhirString,
	extension_ *[]FhirExtension,
	modifierExtension *[]FhirExtension,
	event *[]FhirDateTime,
	repeat *TimingRepeat,
	code *CodeableConcept,
) *Timing {
	return &Timing{
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
		Event: func() []FhirDateTime {
			if event != nil { return *event }
			return m.Event
		}(),
		Repeat: func() TimingRepeat {
			if repeat != nil { return *repeat }
			return m.Repeat
		}(),
		Code: func() CodeableConcept {
			if code != nil { return *code }
			return m.Code
		}(),
	}
}
// TimingRepeat
// A set of rules that describe when the event is scheduled.
type TimingRepeat struct {
	Element
	// id
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id FhirString `json:"id,omitempty"`
	// extension
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension_ []FhirExtension `json:"extension,omitempty"`
	// boundsDuration
	// Either a duration for the length of the timing schedule, a range of possible length, or outer bounds for start and/or end limits of the timing schedule.
	BoundsDuration FhirDuration `json:"boundsDuration,omitempty"`
	// boundsRange
	// Either a duration for the length of the timing schedule, a range of possible length, or outer bounds for start and/or end limits of the timing schedule.
	BoundsRange Range `json:"boundsRange,omitempty"`
	// boundsPeriod
	// Either a duration for the length of the timing schedule, a range of possible length, or outer bounds for start and/or end limits of the timing schedule.
	BoundsPeriod Period `json:"boundsPeriod,omitempty"`
	// count
	// A total count of the desired number of repetitions across the duration of the entire timing specification. If countMax is present, this element indicates the lower bound of the allowed range of count values.
	Count FhirPositiveInt `json:"count,omitempty"`
	// countMax
	// If present, indicates that the count is a range - so to perform the action between [count] and [countMax] times.
	CountMax FhirPositiveInt `json:"countMax,omitempty"`
	// duration
	// How long this thing happens for when it happens. If durationMax is present, this element indicates the lower bound of the allowed range of the duration.
	Duration FhirDecimal `json:"duration,omitempty"`
	// durationMax
	// If present, indicates that the duration is a range - so to perform the action between [duration] and [durationMax] time length.
	DurationMax FhirDecimal `json:"durationMax,omitempty"`
	// durationUnit
	// The units of time for the duration, in UCUM units.
	DurationUnit UnitsOfTime `json:"durationUnit,omitempty"`
	// frequency
	// The number of times to repeat the action within the specified period. If frequencyMax is present, this element indicates the lower bound of the allowed range of the frequency.
	Frequency FhirPositiveInt `json:"frequency,omitempty"`
	// frequencyMax
	// If present, indicates that the frequency is a range - so to repeat between [frequency] and [frequencyMax] times within the period or period range.
	FrequencyMax FhirPositiveInt `json:"frequencyMax,omitempty"`
	// period
	// Indicates the duration of time over which repetitions are to occur; e.g. to express "3 times per day", 3 would be the frequency and "1 day" would be the period. If periodMax is present, this element indicates the lower bound of the allowed range of the period length.
	Period FhirDecimal `json:"period,omitempty"`
	// periodMax
	// If present, indicates that the period is a range from [period] to [periodMax], allowing expressing concepts such as "do this once every 3-5 days.
	PeriodMax FhirDecimal `json:"periodMax,omitempty"`
	// periodUnit
	// The units of time for the period in UCUM units.
	PeriodUnit UnitsOfTime `json:"periodUnit,omitempty"`
	// dayOfWeek
	// If one or more days of week is provided, then the action happens only on the specified day(s).
	DayOfWeek []DaysOfWeek `json:"dayOfWeek,omitempty"`
	// timeOfDay
	// Specified time of day for action to take place.
	TimeOfDay []FhirTime `json:"timeOfDay,omitempty"`
	// when
	// An approximate time period during the day, potentially linked to an event of daily living that indicates when the action should occur.
	When []EventTiming `json:"when,omitempty"`
	// offset
	// The number of minutes from the event. If the event code does not indicate whether the minutes is before or after the event, then the offset is assumed to be after the event.
	Offset FhirUnsignedInt `json:"offset,omitempty"`
}

// NewTimingRepeat creates a new TimingRepeat instance
func NewTimingRepeat(
	id FhirString,
	extension_ []FhirExtension,
	boundsDuration FhirDuration,
	boundsRange Range,
	boundsPeriod Period,
	count FhirPositiveInt,
	countMax FhirPositiveInt,
	duration FhirDecimal,
	durationMax FhirDecimal,
	durationUnit UnitsOfTime,
	frequency FhirPositiveInt,
	frequencyMax FhirPositiveInt,
	period FhirDecimal,
	periodMax FhirDecimal,
	periodUnit UnitsOfTime,
	dayOfWeek []DaysOfWeek,
	timeOfDay []FhirTime,
	when []EventTiming,
	offset FhirUnsignedInt,
) *TimingRepeat {
	return &TimingRepeat{
		Id: id,
		Extension_: extension_,
		BoundsDuration: boundsDuration,
		BoundsRange: boundsRange,
		BoundsPeriod: boundsPeriod,
		Count: count,
		CountMax: countMax,
		Duration: duration,
		DurationMax: durationMax,
		DurationUnit: durationUnit,
		Frequency: frequency,
		FrequencyMax: frequencyMax,
		Period: period,
		PeriodMax: periodMax,
		PeriodUnit: periodUnit,
		DayOfWeek: dayOfWeek,
		TimeOfDay: timeOfDay,
		When: when,
		Offset: offset,
	}
}
// FromJSON populates TimingRepeat from JSON data
func (m *TimingRepeat) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts TimingRepeat to JSON data
func (m *TimingRepeat) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// CopyWith creates a modified copy of TimingRepeat
func (m *TimingRepeat) CopyWith(
	id *FhirString,
	extension_ *[]FhirExtension,
	boundsDuration *FhirDuration,
	boundsRange *Range,
	boundsPeriod *Period,
	count *FhirPositiveInt,
	countMax *FhirPositiveInt,
	duration *FhirDecimal,
	durationMax *FhirDecimal,
	durationUnit *UnitsOfTime,
	frequency *FhirPositiveInt,
	frequencyMax *FhirPositiveInt,
	period *FhirDecimal,
	periodMax *FhirDecimal,
	periodUnit *UnitsOfTime,
	dayOfWeek *[]DaysOfWeek,
	timeOfDay *[]FhirTime,
	when *[]EventTiming,
	offset *FhirUnsignedInt,
) *TimingRepeat {
	return &TimingRepeat{
		Id: func() FhirString {
			if id != nil { return *id }
			return m.Id
		}(),
		Extension_: func() []FhirExtension {
			if extension_ != nil { return *extension_ }
			return m.Extension_
		}(),
		BoundsDuration: func() FhirDuration {
			if boundsDuration != nil { return *boundsDuration }
			return m.BoundsDuration
		}(),
		BoundsRange: func() Range {
			if boundsRange != nil { return *boundsRange }
			return m.BoundsRange
		}(),
		BoundsPeriod: func() Period {
			if boundsPeriod != nil { return *boundsPeriod }
			return m.BoundsPeriod
		}(),
		Count: func() FhirPositiveInt {
			if count != nil { return *count }
			return m.Count
		}(),
		CountMax: func() FhirPositiveInt {
			if countMax != nil { return *countMax }
			return m.CountMax
		}(),
		Duration: func() FhirDecimal {
			if duration != nil { return *duration }
			return m.Duration
		}(),
		DurationMax: func() FhirDecimal {
			if durationMax != nil { return *durationMax }
			return m.DurationMax
		}(),
		DurationUnit: func() UnitsOfTime {
			if durationUnit != nil { return *durationUnit }
			return m.DurationUnit
		}(),
		Frequency: func() FhirPositiveInt {
			if frequency != nil { return *frequency }
			return m.Frequency
		}(),
		FrequencyMax: func() FhirPositiveInt {
			if frequencyMax != nil { return *frequencyMax }
			return m.FrequencyMax
		}(),
		Period: func() FhirDecimal {
			if period != nil { return *period }
			return m.Period
		}(),
		PeriodMax: func() FhirDecimal {
			if periodMax != nil { return *periodMax }
			return m.PeriodMax
		}(),
		PeriodUnit: func() UnitsOfTime {
			if periodUnit != nil { return *periodUnit }
			return m.PeriodUnit
		}(),
		DayOfWeek: func() []DaysOfWeek {
			if dayOfWeek != nil { return *dayOfWeek }
			return m.DayOfWeek
		}(),
		TimeOfDay: func() []FhirTime {
			if timeOfDay != nil { return *timeOfDay }
			return m.TimeOfDay
		}(),
		When: func() []EventTiming {
			if when != nil { return *when }
			return m.When
		}(),
		Offset: func() FhirUnsignedInt {
			if offset != nil { return *offset }
			return m.Offset
		}(),
	}
}