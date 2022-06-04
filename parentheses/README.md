# Valid Parentheses

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
 

Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false
 

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.

## Testing and Benchmarking

* testing

```sh
go test -count 1 -timeout 2s -run ^TestIsValid$ github.com/fernandoocampo/justforfun/parentheses
go test -count 1 -timeout 2s -run ^TestIsValid github.com/fernandoocampo/justforfun/parentheses
```

* benchmarking

```sh
go test -run none -bench IsValid -benchtime=5s -benchmem -cpuprofile=cpu.out -memprofile=mem.out -trace=trace.out github.com/fernandoocampo/justforfun/parentheses

go test -run none -bench IsValidLoop -benchtime=5s -benchmem -cpuprofile=cpu.out -memprofile=mem.out -trace=trace.out github.com/fernandoocampo/justforfun/parentheses

go tool pprof parentheses.test cpu.out
go tool pprof -alloc_space parentheses.test mem.out

rm cpu.out mem.out trace.out parentheses.test
```
