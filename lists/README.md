# Merge Two Sorted Lists

You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

 

Example 1:


Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
Example 2:

Input: list1 = [], list2 = []
Output: []
Example 3:

Input: list1 = [], list2 = [0]
Output: [0]
 

Constraints:

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.

## Testing and Benchmarking

* test

```sh
go test -timeout 3s -count 1 -run ^TestMergeTwoLists$ github.com/fernandoocampo/justforfun/lists
```

* bench

```sh
rm cpu.out mem.out trace.out lists.test

go test -run none -bench MergetTwoLists -benchtime=5s -benchmem -cpuprofile=cpu.out -memprofile=mem.out -trace=trace.out github.com/fernandoocampo/justforfun/lists

go tool pprof lists.test cpu.out
go tool pprof -alloc_space lists.test mem.out
```


# Merge k Sorted Lists

You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.

 

Example 1:

Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
Example 2:

Input: lists = []
Output: []
Example 3:

Input: lists = [[]]
Output: []
 

Constraints:

k == lists.length
0 <= k <= 104
0 <= lists[i].length <= 500
-104 <= lists[i][j] <= 104
lists[i] is sorted in ascending order.
The sum of lists[i].length will not exceed 10^4.

## Testing and Benchmarking

* test

```sh
go test -timeout 3s -count 1 -run ^TestMergeKLists$ github.com/fernandoocampo/justforfun/lists
```

* bench

```sh
rm cpu.out mem.out trace.out lists.test

go test -run none -bench MergetKLists -benchtime=5s -benchmem -cpuprofile=cpu.out -memprofile=mem.out -trace=trace.out github.com/fernandoocampo/justforfun/lists

go tool pprof lists.test cpu.out
go tool pprof -alloc_space lists.test mem.out
```