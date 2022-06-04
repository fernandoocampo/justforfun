package longsubstr

import "strings"

func LengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	var max int
	remainder := s
	for _, l := range s {
		if len(remainder) == 1 || len(remainder) <= max {
			break
		}
		remainder = remainder[1:]
		var newword strings.Builder
		newword.WriteRune(l)
		count := 1
		for _, nl := range remainder {
			if strings.ContainsRune(newword.String(), nl) {
				break
			}
			newword.WriteRune(nl)
			count++
		}
		if count > max {
			max = count
		}
		newword.Reset()
	}
	return max
}

func LengthOfLongestSubstringRec(s string) int {
	if s == "" {
		return 0
	}
	return getLongestWord(s, 0)
}

func getLongestWord(s string, max int) int {
	if len(s) == 1 {
		if max > len(s) {
			return max
		}
		return len(s)
	}
	if len(s) <= max {
		return max
	}
	count := countLetters(s, &strings.Builder{}, 0)
	if count > max {
		max = count
	}
	return getLongestWord(s[1:], max)
}

func countLetters(s string, newword *strings.Builder, count int) int {
	l := rune(s[0])
	if strings.ContainsRune(newword.String(), l) {
		return count
	}
	newword.WriteRune(l)
	count++
	if len(s) == 1 {
		return count
	}
	return countLetters(s[1:], newword, count)
}
