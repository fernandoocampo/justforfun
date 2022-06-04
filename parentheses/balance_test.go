package parentheses_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/parentheses"
)

func TestIsValid(t *testing.T) {
	cases := map[string]struct {
		input string
		want  bool
	}{
		"()": {
			input: "()",
			want:  true,
		},
		"(){}[]": {
			input: "(){}[]",
			want:  true,
		},
		"(]": {
			input: "(]",
			want:  false,
		},
		"(){}}{": {
			input: "(){}}{",
			want:  false,
		},
		"[])": {
			input: "[])",
			want:  false,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := parentheses.IsValid(data.input)
			if got != data.want {
				st.Errorf("want: %t, but got: %t", data.want, got)
			}
		})
	}
}

func TestIsValidLoop(t *testing.T) {
	cases := map[string]struct {
		input string
		want  bool
	}{
		"()": {
			input: "()",
			want:  true,
		},
		"(){}[]": {
			input: "(){}[]",
			want:  true,
		},
		"(]": {
			input: "(]",
			want:  false,
		},
		"(){}}{": {
			input: "(){}}{",
			want:  false,
		},
		"[])": {
			input: "[])",
			want:  false,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := parentheses.IsValidLoop(data.input)
			if got != data.want {
				st.Errorf("want: %t, but got: %t", data.want, got)
			}
		})
	}
}

func BenchmarkIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parentheses.IsValid("(){}}{")
	}
}

func BenchmarkIsValidLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parentheses.IsValidLoop("(){}}{")
	}
}
