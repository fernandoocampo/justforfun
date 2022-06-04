package median_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/median"
)

func TestFindMedianSortedArrays(t *testing.T) {
	cases := map[string]struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		"2.00000": {
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.00000,
		},
		"2.50000": {
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.50000,
		},
		"2.50000_2": {
			nums1: []int{1, 3},
			nums2: []int{2, 7},
			want:  2.50000,
		},
		"nums1_empty": {
			nums1: []int{},
			nums2: []int{1},
			want:  1.00000,
		},
	}

	for name, testdata := range cases {
		t.Run(name, func(st *testing.T) {
			got := median.FindMedianSortedArrays(testdata.nums1, testdata.nums2)
			if got != testdata.want {
				st.Errorf("want: %f, but got: %f", testdata.want, got)
			}
		})
	}
}

func TestFindMedianSortedArraysLoop(t *testing.T) {
	cases := map[string]struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		"2.00000": {
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.00000,
		},
		"2.50000": {
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.50000,
		},
		"2.50000_2": {
			nums1: []int{1, 3},
			nums2: []int{2, 7},
			want:  2.50000,
		},
	}

	for name, testdata := range cases {
		t.Run(name, func(st *testing.T) {
			got := median.FindMedianSortedArraysLoop(testdata.nums1, testdata.nums2)
			if got != testdata.want {
				st.Errorf("want: %f, but got: %f", testdata.want, got)
			}
		})
	}
}

func BenchmarkFindMedianSortedArrays(b *testing.B) {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{6, 7, 8, 9, 10}

	for i := 0; i < b.N; i++ {
		median.FindMedianSortedArrays(nums1, nums2)
	}
}

func BenchmarkFindMedianSortedArrays2(b *testing.B) {
	nums2 := []int{1, 2, 3, 4, 5}
	nums1 := []int{6, 7, 8, 9, 10}

	for i := 0; i < b.N; i++ {
		median.FindMedianSortedArrays(nums1, nums2)
	}
}

func BenchmarkFindMedianSortedArrays3(b *testing.B) {
	nums2 := []int{1, 3, 5, 7, 9}
	nums1 := []int{2, 4, 6, 8, 10}

	for i := 0; i < b.N; i++ {
		median.FindMedianSortedArrays(nums1, nums2)
	}
}

func aBenchmarkFindMedianSortedArraysLoop(b *testing.B) {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{6, 7, 8, 9, 10}

	for i := 0; i < b.N; i++ {
		median.FindMedianSortedArraysLoop(nums1, nums2)
	}
}
