# 4SUM

https://leetcode.com/problems/4sum/

18. 4Sum
Medium

6395

715

Add to List

Share
Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:

0 <= a, b, c, d < n
a, b, c, and d are distinct.
nums[a] + nums[b] + nums[c] + nums[d] == target
You may return the answer in any order.

 

Example 1:

Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
Example 2:

Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]
 

Constraints:

1 <= nums.length <= 200
-109 <= nums[i] <= 109
-109 <= target <= 109

## Testing and Benchmarking

* testing

```sh
go test -count=1 -timeout 2s -run ^TestFourSum$ github.com/fernandoocampo/justforfun/foursum
```

* benchmarking

```sh
go test -run none -benchtime 5s -bench FourSum -benchmem github.com/fernandoocampo/justforfun/foursum
go test -run none -benchtime 5s -bench FourSum -benchmem github.com/fernandoocampo/justforfun/foursum -memprofile mem.out

go test -run none -bench FourSum -benchtime=5s -benchmem -cpuprofile=cpu.out -memprofile=mem.out -trace=trace.out github.com/fernandoocampo/justforfun/foursum
g

go tool pprof foursum.test cpu.out
go tool pprof -alloc_space foursum.test mem.out

go tool trace trace.out

go tool pprof -http :8080 cpu.out
go tool pprof -http :8081 mem.out
```