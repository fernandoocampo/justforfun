package substrings

import (
	"strings"
)

func FindSubstring(s string, words []string) []int {
	if cannotFind(s, words) {
		return nil
	}
	wordsize := len(words[0])
	wordssize := len(words) * wordsize
	initials := initials(words)
	var result []int
	for idx, v := range s {
		if cannotCompare(v, initials) {
			continue
		}
		if len(s[idx:]) < wordssize {
			return result
		}
		if aMatch(s[idx:idx+wordssize], words, wordsize) {
			result = append(result, idx)
		}
	}
	return result
}

func aMatch(s string, words []string, wsize int) bool {
	if len(words) == 0 {
		return len(s) == 0
	}
	fidx := -1
	for j := 0; j < len(s); j += wsize {
		if words[0] == s[j:j+wsize] {
			fidx = j
			break
		}
	}
	if fidx == -1 {
		return false
	}
	if fidx == 0 {
		s = s[wsize:]
	} else {
		s = s[:fidx] + s[fidx+wsize:]

	}
	return aMatch(s, words[1:], wsize)
}

func initials(words []string) string {
	result := make([]rune, len(words))
	for i := 0; i < len(words); i++ {
		result[i] = rune(words[i][0])
	}
	return string(result)
}

func cannotCompare(v rune, initials string) bool {
	return !strings.ContainsRune(initials, v)
}

func cannotFind(s string, words []string) bool {
	return len(words) == 0 ||
		s == "" ||
		len(s) < len(words[0])*len(words)
}
