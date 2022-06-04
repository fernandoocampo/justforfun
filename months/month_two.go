package months

import (
	"log"
)

var monthsTwo = []string{"january", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december"}

func WhichMonthTwo(current string, monthsAhead int) string {
	result := searchMonth(current) + monthsAhead
	log.Println("current", current, "monthAhead", monthsAhead, "result", result)
	remainder := result % len(monthsTwo)
	return monthsTwo[remainder]
}

func searchMonth(current string) int {
	for idx, v := range monthsTwo {
		if v == current {
			return idx
		}
	}
	return -1
}
