package reversers_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/reversers"
)

func TestReverse(t *testing.T) {
	cases := map[string]struct {
		input int
		want  int
	}{
		"123": {
			input: 123,
			want:  321,
		},
		"-123": {
			input: -123,
			want:  -321,
		},
		"120": {
			input: 120,
			want:  21,
		},
		"1534236469": {
			input: 1534236469,
			want:  9646324351,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := reversers.Reverse(data.input)
			if data.want != got {
				st.Errorf("want: %d, but got: %d", data.want, got)
			}
		})
	}
}

func TestReverseLoop(t *testing.T) {
	cases := map[string]struct {
		input int
		want  int
	}{
		"123": {
			input: 123,
			want:  321,
		},
		"-123": {
			input: -123,
			want:  -321,
		},
		"120": {
			input: 120,
			want:  21,
		},
		"1534236469": {
			input: 1534236469,
			want:  9646324351,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := reversers.ReverseLoop(data.input)
			if data.want != got {
				st.Errorf("want: %d, but got: %d", data.want, got)
			}
		})
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reversers.Reverse(-123)
	}
}

func BenchmarkReverseLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reversers.ReverseLoop(-123)
	}
}
