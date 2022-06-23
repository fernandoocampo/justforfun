package duplicates

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
