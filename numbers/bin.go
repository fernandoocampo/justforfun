package numbers

func FakeBin(x string) string {
	if x == "" {
		return x
	}

	result := make([]rune, len(x))

	for idx, v := range x {
		if (v - '0') < 5 {
			result[idx] = '0'
			continue
		}

		result[idx] = '1'
	}

	return string(result)
}

func FakeBinSmart(x string) string {
	var result string
	for _, char := range x {
		if char < '5' {
			result += "0"
		} else {
			result += "1"
		}
	}
	return result
}
