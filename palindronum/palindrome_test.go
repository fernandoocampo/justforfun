package palindronum_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/palindronum"
)

func TestIsPalindrome(t *testing.T) {
	cases := map[string]struct {
		input int
		want  bool
	}{
		"121": {
			121,
			true,
		},
		"-121": {
			-121,
			false,
		},
		"10": {
			10,
			false,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := palindronum.IsPalindrome(data.input)
			if got != data.want {
				st.Errorf("want: %t, but got: %t", data.want, got)
			}
		})
	}
}

func TestIsPalindromeLoop(t *testing.T) {
	cases := map[string]struct {
		input int
		want  bool
	}{
		"121": {
			121,
			true,
		},
		"-121": {
			-121,
			false,
		},
		"10": {
			10,
			false,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := palindronum.IsPalindromeLoop(data.input)
			if got != data.want {
				st.Errorf("want: %t, but got: %t", data.want, got)
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		palindronum.IsPalindrome(123321)
	}
}

func BenchmarkIsPalindromeLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		palindronum.IsPalindromeLoop(123321)
	}
}
