package zigzag

import (
	"strings"
)

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	b := make([]strings.Builder, numRows)
	return organize(s, b, numRows, 0, 1)
}

func organize(s string, b []strings.Builder, numRows, idx, sum int) string {
	if len(s) == 0 {
		for i := 1; i < numRows; i++ {
			b[0].WriteString(b[i].String())
		}
		return b[0].String()
	}
	b[idx].WriteByte(s[0])
	if idx == numRows-1 && sum > 0 {
		sum = -1
	}
	if idx == 0 && sum < 0 {
		sum = 1
	}
	return organize(s[1:], b, numRows, idx+sum, sum)
}

func ConvertLoop(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	b := make([]strings.Builder, numRows)
	var i int
	add := 1
	for _, v := range s {
		b[i].WriteRune(v)
		if i == numRows-1 && add == 1 {
			add = -1
		}
		if i == 0 && add == -1 {
			add = 1
		}
		i = i + add
	}
	for i := 1; i < numRows; i++ {
		b[0].WriteString(b[i].String())
	}
	return b[0].String()
}
