package substrings_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/substrings"
	"github.com/stretchr/testify/assert"
)

func TestFindSubstring(t *testing.T) {
	cases := map[string]struct {
		s     string
		words []string
		want  []int
	}{
		"s_empty": {
			s:     "",
			words: []string{"foo", "bar"},
			want:  nil,
		},
		"words_empty": {
			s:     "barfoothefoobarman",
			words: nil,
			want:  nil,
		},
		"ababaab_ab_ba_ba": {
			s:     "ababaab",
			words: []string{"ab", "ba", "ba"},
			want:  []int{1},
		},
		"s_shorter_than_words": {
			s:     "barfoo",
			words: []string{"foo", "bar", "the"},
			want:  nil,
		},
		"barfoothefoobarman_foo_bar": {
			s:     "barfoothefoobarman",
			words: []string{"foo", "bar"},
			want:  []int{0, 9},
		},
		"barfoofoobarthefoobarman_bar_foo_the": {
			s:     "barfoofoobarthefoobarman",
			words: []string{"bar", "foo", "the"},
			want:  []int{6, 9, 12},
		},
		"wordgoodgoodgoodbestword": {
			s:     "wordgoodgoodgoodbestword",
			words: []string{"word", "good", "best", "word"},
			want:  nil,
		},
		"bcabbcaabbccacacbabccacaababcbb": {
			s:     "bcabbcaabbccacacbabccacaababcbb",
			words: []string{"c", "b", "a", "c", "a", "a", "a", "b", "c"},
			want:  []int{6, 16, 17, 18, 19, 20},
		},
	}
	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := substrings.FindSubstring(data.s, data.words)
			assert.Equal(st, data.want, got)
		})
	}
}
