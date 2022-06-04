package palindromics_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/palindromics"
)

func TestLongestPalindrome(t *testing.T) {
	cases := map[string]struct {
		input string
		want  string
	}{
		"babad": {
			input: "babad",
			want:  "aba",
		},
		"cbbd": {
			input: "cbbd",
			want:  "bb",
		},
		"bb": {
			input: "bb",
			want:  "bb",
		},
		"aacabdkacaa": {
			input: "aacabdkacaa",
			want:  "aca",
		},
		"nothing": {
			input: "ac",
			want:  "c",
		},
		"one": {
			input: "a",
			want:  "a",
		},
		"abcda": {
			input: "abcda",
			want:  "a",
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := palindromics.LongestPalindrome(data.input)
			if got != data.want {
				st.Errorf("want: %s, but got: %s", data.want, got)
			}
		})
	}
}

func TestLongestPalindromeLoop(t *testing.T) {
	cases := map[string]struct {
		input string
		want  string
	}{
		"babad": {
			input: "babad",
			want:  "aba",
		},
		"cbbd": {
			input: "cbbd",
			want:  "bb",
		},
		"bb": {
			input: "bb",
			want:  "bb",
		},
		"aacabdkacaa": {
			input: "aacabdkacaa",
			want:  "aca",
		},
		"nothing": {
			input: "ac",
			want:  "",
		},
		"one": {
			input: "a",
			want:  "a",
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := palindromics.LongestPalindromeLoop(data.input)
			if got != data.want {
				st.Errorf("want: %s, but got: %s", data.want, got)
			}
		})
	}
}

func BenchmarkLongestPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		palindromics.LongestPalindrome("aacabdkacaa")
	}
}

func BenchmarkLongestPalindromeLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		palindromics.LongestPalindromeLoop("aacabdkacaa")
	}
}
