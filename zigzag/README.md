# zigzag conversion

The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
 

Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I
Example 3:

Input: s = "A", numRows = 1
Output: "A"

Example 4:

Input: s = "PAYPALISHIRING", numRows = 2
Output: "PYAIHRNAPLSIIG"

P Y A I H R N
A P L S I I G
 

Constraints:

1 <= s.length <= 1000
s consists of English letters (lower-case and upper-case), ',' and '.'.
1 <= numRows <= 1000

## Testing and Benchmarking

* Test

```sh
go test -timeout 5s -run ^TestConvert$ github.com/fernandoocampo/justforfun/zigzag
go test -timeout 5s -run ^TestConvertLoop$ github.com/fernandoocampo/justforfun/zigzag
```

* BenchMark

```sh
go test -run none -bench Convert$ -benchtime 3s -benchmem github.com/fernandoocampo/justforfun/zigzag
go test -run none -bench Convert$ -benchtime 3s -benchmem -memprofile mem.out github.com/fernandoocampo/justforfun/zigzag
go tool pprof -alloc_space zigzag.test mem.out

go test -run none -bench ConvertLoop$ -benchtime 3s -benchmem github.com/fernandoocampo/justforfun/zigzag
go test -run none -bench ConvertLoop$ -benchtime 3s -benchmem -memprofile mem.out github.com/fernandoocampo/justforfun/zigzag
```