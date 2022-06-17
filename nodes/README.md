# 19. Remove Nth Node From End of List

Given the head of a linked list, remove the nth node from the end of the list and return its head.

 

Example 1:


Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
Example 2:

Input: head = [1], n = 1
Output: []
Example 3:

Input: head = [1,2], n = 1
Output: [1]
 

Constraints:

The number of nodes in the list is sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz
 

Follow up: Could you do this in one pass?


## Testing and Benchmarking

* Testing

```sh
go test -v -count=1 -timeout 4s -run ^TestRemoveNthFromEnd$ github.com/fernandoocampo/justforfun/nodes
```

* Benchmarking

```sh
rm trace.out nodes.test mem.out cpu.out

go test -run none -bench RemoveNthFromEnd -benchtime=5s -benchmem -cpuprofile=cpu.out -memprofile=mem.out -trace=trace.out github.com/fernandoocampo/justforfun/nodes

go tool pprof nodes.test cpu.out
go tool pprof -alloc_space nodes.test mem.out
```

# 24. Swap Nodes in Pairs

Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

 

Example 1:

(1) -> (2) -> (3) -> (4)
        vvvvvv
(2) -> (1) -> (4) -> (3)


Input: head = [1,2,3,4]
Output: [2,1,4,3]
Example 2:

Input: head = []
Output: []
Example 3:

Input: head = [1]
Output: [1]
 

Constraints:

The number of nodes in the list is in the range [0, 100].
0 <= Node.val <= 100

## Testing and Benchmarking

* Testing
```sh
go test -v -count=1 -timeout 3s -run ^TestSwapPairs github.com/fernandoocampo/justforfun/nodes
go test -v -count=1 -timeout 3s -run ^TestSwapPairsInternet github.com/fernandoocampo/justforfun/nodes
```

* Benchmarking
```sh
rm cpu.out nodes.test mem.out trace.out 
go test -run none -benchtime 5s -bench SwapPairs -benchmem -cpuprofile=cpu.out -memprofile=mem.out -trace=trace.out github.com/fernandoocampo/justforfun/nodes

go test -run none -benchtime 5s -bench SwapPairsInternet -benchmem -cpuprofile=cpu.out -memprofile=mem.out -trace=trace.out github.com/fernandoocampo/justforfun/nodes

go tool pprof nodes.test cpu.out
go tool pprof -alloc_space nodes.test mem.out
```

# Reverse Nodes in k-group

Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.

You may not alter the values in the list's nodes, only nodes themselves may be changed.

k = 2
(1)->(2)->(3)->(4)->(5)
        vvvv
(2)->(1)->(4)->(3)->(5)


k = 3
(1)->(2)->(3)->(4)->(5)
        vvvv
(3)->(2)->(1)->(4)->(5)

Constraints:

The number of nodes in the list is n.
1 <= k <= n <= 5000
0 <= Node.val <= 1000
 

Follow-up: Can you solve the problem in O(1) extra memory space?

## Testing and Benchmarking

* Testing
```sh
go test -v -count=1 -timeout 3s -run ^TestReverseKGroup github.com/fernandoocampo/justforfun/nodes
go test -v -count=1 -timeout 3s -run ^TestReverseKGroupInternet github.com/fernandoocampo/justforfun/nodes
```

* Benchmarking
```sh
rm cpu.out nodes.test mem.out trace.out 
go test -run none -benchtime 5s -bench ReverseKGroup -benchmem -cpuprofile=cpu.out -memprofile=mem.out -trace=trace.out github.com/fernandoocampo/justforfun/nodes

go test -run none -benchtime 5s -bench ReverseKGroupInternet -benchmem -cpuprofile=cpu.out -memprofile=mem.out -trace=trace.out github.com/fernandoocampo/justforfun/nodes

go tool pprof nodes.test cpu.out
go tool pprof -alloc_space nodes.test mem.out
```