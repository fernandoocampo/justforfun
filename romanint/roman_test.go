package romanint_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/romanint"
)

func TestIntToRoman(t *testing.T) {
	cases := map[string]struct {
		num  int
		want string
	}{
		"V": {
			num:  5,
			want: "V",
		},
		"III": {
			num:  3,
			want: "III",
		},
		"LVIII": {
			num:  58,
			want: "LVIII",
		},
		"MCMXCIV": {
			num:  1994,
			want: "MCMXCIV",
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := romanint.IntToRoman(data.num)
			if got != data.want {
				st.Errorf("num: %d, want: %s, but got: %s", data.num, data.want, got)
			}
		})
	}
}
