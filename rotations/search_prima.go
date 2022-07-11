package rotations

func SearchI(nums []int, target int) int {
	return 0
}

func findPivotIndex(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[right] < nums[mid] {
			left = mid + 1
			continue
		}
		right = mid
	}
	return left
}
