package countsum

func CountPositivesSumNegatives(numbers []int) []int {
	res := []int{0, 0}
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 0 {
			res[0]++
			continue
		}
		res[1] += numbers[i]
	}
	return res // your code here
}
