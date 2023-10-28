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

func GetMiddle(s string) string {
	if len(s) < 2 {
		return s
	}

	if len(s)%2 != 0 {
		return string(s[(len(s))/2])
	}

	return s[(len(s)/2)-1 : (len(s)/2)+1]
}

func GetMiddleTwo(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}
	return s[(n-1)/2 : n/2+1]
}
