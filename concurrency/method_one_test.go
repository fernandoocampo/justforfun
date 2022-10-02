package concurrency_test

import (
	"sort"
	"testing"

	"github.com/fernandoocampo/justforfun/concurrency"
	"github.com/stretchr/testify/assert"
)

func TestProcessDataOne(t *testing.T) {
	// Given
	values := generateValuesOne(10)
	expectedResult := []concurrency.Result{
		{1, false}, {2, true}, {3, false}, {4, true}, {5, false},
		{6, true}, {7, false}, {8, true}, {9, false}, {10, true},
	}

	// When
	got := concurrency.ProcessDataOne(values)

	// Then
	sort.Slice(got, func(i, j int) bool {
		return got[i].Value < got[j].Value
	})
	assert.Equal(t, expectedResult, got)
}

func generateValuesOne(size int) []concurrency.Data {
	result := make([]concurrency.Data, size)

	for i := 0; i < size; i++ {
		result[i] = concurrency.Data{
			Value: i + 1,
		}
	}

	return result
}
