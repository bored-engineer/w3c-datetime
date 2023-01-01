package datetime

import (
	"testing"
	"time"
)

func TestPrecision(t *testing.T) {
	type test struct {
		input string
		want  Precision
	}

	tests := []test{
		{input: "1997", want: PrecisionYear},
		{input: "1997-07", want: PrecisionMonth},
		{input: "1997-07-16", want: PrecisionDay},
		{input: "1997-07-16T19:20+01:00", want: PrecisionHours},
		{input: "1997-07-16T19:20:30+01:00", want: PrecisionSeconds},
		{input: "1997-07-16T19:20:30.45+01:00", want: PrecisionNanoseconds},
		{input: "invalid", want: PrecisionUnknown},
	}

	for _, tc := range tests {
		have := ParsePrecision(tc.input)
		if have != tc.want {
			t.Fatalf("for input %q, expected: %s, got: %s", tc.input, tc.want, have)
		}
	}
}

func TestTime(t *testing.T) {
	type test struct {
		input string
		want  time.Time
	}

	loc := time.FixedZone("+0100", 60*60)

	tests := []test{
		{input: "1997", want: time.Date(1997, 1, 1, 0, 0, 0, 0, time.UTC)},
		{input: "1997-07", want: time.Date(1997, 7, 1, 0, 0, 0, 0, time.UTC)},
		{input: "1997-07-16", want: time.Date(1997, 7, 16, 0, 0, 0, 0, time.UTC)},
		{input: "1997-07-16T19:20+01:00", want: time.Date(1997, 7, 16, 19, 20, 0, 0, loc)},
		{input: "1997-07-16T19:20:30+01:00", want: time.Date(1997, 7, 16, 19, 20, 30, 0, loc)},
		{input: "1997-07-16T19:20:30.45+01:00", want: time.Date(1997, 7, 16, 19, 20, 30, 450000000, loc)},
		{input: "1994-11-05T08:15:30Z", want: time.Date(1994, 11, 5, 8, 15, 30, 0, time.UTC)},
	}

	for _, tc := range tests {
		var dt Time
		if err := dt.UnmarshalText([]byte(tc.input)); err != nil {
			t.Fatalf("(datetime.Time).UnmarshalText failed for %q: %s", tc.input, err)
		}
		if !tc.want.Equal(dt.Time) {
			t.Fatalf("for input %q, expected: %s, got: %s", tc.input, tc.want, dt.Time)
		}
		text, err := dt.MarshalText()
		if err != nil {
			t.Fatalf("(datetime.Time).MarshalText failed for %q: %s", tc.input, err)
		}
		if tc.input != string(text) {
			t.Fatalf("expected: %s, got: %s", tc.input, string(text))
		}
	}
}
