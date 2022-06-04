package palindronum

import (
	"strconv"
	"strings"
)

func IsPalindrome(x int) bool {
	sx := strconv.Itoa(x)
	xs := reverse(sx, &strings.Builder{})

	return sx == xs
}

func reverse(s string, b *strings.Builder) string {
	b.WriteByte(s[len(s)-1])
	if len(s) == 1 {
		return b.String()
	}
	return reverse(s[:len(s)-1], b)
}

func IsPalindromeLoop(x int) bool {
	sx := strconv.Itoa(x)
	for i, j := len(sx)-1, 0; j < len(sx); i, j = i-1, j+1 {
		if sx[i] != sx[j] {
			return false
		}
	}
	return true
}
