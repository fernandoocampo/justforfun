package trickies_test

import (
	"strings"
	"testing"

	"github.com/fernandoocampo/justforfun/trickies"
	"github.com/stretchr/testify/assert"
)

func TestLoopWithDefer(t *testing.T) {
	givenArray := []int{7, 8, 9}
	expectedResult := "012 987"
	var builder strings.Builder

	stream := trickies.LoopWithDefer(givenArray)

	for v := range stream {
		builder.WriteString(v)
	}

	assert.Equal(t, expectedResult, builder.String())
}

func TestLoopSliceAndArray(t *testing.T) {
	input := [2]int{5, 7}
	expectedResult := "79"

	got := trickies.LoopSliceAndArray(input)

	assert.Equal(t, expectedResult, got)
}

func TestLoopBooks(t *testing.T) {
	input := []trickies.Book{{555}}
	expectedResult := 555

	got := trickies.LoopBooks(input)

	assert.Equal(t, expectedResult, got)
}

func TestLoopBooksWithPointer(t *testing.T) {
	input := []*trickies.Book{{555}}
	expectedResult := 999

	got := trickies.LoopBooksWithPointer(input)

	assert.Equal(t, expectedResult, got)
}
