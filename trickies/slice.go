package trickies

import (
	"fmt"
	"strconv"
	"strings"
)

// UpdateArray takes some items from given array
// and modifies its values.
// let's suppose we have this array: [4]int{0,1,2,3}
// First it creates a slice x with the item at values[0]
// x := values[:1] --> []int{0}
// Next it creates a slice y with the items from pos 2 to the end of values.
// y := values[2:] --> []int{2,3}
// Then it appends y to x
// println(x) --> {0,2,3}
// so, values should be now: [4]{0,2,3,3}
// and it turns out that y is different now
// because it points to values
// println(y) --> {3,3}
// After that it appends y to x again.
// println(x) --> {0,2,3,3,3}
// so, values should be in the end: [4]{0,2,3,3}
func UpdateArray(values [4]int) ([4]int, []int) {
	x := values[:1]
	y := values[2:]
	x = append(x, y...)
	x = append(x, y...)
	return values, x
}

// PrintFuzzyStringArray iterates the given array and get
// the items to create a new string.
// The trick here is that the for range takes a copy
// of the given slice.
// Let's suppose that the function receives this array:
// values = []{"A", "B", "C"}
// the copy of the for range will be: []{"A", "B", "C"}
// despite of the for loop append more items to values
// the values from the for range will see the first
// 3 positions.
// You see that there is append, that append will make
// the slice to be different from the one we have in
// the for range loop.
func PrintFuzzyStringArray(values []string) string {
	var builder strings.Builder

	for i, s := range values {
		builder.WriteString(strconv.Itoa(i))
		builder.WriteString(s)
		builder.WriteString(",")
		values[i+1] = "M"
		values = append(values, "Z")
		values[i+1] = "Z"
	}

	fmt.Println(values)
	return builder.String()
}
