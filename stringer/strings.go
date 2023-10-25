package stringer

import (
	"strings"
	"unicode"
)

func IsUpperCase(value string) bool {
	for _, v := range value {
		println(v)
		if v != ' ' && !unicode.IsUpper(v) {
			return false
		}
	}
	return true
}

func IsUpperCaseSmart(value string) bool {
	return value == strings.ToUpper(value)
}
