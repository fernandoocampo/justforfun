# longest palindromic substring

Given a string s, return the longest palindromic substring in s.

 

Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"
 

Constraints:

1 <= s.length <= 1000
s consist of only digits and English letters.
Accepted
1,834,156
Submissions
5,766,560

## Testing and benchmarks

* testing

```sh
go test -timeout 5s -run ^TestLongestPalindrome$ github.com/fernandoocampo/justforfun/palindromics
go test -run none -bench LongestPalindrome$ -benchtime 3s github.com/fernandoocampo/justforfun/palindromics -benchmem
go test -run none -bench LongestPalindrome$ -benchtime 3s github.com/fernandoocampo/justforfun/palindromics -benchmem -memprofile mem.out
go tool pprof -alloc_space palindromics.test mem.out

go test -timeout 5s -run ^TestLongestPalindromeLoop$ github.com/fernandoocampo/justforfun/palindromics
go test -run none -bench LongestPalindromeLoop$ -benchtime 3s github.com/fernandoocampo/justforfun/palindromics
go test -run none -bench LongestPalindromeLoop$ -benchtime 3s github.com/fernandoocampo/justforfun/palindromics -benchmem -memprofile mem.out

```