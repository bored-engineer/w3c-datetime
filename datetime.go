package datetime

import (
	"fmt"
	"strings"
	"time"
)

// Precision is the granularity of the datetime
type Precision int

// Returns the layout expected by time.Parse or "" if unknown
func (p Precision) Layout() string {
	switch p {
	case PrecisionYear:
		return "2006"
	case PrecisionMonth:
		return "2006-01"
	case PrecisionDay:
		return "2006-01-02"
	case PrecisionHours:
		return "2006-01-02T15:04Z07:00"
	case PrecisionSeconds:
		return "2006-01-02T15:04:05Z07:00"
	case PrecisionNanoseconds:
		return "2006-01-02T15:04:05.999999999Z07:00"
	default:
		return ""
	}
}

// String implements fmt.Stringer
func (p Precision) String() string {
	switch p {
	case PrecisionYear:
		return "YYYY"
	case PrecisionMonth:
		return "YYYY-MM"
	case PrecisionDay:
		return "YYYY-MM-DD"
	case PrecisionHours:
		return "YYYY-MM-DDThh:mmTZD"
	case PrecisionSeconds:
		return "YYYY-MM-DDThh:mm:ssTZD"
	case PrecisionNanoseconds:
		return "YYYY-MM-DDThh:mm:ss.sTZD"
	default:
		return ""
	}
}

// ParsePrecision detects the precision level in the provided string
func ParsePrecision(datetime string) Precision {
	if idx := strings.IndexRune(datetime, 'T'); idx == 10 {
		jdx := strings.LastIndexAny(datetime, "Z-+")
		if jdx == -1 {
			jdx = len(datetime)
		}
		count := strings.Count(datetime[idx:jdx], ":")
		switch {
		case count == 1 && jdx == 16:
			return PrecisionHours
		case count == 2 && jdx >= 19:
			if strings.ContainsRune(datetime[idx:jdx], '.') {
				return PrecisionNanoseconds
			} else {
				return PrecisionSeconds
			}
		}
	} else if idx == -1 {
		count := strings.Count(datetime, "-")
		switch {
		case count == 0 && len(datetime) == 4:
			return PrecisionYear
		case count == 1 && len(datetime) == 7:
			return PrecisionMonth
		case count == 2 && len(datetime) == 10:
			return PrecisionDay
		}
	}
	return PrecisionUnknown
}

const (
	PrecisionUnknown Precision = iota
	// Year: YYYY (eg 1997)
	PrecisionYear
	// Year and month: YYYY-MM (eg 1997-07)
	PrecisionMonth
	// Complete date: YYYY-MM-DD (eg 1997-07-16)
	PrecisionDay
	// Complete date plus hours and minutes: YYYY-MM-DDThh:mmTZD (eg 1997-07-16T19:20+01:00)
	PrecisionHours
	// Complete date plus hours, minutes and seconds: YYYY-MM-DDThh:mm:ssTZD (eg 1997-07-16T19:20:30+01:00)
	PrecisionSeconds
	// Complete date plus hours, minutes, seconds and a decimal fraction of a second YYYY-MM-DDThh:mm:ss.sTZD (eg 1997-07-16T19:20:30.45+01:00)
	PrecisionNanoseconds
)

// Time represents a W3C datetime
// https://www.w3.org/TR/NOTE-datetime
type Time struct {
	time.Time
	Precision Precision
}

// New creates a Time with the highest precision (PrecisionNanoseconds)
func New(t time.Time) Time {
	return NewWithPrecision(t, PrecisionNanoseconds)
}

// New creates a Time with the specified precision
func NewWithPrecision(t time.Time, p Precision) Time {
	return Time{Time: t, Precision: p}
}

// String implements fmt.Stringer
func (t Time) String() string {
	return t.Time.Format(t.Precision.Layout())
}

// UnmarshalText implements encoding.TextUnmarshaler
func (t *Time) UnmarshalText(text []byte) (err error) {
	*t, err = Parse(string(text))
	return
}

// MarshalText implements encoding.TextMarshaler
func (t Time) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

// Parse parses a formatted datetime string and returns the time value/precision it represents.
func Parse(datetime string) (Time, error) {
	p := ParsePrecision(datetime)
	if p == PrecisionUnknown {
		return Time{}, fmt.Errorf("unknown datetime precision: %q", datetime)
	}
	t, err := time.Parse(p.Layout(), datetime)
	if err != nil {
		return Time{}, fmt.Errorf("time.Parse failed: %w", err)
	}
	return Time{Time: t, Precision: p}, nil
}
