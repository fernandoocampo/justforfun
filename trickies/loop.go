package trickies

import (
	"fmt"
	"strconv"
	"strings"
)

type Book struct {
	Pages int
}

func LoopWithDefer(input []int) <-chan string {
	resultStream := make(chan string)

	go func() {
		defer close(resultStream)
		y := [3]*int{}

		for i, v := range input {
			newi := i
			defer func(nv int) {
				resultStream <- strconv.Itoa(nv)
			}(v)
			y[newi] = &newi
		}

		resultStream <- fmt.Sprintf("%d%d%d ", *y[0], *y[1], *y[2])
	}()

	return resultStream
}

func f(input [2]int) string {
	var builder strings.Builder

	for i, v := range input { // creates a copy of input
		if i == 0 {
			input[1] = 9
		} else {
			// here it will get 7 because keep the original array
			builder.WriteString(strconv.Itoa(v))
		}
	}
	fmt.Println("f", input) // changed input

	return builder.String()
}

func g(input [2]int) string {
	var builder strings.Builder

	for i, v := range input[:] { // works over a slice
		if i == 0 {
			input[1] = 9
		} else {
			// here it will print 9 because the
			// slice changed its internal value.
			builder.WriteString(strconv.Itoa(v))
		}
	}

	fmt.Println("g", input) // changed input

	return builder.String()
}

func LoopSliceAndArray(input [2]int) string {
	return fmt.Sprintf("%s%s", f(input), g(input))
}

func LoopBooks(books []Book) int {
	for _, book := range books {
		book.Pages = 999
	}

	return books[0].Pages
}

func LoopBooksWithPointer(books []*Book) int {
	for _, book := range books {
		book.Pages = 999
	}

	return books[0].Pages
}
