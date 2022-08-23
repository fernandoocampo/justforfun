package reversers

import (
	"math"
	"strconv"
	"strings"
)

func Reverse(input int) int {
	if input == 0 {
		return input
	}
	var negative bool
	if input < 0 {
		negative = true
		input = input * -1
	}
	number := strconv.Itoa(input)
	result := reverse(number, &strings.Builder{})
	if negative {
		result = result * -1
	}
	return result
}

func reverse(n string, b *strings.Builder) int {
	if len(n) == 1 {
		b.WriteString(n)
		v, err := strconv.Atoi(b.String())
		if err != nil {
			return 0
		}
		return v
	}
	b.WriteByte(n[len(n)-1])
	return reverse(n[:len(n)-1], b)
}

func ReverseLoop(input int) int {
	number := strconv.Itoa(input)
	if input < 0 {
		number = number[1:]
	}
	b := strings.Builder{}
	for i := len(number) - 1; i > -1; i-- {
		b.WriteByte(number[i])
	}
	result, err := strconv.Atoi(b.String())
	if err != nil {
		return 0
	}
	if input < 0 {
		return result * -1
	}
	return result
}

func ReverseLoopInt32(input int) int {
	number := strconv.Itoa(input)
	if input < 0 {
		number = number[1:]
	}
	b := strings.Builder{}
	for i := len(number) - 1; i > -1; i-- {
		b.WriteByte(number[i])
	}
	result, err := strconv.Atoi(b.String())
	if err != nil {
		return 0
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	if input < 0 {
		return result * -1
	}
	return result
}
