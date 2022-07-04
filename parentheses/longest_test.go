package parentheses_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/parentheses"
)

func TestLongestValidParentheses(t *testing.T) {
	cases := map[string]struct {
		s    string
		want int
	}{
		"empty": {
			s:    "",
			want: 0,
		},
		"()": {
			s:    "()",
			want: 2,
		},
		"(()": {
			s:    "(()",
			want: 2,
		},
		")()())": {
			s:    ")()())",
			want: 4,
		},
		"()(()": {
			s:    "()(()",
			want: 2,
		},
		"(()()": {
			s:    "(()()",
			want: 4,
		},
		"(()())": {
			s:    "(()())",
			want: 6,
		},
		")()())()()(": {
			s:    ")()())()()(",
			want: 4,
		},
		"(())(": {
			s:    "(())(",
			want: 4,
		},
		")(((((()())()()))()(()))(": {
			s:    ")(((((()())()()))()(()))(",
			want: 22,
		},
	}
	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := parentheses.LongestValidParentheses(data.s)
			if got != data.want {
				st.Errorf("want: %d, but got: %d", data.want, got)
			}
		})
	}
}
