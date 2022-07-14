package positions

func SearchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	idx := findMiddle(nums, target)
	if idx > len(nums)-1 || nums[idx] == target {
		return idx
	}
	for i := idx; nums[idx] < target; i++ {
		if nums[i] > target {
			return i
		}
		if nums[i] < target {
			return i + 1
		}
	}
	for i := idx; nums[idx] > target; i-- {
		if nums[i] < target {
			return i + 1
		}
		if nums[i] > target {
			return i
		}
	}
	return -1
}

func findMiddle(nums []int, target int) int {
	if len(nums) == 1 {
		return 0
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
			continue
		}
		if nums[mid] < target {
			left = mid + 1
			continue
		}
		return mid
	}
	return left
}
