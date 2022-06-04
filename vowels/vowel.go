package vowels

import "strings"

var vowelmap = map[rune]interface{}{
	'a': nil,
	'e': nil,
	'i': nil,
	'o': nil,
	'u': nil,
}

func GetCountFor(str string) int {
	var count int
	for _, w := range str {
		if _, ok := vowelmap[w]; ok {
			count++
		}
	}
	return count
}

func GetCountRec(str string) int {
	if str == "" {
		return 0
	}
	if len(str) == 1 {
		if strings.ContainsRune("aeiou", []rune(str)[0]) {
			return 1
		}
		return 0
	}
	result := GetCountRec(string(str[0]))
	result += GetCountRec(str[1:])
	return result
}
