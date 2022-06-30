package permutations

import (
	"sort"
)

func NextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	i := findLastIltIPlus1(nums)
	if i == -1 {
		sort.Ints(nums)
		return
	}
	if i+1 == len(nums)-1 {
		nums[i], nums[i+1] = nums[i+1], nums[i]
		return
	}
	j := findLastValueGTi(i, nums)
	nums[i], nums[j] = nums[j], nums[i]
	sort.Ints(nums[i+1:])
}

func findLastIltIPlus1(nums []int) int {
	idx := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			idx = i
		}
	}
	return idx
}

func findLastValueGTi(i int, nums []int) int {
	var idx int
	for j := i + 1; j < len(nums); j++ {
		if nums[j] > nums[i] {
			idx = j
		}
	}
	return idx
}
