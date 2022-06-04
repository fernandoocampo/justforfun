package foursum

import (
	"sort"
	"strconv"
)

type data struct {
	a, b, c, d int
}

type holder struct {
	now    []int
	result [][]int
	values map[string]interface{}
}

func FourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}
	if equalValues(nums) {
		if nums[0]*4 == target {
			return [][]int{{nums[0], nums[0], nums[0], nums[0]}}
		}
		return [][]int{}
	}
	sort.Ints(nums)
	h := &holder{
		values: make(map[string]interface{}),
	}
	h.listQuadruplets(nums, target, 0)
	return h.result
}

func (h *holder) listQuadruplets(nums []int, target, count int) {
	if len(h.now) == 4 {
		if isEqualToTarget(h.now, target) && h.notExist() {
			h.result = append(h.result, h.now)
		}
		return
	}

	for i := count; i < len(nums); i++ {
		h.now = append(h.now, nums[i])
		h.listQuadruplets(nums, target, count+1)
		dst := make([]int, len(h.now)-1)
		copy(dst, h.now[:len(h.now)-1])
		h.now = dst
		count++
	}
}

func isEqualToTarget(nums []int, target int) bool {
	return sum(nums) == target
}

func sum(nums []int) int {
	var result int
	for _, v := range nums {
		result += v
	}
	return result
}

func equalValues(nums []int) bool {
	equal := true
	value := nums[0]

	for _, v := range nums {
		if value != v {
			equal = false
			break
		}
	}
	return equal
}

func (h *holder) notExist() bool {
	key := strconv.Itoa(h.now[0]) + strconv.Itoa(h.now[1]) + strconv.Itoa(h.now[2]) + strconv.Itoa(h.now[3])
	_, exist := h.values[key]
	if exist {
		return false
	}
	h.values[key] = true
	return true
}
