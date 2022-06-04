package median

import "sort"

func FindMedianSortedArrays(nums1, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0.0
	}
	return calculateMedian(mergeAndOrderSlices(nums1, nums2))
}

func mergeAndOrderSlices(nums1, nums2 []int) []int {
	if len(nums1) == 0 && len(nums2) > 0 {
		sort.Ints(nums2)
		return nums2
	}
	if len(nums2) == 0 && len(nums1) > 0 {
		sort.Ints(nums1)
		return nums1
	}

	major := nums1
	minor := nums2
	if nums2[0] >= nums1[len(nums1)-1] {
		major = nums2
		minor = nums1
	}
	mergedArray := append(minor, major...)
	sort.Ints(mergedArray)
	return mergedArray
}

func FindMedianSortedArraysAlmost(nums1, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0.0
	}
	var sortit bool
	var mergedArray []int
	switch {
	case nums1[0] >= nums2[len(nums2)-1]:
		mergedArray = append(nums2, nums1...)
	case nums2[0] >= nums1[len(nums1)-1]:
		mergedArray = append(nums1, nums2...)
	default:
		sortit = true
		mergedArray = append(nums1, nums2...)
	}
	if sortit {
		sort.Ints(mergedArray)
	}
	return calculateMedian(mergedArray)
}

func FindMedianSortedArraysComplex(nums1, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0.0
	}
	var mergedArray []int
	if nums1[0] >= nums2[len(nums2)-1] {
		mergedArray = append(mergedArray, nums2...)
		mergedArray = append(mergedArray, nums1...)
		return calculateMedian(mergedArray)
	}

	if nums2[0] >= nums1[len(nums1)-1] {
		mergedArray = append(mergedArray, nums1...)
		mergedArray = append(mergedArray, nums2...)
		return calculateMedian(mergedArray)
	}
	var result []int
	if len(nums1) >= len(nums2) {
		mergeArrays(nums1, nums2, result, 0, 0)
	} else {
		mergeArrays(nums2, nums1, result, 0, 0)
	}
	return calculateMedian(result)
}

func calculateMedian(nums []int) float64 {
	if len(nums)%2 != 0 { // odd number
		idx := (len(nums) + 1) / 2
		return float64(nums[idx-1])
	}
	// even number
	idx1 := (len(nums) / 2) - 1
	idx2 := idx1 + 1
	result := float64(nums[idx1]+nums[idx2]) / 2
	return result
}

func mergeArrays(dst, src, result []int, isrc, idst int) {
	if isrc >= len(src) {
		return
	}
	var tmp []int
	if dst[idst] >= src[isrc] {
		tmp = dst[:idst]
		tmp = append(tmp, src[isrc])
		tmp = append(tmp, dst[idst+1:]...)
	}
	if dst[idst] < src[isrc] {
		tmp = dst[:idst+1]
		tmp = append(tmp, src[isrc])
		tmp = append(tmp, dst[idst+2:]...)
	}
	isrc++
	idst++
	mergeArrays(dst, src, tmp, isrc, idst)
}

func FindMedianSortedArraysLoop(nums1, nums2 []int) float64 {
	return 0.0
}

func FindAverageSortedArrays(nums1, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0.0
	}
	count := float64(len(nums1)) + float64(len(nums2))
	sum := sumArrays(nums1) + sumArrays(nums2)
	return float64(sum) / count
}

func sumArrays(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sum := nums[0] + sumArrays(nums[1:])
	return sum
}

func FindAverageSortedArraysLoop(nums1, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0.0
	}
	count := float64(len(nums1)) + float64(len(nums2))
	var sum int
	for i := 0; ; i++ {
		if i >= len(nums1) && i >= len(nums2) {
			break
		}
		if i < len(nums1) {
			sum += nums1[i]
		}
		if i < len(nums2) {
			sum += nums2[i]
		}
	}
	return float64(sum) / count
}
