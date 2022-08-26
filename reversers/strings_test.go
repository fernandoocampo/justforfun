package reversers_test

import (
	"context"
	"testing"
	"time"

	"github.com/fernandoocampo/justforfun/reversers"
)

func TestReverseString(t *testing.T) {
	cases := map[string]struct {
		input string
		want  string
	}{
		"anitalavalatina": {
			input: "anitalavalatina",
			want:  "anitalavalatina",
		},
		"any text": {
			input: "any text",
			want:  "txet yna",
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
			defer cancel()
			got := reversers.ReverseString(ctx, data.input)
			if got != data.want {
				t.Errorf("want: %s, but got: %s", data.want, got)
			}
		})
	}
}

func TestReverseNormalString(t *testing.T) {
	cases := map[string]struct {
		input string
		want  string
	}{
		"anitalavalatina": {
			input: "anitalavalatina",
			want:  "anitalavalatina",
		},
		"any text": {
			input: "any text",
			want:  "txet yna",
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := reversers.ReverseNormalString(data.input)
			if got != data.want {
				t.Errorf("want: %s, but got: %s", data.want, got)
			}
		})
	}
}
