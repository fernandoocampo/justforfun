package months

import "strings"

var months = []string{
	"january", "february", "march", "april", "may", "june",
	"july", "august", "september", "october", "november",
	"december"}

func WhichMonth(currentMonth string, monthsAhead int) string {
	currentMonthN := getMonthPosition(currentMonth)
	if currentMonthN == -1 {
		return ""
	}
	if (currentMonthN + monthsAhead) < 12 {
		return months[currentMonthN+monthsAhead]
	}

	pos := (currentMonthN + monthsAhead) % 12
	return months[pos]
}

func getMonthPosition(monthName string) int {
	for i := 0; i < len(months); i++ {
		if strings.EqualFold(months[i], monthName) {
			return i
		}
	}
	return -1
}
