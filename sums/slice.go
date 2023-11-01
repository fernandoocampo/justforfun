package sums

func Between(a, b int) []int {
	var result []int
	for i := a; i <= b; i++ {
		result = append(result, i)
	}
	return result
}

func SquareSum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	if len(numbers) == 1 {
		return numbers[0] * numbers[0]
	}

	return numbers[0]*numbers[0] + SquareSum(numbers[1:])
}
