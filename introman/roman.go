package introman

var mapa = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func RomanToInt(s string) int {
	return transform(s, 0)
}

func transform(s string, sum int) int {
	if len(s) == 0 {
		return sum
	}
	if len(s) >= 2 {
		val, ok := mapa[s[len(s)-2:]]
		if ok {
			return transform(s[:len(s)-2], sum+val)
		}
	}
	if val, ok := mapa[string(s[len(s)-1])]; ok {
		return transform(s[:len(s)-1], sum+val)
	}
	return transform(s[:len(s)-1], sum)
}

func RomanToIntLoop(s string) int {
	var sum int
	for i := len(s) - 1; i >= 0; i-- {
		if i >= 2 {
			val, ok := mapa[s[i-1:i+1]]
			if ok {
				sum += val
				i--
				continue
			}
		}
		if val, ok := mapa[string(s[i])]; ok {
			sum += val
		}
	}
	return sum
}

func RomanToIntLoop2(s string) int {
	var sum int
	var former int
	for i := len(s) - 1; i >= 0; i-- {
		val := mapa[string(s[i])]
		if former > val {
			sum -= val
		} else {
			sum += val
		}
		former = val
	}
	return sum
}
