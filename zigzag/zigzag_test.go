package zigzag_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/zigzag"
)

func TestConvert(t *testing.T) {
	cases := map[string]struct {
		s       string
		numRows int
		want    string
	}{
		"PAYPALISHIRING_3": {
			"PAYPALISHIRING",
			3,
			"PAHNAPLSIIGYIR",
		},
		"PAYPALISHIRING_4": {
			"PAYPALISHIRING",
			4,
			"PINALSIGYAHRPI",
		},
		"A_1": {
			"A",
			1,
			"A",
		},
		"PAYPALISHIRING_2": {
			"PAYPALISHIRING",
			2,
			"PYAIHRNAPLSIIG",
		},
		"ABCDEFGHIJK_3": {
			"ABCDEFGHIJK",
			3,
			"AEIBDFHJCGK",
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := zigzag.Convert(data.s, data.numRows)
			if got != data.want {
				st.Errorf("want: %q, but got: %q", data.want, got)
			}

		})
	}
}

func TestConvertLoop(t *testing.T) {
	cases := map[string]struct {
		s       string
		numRows int
		want    string
	}{
		"PAYPALISHIRING_3": {
			"PAYPALISHIRING",
			3,
			"PAHNAPLSIIGYIR",
		},
		"PAYPALISHIRING_4": {
			"PAYPALISHIRING",
			4,
			"PINALSIGYAHRPI",
		},
		"A_1": {
			"A",
			1,
			"A",
		},
		"PAYPALISHIRING_2": {
			"PAYPALISHIRING",
			2,
			"PYAIHRNAPLSIIG",
		},
		"ABCDEFGHIJK_3": {
			"ABCDEFGHIJK",
			3,
			"AEIBDFHJCGK",
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := zigzag.ConvertLoop(data.s, data.numRows)
			if got != data.want {
				st.Errorf("want: %q, but got: %q", data.want, got)
			}

		})
	}
}

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zigzag.Convert("PAYPALISHIRING", 4)
	}
}

func BenchmarkConvertLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zigzag.ConvertLoop("PAYPALISHIRING", 4)
	}
}
