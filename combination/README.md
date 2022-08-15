# combination katas

## 39. Combination Sum

Medium

12222

255

Add to List

Share
Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.

It is guaranteed that the number of unique combinations that sum up to target is less than 150 combinations for the given input.

 

Example 1:

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.
Example 2:

Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]
Example 3:

Input: candidates = [2], target = 1
Output: []
 

Constraints:

1 <= candidates.length <= 30
1 <= candidates[i] <= 200
All elements of candidates are distinct.
1 <= target <= 500

### testing and benchmarking

* testing

```sh
go test -race -count 1 -timeout 2s -run ^TestCombinationSum$ github.com/fernandoocampo/justforfun/combination
```

### Report

run a scenario to understand what someone else cleverly did

```log
level: 0-, parent: none :  candidates [2 3 5] partial [] target 8 leaf 
level: 1-1, parent: 0 :  candidates [2 3 5] partial [2] target 6 leaf 1
level: 2-1, parent: 1 :  candidates [2 3 5] partial [2 2] target 4 leaf 1
level: 3-1, parent: 2 :  candidates [2 3 5] partial [2 2 2] target 2 leaf 1
level: 4-1, parent: 3 :  candidates [2 3 5] partial [2 2 2 2] target 0 leaf 1
found one [2 2 2 2]
exiting 4-1
level: 4-2, parent: 3 :  candidates [3 5] partial [2 2 2] target 2 leaf 2
exiting 4-2
exiting 3-
level: 3-2, parent: 2 :  candidates [3 5] partial [2 2] target 4 leaf 2
level: 4-1, parent: 3 :  candidates [3 5] partial [2 2 3] target 1 leaf 1
exiting 4-1
level: 4-2, parent: 3 :  candidates [5] partial [2 2] target 4 leaf 2
exiting 4-2
exiting 3-
exiting 2-
level: 2-2, parent: 1 :  candidates [3 5] partial [2] target 6 leaf 2
level: 3-1, parent: 2 :  candidates [3 5] partial [2 3] target 3 leaf 1
level: 4-1, parent: 3 :  candidates [3 5] partial [2 3 3] target 0 leaf 1
found one [2 3 3]
exiting 4-1
level: 4-2, parent: 3 :  candidates [5] partial [2 3] target 3 leaf 2
exiting 4-2
exiting 3-
level: 3-2, parent: 2 :  candidates [5] partial [2] target 6 leaf 2
level: 4-1, parent: 3 :  candidates [5] partial [2 5] target 1 leaf 1
exiting 4-1
level: 4-2, parent: 3 :  candidates [] partial [2] target 6 leaf 2
exiting 4-2
exiting 3-
exiting 2-
exiting 1-
level: 1-2, parent: 0 :  candidates [3 5] partial [] target 8 leaf 2
level: 2-1, parent: 1 :  candidates [3 5] partial [3] target 5 leaf 1
level: 3-1, parent: 2 :  candidates [3 5] partial [3 3] target 2 leaf 1
exiting 3-1
level: 3-2, parent: 2 :  candidates [5] partial [3] target 5 leaf 2
level: 4-1, parent: 3 :  candidates [5] partial [3 5] target 0 leaf 1
found one [3 5]
exiting 4-1
level: 4-2, parent: 3 :  candidates [] partial [3] target 5 leaf 2
exiting 4-2
exiting 3-
exiting 2-
level: 2-2, parent: 1 :  candidates [5] partial [] target 8 leaf 2
level: 3-1, parent: 2 :  candidates [5] partial [5] target 3 leaf 1
exiting 3-1
level: 3-2, parent: 2 :  candidates [] partial [] target 8 leaf 2
exiting 3-2
exiting 2-
exiting 1-
exiting 0-
got [[2 2 2 2] [2 3 3] [3 5]]
```