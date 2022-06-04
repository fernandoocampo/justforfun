package phones

import (
	"strconv"
	"strings"
)

var buttons = []string{"", "_", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

type combiner struct {
	now    string
	result []string
}

func (c *combiner) combine(s string, i int) {
	if i == len(s) {
		c.result = append(c.result, c.now)
		return
	}
	idx := index(string(s[i]))
	for _, ch := range buttons[idx] {
		c.now = c.now + string(ch)
		c.combine(s, i+1)
		c.now = c.now[:len(c.now)-1]
	}
}

func LetterCombinations(digits string) []string {
	var result []string
	if digits == "" {
		return result
	}
	if len(digits) == 1 {
		idx := index(string(digits[0]))
		return strings.Split(buttons[idx], "")
	}
	c := &combiner{}
	c.combine(digits, 0)
	return c.result
}

func index(digit string) int {
	idx, err := strconv.Atoi(string(digit))
	if err != nil {
		panic(err)
	}
	return idx
}
