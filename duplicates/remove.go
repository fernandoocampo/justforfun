package duplicates

import (
	"slices"
)

func RemoveElement(nums []int, val int) int {
	if len(nums) == 1 && nums[0] != val {
		return 1
	}
	if len(nums) == 1 && nums[0] == val {
		return 0
	}
	left, right, size := 0, 0, len(nums)
	for right < size && left < size {
		if nums[left] != val {
			left, right = left+1, left
			continue
		}
		if right++; right >= size {
			break
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	return left
}

func Solve(data []int) []int {
	if len(data) < 2 {
		return data
	}

	var result []int
	for i := len(data) - 1; i >= 0; i-- {
		if slices.Contains(result, data[i]) {
			continue
		}
		for j := 0; j < len(data); j++ {
			if data[i] == data[j] && !slices.Contains(result, data[i]) {
				result = append([]int{data[i]}, result...)
				break
			}
		}
	}
	return result
}

func SolveSmart(arr []int) (res []int) {
	visited := map[int]bool{}
	for i := len(arr) - 1; i >= 0; i-- {
		n := arr[i]
		if visited[n] {
			continue
		}

		visited[n] = true
		res = append([]int{n}, res...)
	}

	return
}

func SolveNoSlices(data []int) []int {
	if len(data) < 2 {
		return data
	}

	var result []int
	for i := len(data) - 1; i >= 0; i-- {
		if contains(result, data[i]) {
			continue
		}
		for j := 0; j < len(data); j++ {
			if data[i] == data[j] && !contains(result, data[i]) {
				result = append([]int{data[i]}, result...)
				break
			}
		}
	}
	return result
	// https://github.com/ZaytsevNS/python_codewars/blob/main/7KYU/solve_remove_duplicates.py
}

func contains(values []int, value int) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}

	return false
}
