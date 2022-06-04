# Longest common prefix

4. Longest Common Prefix
Easy

8119

3029

Add to List

Share
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

 

Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
 

Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lower-case English letters.

## Testing and Benchmarking

* testing

```sh
go test -timeout 2s -run ^TestLongestCommonPrefix$ github.com/fernandoocampo/justforfun/prefixes
go test -timeout 2s -run ^TestLongestCommonPrefixLoop$ github.com/fernandoocampo/justforfun/prefixes
go test -run none -bench LongestCommonPrefixLoop$ -benchtime 5s  -benchmem github.com/fernandoocampo/justforfun/prefixes
go test -run none -bench LongestCommonPrefix$ -benchtime 5s  -benchmem github.com/fernandoocampo/justforfun/prefixes
```