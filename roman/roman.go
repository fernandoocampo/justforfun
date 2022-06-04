package roman

import (
	"strings"
)

type roman struct {
	value int
	roman string
}

var romansMap = map[string]int{
	"M":  1000,
	"CM": 900,
	"D":  500,
	"CD": 400,
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

var romans = []roman{
	{
		value: 1000,
		roman: "M",
	},
	{
		value: 900,
		roman: "CM",
	},
	{
		value: 500,
		roman: "D",
	},
	{
		value: 400,
		roman: "CD",
	},
	{
		value: 100,
		roman: "C",
	},
	{
		value: 90,
		roman: "XC",
	},
	{
		value: 50,
		roman: "L",
	},
	{
		value: 40,
		roman: "XL",
	},
	{
		value: 10,
		roman: "X",
	},
	{
		value: 9,
		roman: "IX",
	},
	{
		value: 5,
		roman: "V",
	},
	{
		value: 4,
		roman: "IV",
	},
	{
		value: 1,
		roman: "I",
	},
}

func GetRomanNumbers(number int) string {
	var result strings.Builder
	for _, conv := range romans {
		for number >= conv.value {
			result.WriteString(conv.roman)
			number -= conv.value
		}
	}
	return result.String()
}

func GetNumberFromRoman(roman string) int {
	return getNumberFromRoman(roman)
}

func getNumberFromRoman(roman string) int {
	if roman == "" {
		return 0
	}
	result, ok := romansMap[roman]
	if ok {
		return result
	}
	if len(roman) > 2 {
		result = getNumberFromRoman(roman[:2])
		result += getNumberFromRoman(roman[2:])
		return result
	}
	result = getNumberFromRoman(roman[:1])
	result += getNumberFromRoman(roman[1:])

	return result
}
