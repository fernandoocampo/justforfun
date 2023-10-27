package stringer

import (
	"sort"
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

func SortByLength(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	}

	sort.Slice(arr, func(i, j int) bool { return len(arr[i]) < len(arr[j]) })

	return arr
}
