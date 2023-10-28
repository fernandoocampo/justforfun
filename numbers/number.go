package numbers

import (
	"sort"
)

func FindSmallestNumber(data []int) int {
	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	return data[0]
}

func SmallestIntegerFinder(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}

func SmallestIntegerFinderFor(numbers []int) int {
	smallest := numbers[0]

	for i := range numbers {
		if numbers[i] < smallest {
			smallest = numbers[i]
		}
	}

	return smallest
}

func SmallestIntegerFinderForTwo(numbers []int) int {
	smallest := numbers[0]

	for i := 0; i < len(numbers); i++ {
		if numbers[i] < smallest {
			smallest = numbers[i]
		}
	}

	return smallest
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
