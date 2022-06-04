# Palindrome Number

9. Palindrome Number
Easy

5854

2138

Add to List

Share
Given an integer x, return true if x is palindrome integer.

An integer is a palindrome when it reads the same backward as forward.

For example, 121 is a palindrome while 123 is not.
 

Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
 

Constraints:

-231 <= x <= 231 - 1
 

Follow up: Could you solve it without converting the integer to a string?

## Testing and Benchmarking

* testing

```sh
go test -timeout 5s -run ^TestIsPalindrome$ github.com/fernandoocampo/justforfun/palindronum
go test -timeout 5s -run ^TestIsPalindromeLoop$ github.com/fernandoocampo/justforfun/palindronum
```

* benchmarking

```sh
go test -run none -bench IsPalindrome$ -benchtime 4s -benchmem github.com/fernandoocampo/justforfun/palindronum
go test -run none -bench IsPalindromeLoop$ -benchtime 4s -benchmem github.com/fernandoocampo/justforfun/palindronum
```