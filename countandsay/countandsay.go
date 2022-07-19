package countandsay

import (
	"bytes"
	"strconv"
)

func CountAndSay(n int) string {
	if n == 0 {
		return ""
	}
	if n == 1 {
		return "1"
	}
	result := CountAndSay(n - 1)
	if len(result) == 1 {
		return "11"
	}
	var buffer bytes.Buffer
	count := 1
	last := result[0]
	for i := 1; i < len(result); i++ {
		if result[i] == last {
			count++
			continue
		}
		buffer.WriteString(strconv.Itoa(count))
		buffer.WriteByte(last)
		last = result[i]
		count = 1
	}
	buffer.WriteString(strconv.Itoa(count))
	buffer.WriteByte(last)

	return buffer.String()
}
