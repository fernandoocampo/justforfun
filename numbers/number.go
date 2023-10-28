package numbers

import (
	"sort"
)

func FindSmallestNumber(data []int) int {
	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	return data[0]
}

func MakeFibos(numberOfFibos int) []int {
	result := make([]int, numberOfFibos)
	lastNumber := 1
	beforeLastOne := 0
	result[1] = lastNumber
	for i := 2; i < numberOfFibos; i++ {
		result[i] = lastNumber + beforeLastOne
		beforeLastOne = lastNumber
		lastNumber = result[i]
	}

	return result
}
