package introman_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/introman"
)

func TestRomanToInt(t *testing.T) {
	cases := map[string]struct {
		input string
		want  int
	}{
		"III":     {"III", 3},
		"LVIII":   {"LVIII", 58},
		"MCMXCIV": {"MCMXCIV", 1994},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := introman.RomanToInt(data.input)
			if got != data.want {
				st.Errorf("want: %d, but got: %d", data.want, got)
			}
		})
	}
}

func TestRomanToIntLoop(t *testing.T) {
	cases := map[string]struct {
		input string
		want  int
	}{
		"III":     {"III", 3},
		"LVIII":   {"LVIII", 58},
		"MCMXCIV": {"MCMXCIV", 1994},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := introman.RomanToIntLoop(data.input)
			if got != data.want {
				st.Errorf("want: %d, but got: %d", data.want, got)
			}
		})
	}
}

func TestRomanToIntLoop2(t *testing.T) {
	cases := map[string]struct {
		input string
		want  int
	}{
		"III":     {"III", 3},
		"LVIII":   {"LVIII", 58},
		"MCMXCIV": {"MCMXCIV", 1994},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := introman.RomanToIntLoop2(data.input)
			if got != data.want {
				st.Errorf("want: %d, but got: %d", data.want, got)
			}
		})
	}
}

func BenchmarkRomanToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		introman.RomanToInt("MCMXCIV")
	}
}

func BenchmarkRomanToIntLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		introman.RomanToIntLoop("MCMXCIV")
	}
}

func BenchmarkRomanToIntLoop2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		introman.RomanToIntLoop2("MCMXCIV")
	}
}
