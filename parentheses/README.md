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

# 22. Generate Parentheses
Medium
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
Example 2:

Input: n = 1
Output: ["()"]

Constraints:

1 <= n <= 8

## Testing and Benchmarking

* testing

```sh
go test -count 1 -timeout 3s -run ^TestGenerateParenthesis$ github.com/fernandoocampo/justforfun/parentheses
```

* failing

output
```log
["((((()))))","(((()())))","(((())()))","(((()))())","(((())))()","((()(())))","((()()()))","((()())())","((()()))()","((())(()))","((())()())","((())())()","((()))(())","((()))()()","(()((())))","(()(()()))","(()(())())","(()(()))()","(()()(()))","(()()()())","(()()())()","(()())(())","(()())()()","(())((()))","(())(()())","(())(())()","()(((())))","()((()()))","()((())())","()((()))()","()(()(()))","()(()()())","()(()())()","()(())(())","()(())()()","()()((()))","()()(()())","()()(())()","()()()(())","()()()()()"]
```

want
```log
["((((()))))","(((()())))","(((())()))","(((()))())","(((())))()","((()(())))","((()()()))","((()())())","((()()))()","((())(()))","((())()())","((())())()","((()))(())","((()))()()","(()((())))","(()(()()))","(()(())())","(()(()))()","(()()(()))","(()()()())","(()()())()","(()())(())","(()())()()","(())((()))","(())(()())","(())(())()","(())()(())","(())()()()","()(((())))","()((()()))","()((())())","()((()))()","()(()(()))","()(()()())","()(()())()","()(())(())","()(())()()","()()((()))","()()(()())","()()(())()","()()()(())","()()()()()"]
```

"(())(())()"
"(())()(())"
"(())()()()"
"()(((())))"

```log
2022/06/09 17:23:40 stack 0 left 2 right 2 idx 0 bytes 
2022/06/09 17:23:40 stack 1 left 1 right 2 idx 1 bytes (
2022/06/09 17:23:40 stack 2 left 0 right 2 idx 2 bytes ((
2022/06/09 17:23:40 stack 3 left 0 right 1 idx 3 bytes (()
2022/06/09 17:23:40 stack 4 left 0 right 0 idx 4 bytes (())
2022/06/09 17:23:40 res &[(())]
2022/06/09 17:23:40 stack 5 left 1 right 1 idx 2 bytes ()))
2022/06/09 17:23:40 stack 6 left 0 right 1 idx 3 bytes ()()
2022/06/09 17:23:40 stack 7 left 0 right 0 idx 4 bytes ()()
2022/06/09 17:23:40 res &[(()) ()()]
```

```log
2022/06/10 00:35:06 stackL 0 left 2 right 2 idx 0 bytes (
2022/06/10 00:35:06 stackL 1 left 1 right 2 idx 1 bytes ((
2022/06/10 00:35:06 stackR 2 left 0 right 2 idx 2 bytes (()
2022/06/10 00:35:06 stackR 3 left 0 right 1 idx 3 bytes (())
2022/06/10 00:35:06 stack  4 left 0 right 0 idx 4 bytes (())
2022/06/10 00:35:06 res &[(())]
2022/06/10 00:35:06 stackR 1 left 1 right 2 idx 1 bytes ()))
2022/06/10 00:35:06 stackL 5 left 1 right 1 idx 2 bytes ()()
2022/06/10 00:35:06 stackR 6 left 0 right 1 idx 3 bytes ()()
2022/06/10 00:35:06 stack  7 left 0 right 0 idx 4 bytes ()()
2022/06/10 00:35:06 res &[(()) ()()]
```

```log
2022/06/10 12:28:28 stackL 0 left 3 right 3 idx 0 bytes (
2022/06/10 12:28:28 stackL 1 left 2 right 3 idx 1 bytes ((
2022/06/10 12:28:28 stackL 2 left 1 right 3 idx 2 bytes (((
2022/06/10 12:28:28 stackR 3 left 0 right 3 idx 3 bytes ((()
2022/06/10 12:28:28 stackR 4 left 0 right 2 idx 4 bytes ((())
2022/06/10 12:28:28 stackR 5 left 0 right 1 idx 5 bytes ((()))
2022/06/10 12:28:28 stack  6 left 0 right 0 idx 6 bytes ((()))
2022/06/10 12:28:28 res &[((()))]
2022/06/10 12:28:28 stackR 2 left 1 right 3 idx 2 bytes (())))
2022/06/10 12:28:28 stackL 7 left 1 right 2 idx 3 bytes (()())
2022/06/10 12:28:28 stackR 8 left 0 right 2 idx 4 bytes (()())
2022/06/10 12:28:28 stackR 9 left 0 right 1 idx 5 bytes (()())
2022/06/10 12:28:28 stack 10 left 0 right 0 idx 6 bytes (()())
2022/06/10 12:28:28 res &[((())) (()())]
2022/06/10 12:28:28 stackR  7 left 1 right 2 idx 3 bytes (())))
2022/06/10 12:28:28 stackL 11 left 1 right 1 idx 4 bytes (())()
2022/06/10 12:28:28 stackR 12 left 0 right 1 idx 5 bytes (())()
2022/06/10 12:28:28 stack  13 left 0 right 0 idx 6 bytes (())()
2022/06/10 12:28:28 res &[((())) (()()) (())()]
2022/06/10 12:28:28 stackR  1 left 2 right 3 idx 1 bytes ()))()
2022/06/10 12:28:28 stackL 14 left 2 right 2 idx 2 bytes ()()()
2022/06/10 12:28:28 stackL 15 left 1 right 2 idx 3 bytes ()((()
2022/06/10 12:28:28 stackR 16 left 0 right 2 idx 4 bytes ()(())
2022/06/10 12:28:28 stackR 17 left 0 right 1 idx 5 bytes ()(())
2022/06/10 12:28:28 stack  18 left 0 right 0 idx 6 bytes ()(())
2022/06/10 12:28:28 res &[((())) (()()) (())() ()(())]
2022/06/10 12:28:28 stackR 15 left 1 right 2 idx 3 bytes ()()))
2022/06/10 12:28:28 stackL 19 left 1 right 1 idx 4 bytes ()()()
2022/06/10 12:28:28 stackR 20 left 0 right 1 idx 5 bytes ()()()
2022/06/10 12:28:28 stack  21 left 0 right 0 idx 6 bytes ()()()
2022/06/10 12:28:28 res &[((())) (()()) (())() ()(()) ()()()]
```

(())
()()

((()))
(()())
(())()
()(())
()()()