package prefixes

func LongestCommonPrefix(strs []string) string {
	return searchCommonPrefix(strs, make(map[string]int))
}

func searchCommonPrefix(strs []string, data map[string]int) string {
	if len(strs) == 0 {
		return getLongestWord(data, len(strs))
	}
	indexWord(strs[0], 2, data)
	return searchCommonPrefix(strs[1:], data)
}

func getLongestWord(data map[string]int, maxlen int) string {
	var result string
	var max = maxlen
	for k, v := range data {
		if v > 1 && v == max && len(k) > len(result) {
			result = k
		}
	}
	return result
}

func indexWord(w string, idx int, data map[string]int) {
	if idx >= len(w) {
		return
	}
	_, ok := data[w[:idx]]
	if !ok {
		data[w[:idx]] = 1
	} else {
		data[w[:idx]]++
	}
	indexWord(w, idx+1, data)
}

func LongestCommonPrefixLoop(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	data := make(map[string]int)
	for _, w := range strs {
		indexWordLoop(data, w)
	}
	return getLongestWord(data, len(strs))
}

func indexWordLoop(data map[string]int, w string) {
	for i := 1; i <= len(w); i++ {
		_, ok := data[w[:i]]
		if !ok {
			data[w[:i]] = 1
			continue
		}
		data[w[:i]]++
	}
}
