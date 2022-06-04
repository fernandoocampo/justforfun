package balance_test

import (
	"errors"
	"testing"

	"github.com/fernandoocampo/justforfun/balance"
)

func TestValidate(t *testing.T) {
	cases := map[string]struct {
		input string
		want  error
	}{
		"(()(()))": {
			input: "(()(()))",
			want:  nil,
		},
		"(())(()": {
			input: "(())(()",
			want:  balance.ErrInvalid,
		},
		"()": {
			input: "()",
			want:  nil,
		},
		")(": {
			input: ")(",
			want:  balance.ErrInvalid,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := balance.Validate(data.input)
			if !errors.Is(got, data.want) {
				st.Errorf("want: %v, but got: %v", data.want, got)
			}
		})
	}
}

func TestValidateRec(t *testing.T) {
	cases := map[string]struct {
		input string
		want  error
	}{
		"(()(()))": {
			input: "(()(()))",
			want:  nil,
		},
		"(())(()": {
			input: "(())(()",
			want:  balance.ErrInvalid,
		},
		"()": {
			input: "()",
			want:  nil,
		},
		")(": {
			input: ")(",
			want:  balance.ErrInvalid,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := balance.ValidateRec(data.input)
			if !errors.Is(got, data.want) {
				st.Errorf("want: %v, but got: %v", data.want, got)
			}
		})
	}
}
