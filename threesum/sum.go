package threesum

import (
	"fmt"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	// triplets(nums)
	// return nil
	sort.Ints(nums)
	return iterateNums(nums, 0, make([][]int, 0))
	// return iterateNumsa(nums, make([][]int, 0))
	// return threeSumI(nums)
}

func iterateNumsa(nums []int, r [][]int) [][]int {
	if len(nums) < 3 {
		return r
	}
	r = collectTriplets([]int{nums[0], nums[1]}, nums[1:], r)
	return iterateNumsa(nums[1:], r)
}

func iterateNums(nums []int, i int, r [][]int) [][]int {
	if len(nums) < 3 {
		return r
	}
	if i == len(nums)-1 {
		return iterateNums(nums[1:], 0, r)
	}
	r = collectTriplets([]int{nums[0], nums[i+1]}, nums[i+2:], r)
	fmt.Println(r)
	return iterateNums(nums, i+1, r)
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

func collectTriplets(base, nums []int, store [][]int) [][]int {
	if len(nums) == 0 {
		return store
	}
	if base[0]+base[1]+nums[0] != 0 {
		return collectTriplets(base, nums[1:], store)
	}
	item := []int{base[0], base[1], nums[0]}
	sort.Ints(item)
	return collectTriplets(
		base,
		nums[1:],
		append(store, item),
	)
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
