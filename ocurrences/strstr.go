package ocurrences

func StrStr(haystack, needle string) int {
	if needle == "" {
		return 0
	}
	needleSize := len(needle)
	if haystack == "" || len(haystack) < needleSize {
		return -1
	}
	if len(haystack) == needleSize && haystack == needle {
		return 0
	}
	i := 0
	for ; i <= len(haystack)-needleSize; i++ {
		if haystack[i:i+needleSize] == needle {
			return i
		}
	}
	return -1
}
