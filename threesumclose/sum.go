package threesumclose

import (
	"fmt"
	"math"
	"sort"
)

func ThreeSum(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	min := nums[0] + nums[1] + nums[2]
	return iterateNums(nums, 0, target, min)
}

func iterateNums(nums []int, i, target, min int) int {
	if len(nums) < 3 {
		return min
	}
	if i == len(nums)-1 {
		return iterateNums(nums[1:], 0, target, min)
	}
	min = collectTriplets([]int{nums[0], nums[i+1]}, nums[i+2:], target, min)
	return iterateNums(nums, i+1, target, min)
}

func collectTriplets(base, nums []int, target, min int) int {
	if len(nums) == 0 {
		return min
	}
	sum := base[0] + base[1] + nums[0]
	if sum == target {
		return target
	}
	if isClose(target, sum, min) {
		return collectTriplets(base, nums[1:], target, sum)
	}
	return collectTriplets(
		base,
		nums[1:],
		target,
		min,
	)
}

func isClose(target, sum, min int) bool {
	return abs(sum-target) < abs(min-target)
}

func abs(v int) int {
	return int(math.Abs(float64(v)))
}

func threeSumI(nums []int) [][]int {
	n := len(nums)

	// Hashmap to store the triplets
	// Helps avoid duplicate triplets
	tripletMap := make(map[[3]int][]int)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				triplet := [3]int{nums[i], nums[j], nums[k]}
				sort.Ints(triplet[:])

				if nums[i]+nums[j]+nums[k] == 0 {
					tripletMap[triplet] = []int{nums[i], nums[j], nums[k]}
				}
			}
		}
	}

	// Iterate through the hashmap and add the triplets to the result array
	var result [][]int
	for _, triplet := range tripletMap {
		result = append(result, triplet)
	}
	return result
}

func triplets(nums []int) {
	n := len(nums)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				triplet := [3]int{nums[i], nums[j], nums[k]}
				sort.Ints(triplet[:])
				fmt.Println([]int{nums[i], nums[j], nums[k]})

				// if nums[i]+nums[j]+nums[k] == 0 {
				// 	tripletMap[triplet] = []int{nums[i], nums[j], nums[k]}
				// }
			}
		}
	}
}
