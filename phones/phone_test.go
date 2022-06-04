package phones_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/phones"
	"github.com/stretchr/testify/assert"
)

func TestLetterCombinations(t *testing.T) {
	cases := map[string]struct {
		digits string
		want   []string
	}{
		"23": {
			digits: "23",
			want:   []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		"empty": {
			digits: "",
			want:   nil,
		},
		"2": {
			digits: "2",
			want:   []string{"a", "b", "c"},
		},
		"234": {
			digits: "234",
			want:   []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"},
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := phones.LetterCombinations(data.digits)
			assert.Equal(t, data.want, got)
		})
	}
}
