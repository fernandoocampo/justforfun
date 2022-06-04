https://leetcode.com/problems/longest-substring-without-repeating-characters/
3. Longest Substring Without Repeating Characters
Medium

23642

1060

Add to List

Share
Given a string s, find the length of the longest substring without repeating characters.

 

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
 

Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.


## Running benchmark

```sh
go test -run none -bench LengthOfLongestSubstring -benchtime 3s -benchmem ./longsubstr
go test ./longsubstr -run none -bench LengthOfLongestSubstring -benchtime 3s -benchmem
go test ./longsubstr -run none -bench LengthOfLongestSubstring -benchtime 3s -benchmem -memprofile mem.out
go tool pprof -alloc_space longsubstr.test mem.out
list LengthOfLongestSubstring
rm mem.out longsubstr.test

go test ./longsubstr -run none -bench LengthOfLongestSubstringRec -benchtime 3s -benchmem
go test ./longsubstr -run none -bench LengthOfLongestSubstringRec -benchtime 3s -benchmem -memprofile mem.out

```