package prefixes_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/prefixes"
)

func TestLongestCommonPrefix(t *testing.T) {
	cases := map[string]struct {
		strs []string
		want string
	}{
		`["flower","flow","flight"]`: {[]string{"flower", "flow", "flight"}, "fl"},
		`["dog","racecar","car"]`:    {[]string{"dog", "racecar", "car"}, ""},
		`["a"]`:                      {[]string{"a"}, "a"},
		`["ab", "a"]`:                {[]string{"ab", "a"}, "a"},
	}
	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := prefixes.LongestCommonPrefix(data.strs)
			if got != data.want {
				st.Errorf("want: %q, but got: %q", data.want, got)
			}
		})
	}
}

func TestLongestCommonPrefixLoop(t *testing.T) {
	cases := map[string]struct {
		strs []string
		want string
	}{
		`["flower","flow","flight"]`:   {[]string{"flower", "flow", "flight"}, "fl"},
		`["reflower","flow","flight"]`: {[]string{"reflower", "flow", "flight"}, ""},
		`["dog","racecar","car"]`:      {[]string{"dog", "racecar", "car"}, ""},
		`["a"]`:                        {[]string{"a"}, "a"},
		`["ab","a"]`:                   {[]string{"ab", "a"}, "a"},
	}
	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := prefixes.LongestCommonPrefixLoop(data.strs)
			if got != data.want {
				st.Errorf("want: %q, but got: %q", data.want, got)
			}
		})
	}
}

func BenchmarkLongestCommonPrefixLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prefixes.LongestCommonPrefixLoop([]string{"flower", "flow", "flight"})
	}
}

func BenchmarkLongestCommonPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prefixes.LongestCommonPrefix([]string{"flower", "flow", "flight"})
	}
}
