package countandsay_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/countandsay"
)

func TestCountAndSay(t *testing.T) {
	cases := map[string]struct {
		n    int
		want string
	}{
		"1": {
			n:    1,
			want: "1",
		},
		"2": {
			n:    2,
			want: "11",
		},
		"3": {
			n:    3,
			want: "21",
		},
		"4": {
			n:    4,
			want: "1211",
		},
		"5": {
			n:    5,
			want: "111221",
		},
		"6": {
			n:    6,
			want: "312211",
		},
	}
	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := countandsay.CountAndSay(data.n)
			if got != data.want {
				st.Errorf("want: %s, but got: %s", data.want, got)
			}
		})
	}
}
