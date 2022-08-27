package trickies_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/trickies"
	"github.com/stretchr/testify/assert"
)

func TestUpdatingArray(t *testing.T) {
	given := [...]int{0, 1, 2, 3}
	expectedX := []int{0, 2, 3, 3, 3}
	expectedValues := [4]int{0, 2, 3, 3}

	got, x := trickies.UpdateArray(given)

	assert.Equal(t, expectedX, x)
	assert.Equal(t, expectedValues, got)
}

func TestStringArray(t *testing.T) {
	given := []string{"A", "B", "C"}
	expectedResult := "0A,1M,2C,"

	got := trickies.PrintFuzzyStringArray(given)

	assert.Equal(t, expectedResult, got)
}
