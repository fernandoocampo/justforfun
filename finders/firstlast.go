package finders

import "fmt"

func SearchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	idx := findSearchPoint(nums, target)
	fmt.Println("idx", idx)
	if idx == -1 {
		return result
	}
	for i := idx; ; i-- {
		if i == 0 || nums[i-1] < target {
			result[0] = i
			break
		}
	}

	for i := idx; ; i++ {
		if i == len(nums)-1 || nums[i+1] > target {
			result[1] = i
			break
		}
	}
	return result
}

func findSearchPoint(nums []int, target int) int {
	idx, left, right := -1, 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
	}
	return idx
}

func SearchRangeOn(nums []int, target int) []int {
	result := []int{-1, -1}
	for i := 0; i < len(nums); i++ {
		if nums[i] == target && result[0] == -1 {
			result[0] = i
		}
		if nums[i] == target {
			result[1] = i
		}

	}
	return result
}
