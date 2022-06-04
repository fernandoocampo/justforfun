package romanint

import "strings"

var romanMap = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

var romans = []struct {
	n int
	r string
}{
	{n: 1000, r: "M"},
	{n: 900, r: "CM"},
	{n: 500, r: "D"},
	{n: 400, r: "CD"},
	{n: 100, r: "C"},
	{n: 90, r: "XC"},
	{n: 50, r: "L"},
	{n: 40, r: "XL"},
	{n: 10, r: "X"},
	{n: 9, r: "IX"},
	{n: 5, r: "V"},
	{n: 4, r: "IV"},
	{n: 1, r: "I"},
}

func IntToRoman(num int) string {
	if num < 1 {
		return ""
	}
	val, ok := romanMap[num]
	if ok {
		return val
	}
	b := strings.Builder{}
	for _, v := range romans {
		for num >= v.n {
			b.WriteString(v.r)
			num -= v.n
		}
	}
	return b.String()
}
