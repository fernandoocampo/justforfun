package stringer

import (
	"log"
	"sort"
	"strings"
	"unicode"
)

var scrabbleScore = map[rune]int{
	'a': 1, 'b': 3, 'c': 3, 'd': 2, 'e': 1,
	'f': 4, 'g': 2, 'h': 4, 'i': 1, 'j': 8,
	'k': 5, 'l': 1, 'm': 3, 'n': 1, 'o': 1,
	'p': 3, 'q': 10, 'r': 1, 's': 1, 't': 1,
	'u': 1, 'v': 4, 'w': 4, 'x': 8, 'y': 4,
	'z': 10,
}

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

func ScrabbleScore(st string) int {
	input := strings.ToLower(st)
	var result int

	for _, v := range input {
		result += scrabbleScore[v]
	}

	return result
}

func EndsWith(str, ending string) bool {
	if ending == str {
		return true
	}

	if len(ending) > len(str) {
		return false
	}

	log.Println("str", str[len(str)-len(ending):])

	return str[len(str)-len(ending):] == ending
}

func EndsWithBrief(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}

func FlickSwitch(list []string) []bool {
	const flick = "flick"
	result := make([]bool, len(list))
	flag := true
	for i := 0; i < len(list); i++ {
		if list[i] != flick {
			result[i] = flag
			continue
		}

		flag = !flag
		result[i] = flag
	}
	return result
}
