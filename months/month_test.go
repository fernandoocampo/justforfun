package months_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/months"
)

func TestWhichMonth(t *testing.T) {
	cases := map[string]struct {
		monthsAhead  int
		currentMonth string
		want         string
	}{
		"feb_may": {
			monthsAhead:  3,
			currentMonth: "february",
			want:         "may",
		},
		"feb_jun": {
			monthsAhead:  4,
			currentMonth: "february",
			want:         "june",
		},
		"jun_dec": {
			monthsAhead:  6,
			currentMonth: "june",
			want:         "december",
		},
		"nov_jan": {
			monthsAhead:  2,
			currentMonth: "november",
			want:         "january",
		},
		"nov_feb": {
			monthsAhead:  3,
			currentMonth: "november",
			want:         "february",
		},
		"nov_april": {
			monthsAhead:  5,
			currentMonth: "november",
			want:         "april",
		},
		"nov_double_oct": {
			monthsAhead:  23,
			currentMonth: "november",
			want:         "october",
		},
	}

	for name, testdata := range cases {
		t.Run(name, func(subt *testing.T) {
			got := months.WhichMonth(testdata.currentMonth, testdata.monthsAhead)
			if got != testdata.want {
				subt.Errorf("want: %s, but got: %s", testdata.want, got)
			}
		})
	}
}

func TestWhichMonthTwo(t *testing.T) {
	cases := map[string]struct {
		monthsAhead  int
		currentMonth string
		want         string
	}{
		"feb_may": {
			monthsAhead:  3,
			currentMonth: "february",
			want:         "may",
		},
		"feb_jun": {
			monthsAhead:  4,
			currentMonth: "february",
			want:         "june",
		},
		"jun_dec": {
			monthsAhead:  6,
			currentMonth: "june",
			want:         "december",
		},
		"nov_jan": {
			monthsAhead:  2,
			currentMonth: "november",
			want:         "january",
		},
		"nov_feb": {
			monthsAhead:  3,
			currentMonth: "november",
			want:         "february",
		},
		"nov_april": {
			monthsAhead:  5,
			currentMonth: "november",
			want:         "april",
		},
		"nov_double_oct": {
			monthsAhead:  23,
			currentMonth: "november",
			want:         "october",
		},
	}

	for name, testdata := range cases {
		t.Run(name, func(subt *testing.T) {
			got := months.WhichMonthTwo(testdata.currentMonth, testdata.monthsAhead)
			if got != testdata.want {
				subt.Errorf("want: %s with months ahead %d, but got: %s", testdata.want, testdata.monthsAhead, got)
			}
		})
	}
}
