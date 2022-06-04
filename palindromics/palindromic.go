package palindromics

func LongestPalindrome(s string) string {
	if isPalindromic(s) {
		return s
	}
	result := iterateWord(s, "")
	return result
}

func iterateWord(s string, longest string) string {
	if len(s) == 1 || len(s) <= len(longest) {
		return longest
	}
	result := findLongestPalindrome(s, "")
	if len(result) > len(longest) {
		longest = result
	}
	return iterateWord(s[:len(s)-1], longest)
}

func findLongestPalindrome(s, longest string) string {
	if s == "" || len(s) <= len(longest) {
		return longest
	}
	if isPalindromic(s) && len(s) > len(longest) {
		longest = s
	}
	return findLongestPalindrome(s[1:], longest)
}

func isPalindromic(s string) bool {
	if s == "" {
		return false
	}
	if len(s) == 1 {
		return true
	}
	if len(s) == 3 || len(s) == 2 {
		return s[0] == s[len(s)-1]
	}
	return (s[0] == s[len(s)-1]) && isPalindromic(s[1:len(s)-1])
}

func LongestPalindromeLoop(s string) string {
	if isPalindromic(s) {
		return s
	}
	var longest string
	for i, j := 0, 0; j < len(s); i++ {
		if i == len(s)-1 || i >= len(s)-j {
			j++
			i = -1
			continue
		}
		nw := s[i : len(s)-j]
		if len(nw) < len(longest) || !isPalindromic(nw) {
			continue
		}
		if len(nw) > len(longest) {
			longest = nw
		}
	}
	if len(s) > 1 && len(longest) == 1 {
		return ""
	}
	return longest
}
