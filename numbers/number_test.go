package numbers_test

import (
	"slices"
	"testing"

	"github.com/fernandoocampo/justforfun/numbers"
)

func TestFindSmallestNumber(t *testing.T) {
	t.Parallel()
	// Given
	cases := map[string]struct {
		given []int
		want  int
	}{
		"a": {
			given: []int{10, 8, 7, 6, 100, 13, 23},
			want:  6,
		},
	}

	for testName, testData := range cases {
		t.Run(testName, func(st *testing.T) {
			testData := testData
			st.Parallel()
			// when
			got := numbers.FindSmallestNumber(testData.given)
			// Then
			if testData.want != got {
				st.Errorf("want: %d, but got: %d", testData.want, got)
			}
		})
	}
}

func TestGenerateFibon(t *testing.T) {
	t.Parallel()
	// Given
	cases := map[string]struct {
		numbers int
		want    []int
	}{
		"011235": {
			numbers: 6,
			want:    []int{0, 1, 1, 2, 3, 5},
		},
		"0112358": {
			numbers: 7,
			want:    []int{0, 1, 1, 2, 3, 5, 8},
		},
	}

	for testName, testData := range cases {
		t.Run(testName, func(st *testing.T) {
			// When
			got := numbers.MakeFibos(testData.numbers)
			// Then
			if !slices.Equal(testData.want, got) {
				t.Errorf("want: %+v, but got: %+v", testData.want, got)
			}
		})
	}
}

// go test -run=^$ -benchmem -benchtime 3s -bench ^BenchmarkFindSmallestNumber$ github.com/fernandoocampo/justforfun/numbers
// goos: darwin
// goarch: amd64
// pkg: github.com/fernandoocampo/justforfun/numbers
// cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
// BenchmarkFindSmallestNumber-16           9532549               374.9 ns/op            56 B/op          2 allocs/op
// PASS
// ok      github.com/fernandoocampo/justforfun/numbers    4.401s
func BenchmarkFindSmallestNumber(b *testing.B) {
	data := []int{10, 8, 7, 6, 100, 13, 23, 11, 34, 22, 5, 233, 12, 93, 24, 78, 66, 33}
	input := make([]int, len(data))
	for i := 0; i < b.N; i++ {
		copy(input, data)
		_ = numbers.FindSmallestNumber(input)
	}
}

// go test -run=^$ -benchmem -benchtime 3s -bench ^BenchmarkSmallestIntegerFinder$ github.com/fernandoocampo/justforfun/numbers
// goos: darwin
// goarch: amd64
// pkg: github.com/fernandoocampo/justforfun/numbers
// cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
// BenchmarkSmallestIntegerFinder-16       11381394               324.9 ns/op            24 B/op          1 allocs/op
// PASS
// ok      github.com/fernandoocampo/justforfun/numbers    4.276s
func BenchmarkSmallestIntegerFinder(b *testing.B) {
	data := []int{10, 8, 7, 6, 100, 13, 23, 11, 34, 22, 5, 233, 12, 93, 24, 78, 66, 33}
	input := make([]int, len(data))
	for i := 0; i < b.N; i++ {
		copy(input, data)
		_ = numbers.SmallestIntegerFinder(input)
	}
}

// Winner
// go test -run=^$ -benchmem -benchtime 3s -bench ^BenchmarkSmallestIntegerFinderFor$ github.com/fernandoocampo/justforfun/numbers
// goos: darwin
// goarch: amd64
// pkg: github.com/fernandoocampo/justforfun/numbers
// cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
// BenchmarkSmallestIntegerFinderFor-16            251742655               14.13 ns/op            0 B/op          0 allocs/op
// PASS
// ok      github.com/fernandoocampo/justforfun/numbers    5.175s
func BenchmarkSmallestIntegerFinderFor(b *testing.B) {
	data := []int{10, 8, 7, 6, 100, 13, 23, 11, 34, 22, 5, 233, 12, 93, 24, 78, 66, 33}
	input := make([]int, len(data))
	for i := 0; i < b.N; i++ {
		copy(input, data)
		_ = numbers.SmallestIntegerFinderFor(input)
	}
}

// go test -run=^$ -benchmem -benchtime 3s -bench ^BenchmarkSmallestIntegerFinderForTwo$ github.com/fernandoocampo/justforfun/numbers
// goos: darwin
// goarch: amd64
// pkg: github.com/fernandoocampo/justforfun/numbers
// cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
// BenchmarkSmallestIntegerFinderForTwo-16         232770330               15.59 ns/op            0 B/op          0 allocs/op
// PASS
// ok      github.com/fernandoocampo/justforfun/numbers    5.443s
func BenchmarkSmallestIntegerFinderForTwo(b *testing.B) {
	data := []int{10, 8, 7, 6, 100, 13, 23, 11, 34, 22, 5, 233, 12, 93, 24, 78, 66, 33}
	input := make([]int, len(data))
	for i := 0; i < b.N; i++ {
		copy(input, data)
		_ = numbers.SmallestIntegerFinderForTwo(input)
	}
}
