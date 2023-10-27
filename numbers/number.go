package numbers

import (
	"sort"
)

func FindSmallestNumber(data []int) int {
	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	return data[0]
}
