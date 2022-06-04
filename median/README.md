# 4. Median of Two Sorted Arrays
Hard

16503

2009

Add to List

Share
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).

 

Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
 

Constraints:

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106

## Tests and BenchMark

* test

```sh
go test -race -v -count=1 -timeout 5s -run ^TestFindMedianSortedArrays$ github.com/fernandoocampo/justforfun/median
go test -race -v -count=1 -timeout 5s -run ^TestFindMedianSortedArraysLoop$ github.com/fernandoocampo/justforfun/median
```

* bench

```sh
go test ./median -run none -bench FindMedianSortedArrays$ -benchtime 5s -benchmem
go test ./median -run none -bench FindMedianSortedArraysLoop$ -benchtime 5s -benchmem

go test ./median -run none -bench FindMedianSortedArrays$ -benchtime 5s -benchmem -memprofile mem.out
go tool pprof -alloc_space median.test mem.out

go test ./median -run none -bench FindMedianSortedArraysLoop$ -benchtime 5s -benchmem -memprofile mem.out
go tool pprof -alloc_space median.test mem.out
```