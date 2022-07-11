package rotations

func Search(nums []int, target int) int {
	rotated := getIndexMinValue(nums)
	if nums[rotated] == target {
		return rotated
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		rotatedMid := (rotated + mid) % len(nums)
		switch {
		case nums[rotatedMid] < target:
			left = mid + 1
		case target < nums[rotatedMid]:
			right = mid - 1
		default:
			return rotatedMid
		}
	}
	return -1
}

func getIndexMinValue(nums []int) int {
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

// O(N)
func SearchOn(nums []int, target int) int {
	result := -1
	for idx, v := range nums {
		if v == target {
			return idx
		}
	}
	return result
}
