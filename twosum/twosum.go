package twosum

func TwoSum(nums []int, target int) []int {
	var i, j = 0, 1
	for ; j < len(nums); j++ {
		if nums[i]+nums[j] == target {
			return []int{i, j}
		}
		if j == len(nums)-1 {
			i++
			j = i
		}
	}
	return []int{}
}
