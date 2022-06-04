# reverse integers
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

 

Example 1:

Input: x = 123
Output: 321
Example 2:

Input: x = -123
Output: -321
Example 3:

Input: x = 120
Output: 21
 

Constraints:

-231 <= x <= 231 - 1

## Testing and Benchmarking

* testing

```sh
go test -count=1 -timeout 5s -run ^TestReverse$ github.com/fernandoocampo/justforfun/reverseints
go test -count=1 -timeout 5s -run ^TestReverseLoop$ github.com/fernandoocampo/justforfun/reverseints
```

* benchmark

```sh
go test -run none -bench Reverse$ -benchtime 4s -benchmem github.com/fernandoocampo/justforfun/reverseints
go test -run none -bench Reverse$ -benchtime 4s -benchmem -memprofile mem.out github.com/fernandoocampo/justforfun/reverseints
go tool pprof -alloc_space reverseints.test mem.out


go test -run none -bench ReverseLoop$ -benchtime 4s -benchmem github.com/fernandoocampo/justforfun/reverseints
```